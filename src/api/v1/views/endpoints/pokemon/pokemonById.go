package pokemon

import (
	controller "src/api/v1/controllers/pokemon"
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/generic"

	"github.com/gin-gonic/gin"
)

// @Summary		Pokemon
// @Description	get pokemon with id
// @Tags		pokemon
// @Produce		json
// @Success		200	{object}	any
// @Router		/pokemon/:id [get]
func (*getPokemonByIdType) docs(*gin.Context) {}

type getPokemonByIdIn = *controller.PokemonCharactInput
type getPokemonByIdOut = *conventions.ApiResponse[controller.PokemonCharactOutput]
type getPokemonByIdAlias = generic.Endpoint[getPokemonByIdIn, getPokemonByIdOut]
type getPokemonByIdType struct {
	generic.Endpoint[getPokemonByIdIn, getPokemonByIdOut]
}

var getPokemonById = getPokemonByIdType{getPokemonByIdAlias{OkCode: 200}}

func (p *getPokemonByIdType) handler(c *gin.Context, in getPokemonByIdIn) (getPokemonByIdOut, error) {
	return conventions.Respond(c, in, controller.GetPokemonById, p.OkCode)
}
