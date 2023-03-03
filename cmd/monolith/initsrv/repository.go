package initsrv

import (
	"github.com/naofel1/api-golang-template/internal/storage/mailgun/mailerrepo"
	"github.com/naofel1/api-golang-template/internal/storage/mariadb/admin/adminrepo"
	"github.com/naofel1/api-golang-template/internal/storage/mariadb/student/studentrepo"
	"github.com/naofel1/api-golang-template/internal/storage/redis/tokenrepo"
)

// Repository will hold repository that will be injected
// into this Service layer on service initialization
type Repository struct {
	Token   *tokenrepo.Repository
	Mailer  *mailerrepo.Repository
	Admin   *adminrepo.AtomicRepository
	Student *studentrepo.AtomicRepository
}

// InitRepository will initialize all repository
func InitRepository(conf *Client) *Repository {
	return &Repository{
		Token:   tokenrepo.New(conf.Redis),
		Mailer:  mailerrepo.New(conf.Mailer),
		Admin:   adminrepo.NewAtomic(conf.MariaDB),
		Student: studentrepo.NewAtomic(conf.MariaDB),
	}
}
