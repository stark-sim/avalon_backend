package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/stark-sim/avalon_backend/pkg/ent"
)

// ID is the resolver for the id field.
func (r *cardResolver) ID(ctx context.Context, obj *ent.Card) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *gameResolver) ID(ctx context.Context, obj *ent.Game) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *gameUserResolver) ID(ctx context.Context, obj *ent.GameUser) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
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
	panic(fmt.Errorf("not implemented: Cards - cards"))
}

// Games is the resolver for the games field.
func (r *queryResolver) Games(ctx context.Context) ([]*ent.Game, error) {
	panic(fmt.Errorf("not implemented: Games - games"))
}

// GameUsers is the resolver for the gameUsers field.
func (r *queryResolver) GameUsers(ctx context.Context) ([]*ent.GameUser, error) {
	panic(fmt.Errorf("not implemented: GameUsers - gameUsers"))
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context) ([]*ent.Room, error) {
	panic(fmt.Errorf("not implemented: Rooms - rooms"))
}

// RoomUsers is the resolver for the roomUsers field.
func (r *queryResolver) RoomUsers(ctx context.Context) ([]*ent.RoomUser, error) {
	panic(fmt.Errorf("not implemented: RoomUsers - roomUsers"))
}

// ID is the resolver for the id field.
func (r *roomResolver) ID(ctx context.Context, obj *ent.Room) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *roomUserResolver) ID(ctx context.Context, obj *ent.RoomUser) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *cardWhereInputResolver) ID(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *cardWhereInputResolver) IDNeq(ctx context.Context, obj *ent.CardWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *cardWhereInputResolver) IDIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *cardWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.CardWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
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

// GameUserIDs is the resolver for the gameUserIDs field.
func (r *createGameInputResolver) GameUserIDs(ctx context.Context, obj *ent.CreateGameInput, data []string) error {
	panic(fmt.Errorf("not implemented: GameUserIDs - gameUserIDs"))
}

// GameID is the resolver for the gameID field.
func (r *createGameUserInputResolver) GameID(ctx context.Context, obj *ent.CreateGameUserInput, data string) error {
	panic(fmt.Errorf("not implemented: GameID - gameID"))
}

// RoomUserIDs is the resolver for the roomUserIDs field.
func (r *createRoomInputResolver) RoomUserIDs(ctx context.Context, obj *ent.CreateRoomInput, data []string) error {
	panic(fmt.Errorf("not implemented: RoomUserIDs - roomUserIDs"))
}

// RoomID is the resolver for the roomID field.
func (r *createRoomUserInputResolver) RoomID(ctx context.Context, obj *ent.CreateRoomUserInput, data string) error {
	panic(fmt.Errorf("not implemented: RoomID - roomID"))
}

// ID is the resolver for the id field.
func (r *gameUserWhereInputResolver) ID(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *gameUserWhereInputResolver) IDNeq(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *gameUserWhereInputResolver) IDIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *gameUserWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.GameUserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *gameUserWhereInputResolver) IDGt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *gameUserWhereInputResolver) IDGte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *gameUserWhereInputResolver) IDLt(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *gameUserWhereInputResolver) IDLte(ctx context.Context, obj *ent.GameUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ID is the resolver for the id field.
func (r *gameWhereInputResolver) ID(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *gameWhereInputResolver) IDNeq(ctx context.Context, obj *ent.GameWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *gameWhereInputResolver) IDIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *gameWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.GameWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
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

// ID is the resolver for the id field.
func (r *roomUserWhereInputResolver) ID(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *roomUserWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoomUserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *roomUserWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *roomUserWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoomUserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
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

// ID is the resolver for the id field.
func (r *roomWhereInputResolver) ID(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *roomWhereInputResolver) IDNeq(ctx context.Context, obj *ent.RoomWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *roomWhereInputResolver) IDIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *roomWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.RoomWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
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

// AddGameUserIDs is the resolver for the addGameUserIDs field.
func (r *updateGameInputResolver) AddGameUserIDs(ctx context.Context, obj *ent.UpdateGameInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddGameUserIDs - addGameUserIDs"))
}

// RemoveGameUserIDs is the resolver for the removeGameUserIDs field.
func (r *updateGameInputResolver) RemoveGameUserIDs(ctx context.Context, obj *ent.UpdateGameInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveGameUserIDs - removeGameUserIDs"))
}

// GameID is the resolver for the gameID field.
func (r *updateGameUserInputResolver) GameID(ctx context.Context, obj *ent.UpdateGameUserInput, data *string) error {
	panic(fmt.Errorf("not implemented: GameID - gameID"))
}

// AddRoomUserIDs is the resolver for the addRoomUserIDs field.
func (r *updateRoomInputResolver) AddRoomUserIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddRoomUserIDs - addRoomUserIDs"))
}

// RemoveRoomUserIDs is the resolver for the removeRoomUserIDs field.
func (r *updateRoomInputResolver) RemoveRoomUserIDs(ctx context.Context, obj *ent.UpdateRoomInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveRoomUserIDs - removeRoomUserIDs"))
}

// RoomID is the resolver for the roomID field.
func (r *updateRoomUserInputResolver) RoomID(ctx context.Context, obj *ent.UpdateRoomUserInput, data *string) error {
	panic(fmt.Errorf("not implemented: RoomID - roomID"))
}

// Card returns CardResolver implementation.
func (r *Resolver) Card() CardResolver { return &cardResolver{r} }

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// GameUser returns GameUserResolver implementation.
func (r *Resolver) GameUser() GameUserResolver { return &gameUserResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Room returns RoomResolver implementation.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

// RoomUser returns RoomUserResolver implementation.
func (r *Resolver) RoomUser() RoomUserResolver { return &roomUserResolver{r} }

// CardWhereInput returns CardWhereInputResolver implementation.
func (r *Resolver) CardWhereInput() CardWhereInputResolver { return &cardWhereInputResolver{r} }

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
type queryResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
type roomUserResolver struct{ *Resolver }
type cardWhereInputResolver struct{ *Resolver }
type createGameInputResolver struct{ *Resolver }
type createGameUserInputResolver struct{ *Resolver }
type createRoomInputResolver struct{ *Resolver }
type createRoomUserInputResolver struct{ *Resolver }
type gameUserWhereInputResolver struct{ *Resolver }
type gameWhereInputResolver struct{ *Resolver }
type roomUserWhereInputResolver struct{ *Resolver }
type roomWhereInputResolver struct{ *Resolver }
type updateGameInputResolver struct{ *Resolver }
type updateGameUserInputResolver struct{ *Resolver }
type updateRoomInputResolver struct{ *Resolver }
type updateRoomUserInputResolver struct{ *Resolver }
