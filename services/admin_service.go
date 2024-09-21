package services

import (
	"api-service/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminService struct {
	DB *gorm.DB
}

func (s *AdminService) CreateUser(username, password, role string, email string) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
		Email:    email,
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *AdminService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AdminService) DeleteUser(userID uint) error {
	if err := s.DB.Delete(&models.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func (s *AdminService) RevokeToken(userID uint) error {
	var user models.User
	if err := s.DB.First(&user, userID).Error; err != nil {
		return err
	}

	user.Token = "" // Revoke token by clearing it
	if err := s.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
