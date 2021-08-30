package controllers

import (
	"Stachowsky/Teacher_App/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	c.Request.ParseForm()
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	student.FirstName = c.PostForm("firstname")
	student.LastName = c.PostForm("lastname")
	student.Age, _ = strconv.Atoi(c.PostForm("age"))
	student.Grade, _ = strconv.Atoi(c.PostForm("grade"))

	if err := models.AddStudent(&student); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.HTML(http.StatusCreated, "form_data.html", gin.H{"data": student, "message": "Student has been created!"})
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

	c.HTML(http.StatusOK, "form_data.html", gin.H{"data": students, "message": "Students has been read!"})

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
		return
	}
	// c.String(200, "ID: %d %s %s %d %d", student.ID, student.FirstName, student.LastName, student.Grade, student.Age)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student has been read!",
		"data":    student,
	})
}
func DeleteStudentById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	student.ID, _ = strconv.ParseUint(c.Request.FormValue("id"), 2, 10)

	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := models.DeleteStudent(&student, id); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
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
	err := c.Bind(&newStudent)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	student.FirstName = newStudent.FirstName
	student.LastName = newStudent.LastName
	student.Age = newStudent.Age
	student.Grade = newStudent.Grade

	if err := models.UpdateStudent(&newStudent, id); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student has been updated!",
		"data":    student,
	})

}
