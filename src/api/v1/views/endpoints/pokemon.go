package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/controllers/pokemon"
	"src/api/v1/views/conventions"
)

func (*Endpoints) characteristics(c *gin.Context, in *pokemon.PokemonCharactInput) (conventions.ApiResponse[pokemon.PokemonCharactOutput], error) {
	return conventions.Respond(c, in, pokemon.Characteristics, 200)
}
