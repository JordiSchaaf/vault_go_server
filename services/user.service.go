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
	RegisterUser(requestData validators.CreateUser) (*models.User, error)
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

func (c *userService) RegisterUser(requestData validators.CreateUser) (*models.User, error) {
	user := models.User{
		Email:           requestData.Email,
		Password:        hashPassword([]byte(requestData.Password)),
		FirstName:       requestData.FirstName,
		LastName:        requestData.LastName,
		PermissionLevel: 2,
	}
	res := utils.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}

//func (c *userService) DeleteUser(userIdToRemove string, issuerId string) error {
//	user, err := utils.DB.Find(userIdToRemove)
//	if err != nil {
//		return err
//	}
//
//}
