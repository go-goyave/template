package route

import (
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
)

// Routing is an essential part of any Goyave application.
// Defining routes is the action of associating a URI, sometimes having parameters,
// with a handler which will process the request and respond to it.
//
// This file contains your main route registering function that is passed to server.RegisterRoutes().
//
// Learn more here: https://goyave.dev/basics/routing.html

func Register(server *goyave.Server, router *goyave.Router) {
	router.CORS(cors.Default())
	router.GlobalMiddleware(&parse.Middleware{})

	// TODO register routes
}
