package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/shou-nian/EzCashier/docs"
	"github.com/shou-nian/EzCashier/pkg/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerRoutes func for describe group of Swagger routes.
func SwaggerRoutes(router *gin.Engine) {
	// Define server settings:
	serverConnURL, _ := utils.ConnectionURLBuilder("server")

	// Build Swagger route.
	url := ginSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", serverConnURL))
	// Routes for GET method:
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
