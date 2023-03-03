package gincontext

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naofel1/api-golang-template/internal/ent"
)

// SetAdminToContext add admin to the gin context
func SetAdminToContext(c *gin.Context, adm *ent.Admin) {
	c.Set("admin", adm)
}

// GetAdminFromContext extracts the admin from the gin context
func GetAdminFromContext(c *gin.Context) (*ent.Admin, error) {
	// A *ent.Admin will eventually be added to context in middleware
	adminCtx, exists := c.Get("admin")

	// This shouldn't happen, as our middleware ought to throw an error,
	// this is an extra safety measure.
	if !exists {
		return nil, fmt.Errorf("unable to extract admin from request context")
	}

	// Set admin type
	adm, ok := adminCtx.(*ent.Admin)
	if !ok {
		return nil, fmt.Errorf("error type assertion, not admin type")
	}

	return adm, nil
}
