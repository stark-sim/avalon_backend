// Code generated by ent, DO NOT EDIT.

package card

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldTale holds the string denoting the tale field in the database.
	FieldTale = "tale"
	// FieldRed holds the string denoting the red field in the database.
	FieldRed = "red"
	// EdgeGameUsers holds the string denoting the game_users edge name in mutations.
	EdgeGameUsers = "game_users"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// GameUsersTable is the table that holds the game_users relation/edge.
	GameUsersTable = "game_users"
	// GameUsersInverseTable is the table name for the GameUser entity.
	// It exists in this package in order to avoid circular dependency with the "gameuser" package.
	GameUsersInverseTable = "game_users"
	// GameUsersColumn is the table column denoting the game_users relation/edge.
	GameUsersColumn = "card_id"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldRole,
	FieldTale,
	FieldRed,
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
	// DefaultTale holds the default value on creation for the "tale" field.
	DefaultTale string
	// DefaultRed holds the default value on creation for the "red" field.
	DefaultRed bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// Name defines the type for the "name" enum field.
type Name string

// NameMerlin is the default value of the Name enum.
const DefaultName = NameMerlin

// Name values.
const (
	NameMerlin   Name = "Merlin"
	NamePercival Name = "Percival"
	NameGalahad  Name = "Galahad"
	NameBors     Name = "Bors"
	NameBedivere Name = "Bedivere"
	NameGawain   Name = "Gawain"
	NameKay      Name = "Kay"
	NameEctor    Name = "Ector"
	NameMordred  Name = "Mordred"
	NameMorgana  Name = "Morgana"
	NameOberon   Name = "Oberon"
	NameAgravain Name = "Agravain"
	NameLancelot Name = "Lancelot"
	NameKevin    Name = "Kevin"
	NameStuart   Name = "Stuart"
	NameBob      Name = "Bob"
)

func (n Name) String() string {
	return string(n)
}

// NameValidator is a validator for the "name" field enum values. It is called by the builders before save.
func NameValidator(n Name) error {
	switch n {
	case NameMerlin, NamePercival, NameGalahad, NameBors, NameBedivere, NameGawain, NameKay, NameEctor, NameMordred, NameMorgana, NameOberon, NameAgravain, NameLancelot, NameKevin, NameStuart, NameBob:
		return nil
	default:
		return fmt.Errorf("card: invalid enum value for name field: %q", n)
	}
}

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleProphet     Role = "Prophet"
	RoleKnight      Role = "Knight"
	RoleLoyal       Role = "Loyal"
	RoleUsurper     Role = "Usurper"
	RoleEnchantress Role = "Enchantress"
	RoleAssassin    Role = "Assassin"
	RoleErlking     Role = "Erlking"
	RoleAce         Role = "Ace"
	RoleSinner      Role = "Sinner"
	RoleMinion      Role = "Minion"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleProphet, RoleKnight, RoleLoyal, RoleUsurper, RoleEnchantress, RoleAssassin, RoleErlking, RoleAce, RoleSinner, RoleMinion:
		return nil
	default:
		return fmt.Errorf("card: invalid enum value for role field: %q", r)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Name) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Name) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Name(str)
	if err := NameValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Name", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Role(str)
	if err := RoleValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}
