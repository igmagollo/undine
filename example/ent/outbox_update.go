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
	"github.com/igmagollo/undine/example/ent/outbox"
	"github.com/igmagollo/undine/example/ent/predicate"
)

// OutboxUpdate is the builder for updating Outbox entities.
type OutboxUpdate struct {
	config
	hooks    []Hook
	mutation *OutboxMutation
}

// Where appends a list predicates to the OutboxUpdate builder.
func (ou *OutboxUpdate) Where(ps ...predicate.Outbox) *OutboxUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetHeaders sets the "headers" field.
func (ou *OutboxUpdate) SetHeaders(m map[string]string) *OutboxUpdate {
	ou.mutation.SetHeaders(m)
	return ou
}

// ClearHeaders clears the value of the "headers" field.
func (ou *OutboxUpdate) ClearHeaders() *OutboxUpdate {
	ou.mutation.ClearHeaders()
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OutboxUpdate) SetCreatedAt(t time.Time) *OutboxUpdate {
	ou.mutation.SetCreatedAt(t)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OutboxUpdate) SetNillableCreatedAt(t *time.Time) *OutboxUpdate {
	if t != nil {
		ou.SetCreatedAt(*t)
	}
	return ou
}

// Mutation returns the OutboxMutation object of the builder.
func (ou *OutboxUpdate) Mutation() *OutboxMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OutboxUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OutboxUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OutboxUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OutboxUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OutboxUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(outbox.Table, outbox.Columns, sqlgraph.NewFieldSpec(outbox.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Headers(); ok {
		_spec.SetField(outbox.FieldHeaders, field.TypeJSON, value)
	}
	if ou.mutation.HeadersCleared() {
		_spec.ClearField(outbox.FieldHeaders, field.TypeJSON)
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.SetField(outbox.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OutboxUpdateOne is the builder for updating a single Outbox entity.
type OutboxUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OutboxMutation
}

// SetHeaders sets the "headers" field.
func (ouo *OutboxUpdateOne) SetHeaders(m map[string]string) *OutboxUpdateOne {
	ouo.mutation.SetHeaders(m)
	return ouo
}

// ClearHeaders clears the value of the "headers" field.
func (ouo *OutboxUpdateOne) ClearHeaders() *OutboxUpdateOne {
	ouo.mutation.ClearHeaders()
	return ouo
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OutboxUpdateOne) SetCreatedAt(t time.Time) *OutboxUpdateOne {
	ouo.mutation.SetCreatedAt(t)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OutboxUpdateOne) SetNillableCreatedAt(t *time.Time) *OutboxUpdateOne {
	if t != nil {
		ouo.SetCreatedAt(*t)
	}
	return ouo
}

// Mutation returns the OutboxMutation object of the builder.
func (ouo *OutboxUpdateOne) Mutation() *OutboxMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OutboxUpdate builder.
func (ouo *OutboxUpdateOne) Where(ps ...predicate.Outbox) *OutboxUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OutboxUpdateOne) Select(field string, fields ...string) *OutboxUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Outbox entity.
func (ouo *OutboxUpdateOne) Save(ctx context.Context) (*Outbox, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OutboxUpdateOne) SaveX(ctx context.Context) *Outbox {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OutboxUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OutboxUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OutboxUpdateOne) sqlSave(ctx context.Context) (_node *Outbox, err error) {
	_spec := sqlgraph.NewUpdateSpec(outbox.Table, outbox.Columns, sqlgraph.NewFieldSpec(outbox.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Outbox.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outbox.FieldID)
		for _, f := range fields {
			if !outbox.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != outbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Headers(); ok {
		_spec.SetField(outbox.FieldHeaders, field.TypeJSON, value)
	}
	if ouo.mutation.HeadersCleared() {
		_spec.ClearField(outbox.FieldHeaders, field.TypeJSON)
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.SetField(outbox.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Outbox{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
