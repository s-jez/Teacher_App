package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Page(c *gin.Context) {
	if c != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Teacher app in GO!", "message": "Welcome to the Teacher app!"})
	} else {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"title": "Teacher app in GO!", "message": "Error, page doesn't work.."})
	}
}
