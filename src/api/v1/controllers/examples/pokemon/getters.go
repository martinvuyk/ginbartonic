package pokemon

import (
	"src/api/v1/models/examples/pokemon"
	"src/api/v1/models/generic"

	juju "github.com/juju/errors"
)

func getCharacteristics(id int) (*PokemonCharactOutput, error) {
	poke, lookupErr := generic.GetByID(&pokemon.PokemonDB{}, uint(id))
	if lookupErr != nil {
		return nil, juju.NewNotFound(lookupErr, "pokemon not found in the database")
	}

	parsed, parsingErr := generic.FromTo(poke.GetDataRepr(), &PokemonCharactOutput{})
	if parsingErr != nil {
		return nil, juju.NewNotProvisioned(parsingErr, "data was not parsable")
	}
	return parsed, nil
}
