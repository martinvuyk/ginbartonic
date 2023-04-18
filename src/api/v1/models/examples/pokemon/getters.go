package pokemon

import "src/api/v1/models"

func (pokemon *Pokemon) GetRelationships() (relationships, error) {
	var relations relationships
	search := models.Database.Model(&PokemonDB{}).First(&relations, pokemon.ID)
	if err := search.Error; err != nil {
		return relationships{}, err
	}
	return relations, nil
}

func (pokemonDb *PokemonDB) GetDataRepr() *Pokemon {
	poke := Pokemon{
		ID:      pokemonDb.ID,
		Name:    pokemonDb.InterDataAspect.DataAspect.Name,
		Species: pokemonDb.InterDataAspect.Species,
		Height:  pokemonDb.InterDataAspect.DataAspect.Height,
		Weight:  pokemonDb.InterDataAspect.DataAspect.Weight,
		Types:   pokemonDb.InterDataAspect.Types,
	}
	return &poke
}
