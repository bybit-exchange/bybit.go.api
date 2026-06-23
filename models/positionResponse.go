package models

type PositionListInfo struct {
	Category string `json:"category"`
	List     []struct {
		PositionIdx            int    `json:"positionIdx"`
		RiskId                 int    `json:"riskId"`
		RiskLimitValue         string `json:"riskLimitValue"`
		Symbol                 string `json:"symbol"`
		Side                   string `json:"side"`
		Size                   string `json:"size"`
		AvgPrice               string `json:"avgPrice"`
		PositionValue          string `json:"positionValue"`
		TradeMode              int    `json:"tradeMode"`
		AutoAddMargin          int    `json:"autoAddMargin"`
		PositionStatus         string `json:"positionStatus"`
		Leverage               string `json:"leverage"`
		MarkPrice              string `json:"markPrice"`
		LiqPrice               string `json:"liqPrice"`
		BustPrice              string `json:"bustPrice"`
		PositionIM             string `json:"positionIM"`
		PositionMM             string `json:"positionMM"`
		PositionBalance        string `json:"positionBalance"`
		TpslMode               string `json:"tpslMode"`
		TakeProfit             string `json:"takeProfit"`
		StopLoss               string `json:"stopLoss"`
		TrailingStop           string `json:"trailingStop"`
		UnrealisedPnl          string `json:"unrealisedPnl"`
		CumRealisedPnl         string `json:"cumRealisedPnl"`
		AdlRankIndicator       int    `json:"adlRankIndicator"`
		IsReduceOnly           bool   `json:"isReduceOnly"`
		MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
		LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
		CreatedTime            string `json:"createdTime"`
		UpdatedTime            string `json:"updatedTime"`
		Seq                    int64  `json:"seq"`
	} `json:"list"`
	NextPageCursor string `json:"nextPageCursor"`
}

type PositionTpslMode struct {
	TpSlMode string `json:"tpSlMode"`
}

type PositionRiskInfo struct {
	Category       string `json:"category"`
	RiskId         int64  `json:"riskId"`
	RiskLimitValue string `json:"riskLimitValue"`
}

type PositionUpdateMargin struct {
	Category       string `json:"category"`
	Symbol         string `json:"symbol"`
	PositionIdx    int    `json:"positionIdx"`
	RiskId         int    `json:"riskId"`
	RiskLimitValue string `json:"riskLimitValue"`
	Size           string `json:"size"`
	AvgPrice       string `json:"avgPrice"`
	LiqPrice       string `json:"liqPrice"`
	BustPrice      string `json:"bustPrice"`
	MarkPrice      string `json:"markPrice"`
	PositionValue  string `json:"positionValue"`
	Leverage       string `json:"leverage"`
	AutoAddMargin  int    `json:"autoAddMargin"`
	PositionStatus string `json:"positionStatus"`
	PositionIM     string `json:"positionIM"`
	PositionMM     string `json:"positionMM"`
	TakeProfit     string `json:"takeProfit"`
	StopLoss       string `json:"stopLoss"`
	TrailingStop   string `json:"trailingStop"`
	UnrealisedPnl  string `json:"unrealisedPnl"`
	CumRealisedPnl string `json:"cumRealisedPnl"`
	CreatedTime    string `json:"createdTime"`
	UpdatedTime    string `json:"updatedTime"`
}

type PositionExecutionInfo struct {
	Category string                   `json:"category"`
	List     []PositionExecutionEntry `json:"list"`
}

type PositionExecutionEntry struct {
	Symbol          string `json:"symbol"`
	OrderID         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	Side            string `json:"side"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	ExecFee         string `json:"execFee"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	IsMaker         bool   `json:"isMaker"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	BlockTradeId    string `json:"blockTradeId"`
	ClosedSize      string `json:"closedSize"`
	Seq             int64  `json:"seq"`
	NextPageCursor  string `json:"nextPageCursor"`
}

type PositionClosedPnlInfo struct {
	Category       string               `json:"category"`
	List           []ClosedPnlInfoEntry `json:"list"`
	NextPageCursor string               `json:"nextPageCursor"`
}

type ClosedPnlInfoEntry struct {
	Symbol        string `json:"symbol"`
	OrderID       string `json:"orderId"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	OrderPrice    string `json:"orderPrice"`
	OrderType     string `json:"orderType"`
	ExecType      string `json:"execType"`
	ClosedSize    string `json:"closedSize"`
	CumEntryValue string `json:"cumEntryValue"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	CumExitValue  string `json:"cumExitValue"`
	AvgExitPrice  string `json:"avgExitPrice"`
	ClosedPnl     string `json:"closedPnl"`
	FillCount     string `json:"fillCount"`
	Leverage      string `json:"leverage"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}

type AddReduceMarginResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     AddReduceMarginResult  `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type AddReduceMarginResult struct {
	Category       string `json:"category"`
	Symbol         string `json:"symbol"`
	PositionIdx    int    `json:"positionIdx"`
	RiskId         int    `json:"riskId"`
	RiskLimitValue string `json:"riskLimitValue"`
	Size           string `json:"size"`
	AvgPrice       string `json:"avgPrice"`
	LiqPrice       string `json:"liqPrice"`
	BustPrice      string `json:"bustPrice"`
	MarkPrice      string `json:"markPrice"`
	PositionValue  string `json:"positionValue"`
	Leverage       string `json:"leverage"`
	AutoAddMargin  int    `json:"autoAddMargin"`
	PositionStatus string `json:"positionStatus"`
	PositionIM     string `json:"positionIM"`
	PositionMM     string `json:"positionMM"`
	TakeProfit     string `json:"takeProfit"`
	StopLoss       string `json:"stopLoss"`
	TrailingStop   string `json:"trailingStop"`
	UnrealisedPnl  string `json:"unrealisedPnl"`
	CumRealisedPnl string `json:"cumRealisedPnl"`
	CreatedTime    string `json:"createdTime"`
	UpdatedTime    string `json:"updatedTime"`
}

type GetClosePositionResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     GetClosePositionResult `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetClosePositionResult struct {
	Category       string                 `json:"category"`
	List           []GetClosePositionItem `json:"list"`
	NextPageCursor string                 `json:"nextPageCursor"`
}

type GetClosePositionItem struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	AvgExitPrice  string `json:"avgExitPrice"`
	DeliveryPrice string `json:"deliveryPrice"`
	TotalOpenFee  string `json:"totalOpenFee"`
	TotalCloseFee string `json:"totalCloseFee"`
	DeliveryFee   string `json:"deliveryFee"`
	TotalPnl      string `json:"totalPnl"`
	OpenTime      int64  `json:"openTime"`
	CloseTime     int64  `json:"closeTime"`
}

type GetClosedPnlResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     GetClosedPnlResult     `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetClosedPnlResult struct {
	Category       string          `json:"category"`
	List           []ClosedPnlItem `json:"list"`
	NextPageCursor string          `json:"nextPageCursor"`
}

type ClosedPnlItem struct {
	Symbol        string `json:"symbol"`
	OrderId       string `json:"orderId"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	OrderPrice    string `json:"orderPrice"`
	OrderType     string `json:"orderType"`
	ExecType      string `json:"execType"`
	ClosedSize    string `json:"closedSize"`
	CumEntryValue string `json:"cumEntryValue"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	CumExitValue  string `json:"cumExitValue"`
	AvgExitPrice  string `json:"avgExitPrice"`
	ClosedPnl     string `json:"closedPnl"`
	FillCount     string `json:"fillCount"`
	Leverage      string `json:"leverage"`
	OpenFee       string `json:"openFee"`
	CloseFee      string `json:"closeFee"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}

type GetMovePositionHistoryResponse struct {
	RetCode    int                          `json:"retCode"`
	RetMsg     string                       `json:"retMsg"`
	Result     GetMovePositionHistoryResult `json:"result"`
	RetExtInfo map[string]interface{}       `json:"retExtInfo"`
	Time       int64                        `json:"time"`
}

type GetMovePositionHistoryResult struct {
	List           []GetMovePositionHistoryItem `json:"list"`
	NextPageCursor string                       `json:"nextPageCursor"`
}

type GetMovePositionHistoryItem struct {
	BlockTradeId  string `json:"blockTradeId"`
	Category      string `json:"category"`
	OrderId       string `json:"orderId"`
	UserId        int    `json:"userId"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Qty           string `json:"qty"`
	ExecFee       string `json:"execFee"`
	Status        string `json:"status"`
	ExecId        string `json:"execId"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
	RejectParty   string `json:"rejectParty"`
}

type GetPositionInfoResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     PositionInfoResult     `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type PositionInfoResult struct {
	Category       string         `json:"category"`
	List           []PositionItem `json:"list"`
	NextPageCursor string         `json:"nextPageCursor"`
}

type PositionItem struct {
	PositionIdx            int    `json:"positionIdx"`
	RiskId                 int    `json:"riskId"`
	RiskLimitValue         string `json:"riskLimitValue"`
	Symbol                 string `json:"symbol"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	AvgPrice               string `json:"avgPrice"`
	PositionValue          string `json:"positionValue"`
	TradeMode              int    `json:"tradeMode"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	PositionStatus         string `json:"positionStatus"`
	Leverage               string `json:"leverage"`
	MarkPrice              string `json:"markPrice"`
	LiqPrice               string `json:"liqPrice"`
	BustPrice              string `json:"bustPrice"`
	PositionIM             string `json:"positionIM"`
	PositionMM             string `json:"positionMM"`
	PositionBalance        string `json:"positionBalance"`
	TpslMode               string `json:"tpslMode"`
	TakeProfit             string `json:"takeProfit"`
	StopLoss               string `json:"stopLoss"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	CurRealisedPnl         string `json:"curRealisedPnl"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	PositionIMByMp         string `json:"positionIMByMp"`
	PositionMMByMp         string `json:"positionMMByMp"`
	SessionAvgPrice        string `json:"sessionAvgPrice"`
	Delta                  string `json:"delta"`
	Gamma                  string `json:"gamma"`
	Vega                   string `json:"vega"`
	Theta                  string `json:"theta"`
	Seq                    int64  `json:"seq"`
	CreatedTime            string `json:"createdTime"`
	UpdatedTime            string `json:"updatedTime"`
}

type SetLeverageResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type SetTradingStopResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetClosedPositionHistoryResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type ClosedPositionHistoryItem struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	AvgExitPrice  string `json:"avgExitPrice"`
	TotalPnl      string `json:"totalPnl"`
	DeliveryPrice string `json:"deliveryPrice"`
	TotalOpenFee  string `json:"totalOpenFee"`
	TotalCloseFee string `json:"totalCloseFee"`
	DeliveryFee   string `json:"deliveryFee"`
	OpenTime      int64  `json:"openTime"`
	CloseTime     int64  `json:"closeTime"`
}

type GetPositionListResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type PositionListItem struct {
	PositionIdx            int    `json:"positionIdx"`
	RiskId                 int    `json:"riskId"`
	RiskLimitValue         string `json:"riskLimitValue"`
	Symbol                 string `json:"symbol"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	AvgPrice               string `json:"avgPrice"`
	PositionValue          string `json:"positionValue"`
	TradeMode              int    `json:"tradeMode"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	PositionStatus         string `json:"positionStatus"`
	Leverage               string `json:"leverage"`
	MarkPrice              string `json:"markPrice"`
	LiqPrice               string `json:"liqPrice"`
	BustPrice              string `json:"bustPrice"`
	PositionIM             string `json:"positionIM"`
	PositionMM             string `json:"positionMM"`
	PositionBalance        string `json:"positionBalance"`
	TpslMode               string `json:"tpslMode"`
	TakeProfit             string `json:"takeProfit"`
	StopLoss               string `json:"stopLoss"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	CurRealisedPnl         string `json:"curRealisedPnl"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	SessionAvgPrice        string `json:"sessionAvgPrice"`
	Delta                  string `json:"delta"`
	Gamma                  string `json:"gamma"`
	Vega                   string `json:"vega"`
	Theta                  string `json:"theta"`
	CreatedTime            string `json:"createdTime"`
	UpdatedTime            string `json:"updatedTime"`
	Seq                    int64  `json:"seq"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	PositionIMByMp         string `json:"positionIMByMp"`
	PositionMMByMp         string `json:"positionMMByMp"`
	OpenTime               int64  `json:"openTime"`
}

type GetPositionSymbolInfoResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type PositionSymbolInfoItem struct {
	Symbol      string `json:"symbol"`
	Leverage    string `json:"leverage"`
	Side        string `json:"side"`
	PositionIdx int    `json:"positionIdx"`
}

type ConfirmNewRiskLimitResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type MovePositionResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     MovePositionResult     `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type MovePositionResult struct {
	BlockTradeId string `json:"blockTradeId"`
	Status       string `json:"status"`
	RejectParty  string `json:"rejectParty"`
}

type MovePositionLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Side     string `json:"side"`
	Qty      string `json:"qty"`
}

type SetAutoAddMarginResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type ClosedPnlResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     ClosedPnlResult        `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type ClosedPnlResult struct {
	Category       string          `json:"category"`
	List           []ClosedPnlItem `json:"list"`
	NextPageCursor string          `json:"nextPageCursor"`
}

type MovePositionHistoryResponse struct {
	RetCode    int                       `json:"retCode"`
	RetMsg     string                    `json:"retMsg"`
	Result     MovePositionHistoryResult `json:"result"`
	RetExtInfo map[string]interface{}    `json:"retExtInfo"`
	Time       int64                     `json:"time"`
}

type MovePositionHistoryResult struct {
	List           []MovePositionHistoryItem `json:"list"`
	NextPageCursor string                    `json:"nextPageCursor"`
}

type MovePositionHistoryItem struct {
	BlockTradeId  string `json:"blockTradeId"`
	Category      string `json:"category"`
	OrderId       string `json:"orderId"`
	UserId        int    `json:"userId"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Qty           string `json:"qty"`
	ExecFee       string `json:"execFee"`
	Status        string `json:"status"`
	ExecId        string `json:"execId"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
	RejectParty   string `json:"rejectParty"`
}

type GetPositionInfoResult struct {
	Category       string         `json:"category"`
	List           []PositionInfo `json:"list"`
	NextPageCursor string         `json:"nextPageCursor"`
}

type PositionInfo struct {
	PositionIdx            int    `json:"positionIdx"`
	RiskId                 int    `json:"riskId"`
	RiskLimitValue         string `json:"riskLimitValue"`
	Symbol                 string `json:"symbol"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	AvgPrice               string `json:"avgPrice"`
	PositionValue          string `json:"positionValue"`
	TradeMode              int    `json:"tradeMode"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	PositionStatus         string `json:"positionStatus"`
	Leverage               string `json:"leverage"`
	MarkPrice              string `json:"markPrice"`
	LiqPrice               string `json:"liqPrice"`
	BustPrice              string `json:"bustPrice"`
	PositionIM             string `json:"positionIM"`
	PositionMM             string `json:"positionMM"`
	PositionBalance        string `json:"positionBalance"`
	TpslMode               string `json:"tpslMode"`
	TakeProfit             string `json:"takeProfit"`
	StopLoss               string `json:"stopLoss"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	CurRealisedPnl         string `json:"curRealisedPnl"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	PositionIMByMp         string `json:"positionIMByMp"`
	PositionMMByMp         string `json:"positionMMByMp"`
	SessionAvgPrice        string `json:"sessionAvgPrice"`
	Delta                  string `json:"delta"`
	Gamma                  string `json:"gamma"`
	Vega                   string `json:"vega"`
	Theta                  string `json:"theta"`
	Seq                    int64  `json:"seq"`
	CreatedTime            string `json:"createdTime"`
	UpdatedTime            string `json:"updatedTime"`
}

type MovePositionRequestLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Side     string `json:"side"`
	Qty      string `json:"qty"`
}

type SwitchPositionModeResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}
