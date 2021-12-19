package models

import "Stachowsky/Teacher_App/config"

func CreateUser(User *User) error {
	if err := config.DB.Table("users").Create(&User).Error; err != nil {
		return err
	}
	return nil
}
func CheckUserName(User *User, UserName string) error {
	if err := config.DB.Table("users").Where("User_Name = ?", UserName).First(&User).Error; err != nil {
		return err
	}
	return nil
}
func CheckUserEmail(User *User, Email string) error {
	if err := config.DB.Table("users").Where("Email = ?", Email).First(&User).Error; err != nil {
		return err
	}
	return nil
}
func GetUserRole(User *User, userid uint64) *uint64 {
	if err := config.DB.Table("users").Where("ID = ?", userid).First(&User).Error; err != nil {
		return &User.RoleID
	}
	return nil
}
func GetUserEmail(User *User, userid uint64) *string {
	if err := config.DB.Table("users").Where("ID = ?", userid).First(&User).Error; err != nil {
		return &User.Email
	}
	return nil
}
func GetUserId(User *User, userid uint64) *uint64 {
	if err := config.DB.Table("users").Where("ID = ?", userid).First(&User).Error; err != nil {
		return &User.ID
	}
	return nil
}
