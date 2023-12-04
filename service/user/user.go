package user

import (
	"context"

	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/typeutil"
	"goyave.dev/template/database/model"
	"goyave.dev/template/dto"
	"goyave.dev/template/service"
)

// Repository defines the DB functions this service relies on when manipulating users.
type Repository interface {
	First(ctx context.Context, id uint) (*model.User, error)
	Paginate(ctx context.Context, page int, pageSize int) (*database.Paginator[*model.User], error)
}

// Service for the user resource.
type Service struct {
	repository Repository
}

// NewService create a new user Service.
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// First returns the first user identified by the given ID.
func (s *Service) First(ctx context.Context, id uint) (*dto.User, error) {
	u, err := s.repository.First(ctx, id)
	if err != nil {
		return nil, errors.New(err)
	}
	return typeutil.MustConvert[*dto.User](u), err
}

// Paginate returns a paginator containing all the records that match the given filter request.
func (s *Service) Paginate(ctx context.Context, page, pageSize int) (*database.PaginatorDTO[*dto.User], error) {
	paginator, err := s.repository.Paginate(ctx, page, pageSize)
	if err != nil {
		return nil, errors.New(err)
	}
	return typeutil.MustConvert[*database.PaginatorDTO[*dto.User]](paginator), nil
}

// Name returns the service name.
func (s *Service) Name() string {
	return service.User
}
