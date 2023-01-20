package pokemon

import "github.com/gin-gonic/gin"

func Characteristics(c *gin.Context, in *PokemonCharactInput) (*PokemonCharactOutput, error) {
	charact, err := getCharacteristics(in.ID)
	return charact, err
}
