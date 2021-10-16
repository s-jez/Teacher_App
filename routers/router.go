package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.POST("/student", controllers.CreateStudent)
	r.GET("/student", controllers.ReadStudents)
	r.GET("/student/:id", controllers.ReadStudentById)
	r.PUT("/student/:id", controllers.UpdateStudentById)
	r.DELETE("/student/:id", controllers.DeleteStudentById)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"data": "Page not found!"})
	})

	return r
}
