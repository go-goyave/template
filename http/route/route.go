package route

import (
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/middleware/parse"
	"goyave.dev/template/http/controller/user"
)

func Register(_ *goyave.Server, router *goyave.Router) {

	router.CORS(cors.Default())
	router.GlobalMiddleware(&parse.Middleware{})

	router.Controller(&user.UserController{})
}
