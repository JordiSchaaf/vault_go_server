package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"vault/server/models"
	"vault/server/utils"
	"vault/server/validators"
)

type UserService interface {
	GetAllUsers() (*[]*models.User, error)
	GetUser(userId string, issuerId string) (*models.User, error)
	RegisterUser(requestData validators.CreateUser) (*models.User, error)
	DeleteUser(userIdToRemove string, issuerId string) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (c *userService) GetAllUsers() (*[]*models.User, error) {
	var users []*models.User
	res := utils.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return &users, nil
}

func (c *userService) GetUser(userId string, issuerId string) (*models.User, error) {
	if utils.IsUserAuthorized(issuerId) || userId == issuerId {
		user := models.User{}
		res := utils.DB.First(&user, "id = ?", issuerId)
		if res.Error == nil {
			return nil, res.Error
		}
		return &user, nil
	}
	return nil, errors.New("unauthorized")
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

func (c *userService) DeleteUser(userIdToRemove string, issuerId string) error {
	if utils.IsUserAuthorized(issuerId) && userIdToRemove != issuerId {
		userToRemove := models.User{}
		res := utils.DB.Where("id = ?", userIdToRemove).Delete(&userToRemove)
		if res.Error != nil {
			return nil
		}
		return res.Error
	}
	return errors.New("unauthorized")
}
