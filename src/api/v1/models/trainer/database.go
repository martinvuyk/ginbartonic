package trainer

import (
	"gorm.io/gorm"
	// "microservice3/api/v1/models/pokemon"
)

type Relationships struct {
	// Pokemons []pokemon.PokemonDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type trainer struct {
	Trainer       Trainer       `gorm:"embedded"`
	Relationships Relationships `gorm:"embedded"`
}

type TrainerDB struct {
	gorm.Model
	Trainer trainer `gorm:"embedded"`
}
