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
	"github.com/heshaofeng1991/entgo/ent/gen/inventory"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/product"
	"github.com/heshaofeng1991/entgo/ent/gen/productmapping"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
)

// ProductQuery is the builder for querying Product entities.
type ProductQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.Product
	withTenant          *TenantQuery
	withInventories     *InventoryQuery
	withProductMappings *ProductMappingQuery
	withFKs             bool
	modifiers           []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductQuery builder.
func (pq *ProductQuery) Where(ps ...predicate.Product) *ProductQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *ProductQuery) Limit(limit int) *ProductQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *ProductQuery) Offset(offset int) *ProductQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProductQuery) Unique(unique bool) *ProductQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *ProductQuery) Order(o ...OrderFunc) *ProductQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryTenant chains the current query on the "tenant" edge.
func (pq *ProductQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, product.TenantTable, product.TenantColumn),
		)
		schemaConfig := pq.schemaConfig
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.Product
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInventories chains the current query on the "inventories" edge.
func (pq *ProductQuery) QueryInventories() *InventoryQuery {
	query := &InventoryQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(inventory.Table, inventory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.InventoriesTable, product.InventoriesColumn),
		)
		schemaConfig := pq.schemaConfig
		step.To.Schema = schemaConfig.Inventory
		step.Edge.Schema = schemaConfig.Inventory
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductMappings chains the current query on the "product_mappings" edge.
func (pq *ProductQuery) QueryProductMappings() *ProductMappingQuery {
	query := &ProductMappingQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(productmapping.Table, productmapping.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.ProductMappingsTable, product.ProductMappingsColumn),
		)
		schemaConfig := pq.schemaConfig
		step.To.Schema = schemaConfig.ProductMapping
		step.Edge.Schema = schemaConfig.ProductMapping
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Product entity from the query.
// Returns a *NotFoundError when no Product was found.
func (pq *ProductQuery) First(ctx context.Context) (*Product, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{product.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProductQuery) FirstX(ctx context.Context) *Product {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Product ID from the query.
// Returns a *NotFoundError when no Product ID was found.
func (pq *ProductQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{product.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProductQuery) FirstIDX(ctx context.Context) int64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Product entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Product entity is found.
// Returns a *NotFoundError when no Product entities are found.
func (pq *ProductQuery) Only(ctx context.Context) (*Product, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{product.Label}
	default:
		return nil, &NotSingularError{product.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProductQuery) OnlyX(ctx context.Context) *Product {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Product ID in the query.
// Returns a *NotSingularError when more than one Product ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProductQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = &NotSingularError{product.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProductQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Products.
func (pq *ProductQuery) All(ctx context.Context) ([]*Product, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProductQuery) AllX(ctx context.Context) []*Product {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Product IDs.
func (pq *ProductQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := pq.Select(product.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProductQuery) IDsX(ctx context.Context) []int64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProductQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProductQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProductQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProductQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProductQuery) Clone() *ProductQuery {
	if pq == nil {
		return nil
	}
	return &ProductQuery{
		config:              pq.config,
		limit:               pq.limit,
		offset:              pq.offset,
		order:               append([]OrderFunc{}, pq.order...),
		predicates:          append([]predicate.Product{}, pq.predicates...),
		withTenant:          pq.withTenant.Clone(),
		withInventories:     pq.withInventories.Clone(),
		withProductMappings: pq.withProductMappings.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithTenant(opts ...func(*TenantQuery)) *ProductQuery {
	query := &TenantQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withTenant = query
	return pq
}

// WithInventories tells the query-builder to eager-load the nodes that are connected to
// the "inventories" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithInventories(opts ...func(*InventoryQuery)) *ProductQuery {
	query := &InventoryQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withInventories = query
	return pq
}

// WithProductMappings tells the query-builder to eager-load the nodes that are connected to
// the "product_mappings" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithProductMappings(opts ...func(*ProductMappingQuery)) *ProductQuery {
	query := &ProductMappingQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withProductMappings = query
	return pq
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
//	client.Product.Query().
//		GroupBy(product.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (pq *ProductQuery) GroupBy(field string, fields ...string) *ProductGroupBy {
	grbuild := &ProductGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = product.Label
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
//	client.Product.Query().
//		Select(product.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *ProductQuery) Select(fields ...string) *ProductSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &ProductSelect{ProductQuery: pq}
	selbuild.label = product.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ProductSelect configured with the given aggregations.
func (pq *ProductQuery) Aggregate(fns ...AggregateFunc) *ProductSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	if product.Policy == nil {
		return errors.New("gen: uninitialized product.Policy (forgotten import gen/runtime?)")
	}
	if err := product.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *ProductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Product, error) {
	var (
		nodes       = []*Product{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [3]bool{
			pq.withTenant != nil,
			pq.withInventories != nil,
			pq.withProductMappings != nil,
		}
	)
	if pq.withTenant != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, product.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Product).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Product{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = pq.schemaConfig.Product
	ctx = internal.NewSchemaConfigContext(ctx, pq.schemaConfig)
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withTenant; query != nil {
		if err := pq.loadTenant(ctx, query, nodes, nil,
			func(n *Product, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withInventories; query != nil {
		if err := pq.loadInventories(ctx, query, nodes,
			func(n *Product) { n.Edges.Inventories = []*Inventory{} },
			func(n *Product, e *Inventory) { n.Edges.Inventories = append(n.Edges.Inventories, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withProductMappings; query != nil {
		if err := pq.loadProductMappings(ctx, query, nodes,
			func(n *Product) { n.Edges.ProductMappings = []*ProductMapping{} },
			func(n *Product, e *ProductMapping) { n.Edges.ProductMappings = append(n.Edges.ProductMappings, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProductQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*Product, init func(*Product), assign func(*Product, *Tenant)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Product)
	for i := range nodes {
		if nodes[i].product_tenant == nil {
			continue
		}
		fk := *nodes[i].product_tenant
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
			return fmt.Errorf(`unexpected foreign-key "product_tenant" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProductQuery) loadInventories(ctx context.Context, query *InventoryQuery, nodes []*Product, init func(*Product), assign func(*Product, *Inventory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Product)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Inventory(func(s *sql.Selector) {
		s.Where(sql.InValues(product.InventoriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProductQuery) loadProductMappings(ctx context.Context, query *ProductMappingQuery, nodes []*Product, init func(*Product), assign func(*Product, *ProductMapping)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Product)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.ProductMapping(func(s *sql.Selector) {
		s.Where(sql.InValues(product.ProductMappingsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *ProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Schema = pq.schemaConfig.Product
	ctx = internal.NewSchemaConfigContext(ctx, pq.schemaConfig)
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProductQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (pq *ProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: product.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for i := range fields {
			if fields[i] != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(product.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = product.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	t1.Schema(pq.schemaConfig.Product)
	ctx = internal.NewSchemaConfigContext(ctx, pq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *ProductQuery) Modify(modifiers ...func(s *sql.Selector)) *ProductSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// ProductGroupBy is the group-by builder for Product entities.
type ProductGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProductGroupBy) Aggregate(fns ...AggregateFunc) *ProductGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *ProductGroupBy) Scan(ctx context.Context, v any) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *ProductGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range pgb.fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *ProductGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// ProductSelect is the builder for selecting fields of Product entities.
type ProductSelect struct {
	*ProductQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProductSelect) Aggregate(fns ...AggregateFunc) *ProductSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProductSelect) Scan(ctx context.Context, v any) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.ProductQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *ProductSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(ps.sql))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *ProductSelect) Modify(modifiers ...func(s *sql.Selector)) *ProductSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
