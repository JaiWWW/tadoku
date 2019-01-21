package app

import (
	"github.com/tadoku/api/infra"
	"github.com/tadoku/api/interfaces/rdb"
	r "github.com/tadoku/api/interfaces/repositories"
	"github.com/tadoku/api/usecases"
)

// Repositories is a collection of all repositories
type Repositories struct {
	User usecases.UserRepository
}

// NewRepositories initializes all repositories
func NewRepositories(sh rdb.SQLHandler) *Repositories {
	return &Repositories{
		User: r.NewUserRepository(sh, infra.NewPasswordHasher()),
	}
}
