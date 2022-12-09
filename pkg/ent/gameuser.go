// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
)

// GameUser is the model entity for the GameUser schema.
type GameUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id"`
	// GameID holds the value of the "game_id" field.
	GameID int64 `json:"game_id"`
	// CardID holds the value of the "card_id" field.
	CardID int64 `json:"card_id"`
	// Number holds the value of the "number" field.
	Number uint8 `json:"number"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameUserQuery when eager-loading is set.
	Edges GameUserEdges `json:"edges"`
}

// GameUserEdges holds the relations/edges for other nodes in the graph.
type GameUserEdges struct {
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// Card holds the value of the card edge.
	Card *Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameUserEdges) GameOrErr() (*Game, error) {
	if e.loadedTypes[0] {
		if e.Game == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: game.Label}
		}
		return e.Game, nil
	}
	return nil, &NotLoadedError{edge: "game"}
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameUserEdges) CardOrErr() (*Card, error) {
	if e.loadedTypes[1] {
		if e.Card == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: card.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GameUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gameuser.FieldID, gameuser.FieldCreatedBy, gameuser.FieldUpdatedBy, gameuser.FieldUserID, gameuser.FieldGameID, gameuser.FieldCardID, gameuser.FieldNumber:
			values[i] = new(sql.NullInt64)
		case gameuser.FieldCreatedAt, gameuser.FieldUpdatedAt, gameuser.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GameUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GameUser fields.
func (gu *GameUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gameuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gu.ID = int64(value.Int64)
		case gameuser.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gu.CreatedBy = value.Int64
			}
		case gameuser.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gu.UpdatedBy = value.Int64
			}
		case gameuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gu.CreatedAt = value.Time
			}
		case gameuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gu.UpdatedAt = value.Time
			}
		case gameuser.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gu.DeletedAt = value.Time
			}
		case gameuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gu.UserID = value.Int64
			}
		case gameuser.FieldGameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				gu.GameID = value.Int64
			}
		case gameuser.FieldCardID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field card_id", values[i])
			} else if value.Valid {
				gu.CardID = value.Int64
			}
		case gameuser.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				gu.Number = uint8(value.Int64)
			}
		}
	}
	return nil
}

// QueryGame queries the "game" edge of the GameUser entity.
func (gu *GameUser) QueryGame() *GameQuery {
	return (&GameUserClient{config: gu.config}).QueryGame(gu)
}

// QueryCard queries the "card" edge of the GameUser entity.
func (gu *GameUser) QueryCard() *CardQuery {
	return (&GameUserClient{config: gu.config}).QueryCard(gu)
}

// Update returns a builder for updating this GameUser.
// Note that you need to call GameUser.Unwrap() before calling this method if this GameUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *GameUser) Update() *GameUserUpdateOne {
	return (&GameUserClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the GameUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gu *GameUser) Unwrap() *GameUser {
	_tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: GameUser is not a transactional entity")
	}
	gu.config.driver = _tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *GameUser) String() string {
	var builder strings.Builder
	builder.WriteString("GameUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gu.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", gu.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", gu.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(gu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gu.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", gu.UserID))
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(fmt.Sprintf("%v", gu.GameID))
	builder.WriteString(", ")
	builder.WriteString("card_id=")
	builder.WriteString(fmt.Sprintf("%v", gu.CardID))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", gu.Number))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (gu GameUser) IsEntity() {}

// GameUsers is a parsable slice of GameUser.
type GameUsers []*GameUser

func (gu GameUsers) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
