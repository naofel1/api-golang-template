package primitive

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMen     Gender = "men"
	GenderWomen   Gender = "women"
	GenderNeutral Gender = "neutral"
)

// String returns the string value for Gender.
func (e Gender) String() string {
	return string(e)
}

// Values returns all values for Gender for ent GoType.
func (Gender) Values() (classType []string) {
	for _, r := range []Gender{
		GenderMen,
		GenderWomen,
		GenderNeutral,
	} {
		classType = append(classType, string(r))
	}

	return
}
