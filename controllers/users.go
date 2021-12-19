package controllers

import (
	"Stachowsky/Teacher_App/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	var u models.Auth
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}
	var user models.User
	if err := models.CheckUserEmail(&user, user.Email); err != nil {
		c.JSON(400, err.Error())
		return
	}
	// if err := CheckHashPassword(user.Password, u.Password); !err {
	// 	c.JSON(400, err)
	// 	return
	// }
	tokens := models.Tokens{}
	tokens.AccessToken = CreateAccessToken(user.ID, user.RoleID, user)
	tokens.RefreshToken = CreateRefreshToken(user.ID, user.RoleID, user)
	c.JSON(200, tokens)
}

func RegisterUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}
	u.Password, _ = GenerateHashPassword(u.Password)
	u.RoleID = 1
	if err := models.CreateUser(&u); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, u)
}
func CreateAccessToken(userid uint64, role uint64, user models.User) *models.Token {
	var token = &models.Token{}
	token.Expiration = time.Now().Add(time.Minute * 15).Unix()

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["role"] = models.GetUserRole(&user, userid)
	claims["email"] = models.GetUserEmail(&user, userid)
	claims["user_id"] = userid
	claims["exp"] = token.Expiration
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Token, _ = at.SignedString([]byte(os.Getenv("ACCESS_SECRET_TOKEN")))
	return token
}

func CreateRefreshToken(userid uint64, role uint64, user models.User) *models.Token {
	var token = &models.Token{}
	token.Expiration = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["role"] = models.GetUserRole(&user, userid)
	claims["email"] = models.GetUserEmail(&user, userid)
	claims["user_id"] = userid
	claims["exp"] = token.Expiration
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Token, _ = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET_TOKEN")))
	return token
}
func GenerateHashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
func CheckHashPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
