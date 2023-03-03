package apistatus

// InvalidArgument is used to help extract validation errors
type InvalidArgument struct {
	Field   string `json:"field"`
	Value   any    `json:"value"`
	Tag     string `json:"tag"`
	Param   string `json:"param"`
	Message string `json:"message"`
}
