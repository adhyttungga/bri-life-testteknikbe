package dto

type ResponseSuccess struct {
	ResponseCode string `json:"ResponseCode"`
	ResponseDesc string `json:"ResponseDesc"`
	Data         any
}

type AgentData struct {
	AgentId   string `json:"AgentId"`
	AgentName string `json:"AgentName"`
}

type TransactionData struct {
	TransId string `json:"TransId"`
}

type AuthData struct {
	AgentId     string `json:"AgentId"`
	AgentName   string `json:"AgentName"`
	AccessToken string `json:"AccessToken"`
}
