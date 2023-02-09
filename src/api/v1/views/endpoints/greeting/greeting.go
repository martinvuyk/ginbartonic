package greeting

import (
	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

func Setup(router *gin.RouterGroup) {
	router.GET("hello/:name", getGreet.docs, tonic.Handler(getGreet.handler, getGreet.OkCode))
}
