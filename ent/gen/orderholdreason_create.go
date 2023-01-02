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
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/orderholdreason"
)

// OrderHoldReasonCreate is the builder for creating a OrderHoldReason entity.
type OrderHoldReasonCreate struct {
	config
	mutation *OrderHoldReasonMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrderID sets the "order_id" field.
func (ohrc *OrderHoldReasonCreate) SetOrderID(i int64) *OrderHoldReasonCreate {
	ohrc.mutation.SetOrderID(i)
	return ohrc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (ohrc *OrderHoldReasonCreate) SetNillableOrderID(i *int64) *OrderHoldReasonCreate {
	if i != nil {
		ohrc.SetOrderID(*i)
	}
	return ohrc
}

// SetProductID sets the "product_id" field.
func (ohrc *OrderHoldReasonCreate) SetProductID(i int64) *OrderHoldReasonCreate {
	ohrc.mutation.SetProductID(i)
	return ohrc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (ohrc *OrderHoldReasonCreate) SetNillableProductID(i *int64) *OrderHoldReasonCreate {
	if i != nil {
		ohrc.SetProductID(*i)
	}
	return ohrc
}

// SetType sets the "type" field.
func (ohrc *OrderHoldReasonCreate) SetType(s string) *OrderHoldReasonCreate {
	ohrc.mutation.SetType(s)
	return ohrc
}

// SetCode sets the "code" field.
func (ohrc *OrderHoldReasonCreate) SetCode(i int32) *OrderHoldReasonCreate {
	ohrc.mutation.SetCode(i)
	return ohrc
}

// SetReason sets the "reason" field.
func (ohrc *OrderHoldReasonCreate) SetReason(s string) *OrderHoldReasonCreate {
	ohrc.mutation.SetReason(s)
	return ohrc
}

// SetEnReason sets the "en_reason" field.
func (ohrc *OrderHoldReasonCreate) SetEnReason(s string) *OrderHoldReasonCreate {
	ohrc.mutation.SetEnReason(s)
	return ohrc
}

// SetCreatedAt sets the "created_at" field.
func (ohrc *OrderHoldReasonCreate) SetCreatedAt(t time.Time) *OrderHoldReasonCreate {
	ohrc.mutation.SetCreatedAt(t)
	return ohrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ohrc *OrderHoldReasonCreate) SetNillableCreatedAt(t *time.Time) *OrderHoldReasonCreate {
	if t != nil {
		ohrc.SetCreatedAt(*t)
	}
	return ohrc
}

// SetID sets the "id" field.
func (ohrc *OrderHoldReasonCreate) SetID(i int64) *OrderHoldReasonCreate {
	ohrc.mutation.SetID(i)
	return ohrc
}

// SetOrdersID sets the "orders" edge to the Order entity by ID.
func (ohrc *OrderHoldReasonCreate) SetOrdersID(id int64) *OrderHoldReasonCreate {
	ohrc.mutation.SetOrdersID(id)
	return ohrc
}

// SetNillableOrdersID sets the "orders" edge to the Order entity by ID if the given value is not nil.
func (ohrc *OrderHoldReasonCreate) SetNillableOrdersID(id *int64) *OrderHoldReasonCreate {
	if id != nil {
		ohrc = ohrc.SetOrdersID(*id)
	}
	return ohrc
}

// SetOrders sets the "orders" edge to the Order entity.
func (ohrc *OrderHoldReasonCreate) SetOrders(o *Order) *OrderHoldReasonCreate {
	return ohrc.SetOrdersID(o.ID)
}

// Mutation returns the OrderHoldReasonMutation object of the builder.
func (ohrc *OrderHoldReasonCreate) Mutation() *OrderHoldReasonMutation {
	return ohrc.mutation
}

// Save creates the OrderHoldReason in the database.
func (ohrc *OrderHoldReasonCreate) Save(ctx context.Context) (*OrderHoldReason, error) {
	var (
		err  error
		node *OrderHoldReason
	)
	ohrc.defaults()
	if len(ohrc.hooks) == 0 {
		if err = ohrc.check(); err != nil {
			return nil, err
		}
		node, err = ohrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderHoldReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ohrc.check(); err != nil {
				return nil, err
			}
			ohrc.mutation = mutation
			if node, err = ohrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ohrc.hooks) - 1; i >= 0; i-- {
			if ohrc.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = ohrc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ohrc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderHoldReason)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderHoldReasonMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ohrc *OrderHoldReasonCreate) SaveX(ctx context.Context) *OrderHoldReason {
	v, err := ohrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ohrc *OrderHoldReasonCreate) Exec(ctx context.Context) error {
	_, err := ohrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohrc *OrderHoldReasonCreate) ExecX(ctx context.Context) {
	if err := ohrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ohrc *OrderHoldReasonCreate) defaults() {
	if _, ok := ohrc.mutation.CreatedAt(); !ok {
		v := orderholdreason.DefaultCreatedAt()
		ohrc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ohrc *OrderHoldReasonCreate) check() error {
	if _, ok := ohrc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`gen: missing required field "OrderHoldReason.type"`)}
	}
	if _, ok := ohrc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`gen: missing required field "OrderHoldReason.code"`)}
	}
	if _, ok := ohrc.mutation.Reason(); !ok {
		return &ValidationError{Name: "reason", err: errors.New(`gen: missing required field "OrderHoldReason.reason"`)}
	}
	if _, ok := ohrc.mutation.EnReason(); !ok {
		return &ValidationError{Name: "en_reason", err: errors.New(`gen: missing required field "OrderHoldReason.en_reason"`)}
	}
	if _, ok := ohrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`gen: missing required field "OrderHoldReason.created_at"`)}
	}
	return nil
}

func (ohrc *OrderHoldReasonCreate) sqlSave(ctx context.Context) (*OrderHoldReason, error) {
	_node, _spec := ohrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ohrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (ohrc *OrderHoldReasonCreate) createSpec() (*OrderHoldReason, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderHoldReason{config: ohrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderholdreason.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderholdreason.FieldID,
			},
		}
	)
	_spec.Schema = ohrc.schemaConfig.OrderHoldReason
	_spec.OnConflict = ohrc.conflict
	if id, ok := ohrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ohrc.mutation.ProductID(); ok {
		_spec.SetField(orderholdreason.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := ohrc.mutation.GetType(); ok {
		_spec.SetField(orderholdreason.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ohrc.mutation.Code(); ok {
		_spec.SetField(orderholdreason.FieldCode, field.TypeInt32, value)
		_node.Code = value
	}
	if value, ok := ohrc.mutation.Reason(); ok {
		_spec.SetField(orderholdreason.FieldReason, field.TypeString, value)
		_node.Reason = value
	}
	if value, ok := ohrc.mutation.EnReason(); ok {
		_spec.SetField(orderholdreason.FieldEnReason, field.TypeString, value)
		_node.EnReason = value
	}
	if value, ok := ohrc.mutation.CreatedAt(); ok {
		_spec.SetField(orderholdreason.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ohrc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderholdreason.OrdersTable,
			Columns: []string{orderholdreason.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: order.FieldID,
				},
			},
		}
		edge.Schema = ohrc.schemaConfig.OrderHoldReason
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderHoldReason.Create().
//		SetOrderID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderHoldReasonUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
func (ohrc *OrderHoldReasonCreate) OnConflict(opts ...sql.ConflictOption) *OrderHoldReasonUpsertOne {
	ohrc.conflict = opts
	return &OrderHoldReasonUpsertOne{
		create: ohrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderHoldReason.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ohrc *OrderHoldReasonCreate) OnConflictColumns(columns ...string) *OrderHoldReasonUpsertOne {
	ohrc.conflict = append(ohrc.conflict, sql.ConflictColumns(columns...))
	return &OrderHoldReasonUpsertOne{
		create: ohrc,
	}
}

type (
	// OrderHoldReasonUpsertOne is the builder for "upsert"-ing
	//  one OrderHoldReason node.
	OrderHoldReasonUpsertOne struct {
		create *OrderHoldReasonCreate
	}

	// OrderHoldReasonUpsert is the "OnConflict" setter.
	OrderHoldReasonUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrderID sets the "order_id" field.
func (u *OrderHoldReasonUpsert) SetOrderID(v int64) *OrderHoldReasonUpsert {
	u.Set(orderholdreason.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderHoldReasonUpsert) UpdateOrderID() *OrderHoldReasonUpsert {
	u.SetExcluded(orderholdreason.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *OrderHoldReasonUpsert) ClearOrderID() *OrderHoldReasonUpsert {
	u.SetNull(orderholdreason.FieldOrderID)
	return u
}

// SetProductID sets the "product_id" field.
func (u *OrderHoldReasonUpsert) SetProductID(v int64) *OrderHoldReasonUpsert {
	u.Set(orderholdreason.FieldProductID, v)
	return u
}

// UpdateProductID sets the "product_id" field to the value that was provided on create.
func (u *OrderHoldReasonUpsert) UpdateProductID() *OrderHoldReasonUpsert {
	u.SetExcluded(orderholdreason.FieldProductID)
	return u
}

// AddProductID adds v to the "product_id" field.
func (u *OrderHoldReasonUpsert) AddProductID(v int64) *OrderHoldReasonUpsert {
	u.Add(orderholdreason.FieldProductID, v)
	return u
}

// ClearProductID clears the value of the "product_id" field.
func (u *OrderHoldReasonUpsert) ClearProductID() *OrderHoldReasonUpsert {
	u.SetNull(orderholdreason.FieldProductID)
	return u
}

// SetType sets the "type" field.
func (u *OrderHoldReasonUpsert) SetType(v string) *OrderHoldReasonUpsert {
	u.Set(orderholdreason.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OrderHoldReasonUpsert) UpdateType() *OrderHoldReasonUpsert {
	u.SetExcluded(orderholdreason.FieldType)
	return u
}

// SetCode sets the "code" field.
func (u *OrderHoldReasonUpsert) SetCode(v int32) *OrderHoldReasonUpsert {
	u.Set(orderholdreason.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *OrderHoldReasonUpsert) UpdateCode() *OrderHoldReasonUpsert {
	u.SetExcluded(orderholdreason.FieldCode)
	return u
}

// AddCode adds v to the "code" field.
func (u *OrderHoldReasonUpsert) AddCode(v int32) *OrderHoldReasonUpsert {
	u.Add(orderholdreason.FieldCode, v)
	return u
}

// SetReason sets the "reason" field.
func (u *OrderHoldReasonUpsert) SetReason(v string) *OrderHoldReasonUpsert {
	u.Set(orderholdreason.FieldReason, v)
	return u
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *OrderHoldReasonUpsert) UpdateReason() *OrderHoldReasonUpsert {
	u.SetExcluded(orderholdreason.FieldReason)
	return u
}

// SetEnReason sets the "en_reason" field.
func (u *OrderHoldReasonUpsert) SetEnReason(v string) *OrderHoldReasonUpsert {
	u.Set(orderholdreason.FieldEnReason, v)
	return u
}

// UpdateEnReason sets the "en_reason" field to the value that was provided on create.
func (u *OrderHoldReasonUpsert) UpdateEnReason() *OrderHoldReasonUpsert {
	u.SetExcluded(orderholdreason.FieldEnReason)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderHoldReason.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderholdreason.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderHoldReasonUpsertOne) UpdateNewValues() *OrderHoldReasonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orderholdreason.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(orderholdreason.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderHoldReason.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderHoldReasonUpsertOne) Ignore() *OrderHoldReasonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderHoldReasonUpsertOne) DoNothing() *OrderHoldReasonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderHoldReasonCreate.OnConflict
// documentation for more info.
func (u *OrderHoldReasonUpsertOne) Update(set func(*OrderHoldReasonUpsert)) *OrderHoldReasonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderHoldReasonUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderHoldReasonUpsertOne) SetOrderID(v int64) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertOne) UpdateOrderID() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *OrderHoldReasonUpsertOne) ClearOrderID() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.ClearOrderID()
	})
}

// SetProductID sets the "product_id" field.
func (u *OrderHoldReasonUpsertOne) SetProductID(v int64) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetProductID(v)
	})
}

// AddProductID adds v to the "product_id" field.
func (u *OrderHoldReasonUpsertOne) AddProductID(v int64) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.AddProductID(v)
	})
}

// UpdateProductID sets the "product_id" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertOne) UpdateProductID() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateProductID()
	})
}

// ClearProductID clears the value of the "product_id" field.
func (u *OrderHoldReasonUpsertOne) ClearProductID() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.ClearProductID()
	})
}

// SetType sets the "type" field.
func (u *OrderHoldReasonUpsertOne) SetType(v string) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertOne) UpdateType() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateType()
	})
}

// SetCode sets the "code" field.
func (u *OrderHoldReasonUpsertOne) SetCode(v int32) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetCode(v)
	})
}

// AddCode adds v to the "code" field.
func (u *OrderHoldReasonUpsertOne) AddCode(v int32) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.AddCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertOne) UpdateCode() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateCode()
	})
}

// SetReason sets the "reason" field.
func (u *OrderHoldReasonUpsertOne) SetReason(v string) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertOne) UpdateReason() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateReason()
	})
}

// SetEnReason sets the "en_reason" field.
func (u *OrderHoldReasonUpsertOne) SetEnReason(v string) *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetEnReason(v)
	})
}

// UpdateEnReason sets the "en_reason" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertOne) UpdateEnReason() *OrderHoldReasonUpsertOne {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateEnReason()
	})
}

// Exec executes the query.
func (u *OrderHoldReasonUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for OrderHoldReasonCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderHoldReasonUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderHoldReasonUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderHoldReasonUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderHoldReasonCreateBulk is the builder for creating many OrderHoldReason entities in bulk.
type OrderHoldReasonCreateBulk struct {
	config
	builders []*OrderHoldReasonCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderHoldReason entities in the database.
func (ohrcb *OrderHoldReasonCreateBulk) Save(ctx context.Context) ([]*OrderHoldReason, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ohrcb.builders))
	nodes := make([]*OrderHoldReason, len(ohrcb.builders))
	mutators := make([]Mutator, len(ohrcb.builders))
	for i := range ohrcb.builders {
		func(i int, root context.Context) {
			builder := ohrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderHoldReasonMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ohrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ohrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ohrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ohrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ohrcb *OrderHoldReasonCreateBulk) SaveX(ctx context.Context) []*OrderHoldReason {
	v, err := ohrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ohrcb *OrderHoldReasonCreateBulk) Exec(ctx context.Context) error {
	_, err := ohrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohrcb *OrderHoldReasonCreateBulk) ExecX(ctx context.Context) {
	if err := ohrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderHoldReason.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderHoldReasonUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
func (ohrcb *OrderHoldReasonCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderHoldReasonUpsertBulk {
	ohrcb.conflict = opts
	return &OrderHoldReasonUpsertBulk{
		create: ohrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderHoldReason.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ohrcb *OrderHoldReasonCreateBulk) OnConflictColumns(columns ...string) *OrderHoldReasonUpsertBulk {
	ohrcb.conflict = append(ohrcb.conflict, sql.ConflictColumns(columns...))
	return &OrderHoldReasonUpsertBulk{
		create: ohrcb,
	}
}

// OrderHoldReasonUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderHoldReason nodes.
type OrderHoldReasonUpsertBulk struct {
	create *OrderHoldReasonCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderHoldReason.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderholdreason.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderHoldReasonUpsertBulk) UpdateNewValues() *OrderHoldReasonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orderholdreason.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(orderholdreason.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderHoldReason.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderHoldReasonUpsertBulk) Ignore() *OrderHoldReasonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderHoldReasonUpsertBulk) DoNothing() *OrderHoldReasonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderHoldReasonCreateBulk.OnConflict
// documentation for more info.
func (u *OrderHoldReasonUpsertBulk) Update(set func(*OrderHoldReasonUpsert)) *OrderHoldReasonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderHoldReasonUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderHoldReasonUpsertBulk) SetOrderID(v int64) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertBulk) UpdateOrderID() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *OrderHoldReasonUpsertBulk) ClearOrderID() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.ClearOrderID()
	})
}

// SetProductID sets the "product_id" field.
func (u *OrderHoldReasonUpsertBulk) SetProductID(v int64) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetProductID(v)
	})
}

// AddProductID adds v to the "product_id" field.
func (u *OrderHoldReasonUpsertBulk) AddProductID(v int64) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.AddProductID(v)
	})
}

// UpdateProductID sets the "product_id" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertBulk) UpdateProductID() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateProductID()
	})
}

// ClearProductID clears the value of the "product_id" field.
func (u *OrderHoldReasonUpsertBulk) ClearProductID() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.ClearProductID()
	})
}

// SetType sets the "type" field.
func (u *OrderHoldReasonUpsertBulk) SetType(v string) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertBulk) UpdateType() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateType()
	})
}

// SetCode sets the "code" field.
func (u *OrderHoldReasonUpsertBulk) SetCode(v int32) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetCode(v)
	})
}

// AddCode adds v to the "code" field.
func (u *OrderHoldReasonUpsertBulk) AddCode(v int32) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.AddCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertBulk) UpdateCode() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateCode()
	})
}

// SetReason sets the "reason" field.
func (u *OrderHoldReasonUpsertBulk) SetReason(v string) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertBulk) UpdateReason() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateReason()
	})
}

// SetEnReason sets the "en_reason" field.
func (u *OrderHoldReasonUpsertBulk) SetEnReason(v string) *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.SetEnReason(v)
	})
}

// UpdateEnReason sets the "en_reason" field to the value that was provided on create.
func (u *OrderHoldReasonUpsertBulk) UpdateEnReason() *OrderHoldReasonUpsertBulk {
	return u.Update(func(s *OrderHoldReasonUpsert) {
		s.UpdateEnReason()
	})
}

// Exec executes the query.
func (u *OrderHoldReasonUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("gen: OnConflict was set for builder %d. Set it on the OrderHoldReasonCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for OrderHoldReasonCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderHoldReasonUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
