package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/middleware"
	"github.com/shou-nian/EzCashier/pkg/configs"
	"github.com/shou-nian/EzCashier/pkg/routers"
	"github.com/shou-nian/EzCashier/pkg/utils"
	"github.com/shou-nian/EzCashier/repository/migrations"
	"log"

	_ "github.com/joho/godotenv/autoload" // auto load environment variables
)

// @title EzCashier API
// @version 1.0
// @description This is the API server for the EzCashier application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Run auto migration
	err := migrations.AutoMigration()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize a new router.
	router := gin.New()
	// Use Gin's Logger middleware, it will log all requests using the default logger.
	router.Use(gin.Logger())
	// Use Gin's cover all panic errors middleware
	router.Use(gin.Recovery())
	// Use CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080", "http://0.0.0.0:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// List of app routes:
	{
		// Swagger routers:
		routers.SwaggerRoutes(router)

		// Public routers:
		routers.PublicRouters(router)

		// Private routers use the JWT middleware
		router.Use(middleware.JWTMiddleware())
		routers.PrivateRouters(router)
	}

	// Initialize server.
	server := configs.ServerConfig(router)

	// Start API server.
	utils.StartServerWithGracefulShutdown(server)
}
