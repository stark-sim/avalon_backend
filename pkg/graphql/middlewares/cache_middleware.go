package middlewares

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/tools/cache"
	"github.com/vektah/gqlparser/v2/ast"
)

type SubscriptionCacheClientLoader struct{}

func (SubscriptionCacheClientLoader) ExtensionName() string {
	return "SubscriptionCacheClientLoader"
}

func (SubscriptionCacheClientLoader) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (s SubscriptionCacheClientLoader) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)
	logrus.Infof("called !!!")
	if oc.Operation.Operation == ast.Subscription {
		cacheClient := cache.NewRedisClient()
		ctx = context.WithValue(ctx, cache.DefaultClient, cacheClient)
	}
	return next(ctx)
}
