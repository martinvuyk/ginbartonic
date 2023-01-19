package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/controllers/trainer"
	"src/api/v1/views/conventions"
)

func (*Endpoints) pokemonList(c *gin.Context, in *trainer.PokemonListInput) (conventions.ApiResponse[trainer.PokemonListOutput], error) {
	return conventions.Respond(c, in, trainer.PokemonList, 200)
}
