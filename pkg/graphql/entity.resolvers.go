package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: FindUserByID - findUserByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
