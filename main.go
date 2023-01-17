package main

import (
	"fmt"

	"github.com/BoiseITGuru/ArrRequests/internal/controllers"
	"github.com/BoiseITGuru/ArrRequests/internal/database"
	"github.com/BoiseITGuru/ArrRequests/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configurations using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect()
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(fmt.Sprintf(":%v", AppConfig.ServerPort))
}

func initRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		// Token Route - Receives login requests and returns a JWT Token
		auth.POST("/token", controllers.GenerateToken)

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
