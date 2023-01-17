package main

import (
	"fmt"

	"github.com/BoiseITGuru/ArrRequests/internal/config"
	"github.com/BoiseITGuru/ArrRequests/internal/controllers"
	"github.com/BoiseITGuru/ArrRequests/internal/middleware"
	"github.com/BoiseITGuru/ArrRequests/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configurations using Viper
	config.LoadAppConfig()

	// Start Services
	services.Start()

	// Initialize Router
	router := initRouter()
	router.Run(fmt.Sprintf(":%v", config.AppConfig.ServerPort))
}

func initRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		// Login Route - Receives login requests and returns a JWT Token
		auth.POST("/login", controllers.Login)

		// Refresh Token Route - Receives JWT Refresh Token and return new JWT Token
		auth.POST("/refresh-token", controllers.RefreshToken)
	}

	api := router.Group("/api")
	api.Use(middleware.Auth())
	tmdb := api.Group("/tmdb")
	{
		// Trending Route - Get Trending Movies
		tmdb.GET("/trending", controllers.GetTrending)
	}
	return router
}
