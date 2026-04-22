package services

import (
	"go_ecommerce/config"
	"go_ecommerce/models"
	"go_ecommerce/repositories"

	"golang.org/x/crypto/bcrypt"
)

func Register(name, email, password string, role_id int) error {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user := &models.Users{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role_id:  role_id,
	}

	return repositories.CreateUsers(user)

}

func Login(email, password string) (*models.Users, error) {
	user := &models.Users{}
	result := config.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil

}
