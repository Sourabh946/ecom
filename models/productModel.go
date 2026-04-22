package models

type Products struct {
	Id        uint   `gorm:"PrimaryKey" json:"product_id"`
	Vendor_id uint   `gorm:"not null" json:"vendor_id"`
	Name      string `gorm:"size:200;not null" json:"product_name"`
	Price     int    `gorm:"size:10;not null" json:"product_price"`
	Stock     int    `gorm:"size:10; not null" json:"product_stock"`
	Vendors   Users  `json:"vendors" gorm:"foreignKey:Vendor_id"`
}
