package repositories

import (
	"go_ecommerce/config"
	"go_ecommerce/models"
)

func CreateProducts(product *models.Products) error {
	result := config.DB.Create(product)
	return result.Error
}

func GetProductlistbyVendorId(userId int) ([]*models.Products, error) {
	products := []*models.Products{}
	result := config.DB.Where("vendor_id = ?", userId).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func UpdateProduct(product *models.Products, vendorId int) (int64, error) {
	result := config.DB.Model(&models.Products{}).
		Where("vendor_id = ?", vendorId).
		Where("id = ?", product.Id).
		Updates(product)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func DeleteProduct(vendorId int, product_id int) (int64, error) {

	result := config.DB.Delete(&models.Products{}, product_id).
		Where("vendor_id = ?", vendorId)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil

}

func GetallProductOfVendor() ([]*models.Products, error) {
	products := make([]*models.Products, 0)

	result := config.DB.
		Model(&models.Products{}).
		Joins("JOIN users ON users.id = products.vendor_id").
		Where("users.role_id = ?", 3).
		Preload("Vendors").
		Preload("Vendors.Roles").
		Order("id DESC").
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
