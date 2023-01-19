package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/views/conventions"

	"github.com/martinvuyk/gadgeto/tonic"
)

type Endpoints struct{}

func Setup(router *gin.Engine) {
	// Readme for tonic in https://github.com/martinvuyk/gadgeto/tree/master/tonic
	tonic.SetErrorHook(conventions.ErrHook)
	endp := Endpoints{}
	router.GET("hello/:name", tonic.Handler(endp.greeting, 200))
}
