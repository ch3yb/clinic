package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/graph/models"
	"log"
)

func (r *queryResolver) GetPatient(ctx context.Context, patientID int) (*models.Patient, error) {

	patient, err := r.Service.GetPatient(uint(patientID))
	if err != nil {
		return nil, err
	}
	log.Println(patient)

	return patient, nil

}

func (r *mutationResolver) CreatePatient(ctx context.Context, in models.PatientInput) (bool, error) {
	err := r.Service.InsertPatient(&in)
	if err != nil {
		return false, err
	}
	return true, nil

}
func (r *mutationResolver) UpdatePatient(ctx context.Context, in models.Patient) (bool, error) {
	err := r.Service.UpdatePatient(in)
	if err != nil {
		return false, err
	}
	return true, nil

}
