// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/ent/admin"
	"github.com/naofel1/api-golang-template/internal/ent/predicate"
	"github.com/naofel1/api-golang-template/internal/ent/student"
	"github.com/naofel1/api-golang-template/internal/primitive"

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
	TypeAdmin   = "Admin"
	TypeStudent = "Student"
)

// AdminMutation represents an operation that mutates the Admin nodes in the graph.
type AdminMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	pseudo        *string
	password_hash *[]byte
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Admin, error)
	predicates    []predicate.Admin
}

var _ ent.Mutation = (*AdminMutation)(nil)

// adminOption allows management of the mutation configuration using functional options.
type adminOption func(*AdminMutation)

// newAdminMutation creates new mutation for the Admin entity.
func newAdminMutation(c config, op Op, opts ...adminOption) *AdminMutation {
	m := &AdminMutation{
		config:        c,
		op:            op,
		typ:           TypeAdmin,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAdminID sets the ID field of the mutation.
func withAdminID(id uuid.UUID) adminOption {
	return func(m *AdminMutation) {
		var (
			err   error
			once  sync.Once
			value *Admin
		)
		m.oldValue = func(ctx context.Context) (*Admin, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Admin.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAdmin sets the old Admin of the mutation.
func withAdmin(node *Admin) adminOption {
	return func(m *AdminMutation) {
		m.oldValue = func(context.Context) (*Admin, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AdminMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AdminMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Admin entities.
func (m *AdminMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AdminMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AdminMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Admin.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *AdminMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AdminMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AdminMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *AdminMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *AdminMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *AdminMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetPseudo sets the "pseudo" field.
func (m *AdminMutation) SetPseudo(s string) {
	m.pseudo = &s
}

// Pseudo returns the value of the "pseudo" field in the mutation.
func (m *AdminMutation) Pseudo() (r string, exists bool) {
	v := m.pseudo
	if v == nil {
		return
	}
	return *v, true
}

// OldPseudo returns the old "pseudo" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldPseudo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPseudo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPseudo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPseudo: %w", err)
	}
	return oldValue.Pseudo, nil
}

// ResetPseudo resets all changes to the "pseudo" field.
func (m *AdminMutation) ResetPseudo() {
	m.pseudo = nil
}

// SetPasswordHash sets the "password_hash" field.
func (m *AdminMutation) SetPasswordHash(b []byte) {
	m.password_hash = &b
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *AdminMutation) PasswordHash() (r []byte, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldPasswordHash(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *AdminMutation) ResetPasswordHash() {
	m.password_hash = nil
}

// Where appends a list predicates to the AdminMutation builder.
func (m *AdminMutation) Where(ps ...predicate.Admin) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AdminMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AdminMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Admin, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AdminMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AdminMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Admin).
func (m *AdminMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AdminMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, admin.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, admin.FieldUpdatedAt)
	}
	if m.pseudo != nil {
		fields = append(fields, admin.FieldPseudo)
	}
	if m.password_hash != nil {
		fields = append(fields, admin.FieldPasswordHash)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AdminMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case admin.FieldCreatedAt:
		return m.CreatedAt()
	case admin.FieldUpdatedAt:
		return m.UpdatedAt()
	case admin.FieldPseudo:
		return m.Pseudo()
	case admin.FieldPasswordHash:
		return m.PasswordHash()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AdminMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case admin.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case admin.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case admin.FieldPseudo:
		return m.OldPseudo(ctx)
	case admin.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	}
	return nil, fmt.Errorf("unknown Admin field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AdminMutation) SetField(name string, value ent.Value) error {
	switch name {
	case admin.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case admin.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case admin.FieldPseudo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPseudo(v)
		return nil
	case admin.FieldPasswordHash:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	}
	return fmt.Errorf("unknown Admin field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AdminMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AdminMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AdminMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Admin numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AdminMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AdminMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AdminMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Admin nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AdminMutation) ResetField(name string) error {
	switch name {
	case admin.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case admin.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case admin.FieldPseudo:
		m.ResetPseudo()
		return nil
	case admin.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown Admin field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AdminMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AdminMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AdminMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AdminMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AdminMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AdminMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AdminMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Admin unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AdminMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Admin edge %s", name)
}

// StudentMutation represents an operation that mutates the Student nodes in the graph.
type StudentMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	first_name    *string
	last_name     *string
	pseudo        *string
	gender        *primitive.Gender
	birthday      *time.Time
	password_hash *[]byte
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Student, error)
	predicates    []predicate.Student
}

var _ ent.Mutation = (*StudentMutation)(nil)

// studentOption allows management of the mutation configuration using functional options.
type studentOption func(*StudentMutation)

// newStudentMutation creates new mutation for the Student entity.
func newStudentMutation(c config, op Op, opts ...studentOption) *StudentMutation {
	m := &StudentMutation{
		config:        c,
		op:            op,
		typ:           TypeStudent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStudentID sets the ID field of the mutation.
func withStudentID(id uuid.UUID) studentOption {
	return func(m *StudentMutation) {
		var (
			err   error
			once  sync.Once
			value *Student
		)
		m.oldValue = func(ctx context.Context) (*Student, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Student.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStudent sets the old Student of the mutation.
func withStudent(node *Student) studentOption {
	return func(m *StudentMutation) {
		m.oldValue = func(context.Context) (*Student, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StudentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StudentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Student entities.
func (m *StudentMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StudentMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StudentMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Student.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *StudentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *StudentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *StudentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *StudentMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *StudentMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *StudentMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetFirstName sets the "first_name" field.
func (m *StudentMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *StudentMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *StudentMutation) ResetFirstName() {
	m.first_name = nil
}

// SetLastName sets the "last_name" field.
func (m *StudentMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *StudentMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *StudentMutation) ResetLastName() {
	m.last_name = nil
}

// SetPseudo sets the "pseudo" field.
func (m *StudentMutation) SetPseudo(s string) {
	m.pseudo = &s
}

// Pseudo returns the value of the "pseudo" field in the mutation.
func (m *StudentMutation) Pseudo() (r string, exists bool) {
	v := m.pseudo
	if v == nil {
		return
	}
	return *v, true
}

// OldPseudo returns the old "pseudo" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldPseudo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPseudo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPseudo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPseudo: %w", err)
	}
	return oldValue.Pseudo, nil
}

// ResetPseudo resets all changes to the "pseudo" field.
func (m *StudentMutation) ResetPseudo() {
	m.pseudo = nil
}

// SetGender sets the "gender" field.
func (m *StudentMutation) SetGender(pr primitive.Gender) {
	m.gender = &pr
}

// Gender returns the value of the "gender" field in the mutation.
func (m *StudentMutation) Gender() (r primitive.Gender, exists bool) {
	v := m.gender
	if v == nil {
		return
	}
	return *v, true
}

// OldGender returns the old "gender" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldGender(ctx context.Context) (v primitive.Gender, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGender is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGender requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGender: %w", err)
	}
	return oldValue.Gender, nil
}

// ResetGender resets all changes to the "gender" field.
func (m *StudentMutation) ResetGender() {
	m.gender = nil
}

// SetBirthday sets the "birthday" field.
func (m *StudentMutation) SetBirthday(t time.Time) {
	m.birthday = &t
}

// Birthday returns the value of the "birthday" field in the mutation.
func (m *StudentMutation) Birthday() (r time.Time, exists bool) {
	v := m.birthday
	if v == nil {
		return
	}
	return *v, true
}

// OldBirthday returns the old "birthday" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldBirthday(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBirthday is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBirthday requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBirthday: %w", err)
	}
	return oldValue.Birthday, nil
}

// ClearBirthday clears the value of the "birthday" field.
func (m *StudentMutation) ClearBirthday() {
	m.birthday = nil
	m.clearedFields[student.FieldBirthday] = struct{}{}
}

// BirthdayCleared returns if the "birthday" field was cleared in this mutation.
func (m *StudentMutation) BirthdayCleared() bool {
	_, ok := m.clearedFields[student.FieldBirthday]
	return ok
}

// ResetBirthday resets all changes to the "birthday" field.
func (m *StudentMutation) ResetBirthday() {
	m.birthday = nil
	delete(m.clearedFields, student.FieldBirthday)
}

// SetPasswordHash sets the "password_hash" field.
func (m *StudentMutation) SetPasswordHash(b []byte) {
	m.password_hash = &b
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *StudentMutation) PasswordHash() (r []byte, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldPasswordHash(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (m *StudentMutation) ClearPasswordHash() {
	m.password_hash = nil
	m.clearedFields[student.FieldPasswordHash] = struct{}{}
}

// PasswordHashCleared returns if the "password_hash" field was cleared in this mutation.
func (m *StudentMutation) PasswordHashCleared() bool {
	_, ok := m.clearedFields[student.FieldPasswordHash]
	return ok
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *StudentMutation) ResetPasswordHash() {
	m.password_hash = nil
	delete(m.clearedFields, student.FieldPasswordHash)
}

// Where appends a list predicates to the StudentMutation builder.
func (m *StudentMutation) Where(ps ...predicate.Student) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StudentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StudentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Student, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StudentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StudentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Student).
func (m *StudentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StudentMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.created_at != nil {
		fields = append(fields, student.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, student.FieldUpdatedAt)
	}
	if m.first_name != nil {
		fields = append(fields, student.FieldFirstName)
	}
	if m.last_name != nil {
		fields = append(fields, student.FieldLastName)
	}
	if m.pseudo != nil {
		fields = append(fields, student.FieldPseudo)
	}
	if m.gender != nil {
		fields = append(fields, student.FieldGender)
	}
	if m.birthday != nil {
		fields = append(fields, student.FieldBirthday)
	}
	if m.password_hash != nil {
		fields = append(fields, student.FieldPasswordHash)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StudentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case student.FieldCreatedAt:
		return m.CreatedAt()
	case student.FieldUpdatedAt:
		return m.UpdatedAt()
	case student.FieldFirstName:
		return m.FirstName()
	case student.FieldLastName:
		return m.LastName()
	case student.FieldPseudo:
		return m.Pseudo()
	case student.FieldGender:
		return m.Gender()
	case student.FieldBirthday:
		return m.Birthday()
	case student.FieldPasswordHash:
		return m.PasswordHash()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StudentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case student.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case student.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case student.FieldFirstName:
		return m.OldFirstName(ctx)
	case student.FieldLastName:
		return m.OldLastName(ctx)
	case student.FieldPseudo:
		return m.OldPseudo(ctx)
	case student.FieldGender:
		return m.OldGender(ctx)
	case student.FieldBirthday:
		return m.OldBirthday(ctx)
	case student.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	}
	return nil, fmt.Errorf("unknown Student field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case student.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case student.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case student.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case student.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case student.FieldPseudo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPseudo(v)
		return nil
	case student.FieldGender:
		v, ok := value.(primitive.Gender)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGender(v)
		return nil
	case student.FieldBirthday:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBirthday(v)
		return nil
	case student.FieldPasswordHash:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StudentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StudentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Student numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StudentMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(student.FieldBirthday) {
		fields = append(fields, student.FieldBirthday)
	}
	if m.FieldCleared(student.FieldPasswordHash) {
		fields = append(fields, student.FieldPasswordHash)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StudentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StudentMutation) ClearField(name string) error {
	switch name {
	case student.FieldBirthday:
		m.ClearBirthday()
		return nil
	case student.FieldPasswordHash:
		m.ClearPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown Student nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StudentMutation) ResetField(name string) error {
	switch name {
	case student.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case student.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case student.FieldFirstName:
		m.ResetFirstName()
		return nil
	case student.FieldLastName:
		m.ResetLastName()
		return nil
	case student.FieldPseudo:
		m.ResetPseudo()
		return nil
	case student.FieldGender:
		m.ResetGender()
		return nil
	case student.FieldBirthday:
		m.ResetBirthday()
		return nil
	case student.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StudentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StudentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StudentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StudentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StudentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StudentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StudentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Student unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StudentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Student edge %s", name)
}