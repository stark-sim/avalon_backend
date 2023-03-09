package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/internal/logic"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/squad"
	"github.com/stark-sim/avalon_backend/pkg/ent/vote"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"github.com/stark-sim/avalon_backend/tools"
	"github.com/stark-sim/avalon_backend/tools/cache"
)

// User is the resolver for the user field.
func (r *gameUserResolver) User(ctx context.Context, obj *ent.GameUser) (*model.User, error) {
	user, err := GetUserAtResolver(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
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
	// 加入前检查 自己 是否在 其他房间
	_, err = tx.RoomUser.Query().Where(roomuser.UserID(req.UserID), roomuser.DeletedAt(tools.ZeroTime), roomuser.RoomIDNEQ(req.RoomID)).First(ctx)
	if !ent.IsNotFound(err) {
		return nil, errors.New("user be in another room")
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
		// 添加前查看自己是不是第一个加入房间的人，如果是，自己是房主
		notHost, err := tx.RoomUser.Query().Where(roomuser.RoomID(req.RoomID)).Exist(ctx)
		if err != nil {
			logrus.Errorf("error at check if room is empty: %v", err)
			return nil, err
		}
		var newRoomUser *ent.RoomUser
		if notHost {
			newRoomUser, err = tx.RoomUser.
				Create().
				SetUserID(req.UserID).
				SetRoomID(req.RoomID).
				Save(ctx)
		} else {
			newRoomUser, err = tx.RoomUser.
				Create().
				SetUserID(req.UserID).
				SetRoomID(req.RoomID).
				SetHost(true).
				Save(ctx)
		}
		if err != nil {
			logrus.Errorf("error at create roomUser: %v", err)
			return nil, err
		}
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
		roomUser, err = tx.RoomUser.UpdateOneID(roomUser.ID).SetDeletedAt(tools.ZeroTime).Save(ctx)
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
	defer redisClient.Close()
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
		SetHost(false).
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
	} else {
		// 如果还有别人，并且自己是房主，那么要把房主丢给别人
		if roomUser.Host {
			anotherRoomUser, err := tx.RoomUser.
				Query().
				Where(roomuser.RoomID(req.RoomID), roomuser.UserIDNEQ(req.UserID), roomuser.DeletedAt(tools.ZeroTime)).
				First(ctx)
			if err != nil {
				return nil, err
			}
			err = tx.RoomUser.UpdateOne(anotherRoomUser).SetHost(true).Exec(ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
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
func (r *mutationResolver) CreateGame(ctx context.Context, req model.CreateGameRequest) (*ent.Game, error) {
	// 将房间中现有的人加入到一局新游戏里
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	roomID := tools.StringToInt64(req.RoomID)
	// 先检查房间没有在进行游戏
	_room, err := r.client.Room.
		Query().
		Where(room.ID(roomID), room.DeletedAt(tools.ZeroTime), room.GameOn(false)).
		First(ctx)
	if err != nil {
		logrus.Errorf("room %s not exist or gameOn, can't create game", req.RoomID)
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
		SetTheAssassinatedIds([]string{}).
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
		var protected bool
		if i == 4 {
			protected = true
		} else {
			protected = false
		}
		missionCreates[i] = tx.Mission.
			Create().
			SetGameID(_game.ID).
			SetLeaderID(userIDs[i]).
			SetCapacity(logic.GetMissionCapacityByNumAndSeq(playerNum, i+1)).
			SetSequence(uint8(i + 1)).
			SetProtected(protected)
	}
	_, err = tx.Mission.CreateBulk(missionCreates...).Save(ctx)
	if err != nil {
		logrus.Errorf("error at creating missions: %v", err)
		return nil, err
	}
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

// TempPickSquads is the resolver for the tempPickSquads field.
func (r *mutationResolver) TempPickSquads(ctx context.Context, req []*ent.CreateSquadInput) ([]string, error) {
	missionID := ""
	userIDs := make([]string, len(req))
	for i, v := range req {
		if missionID == "" {
			missionID = strconv.FormatInt(v.MissionID, 10)
		} else {
			if strconv.FormatInt(v.MissionID, 10) != missionID {
				return nil, errors.New("missionID not the same")
			}
		}
		userIDs[i] = strconv.FormatInt(v.UserID, 10)
	}
	if missionID == "0" {
		return nil, errors.New("missionID cannot be 0")
	}
	cacheClient := cache.NewRedisClient()
	defer cacheClient.Close()
	err := cacheClient.SetMissionTempPickUserIDs(ctx, missionID, userIDs)
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}

// PickSquads is the resolver for the pickSquads field.
func (r *mutationResolver) PickSquads(ctx context.Context, req []*ent.CreateSquadInput) ([]*ent.Squad, error) {
	// 队长选任务小队人数，选好后任务进去投票阶段，并且为大家创建 Vote
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	// 检查 1. 小队任务的是同一个 2. 小队人数和任务任务相等 3. 小队人员在任务所属游戏中
	missionID := int64(0)
	userIDs := make([]int64, len(req))
	for i, v := range req {
		if missionID == 0 {
			missionID = v.MissionID
		} else if missionID != v.MissionID {
			return nil, errors.New("squads' mission_id not the same")
		}
		userIDs[i] = req[i].UserID
	}
	if missionID == 0 {
		return nil, errors.New("squads' mission_id is 0")
	}
	_mission, err := tx.Mission.
		Query().
		Where(
			mission.ID(missionID),
			mission.DeletedAt(tools.ZeroTime),
			mission.StatusEQ(mission.StatusPicking),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}
	if _mission.Capacity != uint8(len(userIDs)) {
		return nil, errors.New("squad number doesn't match mission's capacity")
	}
	// 通过查目标玩家和这局游戏中的玩家数量和 id 对得上来判断是不是属于这局游戏
	count, err := tx.GameUser.
		Query().
		Where(
			gameuser.GameID(_mission.GameID),
			gameuser.UserIDIn(userIDs...),
			gameuser.DeletedAt(tools.ZeroTime),
		).Count(ctx)
	if err != nil {
		return nil, err
	}
	if count != len(userIDs) {
		return nil, errors.New("squad users not in squad's game")
	}
	// 检查完毕，开始创建 Squad
	squadCreates := make([]*ent.SquadCreate, len(userIDs))
	for i, v := range userIDs {
		squadCreates[i] = tx.Squad.
			Create().
			SetMissionID(missionID).
			SetUserID(v)
	}
	squads, err := tx.Squad.CreateBulk(squadCreates...).Save(ctx)
	if err != nil {
		logrus.Errorf("error at creating squads: %v", err)
		return nil, err
	}
	// 创建 Vote
	var allUserIDs []struct {
		UserID int64 `json:"user_id"`
	}
	err = tx.GameUser.Query().
		Where(gameuser.GameID(_mission.GameID), gameuser.DeletedAt(tools.ZeroTime)).
		Select(gameuser.FieldUserID).
		Scan(ctx, &allUserIDs)
	voteCreates := make([]*ent.VoteCreate, len(allUserIDs))
	for i, v := range allUserIDs {
		// 如果用户是该任务的 leader，那么该 Vote 已经决定且通过
		if v.UserID == _mission.LeaderID {
			voteCreates[i] = tx.Vote.
				Create().
				SetUserID(v.UserID).
				SetMissionID(_mission.ID).
				SetPass(true).
				SetVoted(true)
		} else {
			voteCreates[i] = tx.Vote.
				Create().
				SetUserID(v.UserID).
				SetMissionID(_mission.ID)
		}
	}
	_, err = tx.Vote.CreateBulk(voteCreates...).Save(ctx)
	if err != nil {
		logrus.Errorf("error at creating votes: %v", err)
		return nil, err
	}
	// 新增 Squad 和 Vote 后别忘了修改 Mission 状态为 voting
	_, err = tx.Mission.UpdateOne(_mission).SetStatus(mission.StatusVoting).Save(ctx)
	if err != nil {
		logrus.Errorf("error at updating mission's status to voting: %v", err)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	return squads, nil
}

// Vote is the resolver for the vote field.
func (r *mutationResolver) Vote(ctx context.Context, req model.VoteRequest) (*ent.Vote, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_vote, err := tx.Vote.UpdateOneID(tools.StringToInt64(req.VoteID)).SetPass(req.Pass).SetVoted(true).Save(ctx)
	// 如果大家都投票完了，那么更改 Mission 的状态
	votes, err := tx.Vote.Query().Where(vote.MissionID(_vote.MissionID), vote.DeletedAt(tools.ZeroTime)).All(ctx)
	allVoted := true
	passCount := 0
	notPassCount := 0
	for _, v := range votes {
		if !v.Voted {
			allVoted = false
			break
		}
		if v.Pass {
			passCount += 1
		} else {
			notPassCount += 1
		}
	}
	if allVoted {
		_mission, err := tx.Mission.Query().Where(mission.ID(_vote.MissionID), mission.DeletedAt(tools.ZeroTime)).First(ctx)
		if err != nil {
			logrus.Errorf("error at querying mission by vote: %v", err)
			return nil, err
		}
		// 判断流局还是继续进行
		if notPassCount >= passCount {
			err = tx.Mission.UpdateOne(_mission).SetStatus(mission.StatusDelayed).Exec(ctx)
			if err != nil {
				logrus.Errorf("error at update mission to status delayed: %v", err)
				return nil, err
			}
			// 流局要创建新的任务，并且后续任务的 Leader 都往后延一人
			// 先把后面的 Mission 找出来
			postMissions, err := tx.Mission.Query().Where(mission.SequenceGT(_mission.Sequence), mission.DeletedAt(tools.ZeroTime)).All(ctx)
			if err != nil {
				return nil, err
			}
			// 再把 GameUser 中的 userID 按 number 找出来
			var inGameUserIDs []struct {
				UserID int64 `json:"user_id"`
			}
			err = tx.GameUser.Query().
				Where(gameuser.GameID(_mission.GameID), gameuser.DeletedAt(tools.ZeroTime)).
				Order(ent.Asc(gameuser.FieldNumber)).
				Select(gameuser.FieldUserID).
				Scan(ctx, &inGameUserIDs)
			if err != nil {
				return nil, err
			}
			// 之后的任务的 leader 往后盐，第一个 leader 作为流局重开局的 leader
			newMissionLeaderID := postMissions[0].LeaderID
			for _, postMission := range postMissions {
				for i, v := range inGameUserIDs {
					if postMission.LeaderID == v.UserID {
						err = tx.Mission.UpdateOne(postMission).SetLeaderID(inGameUserIDs[i+1].UserID).Exec(ctx)
						return nil, err
					}
				}
			}
			err = tx.Mission.Create().
				SetGameID(_mission.GameID).
				SetSequence(_mission.Sequence).
				SetLeaderID(newMissionLeaderID).
				SetCapacity(_mission.Capacity).
				Exec(ctx)
			if err != nil {
				return nil, err
			}
		} else {
			// 没有流局，进入任务执行阶段
			err = tx.Mission.UpdateOne(_mission).SetStatus(mission.StatusActing).Exec(ctx)
			if err != nil {
				logrus.Errorf("error at update mission to status closed: %v", err)
				return nil, err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return _vote, nil
}

// Act is the resolver for the act field.
func (r *mutationResolver) Act(ctx context.Context, req model.ActRequest) (*ent.Squad, error) {
	// 执行任务是否破坏
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_squad, err := tx.Squad.UpdateOneID(tools.StringToInt64(req.SquadID)).SetRat(req.Rat).SetActed(true).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 执行后判断是不是小队成员都执行完了
	allActed := true
	squads, err := tx.Squad.Query().Where(squad.MissionID(_squad.MissionID), squad.DeletedAt(tools.ZeroTime)).All(ctx)
	ratCount := 0
	for _, v := range squads {
		if !v.Acted {
			allActed = false
			break
		}
		if v.Rat {
			ratCount += 1
		}
	}
	// 全部执行完毕，任务结束，判断任务成功与否
	if allActed {
		_mission, err := tx.Mission.Query().Where(mission.ID(_squad.MissionID), mission.DeletedAt(tools.ZeroTime)).First(ctx)
		if err != nil {
			return nil, err
		}
		missionFailed := false
		if ratCount > 0 {
			// 保护轮
			if _mission.Protected && ratCount <= 1 {
				missionFailed = false
			} else {
				missionFailed = true
			}
		}
		err = tx.Mission.UpdateOne(_mission).SetStatus(mission.StatusClosed).SetFailed(missionFailed).Exec(ctx)
		if err != nil {
			return nil, err
		}
		// 如果目前任务失败，则判断有没有一共失败了 3 次，如果有，游戏由红方胜利结束
		if missionFailed {
			failedMissionCount, err := tx.Mission.Query().
				Where(
					mission.GameID(_mission.GameID),
					mission.DeletedAt(tools.ZeroTime),
					mission.StatusEQ(mission.StatusClosed),
					mission.Failed(true),
				).
				Count(ctx)
			if err != nil {
				logrus.Errorf("error at query missions when final act is done: %v", err)
			}
			// 如果失败次数达到三次，则红方胜利，游戏结束
			if failedMissionCount >= 3 {
				// 游戏结束，红方胜利，结束方式为刺杀成功
				_game, err := tx.Game.UpdateOneID(_mission.GameID).
					SetEndBy(game.EndByRed).
					Save(ctx)
				if err != nil {
					return nil, err
				}
				// 游戏结束时，房间需要变为无游戏状态
				err = tx.Room.UpdateOneID(_game.RoomID).
					SetGameOn(false).
					Exec(ctx)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return _squad, err
}

// TempAssassinate is the resolver for the tempAssassinate field.
func (r *mutationResolver) TempAssassinate(ctx context.Context, req model.AssassinateRequest) ([]string, error) {
	// 缓存临时刺杀目标
	cacheClient := cache.NewRedisClient()
	defer cacheClient.Close()
	err := cacheClient.SetGameTempAssassinatedIDs(ctx, req.GameID, req.TheAssassinatedIDs)
	if err != nil {
		return nil, err
	}
	return req.TheAssassinatedIDs, nil
}

// Assassinate is the resolver for the assassinate field.
func (r *mutationResolver) Assassinate(ctx context.Context, req model.AssassinateRequest) (*ent.Game, error) {
	cacheClient := cache.NewRedisClient()
	defer cacheClient.Close()
	// 刺杀时可以删除游戏的暂时刺杀目标缓存
	err := cacheClient.DeleteGameTempAssassinatedIDs(ctx, req.GameID)
	if err != nil {
		return nil, err
	}
	// 进行事务
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_game, err := tx.Game.Query().Where(game.ID(tools.StringToInt64(req.GameID)), game.DeletedAt(tools.ZeroTime)).First(ctx)
	if err != nil {
		return nil, err
	}
	// 把刺杀的游戏玩家找到
	theAssassinatedIDs := make([]int64, len(req.TheAssassinatedIDs))
	for i, theAssassinatedID := range req.TheAssassinatedIDs {
		theAssassinatedIDs[i] = tools.StringToInt64(theAssassinatedID)
	}
	gameUsers, err := tx.GameUser.Query().
		Where(gameuser.GameID(_game.ID), gameuser.DeletedAt(tools.ZeroTime), gameuser.UserIDIn(theAssassinatedIDs...)).
		WithCard().
		All(ctx)
	if err != nil {
		logrus.Errorf("error at query target gameUser when assassinate: %v", err)
		return nil, err
	}
	// 看看杀没杀到梅林
	merlinDead := false
	for _, gameUser := range gameUsers {
		tempCard, err := gameUser.Card(ctx)
		if err != nil {
			logrus.Errorf("error at query gameUsers with card: %v", err)
			return nil, err
		}
		if tempCard.Name == card.NameMerlin {
			// 游戏结束，红方胜利，结束方式为刺杀成功
			_game, err = tx.Game.UpdateOne(_game).
				SetEndBy(game.EndByAssassination).
				SetTheAssassinatedIds(req.TheAssassinatedIDs).
				Save(ctx)
			// 游戏结束时，房间需要变为无游戏状态
			_, err = tx.Room.UpdateOneID(_game.RoomID).
				SetGameOn(false).
				Save(ctx)
			if err != nil {
				logrus.Errorf("error at update game when assassinate: %v", err)
				return nil, err
			}
			// 梅林阵亡
			merlinDead = true
			break
		}
	}
	// 梅林没死
	if !merlinDead {
		// 游戏结束，蓝方获胜
		_game, err = tx.Game.UpdateOne(_game).
			SetEndBy(game.EndByBlue).
			SetTheAssassinatedIds(req.TheAssassinatedIDs).
			Save(ctx)
		// 游戏结束时，房间需要变为无游戏状态
		_, err = tx.Room.UpdateOneID(_game.RoomID).
			SetGameOn(false).
			Save(ctx)
		if err != nil {
			logrus.Errorf("error at update game when assassinate: %v", err)
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return _game, nil
}

// JoinRoomByShortCode is the resolver for the joinRoomByShortCode field.
func (r *mutationResolver) JoinRoomByShortCode(ctx context.Context, req model.JoinRoomInput) (*ent.RoomUser, error) {
	panic(fmt.Errorf("not implemented: JoinRoomByShortCode - joinRoomByShortCode"))
}

// TerminateGame is the resolver for the terminateGame field.
func (r *mutationResolver) TerminateGame(ctx context.Context, req model.GameRequest) (*ent.Game, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	// 游戏修改状态
	_game, err := tx.Game.UpdateOneID(tools.StringToInt64(req.ID)).SetEndBy(game.EndByHand).Save(ctx)
	if err != nil {
		logrus.Errorf("error at terminate game: %v", err)
		return nil, err
	}
	// 房间变回无进行游戏状态
	if err = tx.Room.UpdateOneID(_game.RoomID).SetGameOn(false).Exec(ctx); err != nil {
		logrus.Errorf("error at update room when terminate game: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return _game, nil
}

// GetJoinedRoom is the resolver for the getJoinedRoom field.
func (r *queryResolver) GetJoinedRoom(ctx context.Context, req model.UserRequest) (*ent.Room, error) {
	// 查询 room 没有关闭的，且具有 roomUser 的
	userID := tools.StringToInt64(req.ID)
	_room, err := r.client.Room.Query().
		Where(
			room.DeletedAt(tools.ZeroTime),
			room.Closed(false),
			room.HasRoomUsersWith(
				roomuser.DeletedAt(tools.ZeroTime),
				roomuser.UserID(userID),
			),
		).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		} else {
			logrus.Errorf("error at query user joined room: %v", err)
			return nil, err
		}
	}
	return _room, nil
}

// GetVoteInMission is the resolver for the getVoteInMission field.
func (r *queryResolver) GetVoteInMission(ctx context.Context, req ent.VoteWhereInput) (*ent.Vote, error) {
	userID := req.UserID
	missionID := req.MissionID
	if userID == nil || missionID == nil {
		return nil, errors.New("userID and missionID can't be null")
	}
	// 投好票的也返回，让前端可以通过数据恢复自己的投票状态
	_vote, err := r.client.Vote.Query().Where(vote.UserID(*userID), vote.MissionID(*missionID), vote.DeletedAt(tools.ZeroTime)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		logrus.Errorf("error at query vote in mission: %v", err)
		return nil, err
	} else {
		return _vote, nil
	}
}

// GetSquadInMission is the resolver for the getSquadInMission field.
func (r *queryResolver) GetSquadInMission(ctx context.Context, req ent.SquadWhereInput) (*ent.Squad, error) {
	userID := req.UserID
	missionID := req.MissionID
	if userID == nil || missionID == nil {
		return nil, errors.New("userID and missionID can't be null")
	}
	_squad, err := r.client.Squad.Query().Where(squad.UserID(*userID), squad.MissionID(*missionID), squad.DeletedAt(tools.ZeroTime)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		logrus.Errorf("error at query vote in mission: %v", err)
		return nil, err
	} else {
		return _squad, nil
	}
}

// GetEndedGame is the resolver for the getEndedGame field.
func (r *queryResolver) GetEndedGame(ctx context.Context, req model.GameRequest) (*ent.Game, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	// 先看看 Game 的 Missions 是不是分出结果了
	missions, err := tx.Mission.Query().
		Where(
			mission.GameID(tools.StringToInt64(req.ID)),
			mission.DeletedAt(tools.ZeroTime),
			mission.StatusNEQ(mission.StatusDelayed),
		).
		All(ctx)
	closedCount := 0
	failedCount := 0
	for _, v := range missions {
		if v.Status == mission.StatusClosed {
			closedCount += 1
			if v.Failed {
				failedCount += 1
			}
		}
	}
	// 还没结束的话，返回空
	if failedCount != 3 && closedCount-failedCount < 3 {
		return nil, nil
	}
	// 如果失败达到 3 次，游戏状态直接变成 end_by red
	if failedCount == 3 {
		err = tx.Game.UpdateOneID(tools.StringToInt64(req.ID)).SetEndBy(game.EndByRed).Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	_game, err := tx.Game.Query().
		Where(game.ID(tools.StringToInt64(req.ID)), game.DeletedAt(tools.ZeroTime)).
		WithNamedGameUsers("gameUsers", func(query *ent.GameUserQuery) {
			query.Where(gameuser.HasGameWith(game.EndByNEQ(game.EndByNone))).WithCard()
		}).
		First(ctx)
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return _game, nil
}

// GetVagueGameUsers is the resolver for the getVagueGameUsers field.
func (r *queryResolver) GetVagueGameUsers(ctx context.Context, req model.GameRequest) ([]*ent.GameUser, error) {
	gameUsers, err := r.client.GameUser.Query().
		Where(gameuser.GameID(tools.StringToInt64(req.ID)), gameuser.DeletedAt(tools.ZeroTime)).
		WithCard().
		Order(ent.Asc(gameuser.FieldNumber)).
		All(ctx)
	if err != nil {
		logrus.Errorf("error at query gameUsers when getVagueGameUsers: %v", err)
		return nil, err
	}
	return gameUsers, nil
}

// GetGameUsersByGame is the resolver for the getGameUsersByGame field.
func (r *queryResolver) GetGameUsersByGame(ctx context.Context, req model.GameRequest) ([]*ent.GameUser, error) {
	// 先检查 Game 状态，如果已结束，则返回时允许访问 Card，通过事务来限制 graphql 的自行进一步访问
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_game, err := tx.Game.Query().Where(game.ID(tools.StringToInt64(req.ID))).First(ctx)
	if err != nil {
		logrus.Errorf("error at query game by id: %v", err)
		return nil, err
	}
	var gameUsers []*ent.GameUser
	if _game.EndBy != game.EndByNone {
		gameUsers, err = tx.GameUser.Query().
			Where(gameuser.DeletedAt(tools.ZeroTime), gameuser.GameID(tools.StringToInt64(req.ID))).
			Order(ent.Asc(gameuser.FieldNumber)).
			WithCard().
			All(ctx)
	} else {
		gameUsers, err = tx.GameUser.Query().
			Where(gameuser.DeletedAt(tools.ZeroTime), gameuser.GameID(tools.StringToInt64(req.ID))).
			Order(ent.Asc(gameuser.FieldNumber)).
			All(ctx)
	}
	if err != nil {
		logrus.Errorf("error at query gameUsers by gameID: %v", err)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return gameUsers, nil
}

// GetOnesCardInGame is the resolver for the getOnesCardInGame field.
func (r *queryResolver) GetOnesCardInGame(ctx context.Context, req model.GameUserRequest) (*ent.Card, error) {
	gameUser, err := r.client.GameUser.Query().
		Where(
			gameuser.UserID(tools.StringToInt64(req.UserID)),
			gameuser.GameID(tools.StringToInt64(req.GameID)),
			gameuser.DeletedAt(tools.ZeroTime),
		).
		WithCard(func(cardQuery *ent.CardQuery) {
			cardQuery.Where(card.DeletedAt(tools.ZeroTime))
		}).
		First(ctx)
	if err != nil {
		logrus.Errorf("error at get one's GameUser for Card: %v", err)
		return nil, err
	}
	if gameUser.Edges.Card == nil {
		return nil, errors.New("don't have card for this gameUser")
	}
	return gameUser.Edges.Card, nil
}

// ViewOthersInGame is the resolver for the viewOthersInGame field. 凭本人的身份来返回对应的他人数据
func (r *queryResolver) ViewOthersInGame(ctx context.Context, req model.GameUserRequest) ([]*model.OtherView, error) {
	// 先把人都搜出来
	gameUsers, err := r.client.GameUser.Query().
		Where(
			gameuser.GameID(tools.StringToInt64(req.GameID)),
			gameuser.DeletedAt(tools.ZeroTime),
		).
		WithCard().
		All(ctx)
	if err != nil {
		logrus.Errorf("error at query gameUsers")
		return nil, err
	}
	// 找出自己的身份
	var selfCardName card.Name
	var selfCardRed bool
	for _, gameUser := range gameUsers {
		if gameUser.UserID == tools.StringToInt64(req.UserID) {
			selfCardName = gameUser.Edges.Card.Name
			selfCardRed = gameUser.Edges.Card.Red
			break
		}
	}
	// 准备构建返回数据
	res := make([]*model.OtherView, 0)
	for _, gameUser := range gameUsers {
		switch selfCardName {
		// 梅林看所有，看错莫德雷德
		case card.NameMerlin:
			if gameUser.Edges.Card.Red == false {
				res = append(res, &model.OtherView{
					UserID: strconv.FormatInt(gameUser.UserID, 10),
					Type:   "BLUE",
				})
			} else {
				// 红方中看不到莫德雷德
				if gameUser.Edges.Card.Name == card.NameMordred {
					res = append(res, &model.OtherView{
						UserID: strconv.FormatInt(gameUser.UserID, 10),
						Type:   "BLUE",
					})
				} else {
					res = append(res, &model.OtherView{
						UserID: strconv.FormatInt(gameUser.UserID, 10),
						Type:   "RED",
					})
				}
			}
		// 派西看梅林和莫甘娜，但看不懂
		case card.NamePercival:
			if gameUser.Edges.Card.Name == card.NameMerlin || gameUser.Edges.Card.Name == card.NameMorgana {
				res = append(res, &model.OtherView{
					UserID: strconv.FormatInt(gameUser.UserID, 10),
					Type:   "UNKNOWN",
				})
			}
		// 除奥伯伦外，红方看到红方
		default:
			if selfCardRed && selfCardName != card.NameOberon {
				if gameUser.Edges.Card.Red && gameUser.Edges.Card.Name != card.NameOberon {
					res = append(res, &model.OtherView{
						UserID: strconv.FormatInt(gameUser.UserID, 10),
						Type:   "RED",
					})
				}
			}
		}
	}
	return res, err
}

// User is the resolver for the user field.
func (r *roomUserResolver) User(ctx context.Context, obj *ent.RoomUser) (*model.User, error) {
	user, err := GetUserAtResolver(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// User is the resolver for the user field.
func (r *squadResolver) User(ctx context.Context, obj *ent.Squad) (*model.User, error) {
	fc := graphql.GetFieldContext(ctx)
	logrus.Infof("%s\n", fc.Object)
	return &model.User{
		ID: strconv.FormatInt(obj.UserID, 10),
	}, nil
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
	// 前端目前用这个方法，只需要知道轮到哪个任务在进行，和这些任务的状态，不需要知道 Mission 的 Squad 等后续数据
	ch := make(chan []*ent.Mission)
	// 检查入参 gameID
	gameID := tools.StringToInt64(req.ID)
	_, err := r.client.Game.Query().Where(game.ID(gameID), game.DeletedAt(tools.ZeroTime)).First(ctx)
	if err != nil {
		logrus.Errorf("error at query missions, gameID not exist: %v", err)
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
				time.Sleep(time.Second)
			}
		}
	}()
	return ch, nil
}

// GetAssassinationByGame is the resolver for the getAssassinationByGame field.
func (r *subscriptionResolver) GetAssassinationByGame(ctx context.Context, req model.GameRequest) (<-chan *model.AssassinInfo, error) {
	ch := make(chan *model.AssassinInfo)
	go func() {
		cacheClient, ok := ctx.Value(cache.DefaultClient).(cache.Client)
		if !ok {
			logrus.Errorf("error at get cacheClient when get AssassinationInfo")
			return
		}
		defer cacheClient.Close()
		res := model.AssassinInfo{}
		for {
			_game, err := r.client.Game.Query().
				Where(game.ID(tools.StringToInt64(req.ID)), game.DeletedAt(tools.ZeroTime)).
				First(ctx)
			if err != nil {
				return
			}
			if len(_game.TheAssassinatedIds) != 0 {
				res.TheAssassinatedIDs = _game.TheAssassinatedIds
				res.TempPickedIDs = res.TheAssassinatedIDs
			} else {
				// 从 redis 中获取 刺客暂时选定的 人
				tempAssassinatedIDs, err := cacheClient.GetGameTempAssassinatedIDs(ctx, req.ID)
				if err != nil {
					return
				}
				res.TheAssassinatedIDs = []string{}
				res.TempPickedIDs = tempAssassinatedIDs
			}
			select {
			case ch <- &res:
				time.Sleep(time.Second)
			}
		}
	}()
	return ch, nil
}

// GetGame is the resolver for the getGame field.
func (r *subscriptionResolver) GetGame(ctx context.Context, req model.GameRequest) (<-chan *ent.Game, error) {
	ch := make(chan *ent.Game)
	go func() {
		for {
			_game, err := r.client.Game.Get(ctx, tools.StringToInt64(req.ID))
			if err != nil {
				return
			}
			select {
			case ch <- _game:
				time.Sleep(time.Second)
			}
		}
	}()
	return ch, nil
}

// User is the resolver for the user field.
func (r *voteResolver) User(ctx context.Context, obj *ent.Vote) (*model.User, error) {
	fc := graphql.GetFieldContext(ctx)
	logrus.Infof("%s\n", fc.Object)
	return &model.User{
		ID: strconv.FormatInt(obj.UserID, 10),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
