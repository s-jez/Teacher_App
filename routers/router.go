package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("assets/*.html")
	r.Use(static.Serve("/assets", static.LocalFile("./assets/", true)))

	r.GET("/", controllers.Page)
	r.GET("/students/create", controllers.FormCreate)
	r.GET("/students/read", controllers.FormRead)
	r.GET("/students/update", controllers.FormUpdate)
	r.GET("/students/delete", controllers.FormDelete)

	r.POST("/student", controllers.CreateStudent)
	r.GET("/student", controllers.ReadStudents)
	r.GET("/student/:id", controllers.ReadStudentById)
	r.PUT("/student/:id", controllers.UpdateStudentById)
	r.DELETE("/student/:id", controllers.DeleteStudentById)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "error.html", gin.H{"title": "Teacher App in GO!", "error": "Page not found!"})
	})

	return r
}
