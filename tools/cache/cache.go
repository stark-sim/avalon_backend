package cache

import (
	"context"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
)

const DefaultClient = "REDIS"

type Client interface {
	GetUser(ctx context.Context, userID string) (*model.User, error)
	SetUser(ctx context.Context, user *model.User) error
}
