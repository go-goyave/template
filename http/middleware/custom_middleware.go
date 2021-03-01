package middleware

import "goyave.dev/goyave/v3"

// Middleware are handlers executed before the controller handler.
// They are a convenient way to filter, intercept or alter HTTP requests entering your application.
// Learn more here: https://goyave.dev/guide/basics/middleware.html

// MyCustomMiddleware is an example middleware.
//
// To use this middleware, assign it to a router in "http/routes/routes.go"
//
//     router.Middleware(middleware.MyCustomMiddleware)
func MyCustomMiddleware(next goyave.Handler) goyave.Handler {
	return func(response *goyave.Response, request *goyave.Request) {
		// Do something
		next(response, request) // Pass to the next handler
	}
}
