package dto

type RequestAuth struct {
	AgentId  string `json:"agent_id" validate:required`
	Password string `json:"password" validate:required`
}
