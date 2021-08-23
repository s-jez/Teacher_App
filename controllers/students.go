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
	// student.FirstName = c.PostForm("FirstName")
	// student.LastName = c.PostForm("LastName")

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
	id := c.Param("id")

	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := models.ReadStudent(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	if student.ID == 0 {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been read!",
		"data":    student,
	})
}
func DeleteStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(400, err.Error())
	}

	if err := models.DeleteStudent(&student, id); err != nil {
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
	id := c.Param("id")
	models.SelectById(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "We not found this student!",
			"data":    student,
		})
	}
	var newStudent models.Student
	err := c.BindJSON(&newStudent)
	if err != nil {
		c.JSON(400, err.Error())
	}
	student.FirstName = newStudent.FirstName
	student.LastName = newStudent.LastName
	student.Age = newStudent.Age
	student.Grade = newStudent.Grade

	if err := models.UpdateStudent(&newStudent, id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been updated!",
		"data":    student,
	})

}
