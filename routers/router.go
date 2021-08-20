package routers

import (
	"Stachowsky/Teacher_App/controllers"

	"github.com/gin-gonic/gin"
)

// func SetupRouters() { // Słaba nazwa
func CreateUrlMappings() *gin.Engine { // Zwracamy gin.Engine
	r := gin.Default()
	r.LoadHTMLGlob("assets/*")
	r.GET("/", controllers.Page)
	r.POST("/student", controllers.CreateStudent)
	// Jest to funkcja, którą będziemy używać w main dlatego zwracamy router
	return r
}

// Page powinien być w controllersach
