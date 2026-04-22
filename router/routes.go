package router

import (
	"go_ecommerce/controllers"
	"go_ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	user := r.Group("/user")
	user.Use(middleware.AuthRequired())
	{
		user.GET("/", controllers.ShowUserData)
		user.PUT("/", controllers.EditUserData)
		user.PUT("/:id", controllers.EditUserDataById)
		user.GET("/all", controllers.GetAllUserInfo)
		user.DELETE(":id", controllers.DeleteUser)

		//Product page route
		user.GET("/products", controllers.ShowAllProducts) // for user Dashboard

	}

	product := r.Group("/product")
	product.Use(middleware.AuthRequired())
	{
		product.POST("/", controllers.CreateProduct)
		product.GET("/", controllers.GetVendorProducts)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)

	}

}
