package controllers

import (
	"Stachowsky/Teacher_App/models"

	"github.com/gin-gonic/gin"
)

// Implement CRUD controllers to work with database and perform CRUD operations..
// Create, Read, Update, Delete

func CreateStudent(c *gin.Context) {
	var Student models.Student
	err := c.BindJSON(&Student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err = models.AddStudent(&Student); err != nil {
		c.JSON(500, err.Error())
		return
	}
}
func ReadStudent(c *gin.Context) {
	// Read operations
}
func UpdateStudent(c *gin.Context) {
	// Update operations
}
func DeleteStudent(c *gin.Context) {
	// Delete oeprations
}
