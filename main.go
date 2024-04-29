package main

import (
	"embed"
	"fmt"
	"os"

	"goyave.dev/template/http/route"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/fsutil"
	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v5/database/dialect/mysql"
	// _ "goyave.dev/goyave/v5/database/dialect/postgres"
	// _ "goyave.dev/goyave/v5/database/dialect/sqlite"
	// _ "goyave.dev/goyave/v5/database/dialect/mssql"
)

//go:embed resources
var resources embed.FS

func main() {
	resources := fsutil.NewEmbed(resources)
	langFS, err := resources.Sub("resources/lang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	opts := goyave.Options{
		LangFS: langFS,
	}

	server, err := goyave.New(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	server.Logger.Info("Registering hooks")
	server.RegisterSignalHook()

	server.RegisterStartupHook(func(s *goyave.Server) {
		server.Logger.Info("Server is listening", "host", s.Host())
	})

	server.RegisterShutdownHook(func(s *goyave.Server) {
		s.Logger.Info("Server is shutting down")
	})

	registerServices(server)

	server.Logger.Info("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.Logger.Error(err)
		os.Exit(3)
	}
}

func registerServices(server *goyave.Server) {
	server.Logger.Info("Registering services")

	// TODO register services
}
