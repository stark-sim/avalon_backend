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
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
	"github.com/stark-sim/avalon_backend/pkg/ent/record"
	"github.com/stark-sim/avalon_backend/pkg/ent/room"
)

// RecordUpdate is the builder for updating Record entities.
type RecordUpdate struct {
	config
	hooks    []Hook
	mutation *RecordMutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ru *RecordUpdate) Where(ps ...predicate.Record) *RecordUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCreatedBy sets the "created_by" field.
func (ru *RecordUpdate) SetCreatedBy(i int64) *RecordUpdate {
	ru.mutation.ResetCreatedBy()
	ru.mutation.SetCreatedBy(i)
	return ru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableCreatedBy(i *int64) *RecordUpdate {
	if i != nil {
		ru.SetCreatedBy(*i)
	}
	return ru
}

// AddCreatedBy adds i to the "created_by" field.
func (ru *RecordUpdate) AddCreatedBy(i int64) *RecordUpdate {
	ru.mutation.AddCreatedBy(i)
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RecordUpdate) SetUpdatedBy(i int64) *RecordUpdate {
	ru.mutation.ResetUpdatedBy()
	ru.mutation.SetUpdatedBy(i)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableUpdatedBy(i *int64) *RecordUpdate {
	if i != nil {
		ru.SetUpdatedBy(*i)
	}
	return ru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ru *RecordUpdate) AddUpdatedBy(i int64) *RecordUpdate {
	ru.mutation.AddUpdatedBy(i)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RecordUpdate) SetUpdatedAt(t time.Time) *RecordUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RecordUpdate) SetDeletedAt(t time.Time) *RecordUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableDeletedAt(t *time.Time) *RecordUpdate {
	if t != nil {
		ru.SetDeletedAt(*t)
	}
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RecordUpdate) SetUserID(i int64) *RecordUpdate {
	ru.mutation.ResetUserID()
	ru.mutation.SetUserID(i)
	return ru
}

// AddUserID adds i to the "user_id" field.
func (ru *RecordUpdate) AddUserID(i int64) *RecordUpdate {
	ru.mutation.AddUserID(i)
	return ru
}

// SetRoomID sets the "room_id" field.
func (ru *RecordUpdate) SetRoomID(i int64) *RecordUpdate {
	ru.mutation.SetRoomID(i)
	return ru
}

// SetScore sets the "score" field.
func (ru *RecordUpdate) SetScore(i int32) *RecordUpdate {
	ru.mutation.ResetScore()
	ru.mutation.SetScore(i)
	return ru
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableScore(i *int32) *RecordUpdate {
	if i != nil {
		ru.SetScore(*i)
	}
	return ru
}

// AddScore adds i to the "score" field.
func (ru *RecordUpdate) AddScore(i int32) *RecordUpdate {
	ru.mutation.AddScore(i)
	return ru
}

// SetRoom sets the "room" edge to the Room entity.
func (ru *RecordUpdate) SetRoom(r *Room) *RecordUpdate {
	return ru.SetRoomID(r.ID)
}

// Mutation returns the RecordMutation object of the builder.
func (ru *RecordUpdate) Mutation() *RecordMutation {
	return ru.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (ru *RecordUpdate) ClearRoom() *RecordUpdate {
	ru.mutation.ClearRoom()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
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
func (ru *RecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RecordUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := record.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RecordUpdate) check() error {
	if _, ok := ru.mutation.RoomID(); ru.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Record.room"`)
	}
	return nil
}

func (ru *RecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   record.Table,
			Columns: record.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: record.FieldID,
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
		_spec.SetField(record.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedCreatedBy(); ok {
		_spec.AddField(record.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.SetField(record.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(record.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(record.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(record.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UserID(); ok {
		_spec.SetField(record.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedUserID(); ok {
		_spec.AddField(record.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.Score(); ok {
		_spec.SetField(record.FieldScore, field.TypeInt32, value)
	}
	if value, ok := ru.mutation.AddedScore(); ok {
		_spec.AddField(record.FieldScore, field.TypeInt32, value)
	}
	if ru.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.RoomTable,
			Columns: []string{record.RoomColumn},
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
	if nodes := ru.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.RoomTable,
			Columns: []string{record.RoomColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RecordUpdateOne is the builder for updating a single Record entity.
type RecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecordMutation
}

// SetCreatedBy sets the "created_by" field.
func (ruo *RecordUpdateOne) SetCreatedBy(i int64) *RecordUpdateOne {
	ruo.mutation.ResetCreatedBy()
	ruo.mutation.SetCreatedBy(i)
	return ruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableCreatedBy(i *int64) *RecordUpdateOne {
	if i != nil {
		ruo.SetCreatedBy(*i)
	}
	return ruo
}

// AddCreatedBy adds i to the "created_by" field.
func (ruo *RecordUpdateOne) AddCreatedBy(i int64) *RecordUpdateOne {
	ruo.mutation.AddCreatedBy(i)
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RecordUpdateOne) SetUpdatedBy(i int64) *RecordUpdateOne {
	ruo.mutation.ResetUpdatedBy()
	ruo.mutation.SetUpdatedBy(i)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableUpdatedBy(i *int64) *RecordUpdateOne {
	if i != nil {
		ruo.SetUpdatedBy(*i)
	}
	return ruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ruo *RecordUpdateOne) AddUpdatedBy(i int64) *RecordUpdateOne {
	ruo.mutation.AddUpdatedBy(i)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RecordUpdateOne) SetUpdatedAt(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RecordUpdateOne) SetDeletedAt(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableDeletedAt(t *time.Time) *RecordUpdateOne {
	if t != nil {
		ruo.SetDeletedAt(*t)
	}
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RecordUpdateOne) SetUserID(i int64) *RecordUpdateOne {
	ruo.mutation.ResetUserID()
	ruo.mutation.SetUserID(i)
	return ruo
}

// AddUserID adds i to the "user_id" field.
func (ruo *RecordUpdateOne) AddUserID(i int64) *RecordUpdateOne {
	ruo.mutation.AddUserID(i)
	return ruo
}

// SetRoomID sets the "room_id" field.
func (ruo *RecordUpdateOne) SetRoomID(i int64) *RecordUpdateOne {
	ruo.mutation.SetRoomID(i)
	return ruo
}

// SetScore sets the "score" field.
func (ruo *RecordUpdateOne) SetScore(i int32) *RecordUpdateOne {
	ruo.mutation.ResetScore()
	ruo.mutation.SetScore(i)
	return ruo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableScore(i *int32) *RecordUpdateOne {
	if i != nil {
		ruo.SetScore(*i)
	}
	return ruo
}

// AddScore adds i to the "score" field.
func (ruo *RecordUpdateOne) AddScore(i int32) *RecordUpdateOne {
	ruo.mutation.AddScore(i)
	return ruo
}

// SetRoom sets the "room" edge to the Room entity.
func (ruo *RecordUpdateOne) SetRoom(r *Room) *RecordUpdateOne {
	return ruo.SetRoomID(r.ID)
}

// Mutation returns the RecordMutation object of the builder.
func (ruo *RecordUpdateOne) Mutation() *RecordMutation {
	return ruo.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (ruo *RecordUpdateOne) ClearRoom() *RecordUpdateOne {
	ruo.mutation.ClearRoom()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecordUpdateOne) Select(field string, fields ...string) *RecordUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Record entity.
func (ruo *RecordUpdateOne) Save(ctx context.Context) (*Record, error) {
	var (
		err  error
		node *Record
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
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
		nv, ok := v.(*Record)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RecordMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordUpdateOne) SaveX(ctx context.Context) *Record {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RecordUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := record.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RecordUpdateOne) check() error {
	if _, ok := ruo.mutation.RoomID(); ruo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Record.room"`)
	}
	return nil
}

func (ruo *RecordUpdateOne) sqlSave(ctx context.Context) (_node *Record, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   record.Table,
			Columns: record.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: record.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Record.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, record.FieldID)
		for _, f := range fields {
			if !record.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != record.FieldID {
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
		_spec.SetField(record.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(record.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.SetField(record.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(record.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(record.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(record.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UserID(); ok {
		_spec.SetField(record.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedUserID(); ok {
		_spec.AddField(record.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.Score(); ok {
		_spec.SetField(record.FieldScore, field.TypeInt32, value)
	}
	if value, ok := ruo.mutation.AddedScore(); ok {
		_spec.AddField(record.FieldScore, field.TypeInt32, value)
	}
	if ruo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.RoomTable,
			Columns: []string{record.RoomColumn},
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
	if nodes := ruo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.RoomTable,
			Columns: []string{record.RoomColumn},
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
	_node = &Record{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
