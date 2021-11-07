package controllers

import (
	"Stachowsky/Teacher_App/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	student.FirstName = c.Request.FormValue("firstname")
	student.LastName = c.Request.FormValue("lastname")
	student.Age, _ = strconv.Atoi(c.Request.FormValue("age"))
	student.Grade, _ = strconv.Atoi(c.Request.FormValue("grade"))
	if err := models.CreateStudent(&student); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, student)
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
	c.JSON(200, students)
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
	c.JSON(200, student)
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
	nStudent.ID, _ = strconv.ParseUint(c.Request.FormValue("id"), 10, 64)
	nStudent.FirstName = c.Request.FormValue("firstname")
	nStudent.LastName = c.Request.FormValue("lastname")
	nStudent.Age, _ = strconv.Atoi(c.Request.FormValue("age"))
	nStudent.Grade, _ = strconv.Atoi(c.Request.FormValue("grade"))

	student.FirstName = nStudent.FirstName
	student.LastName = nStudent.LastName
	student.Age = nStudent.Age
	student.Grade = nStudent.Grade

	if err := models.UpdateStudent(&nStudent, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, student)
}
func DeleteStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	student.ID, _ = strconv.ParseUint(c.Request.FormValue("id"), 10, 64)
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := models.DeleteStudent(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(204, student)
}
