package controllers

import (
	"Stachowsky/Teacher_App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(400, err.Error())
	}
	if err := models.AddStudent(&student); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Student has been created!",
		"data":    student,
	})
}
func ReadStudents(c *gin.Context) {
	var students []models.Student
	if err := models.ReadStudents(&students); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been read!",
		"data":    students,
	})
}
func ReadStudentById(c *gin.Context) {
	var student models.Student
	var id = c.Param("id")
	if err := models.ReadStudentById(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been read!",
		"data":    student,
	})
}
func DeleteStudentById(c *gin.Context) {
	var student models.Student
	var id = c.Param("id")
	if err := models.DeleteStudentById(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been deleted!",
		"data":    student,
	})
}
func UpdateStudentById(c *gin.Context) {
	var student models.Student
	var id = c.Param("id")
	if err := models.UpdateStudentById(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been updated!",
		"data":    student,
	})

}
