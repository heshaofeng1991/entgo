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
	"github.com/heshaofeng1991/entgo/ent/gen/channeloption"
)

// ChannelOptionCreate is the builder for creating a ChannelOption entity.
type ChannelOptionCreate struct {
	config
	mutation *ChannelOptionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrderID sets the "order_id" field.
func (coc *ChannelOptionCreate) SetOrderID(i int64) *ChannelOptionCreate {
	coc.mutation.SetOrderID(i)
	return coc
}

// SetChannelID sets the "channel_id" field.
func (coc *ChannelOptionCreate) SetChannelID(i int64) *ChannelOptionCreate {
	coc.mutation.SetChannelID(i)
	return coc
}

// SetCountryCode sets the "country_code" field.
func (coc *ChannelOptionCreate) SetCountryCode(s string) *ChannelOptionCreate {
	coc.mutation.SetCountryCode(s)
	return coc
}

// SetCreatedAt sets the "created_at" field.
func (coc *ChannelOptionCreate) SetCreatedAt(t time.Time) *ChannelOptionCreate {
	coc.mutation.SetCreatedAt(t)
	return coc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (coc *ChannelOptionCreate) SetNillableCreatedAt(t *time.Time) *ChannelOptionCreate {
	if t != nil {
		coc.SetCreatedAt(*t)
	}
	return coc
}

// SetUpdatedAt sets the "updated_at" field.
func (coc *ChannelOptionCreate) SetUpdatedAt(t time.Time) *ChannelOptionCreate {
	coc.mutation.SetUpdatedAt(t)
	return coc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (coc *ChannelOptionCreate) SetNillableUpdatedAt(t *time.Time) *ChannelOptionCreate {
	if t != nil {
		coc.SetUpdatedAt(*t)
	}
	return coc
}

// SetDeletedAt sets the "deleted_at" field.
func (coc *ChannelOptionCreate) SetDeletedAt(t time.Time) *ChannelOptionCreate {
	coc.mutation.SetDeletedAt(t)
	return coc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (coc *ChannelOptionCreate) SetNillableDeletedAt(t *time.Time) *ChannelOptionCreate {
	if t != nil {
		coc.SetDeletedAt(*t)
	}
	return coc
}

// SetID sets the "id" field.
func (coc *ChannelOptionCreate) SetID(i int64) *ChannelOptionCreate {
	coc.mutation.SetID(i)
	return coc
}

// Mutation returns the ChannelOptionMutation object of the builder.
func (coc *ChannelOptionCreate) Mutation() *ChannelOptionMutation {
	return coc.mutation
}

// Save creates the ChannelOption in the database.
func (coc *ChannelOptionCreate) Save(ctx context.Context) (*ChannelOption, error) {
	var (
		err  error
		node *ChannelOption
	)
	coc.defaults()
	if len(coc.hooks) == 0 {
		if err = coc.check(); err != nil {
			return nil, err
		}
		node, err = coc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelOptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = coc.check(); err != nil {
				return nil, err
			}
			coc.mutation = mutation
			if node, err = coc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(coc.hooks) - 1; i >= 0; i-- {
			if coc.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = coc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, coc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ChannelOption)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ChannelOptionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (coc *ChannelOptionCreate) SaveX(ctx context.Context) *ChannelOption {
	v, err := coc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (coc *ChannelOptionCreate) Exec(ctx context.Context) error {
	_, err := coc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (coc *ChannelOptionCreate) ExecX(ctx context.Context) {
	if err := coc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (coc *ChannelOptionCreate) defaults() {
	if _, ok := coc.mutation.CreatedAt(); !ok {
		v := channeloption.DefaultCreatedAt()
		coc.mutation.SetCreatedAt(v)
	}
	if _, ok := coc.mutation.UpdatedAt(); !ok {
		v := channeloption.DefaultUpdatedAt()
		coc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (coc *ChannelOptionCreate) check() error {
	if _, ok := coc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`gen: missing required field "ChannelOption.order_id"`)}
	}
	if _, ok := coc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`gen: missing required field "ChannelOption.channel_id"`)}
	}
	if _, ok := coc.mutation.CountryCode(); !ok {
		return &ValidationError{Name: "country_code", err: errors.New(`gen: missing required field "ChannelOption.country_code"`)}
	}
	if _, ok := coc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`gen: missing required field "ChannelOption.created_at"`)}
	}
	if _, ok := coc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`gen: missing required field "ChannelOption.updated_at"`)}
	}
	return nil
}

func (coc *ChannelOptionCreate) sqlSave(ctx context.Context) (*ChannelOption, error) {
	_node, _spec := coc.createSpec()
	if err := sqlgraph.CreateNode(ctx, coc.driver, _spec); err != nil {
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

func (coc *ChannelOptionCreate) createSpec() (*ChannelOption, *sqlgraph.CreateSpec) {
	var (
		_node = &ChannelOption{config: coc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: channeloption.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channeloption.FieldID,
			},
		}
	)
	_spec.Schema = coc.schemaConfig.ChannelOption
	_spec.OnConflict = coc.conflict
	if id, ok := coc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := coc.mutation.OrderID(); ok {
		_spec.SetField(channeloption.FieldOrderID, field.TypeInt64, value)
		_node.OrderID = value
	}
	if value, ok := coc.mutation.ChannelID(); ok {
		_spec.SetField(channeloption.FieldChannelID, field.TypeInt64, value)
		_node.ChannelID = value
	}
	if value, ok := coc.mutation.CountryCode(); ok {
		_spec.SetField(channeloption.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = value
	}
	if value, ok := coc.mutation.CreatedAt(); ok {
		_spec.SetField(channeloption.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := coc.mutation.UpdatedAt(); ok {
		_spec.SetField(channeloption.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := coc.mutation.DeletedAt(); ok {
		_spec.SetField(channeloption.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChannelOption.Create().
//		SetOrderID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChannelOptionUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
func (coc *ChannelOptionCreate) OnConflict(opts ...sql.ConflictOption) *ChannelOptionUpsertOne {
	coc.conflict = opts
	return &ChannelOptionUpsertOne{
		create: coc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChannelOption.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (coc *ChannelOptionCreate) OnConflictColumns(columns ...string) *ChannelOptionUpsertOne {
	coc.conflict = append(coc.conflict, sql.ConflictColumns(columns...))
	return &ChannelOptionUpsertOne{
		create: coc,
	}
}

type (
	// ChannelOptionUpsertOne is the builder for "upsert"-ing
	//  one ChannelOption node.
	ChannelOptionUpsertOne struct {
		create *ChannelOptionCreate
	}

	// ChannelOptionUpsert is the "OnConflict" setter.
	ChannelOptionUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrderID sets the "order_id" field.
func (u *ChannelOptionUpsert) SetOrderID(v int64) *ChannelOptionUpsert {
	u.Set(channeloption.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *ChannelOptionUpsert) UpdateOrderID() *ChannelOptionUpsert {
	u.SetExcluded(channeloption.FieldOrderID)
	return u
}

// AddOrderID adds v to the "order_id" field.
func (u *ChannelOptionUpsert) AddOrderID(v int64) *ChannelOptionUpsert {
	u.Add(channeloption.FieldOrderID, v)
	return u
}

// SetChannelID sets the "channel_id" field.
func (u *ChannelOptionUpsert) SetChannelID(v int64) *ChannelOptionUpsert {
	u.Set(channeloption.FieldChannelID, v)
	return u
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *ChannelOptionUpsert) UpdateChannelID() *ChannelOptionUpsert {
	u.SetExcluded(channeloption.FieldChannelID)
	return u
}

// AddChannelID adds v to the "channel_id" field.
func (u *ChannelOptionUpsert) AddChannelID(v int64) *ChannelOptionUpsert {
	u.Add(channeloption.FieldChannelID, v)
	return u
}

// SetCountryCode sets the "country_code" field.
func (u *ChannelOptionUpsert) SetCountryCode(v string) *ChannelOptionUpsert {
	u.Set(channeloption.FieldCountryCode, v)
	return u
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *ChannelOptionUpsert) UpdateCountryCode() *ChannelOptionUpsert {
	u.SetExcluded(channeloption.FieldCountryCode)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ChannelOptionUpsert) SetCreatedAt(v time.Time) *ChannelOptionUpsert {
	u.Set(channeloption.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ChannelOptionUpsert) UpdateCreatedAt() *ChannelOptionUpsert {
	u.SetExcluded(channeloption.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelOptionUpsert) SetUpdatedAt(v time.Time) *ChannelOptionUpsert {
	u.Set(channeloption.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelOptionUpsert) UpdateUpdatedAt() *ChannelOptionUpsert {
	u.SetExcluded(channeloption.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChannelOptionUpsert) SetDeletedAt(v time.Time) *ChannelOptionUpsert {
	u.Set(channeloption.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChannelOptionUpsert) UpdateDeletedAt() *ChannelOptionUpsert {
	u.SetExcluded(channeloption.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChannelOptionUpsert) ClearDeletedAt() *ChannelOptionUpsert {
	u.SetNull(channeloption.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ChannelOption.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(channeloption.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChannelOptionUpsertOne) UpdateNewValues() *ChannelOptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(channeloption.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChannelOption.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChannelOptionUpsertOne) Ignore() *ChannelOptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChannelOptionUpsertOne) DoNothing() *ChannelOptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChannelOptionCreate.OnConflict
// documentation for more info.
func (u *ChannelOptionUpsertOne) Update(set func(*ChannelOptionUpsert)) *ChannelOptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChannelOptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *ChannelOptionUpsertOne) SetOrderID(v int64) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetOrderID(v)
	})
}

// AddOrderID adds v to the "order_id" field.
func (u *ChannelOptionUpsertOne) AddOrderID(v int64) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.AddOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *ChannelOptionUpsertOne) UpdateOrderID() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateOrderID()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *ChannelOptionUpsertOne) SetChannelID(v int64) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetChannelID(v)
	})
}

// AddChannelID adds v to the "channel_id" field.
func (u *ChannelOptionUpsertOne) AddChannelID(v int64) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.AddChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *ChannelOptionUpsertOne) UpdateChannelID() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateChannelID()
	})
}

// SetCountryCode sets the "country_code" field.
func (u *ChannelOptionUpsertOne) SetCountryCode(v string) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetCountryCode(v)
	})
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *ChannelOptionUpsertOne) UpdateCountryCode() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateCountryCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ChannelOptionUpsertOne) SetCreatedAt(v time.Time) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ChannelOptionUpsertOne) UpdateCreatedAt() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelOptionUpsertOne) SetUpdatedAt(v time.Time) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelOptionUpsertOne) UpdateUpdatedAt() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChannelOptionUpsertOne) SetDeletedAt(v time.Time) *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChannelOptionUpsertOne) UpdateDeletedAt() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChannelOptionUpsertOne) ClearDeletedAt() *ChannelOptionUpsertOne {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ChannelOptionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for ChannelOptionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChannelOptionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChannelOptionUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChannelOptionUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ChannelOptionCreateBulk is the builder for creating many ChannelOption entities in bulk.
type ChannelOptionCreateBulk struct {
	config
	builders []*ChannelOptionCreate
	conflict []sql.ConflictOption
}

// Save creates the ChannelOption entities in the database.
func (cocb *ChannelOptionCreateBulk) Save(ctx context.Context) ([]*ChannelOption, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cocb.builders))
	nodes := make([]*ChannelOption, len(cocb.builders))
	mutators := make([]Mutator, len(cocb.builders))
	for i := range cocb.builders {
		func(i int, root context.Context) {
			builder := cocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelOptionMutation)
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
					_, err = mutators[i+1].Mutate(root, cocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cocb *ChannelOptionCreateBulk) SaveX(ctx context.Context) []*ChannelOption {
	v, err := cocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cocb *ChannelOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := cocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cocb *ChannelOptionCreateBulk) ExecX(ctx context.Context) {
	if err := cocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChannelOption.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChannelOptionUpsert) {
//			SetOrderID(v+v).
//		}).
//		Exec(ctx)
func (cocb *ChannelOptionCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChannelOptionUpsertBulk {
	cocb.conflict = opts
	return &ChannelOptionUpsertBulk{
		create: cocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChannelOption.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cocb *ChannelOptionCreateBulk) OnConflictColumns(columns ...string) *ChannelOptionUpsertBulk {
	cocb.conflict = append(cocb.conflict, sql.ConflictColumns(columns...))
	return &ChannelOptionUpsertBulk{
		create: cocb,
	}
}

// ChannelOptionUpsertBulk is the builder for "upsert"-ing
// a bulk of ChannelOption nodes.
type ChannelOptionUpsertBulk struct {
	create *ChannelOptionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ChannelOption.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(channeloption.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChannelOptionUpsertBulk) UpdateNewValues() *ChannelOptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(channeloption.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChannelOption.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChannelOptionUpsertBulk) Ignore() *ChannelOptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChannelOptionUpsertBulk) DoNothing() *ChannelOptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChannelOptionCreateBulk.OnConflict
// documentation for more info.
func (u *ChannelOptionUpsertBulk) Update(set func(*ChannelOptionUpsert)) *ChannelOptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChannelOptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrderID sets the "order_id" field.
func (u *ChannelOptionUpsertBulk) SetOrderID(v int64) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetOrderID(v)
	})
}

// AddOrderID adds v to the "order_id" field.
func (u *ChannelOptionUpsertBulk) AddOrderID(v int64) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.AddOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *ChannelOptionUpsertBulk) UpdateOrderID() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateOrderID()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *ChannelOptionUpsertBulk) SetChannelID(v int64) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetChannelID(v)
	})
}

// AddChannelID adds v to the "channel_id" field.
func (u *ChannelOptionUpsertBulk) AddChannelID(v int64) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.AddChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *ChannelOptionUpsertBulk) UpdateChannelID() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateChannelID()
	})
}

// SetCountryCode sets the "country_code" field.
func (u *ChannelOptionUpsertBulk) SetCountryCode(v string) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetCountryCode(v)
	})
}

// UpdateCountryCode sets the "country_code" field to the value that was provided on create.
func (u *ChannelOptionUpsertBulk) UpdateCountryCode() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateCountryCode()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ChannelOptionUpsertBulk) SetCreatedAt(v time.Time) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ChannelOptionUpsertBulk) UpdateCreatedAt() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelOptionUpsertBulk) SetUpdatedAt(v time.Time) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelOptionUpsertBulk) UpdateUpdatedAt() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChannelOptionUpsertBulk) SetDeletedAt(v time.Time) *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChannelOptionUpsertBulk) UpdateDeletedAt() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChannelOptionUpsertBulk) ClearDeletedAt() *ChannelOptionUpsertBulk {
	return u.Update(func(s *ChannelOptionUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ChannelOptionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("gen: OnConflict was set for builder %d. Set it on the ChannelOptionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for ChannelOptionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChannelOptionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
