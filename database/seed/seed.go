package seed

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/template/database/model"
)

// UserGenerator generates a user with a unique email and a random name.
func UserGenerator() *model.User {
	user := &model.User{}
	user.Name = faker.Name()

	user.Email = faker.Email(
		options.WithGenerateUniqueValues(true),
	)
	return user
}

// Seed migrate all the models and populates the database.
func Seed(db *gorm.DB) {
	if err := Migrate(db); err != nil {
		panic(err)
	}

	userFactory := database.NewFactory(UserGenerator)
	userFactory.Save(db, 10)
}

// Migrate auto-migrate all models.
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return errors.New(err)
	}
	return nil
}
