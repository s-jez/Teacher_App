package models

// Student model
type Student struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"firstname" binding:"required" form:"firstname"`
	LastName  string `json:"lastname" binding:"required" form:"lastname"`
	Age       int    `json:"age" binding:"required" form:"age"`
	Grade     int    `json:"grade" form:"grade"`
}

// Set student table name
func (s *Student) TableName() string {
	return "students"
}
