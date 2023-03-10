// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/ent/admin"
	"github.com/naofel1/api-golang-template/internal/ent/schema"
	"github.com/naofel1/api-golang-template/internal/ent/student"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields0[1].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescID is the schema descriptor for id field.
	adminDescID := adminFields[0].Descriptor()
	// admin.DefaultID holds the default value on creation for the id field.
	admin.DefaultID = adminDescID.Default.(func() uuid.UUID)
	studentMixin := schema.Student{}.Mixin()
	studentMixinFields0 := studentMixin[0].Fields()
	_ = studentMixinFields0
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescCreatedAt is the schema descriptor for created_at field.
	studentDescCreatedAt := studentMixinFields0[0].Descriptor()
	// student.DefaultCreatedAt holds the default value on creation for the created_at field.
	student.DefaultCreatedAt = studentDescCreatedAt.Default.(func() time.Time)
	// studentDescUpdatedAt is the schema descriptor for updated_at field.
	studentDescUpdatedAt := studentMixinFields0[1].Descriptor()
	// student.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	student.DefaultUpdatedAt = studentDescUpdatedAt.Default.(func() time.Time)
	// student.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	student.UpdateDefaultUpdatedAt = studentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// studentDescPseudo is the schema descriptor for pseudo field.
	studentDescPseudo := studentFields[3].Descriptor()
	// student.PseudoValidator is a validator for the "pseudo" field. It is called by the builders before save.
	student.PseudoValidator = studentDescPseudo.Validators[0].(func(string) error)
	// studentDescID is the schema descriptor for id field.
	studentDescID := studentFields[0].Descriptor()
	// student.DefaultID holds the default value on creation for the id field.
	student.DefaultID = studentDescID.Default.(func() uuid.UUID)
}
