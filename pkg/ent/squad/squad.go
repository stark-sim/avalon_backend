// Code generated by ent, DO NOT EDIT.

package squad

import (
	"time"
)

const (
	// Label holds the string label denoting the squad type in the database.
	Label = "squad"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldMissionID holds the string denoting the mission_id field in the database.
	FieldMissionID = "mission_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRat holds the string denoting the rat field in the database.
	FieldRat = "rat"
	// FieldActed holds the string denoting the acted field in the database.
	FieldActed = "acted"
	// EdgeMission holds the string denoting the mission edge name in mutations.
	EdgeMission = "mission"
	// Table holds the table name of the squad in the database.
	Table = "squads"
	// MissionTable is the table that holds the mission relation/edge.
	MissionTable = "squads"
	// MissionInverseTable is the table name for the Mission entity.
	// It exists in this package in order to avoid circular dependency with the "mission" package.
	MissionInverseTable = "missions"
	// MissionColumn is the table column denoting the mission relation/edge.
	MissionColumn = "mission_id"
)

// Columns holds all SQL columns for squad fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMissionID,
	FieldUserID,
	FieldRat,
	FieldActed,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultRat holds the default value on creation for the "rat" field.
	DefaultRat bool
	// DefaultActed holds the default value on creation for the "acted" field.
	DefaultActed bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)
