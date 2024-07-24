package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/graph/models"
)

func (r *queryResolver) Patient(ctx context.Context, patientID int) (*models.Patient, error) {
	patient, err := r.Service.GetPatient(uint(patientID))
	if err != nil {
		return nil, err
	}
	return patient, nil
}
func (r *queryResolver) Patients(ctx context.Context, archived *bool) ([]*models.Patient, error) {
	patients, err := r.Service.GetPatients(archived)
	if err != nil {
		return nil, err
	}
	return patients, nil
}
