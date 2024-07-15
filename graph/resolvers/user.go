package resolvers

import (
	"context"
	"github.com/ch3yb/clinic/graph/models"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {

	r.Service.TestFn()
	var todos = make([]*models.Todo, 0)
	todos = append(todos, &models.Todo{
		ID:   "tes",
		Text: "",
		Done: false,
		User: nil,
	})
	return todos, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, todo models.NewTodo) (*models.Todo, error) {
	return nil, nil
}
