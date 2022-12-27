package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/configs"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"strconv"
	"time"
)

const RedisUserKey = "avalon:graphql:user:%s:json"
const RedisRoomIDShortCodeKey = "avalon:short_code:room:%s:int64"
const RedisRoomIDMutex = "avalon:mutex:room:%s:int64"
const RedisTempAssassinKey = "avalon:cache:assassin:%s:string"
const RedisTempPickKey = "avalon:cache:pick:%s:json"

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
	err = rc.rdb.Set(ctx, fmt.Sprintf(RedisUserKey, user.ID), val, time.Minute*10).Err()
	if err != nil {
		logrus.Errorf("error at redis set User bytes: %v", err)
		return err
	}
	return nil
}

func (rc *RedisClient) GetRoomIDByShortCode(ctx context.Context, shortCode string) (int64, error) {
	val, err := rc.rdb.Get(ctx, fmt.Sprintf(RedisRoomIDShortCodeKey, shortCode)).Int64()
	if err == redis.Nil {
		return 0, err
	} else if err != nil {
		return 0, err
	} else {
		return val, nil
	}
}

func (rc *RedisClient) SetRoomIDByShortCode(ctx context.Context, roomID int64) (string, error) {
	val := strconv.FormatInt(roomID, 10)
	val = val[len(val)-6:]
	err := rc.rdb.Set(ctx, fmt.Sprintf(RedisRoomIDShortCodeKey, val), roomID, 0).Err()
	if err != nil {
		logrus.Errorf("error at redis set RoomID: %v", err)
		return "", err
	}
	return val, nil
}

func (rc *RedisClient) DeleteRoomIDWithShortCode(ctx context.Context, roomID int64) error {
	val := strconv.FormatInt(roomID, 10)
	val = val[len(val)-6:]
	err := rc.rdb.Del(ctx, fmt.Sprintf(RedisRoomIDShortCodeKey, val)).Err()
	if err != nil {
		logrus.Errorf("error at redis delete RoomShortCode: %v", err)
		return err
	}
	return nil
}

// WaitRoomMutex 抢某个房间的信号量，用于解决 RoomUser 新增时的幻读问题
func (rc *RedisClient) WaitRoomMutex(ctx context.Context, roomID int64) error {
	for {
		done, err := rc.rdb.SetNX(ctx, fmt.Sprintf(RedisRoomIDMutex, strconv.FormatInt(roomID, 10)), roomID, 2*time.Second).Result()
		if err != nil {
			return err
		}
		if !done {
			continue
		} else {
			return nil
		}
	}
}

func (rc *RedisClient) ReleaseRoomMutex(ctx context.Context, roomID int64) error {
	_, err := rc.rdb.Del(ctx, fmt.Sprintf(RedisRoomIDMutex, strconv.FormatInt(roomID, 10))).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisClient) SetGameTempAssassinatedID(ctx context.Context, gameID, userID string) error {
	err := rc.rdb.Set(ctx, fmt.Sprintf(RedisTempAssassinKey, gameID), userID, 0).Err()
	if err != nil {
		logrus.Errorf("error at redis set gameTempAssassinatedID: %v", err)
		return err
	}
	return nil
}

func (rc *RedisClient) GetGameTempAssassinatedID(ctx context.Context, gameID string) (string, error) {
	res, err := rc.rdb.Get(ctx, fmt.Sprintf(RedisTempAssassinKey, gameID)).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		logrus.Errorf("error at get gameTempAssassinatedID: %v", err)
		return "", err
	} else {
		return res, nil
	}
}

func (rc *RedisClient) DeleteGameTempAssassinatedID(ctx context.Context, gameID string) error {
	// 真正决定刺杀后，就不需要这个值了
	err := rc.rdb.Del(ctx, fmt.Sprintf(RedisTempAssassinKey, gameID)).Err()
	if err != nil {
		logrus.Errorf("error at delete gameTempAssassinatedID: %v", err)
		return err
	}
	return nil
}

func (rc *RedisClient) SetMissionTempPickUserIDs(ctx context.Context, missionID string, userIDs []string) error {
	val, err := json.Marshal(userIDs)
	if err != nil {
		logrus.Errorf("error at json dump userIDs to bytes: %v", err)
		return err
	}
	err = rc.rdb.Set(ctx, fmt.Sprintf(RedisTempPickKey, missionID), val, 0).Err()
	if err != nil {
		logrus.Errorf("error at redis set tempPickUserIDs: %v", err)
		return err
	}
	return nil
}

func (rc *RedisClient) GetMissionTempPickUserIDs(ctx context.Context, missionID string) ([]string, error) {
	val, err := rc.rdb.Get(ctx, fmt.Sprintf(RedisTempPickKey, missionID)).Bytes()
	if err == redis.Nil {
		return []string{}, nil
	} else if err != nil {
		logrus.Errorf("error at redis get pick userIDs: %v", err)
		return nil, err
	} else {
		res := make([]string, 0)
		err = json.Unmarshal(val, &res)
		if err != nil {
			logrus.Errorf("error at json load redis temp pick userIDs: %v", err)
			return nil, err
		}
		return res, nil
	}
}

func (rc *RedisClient) DeleteMissionTempPickUserIDs(ctx context.Context, missionID string) error {
	err := rc.rdb.Del(ctx, fmt.Sprintf(RedisTempPickKey, missionID)).Err()
	if err != nil {
		logrus.Errorf("error at delete missionTempPickUserIDs: %v", err)
		return err
	}
	return nil
}

func (rc *RedisClient) Close() {
	err := rc.rdb.Close()
	if err != nil {
		logrus.Errorf("error at close redis client: %v", err)
		return
	}
	return
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
