package middlewares

import (
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/naofel1/api-golang-template/pkg/pagination"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// QueryPage contains bound and validated data.
type QueryPage struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// Paginate middleware
func Paginate(logger *otelzap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var querySession QueryPage
		if ok := utils.BindQuery(c, logger, &querySession); !ok {
			return
		}

		// Check if limit is not set
		if querySession.Limit <= 0 || querySession.Limit > 100 {
			querySession.Limit = 10
		}

		// Check if page is not set
		if querySession.Page == 0 {
			querySession.Page = 1
		}

		pag := &pagination.Front{
			CurrentPage:  querySession.Page,
			ItemsPerPage: querySession.Limit,
		}

		// Store the pagination information to use them in handler
		c.Set("pagination", pag)

		// Go to the handler called
		c.Next()
	}
}
