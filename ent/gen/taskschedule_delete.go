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
	"github.com/heshaofeng1991/entgo/ent/gen/taskschedule"
)

// TaskScheduleDelete is the builder for deleting a TaskSchedule entity.
type TaskScheduleDelete struct {
	config
	hooks    []Hook
	mutation *TaskScheduleMutation
}

// Where appends a list predicates to the TaskScheduleDelete builder.
func (tsd *TaskScheduleDelete) Where(ps ...predicate.TaskSchedule) *TaskScheduleDelete {
	tsd.mutation.Where(ps...)
	return tsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tsd *TaskScheduleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tsd.hooks) == 0 {
		affected, err = tsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tsd.mutation = mutation
			affected, err = tsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tsd.hooks) - 1; i >= 0; i-- {
			if tsd.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = tsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsd *TaskScheduleDelete) ExecX(ctx context.Context) int {
	n, err := tsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tsd *TaskScheduleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: taskschedule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: taskschedule.FieldID,
			},
		},
	}
	_spec.Node.Schema = tsd.schemaConfig.TaskSchedule
	ctx = internal.NewSchemaConfigContext(ctx, tsd.schemaConfig)
	if ps := tsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TaskScheduleDeleteOne is the builder for deleting a single TaskSchedule entity.
type TaskScheduleDeleteOne struct {
	tsd *TaskScheduleDelete
}

// Exec executes the deletion query.
func (tsdo *TaskScheduleDeleteOne) Exec(ctx context.Context) error {
	n, err := tsdo.tsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskschedule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tsdo *TaskScheduleDeleteOne) ExecX(ctx context.Context) {
	tsdo.tsd.ExecX(ctx)
}
