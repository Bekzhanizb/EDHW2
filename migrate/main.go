package main

import (
	"log"

	"github.com/Bekzhanizb/EDHW2/initializer"
	"github.com/Bekzhanizb/EDHW2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	log.Println("🚀 Running migrations...")

	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("✅ Migration complete!")
}
