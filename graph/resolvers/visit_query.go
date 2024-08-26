package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/graph/models"
)

func (r *queryResolver) Visit(ctx context.Context, visitID int) (*models.Visit, error) {
	visit, err := r.Service.GetVisit(uint(visitID))
	if err != nil {
		return nil, err
	}
	return visit, nil
}
func (r *queryResolver) Visits(ctx context.Context, patientID int) ([]*models.Visit, error) {
	visits, err := r.Service.GetVisits(uint(patientID))
	if err != nil {
		return nil, err
	}
	return visits, nil
}
