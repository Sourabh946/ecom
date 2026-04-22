package main

import (
	"fmt"
	"go_ecommerce/config"
	"go_ecommerce/middleware"
	"go_ecommerce/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Oops! API route not found",
			"path":    c.Request.URL.Path,
		})
	})

	r.Use(middleware.ResponseFormatter())
	router.Routes(r)

	// r.Run(":8080")
	r.Run("0.0.0.0:8080")
	fmt.Println("Server is running on port 8080")
}
