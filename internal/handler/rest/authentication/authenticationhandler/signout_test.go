package authenticationhandler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"go.opentelemetry.io/otel"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func TestAuthenticationHandler_SignoutStudent(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed init Zap Log: %v", err)
	}
	logger := otelzap.New(l)

	conf := &configs.Config{
		Host: &configs.Host{
			Address: "localhost",
		},
		Jwt: &configs.Jwt{
			TokenDuration:   3600,
			RefreshDuration: 3600,
		},
	}

	// setup mock services, gin engine/router, handler layer
	mockTokenService := tokenservice.NewMockInterface(t)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		ctxStudent := &ent.Student{
			ID: uid,
		}

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// creates a test context for setting a user
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("student", ctxStudent)
		})

		New(&Config{
			R:            router,
			Tracer:       otel.Tracer("Signout Test"),
			TokenService: mockTokenService,
			HostConfig:   conf.Host,
			JwtConfig:    conf.Jwt,
			Logger:       logger,
		})

		mockTokenService.On("Signout", mock.AnythingOfType("*context.valueCtx"), ctxStudent.ID).Return(nil)

		request, _ := http.NewRequest(http.MethodPost, "/student/logout", nil)
		router.ServeHTTP(rr, request)

		respBody, _ := json.Marshal(apistatus.NewSuccessStatus(
			"student signed out successfully!",
		))

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
	})
	t.Run("Signout Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		ctxStudent := &ent.Student{
			ID: uid,
		}

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// creates a test context for setting a student
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("student", ctxStudent)
		})

		mockTokenService.On("Signout", mock.AnythingOfType("*context.valueCtx"), ctxStudent.ID).
			Return(apistatus.NewInternal())

		New(&Config{
			R:            router,
			HostConfig:   conf.Host,
			JwtConfig:    conf.Jwt,
			TokenService: mockTokenService,
			Logger:       logger,
		})

		request, _ := http.NewRequest(http.MethodPost, "/student/logout", nil)
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
