package routes

import (
	"github.com/Niraj1910/school-census-go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.GET("/otp", controllers.SendOTP)
		auth.POST("/otp/verify", controllers.VerifyOTP)
	}
}
