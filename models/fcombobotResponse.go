package models

type CreateBotResponse struct {
	StatusCode    int    `json:"status_code"`
	DebugMsg      string `json:"debug_msg"`
	BotId         int    `json:"bot_id"`
	BanReasonText string `json:"ban_reason_text"`
}
