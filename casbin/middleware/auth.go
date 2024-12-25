package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("username", c.GetHeader("username"))
		c.Set("roles", []string{c.GetHeader("roles")})
		c.Next()
	}
}
