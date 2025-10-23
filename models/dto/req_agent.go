package dto

type RequestAgent struct {
	AgentId   string `json:"agent_id" validate:"required"`
	AgentName string `json:"agent_name" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Active    bool   `json:"active" validate:"required"`
}

type RequestDeleteAgent struct {
	AgentId string `json:"agent_id" validate:"required"`
}
