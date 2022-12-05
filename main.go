package main

import (
	"fmt"

	"github.com/BoiseITGuru/ArrRequests/internal/database"
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
	// ms := router.Group("/ms")
	// {
	// 	// Get Token Route - Receives auth token from MS Azure App
	// 	ms.POST("/auth", controllers.GetRefreshToken)
	// }
	return router
}
