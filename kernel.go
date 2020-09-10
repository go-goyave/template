package main

import (
	"goyave_template/http/route"
	_ "goyave_template/http/validation"
	"os"

	"github.com/System-Glitch/goyave/v3"
	// Import the approriate GORM dialect for the database you're using.
	// _ "github.com/System-Glitch/goyave/v3/database/dialect/mysql"
	// _ "github.com/System-Glitch/goyave/v3/database/dialect/postgres"
	// _ "github.com/System-Glitch/goyave/v3/database/dialect/sqlite"
	// _ "github.com/System-Glitch/goyave/v3/database/dialect/mssql"
)

func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
