package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.Use(static.Serve("/assets", static.LocalFile("/assets", false)))
	r.LoadHTMLGlob("assets/*")

	r.GET("/", controllers.Page)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/student", controllers.ReadStudents)
	r.DELETE("/student/:id", controllers.DeleteStudent)
	r.PUT("/student/:id", controllers.UpdateStudent)

	return r
}
