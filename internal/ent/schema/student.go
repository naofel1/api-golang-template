package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/ent/schema/mixin/timemixin"
	"github.com/naofel1/api-golang-template/internal/primitive"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Mixin is used to set specific pattern of common schema
func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timemixin.TimeMixin{},
	}
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("first_name"),
		field.String("last_name"),
		field.String("pseudo").
			MaxLen(20).
			Unique(),
		field.Enum("gender").
			GoType(primitive.Gender("")),
		field.Time("birthday").
			Optional(),
		field.Bytes("password_hash").
			Optional().
			Sensitive(),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{}
}
