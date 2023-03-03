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
	"github.com/naofel1/api-golang-template/internal/primitive"
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

func TestAuthenticationHandler_SignupStudent(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Sample asset used in working test
	firstname := "Bobby"
	lastname := "Dylan"
	pseudo := "bobDylan"
	gender := "men"
	password := "supertop17"

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

	t.Run("Missing required parameters (pseudo, password...)", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			Tracer:         otel.Tracer("Signup Test"),
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			Logger:         logger,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"pseudo": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockStudentService.AssertNotCalled(t, "Signup")
	})
	t.Run("Invalid pseudo", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			Logger:         logger,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"pseudo":   "bobDylan",
			"password": "supersecret1234",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockStudentService.AssertNotCalled(t, "Signup")
	})
	t.Run("Password too short", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			Tracer:         otel.Tracer("Signup Test"),
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			Logger:         logger,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"firstname": "Bob",
			"lastname":  "Dylan",
			"pseudo":    "bobDylan",
			"gender":    "men",
			"password":  "supe",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockStudentService.AssertNotCalled(t, "Signup")
	})
	t.Run("Password too long", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			Logger:         logger,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"first_name": "Bob",
			"last_name":  "Dylan",
			"pseudo":     "bobDylan",
			"gender":     "men",
			"password":   "super12324jhklafsdjhflkjweyruasdljkfhasdldfjkhasdkljhrleqwwjkrhlqwejrhasdflkjhasdf",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockStudentService.AssertNotCalled(t, "Signup")
	})
	t.Run("Error returned from StudentService", func(t *testing.T) {
		u := &ent.Student{
			FirstName:    firstname,
			LastName:     lastname,
			Pseudo:       pseudo,
			Gender:       primitive.Gender(gender),
			PasswordHash: []byte(password),
		}

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			Logger:         logger,
		})

		mockStudentService.On("Signup", mock.AnythingOfType("*context.valueCtx"), u).
			Return(apistatus.NewConflict("Pseudo Already Exists", u.Pseudo))

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"firstname": firstname,
			"lastname":  lastname,
			"pseudo":    pseudo,
			"gender":    gender,
			"password":  password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 409, rr.Code)
		mockStudentService.AssertExpectations(t)
	})
	t.Run("Successful Token Creation", func(t *testing.T) {
		pseudo = "BobDylan22"

		u := &ent.Student{
			FirstName:    firstname,
			LastName:     lastname,
			Pseudo:       pseudo,
			Gender:       primitive.Gender(gender),
			PasswordHash: []byte(password),
		}

		mockTokenResp := &tokenservice.PairToken{
			IDToken:      &tokenservice.IDToken{SignedToken: "idToken"},
			RefreshToken: &tokenservice.RefreshToken{SignedToken: "refreshToken"},
		}

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			Logger:         logger,
		})

		mockStudentService.
			On("Signup", mock.AnythingOfType("*context.valueCtx"), u).
			Return(nil)
		mockTokenService.
			On("NewPairFromStudent", mock.AnythingOfType("*context.valueCtx"), u, "").
			Return(mockTokenResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"firstname": firstname,
			"lastname":  lastname,
			"pseudo":    pseudo,
			"gender":    gender,
			"password":  password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockStudentService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
	t.Run("Failed Token Creation", func(t *testing.T) {
		u := &ent.Student{
			FirstName:    firstname,
			LastName:     lastname,
			Pseudo:       pseudo,
			Gender:       primitive.Gender(gender),
			PasswordHash: []byte(password),
		}

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		mockStudentService := studentservice.NewMockInterface(t)
		mockTokenService := tokenservice.NewMockInterface(t)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		New(&Config{
			R:              router,
			HostConfig:     conf.Host,
			JwtConfig:      conf.Jwt,
			StudentService: mockStudentService,
			TokenService:   mockTokenService,
			Logger:         logger,
		})

		mockErrorResponse := apistatus.NewInternal()

		mockStudentService.On("Signup", mock.AnythingOfType("*context.valueCtx"), u).
			Return(nil)
		mockTokenService.On("NewPairFromStudent", mock.AnythingOfType("*context.valueCtx"), u, "").
			Return(nil, mockErrorResponse)

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"firstname": firstname,
			"lastname":  lastname,
			"pseudo":    pseudo,
			"gender":    gender,
			"password":  password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/student/register", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(apistatus.NewErrorAPI(mockErrorResponse))
		assert.NoError(t, err)

		assert.Equal(t, mockErrorResponse.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockStudentService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
