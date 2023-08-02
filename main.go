package main

import (
	"fmt"
	"os"

	"goyave.dev/template/database/model"
	"goyave.dev/template/http/route"
	"goyave.dev/template/service/user"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"

	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v5/database/dialect/mysql"
	// _ "goyave.dev/goyave/v5/database/dialect/postgres"
	_ "goyave.dev/goyave/v5/database/dialect/sqlite"
	// _ "goyave.dev/goyave/v5/database/dialect/mssql"
)

func main() {

	server, err := goyave.New()
	if err != nil {
		fmt.Println(err.(*errors.Error).String())
		os.Exit(1)
	}

	if err := server.DB().AutoMigrate(&model.User{}); err != nil {
		server.ErrLogger.Println(errors.New(err).String())
		os.Exit(2)
	}
	factory := database.NewFactory(model.UserGenerator)
	factory.Save(server.DB(), 21)

	server.Logger.Println("Registering hooks")
	server.RegisterSignalHook()

	server.RegisterStartupHook(func(s *goyave.Server) {
		s.Logger.Printf("Server is listening on %s\n", s.Host())
	})

	server.RegisterShutdownHook(func(s *goyave.Server) {
		s.Logger.Println("Server is shutting down")
	})

	server.Logger.Println("Registering services")
	server.RegisterService(&user.Service{})

	server.Logger.Println("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.ErrLogger.Println(err.(*errors.Error).String())
		os.Exit(3)
	}
}
