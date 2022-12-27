package cache

import (
	"context"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
)

const DefaultClient = "REDIS"

type Client interface {
	Close()

	GetUser(ctx context.Context, userID string) (*model.User, error)
	SetUser(ctx context.Context, user *model.User) error

	GetRoomIDByShortCode(ctx context.Context, shortCode string) (int64, error)
	SetRoomIDByShortCode(ctx context.Context, roomID int64) (string, error)
	DeleteRoomIDWithShortCode(ctx context.Context, roomID int64) error

	WaitRoomMutex(ctx context.Context, roomID int64) error
	ReleaseRoomMutex(ctx context.Context, roomID int64) error

	SetGameTempAssassinatedID(ctx context.Context, gameID, userID string) error
	GetGameTempAssassinatedID(ctx context.Context, gameID string) (string, error)
	DeleteGameTempAssassinatedID(ctx context.Context, gameID string) error

	SetMissionTempPickUserIDs(ctx context.Context, missionID string, userIDs []string) error
	GetMissionTempPickUserIDs(ctx context.Context, missionID string) ([]string, error)
	DeleteMissionTempPickUserIDs(ctx context.Context, missionID string) error
}
