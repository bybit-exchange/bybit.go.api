package models

type GetDetailResponse struct {
	StatusCode int         `json:"status_code"`
	DebugMsg   string      `json:"debug_msg"`
	Detail     interface{} `json:"detail"`
}

type BotDetail struct {
	BotId                    int           `json:"bot_id"`
	BotMode                  string        `json:"bot_mode"`
	BotDisplayStatus         string        `json:"bot_display_status"`
	Leverage                 string        `json:"leverage"`
	TotalMargin              string        `json:"total_margin"`
	TotalPnl                 string        `json:"total_pnl"`
	TotalPnlPer              string        `json:"total_pnl_per"`
	Equity                   string        `json:"equity"`
	Roi                      string        `json:"roi"`
	TotalApr                 string        `json:"total_apr"`
	AdjustedPositionNum      int           `json:"adjusted_position_num"`
	LastAdjustPositionTime   int           `json:"last_adjust_position_time"`
	CreateTime               int           `json:"create_time"`
	EndTime                  int           `json:"end_time"`
	CloseCode                string        `json:"close_code"`
	WsToken                  string        `json:"ws_token"`
	SymbolSettings           []interface{} `json:"symbol_settings"`
	SlPercent                string        `json:"sl_percent"`
	TpPercent                string        `json:"tp_percent"`
	TotalPositionValue       string        `json:"total_position_value"`
	MarginBalance            string        `json:"margin_balance"`
	AvailableBalance         string        `json:"available_balance"`
	Imr                      string        `json:"imr"`
	Mmr                      string        `json:"mmr"`
	FundingFee               string        `json:"funding_fee"`
	AdjustPositionMode       string        `json:"adjust_position_mode"`
	AdjustPositionPercent    string        `json:"adjust_position_percent"`
	AdjustPositionTimeInterval int         `json:"adjust_position_time_interval"`
	BotName                  string        `json:"bot_name"`
	RunTimeDuration          int           `json:"run_time_duration"`
	SettlementAssets         string        `json:"settlement_assets"`
	StopType                 string        `json:"stop_type"`
	FollowNum                int           `json:"follow_num"`
	UsedRewardAmount         string        `json:"used_reward_amount"`
	UsedRewardId             string        `json:"used_reward_id"`
	TotalBonus               string        `json:"total_bonus"`
	RealizedPnl              string        `json:"realized_pnl"`
	UnrealizedPnl            string        `json:"unrealized_pnl"`
	TotalExecFee             string        `json:"total_exec_fee"`
	LastModifySettingTime    int           `json:"last_modify_setting_time"`
	MidwayTransfer           string        `json:"midway_transfer"`
	ExitEquity               string        `json:"exit_equity"`
	TrailingStopPercent      string        `json:"trailing_stop_percent"`
	SlExitEquity             string        `json:"sl_exit_equity"`
	TpExitEquity             string        `json:"tp_exit_equity"`
}

type SymbolSetting struct {
	Symbol                string `json:"symbol"`
	BaseToken             string `json:"base_token"`
	QuoteToken            string `json:"quote_token"`
	TargetPositionPercent string `json:"target_position_percent"`
	Side                  int    `json:"side"`
	SymbolId              int    `json:"symbol_id"`
}

type GetLimitResponse struct {
	StatusCode                 int         `json:"status_code"`
	DebugMsg                   string      `json:"debug_msg"`
	CheckCode                  string      `json:"check_code"`
	AdjustPositionPercent      interface{} `json:"adjust_position_percent"`
	AdjustPositionTimeInterval interface{} `json:"adjust_position_time_interval"`
	InitMargin                 interface{} `json:"init_margin"`
	SlPercent                  interface{} `json:"sl_percent"`
	TpPercent                  interface{} `json:"tp_percent"`
	Leverage                   interface{} `json:"leverage"`
	TrailingStopPercent        interface{} `json:"trailing_stop_percent"`
}

type GetLimitRequest struct {
	Leverage                   string        `json:"leverage"`
	InitMargin                 string        `json:"init_margin"`
	AdjustPositionMode         int           `json:"adjust_position_mode"`
	AdjustPositionPercent      string        `json:"adjust_position_percent"`
	AdjustPositionTimeInterval int           `json:"adjust_position_time_interval"`
	SymbolSettings             []interface{} `json:"symbol_settings"`
	SlPercent                  string        `json:"sl_percent"`
	TpPercent                  string        `json:"tp_percent"`
	NeedToSlippage             bool          `json:"need_to_slippage"`
	AppName                    string        `json:"app_name"`
	TrailingStopPercent        string        `json:"trailing_stop_percent"`
}

type Range struct {
	Min string `json:"min"`
	Max string `json:"max"`
}
