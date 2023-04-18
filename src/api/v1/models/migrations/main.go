package migrations

import (
	"src/api/v1/models"
	"src/api/v1/models/examples/pokemon"
	"src/api/v1/models/examples/trainer"
)

func Migrate() {
	models.Database.AutoMigrate(&pokemon.PokemonDB{}, &trainer.TrainerDB{})
}
