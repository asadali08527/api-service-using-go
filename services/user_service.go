package services

import (
	"api-service/models"
	"api-service/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// CreateUser - Create a new user in the DB
func (us *UserService) CreateUser(user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save the user
	return us.DB.Create(user).Error
}

// Authenticate - Authenticate user credentials
func (us *UserService) Authenticate(username, password string) (*models.User, error) {
	var user models.User
	err := us.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	print(password)
	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
func (s *UserService) Login(username, password string) (string, error) {
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetProfile(username string) (models.User, error) {
	var user models.User
	print(username)
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateProfile(username string, mobile, address string) (models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}

	user.Mobile = mobile
	user.Address = address

	if err := s.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
