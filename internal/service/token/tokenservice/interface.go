package tokenservice

import (
	"context"
	"time"

	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/primitive"

	"github.com/google/uuid"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

// Interface specifies the business operations of the service.
type Interface interface {
	NewPairFromStudent(ctx context.Context, u *ent.Student, prevTokenID string) (*PairToken, error)
	NewPairFromAdmin(ctx context.Context, u *ent.Admin, prevTokenID string) (*PairToken, error)
	Signout(ctx context.Context, uid uuid.UUID) error
	ValidateStudentIDToken(ctx context.Context, tokenString string) (*ent.Student, error)
	ValidateAdminIDToken(ctx context.Context, tokenString string) (*ent.Admin, error)
	GetRoleFromIDToken(tokenString string) (primitive.Roles, error)
	ValidateRefreshToken(ctx context.Context, refreshTokenString string) (*RefreshToken, error)
}

// Config will hold repository and used utils that will be injected
// into this Service layer on service initialization
type Config struct {
	TokenRepository Repository
	StudCert        *configs.CertPair
	AdminCert       *configs.CertPair
	Logger          *otelzap.Logger
	Tracer          trace.Tracer
	Conf            *configs.Jwt
}

// New configures and returns an Interface implementation.
func New(c *Config) Interface {
	return &tokenService{
		TokenRepository:   c.TokenRepository,
		StudCert:          c.StudCert,
		AdminCert:         c.AdminCert,
		RefreshSecret:     c.Conf.RefreshSecret,
		IDExpiration:      c.Conf.TokenDuration,
		RefreshExpiration: c.Conf.RefreshDuration,
		Logger:            c.Logger,
		Tracer:            c.Tracer,
	}
}

// tokenService implements tokenService.Interface.
type tokenService struct {
	TokenRepository   Repository
	StudCert          *configs.CertPair
	AdminCert         *configs.CertPair
	Tracer            trace.Tracer
	Logger            *otelzap.Logger
	RefreshSecret     string
	IDExpiration      time.Duration
	RefreshExpiration time.Duration
}

// Repository contain all the function available in the defined domain
type Repository interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}
