package controllers

import (
	"Stachowsky/Teacher_App/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateAccessToken(userid uint64, role int64, user *models.User) *models.Token {
	var token = &models.Token{}
	token.Expiration = time.Now().Add(time.Minute * 15).Unix()
	// create the claims
	claims := jwt.MapClaims{
		"authorized": true,
		"role":       models.GetUserRole(userid),
		"user_id":    userid,
		"exp":        token.Expiration,
	}
	// create the JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Token, _ = at.SignedString([]byte(os.Getenv("ACCESS_SECRET_TOKEN")))
	return token
}

func CreateRefreshToken(userid uint64, role int64, user *models.User) *models.Token {
	var token = &models.Token{}
	token.Expiration = time.Now().Add(time.Hour * 24 * 7).Unix()
	// create the claims
	claims := jwt.MapClaims{
		"role":    models.GetUserRole(userid),
		"user_id": userid,
		"exp":     token.Expiration,
	}
	// create the jwt
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
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	stringToken := authHeader[len(BEARER_SCHEMA):]
	token, err := ExtractRefreshToken(stringToken)
	if err != nil {
		c.JSON(404, "Access deny!")
		return
	}
	// cast claims to MapClaims
	claims := token.Claims.(jwt.MapClaims)
	// extract sub-claim
	id := claims["user_id"].(float64)
	role := claims["role"].(float64)
	// create pair of tokens
	newPairOfTokens := models.Tokens{}
	newUser := models.User{}
	newPairOfTokens.AccessToken = CreateAccessToken(uint64(id), int64(role), &newUser)
	newPairOfTokens.RefreshToken = CreateRefreshToken(uint64(id), int64(role), &newUser)
	c.JSON(200, newPairOfTokens)
}
