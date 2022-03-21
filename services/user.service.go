package services

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"vault/server/models"
	"vault/server/utils"
	"vault/server/validators"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(requestData validators.CreateUser)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (c *userService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	utils.DB.Find(&users)
	return users, nil
}

func (c *userService) CreateUser(requestData validators.CreateUser) {
	user := models.User{
		Email:           requestData.Email,
		Password:        hashPassword([]byte(requestData.Password)),
		FirstName:       requestData.FirstName,
		LastName:        requestData.LastName,
		PermissionLevel: 2,
	}
	utils.DB.Create(&user)
}

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
