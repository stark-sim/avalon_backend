package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/configs"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"time"
)

const RedisUserKey = "avalon:graphql:user:%s:json"

type RedisClient struct {
	rdb *redis.Client
}

func (rc *RedisClient) GetUser(ctx context.Context, userID string) (*model.User, error) {
	val, err := rc.rdb.Get(ctx, fmt.Sprintf(RedisUserKey, userID)).Bytes()
	if err == redis.Nil {
		logrus.Debugf("key %s does not exist", fmt.Sprintf(RedisUserKey, userID))
		return nil, err
	} else if err != nil {
		logrus.Errorf("error at redis get User: %v", err)
		return nil, err
	} else {
		user := new(model.User)
		err = json.Unmarshal(val, user)
		if err != nil {
			logrus.Errorf("error at json load redis value to User: %v", err)
			return nil, err
		}
		return user, nil
	}
}

func (rc *RedisClient) SetUser(ctx context.Context, user *model.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		logrus.Errorf("error at json dump User to bytes: %v", err)
		return err
	}
	err = rc.rdb.Set(ctx, fmt.Sprintf(RedisUserKey, user.ID), val, time.Second*10).Err()
	if err != nil {
		logrus.Errorf("error at redis set User bytes: %v", err)
		return err
	}
	return nil
}

func NewRedisClient() *RedisClient {
	rc := RedisClient{rdb: redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  fmt.Sprintf("%s:%d", configs.Conf.RedisConfig.Host, configs.Conf.RedisConfig.Port),
		Dialer:                nil,
		OnConnect:             nil,
		Username:              "",
		Password:              configs.Conf.RedisConfig.Password,
		CredentialsProvider:   nil,
		DB:                    configs.Conf.RedisConfig.DB,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
	})}
	return &rc
}
