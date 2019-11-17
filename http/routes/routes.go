package routes

import (
	"github.com/System-Glitch/goyave"
	helloController "goyave_template/http/controllers/hello"
)

// Register all the application routes.
func Register(router *goyave.Router) {
	// Register your routes here
	router.Route("GET", "/hello", helloController.SayHi, nil)
}
