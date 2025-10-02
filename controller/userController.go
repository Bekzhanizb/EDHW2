package controller

import (
	"log"

	"github.com/Bekzhanizb/EDHW2/initializer"
	"github.com/Bekzhanizb/EDHW2/models"
)

func CreateUser(user models.User) {
	if err := initializer.DB.Create(&user).Error; err != nil {
		log.Println("Creating new User failed:", err)
		return
	}
	log.Println("New User created.")
}
func GetUsers(minAge int) []*models.User {
	var users []*models.User
	if err := initializer.DB.Where("age > ?", minAge).Find(&users).Error; err != nil {
		log.Println("Query failed:", err)
	}
	return users
}
func UpdateUser(id uint, newName string, newAge int) {
	if err := initializer.DB.Model(&models.User{}).
		Where("id = ?", id).
		Updates(models.User{Name: newName, Age: newAge}).Error; err != nil {
		log.Println("Updating the User dates failed:", err)
	} else {
		log.Println("User updated.")
	}
}
func DeleteUser(id uint) {
	if err := initializer.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		log.Println("Deleting the User failed:", err)
	} else {
		log.Println("User deleted.")
	}
}
