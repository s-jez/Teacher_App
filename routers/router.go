package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("assets/*.html")
	r.Use(static.Serve("/assets", static.LocalFile("./assets/", true)))
	r.Use(sessions.Sessions("session", sessions.NewCookieStore([]byte("secretkey"))))
	app := r.Group("/students")
	{
		app.GET("/create", controllers.FormCreate)
		app.POST("/login", controllers.Login)
		app.GET("/logout", controllers.Logout)
		app.GET("/account", controllers.Account)
		app.GET("/status", controllers.Status)
		// app.GET("/update", controllers.FormUpdate)
		// app.GET("/delete", controllers.FormDelete)
	}
	r.GET("/", controllers.Page)
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
