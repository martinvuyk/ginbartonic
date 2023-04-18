package trainer

import (
	controller "src/api/v1/controllers/examples/trainer"
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/generic"

	"github.com/gin-gonic/gin"
)

type getPokeListIn = *controller.PokemonListInput
type getPokeListOut = *conventions.ApiResponse[controller.PokemonListOutput]
type getPokeListEndp = generic.Endpoint[getPokeListIn, getPokeListOut]
type getPokeListType struct {
	getPokeListEndp
}

var getPokeList = getPokeListType{getPokeListEndp{OkCode: 200}}

// @Summary		PokemonList
// @Description	get a list of pokemon which the trainer owns
// @Tags		PokemonList
// @Produce		json
// @Success		200	{object}	getPokeListOut
// @Router		/trainer/pokemonlist [get]
func (*getPokeListType) docs(*gin.Context) {}

func (*getPokeListType) handler(c *gin.Context, in getPokeListIn) (getPokeListOut, error) {
	return conventions.Respond(c, in, controller.PokemonList, 200)
}
