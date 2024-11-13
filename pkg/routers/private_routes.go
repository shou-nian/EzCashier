package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/controllers"
	"github.com/shou-nian/EzCashier/app/middleware"
)

const routerPath = "/api/v1/user"

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(router *gin.Engine) {
	// Use JWT middleware.
	router.Use(middleware.PrivateAuthorizationMiddleware())
	{
		router.POST(routerPath, controllers.CreateUser)
		router.PUT(routerPath, controllers.UpdateUser)
		router.DELETE(routerPath, controllers.DeleteUser)
	}
}
