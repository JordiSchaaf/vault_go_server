package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vault/server/controllers"
	"vault/server/utils"
)

func main() {
	utils.ConnectDatabase()
	defer utils.DisconnectDatabase()

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hoi"})
	})

	authRoutes := server.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
	}

	userRoutes := server.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/me", controllers.GetMe)
		userRoutes.GET("/:userId", controllers.GetUser)
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.DELETE("/:userId", controllers.DeleteUser)
	}

	server.Run()
}
