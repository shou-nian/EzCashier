package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/middleware"
	"github.com/shou-nian/EzCashier/pkg/configs"
	"github.com/shou-nian/EzCashier/pkg/routers"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/migrations"
	"log"

	_ "github.com/joho/godotenv/autoload" // auto load environment variables
)

func main() {
	// Run auto migration
	err := migrations.AutoMigration()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a new router.
	router := gin.New()

	// Use JWT middleware
	router.Use(middleware.JWTMiddleware())

	// List of app routes:
	routers.PrivateRoutes(router)

	// Initialize server.
	server := configs.ServerConfig(router)

	// Start API server.
	utils.StartServerWithGracefulShutdown(server)
}
