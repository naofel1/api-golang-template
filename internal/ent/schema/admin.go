package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/ent/schema/mixin/timemixin"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Mixin is used to set specific pattern of common schema
func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timemixin.TimeMixin{},
	}
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("pseudo").
			Unique(),
		field.Bytes("password_hash").
			Sensitive(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}
