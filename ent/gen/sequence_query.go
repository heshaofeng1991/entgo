// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/sequence"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
)

// SequenceQuery is the builder for querying Sequence entities.
type SequenceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Sequence
	withTenant *TenantQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SequenceQuery builder.
func (sq *SequenceQuery) Where(ps ...predicate.Sequence) *SequenceQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SequenceQuery) Limit(limit int) *SequenceQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SequenceQuery) Offset(offset int) *SequenceQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SequenceQuery) Unique(unique bool) *SequenceQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SequenceQuery) Order(o ...OrderFunc) *SequenceQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryTenant chains the current query on the "tenant" edge.
func (sq *SequenceQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sequence.Table, sequence.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, sequence.TenantTable, sequence.TenantColumn),
		)
		schemaConfig := sq.schemaConfig
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.Sequence
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Sequence entity from the query.
// Returns a *NotFoundError when no Sequence was found.
func (sq *SequenceQuery) First(ctx context.Context) (*Sequence, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sequence.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SequenceQuery) FirstX(ctx context.Context) *Sequence {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Sequence ID from the query.
// Returns a *NotFoundError when no Sequence ID was found.
func (sq *SequenceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sequence.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SequenceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Sequence entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Sequence entity is found.
// Returns a *NotFoundError when no Sequence entities are found.
func (sq *SequenceQuery) Only(ctx context.Context) (*Sequence, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sequence.Label}
	default:
		return nil, &NotSingularError{sequence.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SequenceQuery) OnlyX(ctx context.Context) *Sequence {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Sequence ID in the query.
// Returns a *NotSingularError when more than one Sequence ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SequenceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sequence.Label}
	default:
		err = &NotSingularError{sequence.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SequenceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Sequences.
func (sq *SequenceQuery) All(ctx context.Context) ([]*Sequence, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SequenceQuery) AllX(ctx context.Context) []*Sequence {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Sequence IDs.
func (sq *SequenceQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := sq.Select(sequence.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SequenceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SequenceQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SequenceQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SequenceQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SequenceQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SequenceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SequenceQuery) Clone() *SequenceQuery {
	if sq == nil {
		return nil
	}
	return &SequenceQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		predicates: append([]predicate.Sequence{}, sq.predicates...),
		withTenant: sq.withTenant.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SequenceQuery) WithTenant(opts ...func(*TenantQuery)) *SequenceQuery {
	query := &TenantQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withTenant = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Sequence.Query().
//		GroupBy(sequence.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (sq *SequenceQuery) GroupBy(field string, fields ...string) *SequenceGroupBy {
	grbuild := &SequenceGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = sequence.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Sequence.Query().
//		Select(sequence.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *SequenceQuery) Select(fields ...string) *SequenceSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SequenceSelect{SequenceQuery: sq}
	selbuild.label = sequence.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SequenceSelect configured with the given aggregations.
func (sq *SequenceQuery) Aggregate(fns ...AggregateFunc) *SequenceSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SequenceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !sequence.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	if sequence.Policy == nil {
		return errors.New("gen: uninitialized sequence.Policy (forgotten import gen/runtime?)")
	}
	if err := sequence.Policy.EvalQuery(ctx, sq); err != nil {
		return err
	}
	return nil
}

func (sq *SequenceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Sequence, error) {
	var (
		nodes       = []*Sequence{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withTenant != nil,
		}
	)
	if sq.withTenant != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, sequence.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Sequence).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Sequence{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = sq.schemaConfig.Sequence
	ctx = internal.NewSchemaConfigContext(ctx, sq.schemaConfig)
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
	if query := sq.withTenant; query != nil {
		if err := sq.loadTenant(ctx, query, nodes, nil,
			func(n *Sequence, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SequenceQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*Sequence, init func(*Sequence), assign func(*Sequence, *Tenant)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Sequence)
	for i := range nodes {
		if nodes[i].sequence_tenant == nil {
			continue
		}
		fk := *nodes[i].sequence_tenant
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tenant.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sequence_tenant" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SequenceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Schema = sq.schemaConfig.Sequence
	ctx = internal.NewSchemaConfigContext(ctx, sq.schemaConfig)
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SequenceQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SequenceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sequence.Table,
			Columns: sequence.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sequence.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, sequence.FieldID)
		for i := range fields {
			if fields[i] != sequence.FieldID {
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

func (sq *SequenceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(sequence.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = sequence.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	t1.Schema(sq.schemaConfig.Sequence)
	ctx = internal.NewSchemaConfigContext(ctx, sq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range sq.modifiers {
		m(selector)
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

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *SequenceQuery) Modify(modifiers ...func(s *sql.Selector)) *SequenceSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// SequenceGroupBy is the group-by builder for Sequence entities.
type SequenceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SequenceGroupBy) Aggregate(fns ...AggregateFunc) *SequenceGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SequenceGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SequenceGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !sequence.ValidColumn(f) {
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

func (sgb *SequenceGroupBy) sqlQuery() *sql.Selector {
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

// SequenceSelect is the builder for selecting fields of Sequence entities.
type SequenceSelect struct {
	*SequenceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SequenceSelect) Aggregate(fns ...AggregateFunc) *SequenceSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SequenceSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SequenceQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SequenceSelect) sqlScan(ctx context.Context, v any) error {
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

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *SequenceSelect) Modify(modifiers ...func(s *sql.Selector)) *SequenceSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}
