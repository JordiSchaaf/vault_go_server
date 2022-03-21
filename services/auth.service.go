package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"vault/server/models"
	"vault/server/utils"
)

var (
	_jwtService JWTService = NewJWTService()
)

type AuthService interface {
	VerifyCredentials(email string, password string) (string, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (c *authService) VerifyCredentials(email string, password string) (string, error) {
	user, err := findByEmail(email)
	if err != nil {
		println(err.Error())
		return "", errors.New("failed to login, check your credentials")
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return "", errors.New("failed to login, check your credentials")
	}

	token := _jwtService.GenerateToken(strconv.FormatInt(user.Id, 10), user.Email)

	return token, nil
}

func findByEmail(email string) (models.User, error) {
	var user models.User
	res := utils.DB.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func comparePassword(hashedPassword string, inputPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, inputPassword)
	if err != nil {
		return false
	}
	return true
}
