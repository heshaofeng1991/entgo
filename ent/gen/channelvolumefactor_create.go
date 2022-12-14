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
)

// ChannelVolumeFactorCreate is the builder for creating a ChannelVolumeFactor entity.
type ChannelVolumeFactorCreate struct {
	config
	mutation *ChannelVolumeFactorMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cvfc *ChannelVolumeFactorCreate) SetCreatedAt(t time.Time) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetCreatedAt(t)
	return cvfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cvfc *ChannelVolumeFactorCreate) SetNillableCreatedAt(t *time.Time) *ChannelVolumeFactorCreate {
	if t != nil {
		cvfc.SetCreatedAt(*t)
	}
	return cvfc
}

// SetUpdatedAt sets the "updated_at" field.
func (cvfc *ChannelVolumeFactorCreate) SetUpdatedAt(t time.Time) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetUpdatedAt(t)
	return cvfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cvfc *ChannelVolumeFactorCreate) SetNillableUpdatedAt(t *time.Time) *ChannelVolumeFactorCreate {
	if t != nil {
		cvfc.SetUpdatedAt(*t)
	}
	return cvfc
}

// SetDeletedAt sets the "deleted_at" field.
func (cvfc *ChannelVolumeFactorCreate) SetDeletedAt(t time.Time) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetDeletedAt(t)
	return cvfc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cvfc *ChannelVolumeFactorCreate) SetNillableDeletedAt(t *time.Time) *ChannelVolumeFactorCreate {
	if t != nil {
		cvfc.SetDeletedAt(*t)
	}
	return cvfc
}

// SetCountryCode sets the "country_code" field.
func (cvfc *ChannelVolumeFactorCreate) SetCountryCode(s string) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetCountryCode(s)
	return cvfc
}

// SetChannelID sets the "channel_id" field.
func (cvfc *ChannelVolumeFactorCreate) SetChannelID(i int64) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetChannelID(i)
	return cvfc
}

// SetVolumeFactor sets the "volume_factor" field.
func (cvfc *ChannelVolumeFactorCreate) SetVolumeFactor(i int) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetVolumeFactor(i)
	return cvfc
}

// SetNillableVolumeFactor sets the "volume_factor" field if the given value is not nil.
func (cvfc *ChannelVolumeFactorCreate) SetNillableVolumeFactor(i *int) *ChannelVolumeFactorCreate {
	if i != nil {
		cvfc.SetVolumeFactor(*i)
	}
	return cvfc
}

// SetStatus sets the "status" field.
func (cvfc *ChannelVolumeFactorCreate) SetStatus(i int8) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetStatus(i)
	return cvfc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cvfc *ChannelVolumeFactorCreate) SetNillableStatus(i *int8) *ChannelVolumeFactorCreate {
	if i != nil {
		cvfc.SetStatus(*i)
	}
	return cvfc
}

// SetID sets the "id" field.
func (cvfc *ChannelVolumeFactorCreate) SetID(i int64) *ChannelVolumeFactorCreate {
	cvfc.mutation.SetID(i)
	return cvfc
}

// Mutation returns the ChannelVolumeFactorMutation object of the builder.
func (cvfc *ChannelVolumeFactorCreate) Mutation() *ChannelVolumeFactorMutation {
	return cvfc.mutation
}

// Save creates the ChannelVolumeFactor in the database.
func (cvfc *ChannelVolumeFactorCreate) Save(ctx context.Context) (*ChannelVolumeFactor, error) {
	var (
		err  error
		node *ChannelVolumeFactor
	)
	cvfc.defaults()
	if len(cvfc.hooks) == 0 {
		if err = cvfc.check(); err != nil {
			return nil, err
		}
		node, err = cvfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelVolumeFactorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cvfc.check(); err != nil {
				return nil, err
			}
			cvfc.mutation = mutation
			if node, err = cvfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cvfc.hooks) - 1; i >= 0; i-- {
			if cvfc.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = cvfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cvfc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cvfc *ChannelVolumeFactorCreate) SaveX(ctx context.Context) *ChannelVolumeFactor {
	v, err := cvfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cvfc *ChannelVolumeFactorCreate) Exec(ctx context.Context) error {
	_, err := cvfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvfc *ChannelVolumeFactorCreate) ExecX(ctx context.Context) {
	if err := cvfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cvfc *ChannelVolumeFactorCreate) defaults() {
	if _, ok := cvfc.mutation.CreatedAt(); !ok {
		v := channelvolumefactor.DefaultCreatedAt()
		cvfc.mutation.SetCreatedAt(v)
	}
	if _, ok := cvfc.mutation.UpdatedAt(); !ok {
		v := channelvolumefactor.DefaultUpdatedAt()
		cvfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cvfc.mutation.VolumeFactor(); !ok {
		v := channelvolumefactor.DefaultVolumeFactor
		cvfc.mutation.SetVolumeFactor(v)
	}
	if _, ok := cvfc.mutation.Status(); !ok {
		v := channelvolumefactor.DefaultStatus
		cvfc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cvfc *ChannelVolumeFactorCreate) check() error {
	if _, ok := cvfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`gen: missing required field "ChannelVolumeFactor.created_at"`)}
	}
	if _, ok := cvfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`gen: missing required field "ChannelVolumeFactor.updated_at"`)}
	}
	if _, ok := cvfc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`gen: missing required field "ChannelVolumeFactor.country_code"`)}
	}
	if _, ok := cvfc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`gen: missing required field "ChannelVolumeFactor.channel_id"`)}
	}
	if _, ok := cvfc.mutation.VolumeFactor(); !ok {
		return &ValidationError{Name: "volume_factor", err: errors.New(`gen: missing required field "ChannelVolumeFactor.volume_factor"`)}
	}
	if _, ok := cvfc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`gen: missing required field "ChannelVolumeFactor.status"`)}
	}
	return nil
}

func (cvfc *ChannelVolumeFactorCreate) sqlSave(ctx context.Context) (*ChannelVolumeFactor, error) {
	_node, _spec := cvfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cvfc.driver, _spec); err != nil {
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

func (cvfc *ChannelVolumeFactorCreate) createSpec() (*ChannelVolumeFactor, *sqlgraph.CreateSpec) {
	var (
		_node = &ChannelVolumeFactor{config: cvfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: channelvolumefactor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channelvolumefactor.FieldID,
			},
		}
	)
	_spec.Schema = cvfc.schemaConfig.ChannelVolumeFactor
	_spec.OnConflict = cvfc.conflict
	if id, ok := cvfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cvfc.mutation.CreatedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cvfc.mutation.UpdatedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cvfc.mutation.DeletedAt(); ok {
		_spec.SetField(channelvolumefactor.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cvfc.mutation.CountryCode(); ok {
		_spec.SetField(channelvolumefactor.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = value
	}
	if value, ok := cvfc.mutation.ChannelID(); ok {
		_spec.SetField(channelvolumefactor.FieldChannelID, field.TypeInt64, value)
		_node.ChannelID = value
	}
	if value, ok := cvfc.mutation.VolumeFactor(); ok {
		_spec.SetField(channelvolumefactor.FieldVolumeFactor, field.TypeInt, value)
		_node.VolumeFactor = value
	}
	if value, ok := cvfc.mutation.Status(); ok {
		_spec.SetField(channelvolumefactor.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChannelVolumeFactor.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChannelVolumeFactorUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cvfc *ChannelVolumeFactorCreate) OnConflict(opts ...sql.ConflictOption) *ChannelVolumeFactorUpsertOne {
	cvfc.conflict = opts
	return &ChannelVolumeFactorUpsertOne{
		create: cvfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChannelVolumeFactor.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cvfc *ChannelVolumeFactorCreate) OnConflictColumns(columns ...string) *ChannelVolumeFactorUpsertOne {
	cvfc.conflict = append(cvfc.conflict, sql.ConflictColumns(columns...))
	return &ChannelVolumeFactorUpsertOne{
		create: cvfc,
	}
}

type (
	// ChannelVolumeFactorUpsertOne is the builder for "upsert"-ing
	//  one ChannelVolumeFactor node.
	ChannelVolumeFactorUpsertOne struct {
		create *ChannelVolumeFactorCreate
	}

	// ChannelVolumeFactorUpsert is the "OnConflict" setter.
	ChannelVolumeFactorUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelVolumeFactorUpsert) SetUpdatedAt(v time.Time) *ChannelVolumeFactorUpsert {
	u.Set(channelvolumefactor.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsert) UpdateUpdatedAt() *ChannelVolumeFactorUpsert {
	u.SetExcluded(channelvolumefactor.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChannelVolumeFactorUpsert) SetDeletedAt(v time.Time) *ChannelVolumeFactorUpsert {
	u.Set(channelvolumefactor.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsert) UpdateDeletedAt() *ChannelVolumeFactorUpsert {
	u.SetExcluded(channelvolumefactor.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChannelVolumeFactorUpsert) ClearDeletedAt() *ChannelVolumeFactorUpsert {
	u.SetNull(channelvolumefactor.FieldDeletedAt)
	return u
}

// SetCountryCode sets the "country_code" field.
func (u *ChannelVolumeFactorUpsert) SetCountryCode(v string) *ChannelVolumeFactorUpsert {
	u.Set(channelvolumefactor.FieldCountryCode, v)
	return u
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsert) UpdateCountryCode() *ChannelVolumeFactorUpsert {
	u.SetExcluded(channelvolumefactor.FieldCountryCode)
	return u
}

// SetChannelID sets the "channel_id" field.
func (u *ChannelVolumeFactorUpsert) SetChannelID(v int64) *ChannelVolumeFactorUpsert {
	u.Set(channelvolumefactor.FieldChannelID, v)
	return u
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsert) UpdateChannelID() *ChannelVolumeFactorUpsert {
	u.SetExcluded(channelvolumefactor.FieldChannelID)
	return u
}

// AddChannelID adds v to the "channel_id" field.
func (u *ChannelVolumeFactorUpsert) AddChannelID(v int64) *ChannelVolumeFactorUpsert {
	u.Add(channelvolumefactor.FieldChannelID, v)
	return u
}

// SetVolumeFactor sets the "volume_factor" field.
func (u *ChannelVolumeFactorUpsert) SetVolumeFactor(v int) *ChannelVolumeFactorUpsert {
	u.Set(channelvolumefactor.FieldVolumeFactor, v)
	return u
}

// UpdateVolumeFactor sets the "volume_factor" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsert) UpdateVolumeFactor() *ChannelVolumeFactorUpsert {
	u.SetExcluded(channelvolumefactor.FieldVolumeFactor)
	return u
}

// AddVolumeFactor adds v to the "volume_factor" field.
func (u *ChannelVolumeFactorUpsert) AddVolumeFactor(v int) *ChannelVolumeFactorUpsert {
	u.Add(channelvolumefactor.FieldVolumeFactor, v)
	return u
}

// SetStatus sets the "status" field.
func (u *ChannelVolumeFactorUpsert) SetStatus(v int8) *ChannelVolumeFactorUpsert {
	u.Set(channelvolumefactor.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsert) UpdateStatus() *ChannelVolumeFactorUpsert {
	u.SetExcluded(channelvolumefactor.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *ChannelVolumeFactorUpsert) AddStatus(v int8) *ChannelVolumeFactorUpsert {
	u.Add(channelvolumefactor.FieldStatus, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ChannelVolumeFactor.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(channelvolumefactor.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChannelVolumeFactorUpsertOne) UpdateNewValues() *ChannelVolumeFactorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(channelvolumefactor.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(channelvolumefactor.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChannelVolumeFactor.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChannelVolumeFactorUpsertOne) Ignore() *ChannelVolumeFactorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChannelVolumeFactorUpsertOne) DoNothing() *ChannelVolumeFactorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChannelVolumeFactorCreate.OnConflict
// documentation for more info.
func (u *ChannelVolumeFactorUpsertOne) Update(set func(*ChannelVolumeFactorUpsert)) *ChannelVolumeFactorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChannelVolumeFactorUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelVolumeFactorUpsertOne) SetUpdatedAt(v time.Time) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertOne) UpdateUpdatedAt() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChannelVolumeFactorUpsertOne) SetDeletedAt(v time.Time) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertOne) UpdateDeletedAt() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChannelVolumeFactorUpsertOne) ClearDeletedAt() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.ClearDeletedAt()
	})
}

// SetCountryCode sets the "country_code" field.
func (u *ChannelVolumeFactorUpsertOne) SetCountryCode(v string) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetCountryCode(v)
	})
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertOne) UpdateCountryCode() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateCountryCode()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *ChannelVolumeFactorUpsertOne) SetChannelID(v int64) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetChannelID(v)
	})
}

// AddChannelID adds v to the "channel_id" field.
func (u *ChannelVolumeFactorUpsertOne) AddChannelID(v int64) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.AddChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertOne) UpdateChannelID() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateChannelID()
	})
}

// SetVolumeFactor sets the "volume_factor" field.
func (u *ChannelVolumeFactorUpsertOne) SetVolumeFactor(v int) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetVolumeFactor(v)
	})
}

// AddVolumeFactor adds v to the "volume_factor" field.
func (u *ChannelVolumeFactorUpsertOne) AddVolumeFactor(v int) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.AddVolumeFactor(v)
	})
}

// UpdateVolumeFactor sets the "volume_factor" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertOne) UpdateVolumeFactor() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateVolumeFactor()
	})
}

// SetStatus sets the "status" field.
func (u *ChannelVolumeFactorUpsertOne) SetStatus(v int8) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *ChannelVolumeFactorUpsertOne) AddStatus(v int8) *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertOne) UpdateStatus() *ChannelVolumeFactorUpsertOne {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *ChannelVolumeFactorUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for ChannelVolumeFactorCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChannelVolumeFactorUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChannelVolumeFactorUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChannelVolumeFactorUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ChannelVolumeFactorCreateBulk is the builder for creating many ChannelVolumeFactor entities in bulk.
type ChannelVolumeFactorCreateBulk struct {
	config
	builders []*ChannelVolumeFactorCreate
	conflict []sql.ConflictOption
}

// Save creates the ChannelVolumeFactor entities in the database.
func (cvfcb *ChannelVolumeFactorCreateBulk) Save(ctx context.Context) ([]*ChannelVolumeFactor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cvfcb.builders))
	nodes := make([]*ChannelVolumeFactor, len(cvfcb.builders))
	mutators := make([]Mutator, len(cvfcb.builders))
	for i := range cvfcb.builders {
		func(i int, root context.Context) {
			builder := cvfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelVolumeFactorMutation)
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
					_, err = mutators[i+1].Mutate(root, cvfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cvfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cvfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cvfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cvfcb *ChannelVolumeFactorCreateBulk) SaveX(ctx context.Context) []*ChannelVolumeFactor {
	v, err := cvfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cvfcb *ChannelVolumeFactorCreateBulk) Exec(ctx context.Context) error {
	_, err := cvfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvfcb *ChannelVolumeFactorCreateBulk) ExecX(ctx context.Context) {
	if err := cvfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChannelVolumeFactor.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChannelVolumeFactorUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cvfcb *ChannelVolumeFactorCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChannelVolumeFactorUpsertBulk {
	cvfcb.conflict = opts
	return &ChannelVolumeFactorUpsertBulk{
		create: cvfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChannelVolumeFactor.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cvfcb *ChannelVolumeFactorCreateBulk) OnConflictColumns(columns ...string) *ChannelVolumeFactorUpsertBulk {
	cvfcb.conflict = append(cvfcb.conflict, sql.ConflictColumns(columns...))
	return &ChannelVolumeFactorUpsertBulk{
		create: cvfcb,
	}
}

// ChannelVolumeFactorUpsertBulk is the builder for "upsert"-ing
// a bulk of ChannelVolumeFactor nodes.
type ChannelVolumeFactorUpsertBulk struct {
	create *ChannelVolumeFactorCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ChannelVolumeFactor.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(channelvolumefactor.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChannelVolumeFactorUpsertBulk) UpdateNewValues() *ChannelVolumeFactorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(channelvolumefactor.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(channelvolumefactor.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChannelVolumeFactor.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChannelVolumeFactorUpsertBulk) Ignore() *ChannelVolumeFactorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChannelVolumeFactorUpsertBulk) DoNothing() *ChannelVolumeFactorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChannelVolumeFactorCreateBulk.OnConflict
// documentation for more info.
func (u *ChannelVolumeFactorUpsertBulk) Update(set func(*ChannelVolumeFactorUpsert)) *ChannelVolumeFactorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChannelVolumeFactorUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelVolumeFactorUpsertBulk) SetUpdatedAt(v time.Time) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertBulk) UpdateUpdatedAt() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChannelVolumeFactorUpsertBulk) SetDeletedAt(v time.Time) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertBulk) UpdateDeletedAt() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChannelVolumeFactorUpsertBulk) ClearDeletedAt() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.ClearDeletedAt()
	})
}

// SetCountryCode sets the "country_code" field.
func (u *ChannelVolumeFactorUpsertBulk) SetCountryCode(v string) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetCountryCode(v)
	})
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertBulk) UpdateCountryCode() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateCountryCode()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *ChannelVolumeFactorUpsertBulk) SetChannelID(v int64) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetChannelID(v)
	})
}

// AddChannelID adds v to the "channel_id" field.
func (u *ChannelVolumeFactorUpsertBulk) AddChannelID(v int64) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.AddChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertBulk) UpdateChannelID() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateChannelID()
	})
}

// SetVolumeFactor sets the "volume_factor" field.
func (u *ChannelVolumeFactorUpsertBulk) SetVolumeFactor(v int) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetVolumeFactor(v)
	})
}

// AddVolumeFactor adds v to the "volume_factor" field.
func (u *ChannelVolumeFactorUpsertBulk) AddVolumeFactor(v int) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.AddVolumeFactor(v)
	})
}

// UpdateVolumeFactor sets the "volume_factor" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertBulk) UpdateVolumeFactor() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateVolumeFactor()
	})
}

// SetStatus sets the "status" field.
func (u *ChannelVolumeFactorUpsertBulk) SetStatus(v int8) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *ChannelVolumeFactorUpsertBulk) AddStatus(v int8) *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ChannelVolumeFactorUpsertBulk) UpdateStatus() *ChannelVolumeFactorUpsertBulk {
	return u.Update(func(s *ChannelVolumeFactorUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *ChannelVolumeFactorUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("gen: OnConflict was set for builder %d. Set it on the ChannelVolumeFactorCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for ChannelVolumeFactorCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChannelVolumeFactorUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
