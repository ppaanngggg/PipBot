// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ppaanngggg/PipBot/ent/conversation"
	openai "github.com/sashabaranov/go-openai"
)

// ConversationCreate is the builder for creating a Conversation entity.
type ConversationCreate struct {
	config
	mutation *ConversationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (cc *ConversationCreate) SetTitle(s string) *ConversationCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cc *ConversationCreate) SetNillableTitle(s *string) *ConversationCreate {
	if s != nil {
		cc.SetTitle(*s)
	}
	return cc
}

// SetMessages sets the "messages" field.
func (cc *ConversationCreate) SetMessages(ocm []openai.ChatCompletionMessage) *ConversationCreate {
	cc.mutation.SetMessages(ocm)
	return cc
}

// Mutation returns the ConversationMutation object of the builder.
func (cc *ConversationCreate) Mutation() *ConversationMutation {
	return cc.mutation
}

// Save creates the Conversation in the database.
func (cc *ConversationCreate) Save(ctx context.Context) (*Conversation, error) {
	return withHooks[*Conversation, ConversationMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConversationCreate) SaveX(ctx context.Context) *Conversation {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConversationCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConversationCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConversationCreate) check() error {
	if _, ok := cc.mutation.Messages(); !ok {
		return &ValidationError{Name: "messages", err: errors.New(`ent: missing required field "Conversation.messages"`)}
	}
	return nil
}

func (cc *ConversationCreate) sqlSave(ctx context.Context) (*Conversation, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConversationCreate) createSpec() (*Conversation, *sqlgraph.CreateSpec) {
	var (
		_node = &Conversation{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(conversation.Table, sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(conversation.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := cc.mutation.Messages(); ok {
		_spec.SetField(conversation.FieldMessages, field.TypeJSON, value)
		_node.Messages = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Conversation.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ConversationUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (cc *ConversationCreate) OnConflict(opts ...sql.ConflictOption) *ConversationUpsertOne {
	cc.conflict = opts
	return &ConversationUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Conversation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ConversationCreate) OnConflictColumns(columns ...string) *ConversationUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ConversationUpsertOne{
		create: cc,
	}
}

type (
	// ConversationUpsertOne is the builder for "upsert"-ing
	//  one Conversation node.
	ConversationUpsertOne struct {
		create *ConversationCreate
	}

	// ConversationUpsert is the "OnConflict" setter.
	ConversationUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *ConversationUpsert) SetTitle(v string) *ConversationUpsert {
	u.Set(conversation.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ConversationUpsert) UpdateTitle() *ConversationUpsert {
	u.SetExcluded(conversation.FieldTitle)
	return u
}

// ClearTitle clears the value of the "title" field.
func (u *ConversationUpsert) ClearTitle() *ConversationUpsert {
	u.SetNull(conversation.FieldTitle)
	return u
}

// SetMessages sets the "messages" field.
func (u *ConversationUpsert) SetMessages(v []openai.ChatCompletionMessage) *ConversationUpsert {
	u.Set(conversation.FieldMessages, v)
	return u
}

// UpdateMessages sets the "messages" field to the value that was provided on create.
func (u *ConversationUpsert) UpdateMessages() *ConversationUpsert {
	u.SetExcluded(conversation.FieldMessages)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Conversation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ConversationUpsertOne) UpdateNewValues() *ConversationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Conversation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ConversationUpsertOne) Ignore() *ConversationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ConversationUpsertOne) DoNothing() *ConversationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ConversationCreate.OnConflict
// documentation for more info.
func (u *ConversationUpsertOne) Update(set func(*ConversationUpsert)) *ConversationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ConversationUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *ConversationUpsertOne) SetTitle(v string) *ConversationUpsertOne {
	return u.Update(func(s *ConversationUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ConversationUpsertOne) UpdateTitle() *ConversationUpsertOne {
	return u.Update(func(s *ConversationUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *ConversationUpsertOne) ClearTitle() *ConversationUpsertOne {
	return u.Update(func(s *ConversationUpsert) {
		s.ClearTitle()
	})
}

// SetMessages sets the "messages" field.
func (u *ConversationUpsertOne) SetMessages(v []openai.ChatCompletionMessage) *ConversationUpsertOne {
	return u.Update(func(s *ConversationUpsert) {
		s.SetMessages(v)
	})
}

// UpdateMessages sets the "messages" field to the value that was provided on create.
func (u *ConversationUpsertOne) UpdateMessages() *ConversationUpsertOne {
	return u.Update(func(s *ConversationUpsert) {
		s.UpdateMessages()
	})
}

// Exec executes the query.
func (u *ConversationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ConversationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ConversationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ConversationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ConversationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ConversationCreateBulk is the builder for creating many Conversation entities in bulk.
type ConversationCreateBulk struct {
	config
	builders []*ConversationCreate
	conflict []sql.ConflictOption
}

// Save creates the Conversation entities in the database.
func (ccb *ConversationCreateBulk) Save(ctx context.Context) ([]*Conversation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Conversation, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConversationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConversationCreateBulk) SaveX(ctx context.Context) []*Conversation {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConversationCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConversationCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Conversation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ConversationUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (ccb *ConversationCreateBulk) OnConflict(opts ...sql.ConflictOption) *ConversationUpsertBulk {
	ccb.conflict = opts
	return &ConversationUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Conversation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ConversationCreateBulk) OnConflictColumns(columns ...string) *ConversationUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ConversationUpsertBulk{
		create: ccb,
	}
}

// ConversationUpsertBulk is the builder for "upsert"-ing
// a bulk of Conversation nodes.
type ConversationUpsertBulk struct {
	create *ConversationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Conversation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ConversationUpsertBulk) UpdateNewValues() *ConversationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Conversation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ConversationUpsertBulk) Ignore() *ConversationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ConversationUpsertBulk) DoNothing() *ConversationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ConversationCreateBulk.OnConflict
// documentation for more info.
func (u *ConversationUpsertBulk) Update(set func(*ConversationUpsert)) *ConversationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ConversationUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *ConversationUpsertBulk) SetTitle(v string) *ConversationUpsertBulk {
	return u.Update(func(s *ConversationUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ConversationUpsertBulk) UpdateTitle() *ConversationUpsertBulk {
	return u.Update(func(s *ConversationUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *ConversationUpsertBulk) ClearTitle() *ConversationUpsertBulk {
	return u.Update(func(s *ConversationUpsert) {
		s.ClearTitle()
	})
}

// SetMessages sets the "messages" field.
func (u *ConversationUpsertBulk) SetMessages(v []openai.ChatCompletionMessage) *ConversationUpsertBulk {
	return u.Update(func(s *ConversationUpsert) {
		s.SetMessages(v)
	})
}

// UpdateMessages sets the "messages" field to the value that was provided on create.
func (u *ConversationUpsertBulk) UpdateMessages() *ConversationUpsertBulk {
	return u.Update(func(s *ConversationUpsert) {
		s.UpdateMessages()
	})
}

// Exec executes the query.
func (u *ConversationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ConversationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ConversationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ConversationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
