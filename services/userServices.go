package services

import (
	"go_ecommerce/config"
	"go_ecommerce/models"
)

func GetUserByID(userID int) (*models.Users, error) {
	user := &models.Users{}
	result := config.DB.
		Preload("Roles").
		Where("id = ?", userID).
		First(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func UpdateUsers(userID int, username string) (*models.Users, error) {

	result := config.DB.Model(&models.Users{}).
		Where("id = ?", userID).
		Update("name", username)
	if result.Error != nil {
		return nil, result.Error
	}

	return GetUserByID(userID)

}

func CheckIfUserAdmin(userId int) (bool, error) {

	user, err := GetUserByID(userId)
	if err != nil {
		return false, err // user not found or DB error
	}

	if user.Role_id == 1 {
		return true, nil
	}

	return false, nil
}

func CheckIfUserVendor(userId int) (bool, error) {

	user, err := GetUserByID(userId)
	if err != nil {
		return false, err // user not found or DB error
	}

	if user.Role_id == 3 {
		return true, nil
	}

	return false, nil
}

func GetallUserInfo() ([]*models.Users, error) {

	users := []*models.Users{}

	result := config.DB.
		Preload("Roles").
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func DeleteUser(userID int) error {
	result := config.DB.Delete(&models.Users{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
