package models

type FGridBotSingleCloseResponse struct {
	StatusCode int    `json:"status_code"`
	DebugMsg   string `json:"debug_msg"`
	BotId      int    `json:"bot_id"`
}
