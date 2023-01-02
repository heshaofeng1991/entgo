// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/inbounditem"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// InboundItemDelete is the builder for deleting a InboundItem entity.
type InboundItemDelete struct {
	config
	hooks    []Hook
	mutation *InboundItemMutation
}

// Where appends a list predicates to the InboundItemDelete builder.
func (iid *InboundItemDelete) Where(ps ...predicate.InboundItem) *InboundItemDelete {
	iid.mutation.Where(ps...)
	return iid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (iid *InboundItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iid.hooks) == 0 {
		affected, err = iid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InboundItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iid.mutation = mutation
			affected, err = iid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iid.hooks) - 1; i >= 0; i-- {
			if iid.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = iid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (iid *InboundItemDelete) ExecX(ctx context.Context) int {
	n, err := iid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (iid *InboundItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: inbounditem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: inbounditem.FieldID,
			},
		},
	}
	_spec.Node.Schema = iid.schemaConfig.InboundItem
	ctx = internal.NewSchemaConfigContext(ctx, iid.schemaConfig)
	if ps := iid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, iid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// InboundItemDeleteOne is the builder for deleting a single InboundItem entity.
type InboundItemDeleteOne struct {
	iid *InboundItemDelete
}

// Exec executes the deletion query.
func (iido *InboundItemDeleteOne) Exec(ctx context.Context) error {
	n, err := iido.iid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inbounditem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iido *InboundItemDeleteOne) ExecX(ctx context.Context) {
	iido.iid.ExecX(ctx)
}
