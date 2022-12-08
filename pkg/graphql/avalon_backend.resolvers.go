package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"strconv"

	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"github.com/stark-sim/avalon_backend/tools"
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
	return &model.User{ID: strconv.FormatInt(obj.UserID, 10)}, nil
}

// CreateRoom is the resolver for the createRoom field.
func (r *mutationResolver) CreateRoom(ctx context.Context, req ent.CreateRoomInput) (*ent.Room, error) {
	return r.client.Room.Create().SetInput(req).Save(ctx)
}

// JoinRoom is the resolver for the joinRoom field.
func (r *mutationResolver) JoinRoom(ctx context.Context, req ent.CreateRoomUserInput) (*ent.RoomUser, error) {
	return r.client.RoomUser.Create().SetInput(req).Save(ctx)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	tempID := tools.StringToInt64(id)
	_user, err := r.client.Game.Query().Where(game.ID(tempID), game.DeletedAtEQ(tools.ZeroTime)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_role, err := r.client.Room.Query().Where(room.ID(tempID), room.DeletedAtEQ(tools.ZeroTime)).First(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					userRole, err := r.client.Card.Query().Where(card.ID(tempID), card.DeletedAtEQ(tools.ZeroTime)).First(ctx)
					if err != nil {
						return nil, err
					}
					return userRole, nil
				}
				return nil, err
			}
			return _role, nil
		}
		return nil, err
	}
	return _user, nil
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	res := make([]ent.Noder, 0)
	tempIDs := make([]int64, len(ids))
	for _, v := range ids {
		tempIDs = append(tempIDs, tools.StringToInt64(v))
	}
	games, err := r.client.Game.Query().Where(game.IDIn(tempIDs...), game.DeletedAtEQ(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range games {
		res = append(res, v)
	}
	rooms, err := r.client.Room.Query().Where(room.IDIn(tempIDs...), room.DeletedAtEQ(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range rooms {
		res = append(res, v)
	}
	cards, err := r.client.Card.Query().Where(card.IDIn(tempIDs...), card.DeletedAtEQ(tools.ZeroTime)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range cards {
		res = append(res, v)
	}
	// return
	return res, nil
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

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context) ([]*ent.Room, error) {
	return r.client.Room.Query().All(ctx)
}

// RoomUsers is the resolver for the roomUsers field.
func (r *queryResolver) RoomUsers(ctx context.Context) ([]*ent.RoomUser, error) {
	return r.client.RoomUser.Query().All(ctx)
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
	return &model.User{ID: strconv.FormatInt(obj.UserID, 10)}, nil
}

// ID is the resolver for the id field.
func (r *cardWhereInputResolver) ID(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.ID = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDNeq is the resolver for the idNEQ field.
func (r *cardWhereInputResolver) IDNeq(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data == nil {
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
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDGte is the resolver for the idGTE field.
func (r *cardWhereInputResolver) IDGte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLt is the resolver for the idLT field.
func (r *cardWhereInputResolver) IDLt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLte is the resolver for the idLTE field.
func (r *cardWhereInputResolver) IDLte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *cardWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *cardWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *cardWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
		obj.UpdatedByIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *cardWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *cardWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *cardWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *cardWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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

// GameUserIDs is the resolver for the gameUserIDs field.
func (r *createGameInputResolver) GameUserIDs(ctx context.Context, obj *ent.CreateGameInput, data []string) error {
	for _, v := range data {
		obj.GameUserIDs = append(obj.GameUserIDs, tools.StringToInt64(v))
	}
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
		obj.UpdatedByIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *gameUserWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *gameUserWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *gameUserWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *gameUserWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *gameUserWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDLt is the resolver for the userIDLT field.
func (r *gameUserWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *gameUserWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
		tempNum := uint8(*data)
		obj.NumberGT = &tempNum
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberGte is the resolver for the numberGTE field.
func (r *gameUserWhereInputResolver) NumberGte(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNum := uint8(*data)
		obj.NumberGTE = &tempNum
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberLt is the resolver for the numberLT field.
func (r *gameUserWhereInputResolver) NumberLt(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNum := uint8(*data)
		obj.NumberLT = &tempNum
		return nil
	} else {
		return errors.New("null")
	}
}

// NumberLte is the resolver for the numberLTE field.
func (r *gameUserWhereInputResolver) NumberLte(ctx context.Context, obj *ent.GameUserWhereInput, data *int) error {
	if data != nil {
		tempNum := uint8(*data)
		obj.NumberLTE = &tempNum
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
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDGte is the resolver for the idGTE field.
func (r *gameWhereInputResolver) IDGte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLt is the resolver for the idLT field.
func (r *gameWhereInputResolver) IDLt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLte is the resolver for the idLTE field.
func (r *gameWhereInputResolver) IDLte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *gameWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *gameWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *gameWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
		obj.UpdatedByIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *gameWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *gameWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *gameWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDGte is the resolver for the idGTE field.
func (r *roomUserWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLt is the resolver for the idLT field.
func (r *roomUserWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLte is the resolver for the idLTE field.
func (r *roomUserWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *roomUserWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *roomUserWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *roomUserWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
		obj.UpdatedByIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *roomUserWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *roomUserWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *roomUserWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLte is the resolver for the updatedByLTE field.
func (r *roomUserWhereInputResolver) UpdatedByLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDGte is the resolver for the userIDGTE field.
func (r *roomUserWhereInputResolver) UserIDGte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDLt is the resolver for the userIDLT field.
func (r *roomUserWhereInputResolver) UserIDLt(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UserIDLte is the resolver for the userIDLTE field.
func (r *roomUserWhereInputResolver) UserIDLte(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UserIDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDGte is the resolver for the idGTE field.
func (r *roomWhereInputResolver) IDGte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLt is the resolver for the idLT field.
func (r *roomWhereInputResolver) IDLt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// IDLte is the resolver for the idLTE field.
func (r *roomWhereInputResolver) IDLte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data == nil {
		tempID := tools.StringToInt64(*data)
		obj.IDLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByGte is the resolver for the createdByGTE field.
func (r *roomWhereInputResolver) CreatedByGte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLt is the resolver for the createdByLT field.
func (r *roomWhereInputResolver) CreatedByLt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// CreatedByLte is the resolver for the createdByLTE field.
func (r *roomWhereInputResolver) CreatedByLte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.CreatedByLTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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
		obj.UpdatedByIn = append(obj.UpdatedByNotIn, tools.StringToInt64(v))
	}
	return nil
}

// UpdatedByGt is the resolver for the updatedByGT field.
func (r *roomWhereInputResolver) UpdatedByGt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByGte is the resolver for the updatedByGTE field.
func (r *roomWhereInputResolver) UpdatedByGte(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByGTE = &tempID
		return nil
	} else {
		return errors.New("null")
	}
}

// UpdatedByLt is the resolver for the updatedByLT field.
func (r *roomWhereInputResolver) UpdatedByLt(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	if data != nil {
		tempID := tools.StringToInt64(*data)
		obj.UpdatedByLT = &tempID
		return nil
	} else {
		return errors.New("null")
	}
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

// Card returns CardResolver implementation.
func (r *Resolver) Card() CardResolver { return &cardResolver{r} }

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// GameUser returns GameUserResolver implementation.
func (r *Resolver) GameUser() GameUserResolver { return &gameUserResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Room returns RoomResolver implementation.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

// RoomUser returns RoomUserResolver implementation.
func (r *Resolver) RoomUser() RoomUserResolver { return &roomUserResolver{r} }

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

// CreateRoomInput returns CreateRoomInputResolver implementation.
func (r *Resolver) CreateRoomInput() CreateRoomInputResolver { return &createRoomInputResolver{r} }

// CreateRoomUserInput returns CreateRoomUserInputResolver implementation.
func (r *Resolver) CreateRoomUserInput() CreateRoomUserInputResolver {
	return &createRoomUserInputResolver{r}
}

// GameUserWhereInput returns GameUserWhereInputResolver implementation.
func (r *Resolver) GameUserWhereInput() GameUserWhereInputResolver {
	return &gameUserWhereInputResolver{r}
}

// GameWhereInput returns GameWhereInputResolver implementation.
func (r *Resolver) GameWhereInput() GameWhereInputResolver { return &gameWhereInputResolver{r} }

// RoomUserWhereInput returns RoomUserWhereInputResolver implementation.
func (r *Resolver) RoomUserWhereInput() RoomUserWhereInputResolver {
	return &roomUserWhereInputResolver{r}
}

// RoomWhereInput returns RoomWhereInputResolver implementation.
func (r *Resolver) RoomWhereInput() RoomWhereInputResolver { return &roomWhereInputResolver{r} }

// UpdateCardInput returns UpdateCardInputResolver implementation.
func (r *Resolver) UpdateCardInput() UpdateCardInputResolver { return &updateCardInputResolver{r} }

// UpdateGameInput returns UpdateGameInputResolver implementation.
func (r *Resolver) UpdateGameInput() UpdateGameInputResolver { return &updateGameInputResolver{r} }

// UpdateGameUserInput returns UpdateGameUserInputResolver implementation.
func (r *Resolver) UpdateGameUserInput() UpdateGameUserInputResolver {
	return &updateGameUserInputResolver{r}
}

// UpdateRoomInput returns UpdateRoomInputResolver implementation.
func (r *Resolver) UpdateRoomInput() UpdateRoomInputResolver { return &updateRoomInputResolver{r} }

// UpdateRoomUserInput returns UpdateRoomUserInputResolver implementation.
func (r *Resolver) UpdateRoomUserInput() UpdateRoomUserInputResolver {
	return &updateRoomUserInputResolver{r}
}

type cardResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type gameUserResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
type cardWhereInputResolver struct{ *Resolver }
type createCardInputResolver struct{ *Resolver }
type createGameInputResolver struct{ *Resolver }
type createGameUserInputResolver struct{ *Resolver }
type createRoomInputResolver struct{ *Resolver }
type createRoomUserInputResolver struct{ *Resolver }
type gameUserWhereInputResolver struct{ *Resolver }
type gameWhereInputResolver struct{ *Resolver }
type roomUserWhereInputResolver struct{ *Resolver }
type roomWhereInputResolver struct{ *Resolver }
type updateCardInputResolver struct{ *Resolver }
type updateGameInputResolver struct{ *Resolver }
type updateGameUserInputResolver struct{ *Resolver }
type updateRoomInputResolver struct{ *Resolver }
type updateRoomUserInputResolver struct{ *Resolver }
