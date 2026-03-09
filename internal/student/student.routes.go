package student

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	repo := NewRepository()

	service := NewService(repo)

	controller := NewController(service)

	students := router.Group("/students")

	{
		students.POST("", controller.CreateStudent)
		students.GET("", controller.GetAllStudents)
		students.DELETE("/:id", controller.DeleteStudent)
	}
}
