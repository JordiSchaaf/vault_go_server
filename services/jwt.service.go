package services

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userId string, email string) string
}

type jwtService struct {
	issuer    string
	secretKey string
}

type jwtCustomClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "vaultress_system",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "system"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userId string, email string) string {
	claims := &jwtCustomClaim{
		userId,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
