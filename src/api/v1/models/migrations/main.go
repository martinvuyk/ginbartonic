package migrations

import (
	"src/api/v1/models"
	"src/api/v1/models/pokemon"
	"src/api/v1/models/trainer"
)

func Migrate() {
	models.Database.AutoMigrate(&pokemon.PokemonDB{}, &trainer.TrainerDB{})
}
