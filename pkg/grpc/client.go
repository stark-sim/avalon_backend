package grpc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/configs"
	"github.com/stark-sim/cas/pkg/grpc/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCASGrpcConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", configs.Conf.APIConfig.CasHost, configs.Conf.APIConfig.CasGrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("failed connecting to server: %v", err)
		return nil, err
	}
	return conn, nil
}

func NewCASClient(conn *grpc.ClientConn) entpb.UserServiceClient {
	// 导入服务端提供的 Client，有点花活，应该把 proto 拿过来自己生成就好的
	client := entpb.NewUserServiceClient(conn)
	return client
}
