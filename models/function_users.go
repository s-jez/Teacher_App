package models

import "Stachowsky/Teacher_App/config"

func CreateUser(User *User) error {
	if err := config.DB.Table("users").Create(&User).Error; err != nil {
		return err
	}
	return nil
}
func CheckUserEmail(User *User, Email string) error {
	if err := config.DB.Table("users").Where("Email = ?", Email).Find(&User).Error; err != nil {
		return err
	}
	return nil
}
func GetUserRole(User *User, userid uint64) error {
	if err := config.DB.Table("users").Where("ID = ?", userid).Find(&User.RoleID).Error; err != nil {
		return err
	}
	return nil
}
func GetUserEmail(User *User, userid uint64) error {
	if err := config.DB.Table("users").Where("ID = ?", userid).Find(&User.Email).Error; err != nil {
		return err
	}
	return nil
}
func GetUserId(User *User, userid uint64) error {
	if err := config.DB.Table("users").Where("ID = ?", userid).Find(&User.ID).Error; err != nil {
		return err
	}
	return nil
}
