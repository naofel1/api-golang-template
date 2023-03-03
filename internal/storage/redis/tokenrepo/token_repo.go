package tokenrepo

import (
	"github.com/naofel1/api-golang-template/internal/service/token/tokenservice"

	"github.com/redis/go-redis/v9"
)

// Repository is data/repository implementation
// of service layer TokenRepository
type Repository struct {
	Redis *redis.Client
}

// New is a factory for initializing Token Repositories
func New(redisClient *redis.Client) *Repository {
	return &Repository{
		Redis: redisClient,
	}
}

var _ tokenservice.Repository = (*Repository)(nil)
