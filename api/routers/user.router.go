package routers

import (
	"trigonal/backend-auth/api/controllers"
	"trigonal/backend-auth/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	user.GET("", middleware.ValidateToken, controllers.GetAllUsers)
	user.GET(":id",middleware.ValidateToken, controllers.GetUser)
	user.PUT("", middleware.ValidateToken, controllers.UpdateUser)
	user.DELETE(":id", middleware.ValidateToken, controllers.DeleteUser)
}