// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/avalon_backend/pkg/ent/mission"
	"github.com/stark-sim/avalon_backend/pkg/ent/predicate"
	"github.com/stark-sim/avalon_backend/pkg/ent/vote"
)

// VoteQuery is the builder for querying Vote entities.
type VoteQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.Vote
	withMission *MissionQuery
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*Vote) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoteQuery builder.
func (vq *VoteQuery) Where(ps ...predicate.Vote) *VoteQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VoteQuery) Limit(limit int) *VoteQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VoteQuery) Offset(offset int) *VoteQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VoteQuery) Unique(unique bool) *VoteQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *VoteQuery) Order(o ...OrderFunc) *VoteQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryMission chains the current query on the "mission" edge.
func (vq *VoteQuery) QueryMission() *MissionQuery {
	query := &MissionQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vote.Table, vote.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vote.MissionTable, vote.MissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vote entity from the query.
// Returns a *NotFoundError when no Vote was found.
func (vq *VoteQuery) First(ctx context.Context) (*Vote, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VoteQuery) FirstX(ctx context.Context) *Vote {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vote ID from the query.
// Returns a *NotFoundError when no Vote ID was found.
func (vq *VoteQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VoteQuery) FirstIDX(ctx context.Context) int64 {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vote entity is found.
// Returns a *NotFoundError when no Vote entities are found.
func (vq *VoteQuery) Only(ctx context.Context) (*Vote, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vote.Label}
	default:
		return nil, &NotSingularError{vote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VoteQuery) OnlyX(ctx context.Context) *Vote {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vote ID in the query.
// Returns a *NotSingularError when more than one Vote ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VoteQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vote.Label}
	default:
		err = &NotSingularError{vote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VoteQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Votes.
func (vq *VoteQuery) All(ctx context.Context) ([]*Vote, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VoteQuery) AllX(ctx context.Context) []*Vote {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vote IDs.
func (vq *VoteQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := vq.Select(vote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VoteQuery) IDsX(ctx context.Context) []int64 {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VoteQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VoteQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VoteQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VoteQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VoteQuery) Clone() *VoteQuery {
	if vq == nil {
		return nil
	}
	return &VoteQuery{
		config:      vq.config,
		limit:       vq.limit,
		offset:      vq.offset,
		order:       append([]OrderFunc{}, vq.order...),
		predicates:  append([]predicate.Vote{}, vq.predicates...),
		withMission: vq.withMission.Clone(),
		// clone intermediate query.
		sql:    vq.sql.Clone(),
		path:   vq.path,
		unique: vq.unique,
	}
}

// WithMission tells the query-builder to eager-load the nodes that are connected to
// the "mission" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoteQuery) WithMission(opts ...func(*MissionQuery)) *VoteQuery {
	query := &MissionQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withMission = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vote.Query().
//		GroupBy(vote.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VoteQuery) GroupBy(field string, fields ...string) *VoteGroupBy {
	grbuild := &VoteGroupBy{config: vq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	grbuild.label = vote.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.Vote.Query().
//		Select(vote.FieldCreatedBy).
//		Scan(ctx, &v)
func (vq *VoteQuery) Select(fields ...string) *VoteSelect {
	vq.fields = append(vq.fields, fields...)
	selbuild := &VoteSelect{VoteQuery: vq}
	selbuild.label = vote.Label
	selbuild.flds, selbuild.scan = &vq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a VoteSelect configured with the given aggregations.
func (vq *VoteQuery) Aggregate(fns ...AggregateFunc) *VoteSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VoteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !vote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VoteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vote, error) {
	var (
		nodes       = []*Vote{}
		_spec       = vq.querySpec()
		loadedTypes = [1]bool{
			vq.withMission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vote).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vote{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withMission; query != nil {
		if err := vq.loadMission(ctx, query, nodes, nil,
			func(n *Vote, e *Mission) { n.Edges.Mission = e }); err != nil {
			return nil, err
		}
	}
	for i := range vq.loadTotal {
		if err := vq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VoteQuery) loadMission(ctx context.Context, query *MissionQuery, nodes []*Vote, init func(*Vote), assign func(*Vote, *Mission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Vote)
	for i := range nodes {
		fk := nodes[i].MissionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(mission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mission_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vq *VoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.fields
	if len(vq.fields) > 0 {
		_spec.Unique = vq.unique != nil && *vq.unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VoteQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (vq *VoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vote.Table,
			Columns: vote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: vote.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vote.FieldID)
		for i := range fields {
			if fields[i] != vote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vote.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = vote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.unique != nil && *vq.unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VoteGroupBy is the group-by builder for Vote entities.
type VoteGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VoteGroupBy) Aggregate(fns ...AggregateFunc) *VoteGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VoteGroupBy) Scan(ctx context.Context, v any) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

func (vgb *VoteGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range vgb.fields {
		if !vote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VoteGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}

// VoteSelect is the builder for selecting fields of Vote entities.
type VoteSelect struct {
	*VoteQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VoteSelect) Aggregate(fns ...AggregateFunc) *VoteSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VoteSelect) Scan(ctx context.Context, v any) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VoteQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

func (vs *VoteSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(vs.sql))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		vs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		vs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
