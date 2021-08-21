package models

import (
	"Stachowsky/Teacher_App/config"
)

func AddStudent(Student *Student) error {
	if err := config.DB.Create(&Student).Error; err != nil {
		return err
	}
	return nil
}
func ReadStudents(Students *[]Student) error {
	if err := config.DB.Find(&Students).Error; err != nil {
		return err
	}
	return nil
}
func ReadStudentById(Student *Student, ID string) error {
	if err := config.DB.Where("ID = ?", ID).First(&Student).Find(&Student).Error; err != nil {
		return err
	}
	return nil
}
func UpdateStudentById(Student *Student, ID string) error {
	if err := config.DB.Where("ID = ?", ID).Updates(&Student).Save(&Student).Error; err != nil {
		return err
	}
	return nil
}
func DeleteStudentById(Student *Student, ID string) error {
	if err := config.DB.Where("ID = ?", ID).First(&Student).Delete(&Student).Error; err != nil {
		return err
	}
	return nil
}
