package initsrv

import (
	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/service/admin/adminservice"
	"github.com/naofel1/api-golang-template/internal/service/mailerservice"
	"github.com/naofel1/api-golang-template/internal/service/student/studentservice"
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
)

// Service will hold service that will be injected
// into this Handdler layer on handler initialization
type Service struct {
	Student studentservice.Interface
	Mailer  mailerservice.Interface
	Token   tokenservice.Interface
	Admin   adminservice.Interface
}

// ServiceConfig contain parameter that will be injected into Service layer
type ServiceConfig struct {
	Logger *otelzap.Logger
	Conf   *configs.Config
	Repo   *Repository
	Utils  *UtilsConfig
}

// InitService will initialize all service needed by application
func InitService(c *ServiceConfig) *Service {
	studentService := studentservice.New(&studentservice.Config{
		Tracer: otel.Tracer("Student Service"),
		Logger: c.Logger,

		StudentRepository: c.Repo.Student,
	})
	mailerService := mailerservice.New(&mailerservice.Config{
		Tracer:     otel.Tracer("Mailer Service"),
		Logger:     c.Logger,
		Conf:       c.Conf.Host,
		MailConfig: c.Conf.Mailgun.MailConfig,

		MailerRepo: c.Repo.Mailer,
	})

	tokenService := tokenservice.New(&tokenservice.Config{
		Tracer:    otel.Tracer("Token Service"),
		Logger:    c.Logger,
		Conf:      c.Conf.Jwt,
		AdminCert: c.Utils.Certs.Admin,
		StudCert:  c.Utils.Certs.Student,

		TokenRepository: c.Repo.Token,
	})
	adminService := adminservice.New(&adminservice.Config{
		Tracer: otel.Tracer("Admin Service"),
		Logger: c.Logger,

		AdminRepository: c.Repo.Admin,
	})

	return &Service{
		Student: studentService,
		Mailer:  mailerService,
		Token:   tokenService,
		Admin:   adminService,
	}
}
