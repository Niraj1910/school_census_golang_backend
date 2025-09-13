package routes

import (
	"github.com/Niraj1910/school-census-go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSchoolRoutes(router *gin.Engine) {
	school := router.Group("/api/schools")
	{
		school.GET("/", controllers.GetSchools)
		school.GET("/:id", controllers.GetSchoolByID)
		school.POST("/", controllers.CreateSchool)
		school.PUT("/:id", controllers.UpdateSchool)
		school.DELETE("/:id", controllers.DeleteSchool)
	}
}
