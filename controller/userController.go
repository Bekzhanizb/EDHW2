package controller

import (
	"log"

	"github.com/Bekzhanizb/EDHW2/initializer"
	"github.com/Bekzhanizb/EDHW2/models"
)

func CreateUSer(user models.User) {
	if err := initializer.DB.Create(&user).Error; err != nil {
		log.Println("Insert failed:", err)
		return
	}
	log.Println("User created.")
}
func GetUsers(minAge int) []*models.User {
	var users []*models.User
	if err := initializer.DB.Where("age > ?", minAge).Find(&users).Error; err != nil {
		log.Println("Query failed:", err)
	}
	return users
}
func UpdateUSer(id uint, newName string, newAge int) {
	if err := initializer.DB.Model(&models.User{}).
		Where("id = ?", id).
		Updates(models.User{Name: newName, Age: newAge}).Error; err != nil {
		log.Println("Update failed:", err)
	} else {
		log.Println("User updated.")
	}
}
func DeleteUser(id uint) {
	if err := initializer.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		log.Println("Delete failed:", err)
	} else {
		log.Println("User deleted.")
	}
}
