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
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
)

// GameUserUpdate is the builder for updating GameUser entities.
type GameUserUpdate struct {
	config
	hooks    []Hook
	mutation *GameUserMutation
}

// Where appends a list predicates to the GameUserUpdate builder.
func (guu *GameUserUpdate) Where(ps ...predicate.GameUser) *GameUserUpdate {
	guu.mutation.Where(ps...)
	return guu
}

// SetCreatedBy sets the "created_by" field.
func (guu *GameUserUpdate) SetCreatedBy(i int64) *GameUserUpdate {
	guu.mutation.ResetCreatedBy()
	guu.mutation.SetCreatedBy(i)
	return guu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (guu *GameUserUpdate) SetNillableCreatedBy(i *int64) *GameUserUpdate {
	if i != nil {
		guu.SetCreatedBy(*i)
	}
	return guu
}

// AddCreatedBy adds i to the "created_by" field.
func (guu *GameUserUpdate) AddCreatedBy(i int64) *GameUserUpdate {
	guu.mutation.AddCreatedBy(i)
	return guu
}

// SetUpdatedBy sets the "updated_by" field.
func (guu *GameUserUpdate) SetUpdatedBy(i int64) *GameUserUpdate {
	guu.mutation.ResetUpdatedBy()
	guu.mutation.SetUpdatedBy(i)
	return guu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (guu *GameUserUpdate) SetNillableUpdatedBy(i *int64) *GameUserUpdate {
	if i != nil {
		guu.SetUpdatedBy(*i)
	}
	return guu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (guu *GameUserUpdate) AddUpdatedBy(i int64) *GameUserUpdate {
	guu.mutation.AddUpdatedBy(i)
	return guu
}

// SetUpdatedAt sets the "updated_at" field.
func (guu *GameUserUpdate) SetUpdatedAt(t time.Time) *GameUserUpdate {
	guu.mutation.SetUpdatedAt(t)
	return guu
}

// SetDeletedAt sets the "deleted_at" field.
func (guu *GameUserUpdate) SetDeletedAt(t time.Time) *GameUserUpdate {
	guu.mutation.SetDeletedAt(t)
	return guu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guu *GameUserUpdate) SetNillableDeletedAt(t *time.Time) *GameUserUpdate {
	if t != nil {
		guu.SetDeletedAt(*t)
	}
	return guu
}

// SetUserID sets the "user_id" field.
func (guu *GameUserUpdate) SetUserID(i int64) *GameUserUpdate {
	guu.mutation.ResetUserID()
	guu.mutation.SetUserID(i)
	return guu
}

// AddUserID adds i to the "user_id" field.
func (guu *GameUserUpdate) AddUserID(i int64) *GameUserUpdate {
	guu.mutation.AddUserID(i)
	return guu
}

// SetGameID sets the "game_id" field.
func (guu *GameUserUpdate) SetGameID(i int64) *GameUserUpdate {
	guu.mutation.SetGameID(i)
	return guu
}

// SetGame sets the "game" edge to the Game entity.
func (guu *GameUserUpdate) SetGame(g *Game) *GameUserUpdate {
	return guu.SetGameID(g.ID)
}

// Mutation returns the GameUserMutation object of the builder.
func (guu *GameUserUpdate) Mutation() *GameUserMutation {
	return guu.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (guu *GameUserUpdate) ClearGame() *GameUserUpdate {
	guu.mutation.ClearGame()
	return guu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (guu *GameUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	guu.defaults()
	if len(guu.hooks) == 0 {
		if err = guu.check(); err != nil {
			return 0, err
		}
		affected, err = guu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guu.check(); err != nil {
				return 0, err
			}
			guu.mutation = mutation
			affected, err = guu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(guu.hooks) - 1; i >= 0; i-- {
			if guu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (guu *GameUserUpdate) SaveX(ctx context.Context) int {
	affected, err := guu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (guu *GameUserUpdate) Exec(ctx context.Context) error {
	_, err := guu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guu *GameUserUpdate) ExecX(ctx context.Context) {
	if err := guu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guu *GameUserUpdate) defaults() {
	if _, ok := guu.mutation.UpdatedAt(); !ok {
		v := gameuser.UpdateDefaultUpdatedAt()
		guu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guu *GameUserUpdate) check() error {
	if _, ok := guu.mutation.GameID(); guu.mutation.GameCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GameUser.game"`)
	}
	return nil
}

func (guu *GameUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gameuser.Table,
			Columns: gameuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gameuser.FieldID,
			},
		},
	}
	if ps := guu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guu.mutation.CreatedBy(); ok {
		_spec.SetField(gameuser.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(gameuser.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guu.mutation.UpdatedBy(); ok {
		_spec.SetField(gameuser.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(gameuser.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guu.mutation.UpdatedAt(); ok {
		_spec.SetField(gameuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guu.mutation.DeletedAt(); ok {
		_spec.SetField(gameuser.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := guu.mutation.UserID(); ok {
		_spec.SetField(gameuser.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := guu.mutation.AddedUserID(); ok {
		_spec.AddField(gameuser.FieldUserID, field.TypeInt64, value)
	}
	if guu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameuser.GameTable,
			Columns: []string{gameuser.GameColumn},
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
	if nodes := guu.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameuser.GameTable,
			Columns: []string{gameuser.GameColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, guu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GameUserUpdateOne is the builder for updating a single GameUser entity.
type GameUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameUserMutation
}

// SetCreatedBy sets the "created_by" field.
func (guuo *GameUserUpdateOne) SetCreatedBy(i int64) *GameUserUpdateOne {
	guuo.mutation.ResetCreatedBy()
	guuo.mutation.SetCreatedBy(i)
	return guuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (guuo *GameUserUpdateOne) SetNillableCreatedBy(i *int64) *GameUserUpdateOne {
	if i != nil {
		guuo.SetCreatedBy(*i)
	}
	return guuo
}

// AddCreatedBy adds i to the "created_by" field.
func (guuo *GameUserUpdateOne) AddCreatedBy(i int64) *GameUserUpdateOne {
	guuo.mutation.AddCreatedBy(i)
	return guuo
}

// SetUpdatedBy sets the "updated_by" field.
func (guuo *GameUserUpdateOne) SetUpdatedBy(i int64) *GameUserUpdateOne {
	guuo.mutation.ResetUpdatedBy()
	guuo.mutation.SetUpdatedBy(i)
	return guuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (guuo *GameUserUpdateOne) SetNillableUpdatedBy(i *int64) *GameUserUpdateOne {
	if i != nil {
		guuo.SetUpdatedBy(*i)
	}
	return guuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (guuo *GameUserUpdateOne) AddUpdatedBy(i int64) *GameUserUpdateOne {
	guuo.mutation.AddUpdatedBy(i)
	return guuo
}

// SetUpdatedAt sets the "updated_at" field.
func (guuo *GameUserUpdateOne) SetUpdatedAt(t time.Time) *GameUserUpdateOne {
	guuo.mutation.SetUpdatedAt(t)
	return guuo
}

// SetDeletedAt sets the "deleted_at" field.
func (guuo *GameUserUpdateOne) SetDeletedAt(t time.Time) *GameUserUpdateOne {
	guuo.mutation.SetDeletedAt(t)
	return guuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guuo *GameUserUpdateOne) SetNillableDeletedAt(t *time.Time) *GameUserUpdateOne {
	if t != nil {
		guuo.SetDeletedAt(*t)
	}
	return guuo
}

// SetUserID sets the "user_id" field.
func (guuo *GameUserUpdateOne) SetUserID(i int64) *GameUserUpdateOne {
	guuo.mutation.ResetUserID()
	guuo.mutation.SetUserID(i)
	return guuo
}

// AddUserID adds i to the "user_id" field.
func (guuo *GameUserUpdateOne) AddUserID(i int64) *GameUserUpdateOne {
	guuo.mutation.AddUserID(i)
	return guuo
}

// SetGameID sets the "game_id" field.
func (guuo *GameUserUpdateOne) SetGameID(i int64) *GameUserUpdateOne {
	guuo.mutation.SetGameID(i)
	return guuo
}

// SetGame sets the "game" edge to the Game entity.
func (guuo *GameUserUpdateOne) SetGame(g *Game) *GameUserUpdateOne {
	return guuo.SetGameID(g.ID)
}

// Mutation returns the GameUserMutation object of the builder.
func (guuo *GameUserUpdateOne) Mutation() *GameUserMutation {
	return guuo.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (guuo *GameUserUpdateOne) ClearGame() *GameUserUpdateOne {
	guuo.mutation.ClearGame()
	return guuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guuo *GameUserUpdateOne) Select(field string, fields ...string) *GameUserUpdateOne {
	guuo.fields = append([]string{field}, fields...)
	return guuo
}

// Save executes the query and returns the updated GameUser entity.
func (guuo *GameUserUpdateOne) Save(ctx context.Context) (*GameUser, error) {
	var (
		err  error
		node *GameUser
	)
	guuo.defaults()
	if len(guuo.hooks) == 0 {
		if err = guuo.check(); err != nil {
			return nil, err
		}
		node, err = guuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guuo.check(); err != nil {
				return nil, err
			}
			guuo.mutation = mutation
			node, err = guuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guuo.hooks) - 1; i >= 0; i-- {
			if guuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GameUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GameUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guuo *GameUserUpdateOne) SaveX(ctx context.Context) *GameUser {
	node, err := guuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guuo *GameUserUpdateOne) Exec(ctx context.Context) error {
	_, err := guuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guuo *GameUserUpdateOne) ExecX(ctx context.Context) {
	if err := guuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guuo *GameUserUpdateOne) defaults() {
	if _, ok := guuo.mutation.UpdatedAt(); !ok {
		v := gameuser.UpdateDefaultUpdatedAt()
		guuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guuo *GameUserUpdateOne) check() error {
	if _, ok := guuo.mutation.GameID(); guuo.mutation.GameCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GameUser.game"`)
	}
	return nil
}

func (guuo *GameUserUpdateOne) sqlSave(ctx context.Context) (_node *GameUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gameuser.Table,
			Columns: gameuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gameuser.FieldID,
			},
		},
	}
	id, ok := guuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GameUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameuser.FieldID)
		for _, f := range fields {
			if !gameuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gameuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guuo.mutation.CreatedBy(); ok {
		_spec.SetField(gameuser.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(gameuser.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guuo.mutation.UpdatedBy(); ok {
		_spec.SetField(gameuser.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(gameuser.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guuo.mutation.UpdatedAt(); ok {
		_spec.SetField(gameuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guuo.mutation.DeletedAt(); ok {
		_spec.SetField(gameuser.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := guuo.mutation.UserID(); ok {
		_spec.SetField(gameuser.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := guuo.mutation.AddedUserID(); ok {
		_spec.AddField(gameuser.FieldUserID, field.TypeInt64, value)
	}
	if guuo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameuser.GameTable,
			Columns: []string{gameuser.GameColumn},
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
	if nodes := guuo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameuser.GameTable,
			Columns: []string{gameuser.GameColumn},
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
	_node = &GameUser{config: guuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
