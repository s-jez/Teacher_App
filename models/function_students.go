package models

import (
	"Stachowsky/Teacher_App/config"
)

func CreateStudent(Student *Student) error {
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
func ReadStudent(Student *Student, ID string) error {
	SelectById(Student, ID)
	if err := config.DB.Find(&Student).Error; err != nil {
		return err
	}
	return nil
}
func UpdateStudent(Student *Student, ID string) error {
	if err := config.DB.Where("id = ?", ID).Save(&Student).Error; err != nil {
		return err
	}
	return nil
}
func DeleteStudent(Student *Student, ID string) error {
	SelectById(Student, ID)
	if err := config.DB.Delete(&Student).Error; err != nil {
		return err
	}
	return nil
}
func SelectById(Student *Student, ID string) error {
	if err := config.DB.First(&Student, ID).Error; err != nil {
		return err
	}
	return nil
}
