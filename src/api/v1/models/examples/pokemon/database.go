package pokemon

import "src/api/v1/models/generic"

type relationships struct {
	// Friends []PokemonDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type interdataAspect struct {
	DataAspect    Pokemon       `gorm:"embedded"`
	Species       string        `gorm:"type:text"`
	Types         string        `gorm:"type:text"`
	Relationships relationships `gorm:"embedded"`
}

type PokemonDB struct {
	generic.DbAspect[Pokemon] `gorm:"embeded"`
	InterDataAspect           interdataAspect `gorm:"embedded"`
}
