// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"movie-app/ent/predicate"
	"movie-app/ent/userpost"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPostDelete is the builder for deleting a UserPost entity.
type UserPostDelete struct {
	config
	hooks    []Hook
	mutation *UserPostMutation
}

// Where appends a list predicates to the UserPostDelete builder.
func (upd *UserPostDelete) Where(ps ...predicate.UserPost) *UserPostDelete {
	upd.mutation.Where(ps...)
	return upd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upd *UserPostDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, upd.sqlExec, upd.mutation, upd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (upd *UserPostDelete) ExecX(ctx context.Context) int {
	n, err := upd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upd *UserPostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userpost.Table, sqlgraph.NewFieldSpec(userpost.FieldID, field.TypeInt))
	if ps := upd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, upd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	upd.mutation.done = true
	return affected, err
}

// UserPostDeleteOne is the builder for deleting a single UserPost entity.
type UserPostDeleteOne struct {
	upd *UserPostDelete
}

// Where appends a list predicates to the UserPostDelete builder.
func (updo *UserPostDeleteOne) Where(ps ...predicate.UserPost) *UserPostDeleteOne {
	updo.upd.mutation.Where(ps...)
	return updo
}

// Exec executes the deletion query.
func (updo *UserPostDeleteOne) Exec(ctx context.Context) error {
	n, err := updo.upd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userpost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (updo *UserPostDeleteOne) ExecX(ctx context.Context) {
	if err := updo.Exec(ctx); err != nil {
		panic(err)
	}
}
