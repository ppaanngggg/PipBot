// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myproject/ent/predicate"
	"myproject/ent/settings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// SettingsUpdate is the builder for updating Settings entities.
type SettingsUpdate struct {
	config
	hooks    []Hook
	mutation *SettingsMutation
}

// Where appends a list predicates to the SettingsUpdate builder.
func (su *SettingsUpdate) Where(ps ...predicate.Settings) *SettingsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetKey sets the "key" field.
func (su *SettingsUpdate) SetKey(s string) *SettingsUpdate {
	su.mutation.SetKey(s)
	return su
}

// SetValues sets the "values" field.
func (su *SettingsUpdate) SetValues(s []string) *SettingsUpdate {
	su.mutation.SetValues(s)
	return su
}

// AppendValues appends s to the "values" field.
func (su *SettingsUpdate) AppendValues(s []string) *SettingsUpdate {
	su.mutation.AppendValues(s)
	return su
}

// Mutation returns the SettingsMutation object of the builder.
func (su *SettingsUpdate) Mutation() *SettingsMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SettingsMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SettingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(settings.Table, settings.Columns, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Key(); ok {
		_spec.SetField(settings.FieldKey, field.TypeString, value)
	}
	if value, ok := su.mutation.Values(); ok {
		_spec.SetField(settings.FieldValues, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedValues(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, settings.FieldValues, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{settings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SettingsUpdateOne is the builder for updating a single Settings entity.
type SettingsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SettingsMutation
}

// SetKey sets the "key" field.
func (suo *SettingsUpdateOne) SetKey(s string) *SettingsUpdateOne {
	suo.mutation.SetKey(s)
	return suo
}

// SetValues sets the "values" field.
func (suo *SettingsUpdateOne) SetValues(s []string) *SettingsUpdateOne {
	suo.mutation.SetValues(s)
	return suo
}

// AppendValues appends s to the "values" field.
func (suo *SettingsUpdateOne) AppendValues(s []string) *SettingsUpdateOne {
	suo.mutation.AppendValues(s)
	return suo
}

// Mutation returns the SettingsMutation object of the builder.
func (suo *SettingsUpdateOne) Mutation() *SettingsMutation {
	return suo.mutation
}

// Where appends a list predicates to the SettingsUpdate builder.
func (suo *SettingsUpdateOne) Where(ps ...predicate.Settings) *SettingsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingsUpdateOne) Select(field string, fields ...string) *SettingsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Settings entity.
func (suo *SettingsUpdateOne) Save(ctx context.Context) (*Settings, error) {
	return withHooks[*Settings, SettingsMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingsUpdateOne) SaveX(ctx context.Context) *Settings {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SettingsUpdateOne) sqlSave(ctx context.Context) (_node *Settings, err error) {
	_spec := sqlgraph.NewUpdateSpec(settings.Table, settings.Columns, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Settings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, settings.FieldID)
		for _, f := range fields {
			if !settings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != settings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Key(); ok {
		_spec.SetField(settings.FieldKey, field.TypeString, value)
	}
	if value, ok := suo.mutation.Values(); ok {
		_spec.SetField(settings.FieldValues, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedValues(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, settings.FieldValues, value)
		})
	}
	_node = &Settings{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{settings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
