// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetCreatedBy sets the "created_by" field.
func (gu *GameUpdate) SetCreatedBy(i int64) *GameUpdate {
	gu.mutation.ResetCreatedBy()
	gu.mutation.SetCreatedBy(i)
	return gu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gu *GameUpdate) SetNillableCreatedBy(i *int64) *GameUpdate {
	if i != nil {
		gu.SetCreatedBy(*i)
	}
	return gu
}

// AddCreatedBy adds i to the "created_by" field.
func (gu *GameUpdate) AddCreatedBy(i int64) *GameUpdate {
	gu.mutation.AddCreatedBy(i)
	return gu
}

// SetUpdatedBy sets the "updated_by" field.
func (gu *GameUpdate) SetUpdatedBy(i int64) *GameUpdate {
	gu.mutation.ResetUpdatedBy()
	gu.mutation.SetUpdatedBy(i)
	return gu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gu *GameUpdate) SetNillableUpdatedBy(i *int64) *GameUpdate {
	if i != nil {
		gu.SetUpdatedBy(*i)
	}
	return gu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (gu *GameUpdate) AddUpdatedBy(i int64) *GameUpdate {
	gu.mutation.AddUpdatedBy(i)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GameUpdate) SetUpdatedAt(t time.Time) *GameUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GameUpdate) SetDeletedAt(t time.Time) *GameUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GameUpdate) SetNillableDeletedAt(t *time.Time) *GameUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// SetRoomID sets the "room_id" field.
func (gu *GameUpdate) SetRoomID(i int64) *GameUpdate {
	gu.mutation.SetRoomID(i)
	return gu
}

// SetResult sets the "result" field.
func (gu *GameUpdate) SetResult(ga game.Result) *GameUpdate {
	gu.mutation.SetResult(ga)
	return gu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (gu *GameUpdate) SetNillableResult(ga *game.Result) *GameUpdate {
	if ga != nil {
		gu.SetResult(*ga)
	}
	return gu
}

// SetCapacity sets the "capacity" field.
func (gu *GameUpdate) SetCapacity(u uint8) *GameUpdate {
	gu.mutation.ResetCapacity()
	gu.mutation.SetCapacity(u)
	return gu
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (gu *GameUpdate) SetNillableCapacity(u *uint8) *GameUpdate {
	if u != nil {
		gu.SetCapacity(*u)
	}
	return gu
}

// AddCapacity adds u to the "capacity" field.
func (gu *GameUpdate) AddCapacity(u int8) *GameUpdate {
	gu.mutation.AddCapacity(u)
	return gu
}

// SetTheAssassinatedIds sets the "the_assassinated_ids" field.
func (gu *GameUpdate) SetTheAssassinatedIds(s []string) *GameUpdate {
	gu.mutation.SetTheAssassinatedIds(s)
	return gu
}

// AppendTheAssassinatedIds appends s to the "the_assassinated_ids" field.
func (gu *GameUpdate) AppendTheAssassinatedIds(s []string) *GameUpdate {
	gu.mutation.AppendTheAssassinatedIds(s)
	return gu
}

// ClearTheAssassinatedIds clears the value of the "the_assassinated_ids" field.
func (gu *GameUpdate) ClearTheAssassinatedIds() *GameUpdate {
	gu.mutation.ClearTheAssassinatedIds()
	return gu
}

// SetAssassinChance sets the "assassin_chance" field.
func (gu *GameUpdate) SetAssassinChance(u uint8) *GameUpdate {
	gu.mutation.ResetAssassinChance()
	gu.mutation.SetAssassinChance(u)
	return gu
}

// SetNillableAssassinChance sets the "assassin_chance" field if the given value is not nil.
func (gu *GameUpdate) SetNillableAssassinChance(u *uint8) *GameUpdate {
	if u != nil {
		gu.SetAssassinChance(*u)
	}
	return gu
}

// AddAssassinChance adds u to the "assassin_chance" field.
func (gu *GameUpdate) AddAssassinChance(u int8) *GameUpdate {
	gu.mutation.AddAssassinChance(u)
	return gu
}

// SetClosed sets the "closed" field.
func (gu *GameUpdate) SetClosed(b bool) *GameUpdate {
	gu.mutation.SetClosed(b)
	return gu
}

// SetNillableClosed sets the "closed" field if the given value is not nil.
func (gu *GameUpdate) SetNillableClosed(b *bool) *GameUpdate {
	if b != nil {
		gu.SetClosed(*b)
	}
	return gu
}

// AddGameUserIDs adds the "game_users" edge to the GameUser entity by IDs.
func (gu *GameUpdate) AddGameUserIDs(ids ...int64) *GameUpdate {
	gu.mutation.AddGameUserIDs(ids...)
	return gu
}

// AddGameUsers adds the "game_users" edges to the GameUser entity.
func (gu *GameUpdate) AddGameUsers(g ...*GameUser) *GameUpdate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddGameUserIDs(ids...)
}

// AddMissionIDs adds the "missions" edge to the Mission entity by IDs.
func (gu *GameUpdate) AddMissionIDs(ids ...int64) *GameUpdate {
	gu.mutation.AddMissionIDs(ids...)
	return gu
}

// AddMissions adds the "missions" edges to the Mission entity.
func (gu *GameUpdate) AddMissions(m ...*Mission) *GameUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMissionIDs(ids...)
}

// SetRoom sets the "room" edge to the Room entity.
func (gu *GameUpdate) SetRoom(r *Room) *GameUpdate {
	return gu.SetRoomID(r.ID)
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// ClearGameUsers clears all "game_users" edges to the GameUser entity.
func (gu *GameUpdate) ClearGameUsers() *GameUpdate {
	gu.mutation.ClearGameUsers()
	return gu
}

// RemoveGameUserIDs removes the "game_users" edge to GameUser entities by IDs.
func (gu *GameUpdate) RemoveGameUserIDs(ids ...int64) *GameUpdate {
	gu.mutation.RemoveGameUserIDs(ids...)
	return gu
}

// RemoveGameUsers removes "game_users" edges to GameUser entities.
func (gu *GameUpdate) RemoveGameUsers(g ...*GameUser) *GameUpdate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveGameUserIDs(ids...)
}

// ClearMissions clears all "missions" edges to the Mission entity.
func (gu *GameUpdate) ClearMissions() *GameUpdate {
	gu.mutation.ClearMissions()
	return gu
}

// RemoveMissionIDs removes the "missions" edge to Mission entities by IDs.
func (gu *GameUpdate) RemoveMissionIDs(ids ...int64) *GameUpdate {
	gu.mutation.RemoveMissionIDs(ids...)
	return gu
}

// RemoveMissions removes "missions" edges to Mission entities.
func (gu *GameUpdate) RemoveMissions(m ...*Mission) *GameUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMissionIDs(ids...)
}

// ClearRoom clears the "room" edge to the Room entity.
func (gu *GameUpdate) ClearRoom() *GameUpdate {
	gu.mutation.ClearRoom()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GameUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := game.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GameUpdate) check() error {
	if v, ok := gu.mutation.Result(); ok {
		if err := game.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "Game.result": %w`, err)}
		}
	}
	if _, ok := gu.mutation.RoomID(); gu.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.room"`)
	}
	return nil
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: game.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.CreatedBy(); ok {
		_spec.SetField(game.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(game.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.UpdatedBy(); ok {
		_spec.SetField(game.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(game.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(game.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.SetField(game.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.Result(); ok {
		_spec.SetField(game.FieldResult, field.TypeEnum, value)
	}
	if value, ok := gu.mutation.Capacity(); ok {
		_spec.SetField(game.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := gu.mutation.AddedCapacity(); ok {
		_spec.AddField(game.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := gu.mutation.TheAssassinatedIds(); ok {
		_spec.SetField(game.FieldTheAssassinatedIds, field.TypeJSON, value)
	}
	if value, ok := gu.mutation.AppendedTheAssassinatedIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, game.FieldTheAssassinatedIds, value)
		})
	}
	if gu.mutation.TheAssassinatedIdsCleared() {
		_spec.ClearField(game.FieldTheAssassinatedIds, field.TypeJSON)
	}
	if value, ok := gu.mutation.AssassinChance(); ok {
		_spec.SetField(game.FieldAssassinChance, field.TypeUint8, value)
	}
	if value, ok := gu.mutation.AddedAssassinChance(); ok {
		_spec.AddField(game.FieldAssassinChance, field.TypeUint8, value)
	}
	if value, ok := gu.mutation.Closed(); ok {
		_spec.SetField(game.FieldClosed, field.TypeBool, value)
	}
	if gu.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.GameUsersTable,
			Columns: []string{game.GameUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: gameuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedGameUsersIDs(); len(nodes) > 0 && !gu.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.GameUsersTable,
			Columns: []string{game.GameUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: gameuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GameUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.GameUsersTable,
			Columns: []string{game.GameUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: gameuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.MissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.MissionsTable,
			Columns: []string{game.MissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: mission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedMissionsIDs(); len(nodes) > 0 && !gu.mutation.MissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.MissionsTable,
			Columns: []string{game.MissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: mission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.MissionsTable,
			Columns: []string{game.MissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: mission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.RoomTable,
			Columns: []string{game.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.RoomTable,
			Columns: []string{game.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// SetCreatedBy sets the "created_by" field.
func (guo *GameUpdateOne) SetCreatedBy(i int64) *GameUpdateOne {
	guo.mutation.ResetCreatedBy()
	guo.mutation.SetCreatedBy(i)
	return guo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableCreatedBy(i *int64) *GameUpdateOne {
	if i != nil {
		guo.SetCreatedBy(*i)
	}
	return guo
}

// AddCreatedBy adds i to the "created_by" field.
func (guo *GameUpdateOne) AddCreatedBy(i int64) *GameUpdateOne {
	guo.mutation.AddCreatedBy(i)
	return guo
}

// SetUpdatedBy sets the "updated_by" field.
func (guo *GameUpdateOne) SetUpdatedBy(i int64) *GameUpdateOne {
	guo.mutation.ResetUpdatedBy()
	guo.mutation.SetUpdatedBy(i)
	return guo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableUpdatedBy(i *int64) *GameUpdateOne {
	if i != nil {
		guo.SetUpdatedBy(*i)
	}
	return guo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (guo *GameUpdateOne) AddUpdatedBy(i int64) *GameUpdateOne {
	guo.mutation.AddUpdatedBy(i)
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GameUpdateOne) SetUpdatedAt(t time.Time) *GameUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GameUpdateOne) SetDeletedAt(t time.Time) *GameUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableDeletedAt(t *time.Time) *GameUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// SetRoomID sets the "room_id" field.
func (guo *GameUpdateOne) SetRoomID(i int64) *GameUpdateOne {
	guo.mutation.SetRoomID(i)
	return guo
}

// SetResult sets the "result" field.
func (guo *GameUpdateOne) SetResult(ga game.Result) *GameUpdateOne {
	guo.mutation.SetResult(ga)
	return guo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableResult(ga *game.Result) *GameUpdateOne {
	if ga != nil {
		guo.SetResult(*ga)
	}
	return guo
}

// SetCapacity sets the "capacity" field.
func (guo *GameUpdateOne) SetCapacity(u uint8) *GameUpdateOne {
	guo.mutation.ResetCapacity()
	guo.mutation.SetCapacity(u)
	return guo
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableCapacity(u *uint8) *GameUpdateOne {
	if u != nil {
		guo.SetCapacity(*u)
	}
	return guo
}

// AddCapacity adds u to the "capacity" field.
func (guo *GameUpdateOne) AddCapacity(u int8) *GameUpdateOne {
	guo.mutation.AddCapacity(u)
	return guo
}

// SetTheAssassinatedIds sets the "the_assassinated_ids" field.
func (guo *GameUpdateOne) SetTheAssassinatedIds(s []string) *GameUpdateOne {
	guo.mutation.SetTheAssassinatedIds(s)
	return guo
}

// AppendTheAssassinatedIds appends s to the "the_assassinated_ids" field.
func (guo *GameUpdateOne) AppendTheAssassinatedIds(s []string) *GameUpdateOne {
	guo.mutation.AppendTheAssassinatedIds(s)
	return guo
}

// ClearTheAssassinatedIds clears the value of the "the_assassinated_ids" field.
func (guo *GameUpdateOne) ClearTheAssassinatedIds() *GameUpdateOne {
	guo.mutation.ClearTheAssassinatedIds()
	return guo
}

// SetAssassinChance sets the "assassin_chance" field.
func (guo *GameUpdateOne) SetAssassinChance(u uint8) *GameUpdateOne {
	guo.mutation.ResetAssassinChance()
	guo.mutation.SetAssassinChance(u)
	return guo
}

// SetNillableAssassinChance sets the "assassin_chance" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableAssassinChance(u *uint8) *GameUpdateOne {
	if u != nil {
		guo.SetAssassinChance(*u)
	}
	return guo
}

// AddAssassinChance adds u to the "assassin_chance" field.
func (guo *GameUpdateOne) AddAssassinChance(u int8) *GameUpdateOne {
	guo.mutation.AddAssassinChance(u)
	return guo
}

// SetClosed sets the "closed" field.
func (guo *GameUpdateOne) SetClosed(b bool) *GameUpdateOne {
	guo.mutation.SetClosed(b)
	return guo
}

// SetNillableClosed sets the "closed" field if the given value is not nil.
func (guo *GameUpdateOne) SetNillableClosed(b *bool) *GameUpdateOne {
	if b != nil {
		guo.SetClosed(*b)
	}
	return guo
}

// AddGameUserIDs adds the "game_users" edge to the GameUser entity by IDs.
func (guo *GameUpdateOne) AddGameUserIDs(ids ...int64) *GameUpdateOne {
	guo.mutation.AddGameUserIDs(ids...)
	return guo
}

// AddGameUsers adds the "game_users" edges to the GameUser entity.
func (guo *GameUpdateOne) AddGameUsers(g ...*GameUser) *GameUpdateOne {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddGameUserIDs(ids...)
}

// AddMissionIDs adds the "missions" edge to the Mission entity by IDs.
func (guo *GameUpdateOne) AddMissionIDs(ids ...int64) *GameUpdateOne {
	guo.mutation.AddMissionIDs(ids...)
	return guo
}

// AddMissions adds the "missions" edges to the Mission entity.
func (guo *GameUpdateOne) AddMissions(m ...*Mission) *GameUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMissionIDs(ids...)
}

// SetRoom sets the "room" edge to the Room entity.
func (guo *GameUpdateOne) SetRoom(r *Room) *GameUpdateOne {
	return guo.SetRoomID(r.ID)
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// ClearGameUsers clears all "game_users" edges to the GameUser entity.
func (guo *GameUpdateOne) ClearGameUsers() *GameUpdateOne {
	guo.mutation.ClearGameUsers()
	return guo
}

// RemoveGameUserIDs removes the "game_users" edge to GameUser entities by IDs.
func (guo *GameUpdateOne) RemoveGameUserIDs(ids ...int64) *GameUpdateOne {
	guo.mutation.RemoveGameUserIDs(ids...)
	return guo
}

// RemoveGameUsers removes "game_users" edges to GameUser entities.
func (guo *GameUpdateOne) RemoveGameUsers(g ...*GameUser) *GameUpdateOne {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveGameUserIDs(ids...)
}

// ClearMissions clears all "missions" edges to the Mission entity.
func (guo *GameUpdateOne) ClearMissions() *GameUpdateOne {
	guo.mutation.ClearMissions()
	return guo
}

// RemoveMissionIDs removes the "missions" edge to Mission entities by IDs.
func (guo *GameUpdateOne) RemoveMissionIDs(ids ...int64) *GameUpdateOne {
	guo.mutation.RemoveMissionIDs(ids...)
	return guo
}

// RemoveMissions removes "missions" edges to Mission entities.
func (guo *GameUpdateOne) RemoveMissions(m ...*Mission) *GameUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMissionIDs(ids...)
}

// ClearRoom clears the "room" edge to the Room entity.
func (guo *GameUpdateOne) ClearRoom() *GameUpdateOne {
	guo.mutation.ClearRoom()
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Game)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GameMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GameUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := game.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GameUpdateOne) check() error {
	if v, ok := guo.mutation.Result(); ok {
		if err := game.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "Game.result": %w`, err)}
		}
	}
	if _, ok := guo.mutation.RoomID(); guo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Game.room"`)
	}
	return nil
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: game.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.CreatedBy(); ok {
		_spec.SetField(game.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(game.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.UpdatedBy(); ok {
		_spec.SetField(game.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(game.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(game.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.SetField(game.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.Result(); ok {
		_spec.SetField(game.FieldResult, field.TypeEnum, value)
	}
	if value, ok := guo.mutation.Capacity(); ok {
		_spec.SetField(game.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := guo.mutation.AddedCapacity(); ok {
		_spec.AddField(game.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := guo.mutation.TheAssassinatedIds(); ok {
		_spec.SetField(game.FieldTheAssassinatedIds, field.TypeJSON, value)
	}
	if value, ok := guo.mutation.AppendedTheAssassinatedIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, game.FieldTheAssassinatedIds, value)
		})
	}
	if guo.mutation.TheAssassinatedIdsCleared() {
		_spec.ClearField(game.FieldTheAssassinatedIds, field.TypeJSON)
	}
	if value, ok := guo.mutation.AssassinChance(); ok {
		_spec.SetField(game.FieldAssassinChance, field.TypeUint8, value)
	}
	if value, ok := guo.mutation.AddedAssassinChance(); ok {
		_spec.AddField(game.FieldAssassinChance, field.TypeUint8, value)
	}
	if value, ok := guo.mutation.Closed(); ok {
		_spec.SetField(game.FieldClosed, field.TypeBool, value)
	}
	if guo.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.GameUsersTable,
			Columns: []string{game.GameUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: gameuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedGameUsersIDs(); len(nodes) > 0 && !guo.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.GameUsersTable,
			Columns: []string{game.GameUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: gameuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GameUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.GameUsersTable,
			Columns: []string{game.GameUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: gameuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.MissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.MissionsTable,
			Columns: []string{game.MissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: mission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedMissionsIDs(); len(nodes) > 0 && !guo.mutation.MissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.MissionsTable,
			Columns: []string{game.MissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: mission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.MissionsTable,
			Columns: []string{game.MissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: mission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.RoomTable,
			Columns: []string{game.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.RoomTable,
			Columns: []string{game.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
