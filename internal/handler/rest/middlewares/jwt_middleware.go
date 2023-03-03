package middlewares

import (
	"strings"

	"github.com/naofel1/api-golang-template/internal/handler/rest/gincontext"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/naofel1/api-golang-template/internal/utils"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const splitHeader = 2

// authHeader get the Authorization header to take the bearer token
type authHeader struct {
	IDToken string `header:"Authorization"`
}

// AuthStudent middleware check if the student token is valid
func AuthStudent(logger *otelzap.Logger, tracer trace.Tracer, s tokenservice.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request context from gin
		ctx := c.Request.Context()

		// Start to trace middleware
		ctx, span := tracer.Start(ctx, "AuthStudent Middleware")
		defer span.End()

		token, err := c.Cookie("idToken")
		if err != nil {
			logger.Ctx(ctx).Info("No idToken cookie set",
				zap.Error(err),
			)
		}

		if token == "" {
			var h authHeader

			// Bind the header to get Bearer token
			if !utils.BindHeader(c, logger, &h) {
				return
			}

			// Split the token
			idTokenHeader := strings.Split(h.IDToken, "Bearer ")

			// If there are only Bearer with no token, the token is not valid
			if len(idTokenHeader) < splitHeader {
				logger.Ctx(ctx).Info("No token are provided",
					zap.String("Authorization", h.IDToken),
				)

				err := apistatus.NewAuthorization("Must provide Authorization header with format `Bearer {token}`")
				c.AbortWithStatusJSON(err.Status(), apistatus.NewErrorAPI(err))

				return
			}

			token = idTokenHeader[1]
		}

		// Validate the token
		stud, err := s.ValidateStudentIDToken(ctx, token)
		if err != nil {
			logger.Ctx(ctx).Info("Cannot validate Student token",
				zap.Error(err),
			)

			err := apistatus.NewAuthorization("Provided token is invalid")
			c.AbortWithStatusJSON(err.Status(), apistatus.NewErrorAPI(err))

			return
		}

		// Store the student information to use them in handler
		gincontext.SetStudentToContext(c, stud)

		// Go to the handler called
		c.Next()
	}
}

// AuthAdmin middleware check if the admin token is valid
func AuthAdmin(logger *otelzap.Logger, tracer trace.Tracer, s tokenservice.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request context from gin
		ctx := c.Request.Context()

		// Start to trace middleware
		ctx, span := tracer.Start(ctx, "AuthAdmin Middleware")
		defer span.End()

		token, err := c.Cookie("idToken")
		if err != nil {
			logger.Ctx(ctx).Info("No idToken cookie set",
				zap.Error(err),
			)
		}

		if token == "" {
			var h authHeader

			// Bind the header to get Bearer token
			if !utils.BindHeader(c, logger, &h) {
				return
			}

			// Split the token
			idTokenHeader := strings.Split(h.IDToken, "Bearer ")

			// If there are only Bearer with no token, the token is not valid
			if len(idTokenHeader) < splitHeader {
				logger.Ctx(ctx).Info("No token are provided",
					zap.String("Authorization", h.IDToken),
				)

				err := apistatus.NewAuthorization("Must provide Authorization header with format `Bearer {token}`")
				c.AbortWithStatusJSON(err.Status(), apistatus.NewErrorAPI(err))

				return
			}

			token = idTokenHeader[1]
		}

		// Validate the token
		adm, err := s.ValidateAdminIDToken(ctx, token)
		if err != nil {
			logger.Ctx(ctx).Info("Cannot validate Admin token",
				zap.Error(err),
			)

			err := apistatus.NewAuthorization("Provided token is invalid")
			c.AbortWithStatusJSON(err.Status(), apistatus.NewErrorAPI(err))

			return
		}

		// Store the admin information to use them in handler
		gincontext.SetAdminToContext(c, adm)

		// Go to the handler called
		c.Next()
	}
}
