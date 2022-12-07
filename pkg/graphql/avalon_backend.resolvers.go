package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	"github.com/stark-sim/avalon_backend/tools"
)

// ID is the resolver for the id field.
func (r *cardResolver) ID(ctx context.Context, obj *ent.Card) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// ID is the resolver for the id field.
func (r *gameResolver) ID(ctx context.Context, obj *ent.Game) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// ID is the resolver for the id field.
func (r *gameUserResolver) ID(ctx context.Context, obj *ent.GameUser) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
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
	panic(fmt.Errorf("not implemented: Node - node"))
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

// ID is the resolver for the id field.
func (r *roomUserResolver) ID(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// RoomID is the resolver for the roomID field.
func (r *roomUserResolver) RoomID(ctx context.Context, obj *ent.RoomUser) (string, error) {
	return strconv.FormatInt(obj.RoomID, 10), nil
}

// User is the resolver for the user field.
func (r *roomUserResolver) User(ctx context.Context, obj *ent.RoomUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
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

// GameUserIDs is the resolver for the gameUserIDs field.
func (r *createCardInputResolver) GameUserIDs(ctx context.Context, obj *ent.CreateCardInput, data []string) error {
	for _, v := range data {
		obj.GameUserIDs = append(obj.GameUserIDs, tools.StringToInt64(v))
	}
	return nil
}

// GameUserIDs is the resolver for the gameUserIDs field.
func (r *createGameInputResolver) GameUserIDs(ctx context.Context, obj *ent.CreateGameInput, data []string) error {
	for _, v := range data {
		obj.GameUserIDs = append(obj.GameUserIDs, tools.StringToInt64(v))
	}
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

// RoomUserIDs is the resolver for the roomUserIDs field.
func (r *createRoomInputResolver) RoomUserIDs(ctx context.Context, obj *ent.CreateRoomInput, data []string) error {
	for _, v := range data {
		obj.RoomUserIDs = append(obj.RoomUserIDs, tools.StringToInt64(v))
	}
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
	panic(fmt.Errorf("not implemented: NumberNeq - numberNEQ"))
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
