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
	"github.com/heshaofeng1991/entgo/ent/gen/channelvolumefactor"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ChannelVolumeFactorUpdate is the builder for updating ChannelVolumeFactor entities.
type ChannelVolumeFactorUpdate struct {
	config
	hooks     []Hook
	mutation  *ChannelVolumeFactorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ChannelVolumeFactorUpdate builder.
func (cvfu *ChannelVolumeFactorUpdate) Where(ps ...predicate.ChannelVolumeFactor) *ChannelVolumeFactorUpdate {
	cvfu.mutation.Where(ps...)
	return cvfu
}

// SetUpdatedAt sets the "updated_at" field.
func (cvfu *ChannelVolumeFactorUpdate) SetUpdatedAt(t time.Time) *ChannelVolumeFactorUpdate {
	cvfu.mutation.SetUpdatedAt(t)
	return cvfu
}

// SetDeletedAt sets the "deleted_at" field.
func (cvfu *ChannelVolumeFactorUpdate) SetDeletedAt(t time.Time) *ChannelVolumeFactorUpdate {
	cvfu.mutation.SetDeletedAt(t)
	return cvfu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cvfu *ChannelVolumeFactorUpdate) SetNillableDeletedAt(t *time.Time) *ChannelVolumeFactorUpdate {
	if t != nil {
		cvfu.SetDeletedAt(*t)
	}
	return cvfu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cvfu *ChannelVolumeFactorUpdate) ClearDeletedAt() *ChannelVolumeFactorUpdate {
	cvfu.mutation.ClearDeletedAt()
	return cvfu
}

// SetCountryCode sets the "country_code" field.
func (cvfu *ChannelVolumeFactorUpdate) SetCountryCode(s string) *ChannelVolumeFactorUpdate {
	cvfu.mutation.SetCountryCode(s)
	return cvfu
}

// SetChannelID sets the "channel_id" field.
func (cvfu *ChannelVolumeFactorUpdate) SetChannelID(i int64) *ChannelVolumeFactorUpdate {
	cvfu.mutation.ResetChannelID()
	cvfu.mutation.SetChannelID(i)
	return cvfu
}

// AddChannelID adds i to the "channel_id" field.
func (cvfu *ChannelVolumeFactorUpdate) AddChannelID(i int64) *ChannelVolumeFactorUpdate {
	cvfu.mutation.AddChannelID(i)
	return cvfu
}

// SetVolumeFactor sets the "volume_factor" field.
func (cvfu *ChannelVolumeFactorUpdate) SetVolumeFactor(i int) *ChannelVolumeFactorUpdate {
	cvfu.mutation.ResetVolumeFactor()
	cvfu.mutation.SetVolumeFactor(i)
	return cvfu
}

// SetNillableVolumeFactor sets the "volume_factor" field if the given value is not nil.
func (cvfu *ChannelVolumeFactorUpdate) SetNillableVolumeFactor(i *int) *ChannelVolumeFactorUpdate {
	if i != nil {
		cvfu.SetVolumeFactor(*i)
	}
	return cvfu
}

// AddVolumeFactor adds i to the "volume_factor" field.
func (cvfu *ChannelVolumeFactorUpdate) AddVolumeFactor(i int) *ChannelVolumeFactorUpdate {
	cvfu.mutation.AddVolumeFactor(i)
	return cvfu
}

// SetStatus sets the "status" field.
func (cvfu *ChannelVolumeFactorUpdate) SetStatus(i int8) *ChannelVolumeFactorUpdate {
	cvfu.mutation.ResetStatus()
	cvfu.mutation.SetStatus(i)
	return cvfu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cvfu *ChannelVolumeFactorUpdate) SetNillableStatus(i *int8) *ChannelVolumeFactorUpdate {
	if i != nil {
		cvfu.SetStatus(*i)
	}
	return cvfu
}

// AddStatus adds i to the "status" field.
func (cvfu *ChannelVolumeFactorUpdate) AddStatus(i int8) *ChannelVolumeFactorUpdate {
	cvfu.mutation.AddStatus(i)
	return cvfu
}

// Mutation returns the ChannelVolumeFactorMutation object of the builder.
func (cvfu *ChannelVolumeFactorUpdate) Mutation() *ChannelVolumeFactorMutation {
	return cvfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cvfu *ChannelVolumeFactorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cvfu.defaults()
	if len(cvfu.hooks) == 0 {
		affected, err = cvfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelVolumeFactorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cvfu.mutation = mutation
			affected, err = cvfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cvfu.hooks) - 1; i >= 0; i-- {
			if cvfu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = cvfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cvfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cvfu *ChannelVolumeFactorUpdate) SaveX(ctx context.Context) int {
	affected, err := cvfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cvfu *ChannelVolumeFactorUpdate) Exec(ctx context.Context) error {
	_, err := cvfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvfu *ChannelVolumeFactorUpdate) ExecX(ctx context.Context) {
	if err := cvfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cvfu *ChannelVolumeFactorUpdate) defaults() {
	if _, ok := cvfu.mutation.UpdatedAt(); !ok {
		v := channelvolumefactor.UpdateDefaultUpdatedAt()
		cvfu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cvfu *ChannelVolumeFactorUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChannelVolumeFactorUpdate {
	cvfu.modifiers = append(cvfu.modifiers, modifiers...)
	return cvfu
}

func (cvfu *ChannelVolumeFactorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channelvolumefactor.Table,
			Columns: channelvolumefactor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channelvolumefactor.FieldID,
			},
		},
	}
	if ps := cvfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cvfu.mutation.UpdatedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cvfu.mutation.DeletedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldDeletedAt, field.TypeTime, value)
	}
	if cvfu.mutation.DeletedAtCleared() {
		_spec.ClearField(channelvolumefactor.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cvfu.mutation.CountryCode(); ok {
		_spec.SetField(channelvolumefactor.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := cvfu.mutation.ChannelID(); ok {
		_spec.SetField(channelvolumefactor.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := cvfu.mutation.AddedChannelID(); ok {
		_spec.AddField(channelvolumefactor.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := cvfu.mutation.VolumeFactor(); ok {
		_spec.SetField(channelvolumefactor.FieldVolumeFactor, field.TypeInt, value)
	}
	if value, ok := cvfu.mutation.AddedVolumeFactor(); ok {
		_spec.AddField(channelvolumefactor.FieldVolumeFactor, field.TypeInt, value)
	}
	if value, ok := cvfu.mutation.Status(); ok {
		_spec.SetField(channelvolumefactor.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cvfu.mutation.AddedStatus(); ok {
		_spec.AddField(channelvolumefactor.FieldStatus, field.TypeInt8, value)
	}
	_spec.Node.Schema = cvfu.schemaConfig.ChannelVolumeFactor
	ctx = internal.NewSchemaConfigContext(ctx, cvfu.schemaConfig)
	_spec.AddModifiers(cvfu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cvfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channelvolumefactor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ChannelVolumeFactorUpdateOne is the builder for updating a single ChannelVolumeFactor entity.
type ChannelVolumeFactorUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ChannelVolumeFactorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetUpdatedAt(t time.Time) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.SetUpdatedAt(t)
	return cvfuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetDeletedAt(t time.Time) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.SetDeletedAt(t)
	return cvfuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetNillableDeletedAt(t *time.Time) *ChannelVolumeFactorUpdateOne {
	if t != nil {
		cvfuo.SetDeletedAt(*t)
	}
	return cvfuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) ClearDeletedAt() *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.ClearDeletedAt()
	return cvfuo
}

// SetCountryCode sets the "country_code" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetCountryCode(s string) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.SetCountryCode(s)
	return cvfuo
}

// SetChannelID sets the "channel_id" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetChannelID(i int64) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.ResetChannelID()
	cvfuo.mutation.SetChannelID(i)
	return cvfuo
}

// AddChannelID adds i to the "channel_id" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) AddChannelID(i int64) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.AddChannelID(i)
	return cvfuo
}

// SetVolumeFactor sets the "volume_factor" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetVolumeFactor(i int) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.ResetVolumeFactor()
	cvfuo.mutation.SetVolumeFactor(i)
	return cvfuo
}

// SetNillableVolumeFactor sets the "volume_factor" field if the given value is not nil.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetNillableVolumeFactor(i *int) *ChannelVolumeFactorUpdateOne {
	if i != nil {
		cvfuo.SetVolumeFactor(*i)
	}
	return cvfuo
}

// AddVolumeFactor adds i to the "volume_factor" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) AddVolumeFactor(i int) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.AddVolumeFactor(i)
	return cvfuo
}

// SetStatus sets the "status" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetStatus(i int8) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.ResetStatus()
	cvfuo.mutation.SetStatus(i)
	return cvfuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cvfuo *ChannelVolumeFactorUpdateOne) SetNillableStatus(i *int8) *ChannelVolumeFactorUpdateOne {
	if i != nil {
		cvfuo.SetStatus(*i)
	}
	return cvfuo
}

// AddStatus adds i to the "status" field.
func (cvfuo *ChannelVolumeFactorUpdateOne) AddStatus(i int8) *ChannelVolumeFactorUpdateOne {
	cvfuo.mutation.AddStatus(i)
	return cvfuo
}

// Mutation returns the ChannelVolumeFactorMutation object of the builder.
func (cvfuo *ChannelVolumeFactorUpdateOne) Mutation() *ChannelVolumeFactorMutation {
	return cvfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cvfuo *ChannelVolumeFactorUpdateOne) Select(field string, fields ...string) *ChannelVolumeFactorUpdateOne {
	cvfuo.fields = append([]string{field}, fields...)
	return cvfuo
}

// Save executes the query and returns the updated ChannelVolumeFactor entity.
func (cvfuo *ChannelVolumeFactorUpdateOne) Save(ctx context.Context) (*ChannelVolumeFactor, error) {
	var (
		err  error
		node *ChannelVolumeFactor
	)
	cvfuo.defaults()
	if len(cvfuo.hooks) == 0 {
		node, err = cvfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelVolumeFactorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cvfuo.mutation = mutation
			node, err = cvfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cvfuo.hooks) - 1; i >= 0; i-- {
			if cvfuo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = cvfuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cvfuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ChannelVolumeFactor)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ChannelVolumeFactorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cvfuo *ChannelVolumeFactorUpdateOne) SaveX(ctx context.Context) *ChannelVolumeFactor {
	node, err := cvfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cvfuo *ChannelVolumeFactorUpdateOne) Exec(ctx context.Context) error {
	_, err := cvfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvfuo *ChannelVolumeFactorUpdateOne) ExecX(ctx context.Context) {
	if err := cvfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cvfuo *ChannelVolumeFactorUpdateOne) defaults() {
	if _, ok := cvfuo.mutation.UpdatedAt(); !ok {
		v := channelvolumefactor.UpdateDefaultUpdatedAt()
		cvfuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cvfuo *ChannelVolumeFactorUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChannelVolumeFactorUpdateOne {
	cvfuo.modifiers = append(cvfuo.modifiers, modifiers...)
	return cvfuo
}

func (cvfuo *ChannelVolumeFactorUpdateOne) sqlSave(ctx context.Context) (_node *ChannelVolumeFactor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channelvolumefactor.Table,
			Columns: channelvolumefactor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channelvolumefactor.FieldID,
			},
		},
	}
	id, ok := cvfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "ChannelVolumeFactor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cvfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channelvolumefactor.FieldID)
		for _, f := range fields {
			if !channelvolumefactor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != channelvolumefactor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cvfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cvfuo.mutation.UpdatedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cvfuo.mutation.DeletedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldDeletedAt, field.TypeTime, value)
	}
	if cvfuo.mutation.DeletedAtCleared() {
		_spec.ClearField(channelvolumefactor.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cvfuo.mutation.CountryCode(); ok {
		_spec.SetField(channelvolumefactor.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := cvfuo.mutation.ChannelID(); ok {
		_spec.SetField(channelvolumefactor.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := cvfuo.mutation.AddedChannelID(); ok {
		_spec.AddField(channelvolumefactor.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := cvfuo.mutation.VolumeFactor(); ok {
		_spec.SetField(channelvolumefactor.FieldVolumeFactor, field.TypeInt, value)
	}
	if value, ok := cvfuo.mutation.AddedVolumeFactor(); ok {
		_spec.AddField(channelvolumefactor.FieldVolumeFactor, field.TypeInt, value)
	}
	if value, ok := cvfuo.mutation.Status(); ok {
		_spec.SetField(channelvolumefactor.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cvfuo.mutation.AddedStatus(); ok {
		_spec.AddField(channelvolumefactor.FieldStatus, field.TypeInt8, value)
	}
	_spec.Node.Schema = cvfuo.schemaConfig.ChannelVolumeFactor
	ctx = internal.NewSchemaConfigContext(ctx, cvfuo.schemaConfig)
	_spec.AddModifiers(cvfuo.modifiers...)
	_node = &ChannelVolumeFactor{config: cvfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cvfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channelvolumefactor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
