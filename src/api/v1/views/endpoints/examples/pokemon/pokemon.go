package pokemon

import (
	"src/api/v1/views/conventions"

	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

func Setup(router *gin.RouterGroup, m *conventions.Metrics) {
	gePok := getPokemonById(m)
	router.GET(gePok.Url, gePok.docs, tonic.Handler(gePok.handler, gePok.OkCode))
}
