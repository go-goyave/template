package repository

import (
	"context"

	"gorm.io/gorm"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
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
func (r *User) Paginate(ctx context.Context, page int, pageSize int) (*database.Paginator[*model.User], error) {
	users := []*model.User{}

	paginator := database.NewPaginator(session.DB(ctx, r.DB), page, pageSize, &users)
	err := paginator.Find()
	return paginator, err
}

// First returns the user identified by the given ID, or `nil`
func (r *User) First(ctx context.Context, id uint) (*model.User, error) {
	var user *model.User
	db := session.DB(ctx, r.DB).Where("id", id).First(&user)
	return user, errors.New(db.Error)
}
