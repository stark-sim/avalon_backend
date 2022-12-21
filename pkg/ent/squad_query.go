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
	"github.com/stark-sim/avalon_backend/pkg/ent/squad"
)

// SquadQuery is the builder for querying Squad entities.
type SquadQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.Squad
	withMission *MissionQuery
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*Squad) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SquadQuery builder.
func (sq *SquadQuery) Where(ps ...predicate.Squad) *SquadQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SquadQuery) Limit(limit int) *SquadQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SquadQuery) Offset(offset int) *SquadQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SquadQuery) Unique(unique bool) *SquadQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SquadQuery) Order(o ...OrderFunc) *SquadQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryMission chains the current query on the "mission" edge.
func (sq *SquadQuery) QueryMission() *MissionQuery {
	query := &MissionQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(squad.Table, squad.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, squad.MissionTable, squad.MissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Squad entity from the query.
// Returns a *NotFoundError when no Squad was found.
func (sq *SquadQuery) First(ctx context.Context) (*Squad, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{squad.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SquadQuery) FirstX(ctx context.Context) *Squad {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Squad ID from the query.
// Returns a *NotFoundError when no Squad ID was found.
func (sq *SquadQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{squad.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SquadQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Squad entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Squad entity is found.
// Returns a *NotFoundError when no Squad entities are found.
func (sq *SquadQuery) Only(ctx context.Context) (*Squad, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{squad.Label}
	default:
		return nil, &NotSingularError{squad.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SquadQuery) OnlyX(ctx context.Context) *Squad {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Squad ID in the query.
// Returns a *NotSingularError when more than one Squad ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SquadQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{squad.Label}
	default:
		err = &NotSingularError{squad.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SquadQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Squads.
func (sq *SquadQuery) All(ctx context.Context) ([]*Squad, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SquadQuery) AllX(ctx context.Context) []*Squad {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Squad IDs.
func (sq *SquadQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := sq.Select(squad.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SquadQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SquadQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SquadQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SquadQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SquadQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SquadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SquadQuery) Clone() *SquadQuery {
	if sq == nil {
		return nil
	}
	return &SquadQuery{
		config:      sq.config,
		limit:       sq.limit,
		offset:      sq.offset,
		order:       append([]OrderFunc{}, sq.order...),
		predicates:  append([]predicate.Squad{}, sq.predicates...),
		withMission: sq.withMission.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithMission tells the query-builder to eager-load the nodes that are connected to
// the "mission" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SquadQuery) WithMission(opts ...func(*MissionQuery)) *SquadQuery {
	query := &MissionQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withMission = query
	return sq
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
//	client.Squad.Query().
//		GroupBy(squad.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SquadQuery) GroupBy(field string, fields ...string) *SquadGroupBy {
	grbuild := &SquadGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = squad.Label
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
//	client.Squad.Query().
//		Select(squad.FieldCreatedBy).
//		Scan(ctx, &v)
func (sq *SquadQuery) Select(fields ...string) *SquadSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SquadSelect{SquadQuery: sq}
	selbuild.label = squad.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SquadSelect configured with the given aggregations.
func (sq *SquadQuery) Aggregate(fns ...AggregateFunc) *SquadSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SquadQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !squad.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SquadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Squad, error) {
	var (
		nodes       = []*Squad{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withMission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Squad).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Squad{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withMission; query != nil {
		if err := sq.loadMission(ctx, query, nodes, nil,
			func(n *Squad, e *Mission) { n.Edges.Mission = e }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SquadQuery) loadMission(ctx context.Context, query *MissionQuery, nodes []*Squad, init func(*Squad), assign func(*Squad, *Mission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Squad)
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

func (sq *SquadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SquadQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SquadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   squad.Table,
			Columns: squad.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: squad.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, squad.FieldID)
		for i := range fields {
			if fields[i] != squad.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SquadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(squad.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = squad.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SquadGroupBy is the group-by builder for Squad entities.
type SquadGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SquadGroupBy) Aggregate(fns ...AggregateFunc) *SquadGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SquadGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SquadGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !squad.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SquadGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SquadSelect is the builder for selecting fields of Squad entities.
type SquadSelect struct {
	*SquadQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SquadSelect) Aggregate(fns ...AggregateFunc) *SquadSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SquadSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SquadQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SquadSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(ss.sql))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ss.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ss.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}