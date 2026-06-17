package models

type GetFGridDetailResponse struct {
	StatusCode int         `json:"status_code"`
	Detail     interface{} `json:"detail"`
	DebugMsg   string      `json:"debug_msg"`
}

type FutureGridDetail struct {
	BotId                   int    `json:"bot_id"`
	Symbol                  string `json:"symbol"`
	Status                  string `json:"status"`
	BaseToken               string `json:"base_token"`
	QuoteToken              string `json:"quote_token"`
	MinPrice                string `json:"min_price"`
	MaxPrice                string `json:"max_price"`
	CellNumber              int    `json:"cell_number"`
	GridType                string `json:"grid_type"`
	GridMode                string `json:"grid_mode"`
	StopLossPer             string `json:"stop_loss_per"`
	TakeProfitPer           string `json:"take_profit_per"`
	RealLeverage            string `json:"real_leverage"`
	TotalInvestment         string `json:"total_investment"`
	TotalValue              string `json:"total_value"`
	CurrentPosition         string `json:"current_position"`
	ArbitrageNum            int    `json:"arbitrage_num"`
	ArbitrageNum24          int    `json:"arbitrage_num_24"`
	Pnl                     string `json:"pnl"`
	PnlPer                  string `json:"pnl_per"`
	GridProfit              string `json:"grid_profit"`
	GridApr                 string `json:"grid_apr"`
	TotalApr                string `json:"total_apr"`
	LastPrice               string `json:"last_price"`
	LiquidationPrice        string `json:"liquidation_price"`
	RealisedPnl             string `json:"realised_pnl"`
	UnrealisedPnl           string `json:"unrealised_pnl"`
	FundingFee              string `json:"funding_fee"`
	PositionBalance         string `json:"position_balance"`
	AvailableBalance        string `json:"available_balance"`
	TotalOrderBalance       string `json:"total_order_balance"`
	CloseReason             string `json:"close_reason"`
	AllowFollow             int    `json:"allow_follow"`
	OperationTime           int    `json:"operation_time"`
	CreateTime              int    `json:"create_time"`
	ModifyTime              int    `json:"modify_time"`
	EndTime                 int    `json:"end_time"`
	MinProfit               string `json:"min_profit"`
	MaxProfit               string `json:"max_profit"`
	Leverage                string `json:"leverage"`
	MarkPrice               string `json:"mark_price"`
	BotCloseCode            string `json:"bot_close_code"`
	FuturesPosSide          string `json:"futures_pos_side"`
	FollowNum               int    `json:"follow_num"`
	EntryPrice              string `json:"entry_price"`
	StopLossPrice           string `json:"stop_loss_price"`
	TakeProfitPrice         string `json:"take_profit_price"`
	TpSlType                string `json:"tp_sl_type"`
	RealClosePrice          string `json:"real_close_price"`
	MinPricePrecision       string `json:"min_price_precision"`
	TickSize                string `json:"tick_size"`
	UsedRewardAmount        string `json:"used_reward_amount"`
	SettlementAssets        string `json:"settlement_assets"`
	AccountType             string `json:"account_type"`
	CumWithdrewAmount       string `json:"cum_withdrew_amount"`
	CurrentProfit           string `json:"current_profit"`
	CurrentPer              string `json:"current_per"`
	CopyTradeIsMaster       bool   `json:"copy_trade_is_master"`
	CopyTradeFollowerNum    string `json:"copy_trade_follower_num"`
	InitBonus               string `json:"init_bonus"`
	UsedBonusAmount         string `json:"used_bonus_amount"`
	UsedRewardId            string `json:"used_reward_id"`
	TakerFeeRate            string `json:"taker_fee_rate"`
	MakerFeeRate            string `json:"maker_fee_rate"`
	Equity                  string `json:"equity"`
	TrailingStopExitEquity  string `json:"trailing_stop_exit_equity"`
	TrailingStopPer         string `json:"trailing_stop_per"`
	IMRate                  string `json:"i_m_rate"`
	MMRate                  string `json:"m_m_rate"`
	MoveUpPrice             string `json:"move_up_price"`
	MoveDownPrice           string `json:"move_down_price"`
	CurrMinPrice            string `json:"curr_min_price"`
	CurrMaxPrice            string `json:"curr_max_price"`
	AdlRankIndicator        int    `json:"adl_rank_indicator"`
}

type FGridBotSingleCreateResponse struct {
	StatusCode   int    `json:"status_code"`
	DebugMsg     string `json:"debug_msg"`
	BotId        int    `json:"bot_id"`
	CheckCode    string `json:"check_code"`
	BanReasonText string `json:"ban_reason_text"`
}
