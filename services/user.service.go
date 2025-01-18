package userService

import (
	"go-crud/db"
	"go-crud/models"
)

// CreateUser creates a new user.
func CreateUser(user *models.User) (*models.User, error) {
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers retrieves a list of all users.
func GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Find(&users).Error
	return users, err
}

// GetUser retrieve a single user based on userId.
func GetUser(userId int) (models.User, error) {
	var user models.User
	err := db.DB.First(&user, userId).Error
	return user, err
}

// UpdateUser updates an existing user by ID.
func UpdateUser(id int, updatedUser *models.User) (*models.User, error) {
	if err := db.DB.Model(&models.User{}).Where("id = ?", id).Updates(updatedUser).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// DeleteUser deletes a user by ID.
func DeleteUser(id int) error {
	return db.DB.Delete(&models.User{}, id).Error
}
