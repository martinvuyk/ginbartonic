package trainer

import (
	"src/api/v1/models/generic"
)

type relationships struct {
	// Pokemons []pokemon.PokemonDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type interDataAspect struct {
	DataAspect    Trainer       `gorm:"embedded"`
	Relationships relationships `gorm:"embedded"`
}

type TrainerDB struct {
	generic.DbAspect[Trainer] `gorm:"embeded"`
	InterDataAspect           interDataAspect `gorm:"embedded"`
}
