package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shou-nian/EzCashier/app/controllers"
)

// PublicRouters func for describe group of private routes.
func PublicRouters(router *gin.Engine) {
	{
		router.POST("/api/v1/login", controllers.Login)
	}
}
