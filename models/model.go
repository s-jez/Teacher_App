package models

// Student model
type Student struct {
	ID        uint   `gorm:"primaryKey" json:"id" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Grade     int    `json:"grade"`
}

// Set student table name
func (s *Student) TableName() string {
	return "students"
}
