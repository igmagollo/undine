// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/igmagollo/undine/example/ent/predicate"
	"github.com/igmagollo/undine/example/ent/processedmessage"
)

// ProcessedMessageUpdate is the builder for updating ProcessedMessage entities.
type ProcessedMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ProcessedMessageMutation
}

// Where appends a list predicates to the ProcessedMessageUpdate builder.
func (pmu *ProcessedMessageUpdate) Where(ps ...predicate.ProcessedMessage) *ProcessedMessageUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetCreatedAt sets the "created_at" field.
func (pmu *ProcessedMessageUpdate) SetCreatedAt(t time.Time) *ProcessedMessageUpdate {
	pmu.mutation.SetCreatedAt(t)
	return pmu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmu *ProcessedMessageUpdate) SetNillableCreatedAt(t *time.Time) *ProcessedMessageUpdate {
	if t != nil {
		pmu.SetCreatedAt(*t)
	}
	return pmu
}

// Mutation returns the ProcessedMessageMutation object of the builder.
func (pmu *ProcessedMessageUpdate) Mutation() *ProcessedMessageMutation {
	return pmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *ProcessedMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pmu.sqlSave, pmu.mutation, pmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *ProcessedMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *ProcessedMessageUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *ProcessedMessageUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pmu *ProcessedMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(processedmessage.Table, processedmessage.Columns, sqlgraph.NewFieldSpec(processedmessage.FieldID, field.TypeUUID))
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.CreatedAt(); ok {
		_spec.SetField(processedmessage.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{processedmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pmu.mutation.done = true
	return n, nil
}

// ProcessedMessageUpdateOne is the builder for updating a single ProcessedMessage entity.
type ProcessedMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProcessedMessageMutation
}

// SetCreatedAt sets the "created_at" field.
func (pmuo *ProcessedMessageUpdateOne) SetCreatedAt(t time.Time) *ProcessedMessageUpdateOne {
	pmuo.mutation.SetCreatedAt(t)
	return pmuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmuo *ProcessedMessageUpdateOne) SetNillableCreatedAt(t *time.Time) *ProcessedMessageUpdateOne {
	if t != nil {
		pmuo.SetCreatedAt(*t)
	}
	return pmuo
}

// Mutation returns the ProcessedMessageMutation object of the builder.
func (pmuo *ProcessedMessageUpdateOne) Mutation() *ProcessedMessageMutation {
	return pmuo.mutation
}

// Where appends a list predicates to the ProcessedMessageUpdate builder.
func (pmuo *ProcessedMessageUpdateOne) Where(ps ...predicate.ProcessedMessage) *ProcessedMessageUpdateOne {
	pmuo.mutation.Where(ps...)
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *ProcessedMessageUpdateOne) Select(field string, fields ...string) *ProcessedMessageUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated ProcessedMessage entity.
func (pmuo *ProcessedMessageUpdateOne) Save(ctx context.Context) (*ProcessedMessage, error) {
	return withHooks(ctx, pmuo.sqlSave, pmuo.mutation, pmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *ProcessedMessageUpdateOne) SaveX(ctx context.Context) *ProcessedMessage {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *ProcessedMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *ProcessedMessageUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pmuo *ProcessedMessageUpdateOne) sqlSave(ctx context.Context) (_node *ProcessedMessage, err error) {
	_spec := sqlgraph.NewUpdateSpec(processedmessage.Table, processedmessage.Columns, sqlgraph.NewFieldSpec(processedmessage.FieldID, field.TypeUUID))
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProcessedMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, processedmessage.FieldID)
		for _, f := range fields {
			if !processedmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != processedmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.CreatedAt(); ok {
		_spec.SetField(processedmessage.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &ProcessedMessage{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{processedmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pmuo.mutation.done = true
	return _node, nil
}