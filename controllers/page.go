package controllers

import "github.com/gin-gonic/gin"

func WelcomePage(c *gin.Context) {
	c.HTML(200, "welcome.html", gin.H{"title": "Students CRUD"})
}
func MainPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"title": "Students CRUD"})
}
