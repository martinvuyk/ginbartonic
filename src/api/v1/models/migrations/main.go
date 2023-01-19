package migrations

import (
	"src/api/v1/models"
)

func Migrate() {
	// Insert your models here
	models.Database.AutoMigrate()
}
