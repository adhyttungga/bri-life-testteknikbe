package dto

type ResponseError struct {
	ResponseCode string `json:"ResponseCode"`
	ResponseDesc string `json:"ResponseDesc"`
}
