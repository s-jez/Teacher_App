package controllers

import (
	"Stachowsky/Teacher_App/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login
func LoginUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(422, "Invalid json provided")
		return
	}
	var user = models.User{
		ID:       1,
		UserName: "root",
		Password: "root",
		Email:    "root",
		Role: "admin",
	}
	if user.UserName != u.UserName || user.Password != u.Password || user.Email != u.Email || user.Role != u.Role {
		c.JSON(401, "Provide valid login details!")
	}
	token, err := CreateAccessToken(user.Email, user.Role)
	if err != nil {
		c.JSON(422, err.Error())
		return
	}
	c.JSON(200, token)
}

// Register
func Register(c *gin.Context) {

}

// Create Access Token
func CreateAccessToken(email string, role string) (*models.Token, error) {
	var err error
	var token = &models.Token{}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.TokenString, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET_TOKEN")))
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Create Refresh Token
func CreateRefreshToken(email, role map[string]bool) (*models.Token, error) {
	var err error
	var token = &models.Token{}
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.TokenString, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET_TOKEN")))
	if err != nil {
		return nil, err
	}
	return token, nil

}

// Extract Token
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Verify Token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Token Valid
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// func ExtractTokenMetadata(r *http.Request) (*models.Token, error) {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		accessUuid, ok := claims["access_uuid"].(string)
// 		if !ok {
// 			return nil, err
// 		}
// 		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &models.Token{
// 			AccessUuid: accessUuid,
// 			UserId:     userId,
// 		}, nil
// 	}
// 	return nil, err
// }
func GenerateHashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
func CheckHashPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
