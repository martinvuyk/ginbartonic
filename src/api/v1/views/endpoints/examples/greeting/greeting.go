package greeting

import (
	"src/api/v1/views/conventions"

	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

func Setup(router *gin.RouterGroup, m *conventions.Metrics) {
	g := getGreet(m)
	router.GET(g.Url, g.docs, tonic.Handler(g.handler, g.OkCode))
}
