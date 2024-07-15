package resolvers

import (
	"github.com/ch3yb/clinic/api/service"
	"github.com/ch3yb/clinic/graph"
	"sync"
)

type Resolver struct {
	Service *service.Service
	mu      sync.Mutex
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
