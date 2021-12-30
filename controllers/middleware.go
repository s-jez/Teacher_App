package controllers

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(Role []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(404, "No authorization header!")
			c.Abort()
		}
		extractedToken := strings.Split(authHeader, "Bearer")
		if len(extractedToken) != 2 {
			c.JSON(404, "Bearer token in not proper format")
		}
		stringToken := strings.TrimSpace(extractedToken[1])
		token, err := ExtractAccessToken(stringToken)
		if err != nil {
			c.JSON(401, "No access to site!")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(401)
			return
		}
		if len(Role) > 0 {
			RoleID := int(claims["role"].(float64))
			for _, Value := range Role {
				if Value == RoleID {
					return
				}
			}
			c.AbortWithStatus(403)
			return
		}
	}
}
