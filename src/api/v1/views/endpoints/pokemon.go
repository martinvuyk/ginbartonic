package endpoints

import (
	"github.com/gin-gonic/gin"

	controller "src/api/v1/controllers/pokemon"
	"src/api/v1/views/conventions"
)

type pokemonType struct{}

var pokemon = pokemonType{}

func (*pokemonType) getPokemonById(c *gin.Context, in *controller.PokemonCharactInput) (*conventions.ApiResponse[controller.PokemonCharactOutput], error) {
	return conventions.Respond(c, in, controller.GetPokemonById, 200)
}

func (*pokemonType) setup(router *gin.Engine) {
	// endpoint to get charachteristics of a pokemon searching by id
	register(router.GET, "pokemon/:id", pokemon.getPokemonById, 200)
}
