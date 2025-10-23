package dto

type RequestCreateTrans struct {
	AgentId   string `json:"agent_id" validate:required`
	ProductId string `json:"product_id" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Usia      int    `json:"usia" validate:"required"`
	Premium   string `json:"premium" validate:"required"`
}

type RequestTrans struct {
	TransId   int64  `json:"trans_id" validate:required`
	AgentId   string `json:"agent_id" validate:required`
	ProductId string `json:"product_id" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Usia      int    `json:"usia" validate:"required"`
	Premium   string `json:"premium" validate:"required"`
}
