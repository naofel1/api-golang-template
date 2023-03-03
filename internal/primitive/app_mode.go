package primitive

// AppMode defines the type for application.
type AppMode string

// AppMode values.
const (
	AppModeProd AppMode = "production"
	AppModeDev  AppMode = "development"
)

// String returns the string value for AppMode.
func (ro AppMode) String() string {
	return string(ro)
}
