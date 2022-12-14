// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/heshaofeng1991/entgo/ent/gen/channelrecommend"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ChannelRecommendDelete is the builder for deleting a ChannelRecommend entity.
type ChannelRecommendDelete struct {
	config
	hooks    []Hook
	mutation *ChannelRecommendMutation
}

// Where appends a list predicates to the ChannelRecommendDelete builder.
func (crd *ChannelRecommendDelete) Where(ps ...predicate.ChannelRecommend) *ChannelRecommendDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *ChannelRecommendDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(crd.hooks) == 0 {
		affected, err = crd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelRecommendMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			crd.mutation = mutation
			affected, err = crd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(crd.hooks) - 1; i >= 0; i-- {
			if crd.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = crd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, crd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *ChannelRecommendDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *ChannelRecommendDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: channelrecommend.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: channelrecommend.FieldID,
			},
		},
	}
	_spec.Node.Schema = crd.schemaConfig.ChannelRecommend
	ctx = internal.NewSchemaConfigContext(ctx, crd.schemaConfig)
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ChannelRecommendDeleteOne is the builder for deleting a single ChannelRecommend entity.
type ChannelRecommendDeleteOne struct {
	crd *ChannelRecommendDelete
}

// Exec executes the deletion query.
func (crdo *ChannelRecommendDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{channelrecommend.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *ChannelRecommendDeleteOne) ExecX(ctx context.Context) {
	crdo.crd.ExecX(ctx)
}
