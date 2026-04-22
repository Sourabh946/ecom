package controllers

import (
	"go_ecommerce/dto"
	"go_ecommerce/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {

	var req dto.ProductDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("response", gin.H{
			"status":  400,
			"error":   "Invalid request",
			"message": err.Error(),
		})
		c.Abort()
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

	isVendor, err := services.CheckIfUserVendor(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if !isVendor {
		c.JSON(401, gin.H{"error": "Unauthorized User"})
		return
	}

	createProduct, err := services.AddProduct(&req, userID)
	if err != nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Failed to create product",
			"message": err.Error(),
		})
		return
	}

	c.Set("response", gin.H{
		"status":  201,
		"message": "Product created successfully",
		"data":    createProduct,
	})

}

func GetVendorProducts(c *gin.Context) {

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

	isVendor, err := services.CheckIfUserVendor(userID)
	if err != nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Internal server error",
			"message": err.Error(),
		})
		return
	}
	if !isVendor {
		c.Set("response", gin.H{
			"status":  401,
			"error":   "Unauthorized User",
			"message": "Only vendors can access products",
		})
		return
	}

	products, err := services.GetAllVendorProducts(userID)
	if err != nil {
		c.Set("response", gin.H{
			"status":  500,
			"error":   "Failed to fetch products",
			"message": err.Error(),
		})
		return
	}

	c.Set("response", gin.H{
		"status":  200,
		"message": "Products fetched successfully",
		"data":    products,
	})
}

func UpdateProduct(c *gin.Context) {
	vendorId, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := vendorId.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user_id type"})
		return
	}

	isVendor, err := services.CheckIfUserVendor(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if !isVendor {
		c.JSON(401, gin.H{"error": "Unauthorized User"})
		return
	}

	var req dto.ProductDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product id"})
		return
	}

	updatedProduct, err := services.UpdateProduct(&req, productID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}
	if updatedProduct == 0 {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Product updated successfully", "data": updatedProduct})

}

func DeleteProduct(c *gin.Context) {

	vendorId, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized action"})
		return
	}

	userID, ok := vendorId.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user_id type"})
		return
	}

	isVendor, err := services.CheckIfUserVendor(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if !isVendor {
		c.JSON(401, gin.H{"error": "Unauthorized User"})
		return
	}

	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product id"})
		return
	}

	deletedProduct, err := services.DeleteProduct(userID, productID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}
	if deletedProduct == 0 {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Product deleted successfully", "data": deletedProduct})
	return

}
