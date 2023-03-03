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

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func TestAuthenticationHandler_TokensStudent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStudentService := studentservice.NewMockInterface(t)
	mockTokenService := tokenservice.NewMockInterface(t)

	router := gin.Default()

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

	New(&Config{
		R:              router,
		HostConfig:     conf.Host,
		JwtConfig:      conf.Jwt,
		TokenService:   mockTokenService,
		StudentService: mockStudentService,
		Logger:         logger,
	})

	t.Run("Invalid request", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with invalid fields
		reqBody, _ := json.Marshal(gin.H{
			"notRefreshToken": "this key is not valid for this handler!",
		})

		request, _ := http.NewRequest(http.MethodPost, "/student/tokens", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockTokenService.AssertNotCalled(t, "ValidateRefreshToken")
		mockStudentService.AssertNotCalled(t, "GetStudent")
		mockTokenService.AssertNotCalled(t, "NewPairFromStudent")
	})

	t.Run("Invalid token", func(t *testing.T) {
		invalidTokenString := "invalid"
		mockErrorMessage := "authProbs"
		mockError := apistatus.NewAuthorization(mockErrorMessage)

		mockTokenService.
			On("ValidateRefreshToken", mock.AnythingOfType("*context.valueCtx"), invalidTokenString).
			Return(nil, mockError)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with invalid fields
		reqBody, _ := json.Marshal(gin.H{
			"refreshToken": invalidTokenString,
		})

		request, _ := http.NewRequest(http.MethodPost, "/student/tokens", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		respBody, _ := json.Marshal(apistatus.NewErrorAPI(mockError))

		assert.Equal(t, mockError.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockTokenService.AssertCalled(t, "ValidateRefreshToken", mock.AnythingOfType("*context.valueCtx"), invalidTokenString)
		mockStudentService.AssertNotCalled(t, "GetStudent")
		mockTokenService.AssertNotCalled(t, "NewPairFromStudent")
	})

	t.Run("Failure to create new student token pair", func(t *testing.T) {
		validTokenString := "valid"
		mockTokenID, _ := uuid.NewRandom()
		mockStudentID, _ := uuid.NewRandom()

		mockRefreshTokenResp := &tokenservice.RefreshToken{
			SignedToken: validTokenString,
			ID:          mockTokenID,
			UID:         mockStudentID,
		}

		mockTokenService.
			On("ValidateRefreshToken", mock.AnythingOfType("*context.valueCtx"), validTokenString).
			Return(mockRefreshTokenResp, nil)

		mockStudentResp := &ent.Student{
			ID: mockStudentID,
		}
		getArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{
				ID: mockRefreshTokenResp.UID,
			},
		}

		mockStudentService.On("GetStudent", getArgs...).Return(nil)

		mockError := apistatus.NewAuthorization("Invalid refresh token")
		newPairArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			mockStudentResp,
			mockRefreshTokenResp.ID.String(),
		}

		mockTokenService.
			On("NewPairFromStudent", newPairArgs...).
			Return(nil, mockError)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with invalid fields
		reqBody, _ := json.Marshal(gin.H{
			"refreshToken": validTokenString,
		})

		request, _ := http.NewRequest(http.MethodPost, "/student/tokens", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		respBody, _ := json.Marshal(gin.H{
			"error": mockError,
		})

		assert.Equal(t, mockError.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockTokenService.AssertCalled(t, "ValidateRefreshToken", mock.AnythingOfType("*context.valueCtx"), validTokenString)
		mockStudentService.AssertCalled(t, "GetStudent", getArgs...)
		mockTokenService.AssertCalled(t, "NewPairFromStudent", newPairArgs...)
	})

	t.Run("Success", func(t *testing.T) {
		validTokenString := "anothervalid"
		mockTokenID, _ := uuid.NewRandom()
		mockUserID, _ := uuid.NewRandom()

		mockRefreshTokenResp := &tokenservice.RefreshToken{
			SignedToken: validTokenString,
			ID:          mockTokenID,
			UID:         mockUserID,
		}

		mockTokenService.
			On("ValidateRefreshToken", mock.AnythingOfType("*context.valueCtx"), validTokenString).
			Return(mockRefreshTokenResp, nil)

		mockUserResp := &ent.Student{
			ID: mockUserID,
		}
		getArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			&ent.Student{
				ID: mockRefreshTokenResp.UID,
			},
		}

		mockStudentService.
			On("GetStudent", getArgs...).
			Return(nil)

		mockNewTokenID, _ := uuid.NewRandom()
		mockNewUserID, _ := uuid.NewRandom()
		mockTokenPairResp := &tokenservice.PairToken{
			IDToken: &tokenservice.IDToken{SignedToken: "aNewIDToken"},
			RefreshToken: &tokenservice.RefreshToken{
				SignedToken: "aNewRefreshToken",
				ID:          mockNewTokenID,
				UID:         mockNewUserID,
			},
		}

		newPairArgs := mock.Arguments{
			mock.AnythingOfType("*context.valueCtx"),
			mockUserResp,
			mockRefreshTokenResp.ID.String(),
		}

		mockTokenService.
			On("NewPairFromStudent", newPairArgs...).
			Return(mockTokenPairResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with invalid fields
		reqBody, _ := json.Marshal(gin.H{
			"refreshToken": validTokenString,
		})

		request, _ := http.NewRequest(http.MethodPost, "/student/tokens", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		respBody, _ := json.Marshal(tokensResponse{
			Tokens:   mockTokenPairResp.ToFront(),
			Duration: 0,
		})

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockTokenService.AssertCalled(t, "ValidateRefreshToken", mock.AnythingOfType("*context.valueCtx"), validTokenString)
		mockStudentService.AssertCalled(t, "GetStudent", getArgs...)
		mockTokenService.AssertCalled(t, "NewPairFromStudent", newPairArgs...)
	})
}
