package models

type GetFMartDetailResponse struct {
	StatusCode  int         `json:"status_code"`
	DebugMsg    string      `json:"debug_msg"`
	FmartDetail interface{} `json:"fmart_detail"`
}

type FMartBotDetail struct {
	BotId                      int    `json:"bot_id"`
	Symbol                     string `json:"symbol"`
	BotDisplayStatus           string `json:"bot_display_status"`
	FmartMode                  string `json:"fmart_mode"`
	BaseToken                  string `json:"base_token"`
	QuoteToken                 string `json:"quote_token"`
	Leverage                   string `json:"leverage"`
	SingleRoundTargetProfitPer string `json:"single_round_target_profit_per"`
	AddPosNum                  string `json:"add_pos_num"`
	AddPosPer                  string `json:"add_pos_per"`
	PriceFloatPer              string `json:"price_float_per"`
	EntryPrice                 string `json:"entry_price"`
	SlPer                      string `json:"sl_per"`
	TotalMargin                string `json:"total_margin"`
	Pnl                        string `json:"pnl"`
	PnlPer                     string `json:"pnl_per"`
	RealizedPnl                string `json:"realized_pnl"`
	UnrealizedPnl              string `json:"unrealized_pnl"`
	CompletedRounds            int    `json:"completed_rounds"`
	CurrentRound               int    `json:"current_round"`
	CurrentAddedPosNum         int    `json:"current_added_pos_num"`
	PosSize                    string `json:"pos_size"`
	AvgPosPrice                string `json:"avg_pos_price"`
	PosBalance                 string `json:"pos_balance"`
	OrderBalance               string `json:"order_balance"`
	AvailableBalance           string `json:"available_balance"`
	MarkPrice                  string `json:"mark_price"`
	EstimatedLiqPrice          string `json:"estimated_liq_price"`
	ActualLeverage             string `json:"actual_leverage"`
	AutoCycle                  string `json:"auto_cycle"`
	OperationTime              int    `json:"operation_time"`
	CreateTime                 int    `json:"create_time"`
	ModifyTime                 int    `json:"modify_time"`
	EndTime                    int    `json:"end_time"`
	StopType                   string `json:"stop_type"`
	SettlementAssets           string `json:"settlement_assets"`
	TickSize                   string `json:"tick_size"`
	ClosePrice                 string `json:"close_price"`
	CurrentPrice               string `json:"current_price"`
	RunningDuration            int    `json:"running_duration"`
	CloseCode                  string `json:"close_code"`
	NetFundingFee              string `json:"net_funding_fee"`
	TotalApr                   string `json:"total_apr"`
	FollowNum                  int    `json:"follow_num"`
	UsedRewardAmount           string `json:"used_reward_amount"`
	UsedRewardId               string `json:"used_reward_id"`
	TotalBonus                 string `json:"total_bonus"`
	AdlRankIndicator           int    `json:"adl_rank_indicator"`
}

type GetFMartLimitResponse struct {
	StatusCode         int         `json:"status_code"`
	DebugMsg           string      `json:"debug_msg"`
	CheckCode          string      `json:"check_code"`
	PriceFloatPercent  interface{} `json:"price_float_percent"`
	AddPositionPercent interface{} `json:"add_position_percent"`
	AddPositionNum     interface{} `json:"add_position_num"`
	InitMargin         interface{} `json:"init_margin"`
	RoundTpPercent     interface{} `json:"round_tp_percent"`
	SlPercent          interface{} `json:"sl_percent"`
	EntryPrice         interface{} `json:"entry_price"`
	Leverage           interface{} `json:"leverage"`
}

type GetFMartLimitRequest struct {
	Symbol             string `json:"symbol"`
	MartingaleMode     string `json:"martingale_mode"`
	Leverage           string `json:"leverage"`
	PriceFloatPercent  string `json:"price_float_percent"`
	AddPositionPercent string `json:"add_position_percent"`
	AddPositionNum     int    `json:"add_position_num"`
	InitMargin         string `json:"init_margin"`
	RoundTpPercent     string `json:"round_tp_percent"`
	SlPercent          string `json:"sl_percent"`
	EntryPrice         string `json:"entry_price"`
	NeedToSlippage     bool   `json:"need_to_slippage"`
	AppName            string `json:"app_name"`
}

type FMartLimitCheckCode struct {
	Value string `json:"value"`
}

type FMartMode struct {
	Value string `json:"value"`
}

type CreateFMartBotRequest struct {
	Symbol             string `json:"symbol"`
	MartingaleMode     string `json:"martingale_mode"`
	Leverage           string `json:"leverage"`
	PriceFloatPercent  string `json:"price_float_percent"`
	AddPositionPercent string `json:"add_position_percent"`
	AddPositionNum     int    `json:"add_position_num"`
	InitMargin         string `json:"init_margin"`
	RoundTpPercent     string `json:"round_tp_percent"`
	AutoCycleToggle    string `json:"auto_cycle_toggle"`
	SlPercent          string `json:"sl_percent"`
	EntryPrice         string `json:"entry_price"`
	Source             string `json:"source"`
	FollowedBotId      int    `json:"followed_bot_id"`
	BlockSource        string `json:"block_source"`
	CreateType         string `json:"create_type"`
	InitBonus          string `json:"init_bonus"`
	Channel            string `json:"channel"`
}
