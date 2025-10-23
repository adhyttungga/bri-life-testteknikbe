package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransId   int64           `json:"trans_id" gorm:"type:bigint;primary_key;autoIncrement;<-:create"`
	AgentId   string          `json:"agent_id" gorm:"type:varchar(8);not null"`
	ProductId string          `json:"product_id" gorm:"type:varchar(8);not null"`
	Nama      string          `json:"nama" gorm:"type:varchar(100);not null"`
	Usia      int             `json:"usia" gorm:"type:int;not null"`
	Premium   decimal.Decimal `json:"premium" gorm:"type:decimal(17,2);not null"`
	CreateAt  time.Time       `json:"create_at" gorm:"type:datetime"`
}
