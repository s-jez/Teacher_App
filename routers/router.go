package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func CreateUrlMappings() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("sites/html/*.html")
	r.Use(static.Serve("/sites", static.LocalFile("./sites/js", true)))
	userRoles := map[string]int{
		"user":  1,
		"dev":   2,
		"admin": 3,
	}
	r.GET("/", controllers.WelcomePage)
	r.GET("/logged", controllers.MainPage)
	r.GET("/student", controllers.ReadStudents)
	r.GET("/student/:id", controllers.ReadStudentById)
	r.POST("/student", controllers.AuthMiddleware([]int{userRoles["admin"], userRoles["dev"]}), controllers.CreateStudent)
	r.PUT("/student/:id", controllers.AuthMiddleware([]int{userRoles["admin"], userRoles["dev"]}), controllers.UpdateStudentById)
	r.DELETE("/student/:id", controllers.AuthMiddleware([]int{userRoles["admin"]}), controllers.DeleteStudentById)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "error.html", gin.H{"title": "Page not found!"})
	})
	r.POST("/login", controllers.LoginUser)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/refresh", controllers.RefreshToken)
	return r
}
