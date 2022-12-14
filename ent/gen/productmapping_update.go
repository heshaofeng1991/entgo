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
	"github.com/heshaofeng1991/entgo/ent/gen/platformproduct"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/product"
	"github.com/heshaofeng1991/entgo/ent/gen/productmapping"
)

// ProductMappingUpdate is the builder for updating ProductMapping entities.
type ProductMappingUpdate struct {
	config
	hooks     []Hook
	mutation  *ProductMappingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProductMappingUpdate builder.
func (pmu *ProductMappingUpdate) Where(ps ...predicate.ProductMapping) *ProductMappingUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetUpdatedAt sets the "updated_at" field.
func (pmu *ProductMappingUpdate) SetUpdatedAt(t time.Time) *ProductMappingUpdate {
	pmu.mutation.SetUpdatedAt(t)
	return pmu
}

// SetDeletedAt sets the "deleted_at" field.
func (pmu *ProductMappingUpdate) SetDeletedAt(t time.Time) *ProductMappingUpdate {
	pmu.mutation.SetDeletedAt(t)
	return pmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pmu *ProductMappingUpdate) SetNillableDeletedAt(t *time.Time) *ProductMappingUpdate {
	if t != nil {
		pmu.SetDeletedAt(*t)
	}
	return pmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pmu *ProductMappingUpdate) ClearDeletedAt() *ProductMappingUpdate {
	pmu.mutation.ClearDeletedAt()
	return pmu
}

// SetPlatformProductID sets the "platform_product_id" field.
func (pmu *ProductMappingUpdate) SetPlatformProductID(i int64) *ProductMappingUpdate {
	pmu.mutation.SetPlatformProductID(i)
	return pmu
}

// SetNillablePlatformProductID sets the "platform_product_id" field if the given value is not nil.
func (pmu *ProductMappingUpdate) SetNillablePlatformProductID(i *int64) *ProductMappingUpdate {
	if i != nil {
		pmu.SetPlatformProductID(*i)
	}
	return pmu
}

// ClearPlatformProductID clears the value of the "platform_product_id" field.
func (pmu *ProductMappingUpdate) ClearPlatformProductID() *ProductMappingUpdate {
	pmu.mutation.ClearPlatformProductID()
	return pmu
}

// SetProductID sets the "product_id" field.
func (pmu *ProductMappingUpdate) SetProductID(i int64) *ProductMappingUpdate {
	pmu.mutation.SetProductID(i)
	return pmu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (pmu *ProductMappingUpdate) SetNillableProductID(i *int64) *ProductMappingUpdate {
	if i != nil {
		pmu.SetProductID(*i)
	}
	return pmu
}

// ClearProductID clears the value of the "product_id" field.
func (pmu *ProductMappingUpdate) ClearProductID() *ProductMappingUpdate {
	pmu.mutation.ClearProductID()
	return pmu
}

// SetQty sets the "qty" field.
func (pmu *ProductMappingUpdate) SetQty(i int) *ProductMappingUpdate {
	pmu.mutation.ResetQty()
	pmu.mutation.SetQty(i)
	return pmu
}

// AddQty adds i to the "qty" field.
func (pmu *ProductMappingUpdate) AddQty(i int) *ProductMappingUpdate {
	pmu.mutation.AddQty(i)
	return pmu
}

// SetCreatedBy sets the "created_by" field.
func (pmu *ProductMappingUpdate) SetCreatedBy(i int64) *ProductMappingUpdate {
	pmu.mutation.ResetCreatedBy()
	pmu.mutation.SetCreatedBy(i)
	return pmu
}

// AddCreatedBy adds i to the "created_by" field.
func (pmu *ProductMappingUpdate) AddCreatedBy(i int64) *ProductMappingUpdate {
	pmu.mutation.AddCreatedBy(i)
	return pmu
}

// SetPlatformProductsID sets the "platform_products" edge to the PlatformProduct entity by ID.
func (pmu *ProductMappingUpdate) SetPlatformProductsID(id int64) *ProductMappingUpdate {
	pmu.mutation.SetPlatformProductsID(id)
	return pmu
}

// SetNillablePlatformProductsID sets the "platform_products" edge to the PlatformProduct entity by ID if the given value is not nil.
func (pmu *ProductMappingUpdate) SetNillablePlatformProductsID(id *int64) *ProductMappingUpdate {
	if id != nil {
		pmu = pmu.SetPlatformProductsID(*id)
	}
	return pmu
}

// SetPlatformProducts sets the "platform_products" edge to the PlatformProduct entity.
func (pmu *ProductMappingUpdate) SetPlatformProducts(p *PlatformProduct) *ProductMappingUpdate {
	return pmu.SetPlatformProductsID(p.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (pmu *ProductMappingUpdate) SetProductsID(id int64) *ProductMappingUpdate {
	pmu.mutation.SetProductsID(id)
	return pmu
}

// SetNillableProductsID sets the "products" edge to the Product entity by ID if the given value is not nil.
func (pmu *ProductMappingUpdate) SetNillableProductsID(id *int64) *ProductMappingUpdate {
	if id != nil {
		pmu = pmu.SetProductsID(*id)
	}
	return pmu
}

// SetProducts sets the "products" edge to the Product entity.
func (pmu *ProductMappingUpdate) SetProducts(p *Product) *ProductMappingUpdate {
	return pmu.SetProductsID(p.ID)
}

// Mutation returns the ProductMappingMutation object of the builder.
func (pmu *ProductMappingUpdate) Mutation() *ProductMappingMutation {
	return pmu.mutation
}

// ClearPlatformProducts clears the "platform_products" edge to the PlatformProduct entity.
func (pmu *ProductMappingUpdate) ClearPlatformProducts() *ProductMappingUpdate {
	pmu.mutation.ClearPlatformProducts()
	return pmu
}

// ClearProducts clears the "products" edge to the Product entity.
func (pmu *ProductMappingUpdate) ClearProducts() *ProductMappingUpdate {
	pmu.mutation.ClearProducts()
	return pmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *ProductMappingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pmu.defaults()
	if len(pmu.hooks) == 0 {
		affected, err = pmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMappingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pmu.mutation = mutation
			affected, err = pmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pmu.hooks) - 1; i >= 0; i-- {
			if pmu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = pmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *ProductMappingUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *ProductMappingUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *ProductMappingUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmu *ProductMappingUpdate) defaults() {
	if _, ok := pmu.mutation.UpdatedAt(); !ok {
		v := productmapping.UpdateDefaultUpdatedAt()
		pmu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pmu *ProductMappingUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProductMappingUpdate {
	pmu.modifiers = append(pmu.modifiers, modifiers...)
	return pmu
}

func (pmu *ProductMappingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productmapping.Table,
			Columns: productmapping.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: productmapping.FieldID,
			},
		},
	}
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.UpdatedAt(); ok {
		_spec.SetField(productmapping.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pmu.mutation.DeletedAt(); ok {
		_spec.SetField(productmapping.FieldDeletedAt, field.TypeTime, value)
	}
	if pmu.mutation.DeletedAtCleared() {
		_spec.ClearField(productmapping.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pmu.mutation.Qty(); ok {
		_spec.SetField(productmapping.FieldQty, field.TypeInt, value)
	}
	if value, ok := pmu.mutation.AddedQty(); ok {
		_spec.AddField(productmapping.FieldQty, field.TypeInt, value)
	}
	if value, ok := pmu.mutation.CreatedBy(); ok {
		_spec.SetField(productmapping.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pmu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(productmapping.FieldCreatedBy, field.TypeInt64, value)
	}
	if pmu.mutation.PlatformProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.PlatformProductsTable,
			Columns: []string{productmapping.PlatformProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: platformproduct.FieldID,
				},
			},
		}
		edge.Schema = pmu.schemaConfig.ProductMapping
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.PlatformProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.PlatformProductsTable,
			Columns: []string{productmapping.PlatformProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: platformproduct.FieldID,
				},
			},
		}
		edge.Schema = pmu.schemaConfig.ProductMapping
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pmu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.ProductsTable,
			Columns: []string{productmapping.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: product.FieldID,
				},
			},
		}
		edge.Schema = pmu.schemaConfig.ProductMapping
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.ProductsTable,
			Columns: []string{productmapping.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: product.FieldID,
				},
			},
		}
		edge.Schema = pmu.schemaConfig.ProductMapping
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = pmu.schemaConfig.ProductMapping
	ctx = internal.NewSchemaConfigContext(ctx, pmu.schemaConfig)
	_spec.AddModifiers(pmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ProductMappingUpdateOne is the builder for updating a single ProductMapping entity.
type ProductMappingUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProductMappingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (pmuo *ProductMappingUpdateOne) SetUpdatedAt(t time.Time) *ProductMappingUpdateOne {
	pmuo.mutation.SetUpdatedAt(t)
	return pmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pmuo *ProductMappingUpdateOne) SetDeletedAt(t time.Time) *ProductMappingUpdateOne {
	pmuo.mutation.SetDeletedAt(t)
	return pmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pmuo *ProductMappingUpdateOne) SetNillableDeletedAt(t *time.Time) *ProductMappingUpdateOne {
	if t != nil {
		pmuo.SetDeletedAt(*t)
	}
	return pmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pmuo *ProductMappingUpdateOne) ClearDeletedAt() *ProductMappingUpdateOne {
	pmuo.mutation.ClearDeletedAt()
	return pmuo
}

// SetPlatformProductID sets the "platform_product_id" field.
func (pmuo *ProductMappingUpdateOne) SetPlatformProductID(i int64) *ProductMappingUpdateOne {
	pmuo.mutation.SetPlatformProductID(i)
	return pmuo
}

// SetNillablePlatformProductID sets the "platform_product_id" field if the given value is not nil.
func (pmuo *ProductMappingUpdateOne) SetNillablePlatformProductID(i *int64) *ProductMappingUpdateOne {
	if i != nil {
		pmuo.SetPlatformProductID(*i)
	}
	return pmuo
}

// ClearPlatformProductID clears the value of the "platform_product_id" field.
func (pmuo *ProductMappingUpdateOne) ClearPlatformProductID() *ProductMappingUpdateOne {
	pmuo.mutation.ClearPlatformProductID()
	return pmuo
}

// SetProductID sets the "product_id" field.
func (pmuo *ProductMappingUpdateOne) SetProductID(i int64) *ProductMappingUpdateOne {
	pmuo.mutation.SetProductID(i)
	return pmuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (pmuo *ProductMappingUpdateOne) SetNillableProductID(i *int64) *ProductMappingUpdateOne {
	if i != nil {
		pmuo.SetProductID(*i)
	}
	return pmuo
}

// ClearProductID clears the value of the "product_id" field.
func (pmuo *ProductMappingUpdateOne) ClearProductID() *ProductMappingUpdateOne {
	pmuo.mutation.ClearProductID()
	return pmuo
}

// SetQty sets the "qty" field.
func (pmuo *ProductMappingUpdateOne) SetQty(i int) *ProductMappingUpdateOne {
	pmuo.mutation.ResetQty()
	pmuo.mutation.SetQty(i)
	return pmuo
}

// AddQty adds i to the "qty" field.
func (pmuo *ProductMappingUpdateOne) AddQty(i int) *ProductMappingUpdateOne {
	pmuo.mutation.AddQty(i)
	return pmuo
}

// SetCreatedBy sets the "created_by" field.
func (pmuo *ProductMappingUpdateOne) SetCreatedBy(i int64) *ProductMappingUpdateOne {
	pmuo.mutation.ResetCreatedBy()
	pmuo.mutation.SetCreatedBy(i)
	return pmuo
}

// AddCreatedBy adds i to the "created_by" field.
func (pmuo *ProductMappingUpdateOne) AddCreatedBy(i int64) *ProductMappingUpdateOne {
	pmuo.mutation.AddCreatedBy(i)
	return pmuo
}

// SetPlatformProductsID sets the "platform_products" edge to the PlatformProduct entity by ID.
func (pmuo *ProductMappingUpdateOne) SetPlatformProductsID(id int64) *ProductMappingUpdateOne {
	pmuo.mutation.SetPlatformProductsID(id)
	return pmuo
}

// SetNillablePlatformProductsID sets the "platform_products" edge to the PlatformProduct entity by ID if the given value is not nil.
func (pmuo *ProductMappingUpdateOne) SetNillablePlatformProductsID(id *int64) *ProductMappingUpdateOne {
	if id != nil {
		pmuo = pmuo.SetPlatformProductsID(*id)
	}
	return pmuo
}

// SetPlatformProducts sets the "platform_products" edge to the PlatformProduct entity.
func (pmuo *ProductMappingUpdateOne) SetPlatformProducts(p *PlatformProduct) *ProductMappingUpdateOne {
	return pmuo.SetPlatformProductsID(p.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (pmuo *ProductMappingUpdateOne) SetProductsID(id int64) *ProductMappingUpdateOne {
	pmuo.mutation.SetProductsID(id)
	return pmuo
}

// SetNillableProductsID sets the "products" edge to the Product entity by ID if the given value is not nil.
func (pmuo *ProductMappingUpdateOne) SetNillableProductsID(id *int64) *ProductMappingUpdateOne {
	if id != nil {
		pmuo = pmuo.SetProductsID(*id)
	}
	return pmuo
}

// SetProducts sets the "products" edge to the Product entity.
func (pmuo *ProductMappingUpdateOne) SetProducts(p *Product) *ProductMappingUpdateOne {
	return pmuo.SetProductsID(p.ID)
}

// Mutation returns the ProductMappingMutation object of the builder.
func (pmuo *ProductMappingUpdateOne) Mutation() *ProductMappingMutation {
	return pmuo.mutation
}

// ClearPlatformProducts clears the "platform_products" edge to the PlatformProduct entity.
func (pmuo *ProductMappingUpdateOne) ClearPlatformProducts() *ProductMappingUpdateOne {
	pmuo.mutation.ClearPlatformProducts()
	return pmuo
}

// ClearProducts clears the "products" edge to the Product entity.
func (pmuo *ProductMappingUpdateOne) ClearProducts() *ProductMappingUpdateOne {
	pmuo.mutation.ClearProducts()
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *ProductMappingUpdateOne) Select(field string, fields ...string) *ProductMappingUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated ProductMapping entity.
func (pmuo *ProductMappingUpdateOne) Save(ctx context.Context) (*ProductMapping, error) {
	var (
		err  error
		node *ProductMapping
	)
	pmuo.defaults()
	if len(pmuo.hooks) == 0 {
		node, err = pmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMappingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pmuo.mutation = mutation
			node, err = pmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pmuo.hooks) - 1; i >= 0; i-- {
			if pmuo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = pmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProductMapping)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProductMappingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *ProductMappingUpdateOne) SaveX(ctx context.Context) *ProductMapping {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *ProductMappingUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *ProductMappingUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmuo *ProductMappingUpdateOne) defaults() {
	if _, ok := pmuo.mutation.UpdatedAt(); !ok {
		v := productmapping.UpdateDefaultUpdatedAt()
		pmuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pmuo *ProductMappingUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProductMappingUpdateOne {
	pmuo.modifiers = append(pmuo.modifiers, modifiers...)
	return pmuo
}

func (pmuo *ProductMappingUpdateOne) sqlSave(ctx context.Context) (_node *ProductMapping, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productmapping.Table,
			Columns: productmapping.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: productmapping.FieldID,
			},
		},
	}
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "ProductMapping.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productmapping.FieldID)
		for _, f := range fields {
			if !productmapping.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != productmapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productmapping.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pmuo.mutation.DeletedAt(); ok {
		_spec.SetField(productmapping.FieldDeletedAt, field.TypeTime, value)
	}
	if pmuo.mutation.DeletedAtCleared() {
		_spec.ClearField(productmapping.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pmuo.mutation.Qty(); ok {
		_spec.SetField(productmapping.FieldQty, field.TypeInt, value)
	}
	if value, ok := pmuo.mutation.AddedQty(); ok {
		_spec.AddField(productmapping.FieldQty, field.TypeInt, value)
	}
	if value, ok := pmuo.mutation.CreatedBy(); ok {
		_spec.SetField(productmapping.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pmuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(productmapping.FieldCreatedBy, field.TypeInt64, value)
	}
	if pmuo.mutation.PlatformProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.PlatformProductsTable,
			Columns: []string{productmapping.PlatformProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: platformproduct.FieldID,
				},
			},
		}
		edge.Schema = pmuo.schemaConfig.ProductMapping
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.PlatformProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.PlatformProductsTable,
			Columns: []string{productmapping.PlatformProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: platformproduct.FieldID,
				},
			},
		}
		edge.Schema = pmuo.schemaConfig.ProductMapping
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pmuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.ProductsTable,
			Columns: []string{productmapping.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: product.FieldID,
				},
			},
		}
		edge.Schema = pmuo.schemaConfig.ProductMapping
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmapping.ProductsTable,
			Columns: []string{productmapping.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: product.FieldID,
				},
			},
		}
		edge.Schema = pmuo.schemaConfig.ProductMapping
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = pmuo.schemaConfig.ProductMapping
	ctx = internal.NewSchemaConfigContext(ctx, pmuo.schemaConfig)
	_spec.AddModifiers(pmuo.modifiers...)
	_node = &ProductMapping{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
