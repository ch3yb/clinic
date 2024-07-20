package service

import (
	"github.com/ch3yb/clinic/graph/models"
	"github.com/mattn/go-sqlite3"
	"time"
)

func (s *Service) GetPatient(patientID uint) (*models.Patient, error) {
	var patient = new(models.Patient)
	err := s.db.QueryRow(`SELECT patient_id, notes, blood_type, emergency_contact_name, emergency_contact_phone, insurance_provider, insurance_policy_number, created_at, first_name, last_name, email, phone_number, address, date_of_birth, gender, status, profile_picture FROM patients where patient_id = $1`, patientID).Scan(&patient.PatientID, &patient.Notes, &patient.BloodType, &patient.EmergencyContactName, &patient.EmergencyContactPhone, &patient.InsuranceProvider, &patient.InsurancePolicyNumber, &patient.CreatedAt, &patient.FirstName, &patient.LastName, &patient.Email, &patient.PhoneNumber, &patient.Address, &patient.DateOfBirth, &patient.Gender, &patient.Status, &patient.ProfilePicture)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	return patient, nil
}

func (s *Service) InsertPatient(patient *models.PatientInput) error {
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

	_, err := s.db.Exec(query,
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

func (s *Service) UpdatePatient(patient models.Patient) error {
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
		patient.PatientID,
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
