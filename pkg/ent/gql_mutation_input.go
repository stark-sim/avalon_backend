// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
)

// CreateCardInput represents a mutation input for creating cards.
type CreateCardInput struct {
	CreatedBy   *int64
	UpdatedBy   *int64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Name        *card.Name
	Role        card.Role
	Tale        *string
	Red         *bool
	GameUserIDs []int64
}

// Mutate applies the CreateCardInput on the CardMutation builder.
func (i *CreateCardInput) Mutate(m *CardMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	m.SetRole(i.Role)
	if v := i.Tale; v != nil {
		m.SetTale(*v)
	}
	if v := i.Red; v != nil {
		m.SetRed(*v)
	}
	if v := i.GameUserIDs; len(v) > 0 {
		m.AddGameUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCardInput on the CardCreate builder.
func (c *CardCreate) SetInput(i CreateCardInput) *CardCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCardInput represents a mutation input for updating cards.
type UpdateCardInput struct {
	CreatedBy         *int64
	UpdatedBy         *int64
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
	Name              *card.Name
	Role              *card.Role
	Tale              *string
	Red               *bool
	AddGameUserIDs    []int64
	RemoveGameUserIDs []int64
}

// Mutate applies the UpdateCardInput on the CardMutation builder.
func (i *UpdateCardInput) Mutate(m *CardMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	if v := i.Tale; v != nil {
		m.SetTale(*v)
	}
	if v := i.Red; v != nil {
		m.SetRed(*v)
	}
	if v := i.AddGameUserIDs; len(v) > 0 {
		m.AddGameUserIDs(v...)
	}
	if v := i.RemoveGameUserIDs; len(v) > 0 {
		m.RemoveGameUserIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCardInput on the CardUpdate builder.
func (c *CardUpdate) SetInput(i UpdateCardInput) *CardUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCardInput on the CardUpdateOne builder.
func (c *CardUpdateOne) SetInput(i UpdateCardInput) *CardUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateGameInput represents a mutation input for creating games.
type CreateGameInput struct {
	CreatedBy          *int64
	UpdatedBy          *int64
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	DeletedAt          *time.Time
	Result             *game.Result
	Capacity           *uint8
	TheAssassinatedIds []string
	AssassinChance     *uint8
	Closed             *bool
	GameUserIDs        []int64
	MissionIDs         []int64
	RoomID             int64
}

// Mutate applies the CreateGameInput on the GameMutation builder.
func (i *CreateGameInput) Mutate(m *GameMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Result; v != nil {
		m.SetResult(*v)
	}
	if v := i.Capacity; v != nil {
		m.SetCapacity(*v)
	}
	if v := i.TheAssassinatedIds; v != nil {
		m.SetTheAssassinatedIds(v)
	}
	if v := i.AssassinChance; v != nil {
		m.SetAssassinChance(*v)
	}
	if v := i.Closed; v != nil {
		m.SetClosed(*v)
	}
	if v := i.GameUserIDs; len(v) > 0 {
		m.AddGameUserIDs(v...)
	}
	if v := i.MissionIDs; len(v) > 0 {
		m.AddMissionIDs(v...)
	}
	m.SetRoomID(i.RoomID)
}

// SetInput applies the change-set in the CreateGameInput on the GameCreate builder.
func (c *GameCreate) SetInput(i CreateGameInput) *GameCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateGameInput represents a mutation input for updating games.
type UpdateGameInput struct {
	CreatedBy                *int64
	UpdatedBy                *int64
	UpdatedAt                *time.Time
	DeletedAt                *time.Time
	Result                   *game.Result
	Capacity                 *uint8
	ClearTheAssassinatedIds  bool
	TheAssassinatedIds       []string
	AppendTheAssassinatedIds []string
	AssassinChance           *uint8
	Closed                   *bool
	AddGameUserIDs           []int64
	RemoveGameUserIDs        []int64
	AddMissionIDs            []int64
	RemoveMissionIDs         []int64
	ClearRoom                bool
	RoomID                   *int64
}

// Mutate applies the UpdateGameInput on the GameMutation builder.
func (i *UpdateGameInput) Mutate(m *GameMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Result; v != nil {
		m.SetResult(*v)
	}
	if v := i.Capacity; v != nil {
		m.SetCapacity(*v)
	}
	if i.ClearTheAssassinatedIds {
		m.ClearTheAssassinatedIds()
	}
	if v := i.TheAssassinatedIds; v != nil {
		m.SetTheAssassinatedIds(v)
	}
	if i.AppendTheAssassinatedIds != nil {
		m.AppendTheAssassinatedIds(i.TheAssassinatedIds)
	}
	if v := i.AssassinChance; v != nil {
		m.SetAssassinChance(*v)
	}
	if v := i.Closed; v != nil {
		m.SetClosed(*v)
	}
	if v := i.AddGameUserIDs; len(v) > 0 {
		m.AddGameUserIDs(v...)
	}
	if v := i.RemoveGameUserIDs; len(v) > 0 {
		m.RemoveGameUserIDs(v...)
	}
	if v := i.AddMissionIDs; len(v) > 0 {
		m.AddMissionIDs(v...)
	}
	if v := i.RemoveMissionIDs; len(v) > 0 {
		m.RemoveMissionIDs(v...)
	}
	if i.ClearRoom {
		m.ClearRoom()
	}
	if v := i.RoomID; v != nil {
		m.SetRoomID(*v)
	}
}

// SetInput applies the change-set in the UpdateGameInput on the GameUpdate builder.
func (c *GameUpdate) SetInput(i UpdateGameInput) *GameUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateGameInput on the GameUpdateOne builder.
func (c *GameUpdateOne) SetInput(i UpdateGameInput) *GameUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateGameUserInput represents a mutation input for creating gameusers.
type CreateGameUserInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    int64
	Number    uint8
	GameID    int64
	CardID    int64
}

// Mutate applies the CreateGameUserInput on the GameUserMutation builder.
func (i *CreateGameUserInput) Mutate(m *GameUserMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetUserID(i.UserID)
	m.SetNumber(i.Number)
	m.SetGameID(i.GameID)
	m.SetCardID(i.CardID)
}

// SetInput applies the change-set in the CreateGameUserInput on the GameUserCreate builder.
func (c *GameUserCreate) SetInput(i CreateGameUserInput) *GameUserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateGameUserInput represents a mutation input for updating gameusers.
type UpdateGameUserInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    *int64
	Number    *uint8
	ClearGame bool
	GameID    *int64
	ClearCard bool
	CardID    *int64
}

// Mutate applies the UpdateGameUserInput on the GameUserMutation builder.
func (i *UpdateGameUserInput) Mutate(m *GameUserMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Number; v != nil {
		m.SetNumber(*v)
	}
	if i.ClearGame {
		m.ClearGame()
	}
	if v := i.GameID; v != nil {
		m.SetGameID(*v)
	}
	if i.ClearCard {
		m.ClearCard()
	}
	if v := i.CardID; v != nil {
		m.SetCardID(*v)
	}
}

// SetInput applies the change-set in the UpdateGameUserInput on the GameUserUpdate builder.
func (c *GameUserUpdate) SetInput(i UpdateGameUserInput) *GameUserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateGameUserInput on the GameUserUpdateOne builder.
func (c *GameUserUpdateOne) SetInput(i UpdateGameUserInput) *GameUserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMissionInput represents a mutation input for creating missions.
type CreateMissionInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Sequence  uint8
	Status    *mission.Status
	Failed    *bool
	Capacity  *uint8
	LeaderID  *int64
	Protected *bool
	GameID    int64
	SquadIDs  []int64
	VoteIDs   []int64
}

// Mutate applies the CreateMissionInput on the MissionMutation builder.
func (i *CreateMissionInput) Mutate(m *MissionMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetSequence(i.Sequence)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Failed; v != nil {
		m.SetFailed(*v)
	}
	if v := i.Capacity; v != nil {
		m.SetCapacity(*v)
	}
	if v := i.LeaderID; v != nil {
		m.SetLeaderID(*v)
	}
	if v := i.Protected; v != nil {
		m.SetProtected(*v)
	}
	m.SetGameID(i.GameID)
	if v := i.SquadIDs; len(v) > 0 {
		m.AddSquadIDs(v...)
	}
	if v := i.VoteIDs; len(v) > 0 {
		m.AddVoteIDs(v...)
	}
}

// SetInput applies the change-set in the CreateMissionInput on the MissionCreate builder.
func (c *MissionCreate) SetInput(i CreateMissionInput) *MissionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMissionInput represents a mutation input for updating missions.
type UpdateMissionInput struct {
	CreatedBy      *int64
	UpdatedBy      *int64
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	Sequence       *uint8
	Status         *mission.Status
	Failed         *bool
	Capacity       *uint8
	LeaderID       *int64
	Protected      *bool
	ClearGame      bool
	GameID         *int64
	AddSquadIDs    []int64
	RemoveSquadIDs []int64
	AddVoteIDs     []int64
	RemoveVoteIDs  []int64
}

// Mutate applies the UpdateMissionInput on the MissionMutation builder.
func (i *UpdateMissionInput) Mutate(m *MissionMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Sequence; v != nil {
		m.SetSequence(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Failed; v != nil {
		m.SetFailed(*v)
	}
	if v := i.Capacity; v != nil {
		m.SetCapacity(*v)
	}
	if v := i.LeaderID; v != nil {
		m.SetLeaderID(*v)
	}
	if v := i.Protected; v != nil {
		m.SetProtected(*v)
	}
	if i.ClearGame {
		m.ClearGame()
	}
	if v := i.GameID; v != nil {
		m.SetGameID(*v)
	}
	if v := i.AddSquadIDs; len(v) > 0 {
		m.AddSquadIDs(v...)
	}
	if v := i.RemoveSquadIDs; len(v) > 0 {
		m.RemoveSquadIDs(v...)
	}
	if v := i.AddVoteIDs; len(v) > 0 {
		m.AddVoteIDs(v...)
	}
	if v := i.RemoveVoteIDs; len(v) > 0 {
		m.RemoveVoteIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateMissionInput on the MissionUpdate builder.
func (c *MissionUpdate) SetInput(i UpdateMissionInput) *MissionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMissionInput on the MissionUpdateOne builder.
func (c *MissionUpdateOne) SetInput(i UpdateMissionInput) *MissionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRecordInput represents a mutation input for creating records.
type CreateRecordInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    int64
	Score     *int32
	RoomID    int64
}

// Mutate applies the CreateRecordInput on the RecordMutation builder.
func (i *CreateRecordInput) Mutate(m *RecordMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetUserID(i.UserID)
	if v := i.Score; v != nil {
		m.SetScore(*v)
	}
	m.SetRoomID(i.RoomID)
}

// SetInput applies the change-set in the CreateRecordInput on the RecordCreate builder.
func (c *RecordCreate) SetInput(i CreateRecordInput) *RecordCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRecordInput represents a mutation input for updating records.
type UpdateRecordInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    *int64
	Score     *int32
	ClearRoom bool
	RoomID    *int64
}

// Mutate applies the UpdateRecordInput on the RecordMutation builder.
func (i *UpdateRecordInput) Mutate(m *RecordMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Score; v != nil {
		m.SetScore(*v)
	}
	if i.ClearRoom {
		m.ClearRoom()
	}
	if v := i.RoomID; v != nil {
		m.SetRoomID(*v)
	}
}

// SetInput applies the change-set in the UpdateRecordInput on the RecordUpdate builder.
func (c *RecordUpdate) SetInput(i UpdateRecordInput) *RecordUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRecordInput on the RecordUpdateOne builder.
func (c *RecordUpdateOne) SetInput(i UpdateRecordInput) *RecordUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRoomInput represents a mutation input for creating rooms.
type CreateRoomInput struct {
	CreatedBy   *int64
	UpdatedBy   *int64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Name        *string
	Closed      *bool
	GameOn      *bool
	RoomUserIDs []int64
	GameIDs     []int64
	RecordIDs   []int64
}

// Mutate applies the CreateRoomInput on the RoomMutation builder.
func (i *CreateRoomInput) Mutate(m *RoomMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Closed; v != nil {
		m.SetClosed(*v)
	}
	if v := i.GameOn; v != nil {
		m.SetGameOn(*v)
	}
	if v := i.RoomUserIDs; len(v) > 0 {
		m.AddRoomUserIDs(v...)
	}
	if v := i.GameIDs; len(v) > 0 {
		m.AddGameIDs(v...)
	}
	if v := i.RecordIDs; len(v) > 0 {
		m.AddRecordIDs(v...)
	}
}

// SetInput applies the change-set in the CreateRoomInput on the RoomCreate builder.
func (c *RoomCreate) SetInput(i CreateRoomInput) *RoomCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRoomInput represents a mutation input for updating rooms.
type UpdateRoomInput struct {
	CreatedBy         *int64
	UpdatedBy         *int64
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
	Name              *string
	Closed            *bool
	GameOn            *bool
	AddRoomUserIDs    []int64
	RemoveRoomUserIDs []int64
	AddGameIDs        []int64
	RemoveGameIDs     []int64
	AddRecordIDs      []int64
	RemoveRecordIDs   []int64
}

// Mutate applies the UpdateRoomInput on the RoomMutation builder.
func (i *UpdateRoomInput) Mutate(m *RoomMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Closed; v != nil {
		m.SetClosed(*v)
	}
	if v := i.GameOn; v != nil {
		m.SetGameOn(*v)
	}
	if v := i.AddRoomUserIDs; len(v) > 0 {
		m.AddRoomUserIDs(v...)
	}
	if v := i.RemoveRoomUserIDs; len(v) > 0 {
		m.RemoveRoomUserIDs(v...)
	}
	if v := i.AddGameIDs; len(v) > 0 {
		m.AddGameIDs(v...)
	}
	if v := i.RemoveGameIDs; len(v) > 0 {
		m.RemoveGameIDs(v...)
	}
	if v := i.AddRecordIDs; len(v) > 0 {
		m.AddRecordIDs(v...)
	}
	if v := i.RemoveRecordIDs; len(v) > 0 {
		m.RemoveRecordIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateRoomInput on the RoomUpdate builder.
func (c *RoomUpdate) SetInput(i UpdateRoomInput) *RoomUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRoomInput on the RoomUpdateOne builder.
func (c *RoomUpdateOne) SetInput(i UpdateRoomInput) *RoomUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRoomUserInput represents a mutation input for creating roomusers.
type CreateRoomUserInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    int64
	Host      *bool
	RoomID    int64
}

// Mutate applies the CreateRoomUserInput on the RoomUserMutation builder.
func (i *CreateRoomUserInput) Mutate(m *RoomUserMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetUserID(i.UserID)
	if v := i.Host; v != nil {
		m.SetHost(*v)
	}
	m.SetRoomID(i.RoomID)
}

// SetInput applies the change-set in the CreateRoomUserInput on the RoomUserCreate builder.
func (c *RoomUserCreate) SetInput(i CreateRoomUserInput) *RoomUserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRoomUserInput represents a mutation input for updating roomusers.
type UpdateRoomUserInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    *int64
	Host      *bool
	ClearRoom bool
	RoomID    *int64
}

// Mutate applies the UpdateRoomUserInput on the RoomUserMutation builder.
func (i *UpdateRoomUserInput) Mutate(m *RoomUserMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Host; v != nil {
		m.SetHost(*v)
	}
	if i.ClearRoom {
		m.ClearRoom()
	}
	if v := i.RoomID; v != nil {
		m.SetRoomID(*v)
	}
}

// SetInput applies the change-set in the UpdateRoomUserInput on the RoomUserUpdate builder.
func (c *RoomUserUpdate) SetInput(i UpdateRoomUserInput) *RoomUserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRoomUserInput on the RoomUserUpdateOne builder.
func (c *RoomUserUpdateOne) SetInput(i UpdateRoomUserInput) *RoomUserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateSquadInput represents a mutation input for creating squads.
type CreateSquadInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    int64
	Rat       *bool
	Acted     *bool
	MissionID int64
}

// Mutate applies the CreateSquadInput on the SquadMutation builder.
func (i *CreateSquadInput) Mutate(m *SquadMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetUserID(i.UserID)
	if v := i.Rat; v != nil {
		m.SetRat(*v)
	}
	if v := i.Acted; v != nil {
		m.SetActed(*v)
	}
	m.SetMissionID(i.MissionID)
}

// SetInput applies the change-set in the CreateSquadInput on the SquadCreate builder.
func (c *SquadCreate) SetInput(i CreateSquadInput) *SquadCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateSquadInput represents a mutation input for updating squads.
type UpdateSquadInput struct {
	CreatedBy    *int64
	UpdatedBy    *int64
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	UserID       *int64
	Rat          *bool
	Acted        *bool
	ClearMission bool
	MissionID    *int64
}

// Mutate applies the UpdateSquadInput on the SquadMutation builder.
func (i *UpdateSquadInput) Mutate(m *SquadMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Rat; v != nil {
		m.SetRat(*v)
	}
	if v := i.Acted; v != nil {
		m.SetActed(*v)
	}
	if i.ClearMission {
		m.ClearMission()
	}
	if v := i.MissionID; v != nil {
		m.SetMissionID(*v)
	}
}

// SetInput applies the change-set in the UpdateSquadInput on the SquadUpdate builder.
func (c *SquadUpdate) SetInput(i UpdateSquadInput) *SquadUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateSquadInput on the SquadUpdateOne builder.
func (c *SquadUpdateOne) SetInput(i UpdateSquadInput) *SquadUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateVoteInput represents a mutation input for creating votes.
type CreateVoteInput struct {
	CreatedBy *int64
	UpdatedBy *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	UserID    int64
	Pass      *bool
	Voted     *bool
	MissionID int64
}

// Mutate applies the CreateVoteInput on the VoteMutation builder.
func (i *CreateVoteInput) Mutate(m *VoteMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetUserID(i.UserID)
	if v := i.Pass; v != nil {
		m.SetPass(*v)
	}
	if v := i.Voted; v != nil {
		m.SetVoted(*v)
	}
	m.SetMissionID(i.MissionID)
}

// SetInput applies the change-set in the CreateVoteInput on the VoteCreate builder.
func (c *VoteCreate) SetInput(i CreateVoteInput) *VoteCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateVoteInput represents a mutation input for updating votes.
type UpdateVoteInput struct {
	CreatedBy    *int64
	UpdatedBy    *int64
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	UserID       *int64
	Pass         *bool
	Voted        *bool
	ClearMission bool
	MissionID    *int64
}

// Mutate applies the UpdateVoteInput on the VoteMutation builder.
func (i *UpdateVoteInput) Mutate(m *VoteMutation) {
	if v := i.CreatedBy; v != nil {
		m.SetCreatedBy(*v)
	}
	if v := i.UpdatedBy; v != nil {
		m.SetUpdatedBy(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Pass; v != nil {
		m.SetPass(*v)
	}
	if v := i.Voted; v != nil {
		m.SetVoted(*v)
	}
	if i.ClearMission {
		m.ClearMission()
	}
	if v := i.MissionID; v != nil {
		m.SetMissionID(*v)
	}
}

// SetInput applies the change-set in the UpdateVoteInput on the VoteUpdate builder.
func (c *VoteUpdate) SetInput(i UpdateVoteInput) *VoteUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateVoteInput on the VoteUpdateOne builder.
func (c *VoteUpdateOne) SetInput(i UpdateVoteInput) *VoteUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
