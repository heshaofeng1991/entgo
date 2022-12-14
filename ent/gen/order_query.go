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
	"github.com/heshaofeng1991/entgo/ent/gen/channel"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/orderholdreason"
	"github.com/heshaofeng1991/entgo/ent/gen/orderitem"
	"github.com/heshaofeng1991/entgo/ent/gen/ordertaxation"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/store"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
	"github.com/heshaofeng1991/entgo/ent/gen/trackmapping"
	"github.com/heshaofeng1991/entgo/ent/gen/warehouse"
)

// OrderQuery is the builder for querying Order entities.
type OrderQuery struct {
	config
	limit                *int
	offset               *int
	unique               *bool
	order                []OrderFunc
	fields               []string
	predicates           []predicate.Order
	withTenant           *TenantQuery
	withOrderItems       *OrderItemQuery
	withOrderHoldReasons *OrderHoldReasonQuery
	withOrderTaxations   *OrderTaxationQuery
	withStores           *StoreQuery
	withWarehouses       *WarehouseQuery
	withChannels         *ChannelQuery
	withTrackMappings    *TrackMappingQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderQuery builder.
func (oq *OrderQuery) Where(ps ...predicate.Order) *OrderQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OrderQuery) Limit(limit int) *OrderQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OrderQuery) Offset(offset int) *OrderQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OrderQuery) Unique(unique bool) *OrderQuery {
	oq.unique = &unique
	return oq
}

// Order adds an order step to the query.
func (oq *OrderQuery) Order(o ...OrderFunc) *OrderQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryTenant chains the current query on the "tenant" edge.
func (oq *OrderQuery) QueryTenant() *TenantQuery {
	query := &TenantQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, order.TenantTable, order.TenantColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.Order
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItems chains the current query on the "order_items" edge.
func (oq *OrderQuery) QueryOrderItems() *OrderItemQuery {
	query := &OrderItemQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.OrderItemsTable, order.OrderItemsColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.OrderItem
		step.Edge.Schema = schemaConfig.OrderItem
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderHoldReasons chains the current query on the "order_hold_reasons" edge.
func (oq *OrderQuery) QueryOrderHoldReasons() *OrderHoldReasonQuery {
	query := &OrderHoldReasonQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(orderholdreason.Table, orderholdreason.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.OrderHoldReasonsTable, order.OrderHoldReasonsColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.OrderHoldReason
		step.Edge.Schema = schemaConfig.OrderHoldReason
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderTaxations chains the current query on the "order_taxations" edge.
func (oq *OrderQuery) QueryOrderTaxations() *OrderTaxationQuery {
	query := &OrderTaxationQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(ordertaxation.Table, ordertaxation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.OrderTaxationsTable, order.OrderTaxationsColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.OrderTaxation
		step.Edge.Schema = schemaConfig.OrderTaxation
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStores chains the current query on the "stores" edge.
func (oq *OrderQuery) QueryStores() *StoreQuery {
	query := &StoreQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(store.Table, store.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.StoresTable, order.StoresColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.Store
		step.Edge.Schema = schemaConfig.Order
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWarehouses chains the current query on the "warehouses" edge.
func (oq *OrderQuery) QueryWarehouses() *WarehouseQuery {
	query := &WarehouseQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(warehouse.Table, warehouse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.WarehousesTable, order.WarehousesColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.Warehouse
		step.Edge.Schema = schemaConfig.Order
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChannels chains the current query on the "channels" edge.
func (oq *OrderQuery) QueryChannels() *ChannelQuery {
	query := &ChannelQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.ChannelsTable, order.ChannelsColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.Channel
		step.Edge.Schema = schemaConfig.Order
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTrackMappings chains the current query on the "track_mappings" edge.
func (oq *OrderQuery) QueryTrackMappings() *TrackMappingQuery {
	query := &TrackMappingQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, selector),
			sqlgraph.To(trackmapping.Table, trackmapping.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.TrackMappingsTable, order.TrackMappingsColumn),
		)
		schemaConfig := oq.schemaConfig
		step.To.Schema = schemaConfig.TrackMapping
		step.Edge.Schema = schemaConfig.TrackMapping
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Order entity from the query.
// Returns a *NotFoundError when no Order was found.
func (oq *OrderQuery) First(ctx context.Context) (*Order, error) {
	nodes, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{order.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OrderQuery) FirstX(ctx context.Context) *Order {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Order ID from the query.
// Returns a *NotFoundError when no Order ID was found.
func (oq *OrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{order.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Order entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Order entity is found.
// Returns a *NotFoundError when no Order entities are found.
func (oq *OrderQuery) Only(ctx context.Context) (*Order, error) {
	nodes, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{order.Label}
	default:
		return nil, &NotSingularError{order.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OrderQuery) OnlyX(ctx context.Context) *Order {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Order ID in the query.
// Returns a *NotSingularError when more than one Order ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{order.Label}
	default:
		err = &NotSingularError{order.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Orders.
func (oq *OrderQuery) All(ctx context.Context) ([]*Order, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *OrderQuery) AllX(ctx context.Context) []*Order {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Order IDs.
func (oq *OrderQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := oq.Select(order.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OrderQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OrderQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OrderQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OrderQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OrderQuery) Clone() *OrderQuery {
	if oq == nil {
		return nil
	}
	return &OrderQuery{
		config:               oq.config,
		limit:                oq.limit,
		offset:               oq.offset,
		order:                append([]OrderFunc{}, oq.order...),
		predicates:           append([]predicate.Order{}, oq.predicates...),
		withTenant:           oq.withTenant.Clone(),
		withOrderItems:       oq.withOrderItems.Clone(),
		withOrderHoldReasons: oq.withOrderHoldReasons.Clone(),
		withOrderTaxations:   oq.withOrderTaxations.Clone(),
		withStores:           oq.withStores.Clone(),
		withWarehouses:       oq.withWarehouses.Clone(),
		withChannels:         oq.withChannels.Clone(),
		withTrackMappings:    oq.withTrackMappings.Clone(),
		// clone intermediate query.
		sql:    oq.sql.Clone(),
		path:   oq.path,
		unique: oq.unique,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithTenant(opts ...func(*TenantQuery)) *OrderQuery {
	query := &TenantQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withTenant = query
	return oq
}

// WithOrderItems tells the query-builder to eager-load the nodes that are connected to
// the "order_items" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithOrderItems(opts ...func(*OrderItemQuery)) *OrderQuery {
	query := &OrderItemQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withOrderItems = query
	return oq
}

// WithOrderHoldReasons tells the query-builder to eager-load the nodes that are connected to
// the "order_hold_reasons" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithOrderHoldReasons(opts ...func(*OrderHoldReasonQuery)) *OrderQuery {
	query := &OrderHoldReasonQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withOrderHoldReasons = query
	return oq
}

// WithOrderTaxations tells the query-builder to eager-load the nodes that are connected to
// the "order_taxations" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithOrderTaxations(opts ...func(*OrderTaxationQuery)) *OrderQuery {
	query := &OrderTaxationQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withOrderTaxations = query
	return oq
}

// WithStores tells the query-builder to eager-load the nodes that are connected to
// the "stores" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithStores(opts ...func(*StoreQuery)) *OrderQuery {
	query := &StoreQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withStores = query
	return oq
}

// WithWarehouses tells the query-builder to eager-load the nodes that are connected to
// the "warehouses" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithWarehouses(opts ...func(*WarehouseQuery)) *OrderQuery {
	query := &WarehouseQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withWarehouses = query
	return oq
}

// WithChannels tells the query-builder to eager-load the nodes that are connected to
// the "channels" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithChannels(opts ...func(*ChannelQuery)) *OrderQuery {
	query := &ChannelQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withChannels = query
	return oq
}

// WithTrackMappings tells the query-builder to eager-load the nodes that are connected to
// the "track_mappings" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrderQuery) WithTrackMappings(opts ...func(*TrackMappingQuery)) *OrderQuery {
	query := &TrackMappingQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withTrackMappings = query
	return oq
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
//	client.Order.Query().
//		GroupBy(order.FieldCreatedAt).
//		Aggregate(gen.Count()).
//		Scan(ctx, &v)
func (oq *OrderQuery) GroupBy(field string, fields ...string) *OrderGroupBy {
	grbuild := &OrderGroupBy{config: oq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(ctx), nil
	}
	grbuild.label = order.Label
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
//	client.Order.Query().
//		Select(order.FieldCreatedAt).
//		Scan(ctx, &v)
func (oq *OrderQuery) Select(fields ...string) *OrderSelect {
	oq.fields = append(oq.fields, fields...)
	selbuild := &OrderSelect{OrderQuery: oq}
	selbuild.label = order.Label
	selbuild.flds, selbuild.scan = &oq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a OrderSelect configured with the given aggregations.
func (oq *OrderQuery) Aggregate(fns ...AggregateFunc) *OrderSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OrderQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oq.fields {
		if !order.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	if order.Policy == nil {
		return errors.New("gen: uninitialized order.Policy (forgotten import gen/runtime?)")
	}
	if err := order.Policy.EvalQuery(ctx, oq); err != nil {
		return err
	}
	return nil
}

func (oq *OrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Order, error) {
	var (
		nodes       = []*Order{}
		withFKs     = oq.withFKs
		_spec       = oq.querySpec()
		loadedTypes = [8]bool{
			oq.withTenant != nil,
			oq.withOrderItems != nil,
			oq.withOrderHoldReasons != nil,
			oq.withOrderTaxations != nil,
			oq.withStores != nil,
			oq.withWarehouses != nil,
			oq.withChannels != nil,
			oq.withTrackMappings != nil,
		}
	)
	if oq.withTenant != nil || oq.withStores != nil || oq.withWarehouses != nil || oq.withChannels != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, order.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Order).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Order{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = oq.schemaConfig.Order
	ctx = internal.NewSchemaConfigContext(ctx, oq.schemaConfig)
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withTenant; query != nil {
		if err := oq.loadTenant(ctx, query, nodes, nil,
			func(n *Order, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withOrderItems; query != nil {
		if err := oq.loadOrderItems(ctx, query, nodes,
			func(n *Order) { n.Edges.OrderItems = []*OrderItem{} },
			func(n *Order, e *OrderItem) { n.Edges.OrderItems = append(n.Edges.OrderItems, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withOrderHoldReasons; query != nil {
		if err := oq.loadOrderHoldReasons(ctx, query, nodes,
			func(n *Order) { n.Edges.OrderHoldReasons = []*OrderHoldReason{} },
			func(n *Order, e *OrderHoldReason) { n.Edges.OrderHoldReasons = append(n.Edges.OrderHoldReasons, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withOrderTaxations; query != nil {
		if err := oq.loadOrderTaxations(ctx, query, nodes,
			func(n *Order) { n.Edges.OrderTaxations = []*OrderTaxation{} },
			func(n *Order, e *OrderTaxation) { n.Edges.OrderTaxations = append(n.Edges.OrderTaxations, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withStores; query != nil {
		if err := oq.loadStores(ctx, query, nodes, nil,
			func(n *Order, e *Store) { n.Edges.Stores = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withWarehouses; query != nil {
		if err := oq.loadWarehouses(ctx, query, nodes, nil,
			func(n *Order, e *Warehouse) { n.Edges.Warehouses = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withChannels; query != nil {
		if err := oq.loadChannels(ctx, query, nodes, nil,
			func(n *Order, e *Channel) { n.Edges.Channels = e }); err != nil {
			return nil, err
		}
	}
	if query := oq.withTrackMappings; query != nil {
		if err := oq.loadTrackMappings(ctx, query, nodes,
			func(n *Order) { n.Edges.TrackMappings = []*TrackMapping{} },
			func(n *Order, e *TrackMapping) { n.Edges.TrackMappings = append(n.Edges.TrackMappings, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OrderQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*Order, init func(*Order), assign func(*Order, *Tenant)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Order)
	for i := range nodes {
		if nodes[i].order_tenant == nil {
			continue
		}
		fk := *nodes[i].order_tenant
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
			return fmt.Errorf(`unexpected foreign-key "order_tenant" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *OrderQuery) loadOrderItems(ctx context.Context, query *OrderItemQuery, nodes []*Order, init func(*Order), assign func(*Order, *OrderItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Order)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.InValues(order.OrderItemsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrderQuery) loadOrderHoldReasons(ctx context.Context, query *OrderHoldReasonQuery, nodes []*Order, init func(*Order), assign func(*Order, *OrderHoldReason)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Order)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.InValues(order.OrderHoldReasonsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrderQuery) loadOrderTaxations(ctx context.Context, query *OrderTaxationQuery, nodes []*Order, init func(*Order), assign func(*Order, *OrderTaxation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Order)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.OrderTaxation(func(s *sql.Selector) {
		s.Where(sql.InValues(order.OrderTaxationsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrderQuery) loadStores(ctx context.Context, query *StoreQuery, nodes []*Order, init func(*Order), assign func(*Order, *Store)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Order)
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
func (oq *OrderQuery) loadWarehouses(ctx context.Context, query *WarehouseQuery, nodes []*Order, init func(*Order), assign func(*Order, *Warehouse)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Order)
	for i := range nodes {
		fk := nodes[i].WarehouseID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(warehouse.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "warehouse_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *OrderQuery) loadChannels(ctx context.Context, query *ChannelQuery, nodes []*Order, init func(*Order), assign func(*Order, *Channel)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Order)
	for i := range nodes {
		fk := nodes[i].ChannelID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(channel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "channel_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (oq *OrderQuery) loadTrackMappings(ctx context.Context, query *TrackMappingQuery, nodes []*Order, init func(*Order), assign func(*Order, *TrackMapping)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Order)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.TrackMapping(func(s *sql.Selector) {
		s.Where(sql.InValues(order.TrackMappingsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oq *OrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Schema = oq.schemaConfig.Order
	ctx = internal.NewSchemaConfigContext(ctx, oq.schemaConfig)
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	_spec.Node.Columns = oq.fields
	if len(oq.fields) > 0 {
		_spec.Unique = oq.unique != nil && *oq.unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OrderQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("gen: check existence: %w", err)
	default:
		return true, nil
	}
}

func (oq *OrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: order.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for i := range fields {
			if fields[i] != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(order.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = order.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.unique != nil && *oq.unique {
		selector.Distinct()
	}
	t1.Schema(oq.schemaConfig.Order)
	ctx = internal.NewSchemaConfigContext(ctx, oq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range oq.modifiers {
		m(selector)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oq *OrderQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderSelect {
	oq.modifiers = append(oq.modifiers, modifiers...)
	return oq.Select()
}

// OrderGroupBy is the group-by builder for Order entities.
type OrderGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OrderGroupBy) Aggregate(fns ...AggregateFunc) *OrderGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scans the result into the given value.
func (ogb *OrderGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

func (ogb *OrderGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ogb.fields {
		if !order.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OrderGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql.Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
		for _, f := range ogb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ogb.fields...)...)
}

// OrderSelect is the builder for selecting fields of Order entities.
type OrderSelect struct {
	*OrderQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OrderSelect) Aggregate(fns ...AggregateFunc) *OrderSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OrderSelect) Scan(ctx context.Context, v any) error {
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	os.sql = os.OrderQuery.sqlQuery(ctx)
	return os.sqlScan(ctx, v)
}

func (os *OrderSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(os.sql))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		os.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		os.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := os.sql.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (os *OrderSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderSelect {
	os.modifiers = append(os.modifiers, modifiers...)
	return os
}
