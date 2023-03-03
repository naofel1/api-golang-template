package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/pagination"

	"github.com/google/uuid"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

// Interface specifies the business operations of the service.
type Interface interface {
	GetStudent(ctx context.Context, u *ent.Student) error
	GetStudentsWithSearchWithPagination(ctx context.Context, u *ent.Students, us *ent.Student, pagination *pagination.Front) error
	GetStudentFromLogin(ctx context.Context, u *ent.Student) error
	GetStudentsWithPagination(ctx context.Context, u *ent.Students, pagination *pagination.Front) error

	ModifyProfile(ctx context.Context, u *ent.Student) error

	DeleteStudent(ctx context.Context, u *ent.Student) error

	Signup(ctx context.Context, u *ent.Student) error
	Signin(ctx context.Context, u *ent.Student) error
}

// Config will hold repository and used utils that will be injected
// into this Service layer on service initialization
type Config struct {
	StudentRepository AtomicRepository
	Tracer            trace.Tracer
	Logger            *otelzap.Logger
}

// New configures and returns an Interface implementation.
func New(c *Config) Interface {
	return &studentService{
		StudentRepository: c.StudentRepository,
		Tracer:            c.Tracer,
		Logger:            c.Logger,
	}
}

// userService implements userService.Interface.
type studentService struct {
	StudentRepository AtomicRepository
	Tracer            trace.Tracer
	Logger            *otelzap.Logger
}

// AtomicOperation is the format waited by AtomicRepository
type AtomicOperation func(context.Context, Repository) error

// AtomicRepository will execute the repository with Transaction
type AtomicRepository interface {
	Execute(context.Context, AtomicOperation) error
}

// Repository contain all the function available in the defined domain
type Repository interface {
	CreateStudent(ctx context.Context, u *ent.Student) error
	UpdateStudent(ctx context.Context, u *ent.Student) error
	UpdatePseudoUsernameAndPassword(ctx context.Context, u *ent.Student) error

	GetAllStudentsAtOffset(ctx context.Context, u *ent.Students, off int, limit int) error

	GetTotalStudentCount(ctx context.Context) (int, error)
	GetTotalStudentByPseudoOrFirstNameOrLastNameCount(ctx context.Context, n *ent.Student) (int, error)

	FindStudentByID(ctx context.Context, uid uuid.UUID) (*ent.Student, error)
	FindStudentByLogin(ctx context.Context, login string) (*ent.Student, error)
	FindStudentsByLoginOrFirstNameOrLastNameAtOffset(ctx context.Context, stud *ent.Student, off int, limit int) (ent.Students, error)

	CheckStudentPseudo(ctx context.Context, u *ent.Student) (bool, error)

	DeleteStudent(ctx context.Context, u *ent.Student) error
}
