package user

import (
	"goyave.dev/goyave/v5/database"
	"goyave.dev/template/database/model"
	"goyave.dev/template/service"
)

// Repository defines the DB functions this service relies on when manipulating users.
type Repository interface {
	First(id int64) (*model.User, error)
	Paginate(page int, pageSize int) (*database.Paginator[*model.User], error)
}

// Service for the user resource.
type Service struct {
	Repository
}

// NewService create a new user Service.
func NewService(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

// Name returns the service name.
func (s *Service) Name() string {
	return service.User
}
