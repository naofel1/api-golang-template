package primitive

// Roles defines the type for the "roles" enum field.
type Roles string

// Roles values.
const (
	RoleAdmin   Roles = "admin"
	RoleStudent Roles = "student"
)

// String returns the string value for Roles.
func (ro Roles) String() string {
	return string(ro)
}
