// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"movie-app/ent/user"
	"movie-app/ent/userpost"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPostCreate is the builder for creating a UserPost entity.
type UserPostCreate struct {
	config
	mutation *UserPostMutation
	hooks    []Hook
}

// SetPostText sets the "post_text" field.
func (upc *UserPostCreate) SetPostText(s string) *UserPostCreate {
	upc.mutation.SetPostText(s)
	return upc
}

// SetPostImages sets the "post_images" field.
func (upc *UserPostCreate) SetPostImages(m map[string]string) *UserPostCreate {
	upc.mutation.SetPostImages(m)
	return upc
}

// SetCreateTime sets the "create_time" field.
func (upc *UserPostCreate) SetCreateTime(t time.Time) *UserPostCreate {
	upc.mutation.SetCreateTime(t)
	return upc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (upc *UserPostCreate) SetNillableCreateTime(t *time.Time) *UserPostCreate {
	if t != nil {
		upc.SetCreateTime(*t)
	}
	return upc
}

// SetID sets the "id" field.
func (upc *UserPostCreate) SetID(i int) *UserPostCreate {
	upc.mutation.SetID(i)
	return upc
}

// SetUserIDID sets the "user_id" edge to the User entity by ID.
func (upc *UserPostCreate) SetUserIDID(id int) *UserPostCreate {
	upc.mutation.SetUserIDID(id)
	return upc
}

// SetNillableUserIDID sets the "user_id" edge to the User entity by ID if the given value is not nil.
func (upc *UserPostCreate) SetNillableUserIDID(id *int) *UserPostCreate {
	if id != nil {
		upc = upc.SetUserIDID(*id)
	}
	return upc
}

// SetUserID sets the "user_id" edge to the User entity.
func (upc *UserPostCreate) SetUserID(u *User) *UserPostCreate {
	return upc.SetUserIDID(u.ID)
}

// Mutation returns the UserPostMutation object of the builder.
func (upc *UserPostCreate) Mutation() *UserPostMutation {
	return upc.mutation
}

// Save creates the UserPost in the database.
func (upc *UserPostCreate) Save(ctx context.Context) (*UserPost, error) {
	upc.defaults()
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserPostCreate) SaveX(ctx context.Context) *UserPost {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserPostCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserPostCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserPostCreate) defaults() {
	if _, ok := upc.mutation.CreateTime(); !ok {
		v := userpost.DefaultCreateTime
		upc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserPostCreate) check() error {
	if _, ok := upc.mutation.PostText(); !ok {
		return &ValidationError{Name: "post_text", err: errors.New(`ent: missing required field "UserPost.post_text"`)}
	}
	if _, ok := upc.mutation.PostImages(); !ok {
		return &ValidationError{Name: "post_images", err: errors.New(`ent: missing required field "UserPost.post_images"`)}
	}
	if _, ok := upc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "UserPost.create_time"`)}
	}
	if v, ok := upc.mutation.ID(); ok {
		if err := userpost.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "UserPost.id": %w`, err)}
		}
	}
	return nil
}

func (upc *UserPostCreate) sqlSave(ctx context.Context) (*UserPost, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserPostCreate) createSpec() (*UserPost, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPost{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userpost.Table, sqlgraph.NewFieldSpec(userpost.FieldID, field.TypeInt))
	)
	if id, ok := upc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := upc.mutation.PostText(); ok {
		_spec.SetField(userpost.FieldPostText, field.TypeString, value)
		_node.PostText = value
	}
	if value, ok := upc.mutation.PostImages(); ok {
		_spec.SetField(userpost.FieldPostImages, field.TypeJSON, value)
		_node.PostImages = value
	}
	if value, ok := upc.mutation.CreateTime(); ok {
		_spec.SetField(userpost.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if nodes := upc.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userpost.UserIDTable,
			Columns: []string{userpost.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserPostCreateBulk is the builder for creating many UserPost entities in bulk.
type UserPostCreateBulk struct {
	config
	builders []*UserPostCreate
}

// Save creates the UserPost entities in the database.
func (upcb *UserPostCreateBulk) Save(ctx context.Context) ([]*UserPost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserPost, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserPostCreateBulk) SaveX(ctx context.Context) []*UserPost {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserPostCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserPostCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
