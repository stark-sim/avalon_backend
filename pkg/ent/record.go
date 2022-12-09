// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/avalon_backend/pkg/ent/record"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
)

// Record is the model entity for the Record schema.
type Record struct {
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
	// RoomID holds the value of the "room_id" field.
	RoomID int64 `json:"room_id"`
	// Score holds the value of the "score" field.
	Score int32 `json:"score"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecordQuery when eager-loading is set.
	Edges RecordEdges `json:"edges"`
}

// RecordEdges holds the relations/edges for other nodes in the graph.
type RecordEdges struct {
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[0] {
		if e.Room == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Record) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case record.FieldID, record.FieldCreatedBy, record.FieldUpdatedBy, record.FieldUserID, record.FieldRoomID, record.FieldScore:
			values[i] = new(sql.NullInt64)
		case record.FieldCreatedAt, record.FieldUpdatedAt, record.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Record", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Record fields.
func (r *Record) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case record.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				r.CreatedBy = value.Int64
			}
		case record.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				r.UpdatedBy = value.Int64
			}
		case record.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case record.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case record.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = value.Time
			}
		case record.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = value.Int64
			}
		case record.FieldRoomID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field room_id", values[i])
			} else if value.Valid {
				r.RoomID = value.Int64
			}
		case record.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				r.Score = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryRoom queries the "room" edge of the Record entity.
func (r *Record) QueryRoom() *RoomQuery {
	return (&RecordClient{config: r.config}).QueryRoom(r)
}

// Update returns a builder for updating this Record.
// Note that you need to call Record.Unwrap() before calling this method if this Record
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Record) Update() *RecordUpdateOne {
	return (&RecordClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Record entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Record) Unwrap() *Record {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Record is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Record) String() string {
	var builder strings.Builder
	builder.WriteString("Record(")
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
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("room_id=")
	builder.WriteString(fmt.Sprintf("%v", r.RoomID))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", r.Score))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (r Record) IsEntity() {}

// Records is a parsable slice of Record.
type Records []*Record

func (r Records) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
