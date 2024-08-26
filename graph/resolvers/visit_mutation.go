package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/graph/models"
)

func (r *mutationResolver) CreateVisit(ctx context.Context, in models.VisitInput) (bool, error) {
	err := r.Service.CreateVisit(&in)
	if err != nil {
		return false, err
	}
	return true, nil
}
