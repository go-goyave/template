package main

import (
	"goyave_template/http/route"
	_ "goyave_template/http/validation"
	"os"

	"github.com/System-Glitch/goyave/v2"
	// Import the approriate GORM dialect for the database you're using.
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
