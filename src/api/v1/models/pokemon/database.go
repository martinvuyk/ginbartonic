package pokemon

import (
	"src/api/v1/models/generic"
)

type Relationships struct {
	generic.Relationships
	// Friends []PokemonDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type pokemon struct {
	generic.InterDataAspect[Pokemon] `gorm:"embedded"`
	Species                          string `gorm:"type:text"`
	Types                            string `gorm:"type:text"`
}

type PokemonDB struct {
	generic.DbAspect[Pokemon, PokemonDB] `gorm:"embeded"`
}
