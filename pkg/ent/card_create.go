// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/pkg/ent/gameuser"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (cc *CardCreate) SetCreatedBy(i int64) *CardCreate {
	cc.mutation.SetCreatedBy(i)
	return cc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cc *CardCreate) SetNillableCreatedBy(i *int64) *CardCreate {
	if i != nil {
		cc.SetCreatedBy(*i)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *CardCreate) SetUpdatedBy(i int64) *CardCreate {
	cc.mutation.SetUpdatedBy(i)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdatedBy(i *int64) *CardCreate {
	if i != nil {
		cc.SetUpdatedBy(*i)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CardCreate) SetCreatedAt(t time.Time) *CardCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CardCreate) SetNillableCreatedAt(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CardCreate) SetUpdatedAt(t time.Time) *CardCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdatedAt(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CardCreate) SetDeletedAt(t time.Time) *CardCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CardCreate) SetNillableDeletedAt(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CardCreate) SetName(c card.Name) *CardCreate {
	cc.mutation.SetName(c)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CardCreate) SetNillableName(c *card.Name) *CardCreate {
	if c != nil {
		cc.SetName(*c)
	}
	return cc
}

// SetRole sets the "role" field.
func (cc *CardCreate) SetRole(c card.Role) *CardCreate {
	cc.mutation.SetRole(c)
	return cc
}

// SetTale sets the "tale" field.
func (cc *CardCreate) SetTale(s string) *CardCreate {
	cc.mutation.SetTale(s)
	return cc
}

// SetNillableTale sets the "tale" field if the given value is not nil.
func (cc *CardCreate) SetNillableTale(s *string) *CardCreate {
	if s != nil {
		cc.SetTale(*s)
	}
	return cc
}

// SetRed sets the "red" field.
func (cc *CardCreate) SetRed(b bool) *CardCreate {
	cc.mutation.SetRed(b)
	return cc
}

// SetNillableRed sets the "red" field if the given value is not nil.
func (cc *CardCreate) SetNillableRed(b *bool) *CardCreate {
	if b != nil {
		cc.SetRed(*b)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CardCreate) SetID(i int64) *CardCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CardCreate) SetNillableID(i *int64) *CardCreate {
	if i != nil {
		cc.SetID(*i)
	}
	return cc
}

// AddGameUserIDs adds the "game_users" edge to the GameUser entity by IDs.
func (cc *CardCreate) AddGameUserIDs(ids ...int64) *CardCreate {
	cc.mutation.AddGameUserIDs(ids...)
	return cc
}

// AddGameUsers adds the "game_users" edges to the GameUser entity.
func (cc *CardCreate) AddGameUsers(g ...*GameUser) *CardCreate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cc.AddGameUserIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	var (
		err  error
		node *Card
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CardCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CardCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.CreatedBy(); !ok {
		v := card.DefaultCreatedBy
		cc.mutation.SetCreatedBy(v)
	}
	if _, ok := cc.mutation.UpdatedBy(); !ok {
		v := card.DefaultUpdatedBy
		cc.mutation.SetUpdatedBy(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := card.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := card.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		v := card.DefaultDeletedAt
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.Name(); !ok {
		v := card.DefaultName
		cc.mutation.SetName(v)
	}
	if _, ok := cc.mutation.Tale(); !ok {
		v := card.DefaultTale
		cc.mutation.SetTale(v)
	}
	if _, ok := cc.mutation.Red(); !ok {
		v := card.DefaultRed
		cc.mutation.SetRed(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := card.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Card.created_by"`)}
	}
	if _, ok := cc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Card.updated_by"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Card.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Card.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Card.deleted_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Card.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "Card.role"`)}
	}
	if v, ok := cc.mutation.Role(); ok {
		if err := card.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Card.role": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Tale(); !ok {
		return &ValidationError{Name: "tale", err: errors.New(`ent: missing required field "Card.tale"`)}
	}
	if _, ok := cc.mutation.Red(); !ok {
		return &ValidationError{Name: "red", err: errors.New(`ent: missing required field "Card.red"`)}
	}
	return nil
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: card.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: card.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(card.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(card.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(card.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(card.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(card.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeEnum, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Role(); ok {
		_spec.SetField(card.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := cc.mutation.Tale(); ok {
		_spec.SetField(card.FieldTale, field.TypeString, value)
		_node.Tale = value
	}
	if value, ok := cc.mutation.Red(); ok {
		_spec.SetField(card.FieldRed, field.TypeBool, value)
		_node.Red = value
	}
	if nodes := cc.mutation.GameUsersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	builders []*CardCreate
}

// Save creates the Card entities in the database.
func (ccb *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Card, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CardCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CardCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
