package models

import (
	"github.com/System-Glitch/goyave/database"
	"github.com/jinzhu/gorm"
)

func init() {
	database.RegisterModel(&User{})
}

// User represents a user.
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique_index"`
}
