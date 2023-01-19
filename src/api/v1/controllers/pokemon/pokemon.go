package pokemon

import "github.com/gin-gonic/gin"

func Characteristics(c *gin.Context, in *PokemonCharactInput) (*PokemonCharactOutput, error) {
	return &PokemonCharactOutput{}, nil
}
