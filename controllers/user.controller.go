package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"vault/server/services"
	"vault/server/validators"
)

var (
	jwtService  services.JWTService  = services.NewJWTService()
	userService services.UserService = services.NewUserService()
)

func GetUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Could not get users"})
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func RegisterUser(c *gin.Context) {
	// Validate input
	var newUser validators.CreateUser
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := userService.RegisterUser(newUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"email": res.Email})
}

func DeleteUser(c *gin.Context) {
	userIdToRemove := c.Param("userId")

	token := jwtService.ValidateToken(c.GetHeader("Authorization"))
	claims := token.Claims.(jwt.MapClaims)
	issuerId := fmt.Sprintf("%v", claims["user_id"])

	err := userService.DeleteUser(userIdToRemove, issuerId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed"})
}
