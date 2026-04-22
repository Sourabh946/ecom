package controllers

import (
	"go_ecommerce/dto"
	"go_ecommerce/helpers"
	"go_ecommerce/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req dto.Login_request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("response", gin.H{
			"status":  400,
			"error":   "Invalid request",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	user, err := services.Login(req.Email, req.Password)
	if err != nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Login failed",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	token_data, refresh_token, err := helpers.GenerateToken(int(user.ID), user.Name)
	if err != nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Token generation failed",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("response", gin.H{
		"status":        200,
		"message":       "Login successful",
		"access_token":  token_data,
		"refresh_token": refresh_token,
	})
}

func Register(c *gin.Context) {

	var req dto.Registraton_request
	if err := c.ShouldBindJSON(&req); err != nil {
		// Instead of c.JSON directly, set error in context for middleware
		c.Set("response", gin.H{
			"status":  400,
			"error":   "Invalid request",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	// call the service layer
	// Call service layer
	err := services.Register(req.Name, req.Email, req.Password, req.Role_id)
	if err != nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Registration failed",
			"message": err.Error(),
		})
		// c.Abort()
		return
	}

	// Set successful response in context
	c.Set("response", gin.H{
		"status":  200,
		"message": "Registration successful",
	})
}
