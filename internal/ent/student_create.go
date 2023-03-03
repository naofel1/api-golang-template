// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/ent/student"
	"github.com/naofel1/api-golang-template/internal/primitive"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *StudentCreate) SetCreatedAt(t time.Time) *StudentCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StudentCreate) SetNillableCreatedAt(t *time.Time) *StudentCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StudentCreate) SetUpdatedAt(t time.Time) *StudentCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StudentCreate) SetNillableUpdatedAt(t *time.Time) *StudentCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetFirstName sets the "first_name" field.
func (sc *StudentCreate) SetFirstName(s string) *StudentCreate {
	sc.mutation.SetFirstName(s)
	return sc
}

// SetLastName sets the "last_name" field.
func (sc *StudentCreate) SetLastName(s string) *StudentCreate {
	sc.mutation.SetLastName(s)
	return sc
}

// SetPseudo sets the "pseudo" field.
func (sc *StudentCreate) SetPseudo(s string) *StudentCreate {
	sc.mutation.SetPseudo(s)
	return sc
}

// SetGender sets the "gender" field.
func (sc *StudentCreate) SetGender(pr primitive.Gender) *StudentCreate {
	sc.mutation.SetGender(pr)
	return sc
}

// SetBirthday sets the "birthday" field.
func (sc *StudentCreate) SetBirthday(t time.Time) *StudentCreate {
	sc.mutation.SetBirthday(t)
	return sc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (sc *StudentCreate) SetNillableBirthday(t *time.Time) *StudentCreate {
	if t != nil {
		sc.SetBirthday(*t)
	}
	return sc
}

// SetPasswordHash sets the "password_hash" field.
func (sc *StudentCreate) SetPasswordHash(b []byte) *StudentCreate {
	sc.mutation.SetPasswordHash(b)
	return sc
}

// SetID sets the "id" field.
func (sc *StudentCreate) SetID(u uuid.UUID) *StudentCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StudentCreate) SetNillableID(u *uuid.UUID) *StudentCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	sc.defaults()
	return withHooks[*Student, StudentMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StudentCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := student.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := student.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := student.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Student.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Student.updated_at"`)}
	}
	if _, ok := sc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Student.first_name"`)}
	}
	if _, ok := sc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Student.last_name"`)}
	}
	if _, ok := sc.mutation.Pseudo(); !ok {
		return &ValidationError{Name: "pseudo", err: errors.New(`ent: missing required field "Student.pseudo"`)}
	}
	if v, ok := sc.mutation.Pseudo(); ok {
		if err := student.PseudoValidator(v); err != nil {
			return &ValidationError{Name: "pseudo", err: fmt.Errorf(`ent: validator failed for field "Student.pseudo": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "Student.gender"`)}
	}
	if v, ok := sc.mutation.Gender(); ok {
		if err := student.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Student.gender": %w`, err)}
		}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(student.Table, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(student.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.FirstName(); ok {
		_spec.SetField(student.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := sc.mutation.LastName(); ok {
		_spec.SetField(student.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := sc.mutation.Pseudo(); ok {
		_spec.SetField(student.FieldPseudo, field.TypeString, value)
		_node.Pseudo = value
	}
	if value, ok := sc.mutation.Gender(); ok {
		_spec.SetField(student.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := sc.mutation.Birthday(); ok {
		_spec.SetField(student.FieldBirthday, field.TypeTime, value)
		_node.Birthday = value
	}
	if value, ok := sc.mutation.PasswordHash(); ok {
		_spec.SetField(student.FieldPasswordHash, field.TypeBytes, value)
		_node.PasswordHash = value
	}
	return _node, _spec
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	builders []*StudentCreate
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
