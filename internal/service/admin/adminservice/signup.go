package adminservice

import (
	"context"
	"fmt"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
)

// Signup reaches our AdminRepository to verify the
// pseudo is available and signs up the user if this is the case
func (s *adminService) Signup(ctx context.Context, u *ent.Admin) error {
	/*
		I love that you're tracing! Awesome. Notice how tracing and logging it makes the domain logic quite "noisy"
		though. We can use the Decorator pattern to make this cleaner, a lot like how HTTP middleware works. I've added
		an example at the end of this file!

		Yes you are fully right, it's noisy to deplicate this logic in every service. I will take a look at the Decorator pattern
		Thanks for the example!
	*/
	ctx, span := s.Tracer.Start(ctx, "signup")
	defer span.End()

	pw, err := generatePasswordHash(string(u.PasswordHash))
	if err != nil {
		s.Logger.Ctx(ctx).Error("Unable to signup admin profile",
			zap.String("Pseudo", u.Pseudo),
		)

		return apistatus.NewInternal()
	}

	// now I realize why I originally used Signup(ctx, email, password)
	// then created a user. It's somewhat un-natural to mutate the user here
	/*
		Angus: Agreed! Consider accepting a domain-defined SignupRequest instead. Rather than mutating the original
		input, have the Repository return a new Admin object, and pass that back to the handler.

		Naofel:  Absolutely! Adopting a domain-defined SignupRequest does make the process more streamlined and intelligible.
		By the way, I just have to tell you - I'm really enjoying and learning so much from each comment you share. Your reviews
		are genuinely enlightening and have given me a real grasp of hexagonal architecture implementation. It's been a game changer
		for me. Thanks so much!
	*/
	u.PasswordHash = pw

	// Angus: No need to do this with an AtomicRepository. I say more about this in interface.go :)
	enroll := func(ctx context.Context, repo Repository) error {
		if err := repo.CreateAdmin(ctx, u); err != nil {
			s.Logger.Ctx(ctx).Info("Error when Register User",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)

			if ent.IsConstraintError(err) {
				return apistatus.NewConflict("pseudo", u.Pseudo)
			}

			return apistatus.NewInternal()
		}

		return nil
	}

	if err := s.AdminRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}

/*
Angus: Here's an example of how we can use the Decorator pattern to move noisy tracing and logging out of your
primary service implementation. We are effectively creating middleware for our domain interfaces.
Notice how easy all of this is to test!

Naofel: Wow, that absolutely insane how much cleaner and more efficient this is. I'm definitely going to implement this
in my template. I take a bit of time to understand this chain but that very perfect. Yes it's very easy to test and
to understand. Thanks a lot for this insane example! 🤯

I also have a question about folder structure, do you think that we need to decouple the service in differente files or
keep all in one file? I think that it's better to decouple in different files because with a very big service it
can be difficult to find the good function, but maybe it's better to keep all in one file to have a better overview of
the service. What do you think about that?
*/
type Angus struct {
	name string
}

type CreateAngusRequest struct {
	name string
}

type AngusService interface {
	Create(ctx context.Context, request *CreateAngusRequest) (*Angus, error)
}

type AngusRepository interface {
	Create(ctx context.Context, request *CreateAngusRequest) (*Angus, error)
}

func NewAngusService(repository AngusRepository, tracer trace.Tracer, logger *otelzap.Logger) AngusService {
	// Let's build the middleware chain...
	var service AngusService
	{
		service = &angusService{angusRepository: repository}
		service = &loggingAngusService{next: service, logger: logger}
		service = &tracingAngusService{next: service, tracer: tracer}
	}

	return service
}

type tracingAngusService struct {
	next   AngusService
	tracer trace.Tracer
}

func (s *tracingAngusService) Create(ctx context.Context, request *CreateAngusRequest) (*Angus, error) {
	ctx, span := s.tracer.Start(ctx, "AngusService.Create")
	defer span.End()

	return s.next.Create(ctx, request)
}

type loggingAngusService struct {
	next   AngusService
	logger *otelzap.Logger
}

func (s *loggingAngusService) Create(ctx context.Context, request *CreateAngusRequest) (*Angus, error) {
	angus, err := s.next.Create(ctx, request)
	if err != nil {
		s.logger.Ctx(ctx).Info("Error creating new Angus",
			zap.String("name", request.name),
			zap.Error(err),
		)
	}

	return angus, err
}

// The real service implementation!
type angusService struct {
	angusRepository AngusRepository
}

var ErrNameNotAngus = fmt.Errorf("name must be %q", "Angus")

// Look how short and clean this is.
func (s *angusService) Create(ctx context.Context, request *CreateAngusRequest) (*Angus, error) {
	if request.name != "Angus" {
		return nil, ErrNameNotAngus
	}

	return s.angusRepository.Create(ctx, request)
}
