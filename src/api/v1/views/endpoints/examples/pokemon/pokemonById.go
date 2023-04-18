package pokemon

import (
	controller "src/api/v1/controllers/examples/pokemon"
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/generic"

	"github.com/gin-gonic/gin"
)

type getPokeByIdIn = *controller.PokemonCharactInput
type getPokeByIdOut = *conventions.ApiResponse[controller.PokemonCharactOutput]
type getPokeByIdEndpoint = generic.Endpoint[getPokeByIdIn, getPokeByIdOut]
type getPokeByIdType struct {
	getPokeByIdEndpoint
}

var getPokemonById = getPokeByIdType{getPokeByIdEndpoint{OkCode: 200}}

// @Summary		Pokemon
// @Description	get pokemon with id
// @Tags		pokemon
// @Produce		json
// @Success		200	{object}	getPokeByIdOut
// @Router		/pokemon/:id [get]
func (*getPokeByIdType) docs(*gin.Context) {}

func (p *getPokeByIdType) handler(c *gin.Context, in getPokeByIdIn) (getPokeByIdOut, error) {
	return conventions.Respond(c, in, controller.GetPokemonById, p.OkCode)
}
