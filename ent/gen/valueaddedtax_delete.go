// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/heshaofeng1991/entgo/ent/gen/valueaddedtax"
)

// ValueAddedTaxDelete is the builder for deleting a ValueAddedTax entity.
type ValueAddedTaxDelete struct {
	config
	hooks    []Hook
	mutation *ValueAddedTaxMutation
}

// Where appends a list predicates to the ValueAddedTaxDelete builder.
func (vatd *ValueAddedTaxDelete) Where(ps ...predicate.ValueAddedTax) *ValueAddedTaxDelete {
	vatd.mutation.Where(ps...)
	return vatd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vatd *ValueAddedTaxDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vatd.hooks) == 0 {
		affected, err = vatd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ValueAddedTaxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vatd.mutation = mutation
			affected, err = vatd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vatd.hooks) - 1; i >= 0; i-- {
			if vatd.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = vatd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vatd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vatd *ValueAddedTaxDelete) ExecX(ctx context.Context) int {
	n, err := vatd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vatd *ValueAddedTaxDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: valueaddedtax.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: valueaddedtax.FieldID,
			},
		},
	}
	_spec.Node.Schema = vatd.schemaConfig.ValueAddedTax
	ctx = internal.NewSchemaConfigContext(ctx, vatd.schemaConfig)
	if ps := vatd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vatd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ValueAddedTaxDeleteOne is the builder for deleting a single ValueAddedTax entity.
type ValueAddedTaxDeleteOne struct {
	vatd *ValueAddedTaxDelete
}

// Exec executes the deletion query.
func (vatdo *ValueAddedTaxDeleteOne) Exec(ctx context.Context) error {
	n, err := vatdo.vatd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{valueaddedtax.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vatdo *ValueAddedTaxDeleteOne) ExecX(ctx context.Context) {
	vatdo.vatd.ExecX(ctx)
}
