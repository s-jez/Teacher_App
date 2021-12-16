package controllers

import "github.com/gin-gonic/gin"

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
func isAdmin(c *gin.Context) {

}
func isDev(c *gin.Context) {
	
}
