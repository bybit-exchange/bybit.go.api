package models

type ValidateFGridInputResponse struct {
	StatusCode      int         `json:"status_code"`
	DebugMsg        string      `json:"debug_msg"`
	Investment      interface{} `json:"investment"`
	Profit          interface{} `json:"profit"`
	CellNumber      interface{} `json:"cell_number"`
	MinPrice        interface{} `json:"min_price"`
	MaxPrice        interface{} `json:"max_price"`
	Leverage        interface{} `json:"leverage"`
	StopLoss        interface{} `json:"stop_loss"`
	TakeProfit      interface{} `json:"take_profit"`
	CheckCode       string      `json:"check_code"`
	TakeProfitPrice interface{} `json:"take_profit_price"`
	StopLossPrice   interface{} `json:"stop_loss_price"`
	EntryPrice      interface{} `json:"entry_price"`
	CellGtv0Per     string      `json:"cell_gtv0_per"`
	TrailingStopPer interface{} `json:"trailing_stop_per"`
	LongLiqPrice    string      `json:"long_liq_price"`
	ShortLiqPrice   string      `json:"short_liq_price"`
	MoveUpPrice     interface{} `json:"move_up_price"`
	MoveDownPrice   interface{} `json:"move_down_price"`
}

type GridRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}
