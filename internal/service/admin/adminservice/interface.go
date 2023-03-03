package adminservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"

	"github.com/google/uuid"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

// Interface specifies the business operations of the service.
type Interface interface {
	Signin(ctx context.Context, u *ent.Admin) error
	Signup(ctx context.Context, u *ent.Admin) error

	ModifyProfile(ctx context.Context, u *ent.Admin) error
	GetAdmin(ctx context.Context, u *ent.Admin) error
}

// Config will hold repository and used utils that will be injected
// into this Service layer on service initialization.
type Config struct {
	AdminRepository AtomicRepository
	Tracer          trace.Tracer
	Logger          *otelzap.Logger
}

// New configures and returns an Interface implementation.
func New(c *Config) Interface {
	return &adminService{
		AdminRepository: c.AdminRepository,
		Tracer:          c.Tracer,
		Logger:          c.Logger,
	}
}

// adminService implements adminService.Interface.
type adminService struct {
	AdminRepository AtomicRepository
	Tracer          trace.Tracer
	Logger          *otelzap.Logger
}

// AtomicOperation is the format waited by AtomicRepository.
type AtomicOperation func(context.Context, Repository) error

// AtomicRepository will execute the repository with Transaction.
type AtomicRepository interface {
	Execute(context.Context, AtomicOperation) error
}

// Repository contain all the function available in the defined domain.
type Repository interface {
	CreateAdmin(ctx context.Context, u *ent.Admin) error
	UpdateAdmin(ctx context.Context, u *ent.Admin) error

	FindAdminByID(ctx context.Context, uid uuid.UUID) (*ent.Admin, error)
	FindAdminByLogin(ctx context.Context, login string) (*ent.Admin, error)
}
