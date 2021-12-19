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
	RoleID   uint64 `json:"roleid"`
}

// Tokens model
type Tokens struct {
	AccessToken  *Token
	RefreshToken *Token
}

// Token model
type Token struct {
	Token      string
	Expiration int64
}

// Authentication login model
type Auth struct {
	UserName string `json:"username"`
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
