package v1

import (
	"fmt"
	"time"

	"src/api/v1/models"
	"src/api/v1/models/migrations"
)

func InitDB() {
	// Connect to database
	err := models.Connect()

	if err != nil {
		time.Sleep(10 * time.Second)
		err = models.Connect()
	}
	if err != nil {
		panic(fmt.Sprintf("DATABASE CONNECTION TIMEOUT:\n %s", err))
	}

	// migrate Models
	migrations.Migrate()

}

func Init() {
	InitDB()
}
