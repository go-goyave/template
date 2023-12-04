package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"type:char(100)"`
	Email string `gorm:"type:char(100);uniqueIndex"`
}
