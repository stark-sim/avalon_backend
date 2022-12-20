package middlewares

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/stark-sim/avalon_backend/tools/cache"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type SubscriptionCacheClientLoader struct{}

func (SubscriptionCacheClientLoader) ExtensionName() string {
	return "SubscriptionCacheClientLoader"
}

func (SubscriptionCacheClientLoader) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (s SubscriptionCacheClientLoader) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	if rc.Operation.Operation == ast.Subscription {
		cacheClient := cache.NewRedisClient()
		ctx = context.WithValue(ctx, cache.DefaultClient, cacheClient)
	}
	return nil
}
