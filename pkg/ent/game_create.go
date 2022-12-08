// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/avalon_backend/pkg/ent/game"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
)

// GameCreate is the builder for creating a Game entity.
type GameCreate struct {
	config
	mutation *GameMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (gc *GameCreate) SetCreatedBy(i int64) *GameCreate {
	gc.mutation.SetCreatedBy(i)
	return gc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gc *GameCreate) SetNillableCreatedBy(i *int64) *GameCreate {
	if i != nil {
		gc.SetCreatedBy(*i)
	}
	return gc
}

// SetUpdatedBy sets the "updated_by" field.
func (gc *GameCreate) SetUpdatedBy(i int64) *GameCreate {
	gc.mutation.SetUpdatedBy(i)
	return gc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gc *GameCreate) SetNillableUpdatedBy(i *int64) *GameCreate {
	if i != nil {
		gc.SetUpdatedBy(*i)
	}
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GameCreate) SetCreatedAt(t time.Time) *GameCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GameCreate) SetNillableCreatedAt(t *time.Time) *GameCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GameCreate) SetUpdatedAt(t time.Time) *GameCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GameCreate) SetNillableUpdatedAt(t *time.Time) *GameCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GameCreate) SetDeletedAt(t time.Time) *GameCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GameCreate) SetNillableDeletedAt(t *time.Time) *GameCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetRoomID sets the "room_id" field.
func (gc *GameCreate) SetRoomID(i int64) *GameCreate {
	gc.mutation.SetRoomID(i)
	return gc
}

// SetEndBy sets the "end_by" field.
func (gc *GameCreate) SetEndBy(gb game.EndBy) *GameCreate {
	gc.mutation.SetEndBy(gb)
	return gc
}

// SetNillableEndBy sets the "end_by" field if the given value is not nil.
func (gc *GameCreate) SetNillableEndBy(gb *game.EndBy) *GameCreate {
	if gb != nil {
		gc.SetEndBy(*gb)
	}
	return gc
}

// SetCapacity sets the "capacity" field.
func (gc *GameCreate) SetCapacity(u uint8) *GameCreate {
	gc.mutation.SetCapacity(u)
	return gc
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (gc *GameCreate) SetNillableCapacity(u *uint8) *GameCreate {
	if u != nil {
		gc.SetCapacity(*u)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GameCreate) SetID(i int64) *GameCreate {
	gc.mutation.SetID(i)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GameCreate) SetNillableID(i *int64) *GameCreate {
	if i != nil {
		gc.SetID(*i)
	}
	return gc
}

// AddGameUserIDs adds the "game_users" edge to the GameUser entity by IDs.
func (gc *GameCreate) AddGameUserIDs(ids ...int64) *GameCreate {
	gc.mutation.AddGameUserIDs(ids...)
	return gc
}

// AddGameUsers adds the "game_users" edges to the GameUser entity.
func (gc *GameCreate) AddGameUsers(g ...*GameUser) *GameCreate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddGameUserIDs(ids...)
}

// AddMissionIDs adds the "missions" edge to the Mission entity by IDs.
func (gc *GameCreate) AddMissionIDs(ids ...int64) *GameCreate {
	gc.mutation.AddMissionIDs(ids...)
	return gc
}

// AddMissions adds the "missions" edges to the Mission entity.
func (gc *GameCreate) AddMissions(m ...*Mission) *GameCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gc.AddMissionIDs(ids...)
}

// SetRoom sets the "room" edge to the Room entity.
func (gc *GameCreate) SetRoom(r *Room) *GameCreate {
	return gc.SetRoomID(r.ID)
}

// Mutation returns the GameMutation object of the builder.
func (gc *GameCreate) Mutation() *GameMutation {
	return gc.mutation
}

// Save creates the Game in the database.
func (gc *GameCreate) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (gc *GameCreate) SaveX(ctx context.Context) *Game {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GameCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GameCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GameCreate) defaults() {
	if _, ok := gc.mutation.CreatedBy(); !ok {
		v := game.DefaultCreatedBy
		gc.mutation.SetCreatedBy(v)
	}
	if _, ok := gc.mutation.UpdatedBy(); !ok {
		v := game.DefaultUpdatedBy
		gc.mutation.SetUpdatedBy(v)
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := game.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := game.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.DeletedAt(); !ok {
		v := game.DefaultDeletedAt
		gc.mutation.SetDeletedAt(v)
	}
	if _, ok := gc.mutation.EndBy(); !ok {
		v := game.DefaultEndBy
		gc.mutation.SetEndBy(v)
	}
	if _, ok := gc.mutation.Capacity(); !ok {
		v := game.DefaultCapacity
		gc.mutation.SetCapacity(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := game.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GameCreate) check() error {
	if _, ok := gc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Game.created_by"`)}
	}
	if _, ok := gc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Game.updated_by"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Game.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Game.updated_at"`)}
	}
	if _, ok := gc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Game.deleted_at"`)}
	}
	if _, ok := gc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room_id", err: errors.New(`ent: missing required field "Game.room_id"`)}
	}
	if _, ok := gc.mutation.EndBy(); !ok {
		return &ValidationError{Name: "end_by", err: errors.New(`ent: missing required field "Game.end_by"`)}
	}
	if v, ok := gc.mutation.EndBy(); ok {
		if err := game.EndByValidator(v); err != nil {
			return &ValidationError{Name: "end_by", err: fmt.Errorf(`ent: validator failed for field "Game.end_by": %w`, err)}
		}
	}
	if _, ok := gc.mutation.Capacity(); !ok {
		return &ValidationError{Name: "capacity", err: errors.New(`ent: missing required field "Game.capacity"`)}
	}
	if _, ok := gc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room", err: errors.New(`ent: missing required edge "Game.room"`)}
	}
	return nil
}

func (gc *GameCreate) sqlSave(ctx context.Context) (*Game, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (gc *GameCreate) createSpec() (*Game, *sqlgraph.CreateSpec) {
	var (
		_node = &Game{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: game.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: game.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedBy(); ok {
		_spec.SetField(game.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := gc.mutation.UpdatedBy(); ok {
		_spec.SetField(game.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(game.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(game.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.SetField(game.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := gc.mutation.EndBy(); ok {
		_spec.SetField(game.FieldEndBy, field.TypeEnum, value)
		_node.EndBy = value
	}
	if value, ok := gc.mutation.Capacity(); ok {
		_spec.SetField(game.FieldCapacity, field.TypeUint8, value)
		_node.Capacity = value
	}
	if nodes := gc.mutation.GameUsersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.MissionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.RoomIDs(); len(nodes) > 0 {
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
		_node.RoomID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GameCreateBulk is the builder for creating many Game entities in bulk.
type GameCreateBulk struct {
	config
	builders []*GameCreate
}

// Save creates the Game entities in the database.
func (gcb *GameCreateBulk) Save(ctx context.Context) ([]*Game, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Game, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GameMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GameCreateBulk) SaveX(ctx context.Context) []*Game {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GameCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GameCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
