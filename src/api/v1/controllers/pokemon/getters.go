package pokemon

import (
	"fmt"
	"src/api/v1/models/generic"
	"src/api/v1/models/pokemon"

	juju "github.com/juju/errors"
)

func getCharacteristics(id int) (*PokemonCharactOutput, error) {
	poke, lookupErr := generic.GetByID[pokemon.Pokemon](pokemon.PokemonDB{}, uint(id))
	fmt.Println(poke)
	parsed, parsingErr := generic.FromTo[pokemon.Pokemon, PokemonCharactOutput](poke)
	fmt.Println(parsed)
	if lookupErr != nil {
		return nil, juju.NewNotFound(lookupErr, "pokemon not found in the database")
	}
	if parsingErr != nil {
		return nil, juju.NewNotFound(parsingErr, "data was not parsable")
	}
	return parsed, nil
}
