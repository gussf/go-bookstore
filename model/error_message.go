package model

type ErrorReport struct {
	Message  string `json:"message"`
	HttpCode int    `json:"httpCode"`
}
