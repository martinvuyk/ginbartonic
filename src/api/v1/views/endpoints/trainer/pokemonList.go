package trainer

import (
	controller "src/api/v1/controllers/trainer"
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/generic"

	"github.com/gin-gonic/gin"
)

// @Summary		PokemonList
// @Description	get a list of pokemon which the trainer owns
// @Tags		PokemonList
// @Produce		json
// @Success		200	{object}	any
// @Router		/trainer/pokemonlist [get]
func (*getPokemonListType) docs(*gin.Context) {}

type getPokemonListIn = *controller.PokemonListInput
type getPokemonListOut = *conventions.ApiResponse[controller.PokemonListOutput]
type getPokemonListAlias = generic.Endpoint[getPokemonListIn, getPokemonListOut]
type getPokemonListType struct {
	generic.Endpoint[getPokemonListIn, getPokemonListOut]
}

var getPokemonList = getPokemonListType{getPokemonListAlias{OkCode: 200}}

func (*getPokemonListType) handler(c *gin.Context, in getPokemonListIn) (getPokemonListOut, error) {
	return conventions.Respond(c, in, controller.PokemonList, 200)
}
