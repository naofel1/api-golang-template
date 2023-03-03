package gincontext

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naofel1/api-golang-template/internal/ent"
)

// SetStudentToContext add student to the gin context
func SetStudentToContext(c *gin.Context, stud *ent.Student) {
	c.Set("student", stud)
}

// GetStudentFromContext extracts the student from the gin context
func GetStudentFromContext(c *gin.Context) (*ent.Student, error) {
	// A *ent.Student will eventually be added to context in middleware
	studentCtx, exists := c.Get("student")

	// This shouldn't happen, as our middleware ought to throw an error,
	// this is an extra safety measure.
	if !exists {
		return nil, fmt.Errorf("unable to extract student from request context")
	}

	// Set student type
	stud, ok := studentCtx.(*ent.Student)
	if !ok {
		return nil, fmt.Errorf("error type assertion, not student type")
	}

	return stud, nil
}
