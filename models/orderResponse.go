package models

type OrderResult struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type ListOrderResult struct {
	List []OrderResult `json:"list"`
}

type OrderInfo struct {
	OrderId            string `json:"orderId"`
	OrderLinkId        string `json:"orderLinkId"`
	BlockTradeId       string `json:"blockTradeId"`
	Symbol             string `json:"symbol"`
	Price              string `json:"price"`
	Qty                string `json:"qty"`
	Side               string `json:"side"`
	IsLeverage         string `json:"isLeverage"`
	PositionIdx        int    `json:"positionIdx"`
	OrderStatus        string `json:"orderStatus"`
	CancelType         string `json:"cancelType"`
	RejectReason       string `json:"rejectReason"`
	AvgPrice           string `json:"avgPrice"`
	LeavesQty          string `json:"leavesQty"`
	LeavesValue        string `json:"leavesValue"`
	CumExecQty         string `json:"cumExecQty"`
	CumExecValue       string `json:"cumExecValue"`
	CumExecFee         string `json:"cumExecFee"`
	TimeInForce        string `json:"timeInForce"`
	OrderType          string `json:"orderType"`
	StopOrderType      string `json:"stopOrderType"`
	OrderIv            string `json:"orderIv"`
	TriggerPrice       string `json:"triggerPrice"`
	TakeProfit         string `json:"takeProfit"`
	StopLoss           string `json:"stopLoss"`
	TpslMode           string `json:"tpslMode"`
	OcoTriggerType     string `json:"ocoTriggerType"`
	TpLimitPrice       string `json:"tpLimitPrice"`
	SlLimitPrice       string `json:"slLimitPrice"`
	TpTriggerBy        string `json:"tpTriggerBy"`
	SlTriggerBy        string `json:"slTriggerBy"`
	TriggerDirection   int    `json:"triggerDirection"`
	TriggerBy          string `json:"triggerBy"`
	LastPriceOnCreated string `json:"lastPriceOnCreated"`
	ReduceOnly         bool   `json:"reduceOnly"`
	CloseOnTrigger     bool   `json:"closeOnTrigger"`
	PlaceType          string `json:"placeType"`
	SmpType            string `json:"smpType"`
	SmpGroup           int    `json:"smpGroup"`
	SmpOrderId         string `json:"smpOrderId"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
}

type OpenOrdersInfo struct {
	Category       string      `json:"category"`
	NextPageCursor string      `json:"nextPageCursor"`
	List           []OrderInfo `json:"list"`
}

type BorrowQuotaInfo struct {
	Symbol             string `json:"symbol"`
	Side               string `json:"side"`
	MaxTradeQty        string `json:"maxTradeQty"`
	MaxTradeAmount     string `json:"maxTradeAmount"`
	SpotMaxTradeQty    string `json:"spotMaxTradeQty"`
	SpotMaxTradeAmount string `json:"spotMaxTradeAmount"`
	BorrowCoin         string `json:"borrowCoin"`
}

type BatchOrderServerResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Category    string  `json:"category"`
			Symbol      string  `json:"symbol"`
			OrderId     string  `json:"orderId"`
			OrderLinkId string  `json:"orderLinkId"`
			CreateAt    *string `json:"createAt,omitempty"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
		List []struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"list"`
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type CancelQuoteResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetQuotesRealtimeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetQuotesRealtimeResult struct {
	List []interface{} `json:"list"`
}

type QuoteRealtimeItem struct {
	RfqId         string        `json:"rfqId"`
	RfqLinkId     string        `json:"rfqLinkId"`
	QuoteId       string        `json:"quoteId"`
	QuoteLinkId   string        `json:"quoteLinkId"`
	ExpiresAt     string        `json:"expiresAt"`
	Status        string        `json:"status"`
	DeskCode      string        `json:"deskCode"`
	ExecQuoteSide string        `json:"execQuoteSide"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	QuoteBuyList  []interface{} `json:"quoteBuyList"`
	QuoteSellList []interface{} `json:"quoteSellList"`
}

type QuoteLegItem struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Qty      string `json:"qty"`
}

type GetQuotesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetQuotesResult struct {
	Cursor string        `json:"cursor"`
	List   []interface{} `json:"list"`
}

type QuoteItem struct {
	RfqId         string        `json:"rfqId"`
	RfqLinkId     string        `json:"rfqLinkId"`
	QuoteId       string        `json:"quoteId"`
	QuoteLinkId   string        `json:"quoteLinkId"`
	ExpiresAt     string        `json:"expiresAt"`
	DeskCode      string        `json:"deskCode"`
	Status        string        `json:"status"`
	ExecQuoteSide string        `json:"execQuoteSide"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	QuoteBuyList  []interface{} `json:"quoteBuyList"`
	QuoteSellList []interface{} `json:"quoteSellList"`
}

type QuoteLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Qty      string `json:"qty"`
}

type GetRfqsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetRfqsResult struct {
	Cursor string        `json:"cursor"`
	List   []interface{} `json:"list"`
}

type RfqItem struct {
	RfqId                  string        `json:"rfqId"`
	RfqLinkId              string        `json:"rfqLinkId"`
	Counterparties         []string      `json:"counterparties"`
	StrategyType           string        `json:"strategyType"`
	ExpiresAt              string        `json:"expiresAt"`
	Status                 string        `json:"status"`
	AcceptOtherQuoteStatus string        `json:"acceptOtherQuoteStatus"`
	DeskCode               string        `json:"deskCode"`
	CreatedAt              string        `json:"createdAt"`
	UpdatedAt              string        `json:"updatedAt"`
	Legs                   []interface{} `json:"legs"`
}

type RfqLegItem struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	Qty      string `json:"qty"`
}

type AmendOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type BatchAmendOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type BatchAmendOrdersRequest struct {
	Category string                `json:"category"`
	Request  []BatchAmendOrderItem `json:"request"`
}

type BatchAmendOrderItem struct {
	Symbol       string `json:"symbol"`
	OrderId      string `json:"orderId,omitempty"`
	OrderLinkId  string `json:"orderLinkId,omitempty"`
	Qty          string `json:"qty,omitempty"`
	Price        string `json:"price,omitempty"`
	OrderIv      string `json:"orderIv,omitempty"`
	TriggerPrice string `json:"triggerPrice,omitempty"`
	TpslMode     string `json:"tpslMode,omitempty"`
	TakeProfit   string `json:"takeProfit,omitempty"`
	StopLoss     string `json:"stopLoss,omitempty"`
	TpTriggerBy  string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy  string `json:"slTriggerBy,omitempty"`
	TriggerBy    string `json:"triggerBy,omitempty"`
	TpLimitPrice string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice string `json:"slLimitPrice,omitempty"`
}

type BatchCancelOrdersRequest struct {
	Category string                 `json:"category"`
	Request  []BatchCancelOrderItem `json:"request"`
}

type BatchCancelOrderItem struct {
	Symbol      string `json:"symbol"`
	OrderId     string `json:"orderId,omitempty"`
	OrderLinkId string `json:"orderLinkId,omitempty"`
}

type BatchCancelOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type BatchCreateOrdersRequest struct {
	Category string           `json:"category"`
	Request  []BatchOrderItem `json:"request"`
}

type BatchOrderItem struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	OrderType        string `json:"orderType"`
	Qty              string `json:"qty"`
	Price            string `json:"price,omitempty"`
	IsLeverage       int    `json:"isLeverage,omitempty"`
	MarketUnit       string `json:"marketUnit,omitempty"`
	TimeInForce      string `json:"timeInForce,omitempty"`
	TriggerDirection int    `json:"triggerDirection,omitempty"`
	OrderFilter      string `json:"orderFilter,omitempty"`
	TriggerPrice     string `json:"triggerPrice,omitempty"`
	TriggerBy        string `json:"triggerBy,omitempty"`
	OrderIv          string `json:"orderIv,omitempty"`
	PositionIdx      int    `json:"positionIdx,omitempty"`
	OrderLinkId      string `json:"orderLinkId,omitempty"`
	TakeProfit       string `json:"takeProfit,omitempty"`
	StopLoss         string `json:"stopLoss,omitempty"`
	TpTriggerBy      string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy      string `json:"slTriggerBy,omitempty"`
	TpslMode         string `json:"tpslMode,omitempty"`
	TpOrderType      string `json:"tpOrderType,omitempty"`
	SlOrderType      string `json:"slOrderType,omitempty"`
	TpLimitPrice     string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice     string `json:"slLimitPrice,omitempty"`
	ReduceOnly       bool   `json:"reduceOnly,omitempty"`
	CloseOnTrigger   bool   `json:"closeOnTrigger,omitempty"`
	SmpType          string `json:"smpType,omitempty"`
	Mmp              bool   `json:"mmp,omitempty"`
}

type BatchCreateOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type CancelAllOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type CancelOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type CreateOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetOpenOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetOrderHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type OrderHistoryDetail struct {
	OrderId               string      `json:"orderId"`
	OrderLinkId           string      `json:"orderLinkId"`
	ParentOrderLinkId     string      `json:"parentOrderLinkId"`
	BlockTradeId          string      `json:"blockTradeId"`
	Symbol                string      `json:"symbol"`
	IsLeverage            string      `json:"isLeverage"`
	PositionIdx           int         `json:"positionIdx"`
	Price                 string      `json:"price"`
	Qty                   string      `json:"qty"`
	Side                  string      `json:"side"`
	OrderStatus           string      `json:"orderStatus"`
	CreateType            string      `json:"createType"`
	CancelType            string      `json:"cancelType"`
	RejectReason          string      `json:"rejectReason"`
	AvgPrice              string      `json:"avgPrice"`
	CumExecQty            string      `json:"cumExecQty"`
	CumExecValue          string      `json:"cumExecValue"`
	CumExecFee            string      `json:"cumExecFee"`
	LeavesQty             string      `json:"leavesQty"`
	LeavesValue           string      `json:"leavesValue"`
	OrderType             string      `json:"orderType"`
	StopOrderType         string      `json:"stopOrderType"`
	OrderIv               string      `json:"orderIv"`
	MarketUnit            string      `json:"marketUnit"`
	TimeInForce           string      `json:"timeInForce"`
	TriggerPrice          string      `json:"triggerPrice"`
	TakeProfit            string      `json:"takeProfit"`
	StopLoss              string      `json:"stopLoss"`
	TpslMode              string      `json:"tpslMode"`
	TpLimitPrice          string      `json:"tpLimitPrice"`
	SlLimitPrice          string      `json:"slLimitPrice"`
	TpTriggerBy           string      `json:"tpTriggerBy"`
	SlTriggerBy           string      `json:"slTriggerBy"`
	TriggerDirection      int         `json:"triggerDirection"`
	TriggerBy             string      `json:"triggerBy"`
	LastPriceOnCreated    string      `json:"lastPriceOnCreated"`
	ReduceOnly            bool        `json:"reduceOnly"`
	CloseOnTrigger        bool        `json:"closeOnTrigger"`
	PlaceType             string      `json:"placeType"`
	SmpType               string      `json:"smpType"`
	SmpGroup              int         `json:"smpGroup"`
	SmpOrderId            string      `json:"smpOrderId"`
	RpiTakerAccess        bool        `json:"rpiTakerAccess"`
	RpiMatchedQty         string      `json:"rpiMatchedQty"`
	OcoTriggerBy          string      `json:"ocoTriggerBy"`
	BasePrice             string      `json:"basePrice"`
	SlippageToleranceType string      `json:"slippageToleranceType"`
	SlippageTolerance     string      `json:"slippageTolerance"`
	CumFeeDetail          interface{} `json:"cumFeeDetail"`
	ExtraFees             string      `json:"extraFees"`
	CreatedTime           string      `json:"createdTime"`
	UpdatedTime           string      `json:"updatedTime"`
}

type GetSpotBorrowQuotaResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetTradeHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeExecutionDetail struct {
	Symbol          string `json:"symbol"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	Side            string `json:"side"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	CreateType      string `json:"createType"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	ExecFee         string `json:"execFee"`
	ExecFeeV2       string `json:"execFeeV2"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	FeeCurrency     string `json:"feeCurrency"`
	IsMaker         bool   `json:"isMaker"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	BlockTradeId    string `json:"blockTradeId"`
	ClosedSize      string `json:"closedSize"`
	Seq             int    `json:"seq"`
	ExtraFees       string `json:"extraFees"`
}

type PreCheckOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type SetDcpResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PreUpgradeTransactionLogResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PreUpgradeTransactionLogItem struct {
	Symbol          string `json:"symbol"`
	Category        string `json:"category"`
	Side            string `json:"side"`
	TransactionTime string `json:"transactionTime"`
	Type            string `json:"type"`
	Qty             string `json:"qty"`
	Size            string `json:"size"`
	Currency        string `json:"currency"`
	TradePrice      string `json:"tradePrice"`
	Funding         string `json:"funding"`
	Fee             string `json:"fee"`
	CashFlow        string `json:"cashFlow"`
	Change          string `json:"change"`
	CashBalance     string `json:"cashBalance"`
	FeeRate         string `json:"feeRate"`
	BonusChange     string `json:"bonusChange"`
	TradeId         string `json:"tradeId"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
}

type PreUpgradeDeliveryRecordResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PreUpgradeDeliveryRecordItem struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	DeliveryTime  int    `json:"deliveryTime"`
	Strike        string `json:"strike"`
	Fee           string `json:"fee"`
	Position      string `json:"position"`
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryRpl   string `json:"deliveryRpl"`
}

type PreUpgradeSettlementRecordResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PreUpgradeSettlementRecordItem struct {
	Symbol          string `json:"symbol"`
	Side            string `json:"side"`
	Size            string `json:"size"`
	SessionAvgPrice string `json:"sessionAvgPrice"`
	MarkPrice       string `json:"markPrice"`
	RealisedPnl     string `json:"realisedPnl"`
	CreatedTime     string `json:"createdTime"`
}

type GetPreUpgradeClosedPnlResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type PreUpgradeClosedPnlDetail struct {
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
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}

type PreUpgradeExecutionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PreUpgradeExecutionItem struct {
	Symbol          string `json:"symbol"`
	OrderId         string `json:"orderId"`
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
}

type PreUpgradeOrderHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PreUpgradeOrderHistoryItem struct {
	OrderId            string `json:"orderId"`
	OrderLinkId        string `json:"orderLinkId"`
	BlockTradeId       string `json:"blockTradeId"`
	Symbol             string `json:"symbol"`
	Price              string `json:"price"`
	Qty                string `json:"qty"`
	Side               string `json:"side"`
	IsLeverage         string `json:"isLeverage"`
	PositionIdx        int    `json:"positionIdx"`
	OrderStatus        string `json:"orderStatus"`
	CancelType         string `json:"cancelType"`
	RejectReason       string `json:"rejectReason"`
	AvgPrice           string `json:"avgPrice"`
	LeavesQty          string `json:"leavesQty"`
	LeavesValue        string `json:"leavesValue"`
	CumExecQty         string `json:"cumExecQty"`
	CumExecValue       string `json:"cumExecValue"`
	CumExecFee         string `json:"cumExecFee"`
	TimeInForce        string `json:"timeInForce"`
	OrderType          string `json:"orderType"`
	StopOrderType      string `json:"stopOrderType"`
	OrderIv            string `json:"orderIv"`
	TriggerPrice       string `json:"triggerPrice"`
	TakeProfit         string `json:"takeProfit"`
	StopLoss           string `json:"stopLoss"`
	TpTriggerBy        string `json:"tpTriggerBy"`
	SlTriggerBy        string `json:"slTriggerBy"`
	TriggerDirection   int    `json:"triggerDirection"`
	TriggerBy          string `json:"triggerBy"`
	LastPriceOnCreated string `json:"lastPriceOnCreated"`
	ReduceOnly         bool   `json:"reduceOnly"`
	CloseOnTrigger     bool   `json:"closeOnTrigger"`
	PlaceType          string `json:"placeType"`
	SmpType            string `json:"smpType"`
	SmpGroup           int    `json:"smpGroup"`
	SmpOrderId         string `json:"smpOrderId"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
}

type TradeAssetDetailResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeBizTokenDetailsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeBizTokenListResponse struct {
	RetCode    int                `json:"retCode"`
	RetMsg     string             `json:"retMsg"`
	Result     []TradeBizTokenDto `json:"result"`
	RetExtInfo interface{}        `json:"retExtInfo"`
	Time       int64              `json:"time"`
}

type TradeBizTokenDto struct {
	TokenCode         string   `json:"tokenCode"`
	ChainCode         string   `json:"chainCode"`
	ChainIconUrl      string   `json:"chainIconUrl"`
	TokenAddress      string   `json:"tokenAddress"`
	Symbol            string   `json:"symbol"`
	TokenDecimals     int      `json:"tokenDecimals"`
	TokenIconUrlDay   string   `json:"tokenIconUrlDay"`
	TokenIconUrlNight string   `json:"tokenIconUrlNight"`
	CreateTime        int64    `json:"createTime"`
	CreateTimeOnchain int64    `json:"createTimeOnchain"`
	RiskFlag          int      `json:"riskFlag"`
	MinOrderQuantity  float64  `json:"minOrderQuantity"`
	MaxOrderQuantity  float64  `json:"maxOrderQuantity"`
	TokenTags         []int    `json:"tokenTags"`
	PayTokenCodes     []string `json:"payTokenCodes"`
}

type TradeBizTokenPriceListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeTokenPriceDto struct {
	ChainCode    string `json:"chainCode"`
	TokenAddress string `json:"tokenAddress"`
	Price        string `json:"price"`
	Change24h    string `json:"change24h"`
	Vol24h       string `json:"vol24h"`
	MarketCap    string `json:"marketCap"`
	Liquidity    string `json:"liquidity"`
	Holders      string `json:"holders"`
}

type TradeOrderListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeOrderDto struct {
	OrderType             int    `json:"orderType"`
	TradeType             int    `json:"tradeType"`
	OrderNo               string `json:"orderNo"`
	OrderStatus           int    `json:"orderStatus"`
	FromTokenCode         string `json:"fromTokenCode"`
	FromTokenAmount       string `json:"fromTokenAmount"`
	FromTokenSymbol       string `json:"fromTokenSymbol"`
	FromTokenDecimals     int    `json:"fromTokenDecimals"`
	FromTokenIconUrlDay   string `json:"fromTokenIconUrlDay"`
	FromTokenIconUrlNight string `json:"fromTokenIconUrlNight"`
	FromChainCode         string `json:"fromChainCode"`
	FromChainIconUrl      string `json:"fromChainIconUrl"`
	ToTokenCode           string `json:"toTokenCode"`
	ToTokenAmount         string `json:"toTokenAmount"`
	ToTokenSymbol         string `json:"toTokenSymbol"`
	ToTokenDecimals       int    `json:"toTokenDecimals"`
	ToTokenIconUrlDay     string `json:"toTokenIconUrlDay"`
	ToTokenIconUrlNight   string `json:"toTokenIconUrlNight"`
	ToChainCode           string `json:"toChainCode"`
	ToChainIconUrl        string `json:"toChainIconUrl"`
	GasTokenSymbol        string `json:"gasTokenSymbol"`
	GasOnchain            string `json:"gasOnchain"`
	GasUsd                string `json:"gasUsd"`
	PlatformFee           string `json:"platformFee"`
	PlatformFeeUsd        string `json:"platformFeeUsd"`
	QuoteMode             int    `json:"quoteMode"`
	CreateTime            int64  `json:"createTime"`
	ExecutionTime         int64  `json:"executionTime"`
	FailureReasonCode     string `json:"failureReasonCode"`
	Source                string `json:"source"`
	SwapRate              string `json:"swapRate"`
	ActualFromTokenAmount string `json:"actualFromTokenAmount"`
}

type TradePayTokenListResponse struct {
	RetCode    int                `json:"retCode"`
	RetMsg     string             `json:"retMsg"`
	Result     []TradePayTokenDto `json:"result"`
	RetExtInfo interface{}        `json:"retExtInfo"`
	Time       int64              `json:"time"`
}

type TradePayTokenDto struct {
	TokenCode         string   `json:"tokenCode"`
	Symbol            string   `json:"symbol"`
	TokenDecimals     int      `json:"tokenDecimals"`
	TokenIconUrlDay   string   `json:"tokenIconUrlDay"`
	TokenIconUrlNight string   `json:"tokenIconUrlNight"`
	Limit             string   `json:"limit"`
	SupportChains     []string `json:"supportChains"`
}

type TradeExecutionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeQuoteResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeQuoteRequest struct {
	TradeType       int    `json:"tradeType"`
	FromTokenCode   string `json:"fromTokenCode"`
	FromTokenAmount string `json:"fromTokenAmount"`
	ToTokenCode     string `json:"toTokenCode"`
	QuoteMode       int    `json:"quoteMode,omitempty"`
}

type DcpSetTimewindowRequest struct {
	Product    string `json:"product"`
	TimeWindow int    `json:"timeWindow"`
}

type AmendOrderRequest struct {
	Category     string `json:"category"`
	Symbol       string `json:"symbol"`
	OrderId      string `json:"orderId"`
	OrderLinkId  string `json:"orderLinkId"`
	OrderIv      string `json:"orderIv"`
	TriggerPrice string `json:"triggerPrice"`
	Qty          string `json:"qty"`
	Price        string `json:"price"`
	TpslMode     string `json:"tpslMode"`
	TakeProfit   string `json:"takeProfit"`
	StopLoss     string `json:"stopLoss"`
	TpTriggerBy  string `json:"tpTriggerBy"`
	SlTriggerBy  string `json:"slTriggerBy"`
	TriggerBy    string `json:"triggerBy"`
	TpLimitPrice string `json:"tpLimitPrice"`
	SlLimitPrice string `json:"slLimitPrice"`
}

type AmendOrderResult struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type CancelAllOrdersRequest struct {
	Category      string `json:"category"`
	Symbol        string `json:"symbol"`
	BaseCoin      string `json:"baseCoin"`
	SettleCoin    string `json:"settleCoin"`
	OrderFilter   string `json:"orderFilter"`
	StopOrderType string `json:"stopOrderType"`
}

type CancelOrderRequest struct {
	Category    string `json:"category"`
	Symbol      string `json:"symbol"`
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
	OrderFilter string `json:"orderFilter"`
}

type CreateOrderResult struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type PreCheckOrderRequest struct {
	Category     string `json:"category"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	OrderType    string `json:"orderType"`
	Qty          string `json:"qty"`
	Price        string `json:"price"`
	IsLeverage   int    `json:"isLeverage"`
	TimeInForce  string `json:"timeInForce"`
	PositionIdx  int    `json:"positionIdx"`
	OrderLinkId  string `json:"orderLinkId"`
	TakeProfit   string `json:"takeProfit"`
	StopLoss     string `json:"stopLoss"`
	TpTriggerBy  string `json:"tpTriggerBy"`
	SlTriggerBy  string `json:"slTriggerBy"`
	ReduceOnly   bool   `json:"reduceOnly"`
	TpslMode     string `json:"tpslMode"`
	TpLimitPrice string `json:"tpLimitPrice"`
	SlLimitPrice string `json:"slLimitPrice"`
	TpOrderType  string `json:"tpOrderType"`
	SlOrderType  string `json:"slOrderType"`
	OrderIv      string `json:"orderIv"`
}

type SetDcpRequest struct {
	Product    string `json:"product"`
	TimeWindow int    `json:"timeWindow"`
}

type AcceptNonLpQuoteResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     AcceptNonLpQuoteResult `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type AcceptNonLpQuoteResult struct {
	RfqId string `json:"rfqId"`
}

type CancelAllQuotesResponse struct {
	RetCode    int                         `json:"retCode"`
	RetMsg     string                      `json:"retMsg"`
	Result     []CancelAllQuotesResultItem `json:"result"`
	RetExtInfo map[string]interface{}      `json:"retExtInfo"`
	Time       int64                       `json:"time"`
}

type CancelAllQuotesResultItem struct {
	RfqId       string `json:"rfqId"`
	QuoteId     string `json:"quoteId"`
	QuoteLinkId string `json:"quoteLinkId"`
	Code        string `json:"code"`
	Msg         string `json:"msg"`
}

type CancelAllRfqsResponse struct {
	RetCode    int                       `json:"retCode"`
	RetMsg     string                    `json:"retMsg"`
	Result     []CancelAllRfqsResultItem `json:"result"`
	RetExtInfo map[string]interface{}    `json:"retExtInfo"`
	Time       int64                     `json:"time"`
}

type CancelAllRfqsResultItem struct {
	RfqId     string `json:"rfqId"`
	RfqLinkId string `json:"rfqLinkId"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
}

type CancelQuoteResult struct {
	RfqId       string `json:"rfqId"`
	QuoteId     string `json:"quoteId"`
	QuoteLinkId string `json:"quoteLinkId"`
}

type CancelRfqResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     CancelRfqResult        `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type CancelRfqResult struct {
	RfqId     string `json:"rfqId"`
	RfqLinkId string `json:"rfqLinkId"`
}

type CreateQuoteResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     CreateQuoteResult      `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type CreateQuoteResult struct {
	RfqId       string `json:"rfqId"`
	QuoteId     string `json:"quoteId"`
	QuoteLinkId string `json:"quoteLinkId"`
	ExpiresAt   string `json:"expiresAt"`
	DeskCode    string `json:"deskCode"`
	Status      string `json:"status"`
}

type CreateRfqResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     CreateRfqResult        `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type CreateRfqResult struct {
	RfqId     string `json:"rfqId"`
	RfqLinkId string `json:"rfqLinkId"`
	Status    string `json:"status"`
	ExpiresAt string `json:"expiresAt"`
	DeskCode  string `json:"deskCode"`
}

type CreateRfqLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	Qty      string `json:"qty"`
}

type ExecuteQuoteResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     ExecuteQuoteResult     `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type ExecuteQuoteResult struct {
	RfqId     string `json:"rfqId"`
	RfqLinkId string `json:"rfqLinkId"`
	QuoteId   string `json:"quoteId"`
	Status    string `json:"status"`
}

type GetPublicTradesResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     GetPublicTradesResult  `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetPublicTradesResult struct {
	Cursor string        `json:"cursor"`
	List   []PublicTrade `json:"list"`
}

type PublicTrade struct {
	RfqId        string           `json:"rfqId"`
	StrategyType string           `json:"strategyType"`
	CreatedAt    string           `json:"createdAt"`
	UpdatedAt    string           `json:"updatedAt"`
	Legs         []PublicTradeLeg `json:"legs"`
}

type PublicTradeLeg struct {
	Category  string `json:"category"`
	Symbol    string `json:"symbol"`
	Side      string `json:"side"`
	Price     string `json:"price"`
	Qty       string `json:"qty"`
	MarkPrice string `json:"markPrice"`
}

type GetQuotesRealtimeItem struct {
	RfqId         string     `json:"rfqId"`
	RfqLinkId     string     `json:"rfqLinkId"`
	QuoteId       string     `json:"quoteId"`
	QuoteLinkId   string     `json:"quoteLinkId"`
	ExpiresAt     string     `json:"expiresAt"`
	Status        string     `json:"status"`
	DeskCode      string     `json:"deskCode"`
	ExecQuoteSide string     `json:"execQuoteSide"`
	CreatedAt     string     `json:"createdAt"`
	UpdatedAt     string     `json:"updatedAt"`
	QuoteBuyList  []QuoteLeg `json:"quoteBuyList"`
	QuoteSellList []QuoteLeg `json:"quoteSellList"`
}

type GetQuotesItem struct {
	RfqId         string         `json:"rfqId"`
	RfqLinkId     string         `json:"rfqLinkId"`
	QuoteId       string         `json:"quoteId"`
	QuoteLinkId   string         `json:"quoteLinkId"`
	ExpiresAt     string         `json:"expiresAt"`
	DeskCode      string         `json:"deskCode"`
	Status        string         `json:"status"`
	ExecQuoteSide string         `json:"execQuoteSide"`
	CreatedAt     string         `json:"createdAt"`
	UpdatedAt     string         `json:"updatedAt"`
	QuoteBuyList  []GetQuotesLeg `json:"quoteBuyList"`
	QuoteSellList []GetQuotesLeg `json:"quoteSellList"`
}

type GetQuotesLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Qty      string `json:"qty"`
}

type GetRfqConfigResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     GetRfqConfigResult     `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetRfqConfigResult struct {
	DeskCode                 string            `json:"deskCode"`
	MaxLegs                  int               `json:"maxLegs"`
	MaxLP                    int               `json:"maxLP"`
	MaxActiveRfq             int               `json:"maxActiveRfq"`
	RfqExpireTime            int               `json:"rfqExpireTime"`
	MinLimitQtySpotOrder     int               `json:"minLimitQtySpotOrder"`
	MinLimitQtyContractOrder int               `json:"minLimitQtyContractOrder"`
	MinLimitQtyOptionOrder   int               `json:"minLimitQtyOptionOrder"`
	StrategyTypes            []RfqStrategyType `json:"strategyTypes"`
	Counterparties           []RfqCounterparty `json:"counterparties"`
}

type RfqStrategyType struct {
	StrategyName string `json:"strategyName"`
}

type RfqCounterparty struct {
	TraderName string `json:"traderName"`
	DeskCode   string `json:"deskCode"`
	Type       string `json:"type"`
}

type GetRfqsRealtimeResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     GetRfqsRealtimeResult  `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type GetRfqsRealtimeResult struct {
	List []RfqRealtimeItem `json:"list"`
}

type RfqRealtimeItem struct {
	RfqId                  string           `json:"rfqId"`
	RfqLinkId              string           `json:"rfqLinkId"`
	Counterparties         []string         `json:"counterparties"`
	ExpiresAt              string           `json:"expiresAt"`
	StrategyType           string           `json:"strategyType"`
	Status                 string           `json:"status"`
	AcceptOtherQuoteStatus string           `json:"acceptOtherQuoteStatus"`
	DeskCode               string           `json:"deskCode"`
	CreatedAt              string           `json:"createdAt"`
	UpdatedAt              string           `json:"updatedAt"`
	Legs                   []RfqRealtimeLeg `json:"legs"`
}

type RfqRealtimeLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	Qty      string `json:"qty"`
}

type GetRfqsItem struct {
	RfqId                  string       `json:"rfqId"`
	RfqLinkId              string       `json:"rfqLinkId"`
	Counterparties         []string     `json:"counterparties"`
	StrategyType           string       `json:"strategyType"`
	ExpiresAt              string       `json:"expiresAt"`
	Status                 string       `json:"status"`
	AcceptOtherQuoteStatus string       `json:"acceptOtherQuoteStatus"`
	DeskCode               string       `json:"deskCode"`
	CreatedAt              string       `json:"createdAt"`
	UpdatedAt              string       `json:"updatedAt"`
	Legs                   []GetRfqsLeg `json:"legs"`
}

type GetRfqsLeg struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	Qty      string `json:"qty"`
}

type GetTradeHistoryResult struct {
	Cursor string                `json:"cursor"`
	List   []GetTradeHistoryItem `json:"list"`
}

type GetTradeHistoryItem struct {
	RfqId         string               `json:"rfqId"`
	RfqLinkId     string               `json:"rfqLinkId"`
	QuoteId       string               `json:"quoteId"`
	QuoteLinkId   string               `json:"quoteLinkId"`
	QuoteSide     string               `json:"quoteSide"`
	StrategyType  string               `json:"strategyType"`
	Status        string               `json:"status"`
	RfqDeskCode   string               `json:"rfqDeskCode"`
	QuoteDeskCode string               `json:"quoteDeskCode"`
	CreatedAt     string               `json:"createdAt"`
	UpdatedAt     string               `json:"updatedAt"`
	Legs          []GetTradeHistoryLeg `json:"legs"`
}

type GetTradeHistoryLeg struct {
	Category      string `json:"category"`
	OrderId       string `json:"orderId"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	Qty           string `json:"qty"`
	MarkPrice     string `json:"markPrice"`
	ExecFee       string `json:"execFee"`
	ExecId        string `json:"execId"`
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	RejectParty   string `json:"rejectParty"`
}
