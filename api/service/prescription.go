package service

import (
	"database/sql"
	"github.com/ch3yb/clinic/graph/models"
	"time"
)

func (s *Service) CreatePrescription(in *models.PrescriptionInput) error {

	var prescriptionID int32
	err := s.db.QueryRow(`
	   INSERT INTO prescriptions( visit_id,instructions, created_at) VALUES ($1,$2,$3) RETURNING prescription_id
	`,
		in.VisitID,
		in.Instructions,
		time.Now().Unix(),
	).Scan(&prescriptionID)

	err = s.setItems(prescriptionID, in.Items)
	if err != nil {
		s.Logger.Error(err.Error())
		return s.Err.ErrInternal()
	}
	return nil
}

func (s *Service) GetPrescription(visitID uint32) (*models.Prescription, error) {
	var prescription = new(models.Prescription)
	err := s.db.QueryRow(`SELECT prescription_id,instructions,created_at FROM  prescriptions WHERE visit_id = $1`, visitID).Scan(
		&prescription.PrescriptionID,
		&prescription.Instructions,
		&prescription.CreatedAt,
	)
	if err != nil {
		s.Logger.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, s.Err.ErrNotFound()
		}
		return nil, s.Err.ErrInternal()
	}

	prescription.Items, err = s.getItems(prescription.PrescriptionID)
	if err != nil {
		return nil, err
	}

	return prescription, nil
}
func (s *Service) setItems(prescriptionID int32, items []*models.PrescriptionItemInput) error {

	for _, item := range items {

		var med = new(models.Medicament)
		err := s.db.QueryRow(`SELECT id, code, denomination_commune_internationale, nom_de_marque, dosage FROM medicament WHERE id = $1`, item.ID).Scan(
			&med.Code,
			&med.DenominationCommuneInternationale,
			&med.NomDeMarque,
			&med.Dosage,
		)
		if err != nil {
			s.Logger.Error(err.Error())
		}

		_, err = s.db.Exec(`insert into prescription_items (prescription_id, medicament_id,medication_name,dosage,duration,frequency) VALUES ($1,$2,$3,$4,$5,$k6)`, prescriptionID, item.ID, med.DenominationCommuneInternationale, med.Dosage, item.Duration, item.Frequency)
		if err != nil {
			s.Logger.Error(err.Error())
			return err
		}
	}

	return nil
}
func (s *Service) getItems(id int) ([]*models.PrescriptionItem, error) {

	rows, err := s.db.Query(`SELECT id,medicament_id,medication_name, dosage, frequency, duration FROM prescription_items WHERE prescription_id = $1`, id)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, s.Err.ErrInternal()
	}
	var items = make([]*models.PrescriptionItem, 0)
	for rows.Next() {
		var item = new(models.PrescriptionItem)
		err = rows.Scan(&item.ID, &item.MedicationName, &item.Dosage, &item.Frequency, &item.Duration)
		if err != nil {
			s.Logger.Error(err.Error())
			return nil, s.Err.ErrInternal()
		}
		items = append(items, item)
	}

	return items, nil
}
