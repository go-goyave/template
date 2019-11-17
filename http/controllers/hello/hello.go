package hello

import (
	"github.com/System-Glitch/goyave"
	"net/http"
)

func SayHi(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "Hi!")
}
