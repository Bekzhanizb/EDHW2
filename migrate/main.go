package main

import (
	"log"

	"github.com/Bekzhanizb/EDHW2/initializer"
	"github.com/Bekzhanizb/EDHW2/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
}

func main() {
	log.Println("Running migrations...")

	err := initializer.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration complete!")
}
