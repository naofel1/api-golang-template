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
	/*
		 	Angus: Look into how to run tests in parallel with t.Parallel(). My rules that almost all tests should be
			parallelized, all the time, from day one. As your application grows, your test suite will stay super fast, and
			all the developers on your team will get into the habit of writing thread-safe tests. Parallelizing everything
			also allows the Go race detector to detect certain data races automatically when you run your tests. This is
			much better than discovering these issues in production!

			If you find that running tests in parallel causes race conditions (e.g. because they share mutable state), this
			is usually a sign that the code being tested should be improved!

			The only tests that can't be parallelized are ones that interact with the environment (e.g. setting and
			unsetting env vars).

			Even integration tests can be parallelized, so long as you make sure that all fields with unique constraints
			are randomly generated, and that each test cleans up ONLY the records it creates in the repositories.

			Naofel: Very interesting, I didn't know that. While I have implemented t.Parallel() in certain tests, I hadn't
			adopted it as a standard approach until now 😄. I think when the application grow it can really do the difference!
	*/
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
