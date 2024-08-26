package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/graph/models"
)

func (r *mutationResolver) CreatePrescription(ctx context.Context, in models.PrescriptionInput) (bool, error) {
	err := r.Service.CreatePrescription(&in)
	if err != nil {
		return false, err
	}
	return true, nil
}
