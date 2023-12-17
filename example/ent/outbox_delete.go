// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/igmagollo/undine/example/ent/outbox"
	"github.com/igmagollo/undine/example/ent/predicate"
)

// OutboxDelete is the builder for deleting a Outbox entity.
type OutboxDelete struct {
	config
	hooks    []Hook
	mutation *OutboxMutation
}

// Where appends a list predicates to the OutboxDelete builder.
func (od *OutboxDelete) Where(ps ...predicate.Outbox) *OutboxDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OutboxDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OutboxDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OutboxDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(outbox.Table, sqlgraph.NewFieldSpec(outbox.FieldID, field.TypeUUID))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OutboxDeleteOne is the builder for deleting a single Outbox entity.
type OutboxDeleteOne struct {
	od *OutboxDelete
}

// Where appends a list predicates to the OutboxDelete builder.
func (odo *OutboxDeleteOne) Where(ps ...predicate.Outbox) *OutboxDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OutboxDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{outbox.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OutboxDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}
