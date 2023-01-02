// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/countryzone"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// CountryZoneUpdate is the builder for updating CountryZone entities.
type CountryZoneUpdate struct {
	config
	hooks     []Hook
	mutation  *CountryZoneMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CountryZoneUpdate builder.
func (czu *CountryZoneUpdate) Where(ps ...predicate.CountryZone) *CountryZoneUpdate {
	czu.mutation.Where(ps...)
	return czu
}

// SetChannelID sets the "channel_id" field.
func (czu *CountryZoneUpdate) SetChannelID(i int64) *CountryZoneUpdate {
	czu.mutation.ResetChannelID()
	czu.mutation.SetChannelID(i)
	return czu
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (czu *CountryZoneUpdate) SetNillableChannelID(i *int64) *CountryZoneUpdate {
	if i != nil {
		czu.SetChannelID(*i)
	}
	return czu
}

// AddChannelID adds i to the "channel_id" field.
func (czu *CountryZoneUpdate) AddChannelID(i int64) *CountryZoneUpdate {
	czu.mutation.AddChannelID(i)
	return czu
}

// SetCountryCode sets the "country_code" field.
func (czu *CountryZoneUpdate) SetCountryCode(s string) *CountryZoneUpdate {
	czu.mutation.SetCountryCode(s)
	return czu
}

// SetZipCode sets the "zip_code" field.
func (czu *CountryZoneUpdate) SetZipCode(s string) *CountryZoneUpdate {
	czu.mutation.SetZipCode(s)
	return czu
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (czu *CountryZoneUpdate) SetNillableZipCode(s *string) *CountryZoneUpdate {
	if s != nil {
		czu.SetZipCode(*s)
	}
	return czu
}

// SetStartZipCode sets the "start_zip_code" field.
func (czu *CountryZoneUpdate) SetStartZipCode(s string) *CountryZoneUpdate {
	czu.mutation.SetStartZipCode(s)
	return czu
}

// SetNillableStartZipCode sets the "start_zip_code" field if the given value is not nil.
func (czu *CountryZoneUpdate) SetNillableStartZipCode(s *string) *CountryZoneUpdate {
	if s != nil {
		czu.SetStartZipCode(*s)
	}
	return czu
}

// SetEndZipCode sets the "end_zip_code" field.
func (czu *CountryZoneUpdate) SetEndZipCode(s string) *CountryZoneUpdate {
	czu.mutation.SetEndZipCode(s)
	return czu
}

// SetNillableEndZipCode sets the "end_zip_code" field if the given value is not nil.
func (czu *CountryZoneUpdate) SetNillableEndZipCode(s *string) *CountryZoneUpdate {
	if s != nil {
		czu.SetEndZipCode(*s)
	}
	return czu
}

// SetZone sets the "zone" field.
func (czu *CountryZoneUpdate) SetZone(s string) *CountryZoneUpdate {
	czu.mutation.SetZone(s)
	return czu
}

// Mutation returns the CountryZoneMutation object of the builder.
func (czu *CountryZoneUpdate) Mutation() *CountryZoneMutation {
	return czu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (czu *CountryZoneUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(czu.hooks) == 0 {
		if err = czu.check(); err != nil {
			return 0, err
		}
		affected, err = czu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryZoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = czu.check(); err != nil {
				return 0, err
			}
			czu.mutation = mutation
			affected, err = czu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(czu.hooks) - 1; i >= 0; i-- {
			if czu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = czu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, czu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (czu *CountryZoneUpdate) SaveX(ctx context.Context) int {
	affected, err := czu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (czu *CountryZoneUpdate) Exec(ctx context.Context) error {
	_, err := czu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (czu *CountryZoneUpdate) ExecX(ctx context.Context) {
	if err := czu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (czu *CountryZoneUpdate) check() error {
	if v, ok := czu.mutation.CountryCode(); ok {
		if err := countryzone.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`gen: validator failed for field "CountryZone.country_code": %w`, err)}
		}
	}
	if v, ok := czu.mutation.Zone(); ok {
		if err := countryzone.ZoneValidator(v); err != nil {
			return &ValidationError{Name: "zone", err: fmt.Errorf(`gen: validator failed for field "CountryZone.zone": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (czu *CountryZoneUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CountryZoneUpdate {
	czu.modifiers = append(czu.modifiers, modifiers...)
	return czu
}

func (czu *CountryZoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   countryzone.Table,
			Columns: countryzone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: countryzone.FieldID,
			},
		},
	}
	if ps := czu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := czu.mutation.ChannelID(); ok {
		_spec.SetField(countryzone.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := czu.mutation.AddedChannelID(); ok {
		_spec.AddField(countryzone.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := czu.mutation.CountryCode(); ok {
		_spec.SetField(countryzone.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := czu.mutation.ZipCode(); ok {
		_spec.SetField(countryzone.FieldZipCode, field.TypeString, value)
	}
	if value, ok := czu.mutation.StartZipCode(); ok {
		_spec.SetField(countryzone.FieldStartZipCode, field.TypeString, value)
	}
	if value, ok := czu.mutation.EndZipCode(); ok {
		_spec.SetField(countryzone.FieldEndZipCode, field.TypeString, value)
	}
	if value, ok := czu.mutation.Zone(); ok {
		_spec.SetField(countryzone.FieldZone, field.TypeString, value)
	}
	_spec.Node.Schema = czu.schemaConfig.CountryZone
	ctx = internal.NewSchemaConfigContext(ctx, czu.schemaConfig)
	_spec.AddModifiers(czu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, czu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{countryzone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CountryZoneUpdateOne is the builder for updating a single CountryZone entity.
type CountryZoneUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CountryZoneMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetChannelID sets the "channel_id" field.
func (czuo *CountryZoneUpdateOne) SetChannelID(i int64) *CountryZoneUpdateOne {
	czuo.mutation.ResetChannelID()
	czuo.mutation.SetChannelID(i)
	return czuo
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (czuo *CountryZoneUpdateOne) SetNillableChannelID(i *int64) *CountryZoneUpdateOne {
	if i != nil {
		czuo.SetChannelID(*i)
	}
	return czuo
}

// AddChannelID adds i to the "channel_id" field.
func (czuo *CountryZoneUpdateOne) AddChannelID(i int64) *CountryZoneUpdateOne {
	czuo.mutation.AddChannelID(i)
	return czuo
}

// SetCountryCode sets the "country_code" field.
func (czuo *CountryZoneUpdateOne) SetCountryCode(s string) *CountryZoneUpdateOne {
	czuo.mutation.SetCountryCode(s)
	return czuo
}

// SetZipCode sets the "zip_code" field.
func (czuo *CountryZoneUpdateOne) SetZipCode(s string) *CountryZoneUpdateOne {
	czuo.mutation.SetZipCode(s)
	return czuo
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (czuo *CountryZoneUpdateOne) SetNillableZipCode(s *string) *CountryZoneUpdateOne {
	if s != nil {
		czuo.SetZipCode(*s)
	}
	return czuo
}

// SetStartZipCode sets the "start_zip_code" field.
func (czuo *CountryZoneUpdateOne) SetStartZipCode(s string) *CountryZoneUpdateOne {
	czuo.mutation.SetStartZipCode(s)
	return czuo
}

// SetNillableStartZipCode sets the "start_zip_code" field if the given value is not nil.
func (czuo *CountryZoneUpdateOne) SetNillableStartZipCode(s *string) *CountryZoneUpdateOne {
	if s != nil {
		czuo.SetStartZipCode(*s)
	}
	return czuo
}

// SetEndZipCode sets the "end_zip_code" field.
func (czuo *CountryZoneUpdateOne) SetEndZipCode(s string) *CountryZoneUpdateOne {
	czuo.mutation.SetEndZipCode(s)
	return czuo
}

// SetNillableEndZipCode sets the "end_zip_code" field if the given value is not nil.
func (czuo *CountryZoneUpdateOne) SetNillableEndZipCode(s *string) *CountryZoneUpdateOne {
	if s != nil {
		czuo.SetEndZipCode(*s)
	}
	return czuo
}

// SetZone sets the "zone" field.
func (czuo *CountryZoneUpdateOne) SetZone(s string) *CountryZoneUpdateOne {
	czuo.mutation.SetZone(s)
	return czuo
}

// Mutation returns the CountryZoneMutation object of the builder.
func (czuo *CountryZoneUpdateOne) Mutation() *CountryZoneMutation {
	return czuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (czuo *CountryZoneUpdateOne) Select(field string, fields ...string) *CountryZoneUpdateOne {
	czuo.fields = append([]string{field}, fields...)
	return czuo
}

// Save executes the query and returns the updated CountryZone entity.
func (czuo *CountryZoneUpdateOne) Save(ctx context.Context) (*CountryZone, error) {
	var (
		err  error
		node *CountryZone
	)
	if len(czuo.hooks) == 0 {
		if err = czuo.check(); err != nil {
			return nil, err
		}
		node, err = czuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryZoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = czuo.check(); err != nil {
				return nil, err
			}
			czuo.mutation = mutation
			node, err = czuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(czuo.hooks) - 1; i >= 0; i-- {
			if czuo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = czuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, czuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CountryZone)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CountryZoneMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (czuo *CountryZoneUpdateOne) SaveX(ctx context.Context) *CountryZone {
	node, err := czuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (czuo *CountryZoneUpdateOne) Exec(ctx context.Context) error {
	_, err := czuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (czuo *CountryZoneUpdateOne) ExecX(ctx context.Context) {
	if err := czuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (czuo *CountryZoneUpdateOne) check() error {
	if v, ok := czuo.mutation.CountryCode(); ok {
		if err := countryzone.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`gen: validator failed for field "CountryZone.country_code": %w`, err)}
		}
	}
	if v, ok := czuo.mutation.Zone(); ok {
		if err := countryzone.ZoneValidator(v); err != nil {
			return &ValidationError{Name: "zone", err: fmt.Errorf(`gen: validator failed for field "CountryZone.zone": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (czuo *CountryZoneUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CountryZoneUpdateOne {
	czuo.modifiers = append(czuo.modifiers, modifiers...)
	return czuo
}

func (czuo *CountryZoneUpdateOne) sqlSave(ctx context.Context) (_node *CountryZone, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   countryzone.Table,
			Columns: countryzone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: countryzone.FieldID,
			},
		},
	}
	id, ok := czuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "CountryZone.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := czuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, countryzone.FieldID)
		for _, f := range fields {
			if !countryzone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != countryzone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := czuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := czuo.mutation.ChannelID(); ok {
		_spec.SetField(countryzone.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := czuo.mutation.AddedChannelID(); ok {
		_spec.AddField(countryzone.FieldChannelID, field.TypeInt64, value)
	}
	if value, ok := czuo.mutation.CountryCode(); ok {
		_spec.SetField(countryzone.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := czuo.mutation.ZipCode(); ok {
		_spec.SetField(countryzone.FieldZipCode, field.TypeString, value)
	}
	if value, ok := czuo.mutation.StartZipCode(); ok {
		_spec.SetField(countryzone.FieldStartZipCode, field.TypeString, value)
	}
	if value, ok := czuo.mutation.EndZipCode(); ok {
		_spec.SetField(countryzone.FieldEndZipCode, field.TypeString, value)
	}
	if value, ok := czuo.mutation.Zone(); ok {
		_spec.SetField(countryzone.FieldZone, field.TypeString, value)
	}
	_spec.Node.Schema = czuo.schemaConfig.CountryZone
	ctx = internal.NewSchemaConfigContext(ctx, czuo.schemaConfig)
	_spec.AddModifiers(czuo.modifiers...)
	_node = &CountryZone{config: czuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, czuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{countryzone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}