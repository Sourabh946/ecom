package middleware

import (
	"strings"

	"go_ecommerce/helpers"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token from header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Expect: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenStr := parts[1]

		// Verify token
		claims, err := helpers.VerifyJWT(tokenStr)
		if err != nil {

			if err.Error() == "token expired" {
				c.JSON(401, gin.H{"error": "Token expired"})
			} else {
				c.JSON(401, gin.H{"error": "Invalid token"})
			}

			c.Abort()
			return
		}

		// Save user data for next handlers
		c.Set("user_id", claims.User_id)
		c.Set("username", claims.Username)

		c.Next() // VERY IMPORTANT
	}
}
