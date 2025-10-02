package main

import (
	"fmt"
	"log"

	"github.com/Bekzhanizb/EDHW2/controller"
	"github.com/Bekzhanizb/EDHW2/initializer"
	"github.com/Bekzhanizb/EDHW2/migrate"
	"github.com/Bekzhanizb/EDHW2/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
	migrate.RunMigrations()
}

func main() {
	usersToCreate := []models.User{
		{Name: "Bekzhan", Age: 19},
		{Name: "Aigerim", Age: 22},
		{Name: "Dias", Age: 25},
		{Name: "Alina", Age: 18},
	}

	for _, u := range usersToCreate {
		controller.CreateUser(u)
	}

	users := controller.GetUsers(18)
	fmt.Println("\nUsers 18< :")
	for _, u := range users {
		fmt.Printf("ID=%d | Name=%s | Age=%d | CreatedAt=%s\n", u.ID, u.Name, u.Age, u.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	if len(users) > 0 {
		firstUser := users[0]
		controller.UpdateUser(firstUser.ID, "Maria", firstUser.Age+1)
	}

	if len(users) > 1 {
		secondUser := users[1]
		controller.DeleteUser(secondUser.ID)
	}

	updatedUsers := controller.GetUsers(0)
	fmt.Println("\nUsers after update & delete:")
	for _, u := range updatedUsers {
		fmt.Printf("ID=%d | Name=%s | Age=%d | CreatedAt=%s\n", u.ID, u.Name, u.Age, u.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	log.Println("Migration finished")
}
