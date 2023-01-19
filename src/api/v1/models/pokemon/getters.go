package pokemon

import (
	"src/api/v1/models"
)

func (pokemon *Pokemon) Get_relationships() (Relationships, error) {
	var relationships Relationships
	search := models.Database.Model(&PokemonDB{}).First(&relationships, pokemon.ID)
	if err := search.Error; err != nil {
		return Relationships{}, err
	}
	return relationships, nil
}

func (pokemon *PokemonDB) Get_data_repr() Pokemon {
	return pokemon.Pokemon.Pokemon
}
