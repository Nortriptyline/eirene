package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// Remove the "Bearer " prefix
		token = strings.Replace(token, "Bearer ", "", 1)

		c.Set("clientToken", token)
		c.Next()
	}
}
