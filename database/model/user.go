package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:char(100)"`
	Email     string         `gorm:"type:char(100);uniqueIndex"`
}
