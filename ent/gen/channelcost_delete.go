// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/channelcost"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ChannelCostDelete is the builder for deleting a ChannelCost entity.
type ChannelCostDelete struct {
	config
	hooks    []Hook
	mutation *ChannelCostMutation
}

// Where appends a list predicates to the ChannelCostDelete builder.
func (ccd *ChannelCostDelete) Where(ps ...predicate.ChannelCost) *ChannelCostDelete {
	ccd.mutation.Where(ps...)
	return ccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccd *ChannelCostDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccd.hooks) == 0 {
		affected, err = ccd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelCostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccd.mutation = mutation
			affected, err = ccd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccd.hooks) - 1; i >= 0; i-- {
			if ccd.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = ccd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccd *ChannelCostDelete) ExecX(ctx context.Context) int {
	n, err := ccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccd *ChannelCostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: channelcost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channelcost.FieldID,
			},
		},
	}
	_spec.Node.Schema = ccd.schemaConfig.ChannelCost
	ctx = internal.NewSchemaConfigContext(ctx, ccd.schemaConfig)
	if ps := ccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ccd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ChannelCostDeleteOne is the builder for deleting a single ChannelCost entity.
type ChannelCostDeleteOne struct {
	ccd *ChannelCostDelete
}

// Exec executes the deletion query.
func (ccdo *ChannelCostDeleteOne) Exec(ctx context.Context) error {
	n, err := ccdo.ccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{channelcost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccdo *ChannelCostDeleteOne) ExecX(ctx context.Context) {
	ccdo.ccd.ExecX(ctx)
}
