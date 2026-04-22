package models

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey" json:"user_id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:150;not null" json:"email"`
	Password  string    `gorm:"size:200; not null" json:"-"`
	Role_id   int       `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	Roles     Roles     `json:"roles" gorm:"foreignKey:Role_id"`
}
