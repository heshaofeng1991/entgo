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
	"github.com/heshaofeng1991/entgo/ent/gen/sequence"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
)

// SequenceCreate is the builder for creating a Sequence entity.
type SequenceCreate struct {
	config
	mutation *SequenceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SequenceCreate) SetCreatedAt(t time.Time) *SequenceCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SequenceCreate) SetNillableCreatedAt(t *time.Time) *SequenceCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SequenceCreate) SetUpdatedAt(t time.Time) *SequenceCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SequenceCreate) SetNillableUpdatedAt(t *time.Time) *SequenceCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SequenceCreate) SetDeletedAt(t time.Time) *SequenceCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SequenceCreate) SetNillableDeletedAt(t *time.Time) *SequenceCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SequenceCreate) SetName(s string) *SequenceCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetPrefix sets the "prefix" field.
func (sc *SequenceCreate) SetPrefix(s string) *SequenceCreate {
	sc.mutation.SetPrefix(s)
	return sc
}

// SetNillablePrefix sets the "prefix" field if the given value is not nil.
func (sc *SequenceCreate) SetNillablePrefix(s *string) *SequenceCreate {
	if s != nil {
		sc.SetPrefix(*s)
	}
	return sc
}

// SetValue sets the "value" field.
func (sc *SequenceCreate) SetValue(i int64) *SequenceCreate {
	sc.mutation.SetValue(i)
	return sc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (sc *SequenceCreate) SetNillableValue(i *int64) *SequenceCreate {
	if i != nil {
		sc.SetValue(*i)
	}
	return sc
}

// SetDisplayValue sets the "display_value" field.
func (sc *SequenceCreate) SetDisplayValue(s string) *SequenceCreate {
	sc.mutation.SetDisplayValue(s)
	return sc
}

// SetNillableDisplayValue sets the "display_value" field if the given value is not nil.
func (sc *SequenceCreate) SetNillableDisplayValue(s *string) *SequenceCreate {
	if s != nil {
		sc.SetDisplayValue(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SequenceCreate) SetID(i int64) *SequenceCreate {
	sc.mutation.SetID(i)
	return sc
}

// SetTenantID sets the "tenant" edge to the Tenant entity by ID.
func (sc *SequenceCreate) SetTenantID(id int64) *SequenceCreate {
	sc.mutation.SetTenantID(id)
	return sc
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (sc *SequenceCreate) SetTenant(t *Tenant) *SequenceCreate {
	return sc.SetTenantID(t.ID)
}

// Mutation returns the SequenceMutation object of the builder.
func (sc *SequenceCreate) Mutation() *SequenceMutation {
	return sc.mutation
}

// Save creates the Sequence in the database.
func (sc *SequenceCreate) Save(ctx context.Context) (*Sequence, error) {
	var (
		err  error
		node *Sequence
	)
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SequenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Sequence)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SequenceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SequenceCreate) SaveX(ctx context.Context) *Sequence {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SequenceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SequenceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SequenceCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if sequence.DefaultCreatedAt == nil {
			return fmt.Errorf("gen: uninitialized sequence.DefaultCreatedAt (forgotten import gen/runtime?)")
		}
		v := sequence.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if sequence.DefaultUpdatedAt == nil {
			return fmt.Errorf("gen: uninitialized sequence.DefaultUpdatedAt (forgotten import gen/runtime?)")
		}
		v := sequence.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Prefix(); !ok {
		v := sequence.DefaultPrefix
		sc.mutation.SetPrefix(v)
	}
	if _, ok := sc.mutation.Value(); !ok {
		v := sequence.DefaultValue
		sc.mutation.SetValue(v)
	}
	if _, ok := sc.mutation.DisplayValue(); !ok {
		v := sequence.DefaultDisplayValue
		sc.mutation.SetDisplayValue(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SequenceCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`gen: missing required field "Sequence.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`gen: missing required field "Sequence.updated_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`gen: missing required field "Sequence.name"`)}
	}
	if _, ok := sc.mutation.Prefix(); !ok {
		return &ValidationError{Name: "prefix", err: errors.New(`gen: missing required field "Sequence.prefix"`)}
	}
	if _, ok := sc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`gen: missing required field "Sequence.value"`)}
	}
	if _, ok := sc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant", err: errors.New(`gen: missing required edge "Sequence.tenant"`)}
	}
	return nil
}

func (sc *SequenceCreate) sqlSave(ctx context.Context) (*Sequence, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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

func (sc *SequenceCreate) createSpec() (*Sequence, *sqlgraph.CreateSpec) {
	var (
		_node = &Sequence{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sequence.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sequence.FieldID,
			},
		}
	)
	_spec.Schema = sc.schemaConfig.Sequence
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(sequence.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(sequence.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(sequence.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(sequence.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Prefix(); ok {
		_spec.SetField(sequence.FieldPrefix, field.TypeString, value)
		_node.Prefix = value
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(sequence.FieldValue, field.TypeInt64, value)
		_node.Value = value
	}
	if value, ok := sc.mutation.DisplayValue(); ok {
		_spec.SetField(sequence.FieldDisplayValue, field.TypeString, value)
		_node.DisplayValue = value
	}
	if nodes := sc.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sequence.TenantTable,
			Columns: []string{sequence.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: tenant.FieldID,
				},
			},
		}
		edge.Schema = sc.schemaConfig.Sequence
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sequence_tenant = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Sequence.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SequenceUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sc *SequenceCreate) OnConflict(opts ...sql.ConflictOption) *SequenceUpsertOne {
	sc.conflict = opts
	return &SequenceUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Sequence.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SequenceCreate) OnConflictColumns(columns ...string) *SequenceUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SequenceUpsertOne{
		create: sc,
	}
}

type (
	// SequenceUpsertOne is the builder for "upsert"-ing
	//  one Sequence node.
	SequenceUpsertOne struct {
		create *SequenceCreate
	}

	// SequenceUpsert is the "OnConflict" setter.
	SequenceUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *SequenceUpsert) SetUpdatedAt(v time.Time) *SequenceUpsert {
	u.Set(sequence.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SequenceUpsert) UpdateUpdatedAt() *SequenceUpsert {
	u.SetExcluded(sequence.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SequenceUpsert) SetDeletedAt(v time.Time) *SequenceUpsert {
	u.Set(sequence.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SequenceUpsert) UpdateDeletedAt() *SequenceUpsert {
	u.SetExcluded(sequence.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SequenceUpsert) ClearDeletedAt() *SequenceUpsert {
	u.SetNull(sequence.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *SequenceUpsert) SetName(v string) *SequenceUpsert {
	u.Set(sequence.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SequenceUpsert) UpdateName() *SequenceUpsert {
	u.SetExcluded(sequence.FieldName)
	return u
}

// SetPrefix sets the "prefix" field.
func (u *SequenceUpsert) SetPrefix(v string) *SequenceUpsert {
	u.Set(sequence.FieldPrefix, v)
	return u
}

// UpdatePrefix sets the "prefix" field to the value that was provided on create.
func (u *SequenceUpsert) UpdatePrefix() *SequenceUpsert {
	u.SetExcluded(sequence.FieldPrefix)
	return u
}

// SetValue sets the "value" field.
func (u *SequenceUpsert) SetValue(v int64) *SequenceUpsert {
	u.Set(sequence.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SequenceUpsert) UpdateValue() *SequenceUpsert {
	u.SetExcluded(sequence.FieldValue)
	return u
}

// AddValue adds v to the "value" field.
func (u *SequenceUpsert) AddValue(v int64) *SequenceUpsert {
	u.Add(sequence.FieldValue, v)
	return u
}

// SetDisplayValue sets the "display_value" field.
func (u *SequenceUpsert) SetDisplayValue(v string) *SequenceUpsert {
	u.Set(sequence.FieldDisplayValue, v)
	return u
}

// UpdateDisplayValue sets the "display_value" field to the value that was provided on create.
func (u *SequenceUpsert) UpdateDisplayValue() *SequenceUpsert {
	u.SetExcluded(sequence.FieldDisplayValue)
	return u
}

// ClearDisplayValue clears the value of the "display_value" field.
func (u *SequenceUpsert) ClearDisplayValue() *SequenceUpsert {
	u.SetNull(sequence.FieldDisplayValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Sequence.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sequence.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SequenceUpsertOne) UpdateNewValues() *SequenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sequence.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(sequence.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Sequence.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SequenceUpsertOne) Ignore() *SequenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SequenceUpsertOne) DoNothing() *SequenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SequenceCreate.OnConflict
// documentation for more info.
func (u *SequenceUpsertOne) Update(set func(*SequenceUpsert)) *SequenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SequenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SequenceUpsertOne) SetUpdatedAt(v time.Time) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SequenceUpsertOne) UpdateUpdatedAt() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SequenceUpsertOne) SetDeletedAt(v time.Time) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SequenceUpsertOne) UpdateDeletedAt() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SequenceUpsertOne) ClearDeletedAt() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *SequenceUpsertOne) SetName(v string) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SequenceUpsertOne) UpdateName() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateName()
	})
}

// SetPrefix sets the "prefix" field.
func (u *SequenceUpsertOne) SetPrefix(v string) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.SetPrefix(v)
	})
}

// UpdatePrefix sets the "prefix" field to the value that was provided on create.
func (u *SequenceUpsertOne) UpdatePrefix() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdatePrefix()
	})
}

// SetValue sets the "value" field.
func (u *SequenceUpsertOne) SetValue(v int64) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *SequenceUpsertOne) AddValue(v int64) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SequenceUpsertOne) UpdateValue() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateValue()
	})
}

// SetDisplayValue sets the "display_value" field.
func (u *SequenceUpsertOne) SetDisplayValue(v string) *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.SetDisplayValue(v)
	})
}

// UpdateDisplayValue sets the "display_value" field to the value that was provided on create.
func (u *SequenceUpsertOne) UpdateDisplayValue() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateDisplayValue()
	})
}

// ClearDisplayValue clears the value of the "display_value" field.
func (u *SequenceUpsertOne) ClearDisplayValue() *SequenceUpsertOne {
	return u.Update(func(s *SequenceUpsert) {
		s.ClearDisplayValue()
	})
}

// Exec executes the query.
func (u *SequenceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for SequenceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SequenceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SequenceUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SequenceUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SequenceCreateBulk is the builder for creating many Sequence entities in bulk.
type SequenceCreateBulk struct {
	config
	builders []*SequenceCreate
	conflict []sql.ConflictOption
}

// Save creates the Sequence entities in the database.
func (scb *SequenceCreateBulk) Save(ctx context.Context) ([]*Sequence, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sequence, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SequenceMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SequenceCreateBulk) SaveX(ctx context.Context) []*Sequence {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SequenceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SequenceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Sequence.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SequenceUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (scb *SequenceCreateBulk) OnConflict(opts ...sql.ConflictOption) *SequenceUpsertBulk {
	scb.conflict = opts
	return &SequenceUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Sequence.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SequenceCreateBulk) OnConflictColumns(columns ...string) *SequenceUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SequenceUpsertBulk{
		create: scb,
	}
}

// SequenceUpsertBulk is the builder for "upsert"-ing
// a bulk of Sequence nodes.
type SequenceUpsertBulk struct {
	create *SequenceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Sequence.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sequence.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SequenceUpsertBulk) UpdateNewValues() *SequenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sequence.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(sequence.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Sequence.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SequenceUpsertBulk) Ignore() *SequenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SequenceUpsertBulk) DoNothing() *SequenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SequenceCreateBulk.OnConflict
// documentation for more info.
func (u *SequenceUpsertBulk) Update(set func(*SequenceUpsert)) *SequenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SequenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SequenceUpsertBulk) SetUpdatedAt(v time.Time) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SequenceUpsertBulk) UpdateUpdatedAt() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SequenceUpsertBulk) SetDeletedAt(v time.Time) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SequenceUpsertBulk) UpdateDeletedAt() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SequenceUpsertBulk) ClearDeletedAt() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *SequenceUpsertBulk) SetName(v string) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SequenceUpsertBulk) UpdateName() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateName()
	})
}

// SetPrefix sets the "prefix" field.
func (u *SequenceUpsertBulk) SetPrefix(v string) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.SetPrefix(v)
	})
}

// UpdatePrefix sets the "prefix" field to the value that was provided on create.
func (u *SequenceUpsertBulk) UpdatePrefix() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdatePrefix()
	})
}

// SetValue sets the "value" field.
func (u *SequenceUpsertBulk) SetValue(v int64) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *SequenceUpsertBulk) AddValue(v int64) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SequenceUpsertBulk) UpdateValue() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateValue()
	})
}

// SetDisplayValue sets the "display_value" field.
func (u *SequenceUpsertBulk) SetDisplayValue(v string) *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.SetDisplayValue(v)
	})
}

// UpdateDisplayValue sets the "display_value" field to the value that was provided on create.
func (u *SequenceUpsertBulk) UpdateDisplayValue() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.UpdateDisplayValue()
	})
}

// ClearDisplayValue clears the value of the "display_value" field.
func (u *SequenceUpsertBulk) ClearDisplayValue() *SequenceUpsertBulk {
	return u.Update(func(s *SequenceUpsert) {
		s.ClearDisplayValue()
	})
}

// Exec executes the query.
func (u *SequenceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("gen: OnConflict was set for builder %d. Set it on the SequenceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("gen: missing options for SequenceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SequenceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}