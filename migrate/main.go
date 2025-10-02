package migrate

import (
	"log"

	"github.com/Bekzhanizb/EDHW2/initializer"
	"github.com/Bekzhanizb/EDHW2/models"
)

func RunMigrations() {
	err := initializer.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migrations applied")
}
