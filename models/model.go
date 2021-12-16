package models

// Student model
type Student struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Grade     int    `json:"grade"`
}

// User model
type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// Token model
type Token struct {
	Role        string `json:"role`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

// Authentication login model
type Auth struct {
	Email    string `json:"email`
	Password string `json:"password"`
}

// Set students table name
func (s *Student) TableName() string {
	return "students"
}

// Set users table name
func (u *User) TableName() string {
	return "users"
}
