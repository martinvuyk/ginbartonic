package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/views/conventions"

	"github.com/martinvuyk/gadgeto/tonic"
)

type endpoints struct{}

func Setup(router *gin.Engine) {
	// Readme for tonic in https://github.com/martinvuyk/gadgeto/tree/master/tonic
	tonic.SetErrorHook(conventions.ErrHook)
	endp := endpoints{}
	router.GET("hello/:name", tonic.Handler(endp.greeting, 200))
	router.GET("trainer/pokemonlist", tonic.Handler(endp.pokemonList, 200))
	router.GET("pokemon/characteristics", tonic.Handler(endp.characteristics, 200))
}
