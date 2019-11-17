package main

import (
	"github.com/System-Glitch/goyave"
	"goyave_template/http/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	goyave.Start(routes.Register)
}
