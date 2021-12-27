package controllers

import (
	"Stachowsky/Teacher_App/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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
	claims["role"] = models.GetUserRole(&user, userid)
	claims["user_id"] = userid
	claims["exp"] = token.Expiration
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Token, _ = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET_TOKEN")))
	return token
}
func ExtractAccessToken(stringToken string) (*jwt.Token, error) {
	return jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET_TOKEN")), nil
	})
}
func ExtractRefreshToken(stringToken string) (*jwt.Token, error) {
	return jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET_TOKEN")), nil
	})
}
func RefreshToken(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.GetHeader("Authorization")
	stringToken := authHeader[len(BEARER_SCHEMA):]
	token, err := ExtractRefreshToken(stringToken)
	if err != nil {
		c.JSON(404, "Access deny!")
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["user_id"].(uint64)
	role := claims["role"].(uint64)

	newPairOfTokens := models.Tokens{}
	newPairOfTokens.AccessToken = CreateAccessToken(id, role, models.User{})
	newPairOfTokens.RefreshToken = CreateRefreshToken(id, role, models.User{})
	c.JSON(200, newPairOfTokens)
}
