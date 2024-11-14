package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/controllers"
)

const authRouterPath = "/api/v1/login"

// PublicRoutes func for describe group of private routes.
func PublicRoutes(router *gin.Engine) {
	// Use JWT middleware.
	{
		router.POST(authRouterPath, controllers.Login)
	}
}
