package service

import (
	"database/sql"
	"github.com/ch3yb/clinic/api/errors"
	"github.com/ch3yb/clinic/graph/models"
	"github.com/ch3yb/clinic/utils"
	"github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"time"
)

func (s *Service) GetPatient(patientID uint) (*models.Patient, error) {
	var patient = new(models.Patient)
	var createdAt, dob time.Time
	var notes, ec, ecp, ip, ipn, email, pp, address sql.NullString

	err := s.db.QueryRow(`SELECT patient_id, notes, blood_type, emergency_contact_name, emergency_contact_phone, insurance_provider, insurance_policy_number, created_at, first_name, last_name, email, phone_number, address, date_of_birth, gender, status, profile_picture FROM patients WHERE patient_id = $1`, patientID).Scan(
		&patient.PatientID,
		&notes, &patient.BloodType,
		&ec, &ecp, &ip, &ipn,
		&createdAt,
		&patient.FirstName, &patient.LastName,
		&email, &patient.PhoneNumber,
		&address, &dob,
		&patient.Gender, &patient.Status,
		&pp,
	)
	if err != nil {
		s.Logger.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, s.Err.ErrNotFound()
		}
		return nil, s.Err.ErrInternal()
	}
	// Assign values from sql.NullString if they are valid
	if notes.Valid {
		patient.Notes = utils.GetStringPointer(notes.String)
	}
	if ec.Valid {
		patient.EmergencyContactName = utils.GetStringPointer(ec.String)
	}
	if ecp.Valid {
		patient.EmergencyContactPhone = utils.GetStringPointer(ecp.String)
	}
	if ip.Valid {
		patient.InsuranceProvider = utils.GetStringPointer(ip.String)
	}
	if ipn.Valid {
		patient.InsurancePolicyNumber = utils.GetStringPointer(ipn.String)
	}
	if email.Valid {
		patient.Email = utils.GetStringPointer(email.String)
	}
	if pp.Valid {
		patient.ProfilePicture = utils.GetStringPointer(pp.String)
	}

	if address.Valid {
		patient.Address = utils.GetStringPointer(address.String)
	}

	// Handle times
	patient.CreatedAt = int(createdAt.Unix())
	patient.DateOfBirth = int(dob.Unix())

	return patient, nil
}
func (s *Service) GetPatients(archived *bool) ([]*models.Patient, error) {
	var patients = make([]*models.Patient, 0)
	var query = `SELECT patient_id, notes, blood_type, emergency_contact_name, emergency_contact_phone, insurance_provider, insurance_policy_number, created_at, first_name, last_name, email, phone_number, address, date_of_birth, gender, status, profile_picture FROM patients WHERE deleted_at IS NULL;`
	if archived != nil && *archived {
		query = `SELECT patient_id, notes, blood_type, emergency_contact_name, emergency_contact_phone, insurance_provider, insurance_policy_number, created_at, first_name, last_name, email, phone_number, address, date_of_birth, gender, status, profile_picture FROM patients WHERE deleted_at IS NOT NULL;`
	}

	rows, err := s.db.Query(query)

	if err != nil {
		s.Logger.Error(err.Error())
		return nil, s.Err.GetErrorMessage(errors.ErrNotFound)
	}

	s.Logger.Info("", zap.Any("", rows))

	for rows.Next() {
		var patient = new(models.Patient)
		var createdAt, dob time.Time
		var notes, ec, ecp, ip, ipn, email, pp, address sql.NullString
		err = rows.Scan(
			&patient.PatientID,
			&notes, &patient.BloodType,
			&ec, &ecp, &ip, &ipn,
			&createdAt,
			&patient.FirstName, &patient.LastName,
			&email, &patient.PhoneNumber,
			&address, &dob,
			&patient.Gender, &patient.Status,
			&pp,
		)
		if err != nil {
			s.Logger.Error(err.Error())
			if err == sql.ErrNoRows {
				return nil, s.Err.GetErrorMessage(errors.ErrNotFound)
			}
			return nil, err
		}

		if notes.Valid {
			patient.Notes = utils.GetStringPointer(notes.String)
		}
		if ec.Valid {
			patient.EmergencyContactName = utils.GetStringPointer(ec.String)
		}
		if ecp.Valid {
			patient.EmergencyContactPhone = utils.GetStringPointer(ecp.String)
		}
		if ip.Valid {
			patient.InsuranceProvider = utils.GetStringPointer(ip.String)
		}
		if ipn.Valid {
			patient.InsurancePolicyNumber = utils.GetStringPointer(ipn.String)
		}
		if email.Valid {
			patient.Email = utils.GetStringPointer(email.String)
		}
		if pp.Valid {
			patient.ProfilePicture = utils.GetStringPointer(pp.String)
		}

		if address.Valid {
			patient.Address = utils.GetStringPointer(address.String)
		}

		// Handle times
		patient.CreatedAt = int(createdAt.Unix())
		patient.DateOfBirth = int(dob.Unix())

		s.Logger.Info("", zap.Any("", patient))

		patients = append(patients, patient)
	}

	return patients, nil
}
func (s *Service) CreatePatient(patient *models.PatientInput) error {

	var exist bool
	err := s.db.QueryRow(`SELECT TRUE FROM patients WHERE first_name = $1 AND last_name = $2;`, patient.FirstName, patient.LastName).Scan(&exist)
	if err != nil {
		if err != sql.ErrNoRows {
			s.Logger.Error(err.Error())
			return s.Err.GetErrorMessage(errors.ErrInternal)
		}
	}
	if exist {
		return s.Err.GetErrorMessage(errors.ErrExist)
	}
	query := `
        INSERT INTO patients (
            notes,
            blood_type,
            emergency_contact_name,
            emergency_contact_phone,
            insurance_provider,
            insurance_policy_number,
            created_at,
            updated_at,
            first_name,
            last_name,
            email,
            phone_number,
            address,
            date_of_birth,
            gender,
            profile_picture
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err = s.db.Exec(query,
		patient.Notes,
		patient.BloodType,
		patient.EmergencyContactName,
		patient.EmergencyContactPhone,
		patient.InsuranceProvider,
		patient.InsurancePolicyNumber,
		time.Now().Unix(),
		time.Now().Unix(),
		patient.FirstName,
		patient.LastName,
		patient.Email,
		patient.PhoneNumber,
		patient.Address,
		patient.DateOfBirth,
		patient.Gender,
		patient.ProfilePicture,
	)

	if err != nil {
		s.Logger.Error(err.Error())
		return sqlite3.ErrInternal
	}

	return nil
}

func (s *Service) UpdatePatient(patientID int, patient models.PatientInput) error {
	query := `
   UPDATE patients
SET
    blood_type = $2,
    emergency_contact_name = $3,
    emergency_contact_phone = $4,
    insurance_provider = $5,
    insurance_policy_number = $6,
    created_at = $7,
    updated_at = $8,
    first_name = $9,
    last_name = $10,
    email = $11,
    phone_number = $12,
    address = $13,
    date_of_birth = $14,
    gender = $15,
    profile_picture = $16,
    notes = 17
WHERE
    patient_id = $1;
`

	_, err := s.db.Exec(query,
		patientID,
		patient.Notes,
		patient.BloodType,
		patient.EmergencyContactName,
		patient.EmergencyContactPhone,
		patient.InsuranceProvider,
		patient.InsurancePolicyNumber,
		time.Now(),
		time.Now(),
		patient.FirstName,
		patient.LastName,
		patient.Email,
		patient.PhoneNumber,
		patient.Address,
		patient.DateOfBirth,
		patient.Gender,
		patient.ProfilePicture,
	)

	if err != nil {
		s.Logger.Error(err.Error())
		return sqlite3.ErrInternal
	}

	return nil
}

func (s *Service) DeletePatient(patientID int) error {
	_, err := s.db.Exec(`UPDATE patients SET deleted_at = CURRENT_TIMESTAMP WHERE patient_id = $1`, patientID)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	return nil
}
func (s *Service) RestorePatient(patientID int) error {
	_, err := s.db.Exec(`UPDATE patients SET deleted_at = NULL WHERE patient_id = $1`, patientID)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	return nil
}
