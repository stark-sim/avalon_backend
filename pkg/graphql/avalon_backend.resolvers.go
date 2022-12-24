package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	redis "github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/internal/logic"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"github.com/stark-sim/avalon_backend/pkg/grpc"
	"github.com/stark-sim/avalon_backend/tools"
	"github.com/stark-sim/avalon_backend/tools/cache"
	cas "github.com/stark-sim/cas/pkg/grpc/pb"
	"github.com/vektah/gqlparser/v2/ast"
	grpc2 "google.golang.org/grpc"
)

// ID is the resolver for the id field.
func (r *cardResolver) ID(ctx context.Context, obj *ent.Card) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *cardResolver) CreatedBy(ctx context.Context, obj *ent.Card) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *cardResolver) UpdatedBy(ctx context.Context, obj *ent.Card) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// ID is the resolver for the id field.
func (r *gameResolver) ID(ctx context.Context, obj *ent.Game) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *gameResolver) CreatedBy(ctx context.Context, obj *ent.Game) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *gameResolver) UpdatedBy(ctx context.Context, obj *ent.Game) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// RoomID is the resolver for the roomID field.
func (r *gameResolver) RoomID(ctx context.Context, obj *ent.Game) (string, error) {
	return strconv.FormatInt(obj.RoomID, 10), nil
}

// Capacity is the resolver for the capacity field.
func (r *gameResolver) Capacity(ctx context.Context, obj *ent.Game) (int, error) {
	return int(obj.Capacity), nil
}

// ID is the resolver for the id field.
func (r *gameUserResolver) ID(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *gameUserResolver) CreatedBy(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *gameUserResolver) UpdatedBy(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// UserID is the resolver for the userID field.
func (r *gameUserResolver) UserID(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// GameID is the resolver for the gameID field.
func (r *gameUserResolver) GameID(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CardID is the resolver for the cardID field.
func (r *gameUserResolver) CardID(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// Number is the resolver for the number field.
func (r *gameUserResolver) Number(ctx context.Context, obj *ent.GameUser) (int, error) {
	return int(obj.Number), nil
}

// User is the resolver for the user field.
func (r *gameUserResolver) User(ctx context.Context, obj *ent.GameUser) (*model.User, error) {
	fc := graphql.GetFieldContext(ctx)
	logrus.Infof("%s\n", fc.Object)
	return &model.User{ID: strconv.FormatInt(obj.UserID, 10)}, nil
}

// ID is the resolver for the id field.
func (r *missionResolver) ID(ctx context.Context, obj *ent.Mission) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *missionResolver) CreatedBy(ctx context.Context, obj *ent.Mission) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *missionResolver) UpdatedBy(ctx context.Context, obj *ent.Mission) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// Sequence is the resolver for the sequence field.
func (r *missionResolver) Sequence(ctx context.Context, obj *ent.Mission) (int, error) {
	return int(obj.Sequence), nil
}

// GameID is the resolver for the gameID field.
func (r *missionResolver) GameID(ctx context.Context, obj *ent.Mission) (string, error) {
	return strconv.FormatInt(obj.GameID, 10), nil
}

// Capacity is the resolver for the capacity field.
func (r *missionResolver) Capacity(ctx context.Context, obj *ent.Mission) (int, error) {
	return int(obj.Capacity), nil
}

// CreateRoom is the resolver for the createRoom field.
func (r *mutationResolver) CreateRoom(ctx context.Context, req ent.CreateRoomInput) (*ent.Room, error) {
	_room, err := r.client.Room.Create().SetInput(req).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 在 redis 中设置 房间快捷码
	redisClient := cache.NewRedisClient()
	_, err = redisClient.SetRoomIDByShortCode(ctx, _room.ID)
	if err != nil {
		return nil, err
	}
	redisClient.Close()
	return _room, nil
}

// JoinRoom is the resolver for the joinRoom field.
func (r *mutationResolver) JoinRoom(ctx context.Context, req ent.CreateRoomUserInput) (*ent.RoomUser, error) {
	// 加入前检查 房间 是否关闭
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_, err = tx.Room.Query().Where(room.ID(req.RoomID), room.DeletedAt(tools.ZeroTime), room.Closed(false)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, errors.New("room no exist")
	} else if err != nil {
		return nil, err
	}
	// 中间表调整软删除字段来代替创建和删除
	// 需要在 tx Commit 之前把之后有可能会拿的 Room 先拿出来
	roomUser, err := tx.RoomUser.Query().Where(roomuser.RoomID(req.RoomID), roomuser.UserID(req.UserID)).WithRoom().First(ctx)
	if ent.IsNotFound(err) {
		// 要插入 RoomUser 的话需要 Room 信号量
		redisClient := cache.NewRedisClient()
		err = redisClient.WaitRoomMutex(ctx, req.RoomID)
		if err != nil {
			return nil, err
		}
		newRoomUser, err := tx.RoomUser.
			Create().
			SetUserID(req.UserID).
			SetRoomID(req.RoomID).
			Save(ctx)
		newRoomUser, err = tx.RoomUser.Query().Where(roomuser.ID(newRoomUser.ID)).WithRoom().First(ctx)
		if err != nil {
			return nil, err
		}
		err = redisClient.ReleaseRoomMutex(ctx, req.RoomID)
		if err != nil {
			return nil, err
		}
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
		redisClient.Close()
		return newRoomUser, nil
	} else {
		err = tx.RoomUser.UpdateOneID(roomUser.ID).SetDeletedAt(tools.ZeroTime).Exec(ctx)
		if err != nil {
			return nil, err
		}
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
		return roomUser, nil
	}
}

// LeaveRoom is the resolver for the leaveRoom field.
func (r *mutationResolver) LeaveRoom(ctx context.Context, req ent.CreateRoomUserInput) (*ent.RoomUser, error) {
	redisClient := cache.NewRedisClient()
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	// 离开前查看房间剩余人数，要抢信号量锁住 RoomUser 表新增
	err = redisClient.WaitRoomMutex(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}
	count, err := tx.RoomUser.Query().Where(roomuser.RoomID(req.RoomID), roomuser.DeletedAt(tools.ZeroTime)).Count(ctx)
	if err != nil {
		return nil, err
	}
	// 查找后就可以释放信号量，因为加入房间时检查 Room.Closed 依赖 ReadCommitted 就足够，不需要信号量
	err = redisClient.ReleaseRoomMutex(ctx, req.RoomID)
	if err != nil {
		return nil, err
	}
	// 自己退出，由于需要返回，先查询出来
	roomUser, err := tx.RoomUser.
		Query().
		Where(roomuser.RoomID(req.RoomID), roomuser.UserID(req.UserID), roomuser.DeletedAt(tools.ZeroTime)).
		WithRoom().
		First(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.RoomUser.
		UpdateOne(roomUser).
		SetDeletedAt(time.Now()).
		Exec(ctx); err != nil {
		return nil, err
	}
	if count == 1 {
		// 如果只剩自己一个人，还需要关闭房间，并清除 shortCode
		_, err = tx.Room.UpdateOneID(req.RoomID).SetClosed(true).Save(ctx)
		if err != nil {
			return nil, err
		}
		err = redisClient.DeleteRoomIDWithShortCode(ctx, req.RoomID)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	redisClient.Close()
	return roomUser, nil
}

// CloseRoom is the resolver for the closeRoom field.
func (r *mutationResolver) CloseRoom(ctx context.Context, req model.RoomRequest) (*ent.Room, error) {
	roomID := tools.StringToInt64(req.ID)
	_room, err := r.client.Room.UpdateOneID(roomID).SetClosed(true).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 将 redis 中的 房间快捷码 清除
	redisClient := cache.NewRedisClient()
	err = redisClient.DeleteRoomIDWithShortCode(ctx, _room.ID)
	if err != nil {
		return nil, err
	}
	redisClient.Close()
	return _room, nil
}

// CreateGame is the resolver for the createGame field.
func (r *mutationResolver) CreateGame(ctx context.Context, req model.RoomRequest) (*ent.Game, error) {
	// 将房间中现有的人加入到一局新游戏里
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	roomID := tools.StringToInt64(req.ID)
	// 先检查房间没有在进行游戏
	_room, err := r.client.Room.
		Query().
		Where(room.ID(roomID), room.DeletedAt(tools.ZeroTime), room.GameOn(false)).
		First(ctx)
	if err != nil {
		logrus.Errorf("room %s not exist or gameOn, can't create game", req.ID)
		return nil, err
	}
	err = tx.Room.UpdateOne(_room).SetGameOn(true).Exec(ctx)
	if err != nil {
		logrus.Errorf("error at room update: %v", err)
		return nil, err
	}
	// 不锁了，开了后进来的不管，用户离开房间前查一下有没有在游戏里就好，离开和这里的查人会制衡
	roomUsers, err := tx.RoomUser.Query().Where(roomuser.RoomID(roomID), roomuser.DeletedAt(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	playerNum := uint8(len(roomUsers))
	// 创建游戏
	_game, err := tx.Game.
		Create().
		SetRoomID(roomID).
		SetEndBy(game.EndByNone).
		SetCapacity(playerNum).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	// 随机排序房间内用户和洗牌，然后创建 GameUser
	userIDs := make([]int64, playerNum)
	for i, roomUser := range roomUsers {
		userIDs[i] = roomUser.UserID
	}
	for i, v := range tools.Shuffle(userIDs) {
		userIDs[i] = v.(int64)
	}
	// 按人数拿牌，拿的时候已经洗好了
	cards, err := logic.GetShuffledCardsByNum(ctx, playerNum, nil)
	if err != nil {
		return nil, err
	}
	// 创建 GameUser，分牌分号
	gameUserCreates := make([]*ent.GameUserCreate, playerNum)
	for i := 0; i < len(roomUsers); i++ {
		gameUserCreates[i] = tx.GameUser.
			Create().
			SetGameID(_game.ID).
			SetUserID(userIDs[i]).
			SetCardID(cards[i].ID).
			SetNumber(uint8(i + 1))
	}
	_, err = tx.GameUser.CreateBulk(gameUserCreates...).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 创建 5 个 Mission，初始队长为 1-5 号玩家
	missionCreates := make([]*ent.MissionCreate, 5)
	for i := 0; i < 5; i++ {
		missionCreates[i] = tx.Mission.
			Create().
			SetGameID(_game.ID).
			SetLeader(userIDs[i]).
			SetCapacity(logic.GetMissionCapacityByNumAndSeq(playerNum, i)).
			SetSequence(uint8(i))
	}
	_, err = tx.Mission.CreateBulk(missionCreates...).Save(ctx)
	// 创建完毕，现在准备返回，把有可能需要的 EagerLoad 上
	_game, err = tx.Game.
		Query().
		Where(game.ID(_game.ID)).
		WithNamedMissions("missions").
		WithNamedGameUsers("gameUsers", func(query *ent.GameUserQuery) {
			query.WithCard()
		}).
		First(ctx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return _game, nil
}

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, req ent.CreateCardInput) (*ent.Card, error) {
	var tale string
	if req.Tale != nil {
		tale = *req.Tale
	}
	return r.client.Card.Create().SetName(*req.Name).SetRole(req.Role).SetTale(tale).Save(ctx)
}

// JoinRoomByShortCode is the resolver for the joinRoomByShortCode field.
func (r *mutationResolver) JoinRoomByShortCode(ctx context.Context, req model.JoinRoomInput) (*ent.RoomUser, error) {
	panic(fmt.Errorf("not implemented: JoinRoomByShortCode - joinRoomByShortCode"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	tempID := tools.StringToInt64(id)
	_game, err := r.client.Game.Query().Where(game.ID(tempID), game.DeletedAtEQ(tools.ZeroTime)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_room, err := r.client.Room.Query().Where(room.ID(tempID), room.DeletedAtEQ(tools.ZeroTime)).First(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					_card, err := r.client.Card.Query().Where(card.ID(tempID), card.DeletedAtEQ(tools.ZeroTime)).First(ctx)
					if err != nil {
						return nil, err
					}
					return _card, nil
				}
				return nil, err
			}
			return _room, nil
		}
		return nil, err
	}
	return _game, nil
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Cards is the resolver for the cards field.
func (r *queryResolver) Cards(ctx context.Context) ([]*ent.Card, error) {
	return r.client.Card.Query().All(ctx)
}

// Games is the resolver for the games field.
func (r *queryResolver) Games(ctx context.Context) ([]*ent.Game, error) {
	return r.client.Game.Query().All(ctx)
}

// GameUsers is the resolver for the gameUsers field.
func (r *queryResolver) GameUsers(ctx context.Context) ([]*ent.GameUser, error) {
	return r.client.GameUser.Query().All(ctx)
}

// Missions is the resolver for the missions field.
func (r *queryResolver) Missions(ctx context.Context) ([]*ent.Mission, error) {
	return r.client.Mission.Query().All(ctx)
}

// Records is the resolver for the records field.
func (r *queryResolver) Records(ctx context.Context) ([]*ent.Record, error) {
	return r.client.Record.Query().All(ctx)
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context) ([]*ent.Room, error) {
	return r.client.Room.Query().All(ctx)
}

// RoomUsers is the resolver for the roomUsers field.
func (r *queryResolver) RoomUsers(ctx context.Context) ([]*ent.RoomUser, error) {
	return r.client.RoomUser.Query().All(ctx)
}

// Squads is the resolver for the squads field.
func (r *queryResolver) Squads(ctx context.Context) ([]*ent.Squad, error) {
	return r.client.Squad.Query().All(ctx)
}

// Votes is the resolver for the votes field.
func (r *queryResolver) Votes(ctx context.Context) ([]*ent.Vote, error) {
	return r.client.Vote.Query().All(ctx)
}

// ID is the resolver for the id field.
func (r *recordResolver) ID(ctx context.Context, obj *ent.Record) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *recordResolver) CreatedBy(ctx context.Context, obj *ent.Record) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *recordResolver) UpdatedBy(ctx context.Context, obj *ent.Record) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// UserID is the resolver for the userID field.
func (r *recordResolver) UserID(ctx context.Context, obj *ent.Record) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// RoomID is the resolver for the roomID field.
func (r *recordResolver) RoomID(ctx context.Context, obj *ent.Record) (string, error) {
	return strconv.FormatInt(obj.RoomID, 10), nil
}

// User is the resolver for the user field.
func (r *recordResolver) User(ctx context.Context, obj *ent.Record) (*model.User, error) {
	fc := graphql.GetFieldContext(ctx)
	logrus.Infof("%s\n", fc.Object)
	return &model.User{
		ID: strconv.FormatInt(obj.UserID, 10),
	}, nil
}

// ID is the resolver for the id field.
func (r *roomResolver) ID(ctx context.Context, obj *ent.Room) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *roomResolver) CreatedBy(ctx context.Context, obj *ent.Room) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *roomResolver) UpdatedBy(ctx context.Context, obj *ent.Room) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// ID is the resolver for the id field.
func (r *roomUserResolver) ID(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *roomUserResolver) CreatedBy(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *roomUserResolver) UpdatedBy(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// UserID is the resolver for the userID field.
func (r *roomUserResolver) UserID(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// RoomID is the resolver for the roomID field.
func (r *roomUserResolver) RoomID(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.RoomID, 10), nil
}

// User is the resolver for the user field.
func (r *roomUserResolver) User(ctx context.Context, obj *ent.RoomUser) (*model.User, error) {
	oc := graphql.GetOperationContext(ctx)
	if oc.Operation.Operation != ast.Subscription {
		// 如果不是特殊操作，利用 Apollo SuperGraph 获取 User 的剩余数据
		return &model.User{ID: strconv.FormatInt(obj.UserID, 10)}, nil
	} else {
		// 如果是无法使用 Apollo 的特殊操作如 Subscription，用 gRPC 获取 User，需要加缓存
		userID := strconv.FormatInt(obj.UserID, 10)
		// 获取上下文中的 CacheClient
		cacheClient, ok := ctx.Value(cache.DefaultClient).(cache.Client)
		if !ok {
			return nil, errors.New("no redis client")
		}
		user, err := cacheClient.GetUser(ctx, userID)
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
				Id: obj.UserID,
			})
			if err != nil {
				logrus.Errorf("error at get user using grpc: %v", err)
				return nil, err
			}
			user = &model.User{
				ID:    userID,
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

// ID is the resolver for the id field.
func (r *squadResolver) ID(ctx context.Context, obj *ent.Squad) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *squadResolver) CreatedBy(ctx context.Context, obj *ent.Squad) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *squadResolver) UpdatedBy(ctx context.Context, obj *ent.Squad) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// MissionID is the resolver for the missionID field.
func (r *squadResolver) MissionID(ctx context.Context, obj *ent.Squad) (string, error) {
	return strconv.FormatInt(obj.MissionID, 10), nil
}

// UserID is the resolver for the userID field.
func (r *squadResolver) UserID(ctx context.Context, obj *ent.Squad) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// User is the resolver for the user field.
func (r *squadResolver) User(ctx context.Context, obj *ent.Squad) (*model.User, error) {
	fc := graphql.GetFieldContext(ctx)
	logrus.Infof("%s\n", fc.Object)
	return &model.User{
		ID: strconv.FormatInt(obj.UserID, 10),
	}, nil
}

// ID is the resolver for the id field.
func (r *voteResolver) ID(ctx context.Context, obj *ent.Vote) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *voteResolver) CreatedBy(ctx context.Context, obj *ent.Vote) (string, error) {
	return strconv.FormatInt(obj.CreatedBy, 10), nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *voteResolver) UpdatedBy(ctx context.Context, obj *ent.Vote) (string, error) {
	return strconv.FormatInt(obj.UpdatedBy, 10), nil
}

// MissionID is the resolver for the missionID field.
func (r *voteResolver) MissionID(ctx context.Context, obj *ent.Vote) (string, error) {
	return strconv.FormatInt(obj.MissionID, 10), nil
}

// UserID is the resolver for the userID field.
func (r *voteResolver) UserID(ctx context.Context, obj *ent.Vote) (string, error) {
	return strconv.FormatInt(obj.UserID, 10), nil
}

// User is the resolver for the user field.
func (r *voteResolver) User(ctx context.Context, obj *ent.Vote) (*model.User, error) {
	fc := graphql.GetFieldContext(ctx)
	logrus.Infof("%s\n", fc.Object)
	return &model.User{
		ID: strconv.FormatInt(obj.UserID, 10),
	}, nil
}

// ID is the resolver for the id field.
func (r *cardWhereInputResolver) ID(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *cardWhereInputResolver) IDNeq(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *cardWhereInputResolver) IDIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *cardWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *cardWhereInputResolver) IDGt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *cardWhereInputResolver) IDGte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *cardWhereInputResolver) IDLt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *cardWhereInputResolver) IDLte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *cardWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *cardWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *cardWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *cardWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *cardWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *cardWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *cardWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *cardWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *cardWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *cardWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *cardWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *cardWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *cardWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *cardWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *cardWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *cardWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *createCardInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateCardInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createCardInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateCardInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// GameUserIDs is the resolver for the gameUserIDs field.
func (r *createCardInputResolver) GameUserIDs(ctx context.Context, obj *ent.CreateCardInput, data []string) error {
	for _, v := range data {
		obj.GameUserIDs = append(obj.GameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createGameInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateGameInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createGameInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateGameInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// Capacity is the resolver for the capacity field.
func (r *createGameInputResolver) Capacity(ctx context.Context, obj *ent.CreateGameInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.Capacity = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// GameUserIDs is the resolver for the gameUserIDs field.
func (r *createGameInputResolver) GameUserIDs(ctx context.Context, obj *ent.CreateGameInput, data []string) error {
	for _, v := range data {
		obj.GameUserIDs = append(obj.GameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// MissionIDs is the resolver for the missionIDs field.
func (r *createGameInputResolver) MissionIDs(ctx context.Context, obj *ent.CreateGameInput, data []string) error {
	for _, v := range data {
		obj.MissionIDs = append(obj.MissionIDs, tools.StringToInt64(v))
	}
	return nil
}

// RoomID is the resolver for the roomID field.
func (r *createGameInputResolver) RoomID(ctx context.Context, obj *ent.CreateGameInput, data string) error {
	obj.RoomID = tools.StringToInt64(data)
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createGameUserInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateGameUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createGameUserInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateGameUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *createGameUserInputResolver) UserID(ctx context.Context, obj *ent.CreateGameUserInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.UserID = tempID
	return nil
}

// Number is the resolver for the number field.
func (r *createGameUserInputResolver) Number(ctx context.Context, obj *ent.CreateGameUserInput, data int) error {
	tempNum := uint8(data)
	obj.Number = tempNum
	return nil
}

// GameID is the resolver for the gameID field.
func (r *createGameUserInputResolver) GameID(ctx context.Context, obj *ent.CreateGameUserInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.GameID = tempID
	return nil
}

// CardID is the resolver for the cardID field.
func (r *createGameUserInputResolver) CardID(ctx context.Context, obj *ent.CreateGameUserInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.CardID = tempID
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createMissionInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateMissionInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createMissionInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateMissionInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// Sequence is the resolver for the sequence field.
func (r *createMissionInputResolver) Sequence(ctx context.Context, obj *ent.CreateMissionInput, data int) error {
	tempSequence := uint8(data)
	obj.Sequence = tempSequence
	return nil
}

// Capacity is the resolver for the capacity field.
func (r *createMissionInputResolver) Capacity(ctx context.Context, obj *ent.CreateMissionInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.Capacity = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// GameID is the resolver for the gameID field.
func (r *createMissionInputResolver) GameID(ctx context.Context, obj *ent.CreateMissionInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.GameID = tempID
	return nil
}

// SquadIDs is the resolver for the squadIDs field.
func (r *createMissionInputResolver) SquadIDs(ctx context.Context, obj *ent.CreateMissionInput, data []string) error {
	for _, v := range data {
		obj.SquadIDs = append(obj.SquadIDs, tools.StringToInt64(v))
	}
	return nil
}

// MissionVoteIDs is the resolver for the missionVoteIDs field.
func (r *createMissionInputResolver) MissionVoteIDs(ctx context.Context, obj *ent.CreateMissionInput, data []string) error {
	for _, v := range data {
		obj.MissionVoteIDs = append(obj.MissionVoteIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createRecordInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateRecordInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createRecordInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateRecordInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *createRecordInputResolver) UserID(ctx context.Context, obj *ent.CreateRecordInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.UserID = tempID
	return nil
}

// RoomID is the resolver for the roomID field.
func (r *createRecordInputResolver) RoomID(ctx context.Context, obj *ent.CreateRecordInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.RoomID = tempID
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createRoomInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateRoomInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createRoomInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateRoomInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomUserIDs is the resolver for the roomUserIDs field.
func (r *createRoomInputResolver) RoomUserIDs(ctx context.Context, obj *ent.CreateRoomInput, data []string) error {
	for _, v := range data {
		obj.RoomUserIDs = append(obj.RoomUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// GameIDs is the resolver for the gameIDs field.
func (r *createRoomInputResolver) GameIDs(ctx context.Context, obj *ent.CreateRoomInput, data []string) error {
	for _, v := range data {
		obj.GameIDs = append(obj.GameIDs, tools.StringToInt64(v))
	}
	return nil
}

// RecordIDs is the resolver for the recordIDs field.
func (r *createRoomInputResolver) RecordIDs(ctx context.Context, obj *ent.CreateRoomInput, data []string) error {
	for _, v := range data {
		obj.RecordIDs = append(obj.RecordIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createRoomUserInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateRoomUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createRoomUserInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateRoomUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *createRoomUserInputResolver) UserID(ctx context.Context, obj *ent.CreateRoomUserInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.UserID = tempID
	return nil
}

// RoomID is the resolver for the roomID field.
func (r *createRoomUserInputResolver) RoomID(ctx context.Context, obj *ent.CreateRoomUserInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.RoomID = tempID
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createSquadInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateSquadInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createSquadInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateSquadInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *createSquadInputResolver) UserID(ctx context.Context, obj *ent.CreateSquadInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.UserID = tempID
	return nil
}

// MissionID is the resolver for the missionID field.
func (r *createSquadInputResolver) MissionID(ctx context.Context, obj *ent.CreateSquadInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.MissionID = tempID
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *createVoteInputResolver) CreatedBy(ctx context.Context, obj *ent.CreateVoteInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *createVoteInputResolver) UpdatedBy(ctx context.Context, obj *ent.CreateVoteInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *createVoteInputResolver) UserID(ctx context.Context, obj *ent.CreateVoteInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.UserID = tempID
	return nil
}

// MissionID is the resolver for the missionID field.
func (r *createVoteInputResolver) MissionID(ctx context.Context, obj *ent.CreateVoteInput, data string) error {
	tempID := tools.StringToInt64(data)
	obj.MissionID = tempID
	return nil
}

// ID is the resolver for the id field.
func (r *gameUserWhereInputResolver) ID(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *gameUserWhereInputResolver) IDNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *gameUserWhereInputResolver) IDIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *gameUserWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *gameUserWhereInputResolver) IDGt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDGte is the resolver for the idGTE field.
func (r *gameUserWhereInputResolver) IDGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLt is the resolver for the idLT field.
func (r *gameUserWhereInputResolver) IDLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLte is the resolver for the idLTE field.
func (r *gameUserWhereInputResolver) IDLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedBy is the resolver for the createdBy field.
func (r *gameUserWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *gameUserWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *gameUserWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *gameUserWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *gameUserWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *gameUserWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *gameUserWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *gameUserWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *gameUserWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *gameUserWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *gameUserWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *gameUserWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *gameUserWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *gameUserWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *gameUserWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *gameUserWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// UserID is the resolver for the userID field.
func (r *gameUserWhereInputResolver) UserID(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDNeq is the resolver for the userIDNEQ field.
func (r *gameUserWhereInputResolver) UserIDNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDIn is the resolver for the userIDIn field.
func (r *gameUserWhereInputResolver) UserIDIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDIn = append(obj.UserIDIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDNotIn is the resolver for the userIDNotIn field.
func (r *gameUserWhereInputResolver) UserIDNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDNotIn = append(obj.UserIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDGt is the resolver for the userIDGT field.
func (r *gameUserWhereInputResolver) UserIDGt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGt - userIDGT"))
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *gameUserWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGte - userIDGTE"))
}

// UserIDLt is the resolver for the userIDLT field.
func (r *gameUserWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLt - userIDLT"))
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *gameUserWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLte - userIDLTE"))
}

// GameID is the resolver for the gameID field.
func (r *gameUserWhereInputResolver) GameID(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.GameID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// GameIDNeq is the resolver for the gameIDNEQ field.
func (r *gameUserWhereInputResolver) GameIDNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// GameIDIn is the resolver for the gameIDIn field.
func (r *gameUserWhereInputResolver) GameIDIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.GameIDIn = append(obj.GameIDIn, tools.StringToInt64(v))
	}
	return nil
}

// GameIDNotIn is the resolver for the gameIDNotIn field.
func (r *gameUserWhereInputResolver) GameIDNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.GameIDNotIn = append(obj.GameIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CardID is the resolver for the cardID field.
func (r *gameUserWhereInputResolver) CardID(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.CardID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CardIDNeq is the resolver for the cardIDNEQ field.
func (r *gameUserWhereInputResolver) CardIDNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.CardIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CardIDIn is the resolver for the cardIDIn field.
func (r *gameUserWhereInputResolver) CardIDIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.CardIDIn = append(obj.CardIDIn, tools.StringToInt64(v))
	}
	return nil
}

// CardIDNotIn is the resolver for the cardIDNotIn field.
func (r *gameUserWhereInputResolver) CardIDNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	for _, v := range data {
		obj.CardIDNotIn = append(obj.CardIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// Number is the resolver for the number field.
func (r *gameUserWhereInputResolver) Number(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNum := uint8(*data)
		obj.Number = &tempNum
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberNeq is the resolver for the numberNEQ field.
func (r *gameUserWhereInputResolver) NumberNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNum := uint8(*data)
		obj.NumberNEQ = &tempNum
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberIn is the resolver for the numberIn field.
func (r *gameUserWhereInputResolver) NumberIn(ctx context.Context, obj *ent.GameUserWhereInput, data []int) error {
	for _, v := range data {
		obj.NumberIn = append(obj.NumberIn, uint8(v))
	}
	return nil
}

// NumberNotIn is the resolver for the numberNotIn field.
func (r *gameUserWhereInputResolver) NumberNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []int) error {
	for _, v := range data {
		obj.NumberNotIn = append(obj.NumberNotIn, uint8(v))
	}
	return nil
}

// NumberGt is the resolver for the numberGT field.
func (r *gameUserWhereInputResolver) NumberGt(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNumber := uint8(*data)
		obj.NumberGT = &tempNumber
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberGte is the resolver for the numberGTE field.
func (r *gameUserWhereInputResolver) NumberGte(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNumber := uint8(*data)
		obj.NumberGTE = &tempNumber
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberLt is the resolver for the numberLT field.
func (r *gameUserWhereInputResolver) NumberLt(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNumber := uint8(*data)
		obj.NumberLT = &tempNumber
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberLte is the resolver for the numberLTE field.
func (r *gameUserWhereInputResolver) NumberLte(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNumber := uint8(*data)
		obj.NumberLTE = &tempNumber
		return nil
	} else {
		return errors.New("null")
	}
}

// ID is the resolver for the id field.
func (r *gameWhereInputResolver) ID(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *gameWhereInputResolver) IDNeq(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *gameWhereInputResolver) IDIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *gameWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *gameWhereInputResolver) IDGt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *gameWhereInputResolver) IDGte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *gameWhereInputResolver) IDLt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *gameWhereInputResolver) IDLte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *gameWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *gameWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *gameWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *gameWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *gameWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *gameWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *gameWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *gameWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *gameWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *gameWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *gameWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *gameWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *gameWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *gameWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *gameWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *gameWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomID is the resolver for the roomID field.
func (r *gameWhereInputResolver) RoomID(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomIDNeq is the resolver for the roomIDNEQ field.
func (r *gameWhereInputResolver) RoomIDNeq(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomIDIn is the resolver for the roomIDIn field.
func (r *gameWhereInputResolver) RoomIDIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.RoomIDIn = append(obj.RoomIDIn, tools.StringToInt64(v))
	}
	return nil
}

// RoomIDNotIn is the resolver for the roomIDNotIn field.
func (r *gameWhereInputResolver) RoomIDNotIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	for _, v := range data {
		obj.RoomIDNotIn = append(obj.RoomIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// Capacity is the resolver for the capacity field.
func (r *gameWhereInputResolver) Capacity(ctx context.Context, obj *ent.GameWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.Capacity = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityNeq is the resolver for the capacityNEQ field.
func (r *gameWhereInputResolver) CapacityNeq(ctx context.Context, obj *ent.GameWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityNEQ = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityIn is the resolver for the capacityIn field.
func (r *gameWhereInputResolver) CapacityIn(ctx context.Context, obj *ent.GameWhereInput, data []int) error {
	for _, v := range data {
		obj.CapacityIn = append(obj.CapacityIn, uint8(v))
	}
	return nil
}

// CapacityNotIn is the resolver for the capacityNotIn field.
func (r *gameWhereInputResolver) CapacityNotIn(ctx context.Context, obj *ent.GameWhereInput, data []int) error {
	for _, v := range data {
		obj.CapacityNotIn = append(obj.CapacityNotIn, uint8(v))
	}
	return nil
}

// CapacityGt is the resolver for the capacityGT field.
func (r *gameWhereInputResolver) CapacityGt(ctx context.Context, obj *ent.GameWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityGT = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityGte is the resolver for the capacityGTE field.
func (r *gameWhereInputResolver) CapacityGte(ctx context.Context, obj *ent.GameWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityGTE = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityLt is the resolver for the capacityLT field.
func (r *gameWhereInputResolver) CapacityLt(ctx context.Context, obj *ent.GameWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityLT = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityLte is the resolver for the capacityLTE field.
func (r *gameWhereInputResolver) CapacityLte(ctx context.Context, obj *ent.GameWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityLTE = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// ID is the resolver for the id field.
func (r *missionWhereInputResolver) ID(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *missionWhereInputResolver) IDNeq(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *missionWhereInputResolver) IDIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *missionWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *missionWhereInputResolver) IDGt(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *missionWhereInputResolver) IDGte(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *missionWhereInputResolver) IDLt(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *missionWhereInputResolver) IDLte(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *missionWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *missionWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *missionWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *missionWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *missionWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *missionWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *missionWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *missionWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *missionWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *missionWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *missionWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *missionWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *missionWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *missionWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *missionWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *missionWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// Sequence is the resolver for the sequence field.
func (r *missionWhereInputResolver) Sequence(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.Sequence = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// SequenceNeq is the resolver for the sequenceNEQ field.
func (r *missionWhereInputResolver) SequenceNeq(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.SequenceNEQ = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// SequenceIn is the resolver for the sequenceIn field.
func (r *missionWhereInputResolver) SequenceIn(ctx context.Context, obj *ent.MissionWhereInput, data []int) error {
	for _, v := range data {
		obj.SequenceIn = append(obj.SequenceIn, uint8(v))
	}
	return nil
}

// SequenceNotIn is the resolver for the sequenceNotIn field.
func (r *missionWhereInputResolver) SequenceNotIn(ctx context.Context, obj *ent.MissionWhereInput, data []int) error {
	for _, v := range data {
		obj.SequenceIn = append(obj.SequenceIn, uint8(v))
	}
	return nil
}

// SequenceGt is the resolver for the sequenceGT field.
func (r *missionWhereInputResolver) SequenceGt(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.SequenceGT = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// SequenceGte is the resolver for the sequenceGTE field.
func (r *missionWhereInputResolver) SequenceGte(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.SequenceGTE = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// SequenceLt is the resolver for the sequenceLT field.
func (r *missionWhereInputResolver) SequenceLt(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.SequenceLT = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// SequenceLte is the resolver for the sequenceLTE field.
func (r *missionWhereInputResolver) SequenceLte(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.SequenceLTE = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// GameID is the resolver for the gameID field.
func (r *missionWhereInputResolver) GameID(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.GameID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// GameIDNeq is the resolver for the gameIDNEQ field.
func (r *missionWhereInputResolver) GameIDNeq(ctx context.Context, obj *ent.MissionWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.GameIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// GameIDIn is the resolver for the gameIDIn field.
func (r *missionWhereInputResolver) GameIDIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.GameIDIn = append(obj.GameIDIn, tools.StringToInt64(v))
	}
	return nil
}

// GameIDNotIn is the resolver for the gameIDNotIn field.
func (r *missionWhereInputResolver) GameIDNotIn(ctx context.Context, obj *ent.MissionWhereInput, data []string) error {
	for _, v := range data {
		obj.GameIDNotIn = append(obj.GameIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// Capacity is the resolver for the capacity field.
func (r *missionWhereInputResolver) Capacity(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.Capacity = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityNeq is the resolver for the capacityNEQ field.
func (r *missionWhereInputResolver) CapacityNeq(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityNEQ = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityIn is the resolver for the capacityIn field.
func (r *missionWhereInputResolver) CapacityIn(ctx context.Context, obj *ent.MissionWhereInput, data []int) error {
	for _, v := range data {
		obj.CapacityIn = append(obj.CapacityIn, uint8(v))
	}
	return nil
}

// CapacityNotIn is the resolver for the capacityNotIn field.
func (r *missionWhereInputResolver) CapacityNotIn(ctx context.Context, obj *ent.MissionWhereInput, data []int) error {
	for _, v := range data {
		obj.CapacityNotIn = append(obj.CapacityNotIn, uint8(v))
	}
	return nil
}

// CapacityGt is the resolver for the capacityGT field.
func (r *missionWhereInputResolver) CapacityGt(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityGT = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityGte is the resolver for the capacityGTE field.
func (r *missionWhereInputResolver) CapacityGte(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityGTE = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityLt is the resolver for the capacityLT field.
func (r *missionWhereInputResolver) CapacityLt(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityLT = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// CapacityLte is the resolver for the capacityLTE field.
func (r *missionWhereInputResolver) CapacityLte(ctx context.Context, obj *ent.MissionWhereInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.CapacityLTE = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// ID is the resolver for the id field.
func (r *recordWhereInputResolver) ID(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *recordWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *recordWhereInputResolver) IDIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *recordWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *recordWhereInputResolver) IDGt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *recordWhereInputResolver) IDGte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *recordWhereInputResolver) IDLt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *recordWhereInputResolver) IDLte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *recordWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *recordWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *recordWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *recordWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *recordWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *recordWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *recordWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *recordWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *recordWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *recordWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *recordWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *recordWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *recordWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *recordWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *recordWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *recordWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// UserID is the resolver for the userID field.
func (r *recordWhereInputResolver) UserID(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDNeq is the resolver for the userIDNEQ field.
func (r *recordWhereInputResolver) UserIDNeq(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDIn is the resolver for the userIDIn field.
func (r *recordWhereInputResolver) UserIDIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDIn = append(obj.UserIDIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDNotIn is the resolver for the userIDNotIn field.
func (r *recordWhereInputResolver) UserIDNotIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDNotIn = append(obj.UserIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDGt is the resolver for the userIDGT field.
func (r *recordWhereInputResolver) UserIDGt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGt - userIDGT"))
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *recordWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGte - userIDGTE"))
}

// UserIDLt is the resolver for the userIDLT field.
func (r *recordWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLt - userIDLT"))
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *recordWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLte - userIDLTE"))
}

// RoomID is the resolver for the roomID field.
func (r *recordWhereInputResolver) RoomID(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomIDNeq is the resolver for the roomIDNEQ field.
func (r *recordWhereInputResolver) RoomIDNeq(ctx context.Context, obj *ent.RecordWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomIDIn is the resolver for the roomIDIn field.
func (r *recordWhereInputResolver) RoomIDIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.RoomIDIn = append(obj.RoomIDIn, tools.StringToInt64(v))
	}
	return nil
}

// RoomIDNotIn is the resolver for the roomIDNotIn field.
func (r *recordWhereInputResolver) RoomIDNotIn(ctx context.Context, obj *ent.RecordWhereInput, data []string) error {
	for _, v := range data {
		obj.RoomIDNotIn = append(obj.RoomIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// ID is the resolver for the id field.
func (r *roomUserWhereInputResolver) ID(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *roomUserWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *roomUserWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *roomUserWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *roomUserWhereInputResolver) IDGt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *roomUserWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *roomUserWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *roomUserWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *roomUserWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *roomUserWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *roomUserWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *roomUserWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *roomUserWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *roomUserWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *roomUserWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *roomUserWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *roomUserWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *roomUserWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *roomUserWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *roomUserWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *roomUserWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *roomUserWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *roomUserWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *roomUserWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// UserID is the resolver for the userID field.
func (r *roomUserWhereInputResolver) UserID(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDNeq is the resolver for the userIDNEQ field.
func (r *roomUserWhereInputResolver) UserIDNeq(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDIn is the resolver for the userIDIn field.
func (r *roomUserWhereInputResolver) UserIDIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDIn = append(obj.UserIDIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDNotIn is the resolver for the userIDNotIn field.
func (r *roomUserWhereInputResolver) UserIDNotIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDNotIn = append(obj.UserIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDGt is the resolver for the userIDGT field.
func (r *roomUserWhereInputResolver) UserIDGt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGt - userIDGT"))
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *roomUserWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGte - userIDGTE"))
}

// UserIDLt is the resolver for the userIDLT field.
func (r *roomUserWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLt - userIDLT"))
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *roomUserWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLte - userIDLTE"))
}

// RoomID is the resolver for the roomID field.
func (r *roomUserWhereInputResolver) RoomID(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomIDNeq is the resolver for the roomIDNEQ field.
func (r *roomUserWhereInputResolver) RoomIDNeq(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomIDIn is the resolver for the roomIDIn field.
func (r *roomUserWhereInputResolver) RoomIDIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.RoomIDIn = append(obj.RoomIDIn, tools.StringToInt64(v))
	}
	return nil
}

// RoomIDNotIn is the resolver for the roomIDNotIn field.
func (r *roomUserWhereInputResolver) RoomIDNotIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	for _, v := range data {
		obj.RoomIDNotIn = append(obj.RoomIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// ID is the resolver for the id field.
func (r *roomWhereInputResolver) ID(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *roomWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *roomWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *roomWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *roomWhereInputResolver) IDGt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *roomWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *roomWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *roomWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *roomWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *roomWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *roomWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *roomWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *roomWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *roomWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *roomWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *roomWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *roomWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *roomWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *roomWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *roomWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *roomWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *roomWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *roomWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *roomWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// ID is the resolver for the id field.
func (r *squadWhereInputResolver) ID(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *squadWhereInputResolver) IDNeq(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *squadWhereInputResolver) IDIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *squadWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *squadWhereInputResolver) IDGt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *squadWhereInputResolver) IDGte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *squadWhereInputResolver) IDLt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *squadWhereInputResolver) IDLte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *squadWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *squadWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *squadWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *squadWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *squadWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *squadWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *squadWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *squadWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *squadWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *squadWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *squadWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *squadWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *squadWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *squadWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *squadWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *squadWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// MissionID is the resolver for the missionID field.
func (r *squadWhereInputResolver) MissionID(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.MissionID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// MissionIDNeq is the resolver for the missionIDNEQ field.
func (r *squadWhereInputResolver) MissionIDNeq(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.MissionIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// MissionIDIn is the resolver for the missionIDIn field.
func (r *squadWhereInputResolver) MissionIDIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.MissionIDIn = append(obj.MissionIDIn, tools.StringToInt64(v))
	}
	return nil
}

// MissionIDNotIn is the resolver for the missionIDNotIn field.
func (r *squadWhereInputResolver) MissionIDNotIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.MissionIDNotIn = append(obj.MissionIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserID is the resolver for the userID field.
func (r *squadWhereInputResolver) UserID(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDNeq is the resolver for the userIDNEQ field.
func (r *squadWhereInputResolver) UserIDNeq(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDIn is the resolver for the userIDIn field.
func (r *squadWhereInputResolver) UserIDIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDIn = append(obj.UserIDIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDNotIn is the resolver for the userIDNotIn field.
func (r *squadWhereInputResolver) UserIDNotIn(ctx context.Context, obj *ent.SquadWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDNotIn = append(obj.UserIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDGt is the resolver for the userIDGT field.
func (r *squadWhereInputResolver) UserIDGt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGt - userIDGT"))
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *squadWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGte - userIDGTE"))
}

// UserIDLt is the resolver for the userIDLT field.
func (r *squadWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLt - userIDLT"))
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *squadWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.SquadWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLte - userIDLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateCardInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateCardInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateCardInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateCardInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// AddGameUserIDs is the resolver for the addGameUserIDs field.
func (r *updateCardInputResolver) AddGameUserIDs(ctx context.Context, obj *ent.UpdateCardInput, data []string) error {
	for _, v := range data {
		obj.AddGameUserIDs = append(obj.AddGameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveGameUserIDs is the resolver for the removeGameUserIDs field.
func (r *updateCardInputResolver) RemoveGameUserIDs(ctx context.Context, obj *ent.UpdateCardInput, data []string) error {
	for _, v := range data {
		obj.RemoveGameUserIDs = append(obj.RemoveGameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateGameInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateGameInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateGameInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateGameInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// Capacity is the resolver for the capacity field.
func (r *updateGameInputResolver) Capacity(ctx context.Context, obj *ent.UpdateGameInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.Capacity = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// AddGameUserIDs is the resolver for the addGameUserIDs field.
func (r *updateGameInputResolver) AddGameUserIDs(ctx context.Context, obj *ent.UpdateGameInput, data []string) error {
	for _, v := range data {
		obj.AddGameUserIDs = append(obj.AddGameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveGameUserIDs is the resolver for the removeGameUserIDs field.
func (r *updateGameInputResolver) RemoveGameUserIDs(ctx context.Context, obj *ent.UpdateGameInput, data []string) error {
	for _, v := range data {
		obj.RemoveGameUserIDs = append(obj.RemoveGameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// AddMissionIDs is the resolver for the addMissionIDs field.
func (r *updateGameInputResolver) AddMissionIDs(ctx context.Context, obj *ent.UpdateGameInput, data []string) error {
	for _, v := range data {
		obj.AddMissionIDs = append(obj.AddMissionIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveMissionIDs is the resolver for the removeMissionIDs field.
func (r *updateGameInputResolver) RemoveMissionIDs(ctx context.Context, obj *ent.UpdateGameInput, data []string) error {
	for _, v := range data {
		obj.RemoveMissionIDs = append(obj.RemoveMissionIDs, tools.StringToInt64(v))
	}
	return nil
}

// RoomID is the resolver for the roomID field.
func (r *updateGameInputResolver) RoomID(ctx context.Context, obj *ent.UpdateGameInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateGameUserInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateGameUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateGameUserInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateGameUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *updateGameUserInputResolver) UserID(ctx context.Context, obj *ent.UpdateGameUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// Number is the resolver for the number field.
func (r *updateGameUserInputResolver) Number(ctx context.Context, obj *ent.UpdateGameUserInput, data *int) error {
	if data != nil {
		tempNum := uint8(*data)
		obj.Number = &tempNum
		return nil
	} else {
		return errors.New("null")
	}
}

// GameID is the resolver for the gameID field.
func (r *updateGameUserInputResolver) GameID(ctx context.Context, obj *ent.UpdateGameUserInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.GameID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CardID is the resolver for the cardID field.
func (r *updateGameUserInputResolver) CardID(ctx context.Context, obj *ent.UpdateGameUserInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.CardID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateMissionInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateMissionInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateMissionInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateMissionInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// Sequence is the resolver for the sequence field.
func (r *updateMissionInputResolver) Sequence(ctx context.Context, obj *ent.UpdateMissionInput, data *int) error {
	if data != nil {
		tempSequence := uint8(*data)
		obj.Sequence = &tempSequence
		return nil
	} else {
		return errors.New("null")
	}
}

// Capacity is the resolver for the capacity field.
func (r *updateMissionInputResolver) Capacity(ctx context.Context, obj *ent.UpdateMissionInput, data *int) error {
	if data != nil {
		tempCapacity := uint8(*data)
		obj.Capacity = &tempCapacity
		return nil
	} else {
		return errors.New("null")
	}
}

// GameID is the resolver for the gameID field.
func (r *updateMissionInputResolver) GameID(ctx context.Context, obj *ent.UpdateMissionInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.GameID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// AddSquadIDs is the resolver for the addSquadIDs field.
func (r *updateMissionInputResolver) AddSquadIDs(ctx context.Context, obj *ent.UpdateMissionInput, data []string) error {
	for _, v := range data {
		obj.AddSquadIDs = append(obj.AddSquadIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveSquadIDs is the resolver for the removeSquadIDs field.
func (r *updateMissionInputResolver) RemoveSquadIDs(ctx context.Context, obj *ent.UpdateMissionInput, data []string) error {
	for _, v := range data {
		obj.RemoveSquadIDs = append(obj.RemoveSquadIDs, tools.StringToInt64(v))
	}
	return nil
}

// AddMissionVoteIDs is the resolver for the addMissionVoteIDs field.
func (r *updateMissionInputResolver) AddMissionVoteIDs(ctx context.Context, obj *ent.UpdateMissionInput, data []string) error {
	for _, v := range data {
		obj.AddMissionVoteIDs = append(obj.AddMissionVoteIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveMissionVoteIDs is the resolver for the removeMissionVoteIDs field.
func (r *updateMissionInputResolver) RemoveMissionVoteIDs(ctx context.Context, obj *ent.UpdateMissionInput, data []string) error {
	for _, v := range data {
		obj.RemoveMissionVoteIDs = append(obj.RemoveMissionVoteIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateRecordInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateRecordInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateRecordInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateRecordInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *updateRecordInputResolver) UserID(ctx context.Context, obj *ent.UpdateRecordInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomID is the resolver for the roomID field.
func (r *updateRecordInputResolver) RoomID(ctx context.Context, obj *ent.UpdateRecordInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateRoomInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateRoomInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateRoomInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateRoomInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// AddRoomUserIDs is the resolver for the addRoomUserIDs field.
func (r *updateRoomInputResolver) AddRoomUserIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	for _, v := range data {
		obj.AddRoomUserIDs = append(obj.AddRoomUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveRoomUserIDs is the resolver for the removeRoomUserIDs field.
func (r *updateRoomInputResolver) RemoveRoomUserIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	for _, v := range data {
		obj.RemoveRoomUserIDs = append(obj.RemoveRoomUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// AddGameIDs is the resolver for the addGameIDs field.
func (r *updateRoomInputResolver) AddGameIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	for _, v := range data {
		obj.AddGameIDs = append(obj.AddGameIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveGameIDs is the resolver for the removeGameIDs field.
func (r *updateRoomInputResolver) RemoveGameIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	for _, v := range data {
		obj.RemoveGameIDs = append(obj.RemoveGameIDs, tools.StringToInt64(v))
	}
	return nil
}

// AddRecordIDs is the resolver for the addRecordIDs field.
func (r *updateRoomInputResolver) AddRecordIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	for _, v := range data {
		obj.AddRecordIDs = append(obj.AddRecordIDs, tools.StringToInt64(v))
	}
	return nil
}

// RemoveRecordIDs is the resolver for the removeRecordIDs field.
func (r *updateRoomInputResolver) RemoveRecordIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	for _, v := range data {
		obj.RemoveRecordIDs = append(obj.RemoveRecordIDs, tools.StringToInt64(v))
	}
	return nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateRoomUserInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateRoomUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateRoomUserInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateRoomUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *updateRoomUserInputResolver) UserID(ctx context.Context, obj *ent.UpdateRoomUserInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// RoomID is the resolver for the roomID field.
func (r *updateRoomUserInputResolver) RoomID(ctx context.Context, obj *ent.UpdateRoomUserInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.RoomID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateSquadInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateSquadInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateSquadInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateSquadInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *updateSquadInputResolver) UserID(ctx context.Context, obj *ent.UpdateSquadInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// MissionID is the resolver for the missionID field.
func (r *updateSquadInputResolver) MissionID(ctx context.Context, obj *ent.UpdateSquadInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.MissionID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedBy is the resolver for the createdBy field.
func (r *updateVoteInputResolver) CreatedBy(ctx context.Context, obj *ent.UpdateVoteInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *updateVoteInputResolver) UpdatedBy(ctx context.Context, obj *ent.UpdateVoteInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserID is the resolver for the userID field.
func (r *updateVoteInputResolver) UserID(ctx context.Context, obj *ent.UpdateVoteInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// MissionID is the resolver for the missionID field.
func (r *updateVoteInputResolver) MissionID(ctx context.Context, obj *ent.UpdateVoteInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.MissionID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// ID is the resolver for the id field.
func (r *voteWhereInputResolver) ID(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *voteWhereInputResolver) IDNeq(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.IDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDIn is the resolver for the idIn field.
func (r *voteWhereInputResolver) IDIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.IDIn = append(obj.IDIn, tools.StringToInt64(v))
	}
	return nil
}

// IDNotIn is the resolver for the idNotIn field.
func (r *voteWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.IDNotIn = append(obj.IDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// IDGt is the resolver for the idGT field.
func (r *voteWhereInputResolver) IDGt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *voteWhereInputResolver) IDGte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *voteWhereInputResolver) IDLt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *voteWhereInputResolver) IDLte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *voteWhereInputResolver) CreatedBy(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByNeq is the resolver for the createdByNEQ field.
func (r *voteWhereInputResolver) CreatedByNeq(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByIn is the resolver for the createdByIn field.
func (r *voteWhereInputResolver) CreatedByIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByIn = append(obj.CreatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByNotIn is the resolver for the createdByNotIn field.
func (r *voteWhereInputResolver) CreatedByNotIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.CreatedByNotIn = append(obj.CreatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// CreatedByGt is the resolver for the createdByGT field.
func (r *voteWhereInputResolver) CreatedByGt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGt - createdByGT"))
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *voteWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByGte - createdByGTE"))
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *voteWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLt - createdByLT"))
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *voteWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: CreatedByLte - createdByLTE"))
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *voteWhereInputResolver) UpdatedBy(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedBy = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByNeq is the resolver for the updatedByNEQ field.
func (r *voteWhereInputResolver) UpdatedByNeq(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByIn is the resolver for the updatedByIn field.
func (r *voteWhereInputResolver) UpdatedByIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByIn = append(obj.UpdatedByIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByNotIn is the resolver for the updatedByNotIn field.
func (r *voteWhereInputResolver) UpdatedByNotIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.UpdatedByNotIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *voteWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGt - updatedByGT"))
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *voteWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByGte - updatedByGTE"))
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *voteWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLt - updatedByLT"))
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *voteWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UpdatedByLte - updatedByLTE"))
}

// MissionID is the resolver for the missionID field.
func (r *voteWhereInputResolver) MissionID(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.MissionID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// MissionIDNeq is the resolver for the missionIDNEQ field.
func (r *voteWhereInputResolver) MissionIDNeq(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.MissionIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// MissionIDIn is the resolver for the missionIDIn field.
func (r *voteWhereInputResolver) MissionIDIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.MissionIDIn = append(obj.MissionIDIn, tools.StringToInt64(v))
	}
	return nil
}

// MissionIDNotIn is the resolver for the missionIDNotIn field.
func (r *voteWhereInputResolver) MissionIDNotIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.MissionIDNotIn = append(obj.MissionIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserID is the resolver for the userID field.
func (r *voteWhereInputResolver) UserID(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDNeq is the resolver for the userIDNEQ field.
func (r *voteWhereInputResolver) UserIDNeq(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDNEQ = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDIn is the resolver for the userIDIn field.
func (r *voteWhereInputResolver) UserIDIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDIn = append(obj.UserIDIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDNotIn is the resolver for the userIDNotIn field.
func (r *voteWhereInputResolver) UserIDNotIn(ctx context.Context, obj *ent.VoteWhereInput, data []string) error {
	for _, v := range data {
		obj.UserIDNotIn = append(obj.UserIDNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UserIDGt is the resolver for the userIDGT field.
func (r *voteWhereInputResolver) UserIDGt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGt - userIDGT"))
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *voteWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDGte - userIDGTE"))
}

// UserIDLt is the resolver for the userIDLT field.
func (r *voteWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLt - userIDLT"))
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *voteWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.VoteWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDLte - userIDLTE"))
}

// Card returns CardResolver implementation.
func (r *Resolver) Card() CardResolver { return &cardResolver{r} }

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// GameUser returns GameUserResolver implementation.
func (r *Resolver) GameUser() GameUserResolver { return &gameUserResolver{r} }

// Mission returns MissionResolver implementation.
func (r *Resolver) Mission() MissionResolver { return &missionResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Record returns RecordResolver implementation.
func (r *Resolver) Record() RecordResolver { return &recordResolver{r} }

// Room returns RoomResolver implementation.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

// RoomUser returns RoomUserResolver implementation.
func (r *Resolver) RoomUser() RoomUserResolver { return &roomUserResolver{r} }

// Squad returns SquadResolver implementation.
func (r *Resolver) Squad() SquadResolver { return &squadResolver{r} }

// Vote returns VoteResolver implementation.
func (r *Resolver) Vote() VoteResolver { return &voteResolver{r} }

// CardWhereInput returns CardWhereInputResolver implementation.
func (r *Resolver) CardWhereInput() CardWhereInputResolver { return &cardWhereInputResolver{r} }

// CreateCardInput returns CreateCardInputResolver implementation.
func (r *Resolver) CreateCardInput() CreateCardInputResolver { return &createCardInputResolver{r} }

// CreateGameInput returns CreateGameInputResolver implementation.
func (r *Resolver) CreateGameInput() CreateGameInputResolver { return &createGameInputResolver{r} }

// CreateGameUserInput returns CreateGameUserInputResolver implementation.
func (r *Resolver) CreateGameUserInput() CreateGameUserInputResolver {
	return &createGameUserInputResolver{r}
}

// CreateMissionInput returns CreateMissionInputResolver implementation.
func (r *Resolver) CreateMissionInput() CreateMissionInputResolver {
	return &createMissionInputResolver{r}
}

// CreateRecordInput returns CreateRecordInputResolver implementation.
func (r *Resolver) CreateRecordInput() CreateRecordInputResolver {
	return &createRecordInputResolver{r}
}

// CreateRoomInput returns CreateRoomInputResolver implementation.
func (r *Resolver) CreateRoomInput() CreateRoomInputResolver { return &createRoomInputResolver{r} }

// CreateRoomUserInput returns CreateRoomUserInputResolver implementation.
func (r *Resolver) CreateRoomUserInput() CreateRoomUserInputResolver {
	return &createRoomUserInputResolver{r}
}

// CreateSquadInput returns CreateSquadInputResolver implementation.
func (r *Resolver) CreateSquadInput() CreateSquadInputResolver { return &createSquadInputResolver{r} }

// CreateVoteInput returns CreateVoteInputResolver implementation.
func (r *Resolver) CreateVoteInput() CreateVoteInputResolver { return &createVoteInputResolver{r} }

// GameUserWhereInput returns GameUserWhereInputResolver implementation.
func (r *Resolver) GameUserWhereInput() GameUserWhereInputResolver {
	return &gameUserWhereInputResolver{r}
}

// GameWhereInput returns GameWhereInputResolver implementation.
func (r *Resolver) GameWhereInput() GameWhereInputResolver { return &gameWhereInputResolver{r} }

// MissionWhereInput returns MissionWhereInputResolver implementation.
func (r *Resolver) MissionWhereInput() MissionWhereInputResolver {
	return &missionWhereInputResolver{r}
}

// RecordWhereInput returns RecordWhereInputResolver implementation.
func (r *Resolver) RecordWhereInput() RecordWhereInputResolver { return &recordWhereInputResolver{r} }

// RoomUserWhereInput returns RoomUserWhereInputResolver implementation.
func (r *Resolver) RoomUserWhereInput() RoomUserWhereInputResolver {
	return &roomUserWhereInputResolver{r}
}

// RoomWhereInput returns RoomWhereInputResolver implementation.
func (r *Resolver) RoomWhereInput() RoomWhereInputResolver { return &roomWhereInputResolver{r} }

// SquadWhereInput returns SquadWhereInputResolver implementation.
func (r *Resolver) SquadWhereInput() SquadWhereInputResolver { return &squadWhereInputResolver{r} }

// UpdateCardInput returns UpdateCardInputResolver implementation.
func (r *Resolver) UpdateCardInput() UpdateCardInputResolver { return &updateCardInputResolver{r} }

// UpdateGameInput returns UpdateGameInputResolver implementation.
func (r *Resolver) UpdateGameInput() UpdateGameInputResolver { return &updateGameInputResolver{r} }

// UpdateGameUserInput returns UpdateGameUserInputResolver implementation.
func (r *Resolver) UpdateGameUserInput() UpdateGameUserInputResolver {
	return &updateGameUserInputResolver{r}
}

// UpdateMissionInput returns UpdateMissionInputResolver implementation.
func (r *Resolver) UpdateMissionInput() UpdateMissionInputResolver {
	return &updateMissionInputResolver{r}
}

// UpdateRecordInput returns UpdateRecordInputResolver implementation.
func (r *Resolver) UpdateRecordInput() UpdateRecordInputResolver {
	return &updateRecordInputResolver{r}
}

// UpdateRoomInput returns UpdateRoomInputResolver implementation.
func (r *Resolver) UpdateRoomInput() UpdateRoomInputResolver { return &updateRoomInputResolver{r} }

// UpdateRoomUserInput returns UpdateRoomUserInputResolver implementation.
func (r *Resolver) UpdateRoomUserInput() UpdateRoomUserInputResolver {
	return &updateRoomUserInputResolver{r}
}

// UpdateSquadInput returns UpdateSquadInputResolver implementation.
func (r *Resolver) UpdateSquadInput() UpdateSquadInputResolver { return &updateSquadInputResolver{r} }

// UpdateVoteInput returns UpdateVoteInputResolver implementation.
func (r *Resolver) UpdateVoteInput() UpdateVoteInputResolver { return &updateVoteInputResolver{r} }

// VoteWhereInput returns VoteWhereInputResolver implementation.
func (r *Resolver) VoteWhereInput() VoteWhereInputResolver { return &voteWhereInputResolver{r} }

type cardResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type gameUserResolver struct{ *Resolver }
type missionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type recordResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
type squadResolver struct{ *Resolver }
type voteResolver struct{ *Resolver }
type cardWhereInputResolver struct{ *Resolver }
type createCardInputResolver struct{ *Resolver }
type createGameInputResolver struct{ *Resolver }
type createGameUserInputResolver struct{ *Resolver }
type createMissionInputResolver struct{ *Resolver }
type createRecordInputResolver struct{ *Resolver }
type createRoomInputResolver struct{ *Resolver }
type createRoomUserInputResolver struct{ *Resolver }
type createSquadInputResolver struct{ *Resolver }
type createVoteInputResolver struct{ *Resolver }
type gameUserWhereInputResolver struct{ *Resolver }
type gameWhereInputResolver struct{ *Resolver }
type missionWhereInputResolver struct{ *Resolver }
type recordWhereInputResolver struct{ *Resolver }
type roomUserWhereInputResolver struct{ *Resolver }
type roomWhereInputResolver struct{ *Resolver }
type squadWhereInputResolver struct{ *Resolver }
type updateCardInputResolver struct{ *Resolver }
type updateGameInputResolver struct{ *Resolver }
type updateGameUserInputResolver struct{ *Resolver }
type updateMissionInputResolver struct{ *Resolver }
type updateRecordInputResolver struct{ *Resolver }
type updateRoomInputResolver struct{ *Resolver }
type updateRoomUserInputResolver struct{ *Resolver }
type updateSquadInputResolver struct{ *Resolver }
type updateVoteInputResolver struct{ *Resolver }
type voteWhereInputResolver struct{ *Resolver }
