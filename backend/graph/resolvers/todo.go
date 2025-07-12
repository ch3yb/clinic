package resolvers

import (
	"context"
	"errors"
	"fmt"
	"github.com/ch3yb/hanotiy/graph/models"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var gg = make([]*models.Todo, 0)
	fmt.Println("todos")
	return gg, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, in models.NewTodo) (*models.Todo, error) {
	return nil, errors.New("hi")
}
