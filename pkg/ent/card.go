// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
)

// Card is the model entity for the Card schema.
type Card struct {
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
	// Name holds the value of the "name" field.
	Name card.Name `json:"name"`
	// Role holds the value of the "role" field.
	Role card.Role `json:"role"`
	// Tale holds the value of the "tale" field.
	Tale string `json:"tale"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges CardEdges `json:"edges"`
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// GameUsers holds the value of the game_users edge.
	GameUsers []*GameUser `json:"game_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedGameUsers map[string][]*GameUser
}

// GameUsersOrErr returns the GameUsers value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) GameUsersOrErr() ([]*GameUser, error) {
	if e.loadedTypes[0] {
		return e.GameUsers, nil
	}
	return nil, &NotLoadedError{edge: "game_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldID, card.FieldCreatedBy, card.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case card.FieldName, card.FieldRole, card.FieldTale:
			values[i] = new(sql.NullString)
		case card.FieldCreatedAt, card.FieldUpdatedAt, card.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Card", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case card.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.Int64
			}
		case card.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.Int64
			}
		case card.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case card.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case card.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case card.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = card.Name(value.String)
			}
		case card.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				c.Role = card.Role(value.String)
			}
		case card.FieldTale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tale", values[i])
			} else if value.Valid {
				c.Tale = value.String
			}
		}
	}
	return nil
}

// QueryGameUsers queries the "game_users" edge of the Card entity.
func (c *Card) QueryGameUsers() *GameUserQuery {
	return (&CardClient{config: c.config}).QueryGameUsers(c)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return (&CardClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(fmt.Sprintf("%v", c.Name))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", c.Role))
	builder.WriteString(", ")
	builder.WriteString("tale=")
	builder.WriteString(c.Tale)
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (c Card) IsEntity() {}

// NamedGameUsers returns the GameUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Card) NamedGameUsers(name string) ([]*GameUser, error) {
	if c.Edges.namedGameUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedGameUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Card) appendNamedGameUsers(name string, edges ...*GameUser) {
	if c.Edges.namedGameUsers == nil {
		c.Edges.namedGameUsers = make(map[string][]*GameUser)
	}
	if len(edges) == 0 {
		c.Edges.namedGameUsers[name] = []*GameUser{}
	} else {
		c.Edges.namedGameUsers[name] = append(c.Edges.namedGameUsers[name], edges...)
	}
}

// Cards is a parsable slice of Card.
type Cards []*Card

func (c Cards) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
