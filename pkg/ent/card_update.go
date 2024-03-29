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
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedBy sets the "created_by" field.
func (cu *CardUpdate) SetCreatedBy(i int64) *CardUpdate {
	cu.mutation.ResetCreatedBy()
	cu.mutation.SetCreatedBy(i)
	return cu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cu *CardUpdate) SetNillableCreatedBy(i *int64) *CardUpdate {
	if i != nil {
		cu.SetCreatedBy(*i)
	}
	return cu
}

// AddCreatedBy adds i to the "created_by" field.
func (cu *CardUpdate) AddCreatedBy(i int64) *CardUpdate {
	cu.mutation.AddCreatedBy(i)
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *CardUpdate) SetUpdatedBy(i int64) *CardUpdate {
	cu.mutation.ResetUpdatedBy()
	cu.mutation.SetUpdatedBy(i)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *CardUpdate) SetNillableUpdatedBy(i *int64) *CardUpdate {
	if i != nil {
		cu.SetUpdatedBy(*i)
	}
	return cu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (cu *CardUpdate) AddUpdatedBy(i int64) *CardUpdate {
	cu.mutation.AddUpdatedBy(i)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CardUpdate) SetUpdatedAt(t time.Time) *CardUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CardUpdate) SetDeletedAt(t time.Time) *CardUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CardUpdate) SetNillableDeletedAt(t *time.Time) *CardUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// SetName sets the "name" field.
func (cu *CardUpdate) SetName(c card.Name) *CardUpdate {
	cu.mutation.SetName(c)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CardUpdate) SetNillableName(c *card.Name) *CardUpdate {
	if c != nil {
		cu.SetName(*c)
	}
	return cu
}

// SetRole sets the "role" field.
func (cu *CardUpdate) SetRole(c card.Role) *CardUpdate {
	cu.mutation.SetRole(c)
	return cu
}

// SetTale sets the "tale" field.
func (cu *CardUpdate) SetTale(s string) *CardUpdate {
	cu.mutation.SetTale(s)
	return cu
}

// SetNillableTale sets the "tale" field if the given value is not nil.
func (cu *CardUpdate) SetNillableTale(s *string) *CardUpdate {
	if s != nil {
		cu.SetTale(*s)
	}
	return cu
}

// SetRed sets the "red" field.
func (cu *CardUpdate) SetRed(b bool) *CardUpdate {
	cu.mutation.SetRed(b)
	return cu
}

// SetNillableRed sets the "red" field if the given value is not nil.
func (cu *CardUpdate) SetNillableRed(b *bool) *CardUpdate {
	if b != nil {
		cu.SetRed(*b)
	}
	return cu
}

// AddGameUserIDs adds the "game_users" edge to the GameUser entity by IDs.
func (cu *CardUpdate) AddGameUserIDs(ids ...int64) *CardUpdate {
	cu.mutation.AddGameUserIDs(ids...)
	return cu
}

// AddGameUsers adds the "game_users" edges to the GameUser entity.
func (cu *CardUpdate) AddGameUsers(g ...*GameUser) *CardUpdate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cu.AddGameUserIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cu *CardUpdate) Mutation() *CardMutation {
	return cu.mutation
}

// ClearGameUsers clears all "game_users" edges to the GameUser entity.
func (cu *CardUpdate) ClearGameUsers() *CardUpdate {
	cu.mutation.ClearGameUsers()
	return cu
}

// RemoveGameUserIDs removes the "game_users" edge to GameUser entities by IDs.
func (cu *CardUpdate) RemoveGameUserIDs(ids ...int64) *CardUpdate {
	cu.mutation.RemoveGameUserIDs(ids...)
	return cu
}

// RemoveGameUsers removes "game_users" edges to GameUser entities.
func (cu *CardUpdate) RemoveGameUsers(g ...*GameUser) *CardUpdate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cu.RemoveGameUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CardUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := card.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CardUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Role(); ok {
		if err := card.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Card.role": %w`, err)}
		}
	}
	return nil
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: card.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedBy(); ok {
		_spec.SetField(card.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(card.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(card.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(card.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(card.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(card.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Role(); ok {
		_spec.SetField(card.FieldRole, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Tale(); ok {
		_spec.SetField(card.FieldTale, field.TypeString, value)
	}
	if value, ok := cu.mutation.Red(); ok {
		_spec.SetField(card.FieldRed, field.TypeBool, value)
	}
	if cu.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.GameUsersTable,
			Columns: []string{card.GameUsersColumn},
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
	if nodes := cu.mutation.RemovedGameUsersIDs(); len(nodes) > 0 && !cu.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.GameUsersTable,
			Columns: []string{card.GameUsersColumn},
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
	if nodes := cu.mutation.GameUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.GameUsersTable,
			Columns: []string{card.GameUsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetCreatedBy sets the "created_by" field.
func (cuo *CardUpdateOne) SetCreatedBy(i int64) *CardUpdateOne {
	cuo.mutation.ResetCreatedBy()
	cuo.mutation.SetCreatedBy(i)
	return cuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableCreatedBy(i *int64) *CardUpdateOne {
	if i != nil {
		cuo.SetCreatedBy(*i)
	}
	return cuo
}

// AddCreatedBy adds i to the "created_by" field.
func (cuo *CardUpdateOne) AddCreatedBy(i int64) *CardUpdateOne {
	cuo.mutation.AddCreatedBy(i)
	return cuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *CardUpdateOne) SetUpdatedBy(i int64) *CardUpdateOne {
	cuo.mutation.ResetUpdatedBy()
	cuo.mutation.SetUpdatedBy(i)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableUpdatedBy(i *int64) *CardUpdateOne {
	if i != nil {
		cuo.SetUpdatedBy(*i)
	}
	return cuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (cuo *CardUpdateOne) AddUpdatedBy(i int64) *CardUpdateOne {
	cuo.mutation.AddUpdatedBy(i)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CardUpdateOne) SetUpdatedAt(t time.Time) *CardUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CardUpdateOne) SetDeletedAt(t time.Time) *CardUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableDeletedAt(t *time.Time) *CardUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// SetName sets the "name" field.
func (cuo *CardUpdateOne) SetName(c card.Name) *CardUpdateOne {
	cuo.mutation.SetName(c)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableName(c *card.Name) *CardUpdateOne {
	if c != nil {
		cuo.SetName(*c)
	}
	return cuo
}

// SetRole sets the "role" field.
func (cuo *CardUpdateOne) SetRole(c card.Role) *CardUpdateOne {
	cuo.mutation.SetRole(c)
	return cuo
}

// SetTale sets the "tale" field.
func (cuo *CardUpdateOne) SetTale(s string) *CardUpdateOne {
	cuo.mutation.SetTale(s)
	return cuo
}

// SetNillableTale sets the "tale" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableTale(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetTale(*s)
	}
	return cuo
}

// SetRed sets the "red" field.
func (cuo *CardUpdateOne) SetRed(b bool) *CardUpdateOne {
	cuo.mutation.SetRed(b)
	return cuo
}

// SetNillableRed sets the "red" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableRed(b *bool) *CardUpdateOne {
	if b != nil {
		cuo.SetRed(*b)
	}
	return cuo
}

// AddGameUserIDs adds the "game_users" edge to the GameUser entity by IDs.
func (cuo *CardUpdateOne) AddGameUserIDs(ids ...int64) *CardUpdateOne {
	cuo.mutation.AddGameUserIDs(ids...)
	return cuo
}

// AddGameUsers adds the "game_users" edges to the GameUser entity.
func (cuo *CardUpdateOne) AddGameUsers(g ...*GameUser) *CardUpdateOne {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cuo.AddGameUserIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cuo *CardUpdateOne) Mutation() *CardMutation {
	return cuo.mutation
}

// ClearGameUsers clears all "game_users" edges to the GameUser entity.
func (cuo *CardUpdateOne) ClearGameUsers() *CardUpdateOne {
	cuo.mutation.ClearGameUsers()
	return cuo
}

// RemoveGameUserIDs removes the "game_users" edge to GameUser entities by IDs.
func (cuo *CardUpdateOne) RemoveGameUserIDs(ids ...int64) *CardUpdateOne {
	cuo.mutation.RemoveGameUserIDs(ids...)
	return cuo
}

// RemoveGameUsers removes "game_users" edges to GameUser entities.
func (cuo *CardUpdateOne) RemoveGameUsers(g ...*GameUser) *CardUpdateOne {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cuo.RemoveGameUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Card entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	var (
		err  error
		node *Card
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Card)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CardUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := card.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CardUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Role(); ok {
		if err := card.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Card.role": %w`, err)}
		}
	}
	return nil
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: card.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedBy(); ok {
		_spec.SetField(card.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(card.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(card.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(card.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(card.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(card.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Role(); ok {
		_spec.SetField(card.FieldRole, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Tale(); ok {
		_spec.SetField(card.FieldTale, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Red(); ok {
		_spec.SetField(card.FieldRed, field.TypeBool, value)
	}
	if cuo.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.GameUsersTable,
			Columns: []string{card.GameUsersColumn},
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
	if nodes := cuo.mutation.RemovedGameUsersIDs(); len(nodes) > 0 && !cuo.mutation.GameUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.GameUsersTable,
			Columns: []string{card.GameUsersColumn},
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
	if nodes := cuo.mutation.GameUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   card.GameUsersTable,
			Columns: []string{card.GameUsersColumn},
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
	_node = &Card{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
