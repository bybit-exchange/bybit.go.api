package models

type CloseGridResponse struct {
	StatusCode int    `json:"status_code"`
	DebugMsg   string `json:"debug_msg"`
}

type CloseGridRequest struct {
	GridId    int `json:"grid_id"`
	CloseMode int `json:"close_mode"`
}

type CreateGridBotResponse struct {
	StatusCode    int    `json:"status_code"`
	DebugMsg      string `json:"debug_msg"`
	GridId        int    `json:"grid_id"`
	BanReasonText string `json:"ban_reason_text"`
}

type CreateGridBotRequest struct {
	Symbol                  string      `json:"symbol"`
	MaxPrice                string      `json:"max_price"`
	MinPrice                string      `json:"min_price"`
	TotalInvestment         string      `json:"total_investment"`
	CellNumber              int         `json:"cell_number"`
	FollowedGridId          int         `json:"followed_grid_id"`
	Source                  int         `json:"source"`
	EntryPrice              string      `json:"entry_price"`
	StopLossPrice           string      `json:"stop_loss_price"`
	TakeProfitPrice         string      `json:"take_profit_price"`
	ToolsDiscoveryParameter interface{} `json:"toolsDiscoveryParameter"`
	BaseInvestment          string      `json:"base_investment"`
	QuoteInvestment         string      `json:"quote_investment"`
	InvestMode              int         `json:"invest_mode"`
	BlockSource             int         `json:"block_source"`
	CreateType              int         `json:"create_type"`
	TsPercent               string      `json:"ts_percent"`
	EnableTrailing          bool        `json:"enable_trailing"`
	LimitUpPrice            string      `json:"limit_up_price"`
	Channel                 string      `json:"channel"`
}

type QueryGridDetailResponse struct {
	StatusCode int         `json:"status_code"`
	Detail     interface{} `json:"detail"`
	DebugMsg   string      `json:"debug_msg"`
}

type GridDetail struct {
	GridId            int    `json:"grid_id"`
	Symbol            string `json:"symbol"`
	BaseToken         string `json:"base_token"`
	QuoteToken        string `json:"quote_token"`
	TotalInvestment   string `json:"total_investment"`
	TotalProfit       string `json:"total_profit"`
	MaxPrice          string `json:"max_price"`
	MinPrice          string `json:"min_price"`
	CellNumber        int    `json:"cell_number"`
	GridProfit        string `json:"grid_profit"`
	GridApr           string `json:"grid_apr"`
	TotalApr          string `json:"total_apr"`
	ArbitrageNum      int    `json:"arbitrage_num"`
	ArbitrageNum24    int    `json:"arbitrage_num_24"`
	AllowFollow       int    `json:"allow_follow"`
	OperationTime     int    `json:"operation_time"`
	Status            string `json:"status"`
	EntryPrice        string `json:"entry_price"`
	StopLossPrice     string `json:"stop_loss_price"`
	TakeProfitPrice   string `json:"take_profit_price"`
	CurrentPrice      string `json:"current_price"`
	CloseReason       string `json:"close_reason"`
	BotCloseCode      string `json:"bot_close_code"`
	FollowNum         int    `json:"follow_num"`
	CreateTime        int    `json:"create_time"`
	EndTime           int    `json:"end_time"`
	UsedRewardAmount  string `json:"used_reward_amount"`
	SettlementAssets  string `json:"settlement_assets"`
	AccountType       string `json:"account_type"`
	CumWithdrewAmount string `json:"cum_withdrew_amount"`
	CurrentProfit     string `json:"current_profit"`
	CurrentPer        string `json:"current_per"`
	Equity            string `json:"equity"`
	TsPercent         string `json:"ts_percent"`
	IsSupportTs       bool   `json:"is_support_ts"`
	EnableTrailing    bool   `json:"enable_trailing"`
	LimitUpPrice      string `json:"limit_up_price"`
	OriMaxPrice       string `json:"ori_max_price"`
	OriMinPrice       string `json:"ori_min_price"`
	OriCellNumber     int    `json:"ori_cell_number"`
}

type ValidateInputResponse struct {
	StatusCode     int         `json:"status_code"`
	DebugMsg       string      `json:"debug_msg"`
	Investment     interface{} `json:"investment"`
	Profit         interface{} `json:"profit"`
	CellNumber     interface{} `json:"cell_number"`
	MinPrice       interface{} `json:"min_price"`
	MaxPrice       interface{} `json:"max_price"`
	StopLoss       interface{} `json:"stop_loss"`
	TakeProfit     interface{} `json:"take_profit"`
	EntryPrice     interface{} `json:"entry_price"`
	CellDistance   string      `json:"cell_distance"`
	CheckCode      string      `json:"check_code"`
	BaseInvestment interface{} `json:"base_investment"`
	TsPercent      interface{} `json:"ts_percent"`
	LimitUpPrice   interface{} `json:"limit_up_price"`
}

type SpotGridRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ValidateInputRequest struct {
	Symbol          string `json:"symbol"`
	CellNumber      int    `json:"cell_number"`
	MinPrice        string `json:"min_price"`
	MaxPrice        string `json:"max_price"`
	TotalInvestment string `json:"total_investment"`
	StopLoss        string `json:"stop_loss"`
	TakeProfit      string `json:"take_profit"`
	EntryPrice      string `json:"entry_price"`
	BaseInvestment  string `json:"base_investment"`
	QuoteInvestment string `json:"quote_investment"`
	InvestMode      int    `json:"invest_mode"`
	TsPercent       string `json:"ts_percent"`
	EnableTrailing  bool   `json:"enable_trailing"`
	LimitUpPrice    string `json:"limit_up_price"`
}
