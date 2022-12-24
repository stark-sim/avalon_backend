// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
	"github.com/stark-sim/avalon_backend/pkg/ent/record"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// Where appends a list predicates to the RoomUpdate builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedBy sets the "created_by" field.
func (ru *RoomUpdate) SetCreatedBy(i int64) *RoomUpdate {
	ru.mutation.ResetCreatedBy()
	ru.mutation.SetCreatedBy(i)
	return ru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableCreatedBy(i *int64) *RoomUpdate {
	if i != nil {
		ru.SetCreatedBy(*i)
	}
	return ru
}

// AddCreatedBy adds i to the "created_by" field.
func (ru *RoomUpdate) AddCreatedBy(i int64) *RoomUpdate {
	ru.mutation.AddCreatedBy(i)
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RoomUpdate) SetUpdatedBy(i int64) *RoomUpdate {
	ru.mutation.ResetUpdatedBy()
	ru.mutation.SetUpdatedBy(i)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableUpdatedBy(i *int64) *RoomUpdate {
	if i != nil {
		ru.SetUpdatedBy(*i)
	}
	return ru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ru *RoomUpdate) AddUpdatedBy(i int64) *RoomUpdate {
	ru.mutation.AddUpdatedBy(i)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoomUpdate) SetUpdatedAt(t time.Time) *RoomUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RoomUpdate) SetDeletedAt(t time.Time) *RoomUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableDeletedAt(t *time.Time) *RoomUpdate {
	if t != nil {
		ru.SetDeletedAt(*t)
	}
	return ru
}

// SetName sets the "name" field.
func (ru *RoomUpdate) SetName(s string) *RoomUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableName(s *string) *RoomUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetClosed sets the "closed" field.
func (ru *RoomUpdate) SetClosed(b bool) *RoomUpdate {
	ru.mutation.SetClosed(b)
	return ru
}

// SetNillableClosed sets the "closed" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableClosed(b *bool) *RoomUpdate {
	if b != nil {
		ru.SetClosed(*b)
	}
	return ru
}

// SetGameOn sets the "game_on" field.
func (ru *RoomUpdate) SetGameOn(b bool) *RoomUpdate {
	ru.mutation.SetGameOn(b)
	return ru
}

// SetNillableGameOn sets the "game_on" field if the given value is not nil.
func (ru *RoomUpdate) SetNillableGameOn(b *bool) *RoomUpdate {
	if b != nil {
		ru.SetGameOn(*b)
	}
	return ru
}

// AddRoomUserIDs adds the "room_users" edge to the RoomUser entity by IDs.
func (ru *RoomUpdate) AddRoomUserIDs(ids ...int64) *RoomUpdate {
	ru.mutation.AddRoomUserIDs(ids...)
	return ru
}

// AddRoomUsers adds the "room_users" edges to the RoomUser entity.
func (ru *RoomUpdate) AddRoomUsers(r ...*RoomUser) *RoomUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRoomUserIDs(ids...)
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (ru *RoomUpdate) AddGameIDs(ids ...int64) *RoomUpdate {
	ru.mutation.AddGameIDs(ids...)
	return ru
}

// AddGames adds the "games" edges to the Game entity.
func (ru *RoomUpdate) AddGames(g ...*Game) *RoomUpdate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ru.AddGameIDs(ids...)
}

// AddRecordIDs adds the "records" edge to the Record entity by IDs.
func (ru *RoomUpdate) AddRecordIDs(ids ...int64) *RoomUpdate {
	ru.mutation.AddRecordIDs(ids...)
	return ru
}

// AddRecords adds the "records" edges to the Record entity.
func (ru *RoomUpdate) AddRecords(r ...*Record) *RoomUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRecordIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// ClearRoomUsers clears all "room_users" edges to the RoomUser entity.
func (ru *RoomUpdate) ClearRoomUsers() *RoomUpdate {
	ru.mutation.ClearRoomUsers()
	return ru
}

// RemoveRoomUserIDs removes the "room_users" edge to RoomUser entities by IDs.
func (ru *RoomUpdate) RemoveRoomUserIDs(ids ...int64) *RoomUpdate {
	ru.mutation.RemoveRoomUserIDs(ids...)
	return ru
}

// RemoveRoomUsers removes "room_users" edges to RoomUser entities.
func (ru *RoomUpdate) RemoveRoomUsers(r ...*RoomUser) *RoomUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRoomUserIDs(ids...)
}

// ClearGames clears all "games" edges to the Game entity.
func (ru *RoomUpdate) ClearGames() *RoomUpdate {
	ru.mutation.ClearGames()
	return ru
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (ru *RoomUpdate) RemoveGameIDs(ids ...int64) *RoomUpdate {
	ru.mutation.RemoveGameIDs(ids...)
	return ru
}

// RemoveGames removes "games" edges to Game entities.
func (ru *RoomUpdate) RemoveGames(g ...*Game) *RoomUpdate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ru.RemoveGameIDs(ids...)
}

// ClearRecords clears all "records" edges to the Record entity.
func (ru *RoomUpdate) ClearRecords() *RoomUpdate {
	ru.mutation.ClearRecords()
	return ru
}

// RemoveRecordIDs removes the "records" edge to Record entities by IDs.
func (ru *RoomUpdate) RemoveRecordIDs(ids ...int64) *RoomUpdate {
	ru.mutation.RemoveRecordIDs(ids...)
	return ru
}

// RemoveRecords removes "records" edges to Record entities.
func (ru *RoomUpdate) RemoveRecords(r ...*Record) *RoomUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRecordIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoomUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := room.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: room.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CreatedBy(); ok {
		_spec.SetField(room.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedCreatedBy(); ok {
		_spec.AddField(room.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.SetField(room.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(room.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(room.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(room.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Closed(); ok {
		_spec.SetField(room.FieldClosed, field.TypeBool, value)
	}
	if value, ok := ru.mutation.GameOn(); ok {
		_spec.SetField(room.FieldGameOn, field.TypeBool, value)
	}
	if ru.mutation.RoomUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: roomuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRoomUsersIDs(); len(nodes) > 0 && !ru.mutation.RoomUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: roomuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoomUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: roomuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.GamesTable,
			Columns: []string{room.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedGamesIDs(); len(nodes) > 0 && !ru.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.GamesTable,
			Columns: []string{room.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.GamesTable,
			Columns: []string{room.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RecordsTable,
			Columns: []string{room.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRecordsIDs(); len(nodes) > 0 && !ru.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RecordsTable,
			Columns: []string{room.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RecordsTable,
			Columns: []string{room.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomMutation
}

// SetCreatedBy sets the "created_by" field.
func (ruo *RoomUpdateOne) SetCreatedBy(i int64) *RoomUpdateOne {
	ruo.mutation.ResetCreatedBy()
	ruo.mutation.SetCreatedBy(i)
	return ruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableCreatedBy(i *int64) *RoomUpdateOne {
	if i != nil {
		ruo.SetCreatedBy(*i)
	}
	return ruo
}

// AddCreatedBy adds i to the "created_by" field.
func (ruo *RoomUpdateOne) AddCreatedBy(i int64) *RoomUpdateOne {
	ruo.mutation.AddCreatedBy(i)
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RoomUpdateOne) SetUpdatedBy(i int64) *RoomUpdateOne {
	ruo.mutation.ResetUpdatedBy()
	ruo.mutation.SetUpdatedBy(i)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableUpdatedBy(i *int64) *RoomUpdateOne {
	if i != nil {
		ruo.SetUpdatedBy(*i)
	}
	return ruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ruo *RoomUpdateOne) AddUpdatedBy(i int64) *RoomUpdateOne {
	ruo.mutation.AddUpdatedBy(i)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoomUpdateOne) SetUpdatedAt(t time.Time) *RoomUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RoomUpdateOne) SetDeletedAt(t time.Time) *RoomUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableDeletedAt(t *time.Time) *RoomUpdateOne {
	if t != nil {
		ruo.SetDeletedAt(*t)
	}
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoomUpdateOne) SetName(s string) *RoomUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableName(s *string) *RoomUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetClosed sets the "closed" field.
func (ruo *RoomUpdateOne) SetClosed(b bool) *RoomUpdateOne {
	ruo.mutation.SetClosed(b)
	return ruo
}

// SetNillableClosed sets the "closed" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableClosed(b *bool) *RoomUpdateOne {
	if b != nil {
		ruo.SetClosed(*b)
	}
	return ruo
}

// SetGameOn sets the "game_on" field.
func (ruo *RoomUpdateOne) SetGameOn(b bool) *RoomUpdateOne {
	ruo.mutation.SetGameOn(b)
	return ruo
}

// SetNillableGameOn sets the "game_on" field if the given value is not nil.
func (ruo *RoomUpdateOne) SetNillableGameOn(b *bool) *RoomUpdateOne {
	if b != nil {
		ruo.SetGameOn(*b)
	}
	return ruo
}

// AddRoomUserIDs adds the "room_users" edge to the RoomUser entity by IDs.
func (ruo *RoomUpdateOne) AddRoomUserIDs(ids ...int64) *RoomUpdateOne {
	ruo.mutation.AddRoomUserIDs(ids...)
	return ruo
}

// AddRoomUsers adds the "room_users" edges to the RoomUser entity.
func (ruo *RoomUpdateOne) AddRoomUsers(r ...*RoomUser) *RoomUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRoomUserIDs(ids...)
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (ruo *RoomUpdateOne) AddGameIDs(ids ...int64) *RoomUpdateOne {
	ruo.mutation.AddGameIDs(ids...)
	return ruo
}

// AddGames adds the "games" edges to the Game entity.
func (ruo *RoomUpdateOne) AddGames(g ...*Game) *RoomUpdateOne {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ruo.AddGameIDs(ids...)
}

// AddRecordIDs adds the "records" edge to the Record entity by IDs.
func (ruo *RoomUpdateOne) AddRecordIDs(ids ...int64) *RoomUpdateOne {
	ruo.mutation.AddRecordIDs(ids...)
	return ruo
}

// AddRecords adds the "records" edges to the Record entity.
func (ruo *RoomUpdateOne) AddRecords(r ...*Record) *RoomUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRecordIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// ClearRoomUsers clears all "room_users" edges to the RoomUser entity.
func (ruo *RoomUpdateOne) ClearRoomUsers() *RoomUpdateOne {
	ruo.mutation.ClearRoomUsers()
	return ruo
}

// RemoveRoomUserIDs removes the "room_users" edge to RoomUser entities by IDs.
func (ruo *RoomUpdateOne) RemoveRoomUserIDs(ids ...int64) *RoomUpdateOne {
	ruo.mutation.RemoveRoomUserIDs(ids...)
	return ruo
}

// RemoveRoomUsers removes "room_users" edges to RoomUser entities.
func (ruo *RoomUpdateOne) RemoveRoomUsers(r ...*RoomUser) *RoomUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRoomUserIDs(ids...)
}

// ClearGames clears all "games" edges to the Game entity.
func (ruo *RoomUpdateOne) ClearGames() *RoomUpdateOne {
	ruo.mutation.ClearGames()
	return ruo
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (ruo *RoomUpdateOne) RemoveGameIDs(ids ...int64) *RoomUpdateOne {
	ruo.mutation.RemoveGameIDs(ids...)
	return ruo
}

// RemoveGames removes "games" edges to Game entities.
func (ruo *RoomUpdateOne) RemoveGames(g ...*Game) *RoomUpdateOne {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ruo.RemoveGameIDs(ids...)
}

// ClearRecords clears all "records" edges to the Record entity.
func (ruo *RoomUpdateOne) ClearRecords() *RoomUpdateOne {
	ruo.mutation.ClearRecords()
	return ruo
}

// RemoveRecordIDs removes the "records" edge to Record entities by IDs.
func (ruo *RoomUpdateOne) RemoveRecordIDs(ids ...int64) *RoomUpdateOne {
	ruo.mutation.RemoveRecordIDs(ids...)
	return ruo
}

// RemoveRecords removes "records" edges to Record entities.
func (ruo *RoomUpdateOne) RemoveRecords(r ...*Record) *RoomUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRecordIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoomUpdateOne) Select(field string, fields ...string) *RoomUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Room entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	var (
		err  error
		node *Room
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Room)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RoomMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoomUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := room.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (_node *Room, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: room.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Room.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for _, f := range fields {
			if !room.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != room.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CreatedBy(); ok {
		_spec.SetField(room.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(room.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.SetField(room.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(room.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(room.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(room.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(room.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Closed(); ok {
		_spec.SetField(room.FieldClosed, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.GameOn(); ok {
		_spec.SetField(room.FieldGameOn, field.TypeBool, value)
	}
	if ruo.mutation.RoomUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: roomuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRoomUsersIDs(); len(nodes) > 0 && !ruo.mutation.RoomUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: roomuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoomUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: roomuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.GamesTable,
			Columns: []string{room.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedGamesIDs(); len(nodes) > 0 && !ruo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.GamesTable,
			Columns: []string{room.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.GamesTable,
			Columns: []string{room.GamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RecordsTable,
			Columns: []string{room.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRecordsIDs(); len(nodes) > 0 && !ruo.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RecordsTable,
			Columns: []string{room.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RecordsTable,
			Columns: []string{room.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: record.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Room{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
