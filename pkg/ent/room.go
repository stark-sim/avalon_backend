// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
)

// Room is the model entity for the Room schema.
type Room struct {
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
	Name string `json:"name"`
	// Closed holds the value of the "closed" field.
	Closed bool `json:"closed"`
	// GameOn holds the value of the "game_on" field.
	GameOn bool `json:"game_on"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges RoomEdges `json:"edges"`
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// RoomUsers holds the value of the room_users edge.
	RoomUsers []*RoomUser `json:"room_users,omitempty"`
	// Games holds the value of the games edge.
	Games []*Game `json:"games,omitempty"`
	// Records holds the value of the records edge.
	Records []*Record `json:"records,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedRoomUsers map[string][]*RoomUser
	namedGames     map[string][]*Game
	namedRecords   map[string][]*Record
}

// RoomUsersOrErr returns the RoomUsers value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) RoomUsersOrErr() ([]*RoomUser, error) {
	if e.loadedTypes[0] {
		return e.RoomUsers, nil
	}
	return nil, &NotLoadedError{edge: "room_users"}
}

// GamesOrErr returns the Games value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) GamesOrErr() ([]*Game, error) {
	if e.loadedTypes[1] {
		return e.Games, nil
	}
	return nil, &NotLoadedError{edge: "games"}
}

// RecordsOrErr returns the Records value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) RecordsOrErr() ([]*Record, error) {
	if e.loadedTypes[2] {
		return e.Records, nil
	}
	return nil, &NotLoadedError{edge: "records"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldClosed, room.FieldGameOn:
			values[i] = new(sql.NullBool)
		case room.FieldID, room.FieldCreatedBy, room.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case room.FieldName:
			values[i] = new(sql.NullString)
		case room.FieldCreatedAt, room.FieldUpdatedAt, room.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Room", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case room.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				r.CreatedBy = value.Int64
			}
		case room.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				r.UpdatedBy = value.Int64
			}
		case room.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case room.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case room.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = value.Time
			}
		case room.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case room.FieldClosed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field closed", values[i])
			} else if value.Valid {
				r.Closed = value.Bool
			}
		case room.FieldGameOn:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field game_on", values[i])
			} else if value.Valid {
				r.GameOn = value.Bool
			}
		}
	}
	return nil
}

// QueryRoomUsers queries the "room_users" edge of the Room entity.
func (r *Room) QueryRoomUsers() *RoomUserQuery {
	return (&RoomClient{config: r.config}).QueryRoomUsers(r)
}

// QueryGames queries the "games" edge of the Room entity.
func (r *Room) QueryGames() *GameQuery {
	return (&RoomClient{config: r.config}).QueryGames(r)
}

// QueryRecords queries the "records" edge of the Room entity.
func (r *Room) QueryRecords() *RecordQuery {
	return (&RoomClient{config: r.config}).QueryRecords(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(r.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("closed=")
	builder.WriteString(fmt.Sprintf("%v", r.Closed))
	builder.WriteString(", ")
	builder.WriteString("game_on=")
	builder.WriteString(fmt.Sprintf("%v", r.GameOn))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (r Room) IsEntity() {}

// NamedRoomUsers returns the RoomUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Room) NamedRoomUsers(name string) ([]*RoomUser, error) {
	if r.Edges.namedRoomUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedRoomUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Room) appendNamedRoomUsers(name string, edges ...*RoomUser) {
	if r.Edges.namedRoomUsers == nil {
		r.Edges.namedRoomUsers = make(map[string][]*RoomUser)
	}
	if len(edges) == 0 {
		r.Edges.namedRoomUsers[name] = []*RoomUser{}
	} else {
		r.Edges.namedRoomUsers[name] = append(r.Edges.namedRoomUsers[name], edges...)
	}
}

// NamedGames returns the Games named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Room) NamedGames(name string) ([]*Game, error) {
	if r.Edges.namedGames == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedGames[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Room) appendNamedGames(name string, edges ...*Game) {
	if r.Edges.namedGames == nil {
		r.Edges.namedGames = make(map[string][]*Game)
	}
	if len(edges) == 0 {
		r.Edges.namedGames[name] = []*Game{}
	} else {
		r.Edges.namedGames[name] = append(r.Edges.namedGames[name], edges...)
	}
}

// NamedRecords returns the Records named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Room) NamedRecords(name string) ([]*Record, error) {
	if r.Edges.namedRecords == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedRecords[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Room) appendNamedRecords(name string, edges ...*Record) {
	if r.Edges.namedRecords == nil {
		r.Edges.namedRecords = make(map[string][]*Record)
	}
	if len(edges) == 0 {
		r.Edges.namedRecords[name] = []*Record{}
	} else {
		r.Edges.namedRecords[name] = append(r.Edges.namedRecords[name], edges...)
	}
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
