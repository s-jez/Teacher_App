package models

import "Stachowsky/Teacher_App/config"

func AddStudent(Student *Student) error {
	if err := config.DB.Create(&Student).Error; err != nil {
		return err
	}
	return nil
}
