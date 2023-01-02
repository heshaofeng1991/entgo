// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/trackmapping"
)

// TrackMappingUpdate is the builder for updating TrackMapping entities.
type TrackMappingUpdate struct {
	config
	hooks     []Hook
	mutation  *TrackMappingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TrackMappingUpdate builder.
func (tmu *TrackMappingUpdate) Where(ps ...predicate.TrackMapping) *TrackMappingUpdate {
	tmu.mutation.Where(ps...)
	return tmu
}

// SetUpdatedAt sets the "updated_at" field.
func (tmu *TrackMappingUpdate) SetUpdatedAt(t time.Time) *TrackMappingUpdate {
	tmu.mutation.SetUpdatedAt(t)
	return tmu
}

// SetDeletedAt sets the "deleted_at" field.
func (tmu *TrackMappingUpdate) SetDeletedAt(t time.Time) *TrackMappingUpdate {
	tmu.mutation.SetDeletedAt(t)
	return tmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableDeletedAt(t *time.Time) *TrackMappingUpdate {
	if t != nil {
		tmu.SetDeletedAt(*t)
	}
	return tmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tmu *TrackMappingUpdate) ClearDeletedAt() *TrackMappingUpdate {
	tmu.mutation.ClearDeletedAt()
	return tmu
}

// SetOrderID sets the "order_id" field.
func (tmu *TrackMappingUpdate) SetOrderID(i int64) *TrackMappingUpdate {
	tmu.mutation.SetOrderID(i)
	return tmu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableOrderID(i *int64) *TrackMappingUpdate {
	if i != nil {
		tmu.SetOrderID(*i)
	}
	return tmu
}

// ClearOrderID clears the value of the "order_id" field.
func (tmu *TrackMappingUpdate) ClearOrderID() *TrackMappingUpdate {
	tmu.mutation.ClearOrderID()
	return tmu
}

// SetTrackingNumber sets the "tracking_number" field.
func (tmu *TrackMappingUpdate) SetTrackingNumber(s string) *TrackMappingUpdate {
	tmu.mutation.SetTrackingNumber(s)
	return tmu
}

// SetTrackingURL sets the "tracking_url" field.
func (tmu *TrackMappingUpdate) SetTrackingURL(s string) *TrackMappingUpdate {
	tmu.mutation.SetTrackingURL(s)
	return tmu
}

// SetNillableTrackingURL sets the "tracking_url" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableTrackingURL(s *string) *TrackMappingUpdate {
	if s != nil {
		tmu.SetTrackingURL(*s)
	}
	return tmu
}

// SetExtTrackingNumber sets the "ext_tracking_number" field.
func (tmu *TrackMappingUpdate) SetExtTrackingNumber(s string) *TrackMappingUpdate {
	tmu.mutation.SetExtTrackingNumber(s)
	return tmu
}

// SetTrackDetails sets the "track_details" field.
func (tmu *TrackMappingUpdate) SetTrackDetails(s string) *TrackMappingUpdate {
	tmu.mutation.SetTrackDetails(s)
	return tmu
}

// SetNillableTrackDetails sets the "track_details" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableTrackDetails(s *string) *TrackMappingUpdate {
	if s != nil {
		tmu.SetTrackDetails(*s)
	}
	return tmu
}

// ClearTrackDetails clears the value of the "track_details" field.
func (tmu *TrackMappingUpdate) ClearTrackDetails() *TrackMappingUpdate {
	tmu.mutation.ClearTrackDetails()
	return tmu
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (tmu *TrackMappingUpdate) SetLastUpdatedAt(t time.Time) *TrackMappingUpdate {
	tmu.mutation.SetLastUpdatedAt(t)
	return tmu
}

// SetNillableLastUpdatedAt sets the "last_updated_at" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableLastUpdatedAt(t *time.Time) *TrackMappingUpdate {
	if t != nil {
		tmu.SetLastUpdatedAt(*t)
	}
	return tmu
}

// ClearLastUpdatedAt clears the value of the "last_updated_at" field.
func (tmu *TrackMappingUpdate) ClearLastUpdatedAt() *TrackMappingUpdate {
	tmu.mutation.ClearLastUpdatedAt()
	return tmu
}

// SetCourierPlatform sets the "courier_platform" field.
func (tmu *TrackMappingUpdate) SetCourierPlatform(s string) *TrackMappingUpdate {
	tmu.mutation.SetCourierPlatform(s)
	return tmu
}

// SetNillableCourierPlatform sets the "courier_platform" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableCourierPlatform(s *string) *TrackMappingUpdate {
	if s != nil {
		tmu.SetCourierPlatform(*s)
	}
	return tmu
}

// SetStatus sets the "status" field.
func (tmu *TrackMappingUpdate) SetStatus(i int32) *TrackMappingUpdate {
	tmu.mutation.ResetStatus()
	tmu.mutation.SetStatus(i)
	return tmu
}

// AddStatus adds i to the "status" field.
func (tmu *TrackMappingUpdate) AddStatus(i int32) *TrackMappingUpdate {
	tmu.mutation.AddStatus(i)
	return tmu
}

// SetFlag sets the "flag" field.
func (tmu *TrackMappingUpdate) SetFlag(i int8) *TrackMappingUpdate {
	tmu.mutation.ResetFlag()
	tmu.mutation.SetFlag(i)
	return tmu
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableFlag(i *int8) *TrackMappingUpdate {
	if i != nil {
		tmu.SetFlag(*i)
	}
	return tmu
}

// AddFlag adds i to the "flag" field.
func (tmu *TrackMappingUpdate) AddFlag(i int8) *TrackMappingUpdate {
	tmu.mutation.AddFlag(i)
	return tmu
}

// SetOrdersID sets the "orders" edge to the Order entity by ID.
func (tmu *TrackMappingUpdate) SetOrdersID(id int64) *TrackMappingUpdate {
	tmu.mutation.SetOrdersID(id)
	return tmu
}

// SetNillableOrdersID sets the "orders" edge to the Order entity by ID if the given value is not nil.
func (tmu *TrackMappingUpdate) SetNillableOrdersID(id *int64) *TrackMappingUpdate {
	if id != nil {
		tmu = tmu.SetOrdersID(*id)
	}
	return tmu
}

// SetOrders sets the "orders" edge to the Order entity.
func (tmu *TrackMappingUpdate) SetOrders(o *Order) *TrackMappingUpdate {
	return tmu.SetOrdersID(o.ID)
}

// Mutation returns the TrackMappingMutation object of the builder.
func (tmu *TrackMappingUpdate) Mutation() *TrackMappingMutation {
	return tmu.mutation
}

// ClearOrders clears the "orders" edge to the Order entity.
func (tmu *TrackMappingUpdate) ClearOrders() *TrackMappingUpdate {
	tmu.mutation.ClearOrders()
	return tmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmu *TrackMappingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tmu.defaults()
	if len(tmu.hooks) == 0 {
		affected, err = tmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackMappingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmu.mutation = mutation
			affected, err = tmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tmu.hooks) - 1; i >= 0; i-- {
			if tmu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = tmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmu *TrackMappingUpdate) SaveX(ctx context.Context) int {
	affected, err := tmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmu *TrackMappingUpdate) Exec(ctx context.Context) error {
	_, err := tmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmu *TrackMappingUpdate) ExecX(ctx context.Context) {
	if err := tmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmu *TrackMappingUpdate) defaults() {
	if _, ok := tmu.mutation.UpdatedAt(); !ok {
		v := trackmapping.UpdateDefaultUpdatedAt()
		tmu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmu *TrackMappingUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TrackMappingUpdate {
	tmu.modifiers = append(tmu.modifiers, modifiers...)
	return tmu
}

func (tmu *TrackMappingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   trackmapping.Table,
			Columns: trackmapping.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: trackmapping.FieldID,
			},
		},
	}
	if ps := tmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmu.mutation.UpdatedAt(); ok {
		_spec.SetField(trackmapping.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tmu.mutation.DeletedAt(); ok {
		_spec.SetField(trackmapping.FieldDeletedAt, field.TypeTime, value)
	}
	if tmu.mutation.DeletedAtCleared() {
		_spec.ClearField(trackmapping.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tmu.mutation.TrackingNumber(); ok {
		_spec.SetField(trackmapping.FieldTrackingNumber, field.TypeString, value)
	}
	if value, ok := tmu.mutation.TrackingURL(); ok {
		_spec.SetField(trackmapping.FieldTrackingURL, field.TypeString, value)
	}
	if value, ok := tmu.mutation.ExtTrackingNumber(); ok {
		_spec.SetField(trackmapping.FieldExtTrackingNumber, field.TypeString, value)
	}
	if value, ok := tmu.mutation.TrackDetails(); ok {
		_spec.SetField(trackmapping.FieldTrackDetails, field.TypeString, value)
	}
	if tmu.mutation.TrackDetailsCleared() {
		_spec.ClearField(trackmapping.FieldTrackDetails, field.TypeString)
	}
	if value, ok := tmu.mutation.LastUpdatedAt(); ok {
		_spec.SetField(trackmapping.FieldLastUpdatedAt, field.TypeTime, value)
	}
	if tmu.mutation.LastUpdatedAtCleared() {
		_spec.ClearField(trackmapping.FieldLastUpdatedAt, field.TypeTime)
	}
	if value, ok := tmu.mutation.CourierPlatform(); ok {
		_spec.SetField(trackmapping.FieldCourierPlatform, field.TypeString, value)
	}
	if value, ok := tmu.mutation.Status(); ok {
		_spec.SetField(trackmapping.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := tmu.mutation.AddedStatus(); ok {
		_spec.AddField(trackmapping.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := tmu.mutation.Flag(); ok {
		_spec.SetField(trackmapping.FieldFlag, field.TypeInt8, value)
	}
	if value, ok := tmu.mutation.AddedFlag(); ok {
		_spec.AddField(trackmapping.FieldFlag, field.TypeInt8, value)
	}
	if tmu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   trackmapping.OrdersTable,
			Columns: []string{trackmapping.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		edge.Schema = tmu.schemaConfig.TrackMapping
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tmu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   trackmapping.OrdersTable,
			Columns: []string{trackmapping.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		edge.Schema = tmu.schemaConfig.TrackMapping
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tmu.schemaConfig.TrackMapping
	ctx = internal.NewSchemaConfigContext(ctx, tmu.schemaConfig)
	_spec.AddModifiers(tmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trackmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TrackMappingUpdateOne is the builder for updating a single TrackMapping entity.
type TrackMappingUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TrackMappingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tmuo *TrackMappingUpdateOne) SetUpdatedAt(t time.Time) *TrackMappingUpdateOne {
	tmuo.mutation.SetUpdatedAt(t)
	return tmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tmuo *TrackMappingUpdateOne) SetDeletedAt(t time.Time) *TrackMappingUpdateOne {
	tmuo.mutation.SetDeletedAt(t)
	return tmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableDeletedAt(t *time.Time) *TrackMappingUpdateOne {
	if t != nil {
		tmuo.SetDeletedAt(*t)
	}
	return tmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tmuo *TrackMappingUpdateOne) ClearDeletedAt() *TrackMappingUpdateOne {
	tmuo.mutation.ClearDeletedAt()
	return tmuo
}

// SetOrderID sets the "order_id" field.
func (tmuo *TrackMappingUpdateOne) SetOrderID(i int64) *TrackMappingUpdateOne {
	tmuo.mutation.SetOrderID(i)
	return tmuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableOrderID(i *int64) *TrackMappingUpdateOne {
	if i != nil {
		tmuo.SetOrderID(*i)
	}
	return tmuo
}

// ClearOrderID clears the value of the "order_id" field.
func (tmuo *TrackMappingUpdateOne) ClearOrderID() *TrackMappingUpdateOne {
	tmuo.mutation.ClearOrderID()
	return tmuo
}

// SetTrackingNumber sets the "tracking_number" field.
func (tmuo *TrackMappingUpdateOne) SetTrackingNumber(s string) *TrackMappingUpdateOne {
	tmuo.mutation.SetTrackingNumber(s)
	return tmuo
}

// SetTrackingURL sets the "tracking_url" field.
func (tmuo *TrackMappingUpdateOne) SetTrackingURL(s string) *TrackMappingUpdateOne {
	tmuo.mutation.SetTrackingURL(s)
	return tmuo
}

// SetNillableTrackingURL sets the "tracking_url" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableTrackingURL(s *string) *TrackMappingUpdateOne {
	if s != nil {
		tmuo.SetTrackingURL(*s)
	}
	return tmuo
}

// SetExtTrackingNumber sets the "ext_tracking_number" field.
func (tmuo *TrackMappingUpdateOne) SetExtTrackingNumber(s string) *TrackMappingUpdateOne {
	tmuo.mutation.SetExtTrackingNumber(s)
	return tmuo
}

// SetTrackDetails sets the "track_details" field.
func (tmuo *TrackMappingUpdateOne) SetTrackDetails(s string) *TrackMappingUpdateOne {
	tmuo.mutation.SetTrackDetails(s)
	return tmuo
}

// SetNillableTrackDetails sets the "track_details" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableTrackDetails(s *string) *TrackMappingUpdateOne {
	if s != nil {
		tmuo.SetTrackDetails(*s)
	}
	return tmuo
}

// ClearTrackDetails clears the value of the "track_details" field.
func (tmuo *TrackMappingUpdateOne) ClearTrackDetails() *TrackMappingUpdateOne {
	tmuo.mutation.ClearTrackDetails()
	return tmuo
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (tmuo *TrackMappingUpdateOne) SetLastUpdatedAt(t time.Time) *TrackMappingUpdateOne {
	tmuo.mutation.SetLastUpdatedAt(t)
	return tmuo
}

// SetNillableLastUpdatedAt sets the "last_updated_at" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableLastUpdatedAt(t *time.Time) *TrackMappingUpdateOne {
	if t != nil {
		tmuo.SetLastUpdatedAt(*t)
	}
	return tmuo
}

// ClearLastUpdatedAt clears the value of the "last_updated_at" field.
func (tmuo *TrackMappingUpdateOne) ClearLastUpdatedAt() *TrackMappingUpdateOne {
	tmuo.mutation.ClearLastUpdatedAt()
	return tmuo
}

// SetCourierPlatform sets the "courier_platform" field.
func (tmuo *TrackMappingUpdateOne) SetCourierPlatform(s string) *TrackMappingUpdateOne {
	tmuo.mutation.SetCourierPlatform(s)
	return tmuo
}

// SetNillableCourierPlatform sets the "courier_platform" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableCourierPlatform(s *string) *TrackMappingUpdateOne {
	if s != nil {
		tmuo.SetCourierPlatform(*s)
	}
	return tmuo
}

// SetStatus sets the "status" field.
func (tmuo *TrackMappingUpdateOne) SetStatus(i int32) *TrackMappingUpdateOne {
	tmuo.mutation.ResetStatus()
	tmuo.mutation.SetStatus(i)
	return tmuo
}

// AddStatus adds i to the "status" field.
func (tmuo *TrackMappingUpdateOne) AddStatus(i int32) *TrackMappingUpdateOne {
	tmuo.mutation.AddStatus(i)
	return tmuo
}

// SetFlag sets the "flag" field.
func (tmuo *TrackMappingUpdateOne) SetFlag(i int8) *TrackMappingUpdateOne {
	tmuo.mutation.ResetFlag()
	tmuo.mutation.SetFlag(i)
	return tmuo
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableFlag(i *int8) *TrackMappingUpdateOne {
	if i != nil {
		tmuo.SetFlag(*i)
	}
	return tmuo
}

// AddFlag adds i to the "flag" field.
func (tmuo *TrackMappingUpdateOne) AddFlag(i int8) *TrackMappingUpdateOne {
	tmuo.mutation.AddFlag(i)
	return tmuo
}

// SetOrdersID sets the "orders" edge to the Order entity by ID.
func (tmuo *TrackMappingUpdateOne) SetOrdersID(id int64) *TrackMappingUpdateOne {
	tmuo.mutation.SetOrdersID(id)
	return tmuo
}

// SetNillableOrdersID sets the "orders" edge to the Order entity by ID if the given value is not nil.
func (tmuo *TrackMappingUpdateOne) SetNillableOrdersID(id *int64) *TrackMappingUpdateOne {
	if id != nil {
		tmuo = tmuo.SetOrdersID(*id)
	}
	return tmuo
}

// SetOrders sets the "orders" edge to the Order entity.
func (tmuo *TrackMappingUpdateOne) SetOrders(o *Order) *TrackMappingUpdateOne {
	return tmuo.SetOrdersID(o.ID)
}

// Mutation returns the TrackMappingMutation object of the builder.
func (tmuo *TrackMappingUpdateOne) Mutation() *TrackMappingMutation {
	return tmuo.mutation
}

// ClearOrders clears the "orders" edge to the Order entity.
func (tmuo *TrackMappingUpdateOne) ClearOrders() *TrackMappingUpdateOne {
	tmuo.mutation.ClearOrders()
	return tmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tmuo *TrackMappingUpdateOne) Select(field string, fields ...string) *TrackMappingUpdateOne {
	tmuo.fields = append([]string{field}, fields...)
	return tmuo
}

// Save executes the query and returns the updated TrackMapping entity.
func (tmuo *TrackMappingUpdateOne) Save(ctx context.Context) (*TrackMapping, error) {
	var (
		err  error
		node *TrackMapping
	)
	tmuo.defaults()
	if len(tmuo.hooks) == 0 {
		node, err = tmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackMappingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmuo.mutation = mutation
			node, err = tmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tmuo.hooks) - 1; i >= 0; i-- {
			if tmuo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = tmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TrackMapping)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TrackMappingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmuo *TrackMappingUpdateOne) SaveX(ctx context.Context) *TrackMapping {
	node, err := tmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmuo *TrackMappingUpdateOne) Exec(ctx context.Context) error {
	_, err := tmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmuo *TrackMappingUpdateOne) ExecX(ctx context.Context) {
	if err := tmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmuo *TrackMappingUpdateOne) defaults() {
	if _, ok := tmuo.mutation.UpdatedAt(); !ok {
		v := trackmapping.UpdateDefaultUpdatedAt()
		tmuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmuo *TrackMappingUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TrackMappingUpdateOne {
	tmuo.modifiers = append(tmuo.modifiers, modifiers...)
	return tmuo
}

func (tmuo *TrackMappingUpdateOne) sqlSave(ctx context.Context) (_node *TrackMapping, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   trackmapping.Table,
			Columns: trackmapping.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: trackmapping.FieldID,
			},
		},
	}
	id, ok := tmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "TrackMapping.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trackmapping.FieldID)
		for _, f := range fields {
			if !trackmapping.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != trackmapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(trackmapping.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tmuo.mutation.DeletedAt(); ok {
		_spec.SetField(trackmapping.FieldDeletedAt, field.TypeTime, value)
	}
	if tmuo.mutation.DeletedAtCleared() {
		_spec.ClearField(trackmapping.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tmuo.mutation.TrackingNumber(); ok {
		_spec.SetField(trackmapping.FieldTrackingNumber, field.TypeString, value)
	}
	if value, ok := tmuo.mutation.TrackingURL(); ok {
		_spec.SetField(trackmapping.FieldTrackingURL, field.TypeString, value)
	}
	if value, ok := tmuo.mutation.ExtTrackingNumber(); ok {
		_spec.SetField(trackmapping.FieldExtTrackingNumber, field.TypeString, value)
	}
	if value, ok := tmuo.mutation.TrackDetails(); ok {
		_spec.SetField(trackmapping.FieldTrackDetails, field.TypeString, value)
	}
	if tmuo.mutation.TrackDetailsCleared() {
		_spec.ClearField(trackmapping.FieldTrackDetails, field.TypeString)
	}
	if value, ok := tmuo.mutation.LastUpdatedAt(); ok {
		_spec.SetField(trackmapping.FieldLastUpdatedAt, field.TypeTime, value)
	}
	if tmuo.mutation.LastUpdatedAtCleared() {
		_spec.ClearField(trackmapping.FieldLastUpdatedAt, field.TypeTime)
	}
	if value, ok := tmuo.mutation.CourierPlatform(); ok {
		_spec.SetField(trackmapping.FieldCourierPlatform, field.TypeString, value)
	}
	if value, ok := tmuo.mutation.Status(); ok {
		_spec.SetField(trackmapping.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := tmuo.mutation.AddedStatus(); ok {
		_spec.AddField(trackmapping.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := tmuo.mutation.Flag(); ok {
		_spec.SetField(trackmapping.FieldFlag, field.TypeInt8, value)
	}
	if value, ok := tmuo.mutation.AddedFlag(); ok {
		_spec.AddField(trackmapping.FieldFlag, field.TypeInt8, value)
	}
	if tmuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   trackmapping.OrdersTable,
			Columns: []string{trackmapping.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		edge.Schema = tmuo.schemaConfig.TrackMapping
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tmuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   trackmapping.OrdersTable,
			Columns: []string{trackmapping.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		edge.Schema = tmuo.schemaConfig.TrackMapping
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tmuo.schemaConfig.TrackMapping
	ctx = internal.NewSchemaConfigContext(ctx, tmuo.schemaConfig)
	_spec.AddModifiers(tmuo.modifiers...)
	_node = &TrackMapping{config: tmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trackmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}