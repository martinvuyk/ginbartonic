package pokemon

import (
	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

func Setup(router *gin.RouterGroup) {
	gePok := getPokemonById
	router.GET("pokemon/:id", gePok.docs, tonic.Handler(gePok.handler, gePok.OkCode))
}
