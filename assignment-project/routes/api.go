package routes

import (
	"assignmentproject/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	route := gin.Default()

	api := route.Group("/students")
	{
		api.GET("/", controllers.GetAllStudents)
		api.POST("/", controllers.CreateStudent)
		api.GET("/:studentID", controllers.GetStudent)
		api.PUT("/:studentID", controllers.UpdateStudent)
		api.DELETE("/:studentID", controllers.DeleteStudent)
	}

	return route
}
