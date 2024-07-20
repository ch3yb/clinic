package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/errors"
	"github.com/ch3yb/clinic/graph/models"
)

func (r *queryResolver) GetPatient(ctx context.Context, patientID int) (*models.Patient, error) {

	patient, err := r.Service.GetPatient(uint(patientID))
	if err != nil {
		return nil, errors.ErrInternal
	}
	return patient, nil
}

func (r *mutationResolver) CreatePatient(ctx context.Context, in models.PatientInput) (bool, error) {
	err := r.Service.InsertPatient(&in)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (r *mutationResolver) UpdatePatient(ctx context.Context, patientID int, in models.PatientInput) (bool, error) {
	err := r.Service.UpdatePatient(patientID, in)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (r *mutationResolver) DeletePatient(ctx context.Context, patientID int) (bool, error) {
	err := r.Service.DeletePatient(patientID)
	if err != nil {
		return false, err
	}
	return true, nil

}
