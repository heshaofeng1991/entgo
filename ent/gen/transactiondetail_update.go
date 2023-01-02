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
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/transactiondetail"
)

// TransactionDetailUpdate is the builder for updating TransactionDetail entities.
type TransactionDetailUpdate struct {
	config
	hooks     []Hook
	mutation  *TransactionDetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TransactionDetailUpdate builder.
func (tdu *TransactionDetailUpdate) Where(ps ...predicate.TransactionDetail) *TransactionDetailUpdate {
	tdu.mutation.Where(ps...)
	return tdu
}

// SetUpdatedAt sets the "updated_at" field.
func (tdu *TransactionDetailUpdate) SetUpdatedAt(t time.Time) *TransactionDetailUpdate {
	tdu.mutation.SetUpdatedAt(t)
	return tdu
}

// SetDeletedAt sets the "deleted_at" field.
func (tdu *TransactionDetailUpdate) SetDeletedAt(t time.Time) *TransactionDetailUpdate {
	tdu.mutation.SetDeletedAt(t)
	return tdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableDeletedAt(t *time.Time) *TransactionDetailUpdate {
	if t != nil {
		tdu.SetDeletedAt(*t)
	}
	return tdu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tdu *TransactionDetailUpdate) ClearDeletedAt() *TransactionDetailUpdate {
	tdu.mutation.ClearDeletedAt()
	return tdu
}

// SetOrderID sets the "order_id" field.
func (tdu *TransactionDetailUpdate) SetOrderID(i int64) *TransactionDetailUpdate {
	tdu.mutation.ResetOrderID()
	tdu.mutation.SetOrderID(i)
	return tdu
}

// AddOrderID adds i to the "order_id" field.
func (tdu *TransactionDetailUpdate) AddOrderID(i int64) *TransactionDetailUpdate {
	tdu.mutation.AddOrderID(i)
	return tdu
}

// SetTransactionID sets the "transaction_id" field.
func (tdu *TransactionDetailUpdate) SetTransactionID(i int64) *TransactionDetailUpdate {
	tdu.mutation.ResetTransactionID()
	tdu.mutation.SetTransactionID(i)
	return tdu
}

// AddTransactionID adds i to the "transaction_id" field.
func (tdu *TransactionDetailUpdate) AddTransactionID(i int64) *TransactionDetailUpdate {
	tdu.mutation.AddTransactionID(i)
	return tdu
}

// SetTransactionType sets the "transaction_type" field.
func (tdu *TransactionDetailUpdate) SetTransactionType(s string) *TransactionDetailUpdate {
	tdu.mutation.SetTransactionType(s)
	return tdu
}

// SetNillableTransactionType sets the "transaction_type" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableTransactionType(s *string) *TransactionDetailUpdate {
	if s != nil {
		tdu.SetTransactionType(*s)
	}
	return tdu
}

// SetDeliveryCost sets the "delivery_cost" field.
func (tdu *TransactionDetailUpdate) SetDeliveryCost(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetDeliveryCost()
	tdu.mutation.SetDeliveryCost(f)
	return tdu
}

// SetNillableDeliveryCost sets the "delivery_cost" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableDeliveryCost(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetDeliveryCost(*f)
	}
	return tdu
}

// AddDeliveryCost adds f to the "delivery_cost" field.
func (tdu *TransactionDetailUpdate) AddDeliveryCost(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddDeliveryCost(f)
	return tdu
}

// SetMiscFee sets the "misc_fee" field.
func (tdu *TransactionDetailUpdate) SetMiscFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetMiscFee()
	tdu.mutation.SetMiscFee(f)
	return tdu
}

// SetNillableMiscFee sets the "misc_fee" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableMiscFee(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetMiscFee(*f)
	}
	return tdu
}

// AddMiscFee adds f to the "misc_fee" field.
func (tdu *TransactionDetailUpdate) AddMiscFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddMiscFee(f)
	return tdu
}

// SetFuelFee sets the "fuel_fee" field.
func (tdu *TransactionDetailUpdate) SetFuelFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetFuelFee()
	tdu.mutation.SetFuelFee(f)
	return tdu
}

// SetNillableFuelFee sets the "fuel_fee" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableFuelFee(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetFuelFee(*f)
	}
	return tdu
}

// AddFuelFee adds f to the "fuel_fee" field.
func (tdu *TransactionDetailUpdate) AddFuelFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddFuelFee(f)
	return tdu
}

// SetRegistrationFee sets the "registration_fee" field.
func (tdu *TransactionDetailUpdate) SetRegistrationFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetRegistrationFee()
	tdu.mutation.SetRegistrationFee(f)
	return tdu
}

// SetNillableRegistrationFee sets the "registration_fee" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableRegistrationFee(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetRegistrationFee(*f)
	}
	return tdu
}

// AddRegistrationFee adds f to the "registration_fee" field.
func (tdu *TransactionDetailUpdate) AddRegistrationFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddRegistrationFee(f)
	return tdu
}

// SetProcessingFee sets the "processing_fee" field.
func (tdu *TransactionDetailUpdate) SetProcessingFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetProcessingFee()
	tdu.mutation.SetProcessingFee(f)
	return tdu
}

// SetNillableProcessingFee sets the "processing_fee" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableProcessingFee(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetProcessingFee(*f)
	}
	return tdu
}

// AddProcessingFee adds f to the "processing_fee" field.
func (tdu *TransactionDetailUpdate) AddProcessingFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddProcessingFee(f)
	return tdu
}

// SetPackageFee sets the "package_fee" field.
func (tdu *TransactionDetailUpdate) SetPackageFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetPackageFee()
	tdu.mutation.SetPackageFee(f)
	return tdu
}

// SetNillablePackageFee sets the "package_fee" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillablePackageFee(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetPackageFee(*f)
	}
	return tdu
}

// AddPackageFee adds f to the "package_fee" field.
func (tdu *TransactionDetailUpdate) AddPackageFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddPackageFee(f)
	return tdu
}

// SetHandlingFee sets the "handling_fee" field.
func (tdu *TransactionDetailUpdate) SetHandlingFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetHandlingFee()
	tdu.mutation.SetHandlingFee(f)
	return tdu
}

// SetNillableHandlingFee sets the "handling_fee" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableHandlingFee(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetHandlingFee(*f)
	}
	return tdu
}

// AddHandlingFee adds f to the "handling_fee" field.
func (tdu *TransactionDetailUpdate) AddHandlingFee(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddHandlingFee(f)
	return tdu
}

// SetVat sets the "vat" field.
func (tdu *TransactionDetailUpdate) SetVat(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetVat()
	tdu.mutation.SetVat(f)
	return tdu
}

// SetNillableVat sets the "vat" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableVat(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetVat(*f)
	}
	return tdu
}

// AddVat adds f to the "vat" field.
func (tdu *TransactionDetailUpdate) AddVat(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddVat(f)
	return tdu
}

// SetAmount sets the "amount" field.
func (tdu *TransactionDetailUpdate) SetAmount(f float64) *TransactionDetailUpdate {
	tdu.mutation.ResetAmount()
	tdu.mutation.SetAmount(f)
	return tdu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableAmount(f *float64) *TransactionDetailUpdate {
	if f != nil {
		tdu.SetAmount(*f)
	}
	return tdu
}

// AddAmount adds f to the "amount" field.
func (tdu *TransactionDetailUpdate) AddAmount(f float64) *TransactionDetailUpdate {
	tdu.mutation.AddAmount(f)
	return tdu
}

// SetWeight sets the "weight" field.
func (tdu *TransactionDetailUpdate) SetWeight(i int) *TransactionDetailUpdate {
	tdu.mutation.ResetWeight()
	tdu.mutation.SetWeight(i)
	return tdu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (tdu *TransactionDetailUpdate) SetNillableWeight(i *int) *TransactionDetailUpdate {
	if i != nil {
		tdu.SetWeight(*i)
	}
	return tdu
}

// AddWeight adds i to the "weight" field.
func (tdu *TransactionDetailUpdate) AddWeight(i int) *TransactionDetailUpdate {
	tdu.mutation.AddWeight(i)
	return tdu
}

// Mutation returns the TransactionDetailMutation object of the builder.
func (tdu *TransactionDetailUpdate) Mutation() *TransactionDetailMutation {
	return tdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tdu *TransactionDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tdu.defaults()
	if len(tdu.hooks) == 0 {
		affected, err = tdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tdu.mutation = mutation
			affected, err = tdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tdu.hooks) - 1; i >= 0; i-- {
			if tdu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = tdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tdu *TransactionDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := tdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tdu *TransactionDetailUpdate) Exec(ctx context.Context) error {
	_, err := tdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdu *TransactionDetailUpdate) ExecX(ctx context.Context) {
	if err := tdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdu *TransactionDetailUpdate) defaults() {
	if _, ok := tdu.mutation.UpdatedAt(); !ok {
		v := transactiondetail.UpdateDefaultUpdatedAt()
		tdu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tdu *TransactionDetailUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TransactionDetailUpdate {
	tdu.modifiers = append(tdu.modifiers, modifiers...)
	return tdu
}

func (tdu *TransactionDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiondetail.Table,
			Columns: transactiondetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: transactiondetail.FieldID,
			},
		},
	}
	if ps := tdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdu.mutation.UpdatedAt(); ok {
		_spec.SetField(transactiondetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.DeletedAt(); ok {
		_spec.SetField(transactiondetail.FieldDeletedAt, field.TypeTime, value)
	}
	if tdu.mutation.DeletedAtCleared() {
		_spec.ClearField(transactiondetail.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tdu.mutation.OrderID(); ok {
		_spec.SetField(transactiondetail.FieldOrderID, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.AddedOrderID(); ok {
		_spec.AddField(transactiondetail.FieldOrderID, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.TransactionID(); ok {
		_spec.SetField(transactiondetail.FieldTransactionID, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.AddedTransactionID(); ok {
		_spec.AddField(transactiondetail.FieldTransactionID, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.TransactionType(); ok {
		_spec.SetField(transactiondetail.FieldTransactionType, field.TypeString, value)
	}
	if value, ok := tdu.mutation.DeliveryCost(); ok {
		_spec.SetField(transactiondetail.FieldDeliveryCost, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedDeliveryCost(); ok {
		_spec.AddField(transactiondetail.FieldDeliveryCost, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.MiscFee(); ok {
		_spec.SetField(transactiondetail.FieldMiscFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedMiscFee(); ok {
		_spec.AddField(transactiondetail.FieldMiscFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.FuelFee(); ok {
		_spec.SetField(transactiondetail.FieldFuelFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedFuelFee(); ok {
		_spec.AddField(transactiondetail.FieldFuelFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.RegistrationFee(); ok {
		_spec.SetField(transactiondetail.FieldRegistrationFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedRegistrationFee(); ok {
		_spec.AddField(transactiondetail.FieldRegistrationFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.ProcessingFee(); ok {
		_spec.SetField(transactiondetail.FieldProcessingFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedProcessingFee(); ok {
		_spec.AddField(transactiondetail.FieldProcessingFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.PackageFee(); ok {
		_spec.SetField(transactiondetail.FieldPackageFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedPackageFee(); ok {
		_spec.AddField(transactiondetail.FieldPackageFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.HandlingFee(); ok {
		_spec.SetField(transactiondetail.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedHandlingFee(); ok {
		_spec.AddField(transactiondetail.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.Vat(); ok {
		_spec.SetField(transactiondetail.FieldVat, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedVat(); ok {
		_spec.AddField(transactiondetail.FieldVat, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.Amount(); ok {
		_spec.SetField(transactiondetail.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedAmount(); ok {
		_spec.AddField(transactiondetail.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.Weight(); ok {
		_spec.SetField(transactiondetail.FieldWeight, field.TypeInt, value)
	}
	if value, ok := tdu.mutation.AddedWeight(); ok {
		_spec.AddField(transactiondetail.FieldWeight, field.TypeInt, value)
	}
	_spec.Node.Schema = tdu.schemaConfig.TransactionDetail
	ctx = internal.NewSchemaConfigContext(ctx, tdu.schemaConfig)
	_spec.AddModifiers(tdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiondetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TransactionDetailUpdateOne is the builder for updating a single TransactionDetail entity.
type TransactionDetailUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TransactionDetailMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tduo *TransactionDetailUpdateOne) SetUpdatedAt(t time.Time) *TransactionDetailUpdateOne {
	tduo.mutation.SetUpdatedAt(t)
	return tduo
}

// SetDeletedAt sets the "deleted_at" field.
func (tduo *TransactionDetailUpdateOne) SetDeletedAt(t time.Time) *TransactionDetailUpdateOne {
	tduo.mutation.SetDeletedAt(t)
	return tduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableDeletedAt(t *time.Time) *TransactionDetailUpdateOne {
	if t != nil {
		tduo.SetDeletedAt(*t)
	}
	return tduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tduo *TransactionDetailUpdateOne) ClearDeletedAt() *TransactionDetailUpdateOne {
	tduo.mutation.ClearDeletedAt()
	return tduo
}

// SetOrderID sets the "order_id" field.
func (tduo *TransactionDetailUpdateOne) SetOrderID(i int64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetOrderID()
	tduo.mutation.SetOrderID(i)
	return tduo
}

// AddOrderID adds i to the "order_id" field.
func (tduo *TransactionDetailUpdateOne) AddOrderID(i int64) *TransactionDetailUpdateOne {
	tduo.mutation.AddOrderID(i)
	return tduo
}

// SetTransactionID sets the "transaction_id" field.
func (tduo *TransactionDetailUpdateOne) SetTransactionID(i int64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetTransactionID()
	tduo.mutation.SetTransactionID(i)
	return tduo
}

// AddTransactionID adds i to the "transaction_id" field.
func (tduo *TransactionDetailUpdateOne) AddTransactionID(i int64) *TransactionDetailUpdateOne {
	tduo.mutation.AddTransactionID(i)
	return tduo
}

// SetTransactionType sets the "transaction_type" field.
func (tduo *TransactionDetailUpdateOne) SetTransactionType(s string) *TransactionDetailUpdateOne {
	tduo.mutation.SetTransactionType(s)
	return tduo
}

// SetNillableTransactionType sets the "transaction_type" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableTransactionType(s *string) *TransactionDetailUpdateOne {
	if s != nil {
		tduo.SetTransactionType(*s)
	}
	return tduo
}

// SetDeliveryCost sets the "delivery_cost" field.
func (tduo *TransactionDetailUpdateOne) SetDeliveryCost(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetDeliveryCost()
	tduo.mutation.SetDeliveryCost(f)
	return tduo
}

// SetNillableDeliveryCost sets the "delivery_cost" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableDeliveryCost(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetDeliveryCost(*f)
	}
	return tduo
}

// AddDeliveryCost adds f to the "delivery_cost" field.
func (tduo *TransactionDetailUpdateOne) AddDeliveryCost(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddDeliveryCost(f)
	return tduo
}

// SetMiscFee sets the "misc_fee" field.
func (tduo *TransactionDetailUpdateOne) SetMiscFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetMiscFee()
	tduo.mutation.SetMiscFee(f)
	return tduo
}

// SetNillableMiscFee sets the "misc_fee" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableMiscFee(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetMiscFee(*f)
	}
	return tduo
}

// AddMiscFee adds f to the "misc_fee" field.
func (tduo *TransactionDetailUpdateOne) AddMiscFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddMiscFee(f)
	return tduo
}

// SetFuelFee sets the "fuel_fee" field.
func (tduo *TransactionDetailUpdateOne) SetFuelFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetFuelFee()
	tduo.mutation.SetFuelFee(f)
	return tduo
}

// SetNillableFuelFee sets the "fuel_fee" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableFuelFee(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetFuelFee(*f)
	}
	return tduo
}

// AddFuelFee adds f to the "fuel_fee" field.
func (tduo *TransactionDetailUpdateOne) AddFuelFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddFuelFee(f)
	return tduo
}

// SetRegistrationFee sets the "registration_fee" field.
func (tduo *TransactionDetailUpdateOne) SetRegistrationFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetRegistrationFee()
	tduo.mutation.SetRegistrationFee(f)
	return tduo
}

// SetNillableRegistrationFee sets the "registration_fee" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableRegistrationFee(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetRegistrationFee(*f)
	}
	return tduo
}

// AddRegistrationFee adds f to the "registration_fee" field.
func (tduo *TransactionDetailUpdateOne) AddRegistrationFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddRegistrationFee(f)
	return tduo
}

// SetProcessingFee sets the "processing_fee" field.
func (tduo *TransactionDetailUpdateOne) SetProcessingFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetProcessingFee()
	tduo.mutation.SetProcessingFee(f)
	return tduo
}

// SetNillableProcessingFee sets the "processing_fee" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableProcessingFee(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetProcessingFee(*f)
	}
	return tduo
}

// AddProcessingFee adds f to the "processing_fee" field.
func (tduo *TransactionDetailUpdateOne) AddProcessingFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddProcessingFee(f)
	return tduo
}

// SetPackageFee sets the "package_fee" field.
func (tduo *TransactionDetailUpdateOne) SetPackageFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetPackageFee()
	tduo.mutation.SetPackageFee(f)
	return tduo
}

// SetNillablePackageFee sets the "package_fee" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillablePackageFee(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetPackageFee(*f)
	}
	return tduo
}

// AddPackageFee adds f to the "package_fee" field.
func (tduo *TransactionDetailUpdateOne) AddPackageFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddPackageFee(f)
	return tduo
}

// SetHandlingFee sets the "handling_fee" field.
func (tduo *TransactionDetailUpdateOne) SetHandlingFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetHandlingFee()
	tduo.mutation.SetHandlingFee(f)
	return tduo
}

// SetNillableHandlingFee sets the "handling_fee" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableHandlingFee(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetHandlingFee(*f)
	}
	return tduo
}

// AddHandlingFee adds f to the "handling_fee" field.
func (tduo *TransactionDetailUpdateOne) AddHandlingFee(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddHandlingFee(f)
	return tduo
}

// SetVat sets the "vat" field.
func (tduo *TransactionDetailUpdateOne) SetVat(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetVat()
	tduo.mutation.SetVat(f)
	return tduo
}

// SetNillableVat sets the "vat" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableVat(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetVat(*f)
	}
	return tduo
}

// AddVat adds f to the "vat" field.
func (tduo *TransactionDetailUpdateOne) AddVat(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddVat(f)
	return tduo
}

// SetAmount sets the "amount" field.
func (tduo *TransactionDetailUpdateOne) SetAmount(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.ResetAmount()
	tduo.mutation.SetAmount(f)
	return tduo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableAmount(f *float64) *TransactionDetailUpdateOne {
	if f != nil {
		tduo.SetAmount(*f)
	}
	return tduo
}

// AddAmount adds f to the "amount" field.
func (tduo *TransactionDetailUpdateOne) AddAmount(f float64) *TransactionDetailUpdateOne {
	tduo.mutation.AddAmount(f)
	return tduo
}

// SetWeight sets the "weight" field.
func (tduo *TransactionDetailUpdateOne) SetWeight(i int) *TransactionDetailUpdateOne {
	tduo.mutation.ResetWeight()
	tduo.mutation.SetWeight(i)
	return tduo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (tduo *TransactionDetailUpdateOne) SetNillableWeight(i *int) *TransactionDetailUpdateOne {
	if i != nil {
		tduo.SetWeight(*i)
	}
	return tduo
}

// AddWeight adds i to the "weight" field.
func (tduo *TransactionDetailUpdateOne) AddWeight(i int) *TransactionDetailUpdateOne {
	tduo.mutation.AddWeight(i)
	return tduo
}

// Mutation returns the TransactionDetailMutation object of the builder.
func (tduo *TransactionDetailUpdateOne) Mutation() *TransactionDetailMutation {
	return tduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tduo *TransactionDetailUpdateOne) Select(field string, fields ...string) *TransactionDetailUpdateOne {
	tduo.fields = append([]string{field}, fields...)
	return tduo
}

// Save executes the query and returns the updated TransactionDetail entity.
func (tduo *TransactionDetailUpdateOne) Save(ctx context.Context) (*TransactionDetail, error) {
	var (
		err  error
		node *TransactionDetail
	)
	tduo.defaults()
	if len(tduo.hooks) == 0 {
		node, err = tduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tduo.mutation = mutation
			node, err = tduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tduo.hooks) - 1; i >= 0; i-- {
			if tduo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = tduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TransactionDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransactionDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tduo *TransactionDetailUpdateOne) SaveX(ctx context.Context) *TransactionDetail {
	node, err := tduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tduo *TransactionDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := tduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tduo *TransactionDetailUpdateOne) ExecX(ctx context.Context) {
	if err := tduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tduo *TransactionDetailUpdateOne) defaults() {
	if _, ok := tduo.mutation.UpdatedAt(); !ok {
		v := transactiondetail.UpdateDefaultUpdatedAt()
		tduo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tduo *TransactionDetailUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TransactionDetailUpdateOne {
	tduo.modifiers = append(tduo.modifiers, modifiers...)
	return tduo
}

func (tduo *TransactionDetailUpdateOne) sqlSave(ctx context.Context) (_node *TransactionDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactiondetail.Table,
			Columns: transactiondetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: transactiondetail.FieldID,
			},
		},
	}
	id, ok := tduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "TransactionDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactiondetail.FieldID)
		for _, f := range fields {
			if !transactiondetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != transactiondetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tduo.mutation.UpdatedAt(); ok {
		_spec.SetField(transactiondetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.DeletedAt(); ok {
		_spec.SetField(transactiondetail.FieldDeletedAt, field.TypeTime, value)
	}
	if tduo.mutation.DeletedAtCleared() {
		_spec.ClearField(transactiondetail.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tduo.mutation.OrderID(); ok {
		_spec.SetField(transactiondetail.FieldOrderID, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.AddedOrderID(); ok {
		_spec.AddField(transactiondetail.FieldOrderID, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.TransactionID(); ok {
		_spec.SetField(transactiondetail.FieldTransactionID, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.AddedTransactionID(); ok {
		_spec.AddField(transactiondetail.FieldTransactionID, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.TransactionType(); ok {
		_spec.SetField(transactiondetail.FieldTransactionType, field.TypeString, value)
	}
	if value, ok := tduo.mutation.DeliveryCost(); ok {
		_spec.SetField(transactiondetail.FieldDeliveryCost, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedDeliveryCost(); ok {
		_spec.AddField(transactiondetail.FieldDeliveryCost, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.MiscFee(); ok {
		_spec.SetField(transactiondetail.FieldMiscFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedMiscFee(); ok {
		_spec.AddField(transactiondetail.FieldMiscFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.FuelFee(); ok {
		_spec.SetField(transactiondetail.FieldFuelFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedFuelFee(); ok {
		_spec.AddField(transactiondetail.FieldFuelFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.RegistrationFee(); ok {
		_spec.SetField(transactiondetail.FieldRegistrationFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedRegistrationFee(); ok {
		_spec.AddField(transactiondetail.FieldRegistrationFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.ProcessingFee(); ok {
		_spec.SetField(transactiondetail.FieldProcessingFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedProcessingFee(); ok {
		_spec.AddField(transactiondetail.FieldProcessingFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.PackageFee(); ok {
		_spec.SetField(transactiondetail.FieldPackageFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedPackageFee(); ok {
		_spec.AddField(transactiondetail.FieldPackageFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.HandlingFee(); ok {
		_spec.SetField(transactiondetail.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedHandlingFee(); ok {
		_spec.AddField(transactiondetail.FieldHandlingFee, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.Vat(); ok {
		_spec.SetField(transactiondetail.FieldVat, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedVat(); ok {
		_spec.AddField(transactiondetail.FieldVat, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.Amount(); ok {
		_spec.SetField(transactiondetail.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedAmount(); ok {
		_spec.AddField(transactiondetail.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.Weight(); ok {
		_spec.SetField(transactiondetail.FieldWeight, field.TypeInt, value)
	}
	if value, ok := tduo.mutation.AddedWeight(); ok {
		_spec.AddField(transactiondetail.FieldWeight, field.TypeInt, value)
	}
	_spec.Node.Schema = tduo.schemaConfig.TransactionDetail
	ctx = internal.NewSchemaConfigContext(ctx, tduo.schemaConfig)
	_spec.AddModifiers(tduo.modifiers...)
	_node = &TransactionDetail{config: tduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactiondetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
