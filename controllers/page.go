package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Teacher app in GO!", "message": "Welcome to the Teacher app!"})
}
