package user

import (
	"gorm.io/gorm"
	"goyave.dev/goyave/v5"
	"goyave.dev/template/database/model/repository"
	"goyave.dev/template/service"
)

type Service struct {
	db *gorm.DB
	*repository.User
}

func (s *Service) Init(server *goyave.Server) {
	s.db = server.DB()
	s.User = &repository.User{DB: s.db}
}

func (s *Service) Name() string {
	return service.User
}
