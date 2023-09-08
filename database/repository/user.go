package repository

import (
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/template/database/model"
)

// User repository for user manipulation in the database.
type User struct {
	DB *gorm.DB
}

// NewUser create a new user repository.
func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

// Paginate returns a paginator after executing it.
func (r *User) Paginate(page int, pageSize int) (*database.Paginator[*model.User], error) {
	users := []*model.User{}

	paginator := database.NewPaginator(r.DB, page, pageSize, &users)
	result := paginator.Find()
	return paginator, result.Error
}

// First returns the user identified by the given ID, or `nil`
func (r *User) First(id int64) (*model.User, error) {
	var user *model.User
	db := r.DB.Where("id", id).First(&user)
	var err error
	if db.Error != nil {
		err = errors.New(db.Error)
	}
	return user, err
}
