package middleware

import (
	"github.com/gin-gonic/gin"
)

func ResponseFormatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		resp, exists := c.Get("response")
		if !exists {
			return
		}

		// Some Gin handlers may implicitly write a 200 header even if they didn't write a body.
		// Only skip if a non-empty body has already been written.
		if c.Writer.Written() && c.Writer.Size() > 0 {
			return
		}

		h, ok := resp.(gin.H)
		if !ok {
			// Fallback: avoid panic if some handler sets a non-gin.H response.
			c.JSON(500, gin.H{
				"status":  500,
				"error":   "Internal Server Error",
				"message": "Invalid response format",
			})
			return
		}

		statusVal, exists := h["status"]
		if !exists {
			c.JSON(500, gin.H{
				"status":  500,
				"error":   "Internal Server Error",
				"message": "Missing response status",
			})
			return
		}

		status := 500
		switch v := statusVal.(type) {
		case int:
			status = v
		case int32:
			status = int(v)
		case int64:
			status = int(v)
		case float64:
			status = int(v)
		}

		c.JSON(status, resp)
	}
}
