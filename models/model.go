package models

// Student model
type Student struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Grade     int    `json:"grade"`
}

// Set student table name
func (s *Student) TableName() string {
	return "students"
}
