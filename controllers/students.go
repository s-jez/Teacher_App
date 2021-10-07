package controllers

import (
	"Stachowsky/Teacher_App/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.Bind(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.Request.ParseForm()

	if err := models.AddStudent(&student); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": student, "message": "Student has been created!"})
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

	c.JSON(http.StatusOK, gin.H{"data": students, "message": "Students has been read!"})

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
	c.JSON(http.StatusOK, gin.H{"message": "Student has been read!", "data": student})
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
	c.JSON(http.StatusNoContent, gin.H{"message": "Student has been deleted!", "data": student})
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

	c.JSON(http.StatusCreated, gin.H{"message": "Student has been updated!", "data": student})
}
func Login(c *gin.Context) {
	session := sessions.Default(c)
	var user models.User
	// Validation form
	// login := c.PostForm("login")
	// pass := c.PostForm("pass")

	// if strings.Trim(user.Name, " ") == "" || strings.Trim(user.Pass, " ") == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Empty parameters in login or pass!"})
	// 	return
	// }
	// if user.Name != "hello" || user.Pass != "world" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed!"})
	// 	return
	// }
	session.Set(gin.AuthUserKey, user.Name)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session!!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully authenticated user!"})
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(gin.AuthUserKey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token!"})
		return
	}
	if user != nil {
		session.Delete(gin.AuthUserKey)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Sucessfully, logout!"})
	}
}
func AuthenticationRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(gin.AuthUserKey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unautharized!"})
	}
	c.Next()
}
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are loggin in!"})
}
func Account(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(gin.AuthUserKey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
