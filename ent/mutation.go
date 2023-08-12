// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"movie-app/ent/predicate"
	"movie-app/ent/user"
	"movie-app/ent/userpost"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser     = "User"
	TypeUserPost = "UserPost"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op               Op
	typ              string
	id               *int
	name             *string
	password         *string
	email            *string
	create_time      *time.Time
	clearedFields    map[string]struct{}
	userPosts        map[int]struct{}
	removeduserPosts map[int]struct{}
	cleareduserPosts bool
	done             bool
	oldValue         func(context.Context) (*User, error)
	predicates       []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetCreateTime sets the "create_time" field.
func (m *UserMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *UserMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *UserMutation) ResetCreateTime() {
	m.create_time = nil
}

// AddUserPostIDs adds the "userPosts" edge to the UserPost entity by ids.
func (m *UserMutation) AddUserPostIDs(ids ...int) {
	if m.userPosts == nil {
		m.userPosts = make(map[int]struct{})
	}
	for i := range ids {
		m.userPosts[ids[i]] = struct{}{}
	}
}

// ClearUserPosts clears the "userPosts" edge to the UserPost entity.
func (m *UserMutation) ClearUserPosts() {
	m.cleareduserPosts = true
}

// UserPostsCleared reports if the "userPosts" edge to the UserPost entity was cleared.
func (m *UserMutation) UserPostsCleared() bool {
	return m.cleareduserPosts
}

// RemoveUserPostIDs removes the "userPosts" edge to the UserPost entity by IDs.
func (m *UserMutation) RemoveUserPostIDs(ids ...int) {
	if m.removeduserPosts == nil {
		m.removeduserPosts = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.userPosts, ids[i])
		m.removeduserPosts[ids[i]] = struct{}{}
	}
}

// RemovedUserPosts returns the removed IDs of the "userPosts" edge to the UserPost entity.
func (m *UserMutation) RemovedUserPostsIDs() (ids []int) {
	for id := range m.removeduserPosts {
		ids = append(ids, id)
	}
	return
}

// UserPostsIDs returns the "userPosts" edge IDs in the mutation.
func (m *UserMutation) UserPostsIDs() (ids []int) {
	for id := range m.userPosts {
		ids = append(ids, id)
	}
	return
}

// ResetUserPosts resets all changes to the "userPosts" edge.
func (m *UserMutation) ResetUserPosts() {
	m.userPosts = nil
	m.cleareduserPosts = false
	m.removeduserPosts = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.create_time != nil {
		fields = append(fields, user.FieldCreateTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldPassword:
		return m.Password()
	case user.FieldEmail:
		return m.Email()
	case user.FieldCreateTime:
		return m.CreateTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldCreateTime:
		return m.OldCreateTime(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.userPosts != nil {
		edges = append(edges, user.EdgeUserPosts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeUserPosts:
		ids := make([]ent.Value, 0, len(m.userPosts))
		for id := range m.userPosts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeduserPosts != nil {
		edges = append(edges, user.EdgeUserPosts)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeUserPosts:
		ids := make([]ent.Value, 0, len(m.removeduserPosts))
		for id := range m.removeduserPosts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduserPosts {
		edges = append(edges, user.EdgeUserPosts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeUserPosts:
		return m.cleareduserPosts
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeUserPosts:
		m.ResetUserPosts()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}

// UserPostMutation represents an operation that mutates the UserPost nodes in the graph.
type UserPostMutation struct {
	config
	op             Op
	typ            string
	id             *int
	post_text      *string
	post_images    *map[string]string
	create_time    *time.Time
	clearedFields  map[string]struct{}
	user_id        *int
	cleareduser_id bool
	done           bool
	oldValue       func(context.Context) (*UserPost, error)
	predicates     []predicate.UserPost
}

var _ ent.Mutation = (*UserPostMutation)(nil)

// userpostOption allows management of the mutation configuration using functional options.
type userpostOption func(*UserPostMutation)

// newUserPostMutation creates new mutation for the UserPost entity.
func newUserPostMutation(c config, op Op, opts ...userpostOption) *UserPostMutation {
	m := &UserPostMutation{
		config:        c,
		op:            op,
		typ:           TypeUserPost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserPostID sets the ID field of the mutation.
func withUserPostID(id int) userpostOption {
	return func(m *UserPostMutation) {
		var (
			err   error
			once  sync.Once
			value *UserPost
		)
		m.oldValue = func(ctx context.Context) (*UserPost, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserPost.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserPost sets the old UserPost of the mutation.
func withUserPost(node *UserPost) userpostOption {
	return func(m *UserPostMutation) {
		m.oldValue = func(context.Context) (*UserPost, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserPostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserPostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of UserPost entities.
func (m *UserPostMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserPostMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserPostMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().UserPost.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPostText sets the "post_text" field.
func (m *UserPostMutation) SetPostText(s string) {
	m.post_text = &s
}

// PostText returns the value of the "post_text" field in the mutation.
func (m *UserPostMutation) PostText() (r string, exists bool) {
	v := m.post_text
	if v == nil {
		return
	}
	return *v, true
}

// OldPostText returns the old "post_text" field's value of the UserPost entity.
// If the UserPost object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserPostMutation) OldPostText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPostText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPostText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPostText: %w", err)
	}
	return oldValue.PostText, nil
}

// ResetPostText resets all changes to the "post_text" field.
func (m *UserPostMutation) ResetPostText() {
	m.post_text = nil
}

// SetPostImages sets the "post_images" field.
func (m *UserPostMutation) SetPostImages(value map[string]string) {
	m.post_images = &value
}

// PostImages returns the value of the "post_images" field in the mutation.
func (m *UserPostMutation) PostImages() (r map[string]string, exists bool) {
	v := m.post_images
	if v == nil {
		return
	}
	return *v, true
}

// OldPostImages returns the old "post_images" field's value of the UserPost entity.
// If the UserPost object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserPostMutation) OldPostImages(ctx context.Context) (v map[string]string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPostImages is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPostImages requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPostImages: %w", err)
	}
	return oldValue.PostImages, nil
}

// ResetPostImages resets all changes to the "post_images" field.
func (m *UserPostMutation) ResetPostImages() {
	m.post_images = nil
}

// SetCreateTime sets the "create_time" field.
func (m *UserPostMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *UserPostMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the UserPost entity.
// If the UserPost object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserPostMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *UserPostMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUserIDID sets the "user_id" edge to the User entity by id.
func (m *UserPostMutation) SetUserIDID(id int) {
	m.user_id = &id
}

// ClearUserID clears the "user_id" edge to the User entity.
func (m *UserPostMutation) ClearUserID() {
	m.cleareduser_id = true
}

// UserIDCleared reports if the "user_id" edge to the User entity was cleared.
func (m *UserPostMutation) UserIDCleared() bool {
	return m.cleareduser_id
}

// UserIDID returns the "user_id" edge ID in the mutation.
func (m *UserPostMutation) UserIDID() (id int, exists bool) {
	if m.user_id != nil {
		return *m.user_id, true
	}
	return
}

// UserIDIDs returns the "user_id" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserIDID instead. It exists only for internal usage by the builders.
func (m *UserPostMutation) UserIDIDs() (ids []int) {
	if id := m.user_id; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUserID resets all changes to the "user_id" edge.
func (m *UserPostMutation) ResetUserID() {
	m.user_id = nil
	m.cleareduser_id = false
}

// Where appends a list predicates to the UserPostMutation builder.
func (m *UserPostMutation) Where(ps ...predicate.UserPost) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserPostMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserPostMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.UserPost, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserPostMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserPostMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (UserPost).
func (m *UserPostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserPostMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.post_text != nil {
		fields = append(fields, userpost.FieldPostText)
	}
	if m.post_images != nil {
		fields = append(fields, userpost.FieldPostImages)
	}
	if m.create_time != nil {
		fields = append(fields, userpost.FieldCreateTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserPostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case userpost.FieldPostText:
		return m.PostText()
	case userpost.FieldPostImages:
		return m.PostImages()
	case userpost.FieldCreateTime:
		return m.CreateTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserPostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case userpost.FieldPostText:
		return m.OldPostText(ctx)
	case userpost.FieldPostImages:
		return m.OldPostImages(ctx)
	case userpost.FieldCreateTime:
		return m.OldCreateTime(ctx)
	}
	return nil, fmt.Errorf("unknown UserPost field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserPostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case userpost.FieldPostText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPostText(v)
		return nil
	case userpost.FieldPostImages:
		v, ok := value.(map[string]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPostImages(v)
		return nil
	case userpost.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	}
	return fmt.Errorf("unknown UserPost field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserPostMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserPostMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserPostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserPost numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserPostMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserPostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserPostMutation) ClearField(name string) error {
	return fmt.Errorf("unknown UserPost nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserPostMutation) ResetField(name string) error {
	switch name {
	case userpost.FieldPostText:
		m.ResetPostText()
		return nil
	case userpost.FieldPostImages:
		m.ResetPostImages()
		return nil
	case userpost.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	}
	return fmt.Errorf("unknown UserPost field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserPostMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user_id != nil {
		edges = append(edges, userpost.EdgeUserID)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserPostMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case userpost.EdgeUserID:
		if id := m.user_id; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserPostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserPostMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserPostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser_id {
		edges = append(edges, userpost.EdgeUserID)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserPostMutation) EdgeCleared(name string) bool {
	switch name {
	case userpost.EdgeUserID:
		return m.cleareduser_id
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserPostMutation) ClearEdge(name string) error {
	switch name {
	case userpost.EdgeUserID:
		m.ClearUserID()
		return nil
	}
	return fmt.Errorf("unknown UserPost unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserPostMutation) ResetEdge(name string) error {
	switch name {
	case userpost.EdgeUserID:
		m.ResetUserID()
		return nil
	}
	return fmt.Errorf("unknown UserPost edge %s", name)
}
