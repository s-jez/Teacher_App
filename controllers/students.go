package controllers

import (
	"Stachowsky/Teacher_App/config"
	"Stachowsky/Teacher_App/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	config.DB.Create(&student)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successful, student has been created!",
		"data":    student,
	})
}
func ReadStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful, all students has been read!",
		"data":    students,
	})
}
func DeleteStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", c.Param("id")).Find(&student).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}
	config.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful, student has been deleted!",
		"data":    student,
	})
}
func UpdateStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := config.DB.Where("id = ?", c.Param("id")).Find(&student).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}
	config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful, student has been updated!",
		"data":    student,
	})
}
