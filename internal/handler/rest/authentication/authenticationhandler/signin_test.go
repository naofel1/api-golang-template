package authenticationhandler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/service/student/studentservice"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"go.opentelemetry.io/otel"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func TestAuthenticationHandler_SigninStudent(t *testing.T) {
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
	mockStudentService := studentservice.NewMockInterface(t)
	mockTokenService := tokenservice.NewMockInterface(t)

	router := gin.Default()

	New(&Config{
		R:              router,
		Tracer:         otel.Tracer("Signin Test"),
		StudentService: mockStudentService,
		TokenService:   mockTokenService,
		HostConfig:     conf.Host,
		JwtConfig:      conf.Jwt,
		Logger:         logger,
	})

	t.Run("Bad request data", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with invalid fields
		reqBody, err := json.Marshal(gin.H{
			"login":    "naofel1",
			"password": "",
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/student/login", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockStudentService.AssertNotCalled(t, "Signin")
		mockTokenService.AssertNotCalled(t, "NewTokensFromUser")
	})
	t.Run("Successful Token Creation", func(t *testing.T) {
		login := "bob17"
		password := "pwworksgreat123"

		mockStudentArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{Pseudo: login, PasswordHash: []byte(password)},
		}
		mockStudentService.On("Signin", mockStudentArgs...).Return(nil)

		mockTokenArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{Pseudo: login, PasswordHash: []byte(password)},
			"",
		}

		mockTokenPair := &tokenservice.PairToken{
			IDToken:      &tokenservice.IDToken{SignedToken: "idToken"},
			RefreshToken: &tokenservice.RefreshToken{SignedToken: "refreshToken"},
		}

		mockTokenService.On("NewPairFromStudent", mockTokenArgs...).Return(mockTokenPair, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with valid fields
		reqBody, err := json.Marshal(gin.H{
			"login":    login,
			"password": password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/student/login", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(mockTokenPair.ToFront())
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockStudentService.AssertCalled(t, "Signin", mockStudentArgs...)
		mockTokenService.AssertCalled(t, "NewPairFromStudent", mockTokenArgs...)
	})
	t.Run("Failed Token Creation", func(t *testing.T) {
		login := "Cannotproducetoken"
		password := "Cannotproducetoken"

		mockStudentArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{Pseudo: login, PasswordHash: []byte(password)},
		}

		mockStudentService.On("Signin", mockStudentArgs...).Return(nil)

		mockTokenArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{Pseudo: login, PasswordHash: []byte(password)},
			"",
		}

		mockError := apistatus.NewInternal()
		mockTokenService.On("NewPairFromStudent", mockTokenArgs...).Return(nil, mockError)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with valid fields
		reqBody, err := json.Marshal(gin.H{
			"login":    login,
			"password": password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/student/login", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(apistatus.NewErrorAPI(mockError))
		assert.NoError(t, err)

		assert.Equal(t, mockError.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockStudentService.AssertCalled(t, "Signin", mockStudentArgs...)
		mockTokenService.AssertCalled(t, "NewPairFromStudent", mockTokenArgs...)
	})
	t.Run("Error Returned from Student.Signin", func(t *testing.T) {
		login := "bob17000"
		password := "pwdoesnotmatch123"

		mockStudentArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{Pseudo: login, PasswordHash: []byte(password)},
		}

		// so we can check for a known status code
		mockError := apistatus.NewAuthorization("invalid email/password combo")

		mockStudentService.On("Signin", mockStudentArgs...).Return(mockError)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with valid fields
		reqBody, err := json.Marshal(gin.H{
			"login":    login,
			"password": password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/student/login", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		mockStudentService.AssertCalled(t, "Signin", mockStudentArgs...)
		mockTokenService.AssertNotCalled(t, "NewTokensFromStudent")
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}
