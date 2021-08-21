package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("assets/*")
	r.Static("/css", "../assets/css")

	r.GET("/", controllers.Page)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/student", controllers.ReadStudents)
	r.GET("/student/:id", controllers.ReadStudentById)
	r.PUT("/student/:id", controllers.UpdateStudentById)
	r.DELETE("/student/:id", controllers.DeleteStudentById)
	// r.GET("/student", controllers.ReadStudents)
	// r.DELETE("/student/:id", controllers.DeleteStudent)
	// r.PUT("/student/:id", controllers.UpdateStudent)

	return r
}
