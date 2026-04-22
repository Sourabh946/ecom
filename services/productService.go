package services

import (
	"go_ecommerce/dto"
	"go_ecommerce/models"
	"go_ecommerce/repositories"
)

func AddProduct(req *dto.ProductDTO, userID int) (*models.Products, error) {
	product := &models.Products{
		Vendor_id: uint(userID),
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
	}
	err := repositories.CreateProducts(product)

	if err != nil {
		return nil, err
	}
	return product, nil

}

func GetAllVendorProducts(userId int) ([]*models.Products, error) {

	result, err := repositories.GetProductlistbyVendorId(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateProduct(req *dto.ProductDTO, productID int, vendorId int) (int64, error) {
	product := &models.Products{
		Id:    uint(productID),
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
	return repositories.UpdateProduct(product, vendorId)

}

func DeleteProduct(vendorId int, product_id int) (int64, error) {
	return repositories.DeleteProduct(vendorId, product_id)
}

func FetchAllVendorProducts() ([]*models.Products, error) {

	return repositories.GetallProductOfVendor()
}
