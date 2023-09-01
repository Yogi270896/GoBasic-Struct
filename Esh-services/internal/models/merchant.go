package models

import "time"

const MerchantTable = "merchant"

type Merchant struct {
	//Id           int       `gorm:"type : int(11) UNSIGNED NOT NULL AUTO_INCREMENT"`
	Email        string    `json:"email" gorm:"unique;type:varchar(100);not null" binding:"required"`
	MobileNumber string    `json:"mobile_number" gorm:"unique; type:BIGINT(30);not null" binding:"min=10,max=15"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}
