package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"time"

	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"github.com/stark-sim/avalon_backend/tools"
)

// GetRoomUser is the resolver for the GetRoomUser field.
func (r *subscriptionResolver) GetRoomUser(ctx context.Context) (<-chan *ent.RoomUser, error) {
	ch := make(chan *ent.RoomUser)
	go func() {
		roomUsers, err := r.client.RoomUser.Query().All(ctx)
		if err != nil {
			return
		}
		for _, roomUser := range roomUsers {
			select {
			case ch <- roomUser:
				// pass one
			default:
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()
	return ch, nil
}

// GetRoomUsers is the resolver for the getRoomUsers field.
func (r *subscriptionResolver) GetRoomUsers(ctx context.Context, req *model.RoomRequest) (<-chan []*ent.RoomUser, error) {
	roomID := tools.StringToInt64(req.ID)
	ch := make(chan []*ent.RoomUser)
	go func() {
		for {
			roomUsers, err := r.client.RoomUser.
				Query().
				Where(roomuser.RoomID(roomID), roomuser.DeletedAt(tools.ZeroTime)).
				Order(ent.Asc(roomuser.FieldUpdatedAt)).
				All(ctx)
			if err != nil {
				return
			}
			select {
			case ch <- roomUsers:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return ch, nil
}

// GetRoomOngoingGame is the resolver for the getRoomOngoingGame field.
func (r *subscriptionResolver) GetRoomOngoingGame(ctx context.Context, req *model.RoomRequest) (<-chan *ent.Game, error) {
	// 获取某个 Room 中 end_by 状态还是 none 的 Game，按道理说，没有 bug 的情况下，不会能找到两条数据
	ch := make(chan *ent.Game)
	roomID := tools.StringToInt64(req.ID)
	// 首先检查房间是正常的
	_, err := r.client.Room.Query().Where(room.ID(roomID), room.DeletedAt(tools.ZeroTime), room.Closed(false)).First(ctx)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			ongoingGames, err := r.client.Game.
				Query().
				Where(
					game.RoomID(roomID),
					game.DeletedAt(tools.ZeroTime),
					game.EndByEQ(game.EndByNone),
				).
				All(ctx)
			if err != nil {
				logrus.Errorf("error at querying ongoing game within room %s: %v", req.ID, err)
				return
			}
			ongoingGamesLen := len(ongoingGames)
			if ongoingGamesLen > 1 {
				logrus.Errorf("something is wrong, room %s has many ongoing games", req.ID)
				return
			}
			if ongoingGamesLen == 0 {
				// 游戏还没开始
				continue
			} else {
				// 游戏开始，返回 game，带上 gameUsers，让前端判断当前用户是否在局中，以决定进入游戏界面
				_game, err := r.client.Game.
					Query().
					Where(game.ID(ongoingGames[0].ID)).
					WithNamedGameUsers("gameUsers").
					First(ctx)
				if err != nil {
					logrus.Errorf("error at query the ongoing game: %v", err)
				}
				select {
				case ch <- _game:
					// 传输
				}
			}
		}
	}()
	return ch, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
