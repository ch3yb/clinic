package resolvers

import (
	"context"
)

func (r *mutationResolver) MutationTest(ctx context.Context) (string, error) {
	return "hello", nil

}
func (r *queryResolver) QueryTest(ctx context.Context) (string, error) {
	return "hello", nil

}
