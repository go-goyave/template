package repository

import (
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/template/database/model"
)

type User struct {
	DB *gorm.DB
}

func (r *User) Paginate(page int, pageSize int) (*database.Paginator[*model.User], error) {
	users := []*model.User{}

	paginator := database.NewPaginator(r.DB, page, pageSize, &users)
	result := paginator.Find()
	return paginator, result.Error
}
