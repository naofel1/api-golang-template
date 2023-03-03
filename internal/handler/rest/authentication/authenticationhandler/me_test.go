package authenticationhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/primitive"
	"github.com/naofel1/api-golang-template/internal/service/student/studentservice"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"go.opentelemetry.io/otel"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func TestAuthenticationHandler_MeStudent(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed init Zap Log: %v", err)
	}
	logger := otelzap.New(l)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockStudentResp := &ent.Student{
			ID:        uid,
			FirstName: "Bobby",
			LastName:  "Bobson",
			Gender:    primitive.GenderMen,
			Pseudo:    "bob17",
		}
		getArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			mockStudentResp,
		}

		mockStudentService := studentservice.NewMockInterface(t)
		mockStudentService.On("GetStudent", getArgs...).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// use a middleware to set context for test
		// the only claims we care about in this test
		// is the UID
		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("student", mockStudentResp)
		})

		New(&Config{
			R:              router,
			StudentService: mockStudentService,
			Tracer:         otel.Tracer("Authentication Test"),
			Logger:         logger,
		})

		request, err := http.NewRequest(http.MethodGet, "/student/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(mockStudentResp.ToFront())
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockStudentService.AssertExpectations(t) // assert that StudentService.GetStudent was called
	})
	t.Run("NoContextStudent", func(t *testing.T) {
		mockStudentService := studentservice.NewMockInterface(t)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// do not append student to context
		router := gin.Default()
		New(&Config{
			R:              router,
			Logger:         logger,
			Tracer:         otel.Tracer("Authentication Test"),
			StudentService: mockStudentService,
		})

		request, err := http.NewRequest(http.MethodGet, "/student/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockStudentService.AssertNotCalled(t, "GetStudent", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		getArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{
				ID: uid,
			},
		}

		mockStudentService := studentservice.NewMockInterface(t)
		mockStudentService.On("GetStudent", getArgs...).Return(fmt.Errorf("some error down call chain"))

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("student", &ent.Student{
				ID: uid,
			},
			)
		})

		New(&Config{
			R:              router,
			Logger:         logger,
			Tracer:         otel.Tracer("Authentication Test"),
			StudentService: mockStudentService,
		})

		request, err := http.NewRequest(http.MethodGet, "/student/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respErr := apistatus.NewNotFound("student", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockStudentService.AssertExpectations(t) // assert that StudentService.GetStudent was called
	})
}
