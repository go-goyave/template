package model

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"type:char(100)"`
	Email string `gorm:"type:char(100);uniqueIndex"`
}

func UserGenerator() *User {
	user := &User{}
	user.Name = faker.Name()

	faker.SetGenerateUniqueValues(true)
	user.Email = faker.Email()
	faker.SetGenerateUniqueValues(false)
	return user
}
