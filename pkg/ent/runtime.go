// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"github.com/stark-sim/avalon_backend/pkg/ent/record"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/schema"
	"github.com/stark-sim/avalon_backend/pkg/ent/squad"
	"github.com/stark-sim/avalon_backend/pkg/ent/vote"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinFields0 := cardMixin[0].Fields()
	_ = cardMixinFields0
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreatedBy is the schema descriptor for created_by field.
	cardDescCreatedBy := cardMixinFields0[1].Descriptor()
	// card.DefaultCreatedBy holds the default value on creation for the created_by field.
	card.DefaultCreatedBy = cardDescCreatedBy.Default.(int64)
	// cardDescUpdatedBy is the schema descriptor for updated_by field.
	cardDescUpdatedBy := cardMixinFields0[2].Descriptor()
	// card.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	card.DefaultUpdatedBy = cardDescUpdatedBy.Default.(int64)
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardMixinFields0[3].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
	// cardDescUpdatedAt is the schema descriptor for updated_at field.
	cardDescUpdatedAt := cardMixinFields0[4].Descriptor()
	// card.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	card.DefaultUpdatedAt = cardDescUpdatedAt.Default.(func() time.Time)
	// card.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	card.UpdateDefaultUpdatedAt = cardDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cardDescDeletedAt is the schema descriptor for deleted_at field.
	cardDescDeletedAt := cardMixinFields0[5].Descriptor()
	// card.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	card.DefaultDeletedAt = cardDescDeletedAt.Default.(time.Time)
	// cardDescTale is the schema descriptor for tale field.
	cardDescTale := cardFields[2].Descriptor()
	// card.DefaultTale holds the default value on creation for the tale field.
	card.DefaultTale = cardDescTale.Default.(string)
	// cardDescID is the schema descriptor for id field.
	cardDescID := cardMixinFields0[0].Descriptor()
	// card.DefaultID holds the default value on creation for the id field.
	card.DefaultID = cardDescID.Default.(func() int64)
	gameMixin := schema.Game{}.Mixin()
	gameMixinFields0 := gameMixin[0].Fields()
	_ = gameMixinFields0
	gameFields := schema.Game{}.Fields()
	_ = gameFields
	// gameDescCreatedBy is the schema descriptor for created_by field.
	gameDescCreatedBy := gameMixinFields0[1].Descriptor()
	// game.DefaultCreatedBy holds the default value on creation for the created_by field.
	game.DefaultCreatedBy = gameDescCreatedBy.Default.(int64)
	// gameDescUpdatedBy is the schema descriptor for updated_by field.
	gameDescUpdatedBy := gameMixinFields0[2].Descriptor()
	// game.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	game.DefaultUpdatedBy = gameDescUpdatedBy.Default.(int64)
	// gameDescCreatedAt is the schema descriptor for created_at field.
	gameDescCreatedAt := gameMixinFields0[3].Descriptor()
	// game.DefaultCreatedAt holds the default value on creation for the created_at field.
	game.DefaultCreatedAt = gameDescCreatedAt.Default.(func() time.Time)
	// gameDescUpdatedAt is the schema descriptor for updated_at field.
	gameDescUpdatedAt := gameMixinFields0[4].Descriptor()
	// game.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	game.DefaultUpdatedAt = gameDescUpdatedAt.Default.(func() time.Time)
	// game.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	game.UpdateDefaultUpdatedAt = gameDescUpdatedAt.UpdateDefault.(func() time.Time)
	// gameDescDeletedAt is the schema descriptor for deleted_at field.
	gameDescDeletedAt := gameMixinFields0[5].Descriptor()
	// game.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	game.DefaultDeletedAt = gameDescDeletedAt.Default.(time.Time)
	// gameDescCapacity is the schema descriptor for capacity field.
	gameDescCapacity := gameFields[2].Descriptor()
	// game.DefaultCapacity holds the default value on creation for the capacity field.
	game.DefaultCapacity = gameDescCapacity.Default.(uint8)
	// gameDescID is the schema descriptor for id field.
	gameDescID := gameMixinFields0[0].Descriptor()
	// game.DefaultID holds the default value on creation for the id field.
	game.DefaultID = gameDescID.Default.(func() int64)
	gameuserMixin := schema.GameUser{}.Mixin()
	gameuserMixinFields0 := gameuserMixin[0].Fields()
	_ = gameuserMixinFields0
	gameuserFields := schema.GameUser{}.Fields()
	_ = gameuserFields
	// gameuserDescCreatedBy is the schema descriptor for created_by field.
	gameuserDescCreatedBy := gameuserMixinFields0[1].Descriptor()
	// gameuser.DefaultCreatedBy holds the default value on creation for the created_by field.
	gameuser.DefaultCreatedBy = gameuserDescCreatedBy.Default.(int64)
	// gameuserDescUpdatedBy is the schema descriptor for updated_by field.
	gameuserDescUpdatedBy := gameuserMixinFields0[2].Descriptor()
	// gameuser.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	gameuser.DefaultUpdatedBy = gameuserDescUpdatedBy.Default.(int64)
	// gameuserDescCreatedAt is the schema descriptor for created_at field.
	gameuserDescCreatedAt := gameuserMixinFields0[3].Descriptor()
	// gameuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	gameuser.DefaultCreatedAt = gameuserDescCreatedAt.Default.(func() time.Time)
	// gameuserDescUpdatedAt is the schema descriptor for updated_at field.
	gameuserDescUpdatedAt := gameuserMixinFields0[4].Descriptor()
	// gameuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	gameuser.DefaultUpdatedAt = gameuserDescUpdatedAt.Default.(func() time.Time)
	// gameuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	gameuser.UpdateDefaultUpdatedAt = gameuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// gameuserDescDeletedAt is the schema descriptor for deleted_at field.
	gameuserDescDeletedAt := gameuserMixinFields0[5].Descriptor()
	// gameuser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	gameuser.DefaultDeletedAt = gameuserDescDeletedAt.Default.(time.Time)
	// gameuserDescID is the schema descriptor for id field.
	gameuserDescID := gameuserMixinFields0[0].Descriptor()
	// gameuser.DefaultID holds the default value on creation for the id field.
	gameuser.DefaultID = gameuserDescID.Default.(func() int64)
	missionMixin := schema.Mission{}.Mixin()
	missionMixinFields0 := missionMixin[0].Fields()
	_ = missionMixinFields0
	missionFields := schema.Mission{}.Fields()
	_ = missionFields
	// missionDescCreatedBy is the schema descriptor for created_by field.
	missionDescCreatedBy := missionMixinFields0[1].Descriptor()
	// mission.DefaultCreatedBy holds the default value on creation for the created_by field.
	mission.DefaultCreatedBy = missionDescCreatedBy.Default.(int64)
	// missionDescUpdatedBy is the schema descriptor for updated_by field.
	missionDescUpdatedBy := missionMixinFields0[2].Descriptor()
	// mission.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	mission.DefaultUpdatedBy = missionDescUpdatedBy.Default.(int64)
	// missionDescCreatedAt is the schema descriptor for created_at field.
	missionDescCreatedAt := missionMixinFields0[3].Descriptor()
	// mission.DefaultCreatedAt holds the default value on creation for the created_at field.
	mission.DefaultCreatedAt = missionDescCreatedAt.Default.(func() time.Time)
	// missionDescUpdatedAt is the schema descriptor for updated_at field.
	missionDescUpdatedAt := missionMixinFields0[4].Descriptor()
	// mission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mission.DefaultUpdatedAt = missionDescUpdatedAt.Default.(func() time.Time)
	// mission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	mission.UpdateDefaultUpdatedAt = missionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// missionDescDeletedAt is the schema descriptor for deleted_at field.
	missionDescDeletedAt := missionMixinFields0[5].Descriptor()
	// mission.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	mission.DefaultDeletedAt = missionDescDeletedAt.Default.(time.Time)
	// missionDescSequence is the schema descriptor for sequence field.
	missionDescSequence := missionFields[0].Descriptor()
	// mission.SequenceValidator is a validator for the "sequence" field. It is called by the builders before save.
	mission.SequenceValidator = func() func(uint8) error {
		validators := missionDescSequence.Validators
		fns := [...]func(uint8) error{
			validators[0].(func(uint8) error),
			validators[1].(func(uint8) error),
		}
		return func(sequence uint8) error {
			for _, fn := range fns {
				if err := fn(sequence); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// missionDescFailed is the schema descriptor for failed field.
	missionDescFailed := missionFields[2].Descriptor()
	// mission.DefaultFailed holds the default value on creation for the failed field.
	mission.DefaultFailed = missionDescFailed.Default.(bool)
	// missionDescCapacity is the schema descriptor for capacity field.
	missionDescCapacity := missionFields[4].Descriptor()
	// mission.DefaultCapacity holds the default value on creation for the capacity field.
	mission.DefaultCapacity = missionDescCapacity.Default.(uint8)
	// missionDescID is the schema descriptor for id field.
	missionDescID := missionMixinFields0[0].Descriptor()
	// mission.DefaultID holds the default value on creation for the id field.
	mission.DefaultID = missionDescID.Default.(func() int64)
	recordMixin := schema.Record{}.Mixin()
	recordMixinFields0 := recordMixin[0].Fields()
	_ = recordMixinFields0
	recordFields := schema.Record{}.Fields()
	_ = recordFields
	// recordDescCreatedBy is the schema descriptor for created_by field.
	recordDescCreatedBy := recordMixinFields0[1].Descriptor()
	// record.DefaultCreatedBy holds the default value on creation for the created_by field.
	record.DefaultCreatedBy = recordDescCreatedBy.Default.(int64)
	// recordDescUpdatedBy is the schema descriptor for updated_by field.
	recordDescUpdatedBy := recordMixinFields0[2].Descriptor()
	// record.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	record.DefaultUpdatedBy = recordDescUpdatedBy.Default.(int64)
	// recordDescCreatedAt is the schema descriptor for created_at field.
	recordDescCreatedAt := recordMixinFields0[3].Descriptor()
	// record.DefaultCreatedAt holds the default value on creation for the created_at field.
	record.DefaultCreatedAt = recordDescCreatedAt.Default.(func() time.Time)
	// recordDescUpdatedAt is the schema descriptor for updated_at field.
	recordDescUpdatedAt := recordMixinFields0[4].Descriptor()
	// record.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	record.DefaultUpdatedAt = recordDescUpdatedAt.Default.(func() time.Time)
	// record.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	record.UpdateDefaultUpdatedAt = recordDescUpdatedAt.UpdateDefault.(func() time.Time)
	// recordDescDeletedAt is the schema descriptor for deleted_at field.
	recordDescDeletedAt := recordMixinFields0[5].Descriptor()
	// record.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	record.DefaultDeletedAt = recordDescDeletedAt.Default.(time.Time)
	// recordDescScore is the schema descriptor for score field.
	recordDescScore := recordFields[2].Descriptor()
	// record.DefaultScore holds the default value on creation for the score field.
	record.DefaultScore = recordDescScore.Default.(int32)
	// recordDescID is the schema descriptor for id field.
	recordDescID := recordMixinFields0[0].Descriptor()
	// record.DefaultID holds the default value on creation for the id field.
	record.DefaultID = recordDescID.Default.(func() int64)
	roomMixin := schema.Room{}.Mixin()
	roomMixinFields0 := roomMixin[0].Fields()
	_ = roomMixinFields0
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescCreatedBy is the schema descriptor for created_by field.
	roomDescCreatedBy := roomMixinFields0[1].Descriptor()
	// room.DefaultCreatedBy holds the default value on creation for the created_by field.
	room.DefaultCreatedBy = roomDescCreatedBy.Default.(int64)
	// roomDescUpdatedBy is the schema descriptor for updated_by field.
	roomDescUpdatedBy := roomMixinFields0[2].Descriptor()
	// room.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	room.DefaultUpdatedBy = roomDescUpdatedBy.Default.(int64)
	// roomDescCreatedAt is the schema descriptor for created_at field.
	roomDescCreatedAt := roomMixinFields0[3].Descriptor()
	// room.DefaultCreatedAt holds the default value on creation for the created_at field.
	room.DefaultCreatedAt = roomDescCreatedAt.Default.(func() time.Time)
	// roomDescUpdatedAt is the schema descriptor for updated_at field.
	roomDescUpdatedAt := roomMixinFields0[4].Descriptor()
	// room.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	room.DefaultUpdatedAt = roomDescUpdatedAt.Default.(func() time.Time)
	// room.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	room.UpdateDefaultUpdatedAt = roomDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roomDescDeletedAt is the schema descriptor for deleted_at field.
	roomDescDeletedAt := roomMixinFields0[5].Descriptor()
	// room.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	room.DefaultDeletedAt = roomDescDeletedAt.Default.(time.Time)
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomFields[0].Descriptor()
	// room.DefaultName holds the default value on creation for the name field.
	room.DefaultName = roomDescName.Default.(string)
	// roomDescClosed is the schema descriptor for closed field.
	roomDescClosed := roomFields[1].Descriptor()
	// room.DefaultClosed holds the default value on creation for the closed field.
	room.DefaultClosed = roomDescClosed.Default.(bool)
	// roomDescGameOn is the schema descriptor for game_on field.
	roomDescGameOn := roomFields[2].Descriptor()
	// room.DefaultGameOn holds the default value on creation for the game_on field.
	room.DefaultGameOn = roomDescGameOn.Default.(bool)
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomMixinFields0[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() int64)
	roomuserMixin := schema.RoomUser{}.Mixin()
	roomuserMixinFields0 := roomuserMixin[0].Fields()
	_ = roomuserMixinFields0
	roomuserFields := schema.RoomUser{}.Fields()
	_ = roomuserFields
	// roomuserDescCreatedBy is the schema descriptor for created_by field.
	roomuserDescCreatedBy := roomuserMixinFields0[1].Descriptor()
	// roomuser.DefaultCreatedBy holds the default value on creation for the created_by field.
	roomuser.DefaultCreatedBy = roomuserDescCreatedBy.Default.(int64)
	// roomuserDescUpdatedBy is the schema descriptor for updated_by field.
	roomuserDescUpdatedBy := roomuserMixinFields0[2].Descriptor()
	// roomuser.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	roomuser.DefaultUpdatedBy = roomuserDescUpdatedBy.Default.(int64)
	// roomuserDescCreatedAt is the schema descriptor for created_at field.
	roomuserDescCreatedAt := roomuserMixinFields0[3].Descriptor()
	// roomuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	roomuser.DefaultCreatedAt = roomuserDescCreatedAt.Default.(func() time.Time)
	// roomuserDescUpdatedAt is the schema descriptor for updated_at field.
	roomuserDescUpdatedAt := roomuserMixinFields0[4].Descriptor()
	// roomuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	roomuser.DefaultUpdatedAt = roomuserDescUpdatedAt.Default.(func() time.Time)
	// roomuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	roomuser.UpdateDefaultUpdatedAt = roomuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roomuserDescDeletedAt is the schema descriptor for deleted_at field.
	roomuserDescDeletedAt := roomuserMixinFields0[5].Descriptor()
	// roomuser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	roomuser.DefaultDeletedAt = roomuserDescDeletedAt.Default.(time.Time)
	// roomuserDescID is the schema descriptor for id field.
	roomuserDescID := roomuserMixinFields0[0].Descriptor()
	// roomuser.DefaultID holds the default value on creation for the id field.
	roomuser.DefaultID = roomuserDescID.Default.(func() int64)
	squadMixin := schema.Squad{}.Mixin()
	squadMixinFields0 := squadMixin[0].Fields()
	_ = squadMixinFields0
	squadFields := schema.Squad{}.Fields()
	_ = squadFields
	// squadDescCreatedBy is the schema descriptor for created_by field.
	squadDescCreatedBy := squadMixinFields0[1].Descriptor()
	// squad.DefaultCreatedBy holds the default value on creation for the created_by field.
	squad.DefaultCreatedBy = squadDescCreatedBy.Default.(int64)
	// squadDescUpdatedBy is the schema descriptor for updated_by field.
	squadDescUpdatedBy := squadMixinFields0[2].Descriptor()
	// squad.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	squad.DefaultUpdatedBy = squadDescUpdatedBy.Default.(int64)
	// squadDescCreatedAt is the schema descriptor for created_at field.
	squadDescCreatedAt := squadMixinFields0[3].Descriptor()
	// squad.DefaultCreatedAt holds the default value on creation for the created_at field.
	squad.DefaultCreatedAt = squadDescCreatedAt.Default.(func() time.Time)
	// squadDescUpdatedAt is the schema descriptor for updated_at field.
	squadDescUpdatedAt := squadMixinFields0[4].Descriptor()
	// squad.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	squad.DefaultUpdatedAt = squadDescUpdatedAt.Default.(func() time.Time)
	// squad.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	squad.UpdateDefaultUpdatedAt = squadDescUpdatedAt.UpdateDefault.(func() time.Time)
	// squadDescDeletedAt is the schema descriptor for deleted_at field.
	squadDescDeletedAt := squadMixinFields0[5].Descriptor()
	// squad.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	squad.DefaultDeletedAt = squadDescDeletedAt.Default.(time.Time)
	// squadDescRat is the schema descriptor for rat field.
	squadDescRat := squadFields[2].Descriptor()
	// squad.DefaultRat holds the default value on creation for the rat field.
	squad.DefaultRat = squadDescRat.Default.(bool)
	// squadDescID is the schema descriptor for id field.
	squadDescID := squadMixinFields0[0].Descriptor()
	// squad.DefaultID holds the default value on creation for the id field.
	squad.DefaultID = squadDescID.Default.(func() int64)
	voteMixin := schema.Vote{}.Mixin()
	voteMixinFields0 := voteMixin[0].Fields()
	_ = voteMixinFields0
	voteFields := schema.Vote{}.Fields()
	_ = voteFields
	// voteDescCreatedBy is the schema descriptor for created_by field.
	voteDescCreatedBy := voteMixinFields0[1].Descriptor()
	// vote.DefaultCreatedBy holds the default value on creation for the created_by field.
	vote.DefaultCreatedBy = voteDescCreatedBy.Default.(int64)
	// voteDescUpdatedBy is the schema descriptor for updated_by field.
	voteDescUpdatedBy := voteMixinFields0[2].Descriptor()
	// vote.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	vote.DefaultUpdatedBy = voteDescUpdatedBy.Default.(int64)
	// voteDescCreatedAt is the schema descriptor for created_at field.
	voteDescCreatedAt := voteMixinFields0[3].Descriptor()
	// vote.DefaultCreatedAt holds the default value on creation for the created_at field.
	vote.DefaultCreatedAt = voteDescCreatedAt.Default.(func() time.Time)
	// voteDescUpdatedAt is the schema descriptor for updated_at field.
	voteDescUpdatedAt := voteMixinFields0[4].Descriptor()
	// vote.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vote.DefaultUpdatedAt = voteDescUpdatedAt.Default.(func() time.Time)
	// vote.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vote.UpdateDefaultUpdatedAt = voteDescUpdatedAt.UpdateDefault.(func() time.Time)
	// voteDescDeletedAt is the schema descriptor for deleted_at field.
	voteDescDeletedAt := voteMixinFields0[5].Descriptor()
	// vote.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	vote.DefaultDeletedAt = voteDescDeletedAt.Default.(time.Time)
	// voteDescPass is the schema descriptor for pass field.
	voteDescPass := voteFields[2].Descriptor()
	// vote.DefaultPass holds the default value on creation for the pass field.
	vote.DefaultPass = voteDescPass.Default.(bool)
	// voteDescID is the schema descriptor for id field.
	voteDescID := voteMixinFields0[0].Descriptor()
	// vote.DefaultID holds the default value on creation for the id field.
	vote.DefaultID = voteDescID.Default.(func() int64)
}
