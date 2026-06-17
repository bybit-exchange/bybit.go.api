package models

type ApiError struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}
