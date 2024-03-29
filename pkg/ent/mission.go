// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
)

// Mission is the model entity for the Mission schema.
type Mission struct {
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
	// 排序
	Sequence uint8 `json:"sequence"`
	// 任务状态
	Status mission.Status `json:"status"`
	// 是否失败
	Failed bool `json:"failed"`
	// 所属游戏 ID
	GameID int64 `json:"game_id"`
	// 需出征人数
	Capacity uint8 `json:"capacity"`
	// 队长用户 ID
	LeaderID int64 `json:"leader_id"`
	// 是否保护轮
	Protected bool `json:"protected"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionQuery when eager-loading is set.
	Edges MissionEdges `json:"edges"`
}

// MissionEdges holds the relations/edges for other nodes in the graph.
type MissionEdges struct {
	// Game holds the value of the game edge.
	Game *Game `json:"game,omitempty"`
	// Squads holds the value of the squads edge.
	Squads []*Squad `json:"squads,omitempty"`
	// Votes holds the value of the votes edge.
	Votes []*Vote `json:"votes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedSquads map[string][]*Squad
	namedVotes  map[string][]*Vote
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) GameOrErr() (*Game, error) {
	if e.loadedTypes[0] {
		if e.Game == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: game.Label}
		}
		return e.Game, nil
	}
	return nil, &NotLoadedError{edge: "game"}
}

// SquadsOrErr returns the Squads value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) SquadsOrErr() ([]*Squad, error) {
	if e.loadedTypes[1] {
		return e.Squads, nil
	}
	return nil, &NotLoadedError{edge: "squads"}
}

// VotesOrErr returns the Votes value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) VotesOrErr() ([]*Vote, error) {
	if e.loadedTypes[2] {
		return e.Votes, nil
	}
	return nil, &NotLoadedError{edge: "votes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mission.FieldFailed, mission.FieldProtected:
			values[i] = new(sql.NullBool)
		case mission.FieldID, mission.FieldCreatedBy, mission.FieldUpdatedBy, mission.FieldSequence, mission.FieldGameID, mission.FieldCapacity, mission.FieldLeaderID:
			values[i] = new(sql.NullInt64)
		case mission.FieldStatus:
			values[i] = new(sql.NullString)
		case mission.FieldCreatedAt, mission.FieldUpdatedAt, mission.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Mission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mission fields.
func (m *Mission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case mission.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				m.CreatedBy = value.Int64
			}
		case mission.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				m.UpdatedBy = value.Int64
			}
		case mission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case mission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case mission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = value.Time
			}
		case mission.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				m.Sequence = uint8(value.Int64)
			}
		case mission.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = mission.Status(value.String)
			}
		case mission.FieldFailed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field failed", values[i])
			} else if value.Valid {
				m.Failed = value.Bool
			}
		case mission.FieldGameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				m.GameID = value.Int64
			}
		case mission.FieldCapacity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field capacity", values[i])
			} else if value.Valid {
				m.Capacity = uint8(value.Int64)
			}
		case mission.FieldLeaderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field leader_id", values[i])
			} else if value.Valid {
				m.LeaderID = value.Int64
			}
		case mission.FieldProtected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field protected", values[i])
			} else if value.Valid {
				m.Protected = value.Bool
			}
		}
	}
	return nil
}

// QueryGame queries the "game" edge of the Mission entity.
func (m *Mission) QueryGame() *GameQuery {
	return (&MissionClient{config: m.config}).QueryGame(m)
}

// QuerySquads queries the "squads" edge of the Mission entity.
func (m *Mission) QuerySquads() *SquadQuery {
	return (&MissionClient{config: m.config}).QuerySquads(m)
}

// QueryVotes queries the "votes" edge of the Mission entity.
func (m *Mission) QueryVotes() *VoteQuery {
	return (&MissionClient{config: m.config}).QueryVotes(m)
}

// Update returns a builder for updating this Mission.
// Note that you need to call Mission.Unwrap() before calling this method if this Mission
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mission) Update() *MissionUpdateOne {
	return (&MissionClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Mission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mission) Unwrap() *Mission {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mission is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mission) String() string {
	var builder strings.Builder
	builder.WriteString("Mission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", m.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(m.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", m.Sequence))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("failed=")
	builder.WriteString(fmt.Sprintf("%v", m.Failed))
	builder.WriteString(", ")
	builder.WriteString("game_id=")
	builder.WriteString(fmt.Sprintf("%v", m.GameID))
	builder.WriteString(", ")
	builder.WriteString("capacity=")
	builder.WriteString(fmt.Sprintf("%v", m.Capacity))
	builder.WriteString(", ")
	builder.WriteString("leader_id=")
	builder.WriteString(fmt.Sprintf("%v", m.LeaderID))
	builder.WriteString(", ")
	builder.WriteString("protected=")
	builder.WriteString(fmt.Sprintf("%v", m.Protected))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (m Mission) IsEntity() {}

// NamedSquads returns the Squads named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Mission) NamedSquads(name string) ([]*Squad, error) {
	if m.Edges.namedSquads == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedSquads[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Mission) appendNamedSquads(name string, edges ...*Squad) {
	if m.Edges.namedSquads == nil {
		m.Edges.namedSquads = make(map[string][]*Squad)
	}
	if len(edges) == 0 {
		m.Edges.namedSquads[name] = []*Squad{}
	} else {
		m.Edges.namedSquads[name] = append(m.Edges.namedSquads[name], edges...)
	}
}

// NamedVotes returns the Votes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Mission) NamedVotes(name string) ([]*Vote, error) {
	if m.Edges.namedVotes == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedVotes[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Mission) appendNamedVotes(name string, edges ...*Vote) {
	if m.Edges.namedVotes == nil {
		m.Edges.namedVotes = make(map[string][]*Vote)
	}
	if len(edges) == 0 {
		m.Edges.namedVotes[name] = []*Vote{}
	} else {
		m.Edges.namedVotes[name] = append(m.Edges.namedVotes[name], edges...)
	}
}

// Missions is a parsable slice of Mission.
type Missions []*Mission

func (m Missions) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
