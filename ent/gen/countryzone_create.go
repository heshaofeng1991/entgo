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
)

// CountryZoneCreate is the builder for creating a CountryZone entity.
type CountryZoneCreate struct {
	config
	mutation *CountryZoneMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetChannelID sets the "channel_id" field.
func (czc *CountryZoneCreate) SetChannelID(i int64) *CountryZoneCreate {
	czc.mutation.SetChannelID(i)
	return czc
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (czc *CountryZoneCreate) SetNillableChannelID(i *int64) *CountryZoneCreate {
	if i != nil {
		czc.SetChannelID(*i)
	}
	return czc
}

// SetCountryCode sets the "country_code" field.
func (czc *CountryZoneCreate) SetCountryCode(s string) *CountryZoneCreate {
	czc.mutation.SetCountryCode(s)
	return czc
}

// SetZipCode sets the "zip_code" field.
func (czc *CountryZoneCreate) SetZipCode(s string) *CountryZoneCreate {
	czc.mutation.SetZipCode(s)
	return czc
}

// SetNillableZipCode sets the "zip_code" field if the given value is not nil.
func (czc *CountryZoneCreate) SetNillableZipCode(s *string) *CountryZoneCreate {
	if s != nil {
		czc.SetZipCode(*s)
	}
	return czc
}

// SetStartZipCode sets the "start_zip_code" field.
func (czc *CountryZoneCreate) SetStartZipCode(s string) *CountryZoneCreate {
	czc.mutation.SetStartZipCode(s)
	return czc
}

// SetNillableStartZipCode sets the "start_zip_code" field if the given value is not nil.
func (czc *CountryZoneCreate) SetNillableStartZipCode(s *string) *CountryZoneCreate {
	if s != nil {
		czc.SetStartZipCode(*s)
	}
	return czc
}

// SetEndZipCode sets the "end_zip_code" field.
func (czc *CountryZoneCreate) SetEndZipCode(s string) *CountryZoneCreate {
	czc.mutation.SetEndZipCode(s)
	return czc
}

// SetNillableEndZipCode sets the "end_zip_code" field if the given value is not nil.
func (czc *CountryZoneCreate) SetNillableEndZipCode(s *string) *CountryZoneCreate {
	if s != nil {
		czc.SetEndZipCode(*s)
	}
	return czc
}

// SetZone sets the "zone" field.
func (czc *CountryZoneCreate) SetZone(s string) *CountryZoneCreate {
	czc.mutation.SetZone(s)
	return czc
}

// Mutation returns the CountryZoneMutation object of the builder.
func (czc *CountryZoneCreate) Mutation() *CountryZoneMutation {
	return czc.mutation
}

// Save creates the CountryZone in the database.
func (czc *CountryZoneCreate) Save(ctx context.Context) (*CountryZone, error) {
	var (
		err  error
		node *CountryZone
	)
	czc.defaults()
	if len(czc.hooks) == 0 {
		if err = czc.check(); err != nil {
			return nil, err
		}
		node, err = czc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryZoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = czc.check(); err != nil {
				return nil, err
			}
			czc.mutation = mutation
			if node, err = czc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(czc.hooks) - 1; i >= 0; i-- {
			if czc.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = czc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, czc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (czc *CountryZoneCreate) SaveX(ctx context.Context) *CountryZone {
	v, err := czc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (czc *CountryZoneCreate) Exec(ctx context.Context) error {
	_, err := czc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (czc *CountryZoneCreate) ExecX(ctx context.Context) {
	if err := czc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (czc *CountryZoneCreate) defaults() {
	if _, ok := czc.mutation.ChannelID(); !ok {
		v := countryzone.DefaultChannelID
		czc.mutation.SetChannelID(v)
	}
	if _, ok := czc.mutation.ZipCode(); !ok {
		v := countryzone.DefaultZipCode
		czc.mutation.SetZipCode(v)
	}
	if _, ok := czc.mutation.StartZipCode(); !ok {
		v := countryzone.DefaultStartZipCode
		czc.mutation.SetStartZipCode(v)
	}
	if _, ok := czc.mutation.EndZipCode(); !ok {
		v := countryzone.DefaultEndZipCode
		czc.mutation.SetEndZipCode(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (czc *CountryZoneCreate) check() error {
	if _, ok := czc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`gen: missing required field "CountryZone.channel_id"`)}
	}
	if _, ok := czc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`gen: missing required field "CountryZone.country_code"`)}
	}
	if v, ok := czc.mutation.CountryCode(); ok {
		if err := countryzone.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`gen: validator failed for field "CountryZone.country_code": %w`, err)}
		}
	}
	if _, ok := czc.mutation.ZipCode(); !ok {
		return &ValidationError{Name: "zip_code", err: errors.New(`gen: missing required field "CountryZone.zip_code"`)}
	}
	if _, ok := czc.mutation.StartZipCode(); !ok {
		return &ValidationError{Name: "start_zip_code", err: errors.New(`gen: missing required field "CountryZone.start_zip_code"`)}
	}
	if _, ok := czc.mutation.EndZipCode(); !ok {
		return &ValidationError{Name: "end_zip_code", err: errors.New(`gen: missing required field "CountryZone.end_zip_code"`)}
	}
	if _, ok := czc.mutation.Zone(); !ok {
		return &ValidationError{Name: "zone", err: errors.New(`gen: missing required field "CountryZone.zone"`)}
	}
	if v, ok := czc.mutation.Zone(); ok {
		if err := countryzone.ZoneValidator(v); err != nil {
			return &ValidationError{Name: "zone", err: fmt.Errorf(`gen: validator failed for field "CountryZone.zone": %w`, err)}
		}
	}
	return nil
}

func (czc *CountryZoneCreate) sqlSave(ctx context.Context) (*CountryZone, error) {
	_node, _spec := czc.createSpec()
	if err := sqlgraph.CreateNode(ctx, czc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (czc *CountryZoneCreate) createSpec() (*CountryZone, *sqlgraph.CreateSpec) {
	var (
		_node = &CountryZone{config: czc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: countryzone.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: countryzone.FieldID,
			},
		}
	)
	_spec.Schema = czc.schemaConfig.CountryZone
	_spec.OnConflict = czc.conflict
	if value, ok := czc.mutation.ChannelID(); ok {
		_spec.SetField(countryzone.FieldChannelID, field.TypeInt64, value)
		_node.ChannelID = value
	}
	if value, ok := czc.mutation.CountryCode(); ok {
		_spec.SetField(countryzone.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = value
	}
	if value, ok := czc.mutation.ZipCode(); ok {
		_spec.SetField(countryzone.FieldZipCode, field.TypeString, value)
		_node.ZipCode = value
	}
	if value, ok := czc.mutation.StartZipCode(); ok {
		_spec.SetField(countryzone.FieldStartZipCode, field.TypeString, value)
		_node.StartZipCode = value
	}
	if value, ok := czc.mutation.EndZipCode(); ok {
		_spec.SetField(countryzone.FieldEndZipCode, field.TypeString, value)
		_node.EndZipCode = value
	}
	if value, ok := czc.mutation.Zone(); ok {
		_spec.SetField(countryzone.FieldZone, field.TypeString, value)
		_node.Zone = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CountryZone.Create().
//		SetChannelID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CountryZoneUpsert) {
//			SetChannelID(v+v).
//		}).
//		Exec(ctx)
func (czc *CountryZoneCreate) OnConflict(opts ...sql.ConflictOption) *CountryZoneUpsertOne {
	czc.conflict = opts
	return &CountryZoneUpsertOne{
		create: czc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CountryZone.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (czc *CountryZoneCreate) OnConflictColumns(columns ...string) *CountryZoneUpsertOne {
	czc.conflict = append(czc.conflict, sql.ConflictColumns(columns...))
	return &CountryZoneUpsertOne{
		create: czc,
	}
}

type (
	// CountryZoneUpsertOne is the builder for "upsert"-ing
	//  one CountryZone node.
	CountryZoneUpsertOne struct {
		create *CountryZoneCreate
	}

	// CountryZoneUpsert is the "OnConflict" setter.
	CountryZoneUpsert struct {
		*sql.UpdateSet
	}
)

// SetChannelID sets the "channel_id" field.
func (u *CountryZoneUpsert) SetChannelID(v int64) *CountryZoneUpsert {
	u.Set(countryzone.FieldChannelID, v)
	return u
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *CountryZoneUpsert) UpdateChannelID() *CountryZoneUpsert {
	u.SetExcluded(countryzone.FieldChannelID)
	return u
}

// AddChannelID adds v to the "channel_id" field.
func (u *CountryZoneUpsert) AddChannelID(v int64) *CountryZoneUpsert {
	u.Add(countryzone.FieldChannelID, v)
	return u
}

// SetCountryCode sets the "country_code" field.
func (u *CountryZoneUpsert) SetCountryCode(v string) *CountryZoneUpsert {
	u.Set(countryzone.FieldCountryCode, v)
	return u
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *CountryZoneUpsert) UpdateCountryCode() *CountryZoneUpsert {
	u.SetExcluded(countryzone.FieldCountryCode)
	return u
}

// SetZipCode sets the "zip_code" field.
func (u *CountryZoneUpsert) SetZipCode(v string) *CountryZoneUpsert {
	u.Set(countryzone.FieldZipCode, v)
	return u
}

// UpdateZipCode sets the "zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsert) UpdateZipCode() *CountryZoneUpsert {
	u.SetExcluded(countryzone.FieldZipCode)
	return u
}

// SetStartZipCode sets the "start_zip_code" field.
func (u *CountryZoneUpsert) SetStartZipCode(v string) *CountryZoneUpsert {
	u.Set(countryzone.FieldStartZipCode, v)
	return u
}

// UpdateStartZipCode sets the "start_zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsert) UpdateStartZipCode() *CountryZoneUpsert {
	u.SetExcluded(countryzone.FieldStartZipCode)
	return u
}

// SetEndZipCode sets the "end_zip_code" field.
func (u *CountryZoneUpsert) SetEndZipCode(v string) *CountryZoneUpsert {
	u.Set(countryzone.FieldEndZipCode, v)
	return u
}

// UpdateEndZipCode sets the "end_zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsert) UpdateEndZipCode() *CountryZoneUpsert {
	u.SetExcluded(countryzone.FieldEndZipCode)
	return u
}

// SetZone sets the "zone" field.
func (u *CountryZoneUpsert) SetZone(v string) *CountryZoneUpsert {
	u.Set(countryzone.FieldZone, v)
	return u
}

// UpdateZone sets the "zone" field to the value that was provided on create.
func (u *CountryZoneUpsert) UpdateZone() *CountryZoneUpsert {
	u.SetExcluded(countryzone.FieldZone)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.CountryZone.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CountryZoneUpsertOne) UpdateNewValues() *CountryZoneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CountryZone.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CountryZoneUpsertOne) Ignore() *CountryZoneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CountryZoneUpsertOne) DoNothing() *CountryZoneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CountryZoneCreate.OnConflict
// documentation for more info.
func (u *CountryZoneUpsertOne) Update(set func(*CountryZoneUpsert)) *CountryZoneUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CountryZoneUpsert{UpdateSet: update})
	}))
	return u
}

// SetChannelID sets the "channel_id" field.
func (u *CountryZoneUpsertOne) SetChannelID(v int64) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetChannelID(v)
	})
}

// AddChannelID adds v to the "channel_id" field.
func (u *CountryZoneUpsertOne) AddChannelID(v int64) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.AddChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *CountryZoneUpsertOne) UpdateChannelID() *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateChannelID()
	})
}

// SetCountryCode sets the "country_code" field.
func (u *CountryZoneUpsertOne) SetCountryCode(v string) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetCountryCode(v)
	})
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *CountryZoneUpsertOne) UpdateCountryCode() *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateCountryCode()
	})
}

// SetZipCode sets the "zip_code" field.
func (u *CountryZoneUpsertOne) SetZipCode(v string) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetZipCode(v)
	})
}

// UpdateZipCode sets the "zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsertOne) UpdateZipCode() *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateZipCode()
	})
}

// SetStartZipCode sets the "start_zip_code" field.
func (u *CountryZoneUpsertOne) SetStartZipCode(v string) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetStartZipCode(v)
	})
}

// UpdateStartZipCode sets the "start_zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsertOne) UpdateStartZipCode() *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateStartZipCode()
	})
}

// SetEndZipCode sets the "end_zip_code" field.
func (u *CountryZoneUpsertOne) SetEndZipCode(v string) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetEndZipCode(v)
	})
}

// UpdateEndZipCode sets the "end_zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsertOne) UpdateEndZipCode() *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateEndZipCode()
	})
}

// SetZone sets the "zone" field.
func (u *CountryZoneUpsertOne) SetZone(v string) *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetZone(v)
	})
}

// UpdateZone sets the "zone" field to the value that was provided on create.
func (u *CountryZoneUpsertOne) UpdateZone() *CountryZoneUpsertOne {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateZone()
	})
}

// Exec executes the query.
func (u *CountryZoneUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for CountryZoneCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CountryZoneUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CountryZoneUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CountryZoneUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CountryZoneCreateBulk is the builder for creating many CountryZone entities in bulk.
type CountryZoneCreateBulk struct {
	config
	builders []*CountryZoneCreate
	conflict []sql.ConflictOption
}

// Save creates the CountryZone entities in the database.
func (czcb *CountryZoneCreateBulk) Save(ctx context.Context) ([]*CountryZone, error) {
	specs := make([]*sqlgraph.CreateSpec, len(czcb.builders))
	nodes := make([]*CountryZone, len(czcb.builders))
	mutators := make([]Mutator, len(czcb.builders))
	for i := range czcb.builders {
		func(i int, root context.Context) {
			builder := czcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryZoneMutation)
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
					_, err = mutators[i+1].Mutate(root, czcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = czcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, czcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, czcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (czcb *CountryZoneCreateBulk) SaveX(ctx context.Context) []*CountryZone {
	v, err := czcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (czcb *CountryZoneCreateBulk) Exec(ctx context.Context) error {
	_, err := czcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (czcb *CountryZoneCreateBulk) ExecX(ctx context.Context) {
	if err := czcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CountryZone.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CountryZoneUpsert) {
//			SetChannelID(v+v).
//		}).
//		Exec(ctx)
func (czcb *CountryZoneCreateBulk) OnConflict(opts ...sql.ConflictOption) *CountryZoneUpsertBulk {
	czcb.conflict = opts
	return &CountryZoneUpsertBulk{
		create: czcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CountryZone.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (czcb *CountryZoneCreateBulk) OnConflictColumns(columns ...string) *CountryZoneUpsertBulk {
	czcb.conflict = append(czcb.conflict, sql.ConflictColumns(columns...))
	return &CountryZoneUpsertBulk{
		create: czcb,
	}
}

// CountryZoneUpsertBulk is the builder for "upsert"-ing
// a bulk of CountryZone nodes.
type CountryZoneUpsertBulk struct {
	create *CountryZoneCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CountryZone.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CountryZoneUpsertBulk) UpdateNewValues() *CountryZoneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CountryZone.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CountryZoneUpsertBulk) Ignore() *CountryZoneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CountryZoneUpsertBulk) DoNothing() *CountryZoneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CountryZoneCreateBulk.OnConflict
// documentation for more info.
func (u *CountryZoneUpsertBulk) Update(set func(*CountryZoneUpsert)) *CountryZoneUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CountryZoneUpsert{UpdateSet: update})
	}))
	return u
}

// SetChannelID sets the "channel_id" field.
func (u *CountryZoneUpsertBulk) SetChannelID(v int64) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetChannelID(v)
	})
}

// AddChannelID adds v to the "channel_id" field.
func (u *CountryZoneUpsertBulk) AddChannelID(v int64) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.AddChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *CountryZoneUpsertBulk) UpdateChannelID() *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateChannelID()
	})
}

// SetCountryCode sets the "country_code" field.
func (u *CountryZoneUpsertBulk) SetCountryCode(v string) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetCountryCode(v)
	})
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *CountryZoneUpsertBulk) UpdateCountryCode() *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateCountryCode()
	})
}

// SetZipCode sets the "zip_code" field.
func (u *CountryZoneUpsertBulk) SetZipCode(v string) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetZipCode(v)
	})
}

// UpdateZipCode sets the "zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsertBulk) UpdateZipCode() *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateZipCode()
	})
}

// SetStartZipCode sets the "start_zip_code" field.
func (u *CountryZoneUpsertBulk) SetStartZipCode(v string) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetStartZipCode(v)
	})
}

// UpdateStartZipCode sets the "start_zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsertBulk) UpdateStartZipCode() *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateStartZipCode()
	})
}

// SetEndZipCode sets the "end_zip_code" field.
func (u *CountryZoneUpsertBulk) SetEndZipCode(v string) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetEndZipCode(v)
	})
}

// UpdateEndZipCode sets the "end_zip_code" field to the value that was provided on create.
func (u *CountryZoneUpsertBulk) UpdateEndZipCode() *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateEndZipCode()
	})
}

// SetZone sets the "zone" field.
func (u *CountryZoneUpsertBulk) SetZone(v string) *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.SetZone(v)
	})
}

// UpdateZone sets the "zone" field to the value that was provided on create.
func (u *CountryZoneUpsertBulk) UpdateZone() *CountryZoneUpsertBulk {
	return u.Update(func(s *CountryZoneUpsert) {
		s.UpdateZone()
	})
}

// Exec executes the query.
func (u *CountryZoneUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("gen: OnConflict was set for builder %d. Set it on the CountryZoneCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for CountryZoneCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CountryZoneUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
