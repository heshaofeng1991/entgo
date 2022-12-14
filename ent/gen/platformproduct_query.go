// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/platformproduct"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/productmapping"
	"github.com/heshaofeng1991/entgo/ent/gen/store"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
)

// PlatformProductQuery is the builder for querying PlatformProduct entities.
type PlatformProductQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.PlatformProduct
	withTenant          *TenantQuery
	withProductMappings *ProductMappingQuery
	withStores          *StoreQuery
	withFKs             bool
	modifiers           []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlatformProductQuery builder.
func (ppq *PlatformProductQuery) Where(ps ...predicate.PlatformProduct) *PlatformProductQuery {
	ppq.predicates = append(ppq.predicates, ps...)
	return ppq
}

// Limit adds a limit step to the query.
func (ppq *PlatformProductQuery) Limit(limit int) *PlatformProductQuery {
	ppq.limit = &limit
	return ppq
}

// Offset adds an offset step to the query.
func (ppq *PlatformProductQuery) Offset(offset int) *PlatformProductQuery {
	ppq.offset = &offset
	return ppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ppq *PlatformProductQuery) Unique(unique bool) *PlatformProductQuery {
	ppq.unique = &unique
	return ppq
}

// Order adds an order step to the query.
func (ppq *PlatformProductQuery) Order(o ...OrderFunc) *PlatformProductQuery {
	ppq.order = append(ppq.order, o...)
	return ppq
}

// QueryTenant chains the current query on the "tenant" edge.
func (ppq *PlatformProductQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: ppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(platformproduct.Table, platformproduct.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, platformproduct.TenantTable, platformproduct.TenantColumn),
		)
		schemaConfig := ppq.schemaConfig
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.PlatformProduct
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductMappings chains the current query on the "product_mappings" edge.
func (ppq *PlatformProductQuery) QueryProductMappings() *ProductMappingQuery {
	query := &ProductMappingQuery{config: ppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(platformproduct.Table, platformproduct.FieldID, selector),
			sqlgraph.To(productmapping.Table, productmapping.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, platformproduct.ProductMappingsTable, platformproduct.ProductMappingsColumn),
		)
		schemaConfig := ppq.schemaConfig
		step.To.Schema = schemaConfig.ProductMapping
		step.Edge.Schema = schemaConfig.ProductMapping
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStores chains the current query on the "stores" edge.
func (ppq *PlatformProductQuery) QueryStores() *StoreQuery {
	query := &StoreQuery{config: ppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(platformproduct.Table, platformproduct.FieldID, selector),
			sqlgraph.To(store.Table, store.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, platformproduct.StoresTable, platformproduct.StoresColumn),
		)
		schemaConfig := ppq.schemaConfig
		step.To.Schema = schemaConfig.Store
		step.Edge.Schema = schemaConfig.PlatformProduct
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlatformProduct entity from the query.
// Returns a *NotFoundError when no PlatformProduct was found.
func (ppq *PlatformProductQuery) First(ctx context.Context) (*PlatformProduct, error) {
	nodes, err := ppq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{platformproduct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ppq *PlatformProductQuery) FirstX(ctx context.Context) *PlatformProduct {
	node, err := ppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlatformProduct ID from the query.
// Returns a *NotFoundError when no PlatformProduct ID was found.
func (ppq *PlatformProductQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ppq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{platformproduct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ppq *PlatformProductQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlatformProduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlatformProduct entity is found.
// Returns a *NotFoundError when no PlatformProduct entities are found.
func (ppq *PlatformProductQuery) Only(ctx context.Context) (*PlatformProduct, error) {
	nodes, err := ppq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{platformproduct.Label}
	default:
		return nil, &NotSingularError{platformproduct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ppq *PlatformProductQuery) OnlyX(ctx context.Context) *PlatformProduct {
	node, err := ppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlatformProduct ID in the query.
// Returns a *NotSingularError when more than one PlatformProduct ID is found.
// Returns a *NotFoundError when no entities are found.
func (ppq *PlatformProductQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ppq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{platformproduct.Label}
	default:
		err = &NotSingularError{platformproduct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ppq *PlatformProductQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlatformProducts.
func (ppq *PlatformProductQuery) All(ctx context.Context) ([]*PlatformProduct, error) {
	if err := ppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ppq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ppq *PlatformProductQuery) AllX(ctx context.Context) []*PlatformProduct {
	nodes, err := ppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlatformProduct IDs.
func (ppq *PlatformProductQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := ppq.Select(platformproduct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ppq *PlatformProductQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ppq *PlatformProductQuery) Count(ctx context.Context) (int, error) {
	if err := ppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ppq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ppq *PlatformProductQuery) CountX(ctx context.Context) int {
	count, err := ppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ppq *PlatformProductQuery) Exist(ctx context.Context) (bool, error) {
	if err := ppq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ppq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ppq *PlatformProductQuery) ExistX(ctx context.Context) bool {
	exist, err := ppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlatformProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ppq *PlatformProductQuery) Clone() *PlatformProductQuery {
	if ppq == nil {
		return nil
	}
	return &PlatformProductQuery{
		config:              ppq.config,
		limit:               ppq.limit,
		offset:              ppq.offset,
		order:               append([]OrderFunc{}, ppq.order...),
		predicates:          append([]predicate.PlatformProduct{}, ppq.predicates...),
		withTenant:          ppq.withTenant.Clone(),
		withProductMappings: ppq.withProductMappings.Clone(),
		withStores:          ppq.withStores.Clone(),
		// clone intermediate query.
		sql:    ppq.sql.Clone(),
		path:   ppq.path,
		unique: ppq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PlatformProductQuery) WithTenant(opts ...func(*TenantQuery)) *PlatformProductQuery {
	query := &TenantQuery{config: ppq.config}
	for _, opt := range opts {
		opt(query)
	}
	ppq.withTenant = query
	return ppq
}

// WithProductMappings tells the query-builder to eager-load the nodes that are connected to
// the "product_mappings" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PlatformProductQuery) WithProductMappings(opts ...func(*ProductMappingQuery)) *PlatformProductQuery {
	query := &ProductMappingQuery{config: ppq.config}
	for _, opt := range opts {
		opt(query)
	}
	ppq.withProductMappings = query
	return ppq
}

// WithStores tells the query-builder to eager-load the nodes that are connected to
// the "stores" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PlatformProductQuery) WithStores(opts ...func(*StoreQuery)) *PlatformProductQuery {
	query := &StoreQuery{config: ppq.config}
	for _, opt := range opts {
		opt(query)
	}
	ppq.withStores = query
	return ppq
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
//	client.PlatformProduct.Query().
//		GroupBy(platformproduct.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (ppq *PlatformProductQuery) GroupBy(field string, fields ...string) *PlatformProductGroupBy {
	grbuild := &PlatformProductGroupBy{config: ppq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ppq.sqlQuery(ctx), nil
	}
	grbuild.label = platformproduct.Label
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
//	client.PlatformProduct.Query().
//		Select(platformproduct.FieldCreatedAt).
//		Scan(ctx, &v)
func (ppq *PlatformProductQuery) Select(fields ...string) *PlatformProductSelect {
	ppq.fields = append(ppq.fields, fields...)
	selbuild := &PlatformProductSelect{PlatformProductQuery: ppq}
	selbuild.label = platformproduct.Label
	selbuild.flds, selbuild.scan = &ppq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a PlatformProductSelect configured with the given aggregations.
func (ppq *PlatformProductQuery) Aggregate(fns ...AggregateFunc) *PlatformProductSelect {
	return ppq.Select().Aggregate(fns...)
}

func (ppq *PlatformProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ppq.fields {
		if !platformproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if ppq.path != nil {
		prev, err := ppq.path(ctx)
		if err != nil {
			return err
		}
		ppq.sql = prev
	}
	if platformproduct.Policy == nil {
		return errors.New("gen: uninitialized platformproduct.Policy (forgotten import gen/runtime?)")
	}
	if err := platformproduct.Policy.EvalQuery(ctx, ppq); err != nil {
		return err
	}
	return nil
}

func (ppq *PlatformProductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlatformProduct, error) {
	var (
		nodes       = []*PlatformProduct{}
		withFKs     = ppq.withFKs
		_spec       = ppq.querySpec()
		loadedTypes = [3]bool{
			ppq.withTenant != nil,
			ppq.withProductMappings != nil,
			ppq.withStores != nil,
		}
	)
	if ppq.withTenant != nil || ppq.withStores != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, platformproduct.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlatformProduct).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlatformProduct{config: ppq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = ppq.schemaConfig.PlatformProduct
	ctx = internal.NewSchemaConfigContext(ctx, ppq.schemaConfig)
	if len(ppq.modifiers) > 0 {
		_spec.Modifiers = ppq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ppq.withTenant; query != nil {
		if err := ppq.loadTenant(ctx, query, nodes, nil,
			func(n *PlatformProduct, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := ppq.withProductMappings; query != nil {
		if err := ppq.loadProductMappings(ctx, query, nodes,
			func(n *PlatformProduct) { n.Edges.ProductMappings = []*ProductMapping{} },
			func(n *PlatformProduct, e *ProductMapping) {
				n.Edges.ProductMappings = append(n.Edges.ProductMappings, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := ppq.withStores; query != nil {
		if err := ppq.loadStores(ctx, query, nodes, nil,
			func(n *PlatformProduct, e *Store) { n.Edges.Stores = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ppq *PlatformProductQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*PlatformProduct, init func(*PlatformProduct), assign func(*PlatformProduct, *Tenant)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*PlatformProduct)
	for i := range nodes {
		if nodes[i].platform_product_tenant == nil {
			continue
		}
		fk := *nodes[i].platform_product_tenant
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
			return fmt.Errorf(`unexpected foreign-key "platform_product_tenant" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ppq *PlatformProductQuery) loadProductMappings(ctx context.Context, query *ProductMappingQuery, nodes []*PlatformProduct, init func(*PlatformProduct), assign func(*PlatformProduct, *ProductMapping)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*PlatformProduct)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.ProductMapping(func(s *sql.Selector) {
		s.Where(sql.InValues(platformproduct.ProductMappingsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PlatformProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "platform_product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ppq *PlatformProductQuery) loadStores(ctx context.Context, query *StoreQuery, nodes []*PlatformProduct, init func(*PlatformProduct), assign func(*PlatformProduct, *Store)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*PlatformProduct)
	for i := range nodes {
		fk := nodes[i].StoreID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(store.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "store_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ppq *PlatformProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ppq.querySpec()
	_spec.Node.Schema = ppq.schemaConfig.PlatformProduct
	ctx = internal.NewSchemaConfigContext(ctx, ppq.schemaConfig)
	if len(ppq.modifiers) > 0 {
		_spec.Modifiers = ppq.modifiers
	}
	_spec.Node.Columns = ppq.fields
	if len(ppq.fields) > 0 {
		_spec.Unique = ppq.unique != nil && *ppq.unique
	}
	return sqlgraph.CountNodes(ctx, ppq.driver, _spec)
}

func (ppq *PlatformProductQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ppq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ppq *PlatformProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   platformproduct.Table,
			Columns: platformproduct.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: platformproduct.FieldID,
			},
		},
		From:   ppq.sql,
		Unique: true,
	}
	if unique := ppq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ppq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platformproduct.FieldID)
		for i := range fields {
			if fields[i] != platformproduct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ppq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ppq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ppq *PlatformProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ppq.driver.Dialect())
	t1 := builder.Table(platformproduct.Table)
	columns := ppq.fields
	if len(columns) == 0 {
		columns = platformproduct.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ppq.sql != nil {
		selector = ppq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ppq.unique != nil && *ppq.unique {
		selector.Distinct()
	}
	t1.Schema(ppq.schemaConfig.PlatformProduct)
	ctx = internal.NewSchemaConfigContext(ctx, ppq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range ppq.modifiers {
		m(selector)
	}
	for _, p := range ppq.predicates {
		p(selector)
	}
	for _, p := range ppq.order {
		p(selector)
	}
	if offset := ppq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ppq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ppq *PlatformProductQuery) Modify(modifiers ...func(s *sql.Selector)) *PlatformProductSelect {
	ppq.modifiers = append(ppq.modifiers, modifiers...)
	return ppq.Select()
}

// PlatformProductGroupBy is the group-by builder for PlatformProduct entities.
type PlatformProductGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ppgb *PlatformProductGroupBy) Aggregate(fns ...AggregateFunc) *PlatformProductGroupBy {
	ppgb.fns = append(ppgb.fns, fns...)
	return ppgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ppgb *PlatformProductGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ppgb.path(ctx)
	if err != nil {
		return err
	}
	ppgb.sql = query
	return ppgb.sqlScan(ctx, v)
}

func (ppgb *PlatformProductGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ppgb.fields {
		if !platformproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ppgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ppgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ppgb *PlatformProductGroupBy) sqlQuery() *sql.Selector {
	selector := ppgb.sql.Select()
	aggregation := make([]string, 0, len(ppgb.fns))
	for _, fn := range ppgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ppgb.fields)+len(ppgb.fns))
		for _, f := range ppgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ppgb.fields...)...)
}

// PlatformProductSelect is the builder for selecting fields of PlatformProduct entities.
type PlatformProductSelect struct {
	*PlatformProductQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pps *PlatformProductSelect) Aggregate(fns ...AggregateFunc) *PlatformProductSelect {
	pps.fns = append(pps.fns, fns...)
	return pps
}

// Scan applies the selector query and scans the result into the given value.
func (pps *PlatformProductSelect) Scan(ctx context.Context, v any) error {
	if err := pps.prepareQuery(ctx); err != nil {
		return err
	}
	pps.sql = pps.PlatformProductQuery.sqlQuery(ctx)
	return pps.sqlScan(ctx, v)
}

func (pps *PlatformProductSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(pps.fns))
	for _, fn := range pps.fns {
		aggregation = append(aggregation, fn(pps.sql))
	}
	switch n := len(*pps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		pps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		pps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := pps.sql.Query()
	if err := pps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pps *PlatformProductSelect) Modify(modifiers ...func(s *sql.Selector)) *PlatformProductSelect {
	pps.modifiers = append(pps.modifiers, modifiers...)
	return pps
}
