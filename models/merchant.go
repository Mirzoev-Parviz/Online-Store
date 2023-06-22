package models

import "gorm.io/gorm"

type Merchant struct {
	ID          int    `json:"id" gorm:"primarykey"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active" gorm:"not null; default: true"`
}

type MerchantProduct struct {
	gorm.Model
	ProductID  int  `json:"product_id"`
	MerchantID int  `json:"merchant_id"`
	Quantity   int  `json:"quantity" gorm:"not null; default: 1"`
	IsActive   bool `json:"is_active" gorm:"not null; default: true"`
}
