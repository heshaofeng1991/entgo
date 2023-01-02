// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/announcements"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/user"
)

// AnnouncementsQuery is the builder for querying Announcements entities.
type AnnouncementsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Announcements
	withUsers  *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnnouncementsQuery builder.
func (aq *AnnouncementsQuery) Where(ps ...predicate.Announcements) *AnnouncementsQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AnnouncementsQuery) Limit(limit int) *AnnouncementsQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AnnouncementsQuery) Offset(offset int) *AnnouncementsQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AnnouncementsQuery) Unique(unique bool) *AnnouncementsQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AnnouncementsQuery) Order(o ...OrderFunc) *AnnouncementsQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUsers chains the current query on the "users" edge.
func (aq *AnnouncementsQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(announcements.Table, announcements.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, announcements.UsersTable, announcements.UsersColumn),
		)
		schemaConfig := aq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Announcements
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Announcements entity from the query.
// Returns a *NotFoundError when no Announcements was found.
func (aq *AnnouncementsQuery) First(ctx context.Context) (*Announcements, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{announcements.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AnnouncementsQuery) FirstX(ctx context.Context) *Announcements {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Announcements ID from the query.
// Returns a *NotFoundError when no Announcements ID was found.
func (aq *AnnouncementsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{announcements.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AnnouncementsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Announcements entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Announcements entity is found.
// Returns a *NotFoundError when no Announcements entities are found.
func (aq *AnnouncementsQuery) Only(ctx context.Context) (*Announcements, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{announcements.Label}
	default:
		return nil, &NotSingularError{announcements.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AnnouncementsQuery) OnlyX(ctx context.Context) *Announcements {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Announcements ID in the query.
// Returns a *NotSingularError when more than one Announcements ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AnnouncementsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{announcements.Label}
	default:
		err = &NotSingularError{announcements.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AnnouncementsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AnnouncementsSlice.
func (aq *AnnouncementsQuery) All(ctx context.Context) ([]*Announcements, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AnnouncementsQuery) AllX(ctx context.Context) []*Announcements {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Announcements IDs.
func (aq *AnnouncementsQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := aq.Select(announcements.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AnnouncementsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AnnouncementsQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AnnouncementsQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AnnouncementsQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AnnouncementsQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnnouncementsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AnnouncementsQuery) Clone() *AnnouncementsQuery {
	if aq == nil {
		return nil
	}
	return &AnnouncementsQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		predicates: append([]predicate.Announcements{}, aq.predicates...),
		withUsers:  aq.withUsers.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnnouncementsQuery) WithUsers(opts ...func(*UserQuery)) *AnnouncementsQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withUsers = query
	return aq
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
//	client.Announcements.Query().
//		GroupBy(announcements.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (aq *AnnouncementsQuery) GroupBy(field string, fields ...string) *AnnouncementsGroupBy {
	grbuild := &AnnouncementsGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = announcements.Label
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
//	client.Announcements.Query().
//		Select(announcements.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AnnouncementsQuery) Select(fields ...string) *AnnouncementsSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AnnouncementsSelect{AnnouncementsQuery: aq}
	selbuild.label = announcements.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AnnouncementsSelect configured with the given aggregations.
func (aq *AnnouncementsQuery) Aggregate(fns ...AggregateFunc) *AnnouncementsSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AnnouncementsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !announcements.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AnnouncementsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Announcements, error) {
	var (
		nodes       = []*Announcements{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Announcements).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Announcements{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = aq.schemaConfig.Announcements
	ctx = internal.NewSchemaConfigContext(ctx, aq.schemaConfig)
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUsers; query != nil {
		if err := aq.loadUsers(ctx, query, nodes, nil,
			func(n *Announcements, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AnnouncementsQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Announcements, init func(*Announcements), assign func(*Announcements, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Announcements)
	for i := range nodes {
		fk := nodes[i].CreateBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "create_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AnnouncementsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Schema = aq.schemaConfig.Announcements
	ctx = internal.NewSchemaConfigContext(ctx, aq.schemaConfig)
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AnnouncementsQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AnnouncementsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   announcements.Table,
			Columns: announcements.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: announcements.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, announcements.FieldID)
		for i := range fields {
			if fields[i] != announcements.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AnnouncementsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(announcements.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = announcements.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	t1.Schema(aq.schemaConfig.Announcements)
	ctx = internal.NewSchemaConfigContext(ctx, aq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *AnnouncementsQuery) Modify(modifiers ...func(s *sql.Selector)) *AnnouncementsSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// AnnouncementsGroupBy is the group-by builder for Announcements entities.
type AnnouncementsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AnnouncementsGroupBy) Aggregate(fns ...AggregateFunc) *AnnouncementsGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AnnouncementsGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AnnouncementsGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !announcements.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AnnouncementsGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AnnouncementsSelect is the builder for selecting fields of Announcements entities.
type AnnouncementsSelect struct {
	*AnnouncementsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AnnouncementsSelect) Aggregate(fns ...AggregateFunc) *AnnouncementsSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AnnouncementsSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AnnouncementsQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AnnouncementsSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(as.sql))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		as.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		as.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *AnnouncementsSelect) Modify(modifiers ...func(s *sql.Selector)) *AnnouncementsSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}