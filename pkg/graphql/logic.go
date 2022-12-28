package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"github.com/stark-sim/avalon_backend/pkg/grpc"
	"github.com/stark-sim/avalon_backend/tools/cache"
	cas "github.com/stark-sim/cas/pkg/grpc/pb"
	"github.com/vektah/gqlparser/v2/ast"
	grpc2 "google.golang.org/grpc"
	"strconv"
)

func GetUserAtResolver(ctx context.Context, userID int64) (*model.User, error) {
	oc := graphql.GetOperationContext(ctx)
	if oc.Operation.Operation != ast.Subscription {
		// 如果不是特殊操作，利用 Apollo SuperGraph 获取 User 的剩余数据
		return &model.User{ID: strconv.FormatInt(userID, 10)}, nil
	} else {
		// 如果是无法使用 Apollo 的特殊操作如 Subscription，用 gRPC 获取 User，需要加缓存
		strUserID := strconv.FormatInt(userID, 10)
		// 获取上下文中的 CacheClient
		cacheClient, ok := ctx.Value(cache.DefaultClient).(cache.Client)
		if !ok {
			return nil, errors.New("no redis client")
		}
		user, err := cacheClient.GetUser(ctx, strUserID)
		if err == redis.Nil {
			// 缓存里没有就调用 gRPC 获取，然后塞回缓存
			// 建立 grpc 连接
			conn, err := grpc.NewCASGrpcConn()
			if err != nil {
				return nil, err
			}
			defer func(conn *grpc2.ClientConn) {
				err := conn.Close()
				if err != nil {
					logrus.Errorf("error at closing grpc conn: %v", err)
				}
			}(conn)
			grpcClient := grpc.NewCASClient(conn)
			res, err := grpcClient.Get(ctx, &cas.UserGetRequest{
				Id: userID,
			})
			if err != nil {
				logrus.Errorf("error at get user using grpc: %v", err)
				return nil, err
			}
			user = &model.User{
				ID:    strUserID,
				Name:  res.Name,
				Phone: res.Phone,
			}
			if err = cacheClient.SetUser(ctx, user); err != nil {
				return nil, err
			}
			return user, nil
		} else if err != nil {
			return nil, err
		} else {
			return user, nil
		}
	}
}
