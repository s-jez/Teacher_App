package controllers

import (
	"Stachowsky/Teacher_App/models"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.CreateStudent(&student); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, gin.H{"data": student, "msg": "Student has been created!"})
}
func ReadStudents(c *gin.Context) {
	var students []models.Student
	err := c.Bind(&students)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.ReadStudents(&students); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"data": students, "msg": "Students has been read!"})
}
func ReadStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.ReadStudent(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	if student.ID == 0 {
		c.JSON(404, gin.H{"msg": "Student not found!"})
	}
	c.JSON(200, gin.H{"data": student, "msg": "Student has been read by id!"})
}
func UpdateStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	models.SelectById(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{"msg": "Student not found!"})
	}
	var nStudent models.Student
	err := c.Bind(&nStudent)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	student.FirstName = nStudent.FirstName
	student.LastName = nStudent.LastName
	student.Age = nStudent.Age
	student.Grade = nStudent.Grade
	if err := models.UpdateStudent(&nStudent, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, gin.H{"data": student, "msg": "Student has been updated!"})
}
func DeleteStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.DeleteStudent(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(204, gin.H{"data": student, "msg": "Student has been deleted!"})
}
