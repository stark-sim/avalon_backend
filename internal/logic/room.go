package logic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/internal/db"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
	"github.com/stark-sim/avalon_backend/tools"
	"github.com/stark-sim/avalon_backend/tools/cache"
)

func JoinRoom(ctx context.Context, dbClient *ent.Client, roomID, userID int64) (*ent.RoomUser, error) {
	var roomUser *ent.RoomUser
	if err := db.WithTx(ctx, dbClient, func(tx *ent.Tx) error {
		// 加入前检查 房间 是否关闭
		_, err := tx.Room.Query().Where(room.ID(roomID), room.DeletedAt(tools.ZeroTime), room.Closed(false)).First(ctx)
		if ent.IsNotFound(err) {
			return errors.New("room no exist")
		} else if err != nil {
			logrus.Errorf("query room before join room: %v", err)
			return err
		}
		// 加入前检查自己是否在其他房间
		if _, err = tx.RoomUser.Query().Where(roomuser.UserID(userID), roomuser.DeletedAt(tools.ZeroTime), roomuser.RoomIDNEQ(roomID)).First(ctx); err != nil {
			if !ent.IsNotFound(err) {
				// 预期外错误
				logrus.Errorf("check if user in another room already: %v", err)
				return err
			}
		} else {
			// 成功找到数据，说明在别的房间，就是出错
			return errors.New("user be in another room")
		}
		// 中间表调整软删除字段来代替创建和删除
		// 需要在 tx Commit 之前把之后有可能会拿的 Room 先拿出来
		roomUser, err = tx.RoomUser.Query().Where(roomuser.RoomID(roomID), roomuser.UserID(userID)).WithRoom().First(ctx)
		if ent.IsNotFound(err) {
			redisClient := cache.NewRedisClient()
			defer redisClient.Close()
			// 要插入 RoomUser 的话需要 Room 信号量
			if err = redisClient.WaitRoomMutex(ctx, roomID); err != nil {
				return err
			}
			defer func(redisClient *cache.RedisClient, ctx context.Context, roomID int64) {
				if err = redisClient.ReleaseRoomMutex(ctx, roomID); err != nil {
					return
				}
			}(redisClient, ctx, roomID)
			// 添加前查看自己是不是第一个加入房间的人，如果是，自己是房主
			// 不需要区分检查软删除，因为所有人退出时，房间会关闭
			hostTaken, err := tx.RoomUser.Query().Where(roomuser.RoomID(roomID)).Exist(ctx)
			if err != nil {
				logrus.Errorf("error at check if room is empty: %v", err)
				return err
			}
			var newRoomUser *ent.RoomUser
			if hostTaken {
				newRoomUser, err = tx.RoomUser.
					Create().
					SetUserID(userID).
					SetRoomID(roomID).
					Save(ctx)
			} else {
				newRoomUser, err = tx.RoomUser.
					Create().
					SetUserID(userID).
					SetRoomID(roomID).
					SetHost(true).
					Save(ctx)
			}
			if err != nil {
				logrus.Errorf("error at create roomUser: %v", err)
				return err
			}
			// 有可能要带出 Room 的信息
			newRoomUser, err = tx.RoomUser.Query().Where(roomuser.ID(newRoomUser.ID)).WithRoom().First(ctx)
			if err != nil {
				return err
			}
			roomUser = newRoomUser
			return nil
		} else {
			// 不需要新增数据库，修改软删除字段就好
			roomUser, err = tx.RoomUser.UpdateOneID(roomUser.ID).SetDeletedAt(tools.ZeroTime).Save(ctx)
			if err != nil {
				logrus.Errorf("update gameUser %d deleted_at when joinRoom: %v", roomUser.ID, err)
				return err
			}
			return nil
		}
	}); err != nil {
		return nil, err
	}
	return roomUser, nil
}
