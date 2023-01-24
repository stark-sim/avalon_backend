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
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
	"github.com/stark-sim/avalon_backend/pkg/ent/squad"
	"github.com/stark-sim/avalon_backend/pkg/ent/vote"
)

// MissionUpdate is the builder for updating Mission entities.
type MissionUpdate struct {
	config
	hooks    []Hook
	mutation *MissionMutation
}

// Where appends a list predicates to the MissionUpdate builder.
func (mu *MissionUpdate) Where(ps ...predicate.Mission) *MissionUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedBy sets the "created_by" field.
func (mu *MissionUpdate) SetCreatedBy(i int64) *MissionUpdate {
	mu.mutation.ResetCreatedBy()
	mu.mutation.SetCreatedBy(i)
	return mu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableCreatedBy(i *int64) *MissionUpdate {
	if i != nil {
		mu.SetCreatedBy(*i)
	}
	return mu
}

// AddCreatedBy adds i to the "created_by" field.
func (mu *MissionUpdate) AddCreatedBy(i int64) *MissionUpdate {
	mu.mutation.AddCreatedBy(i)
	return mu
}

// SetUpdatedBy sets the "updated_by" field.
func (mu *MissionUpdate) SetUpdatedBy(i int64) *MissionUpdate {
	mu.mutation.ResetUpdatedBy()
	mu.mutation.SetUpdatedBy(i)
	return mu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableUpdatedBy(i *int64) *MissionUpdate {
	if i != nil {
		mu.SetUpdatedBy(*i)
	}
	return mu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mu *MissionUpdate) AddUpdatedBy(i int64) *MissionUpdate {
	mu.mutation.AddUpdatedBy(i)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MissionUpdate) SetUpdatedAt(t time.Time) *MissionUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MissionUpdate) SetDeletedAt(t time.Time) *MissionUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableDeletedAt(t *time.Time) *MissionUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// SetSequence sets the "sequence" field.
func (mu *MissionUpdate) SetSequence(u uint8) *MissionUpdate {
	mu.mutation.ResetSequence()
	mu.mutation.SetSequence(u)
	return mu
}

// AddSequence adds u to the "sequence" field.
func (mu *MissionUpdate) AddSequence(u int8) *MissionUpdate {
	mu.mutation.AddSequence(u)
	return mu
}

// SetStatus sets the "status" field.
func (mu *MissionUpdate) SetStatus(m mission.Status) *MissionUpdate {
	mu.mutation.SetStatus(m)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableStatus(m *mission.Status) *MissionUpdate {
	if m != nil {
		mu.SetStatus(*m)
	}
	return mu
}

// SetFailed sets the "failed" field.
func (mu *MissionUpdate) SetFailed(b bool) *MissionUpdate {
	mu.mutation.SetFailed(b)
	return mu
}

// SetNillableFailed sets the "failed" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableFailed(b *bool) *MissionUpdate {
	if b != nil {
		mu.SetFailed(*b)
	}
	return mu
}

// SetGameID sets the "game_id" field.
func (mu *MissionUpdate) SetGameID(i int64) *MissionUpdate {
	mu.mutation.SetGameID(i)
	return mu
}

// SetCapacity sets the "capacity" field.
func (mu *MissionUpdate) SetCapacity(u uint8) *MissionUpdate {
	mu.mutation.ResetCapacity()
	mu.mutation.SetCapacity(u)
	return mu
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableCapacity(u *uint8) *MissionUpdate {
	if u != nil {
		mu.SetCapacity(*u)
	}
	return mu
}

// AddCapacity adds u to the "capacity" field.
func (mu *MissionUpdate) AddCapacity(u int8) *MissionUpdate {
	mu.mutation.AddCapacity(u)
	return mu
}

// SetLeaderID sets the "leader_id" field.
func (mu *MissionUpdate) SetLeaderID(i int64) *MissionUpdate {
	mu.mutation.ResetLeaderID()
	mu.mutation.SetLeaderID(i)
	return mu
}

// SetNillableLeaderID sets the "leader_id" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableLeaderID(i *int64) *MissionUpdate {
	if i != nil {
		mu.SetLeaderID(*i)
	}
	return mu
}

// AddLeaderID adds i to the "leader_id" field.
func (mu *MissionUpdate) AddLeaderID(i int64) *MissionUpdate {
	mu.mutation.AddLeaderID(i)
	return mu
}

// SetProtected sets the "protected" field.
func (mu *MissionUpdate) SetProtected(b bool) *MissionUpdate {
	mu.mutation.SetProtected(b)
	return mu
}

// SetNillableProtected sets the "protected" field if the given value is not nil.
func (mu *MissionUpdate) SetNillableProtected(b *bool) *MissionUpdate {
	if b != nil {
		mu.SetProtected(*b)
	}
	return mu
}

// SetGame sets the "game" edge to the Game entity.
func (mu *MissionUpdate) SetGame(g *Game) *MissionUpdate {
	return mu.SetGameID(g.ID)
}

// AddSquadIDs adds the "squads" edge to the Squad entity by IDs.
func (mu *MissionUpdate) AddSquadIDs(ids ...int64) *MissionUpdate {
	mu.mutation.AddSquadIDs(ids...)
	return mu
}

// AddSquads adds the "squads" edges to the Squad entity.
func (mu *MissionUpdate) AddSquads(s ...*Squad) *MissionUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.AddSquadIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (mu *MissionUpdate) AddVoteIDs(ids ...int64) *MissionUpdate {
	mu.mutation.AddVoteIDs(ids...)
	return mu
}

// AddVotes adds the "votes" edges to the Vote entity.
func (mu *MissionUpdate) AddVotes(v ...*Vote) *MissionUpdate {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return mu.AddVoteIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (mu *MissionUpdate) Mutation() *MissionMutation {
	return mu.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (mu *MissionUpdate) ClearGame() *MissionUpdate {
	mu.mutation.ClearGame()
	return mu
}

// ClearSquads clears all "squads" edges to the Squad entity.
func (mu *MissionUpdate) ClearSquads() *MissionUpdate {
	mu.mutation.ClearSquads()
	return mu
}

// RemoveSquadIDs removes the "squads" edge to Squad entities by IDs.
func (mu *MissionUpdate) RemoveSquadIDs(ids ...int64) *MissionUpdate {
	mu.mutation.RemoveSquadIDs(ids...)
	return mu
}

// RemoveSquads removes "squads" edges to Squad entities.
func (mu *MissionUpdate) RemoveSquads(s ...*Squad) *MissionUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.RemoveSquadIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (mu *MissionUpdate) ClearVotes() *MissionUpdate {
	mu.mutation.ClearVotes()
	return mu
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (mu *MissionUpdate) RemoveVoteIDs(ids ...int64) *MissionUpdate {
	mu.mutation.RemoveVoteIDs(ids...)
	return mu
}

// RemoveVotes removes "votes" edges to Vote entities.
func (mu *MissionUpdate) RemoveVotes(v ...*Vote) *MissionUpdate {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return mu.RemoveVoteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MissionUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MissionUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MissionUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MissionUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := mission.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MissionUpdate) check() error {
	if v, ok := mu.mutation.Sequence(); ok {
		if err := mission.SequenceValidator(v); err != nil {
			return &ValidationError{Name: "sequence", err: fmt.Errorf(`ent: validator failed for field "Mission.sequence": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Status(); ok {
		if err := mission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Mission.status": %w`, err)}
		}
	}
	if _, ok := mu.mutation.GameID(); mu.mutation.GameCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mission.game"`)
	}
	return nil
}

func (mu *MissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mission.Table,
			Columns: mission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: mission.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedBy(); ok {
		_spec.SetField(mission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(mission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.UpdatedBy(); ok {
		_spec.SetField(mission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(mission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(mission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(mission.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Sequence(); ok {
		_spec.SetField(mission.FieldSequence, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.AddedSequence(); ok {
		_spec.AddField(mission.FieldSequence, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(mission.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.Failed(); ok {
		_spec.SetField(mission.FieldFailed, field.TypeBool, value)
	}
	if value, ok := mu.mutation.Capacity(); ok {
		_spec.SetField(mission.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.AddedCapacity(); ok {
		_spec.AddField(mission.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.LeaderID(); ok {
		_spec.SetField(mission.FieldLeaderID, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedLeaderID(); ok {
		_spec.AddField(mission.FieldLeaderID, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.Protected(); ok {
		_spec.SetField(mission.FieldProtected, field.TypeBool, value)
	}
	if mu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mission.GameTable,
			Columns: []string{mission.GameColumn},
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
	if nodes := mu.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mission.GameTable,
			Columns: []string{mission.GameColumn},
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
	if mu.mutation.SquadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.SquadsTable,
			Columns: []string{mission.SquadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: squad.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedSquadsIDs(); len(nodes) > 0 && !mu.mutation.SquadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.SquadsTable,
			Columns: []string{mission.SquadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: squad.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.SquadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.SquadsTable,
			Columns: []string{mission.SquadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: squad.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.VotesTable,
			Columns: []string{mission.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: vote.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedVotesIDs(); len(nodes) > 0 && !mu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.VotesTable,
			Columns: []string{mission.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: vote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.VotesTable,
			Columns: []string{mission.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: vote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MissionUpdateOne is the builder for updating a single Mission entity.
type MissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MissionMutation
}

// SetCreatedBy sets the "created_by" field.
func (muo *MissionUpdateOne) SetCreatedBy(i int64) *MissionUpdateOne {
	muo.mutation.ResetCreatedBy()
	muo.mutation.SetCreatedBy(i)
	return muo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableCreatedBy(i *int64) *MissionUpdateOne {
	if i != nil {
		muo.SetCreatedBy(*i)
	}
	return muo
}

// AddCreatedBy adds i to the "created_by" field.
func (muo *MissionUpdateOne) AddCreatedBy(i int64) *MissionUpdateOne {
	muo.mutation.AddCreatedBy(i)
	return muo
}

// SetUpdatedBy sets the "updated_by" field.
func (muo *MissionUpdateOne) SetUpdatedBy(i int64) *MissionUpdateOne {
	muo.mutation.ResetUpdatedBy()
	muo.mutation.SetUpdatedBy(i)
	return muo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableUpdatedBy(i *int64) *MissionUpdateOne {
	if i != nil {
		muo.SetUpdatedBy(*i)
	}
	return muo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (muo *MissionUpdateOne) AddUpdatedBy(i int64) *MissionUpdateOne {
	muo.mutation.AddUpdatedBy(i)
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MissionUpdateOne) SetUpdatedAt(t time.Time) *MissionUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MissionUpdateOne) SetDeletedAt(t time.Time) *MissionUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableDeletedAt(t *time.Time) *MissionUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// SetSequence sets the "sequence" field.
func (muo *MissionUpdateOne) SetSequence(u uint8) *MissionUpdateOne {
	muo.mutation.ResetSequence()
	muo.mutation.SetSequence(u)
	return muo
}

// AddSequence adds u to the "sequence" field.
func (muo *MissionUpdateOne) AddSequence(u int8) *MissionUpdateOne {
	muo.mutation.AddSequence(u)
	return muo
}

// SetStatus sets the "status" field.
func (muo *MissionUpdateOne) SetStatus(m mission.Status) *MissionUpdateOne {
	muo.mutation.SetStatus(m)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableStatus(m *mission.Status) *MissionUpdateOne {
	if m != nil {
		muo.SetStatus(*m)
	}
	return muo
}

// SetFailed sets the "failed" field.
func (muo *MissionUpdateOne) SetFailed(b bool) *MissionUpdateOne {
	muo.mutation.SetFailed(b)
	return muo
}

// SetNillableFailed sets the "failed" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableFailed(b *bool) *MissionUpdateOne {
	if b != nil {
		muo.SetFailed(*b)
	}
	return muo
}

// SetGameID sets the "game_id" field.
func (muo *MissionUpdateOne) SetGameID(i int64) *MissionUpdateOne {
	muo.mutation.SetGameID(i)
	return muo
}

// SetCapacity sets the "capacity" field.
func (muo *MissionUpdateOne) SetCapacity(u uint8) *MissionUpdateOne {
	muo.mutation.ResetCapacity()
	muo.mutation.SetCapacity(u)
	return muo
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableCapacity(u *uint8) *MissionUpdateOne {
	if u != nil {
		muo.SetCapacity(*u)
	}
	return muo
}

// AddCapacity adds u to the "capacity" field.
func (muo *MissionUpdateOne) AddCapacity(u int8) *MissionUpdateOne {
	muo.mutation.AddCapacity(u)
	return muo
}

// SetLeaderID sets the "leader_id" field.
func (muo *MissionUpdateOne) SetLeaderID(i int64) *MissionUpdateOne {
	muo.mutation.ResetLeaderID()
	muo.mutation.SetLeaderID(i)
	return muo
}

// SetNillableLeaderID sets the "leader_id" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableLeaderID(i *int64) *MissionUpdateOne {
	if i != nil {
		muo.SetLeaderID(*i)
	}
	return muo
}

// AddLeaderID adds i to the "leader_id" field.
func (muo *MissionUpdateOne) AddLeaderID(i int64) *MissionUpdateOne {
	muo.mutation.AddLeaderID(i)
	return muo
}

// SetProtected sets the "protected" field.
func (muo *MissionUpdateOne) SetProtected(b bool) *MissionUpdateOne {
	muo.mutation.SetProtected(b)
	return muo
}

// SetNillableProtected sets the "protected" field if the given value is not nil.
func (muo *MissionUpdateOne) SetNillableProtected(b *bool) *MissionUpdateOne {
	if b != nil {
		muo.SetProtected(*b)
	}
	return muo
}

// SetGame sets the "game" edge to the Game entity.
func (muo *MissionUpdateOne) SetGame(g *Game) *MissionUpdateOne {
	return muo.SetGameID(g.ID)
}

// AddSquadIDs adds the "squads" edge to the Squad entity by IDs.
func (muo *MissionUpdateOne) AddSquadIDs(ids ...int64) *MissionUpdateOne {
	muo.mutation.AddSquadIDs(ids...)
	return muo
}

// AddSquads adds the "squads" edges to the Squad entity.
func (muo *MissionUpdateOne) AddSquads(s ...*Squad) *MissionUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.AddSquadIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (muo *MissionUpdateOne) AddVoteIDs(ids ...int64) *MissionUpdateOne {
	muo.mutation.AddVoteIDs(ids...)
	return muo
}

// AddVotes adds the "votes" edges to the Vote entity.
func (muo *MissionUpdateOne) AddVotes(v ...*Vote) *MissionUpdateOne {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return muo.AddVoteIDs(ids...)
}

// Mutation returns the MissionMutation object of the builder.
func (muo *MissionUpdateOne) Mutation() *MissionMutation {
	return muo.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (muo *MissionUpdateOne) ClearGame() *MissionUpdateOne {
	muo.mutation.ClearGame()
	return muo
}

// ClearSquads clears all "squads" edges to the Squad entity.
func (muo *MissionUpdateOne) ClearSquads() *MissionUpdateOne {
	muo.mutation.ClearSquads()
	return muo
}

// RemoveSquadIDs removes the "squads" edge to Squad entities by IDs.
func (muo *MissionUpdateOne) RemoveSquadIDs(ids ...int64) *MissionUpdateOne {
	muo.mutation.RemoveSquadIDs(ids...)
	return muo
}

// RemoveSquads removes "squads" edges to Squad entities.
func (muo *MissionUpdateOne) RemoveSquads(s ...*Squad) *MissionUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.RemoveSquadIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (muo *MissionUpdateOne) ClearVotes() *MissionUpdateOne {
	muo.mutation.ClearVotes()
	return muo
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (muo *MissionUpdateOne) RemoveVoteIDs(ids ...int64) *MissionUpdateOne {
	muo.mutation.RemoveVoteIDs(ids...)
	return muo
}

// RemoveVotes removes "votes" edges to Vote entities.
func (muo *MissionUpdateOne) RemoveVotes(v ...*Vote) *MissionUpdateOne {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return muo.RemoveVoteIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MissionUpdateOne) Select(field string, fields ...string) *MissionUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mission entity.
func (muo *MissionUpdateOne) Save(ctx context.Context) (*Mission, error) {
	var (
		err  error
		node *Mission
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Mission)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MissionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MissionUpdateOne) SaveX(ctx context.Context) *Mission {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MissionUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MissionUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MissionUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := mission.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MissionUpdateOne) check() error {
	if v, ok := muo.mutation.Sequence(); ok {
		if err := mission.SequenceValidator(v); err != nil {
			return &ValidationError{Name: "sequence", err: fmt.Errorf(`ent: validator failed for field "Mission.sequence": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Status(); ok {
		if err := mission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Mission.status": %w`, err)}
		}
	}
	if _, ok := muo.mutation.GameID(); muo.mutation.GameCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mission.game"`)
	}
	return nil
}

func (muo *MissionUpdateOne) sqlSave(ctx context.Context) (_node *Mission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mission.Table,
			Columns: mission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: mission.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mission.FieldID)
		for _, f := range fields {
			if !mission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedBy(); ok {
		_spec.SetField(mission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(mission.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.UpdatedBy(); ok {
		_spec.SetField(mission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(mission.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(mission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(mission.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Sequence(); ok {
		_spec.SetField(mission.FieldSequence, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.AddedSequence(); ok {
		_spec.AddField(mission.FieldSequence, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(mission.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.Failed(); ok {
		_spec.SetField(mission.FieldFailed, field.TypeBool, value)
	}
	if value, ok := muo.mutation.Capacity(); ok {
		_spec.SetField(mission.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.AddedCapacity(); ok {
		_spec.AddField(mission.FieldCapacity, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.LeaderID(); ok {
		_spec.SetField(mission.FieldLeaderID, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedLeaderID(); ok {
		_spec.AddField(mission.FieldLeaderID, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.Protected(); ok {
		_spec.SetField(mission.FieldProtected, field.TypeBool, value)
	}
	if muo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mission.GameTable,
			Columns: []string{mission.GameColumn},
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
	if nodes := muo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mission.GameTable,
			Columns: []string{mission.GameColumn},
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
	if muo.mutation.SquadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.SquadsTable,
			Columns: []string{mission.SquadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: squad.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedSquadsIDs(); len(nodes) > 0 && !muo.mutation.SquadsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.SquadsTable,
			Columns: []string{mission.SquadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: squad.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.SquadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.SquadsTable,
			Columns: []string{mission.SquadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: squad.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.VotesTable,
			Columns: []string{mission.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: vote.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedVotesIDs(); len(nodes) > 0 && !muo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.VotesTable,
			Columns: []string{mission.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: vote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mission.VotesTable,
			Columns: []string{mission.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: vote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Mission{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
