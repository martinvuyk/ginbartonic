package endpoints

import (
	"github.com/gin-gonic/gin"

	controller "src/api/v1/controllers/trainer"
	"src/api/v1/views/conventions"
)

type trainerType struct{}

var trainer = trainerType{}

// endpoint to get a list of pokemon which the trainer owns
func (*trainerType) getPokemonList(c *gin.Context, in *controller.PokemonListInput) (*conventions.ApiResponse[controller.PokemonListOutput], error) {
	return conventions.Respond(c, in, controller.PokemonList, 200)
}

func (*trainerType) setup(router *gin.RouterGroup) {
	register(router.GET, "trainer/pokemonlist", trainer.getPokemonList, 200)
}
