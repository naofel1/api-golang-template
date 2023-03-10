// Code generated by ent, DO NOT EDIT.

package student

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/primitive"
)

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldPseudo holds the string denoting the pseudo field in the database.
	FieldPseudo = "pseudo"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// Table holds the table name of the student in the database.
	Table = "students"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFirstName,
	FieldLastName,
	FieldPseudo,
	FieldGender,
	FieldBirthday,
	FieldPasswordHash,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// PseudoValidator is a validator for the "pseudo" field. It is called by the builders before save.
	PseudoValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge primitive.Gender) error {
	switch ge.String() {
	case "men", "women", "neutral":
		return nil
	default:
		return fmt.Errorf("student: invalid enum value for gender field: %q", ge)
	}
}
