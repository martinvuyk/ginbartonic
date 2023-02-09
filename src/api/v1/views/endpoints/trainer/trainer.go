package trainer

import (
	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

func Setup(router *gin.RouterGroup) {
	pokL := getPokemonList
	router.GET("trainer/pokemonlist", pokL.docs, tonic.Handler(pokL.handler, pokL.OkCode))
}
