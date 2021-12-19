package controllers

import (
	"Stachowsky/Teacher_App/models"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VeryfiyToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(404, "No authorization header!")
		c.Abort()
	}
	extractedToken := strings.Split(authHeader, "Bearer ")
	stringToken := authHeader[len(extractedToken):]
	token, err := ExtractAccessToken(stringToken)
	if err != nil {
		c.JSON(401, "No access to site!")
	}
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role
}
func IsAdmin(c *gin.Context) {
	role := VeryfiyToken(c)
	if role == "3" {
		c.Next()
	} else {
		c.JSON(404, "Invalid role!")
	}
}
func IsDev(c *gin.Context) {
	role := VeryfiyToken(c)
	if role == "2" || role == "3" {
		c.Next()
	} else {
		c.JSON(404, "Invalid role")
	}
}
func Logout(c *gin.Context) {
	
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
