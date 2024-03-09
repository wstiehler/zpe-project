package middlewares

import "github.com/gin-gonic/gin"

func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
	}
}
