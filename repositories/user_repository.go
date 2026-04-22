package repositories

import (
	"go_ecommerce/config"
	"go_ecommerce/models"
)

func CreateUsers(user *models.Users) error {
	result := config.DB.Create(user)
	return result.Error
}

func GetUserByEmail(email string) (*models.Users, error) {
	user := &models.Users{}
	result := config.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
