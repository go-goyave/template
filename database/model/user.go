package models

import (
	"github.com/System-Glitch/goyave/v2/database"
	"github.com/jinzhu/gorm"
)

// A model is a structure reflecting a database table structure. An instance of a model
// is a single database record. Each model is defined in its own file inside the database/models directory.
// Models are usually just normal Golang structs, basic Go types, or pointers of them.
// "sql.Scanner" and "driver.Valuer" interfaces are also supported.

// Learn more here: https://system-glitch.github.io/goyave/guide/basics/database.html#models

func init() {
	// All models should be registered in an "init()" function inside their model file.
	database.RegisterModel(&User{})
}

// User represents a user.
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique_index"`
}
