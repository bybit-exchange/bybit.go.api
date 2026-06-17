package models

type CreateFMartBotResponse struct {
	StatusCode    int    `json:"status_code"`
	DebugMsg      string `json:"debug_msg"`
	BanReasonText string `json:"ban_reason_text"`
	BotId         int    `json:"bot_id"`
}
