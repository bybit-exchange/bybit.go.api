package models

type CloseBotResponse struct {
	StatusCode int    `json:"status_code"`
	DebugMsg   string `json:"debug_msg"`
}
