package trainer

import (
	"src/api/v1/views/conventions"

	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

func Setup(router *gin.RouterGroup, m *conventions.Metrics) {
	pokL := getPokeList(m)
	router.GET(pokL.Url, pokL.docs, tonic.Handler(pokL.handler, pokL.OkCode))
}
