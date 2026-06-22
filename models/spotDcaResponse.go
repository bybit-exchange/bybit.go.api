package models

type CloseDCABotResponse struct {
	StatusCode int    `json:"status_code"`
	BotId      int    `json:"bot_id"`
	CloseMode  int    `json:"close_mode"`
	DebugMsg   string `json:"debug_msg"`
}

type CreateDCABotResponse struct {
	StatusCode    int    `json:"status_code"`
	BotId         int    `json:"bot_id"`
	DebugMsg      string `json:"debug_msg"`
	BanReasonText string `json:"ban_reason_text"`
}

type CreateDCABotRequest struct {
	Parameters              interface{} `json:"parameters"`
	ToolsDiscoveryParameter interface{} `json:"toolsDiscoveryParameter"`
	Channel                 string      `json:"channel"`
}

type BotParameters struct {
	FrequencyInSecond int       `json:"frequency_in_second"`
	QuoteCoin         string    `json:"quote_coin"`
	Pairs             []DCAPair `json:"pairs"`
	MaxInvestAmount   string    `json:"max_invest_amount"`
}

type DCAPair struct {
	Base   string `json:"base"`
	Amount string `json:"amount"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	DebugMsg   string `json:"debug_msg"`
}
