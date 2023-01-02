// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/channeloption"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ChannelOptionDelete is the builder for deleting a ChannelOption entity.
type ChannelOptionDelete struct {
	config
	hooks    []Hook
	mutation *ChannelOptionMutation
}

// Where appends a list predicates to the ChannelOptionDelete builder.
func (cod *ChannelOptionDelete) Where(ps ...predicate.ChannelOption) *ChannelOptionDelete {
	cod.mutation.Where(ps...)
	return cod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cod *ChannelOptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cod.hooks) == 0 {
		affected, err = cod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelOptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cod.mutation = mutation
			affected, err = cod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cod.hooks) - 1; i >= 0; i-- {
			if cod.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = cod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cod *ChannelOptionDelete) ExecX(ctx context.Context) int {
	n, err := cod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cod *ChannelOptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: channeloption.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channeloption.FieldID,
			},
		},
	}
	_spec.Node.Schema = cod.schemaConfig.ChannelOption
	ctx = internal.NewSchemaConfigContext(ctx, cod.schemaConfig)
	if ps := cod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ChannelOptionDeleteOne is the builder for deleting a single ChannelOption entity.
type ChannelOptionDeleteOne struct {
	cod *ChannelOptionDelete
}

// Exec executes the deletion query.
func (codo *ChannelOptionDeleteOne) Exec(ctx context.Context) error {
	n, err := codo.cod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{channeloption.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (codo *ChannelOptionDeleteOne) ExecX(ctx context.Context) {
	codo.cod.ExecX(ctx)
}