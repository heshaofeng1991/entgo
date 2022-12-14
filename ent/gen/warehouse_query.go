// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/warehouse"
)

// WarehouseQuery is the builder for querying Warehouse entities.
type WarehouseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Warehouse
	withOrders *OrderQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WarehouseQuery builder.
func (wq *WarehouseQuery) Where(ps ...predicate.Warehouse) *WarehouseQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WarehouseQuery) Limit(limit int) *WarehouseQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WarehouseQuery) Offset(offset int) *WarehouseQuery {
	wq.offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WarehouseQuery) Unique(unique bool) *WarehouseQuery {
	wq.unique = &unique
	return wq
}

// Order adds an order step to the query.
func (wq *WarehouseQuery) Order(o ...OrderFunc) *WarehouseQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryOrders chains the current query on the "orders" edge.
func (wq *WarehouseQuery) QueryOrders() *OrderQuery {
	query := &OrderQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(warehouse.Table, warehouse.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, warehouse.OrdersTable, warehouse.OrdersColumn),
		)
		schemaConfig := wq.schemaConfig
		step.To.Schema = schemaConfig.Order
		step.Edge.Schema = schemaConfig.Order
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Warehouse entity from the query.
// Returns a *NotFoundError when no Warehouse was found.
func (wq *WarehouseQuery) First(ctx context.Context) (*Warehouse, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{warehouse.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WarehouseQuery) FirstX(ctx context.Context) *Warehouse {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Warehouse ID from the query.
// Returns a *NotFoundError when no Warehouse ID was found.
func (wq *WarehouseQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{warehouse.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WarehouseQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Warehouse entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Warehouse entity is found.
// Returns a *NotFoundError when no Warehouse entities are found.
func (wq *WarehouseQuery) Only(ctx context.Context) (*Warehouse, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{warehouse.Label}
	default:
		return nil, &NotSingularError{warehouse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WarehouseQuery) OnlyX(ctx context.Context) *Warehouse {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Warehouse ID in the query.
// Returns a *NotSingularError when more than one Warehouse ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WarehouseQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{warehouse.Label}
	default:
		err = &NotSingularError{warehouse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WarehouseQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Warehouses.
func (wq *WarehouseQuery) All(ctx context.Context) ([]*Warehouse, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WarehouseQuery) AllX(ctx context.Context) []*Warehouse {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Warehouse IDs.
func (wq *WarehouseQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := wq.Select(warehouse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WarehouseQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WarehouseQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WarehouseQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WarehouseQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WarehouseQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WarehouseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WarehouseQuery) Clone() *WarehouseQuery {
	if wq == nil {
		return nil
	}
	return &WarehouseQuery{
		config:     wq.config,
		limit:      wq.limit,
		offset:     wq.offset,
		order:      append([]OrderFunc{}, wq.order...),
		predicates: append([]predicate.Warehouse{}, wq.predicates...),
		withOrders: wq.withOrders.Clone(),
		// clone intermediate query.
		sql:    wq.sql.Clone(),
		path:   wq.path,
		unique: wq.unique,
	}
}

// WithOrders tells the query-builder to eager-load the nodes that are connected to
// the "orders" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WarehouseQuery) WithOrders(opts ...func(*OrderQuery)) *WarehouseQuery {
	query := &OrderQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withOrders = query
	return wq
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
//	client.Warehouse.Query().
//		GroupBy(warehouse.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (wq *WarehouseQuery) GroupBy(field string, fields ...string) *WarehouseGroupBy {
	grbuild := &WarehouseGroupBy{config: wq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(ctx), nil
	}
	grbuild.label = warehouse.Label
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
//	client.Warehouse.Query().
//		Select(warehouse.FieldCreatedAt).
//		Scan(ctx, &v)
func (wq *WarehouseQuery) Select(fields ...string) *WarehouseSelect {
	wq.fields = append(wq.fields, fields...)
	selbuild := &WarehouseSelect{WarehouseQuery: wq}
	selbuild.label = warehouse.Label
	selbuild.flds, selbuild.scan = &wq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a WarehouseSelect configured with the given aggregations.
func (wq *WarehouseQuery) Aggregate(fns ...AggregateFunc) *WarehouseSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WarehouseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wq.fields {
		if !warehouse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WarehouseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Warehouse, error) {
	var (
		nodes       = []*Warehouse{}
		_spec       = wq.querySpec()
		loadedTypes = [1]bool{
			wq.withOrders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Warehouse).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Warehouse{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = wq.schemaConfig.Warehouse
	ctx = internal.NewSchemaConfigContext(ctx, wq.schemaConfig)
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withOrders; query != nil {
		if err := wq.loadOrders(ctx, query, nodes,
			func(n *Warehouse) { n.Edges.Orders = []*Order{} },
			func(n *Warehouse, e *Order) { n.Edges.Orders = append(n.Edges.Orders, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WarehouseQuery) loadOrders(ctx context.Context, query *OrderQuery, nodes []*Warehouse, init func(*Warehouse), assign func(*Warehouse, *Order)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Warehouse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Order(func(s *sql.Selector) {
		s.Where(sql.InValues(warehouse.OrdersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WarehouseID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "warehouse_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WarehouseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	_spec.Node.Schema = wq.schemaConfig.Warehouse
	ctx = internal.NewSchemaConfigContext(ctx, wq.schemaConfig)
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.fields
	if len(wq.fields) > 0 {
		_spec.Unique = wq.unique != nil && *wq.unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WarehouseQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (wq *WarehouseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   warehouse.Table,
			Columns: warehouse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: warehouse.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if unique := wq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, warehouse.FieldID)
		for i := range fields {
			if fields[i] != warehouse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WarehouseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(warehouse.Table)
	columns := wq.fields
	if len(columns) == 0 {
		columns = warehouse.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.unique != nil && *wq.unique {
		selector.Distinct()
	}
	t1.Schema(wq.schemaConfig.Warehouse)
	ctx = internal.NewSchemaConfigContext(ctx, wq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range wq.modifiers {
		m(selector)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wq *WarehouseQuery) Modify(modifiers ...func(s *sql.Selector)) *WarehouseSelect {
	wq.modifiers = append(wq.modifiers, modifiers...)
	return wq.Select()
}

// WarehouseGroupBy is the group-by builder for Warehouse entities.
type WarehouseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WarehouseGroupBy) Aggregate(fns ...AggregateFunc) *WarehouseGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wgb *WarehouseGroupBy) Scan(ctx context.Context, v any) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

func (wgb *WarehouseGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range wgb.fields {
		if !warehouse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WarehouseGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql.Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
		for _, f := range wgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wgb.fields...)...)
}

// WarehouseSelect is the builder for selecting fields of Warehouse entities.
type WarehouseSelect struct {
	*WarehouseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WarehouseSelect) Aggregate(fns ...AggregateFunc) *WarehouseSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WarehouseSelect) Scan(ctx context.Context, v any) error {
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	ws.sql = ws.WarehouseQuery.sqlQuery(ctx)
	return ws.sqlScan(ctx, v)
}

func (ws *WarehouseSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(ws.sql))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ws.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ws.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ws.sql.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ws *WarehouseSelect) Modify(modifiers ...func(s *sql.Selector)) *WarehouseSelect {
	ws.modifiers = append(ws.modifiers, modifiers...)
	return ws
}
