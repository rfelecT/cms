package routes

import (
	"cms/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/forgot-password", controllers.ForgotPassword)
	r.POST("/reset-password", controllers.ResetPassword)
}
