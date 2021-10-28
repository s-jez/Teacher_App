package controllers

import "github.com/gin-gonic/gin"

func Page(c *gin.Context) {
	c.HTML(200, "index.gohtml", gin.H{"title": "Students App"})
}
