package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vault/server/services"
	"vault/server/validators"
)

var (
	authService services.AuthService = services.NewAuthService()
)

func Login(c *gin.Context) {
	// Validate input
	var loginRequest validators.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := authService.VerifyCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auth succesful",
		"token": token})
}
