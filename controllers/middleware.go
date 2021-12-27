package controllers

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(isAuthenticatedRole []uint64) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		if len(isAuthenticatedRole) > 0 {
			RoleID := claims["role_id"].(uint64)
			for _, v := range isAuthenticatedRole {
				if RoleID == v {
					return
				}
			}
			c.AbortWithStatus(404)
			return
		}
	}
}
