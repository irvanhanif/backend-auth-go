package routers

import (
	"trigonal/backend-auth/api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.POST("login", controllers.Login)
	auth.POST("register", controllers.Register)
}
