package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "form_create.html", gin.H{"title": "Teacher app in GO!", "message": "Create student!"})
}
func FormRead(c *gin.Context) {
	c.HTML(http.StatusOK, "form_read.html", gin.H{"title": "Teacher app in GO!", "message": "Read student!"})
}
func FormUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "form_update.html", gin.H{"title": "Teacher app in GO!", "message": "Update student!"})
}
func FormDelete(c *gin.Context) {
	c.HTML(http.StatusOK, "form_delete.html", gin.H{"title": "Teacher app in GO!", "message": "Delete student!"})
}
