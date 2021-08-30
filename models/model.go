package models

// Student model
type Student struct {
	ID        uint64 `gorm:"primaryKey" json:"id" form:"id"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Age       int    `json:"age" form:"age"`
	Grade     int    `json:"grade" form:"grade"`
}

// Set student table name
func (s *Student) TableName() string {
	return "students"
}
