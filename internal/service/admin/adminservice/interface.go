package adminservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"

	"github.com/google/uuid"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

// Interface specifies the business operations of the service.
/*
	Angus: Now this is very interesting. You've chosen to use ent, an ORM, to model your entities, which is a valid
	choice in many circumstances. It's particularly useful for small, simple applications that you just want to get off
	the ground fast. However, this is closer to an example of the Active Record pattern rather than Hexagonal
	Architecture.

	In hexagonal architecture, handlers and repositories MUST depend ON the domain. The domain is the source of truth.
	With ORMs, the domain depends on the repository layer, and these two layers are tightly coupled. If the
	underlying database table changes, your domain model is forced to change too. But we don't want the domain to have
	any concept of changes in the underlying data source. We should be able to swap the SQL repository for NoSQL, S3, a
	websocket or even standard input, and the domain should never know the difference.

	To achieve this, the domain needs to define its own entities. A repository is an abstraction layer that translates
	the representation used by the storage to the representation required by the domain. If the underlying data source
	changes, it's the repository's responsibility to make sure it is still able to translate the data to the domain
	model.

	My recommendation would be to:
	  1. Define your entities (by hand) in the domain packages. Consider specifying separate types for,
         e.g. the input type SigninRequest and the response type Admin. Using the same mutable instance for both input
	     and return value is likely to lead to unexpected bugs ;)
	  2. For your repository layer, look into using [sqlc](https://github.com/kyleconroy/sqlc) to generate table models
		 and queries for you from raw SQL (you can also write these by hand, but it's quite boring...)
	  3. In the repository layer, implement the domain repository interfaces, translating the sqlc models to your
		 domain models.

	This is a lot of work initially, but it's the gold standard for keeping large applications maintainable.
*/
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
	/*
		Angus: Something to consider: By using a concrete logger, not a logger interface, you tightly couple the
		application to both otel and zap. If you wanted to change either of those in future, you'd have to change every
		single reference throughout the application.

		It's unlikely that you'll migrate away from zap, of course. It's industry-standard and very well supported. It
		also has a distinctive interface that would be very hard to hide behind a Logger interface that you design.
		This is a totally acceptable thing to do, just be aware that that's the trade-off you're making :)

		I'll discuss how we can reduce the impact of this decision in another comment.
	*/
	Logger *otelzap.Logger
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

/*
	Angus: If I recall correctly, the words for "wait" and "expect" are the same in French? They mean quite different
 	things in English!
*/
// AtomicOperation is the format expected by AtomicRepository.
type AtomicOperation func(context.Context, Repository) error

// AtomicRepository will execute the repository with Transaction.
/*
	Angus: I think there's a slight misunderstanding about what atomic repositories are for. That's ok – it's a tricky
	concept!

	The domain defines an AtomicRepository interface when it requires business logic to be executed atomically with the
	storage/retrieval of data. I.e. a transaction should only be committed after some business logic defined by the
	domain succeeds. For example, imagine a banking service where we want to deduct a transaction from a user's account,
	but ONLY if they have enough money in their account to cover it. In the same transaction, we must
		1. Fetch a user's account balance (REPO LOGIC)
		2. Check that the balance is greater than or equal to the transaction amount (BUSINESS LOGIC)
		3. Deduct the transaction amount from the balance (BUSINESS LOGIC)
		4. Save the user (REPO LOGIC).

	We can't do this using the normal Repository pattern, because we can't control when or if the BUSINESS LOGIC steps
	get executed by the repository layer.

	AtomicRepository allows us to define a callback function (AtomicOperation) containing business logic that the
	Repository promises to execute atomically with the REPO LOGIC.

	You DON'T need an AtomicRepository to execute a single REPO LOGIC operation. CreateAdmin, UpdateAdmin, FindAdminByID
	and FindAdminByLogin are all examples of REPO LOGIC operations. They don't need to be executed atomically with
	business logic, so they don't need to be wrapped in an AtomicOperation.

	Under the hood, your repository layer should execute all mutating queries inside a transaction. But the domain layer
	doesn't care about this detail. AtomicRepositories are only important when business logic needs to be executed after
	a transaction has begun, but before it is committed.

	AtomicOperations have a small performance overhead (not relevant to most applications, but it can be for
	high-performance systems), so avoiding them when they're not needed is a good idea.
*/
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
