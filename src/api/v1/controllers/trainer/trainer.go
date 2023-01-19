package trainer

import "github.com/gin-gonic/gin"

func PokemonList(c *gin.Context, in *PokemonListInput) (*PokemonListOutput, error) {
	return &PokemonListOutput{}, nil
}
