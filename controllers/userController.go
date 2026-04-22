package controllers

import (
	"go_ecommerce/models"
	"go_ecommerce/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user_data struct {
	Id    int
	Name  string
	Email string
}

func ShowUserData(c *gin.Context) {

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(401, gin.H{"error": "Invalid user id"})
		return
	}

	userData, err := services.GetUserByID(userID)
	if err != nil || userData == nil {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	c.Set("response", gin.H{
		"status":  200,
		"message": nil,
		"data":    userData,
	})
}

func EditUserData(c *gin.Context) {
	type editUserRequest struct {
		Name string `json:"name" binding:"required"`
	}
	var req editUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("response", gin.H{
			"status":  400,
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.Set("response", gin.H{
			"status":  401,
			"error":   "User not found",
			"message": "Missing user context",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.Set("response", gin.H{
			"status":  401,
			"error":   "Invalid user id",
			"message": "Invalid user context type",
		})
		return
	}

	userData, err := services.UpdateUsers(userID, req.Name)
	if err != nil || userData == nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Update failed",
			"message": "Could not update user details",
		})
		return
	}

	c.Set("response", gin.H{
		"status":  200,
		"message": "User updated successfully",
		"data":    userData,
	})
}

func EditUserDataById(c *gin.Context) {
	type editUserRequest struct {
		Name string `json:"name" binding:"required"`
	}

	var req editUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	// Get user_id safely
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDRaw.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user_id type"})
		return
	}

	// Parse ID param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if id != userID {

		// Check admin
		isAdmin, err := services.CheckIfUserAdmin(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		if !isAdmin {
			c.JSON(403, gin.H{"error": "Forbidden"})
			return
		}
	}

	// Update user
	userData, err := services.UpdateUsers(id, req.Name)
	if err != nil || userData == nil {
		c.JSON(500, gin.H{
			"error":   "Update failed",
			"message": "Could not update user details",
		})
		return
	}

	// Response
	c.JSON(200, gin.H{
		"message": "User updated successfully",
		"data": gin.H{
			"id":   userData.ID,
			"name": userData.Name,
		},
	})
}

func GetAllUserInfo(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	isAdmin, err := services.CheckIfUserAdmin(userId.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if !isAdmin {
		userData, err := services.GetUserByID(userId.(int))
		if err != nil || userData == nil {
			c.JSON(401, gin.H{"error": "User not found"})
			return
		}

		c.Set("response", gin.H{
			"status":  200,
			"message": nil,
			"data":    []*models.Users{userData},
		})
		return
	}

	allUserData, err := services.GetallUserInfo()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if allUserData == nil {
		c.JSON(404, gin.H{"error": "No users found"})
		return
	}
	c.Set("response", gin.H{
		"status":  200,
		"message": "Users fetched successfully",
		"data":    allUserData,
	})
}

func DeleteUser(c *gin.Context) {

	// ✅ Get ID from URL param
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// ✅ Get logged-in user from middleware
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	// ✅ Check if admin
	isAdmin, err := services.CheckIfUserAdmin(userId.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if !isAdmin {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// ✅ Delete user
	err = services.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.Set("response", gin.H{
		"status":  200,
		"message": "User deleted successfully",
	})

}

func ShowAllProducts(c *gin.Context) {

	getAllProduct, err := services.FetchAllVendorProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if len(getAllProduct) == 0 {
		c.JSON(404, gin.H{"error": "No products found"})
		return
	}

	c.Set("response", gin.H{
		"status":  200,
		"message": "Products fetched successfully",
		"data":    getAllProduct,
	})

}
