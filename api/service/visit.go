package service

import (
	"database/sql"
	"github.com/ch3yb/clinic/api/errors"
	"github.com/ch3yb/clinic/graph/models"
	"time"
)

func (s *Service) CreateVisit(in *models.VisitInput) error {

	var query = `
	   INSERT INTO visits (
patient_id, details, diagnosis, notes, visit_date
	   ) VALUES ($1,$2,$3,$4,$5)
	`

	_, err := s.db.Exec(query,
		in.PatientID,
		in.Details,
		in.Diagnosis,
		in.Notes,
		time.Now().Unix(),
	)

	if err != nil {
		s.Logger.Error(err.Error())
		return s.Err.ErrInternal()
	}
	return nil
}

func (s *Service) GetVisit(visitID uint) (*models.Visit, error) {
	var visit = new(models.Visit)
	var prescriptionID sql.NullInt32

	err := s.db.QueryRow(`SELECT 
patient_id, prescription_id, details, diagnosis, notes, visit_date FROM visits WHERE visit_id = $1`, visitID).Scan(
		&visit.PatientID,
		&prescriptionID,
		&visit.Details,
		&visit.Diagnosis,
		&visit.Notes,
		&visit.VisitDate,
	)
	if err != nil {
		s.Logger.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, s.Err.ErrNotFound()
		}
		return nil, s.Err.ErrInternal()
	}

	if prescriptionID.Valid {
		// todo get prescrption ...
		visit.VisitPrescription, err = s.GetPrescription(uint32(prescriptionID.Int32))
		visit.VisitPrescription = nil
	}
	return visit, nil
}

func (s *Service) GetVisits(patientID uint) ([]*models.Visit, error) {
	var visits = make([]*models.Visit, 0)

	rows, err := s.db.Query(`SELECT visit_id,
	patient_id, prescription_id, details, diagnosis, notes, visit_date FROM visits WHERE patient_id = $1`, patientID)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, s.Err.GetErrorMessage(errors.ErrNotFound)
	}

	for rows.Next() {
		var visit = new(models.Visit)
		var prescriptionID sql.NullInt32
		rows.Scan(
			&visit.VisitID,
			&visit.PatientID,
			&prescriptionID,
			&visit.Details,
			&visit.Diagnosis,
			&visit.Notes,
			&visit.VisitDate,
		)

		if prescriptionID.Valid {
			visit.VisitPrescription, err = s.GetPrescription(uint32(visit.VisitID))
			if err != nil {
				s.Logger.Error(err.Error())
				return nil, err
			}
		}

		if err != nil {
			s.Logger.Error(err.Error())
			if err == sql.ErrNoRows {
				return nil, s.Err.GetErrorMessage(errors.ErrNotFound)
			}
			return nil, err
		}

		visits = append(visits, visit)
	}

	return visits, nil
}
