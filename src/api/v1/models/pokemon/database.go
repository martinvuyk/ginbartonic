package pokemon

import (
	"gorm.io/gorm"
)

type Relationships struct {
	// Friends []PokemonDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Embed data.go representation of Pokemon and override json
// field definitions so that gorm treats them as text and has no need for serializers
type pokemon struct {
	Pokemon       Pokemon       `gorm:"embedded;<-:create"`
	Species       string        `gorm:"type:text;<-:create"`
	Types         string        `gorm:"type:text;<-:create"`
	Relationships Relationships `gorm:"embedded"`
}

type PokemonDB struct {
	gorm.Model
	Pokemon pokemon `gorm:"embedded"`
}
