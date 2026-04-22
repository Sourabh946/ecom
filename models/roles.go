package models

type Roles struct {
	Id        uint   `gorm:"primarykey" json:"role_id"`
	Role_name string `gorm:"role_name" json:"role_name"`
}
