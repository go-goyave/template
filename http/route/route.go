package route

import (
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
)

func Register(server *goyave.Server, router *goyave.Router) {
	router.CORS(cors.Default())
	router.GlobalMiddleware(&parse.Middleware{})

	// TODO register routes
}
