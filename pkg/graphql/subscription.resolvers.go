package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
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
func (r *subscriptionResolver) GetRoomOngoingGame(ctx context.Context, req model.RoomRequest) (<-chan *ent.Game, error) {
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
			time.Sleep(time.Second)
			if ongoingGamesLen == 0 {
				// 游戏还没开始
				select {
				case ch <- nil:
					// 返回空对象
				}
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

// GetMissionsByGame is the resolver for the getMissionsByGame field.
func (r *subscriptionResolver) GetMissionsByGame(ctx context.Context, req model.GameRequest) (<-chan []*ent.Mission, error) {
	// 前端目前用这个方法，只需要知道轮到哪个任务在进行，并且这些任务的状态，不需要知道 Mission 的 Squad 等后续数据
	ch := make(chan []*ent.Mission)
	// 检查如参 gameID
	gameID := tools.StringToInt64(req.ID)
	_, err := r.client.Game.Query().Where(game.ID(gameID), game.EndByEQ(game.EndByNone), game.DeletedAt(tools.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("error at query missions, gameID not exist or ended: %v", err)
		return nil, err
	}
	go func() {
		for {
			missions, err := r.client.Mission.
				Query().
				Where(mission.GameID(gameID), mission.DeletedAt(tools.ZeroTime)).
				Order(ent.Asc(mission.FieldSequence), ent.Asc(mission.FieldCreatedAt)).
				All(ctx)
			if err != nil {
				logrus.Errorf("error at query missions: %v", err)
				return
			}
			select {
			case ch <- missions:
				// 传输
			}
		}
	}()
	return ch, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
