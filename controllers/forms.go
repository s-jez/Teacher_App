package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "form_create.html", gin.H{"title": "Teacher app in GO!", "message": "Welcome to the Teacher app!"})
}
func FormRead(c *gin.Context) {
	c.HTML(http.StatusOK, "form_read.html", gin.H{"title": "Teacher app in GO!", "message": "Welcome to the Teacher app!"})
}
func FormUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "form_update.html", gin.H{"title": "Teacher app in GO!", "message": "Welcome to the Teacher app!"})
}
func FormDelete(c *gin.Context) {
	c.HTML(http.StatusOK, "form_delete.html", gin.H{"title": "Teacher app in GO!", "message": "Welcome to the Teacher app!"})
}
