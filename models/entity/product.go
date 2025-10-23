package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ProductId   string          `json:"product_id" gorm:"type:varchar(8);not null;unique"`
	ProductName string          `json:"product_name" gorm:"type:varchar(100);not null"`
	Premium     decimal.Decimal `json:"premium" gorm:type:decimal(17,2);not null`
	Active      bool            `json:"active" gorm:"type:tinyint(1);not null"`
	CreateAt    time.Time       `json:"create_at" gorm:"type:datetime"`
}

type ProductParameter struct {
	Id             int       `json:"id" gorm:"type:int;autoIncrement;primary_key"`
	ProductId      string    `json:"product_id" gorm:"varchar(8);not null"`
	ParameterName  string    `json:"parameter_name" gorm:"type:varchar(100);not null"`
	ParameterValue string    `json:"parameter_value" gorm:"type:varchar(8);not null"`
	Active         bool      `json:"active" gorm:"type:tinyint(1);not null"`
	CreateAt       time.Time `json:"create_at" gorm:"type:datetime"`
}
