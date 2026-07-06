package models

type GetSpreadOpenOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SpreadOpenOrder struct {
	Symbol      string `json:"symbol"`
	BaseCoin    string `json:"baseCoin"`
	OrderType   string `json:"orderType"`
	OrderLinkId string `json:"orderLinkId"`
	Side        string `json:"side"`
	TimeInForce string `json:"timeInForce"`
	OrderId     string `json:"orderId"`
	LeavesQty   string `json:"leavesQty"`
	OrderStatus string `json:"orderStatus"`
	CumExecQty  string `json:"cumExecQty"`
	Price       string `json:"price"`
	Qty         string `json:"qty"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type GetSpreadOrderHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SpreadOrderHistoryItem struct {
	Symbol       string `json:"symbol"`
	BaseCoin     string `json:"baseCoin"`
	SettleCoin   string `json:"settleCoin"`
	OrderType    string `json:"orderType"`
	OrderId      string `json:"orderId"`
	OrderLinkId  string `json:"orderLinkId"`
	ContractType string `json:"contractType"`
	OrderStatus  string `json:"orderStatus"`
	Side         string `json:"side"`
	Price        string `json:"price"`
	OrderQty     string `json:"orderQty"`
	Qty          string `json:"qty"`
	TimeInForce  string `json:"timeInForce"`
	CxlRejReason string `json:"cxlRejReason"`
	LeavesQty    string `json:"leavesQty"`
	CumExecQty   string `json:"cumExecQty"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	Leg1Symbol   string `json:"leg1Symbol"`
	Leg1ProdType string `json:"leg1ProdType"`
	Leg1OrderId  string `json:"leg1OrderId"`
	Leg1Side     string `json:"leg1Side"`
	Leg2Symbol   string `json:"leg2Symbol"`
	Leg2ProdType string `json:"leg2ProdType"`
	Leg2OrderId  string `json:"leg2OrderId"`
	Leg2Side     string `json:"leg2Side"`
}

type GetSpreadTradeHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SpreadTradeExecutionRecord struct {
	Symbol      string                    `json:"symbol"`
	OrderLinkId string                    `json:"orderLinkId"`
	Side        string                    `json:"side"`
	OrderId     string                    `json:"orderId"`
	ExecPrice   string                    `json:"execPrice"`
	ExecTime    string                    `json:"execTime"`
	ExecType    string                    `json:"execType"`
	ExecQty     string                    `json:"execQty"`
	ExecId      string                    `json:"execId"`
	Legs        []SpreadTradeExecutionLeg `json:"legs"`
}

type SpreadTradeExecutionLeg struct {
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	ExecPrice   string `json:"execPrice"`
	ExecTime    string `json:"execTime"`
	ExecValue   string `json:"execValue"`
	ExecType    string `json:"execType"`
	Category    string `json:"category"`
	ExecQty     string `json:"execQty"`
	ExecFee     string `json:"execFee"`
	ExecFeeV2   string `json:"execFeeV2"`
	FeeCurrency string `json:"feeCurrency"`
	ExecId      string `json:"execId"`
}
