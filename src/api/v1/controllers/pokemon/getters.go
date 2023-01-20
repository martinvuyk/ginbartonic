package pokemon

import (
	"src/api/v1/models/generic"
	"src/api/v1/models/pokemon"

	juju "github.com/juju/errors"
)

func getCharacteristics(id int) (*PokemonCharactOutput, error) {
	poke, lookupErr := pokemon.GetByID(uint(id))
	data, parsingErr := generic.Copy[PokemonCharactOutput](poke)
	if lookupErr != nil {
		return nil, juju.NewNotFound(lookupErr, "pokemon not found in the database")
	}
	if parsingErr != nil {
		return nil, juju.NewNotFound(parsingErr, "data was not parsable")
	}
	return &data, nil
}
