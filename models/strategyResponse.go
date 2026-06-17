package models

type StrategyCreateResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
	ExtCode string      `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	TimeNow string      `json:"time_now"`
}

type ChaseOrderStrategyRequest struct {
	Category       string `json:"category"`
	Symbol         string `json:"symbol"`
	Side           string `json:"side"`
	Size           string `json:"size"`
	StrategyType   string `json:"strategyType"`
	ChaseDistance  string `json:"chaseDistance"`
	ChasePercentE4 int    `json:"chasePercentE4"`
	MaxChasePrice  string `json:"maxChasePrice"`
	TriggerPrice   string `json:"triggerPrice"`
	ReduceOnly     bool   `json:"reduceOnly"`
	PositionIdx    int    `json:"positionIdx"`
	LeverageType   int    `json:"leverageType"`
}

type OrderQueryResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
	ExtCode string      `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	TimeNow string      `json:"time_now"`
}

type OrderDetail struct {
	StrategyId     string `json:"strategyId"`
	OrderId        string `json:"orderId"`
	Symbol         string `json:"symbol"`
	Side           string `json:"side"`
	Size           string `json:"size"`
	Price          string `json:"price"`
	Status         string `json:"status"`
	ExecutedSize   string `json:"executedSize"`
	DealTimeE3     int    `json:"dealTimeE3"`
	ParentOrderId  string `json:"parentOrderId"`
	CreatedTimeE3  int    `json:"createdTimeE3"`
	UpdatedTimeE3  int    `json:"updatedTimeE3"`
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	Category       string `json:"category"`
	PositionIdx    int    `json:"positionIdx"`
	LeverageType   int    `json:"leverageType"`
}

type PovStrategyRequest struct {
	Category     string      `json:"category"`
	Symbol       string      `json:"symbol"`
	Side         string      `json:"side"`
	Size         string      `json:"size"`
	StrategyType string      `json:"strategyType"`
	Duration     int         `json:"duration"`
	Interval     int         `json:"interval"`
	PovParams    interface{} `json:"povParams"`
	ReduceOnly   bool        `json:"reduceOnly"`
	PositionIdx  int         `json:"positionIdx"`
}

type PovParams struct {
	Mode              string `json:"mode"`
	ParticipationRate string `json:"participationRate"`
	ReferenceWindow   string `json:"referenceWindow"`
	DepthReference    int    `json:"depthReference"`
}

type StrategyQueryResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
	ExtCode string      `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	TimeNow string      `json:"time_now"`
}

type StrategyDetail struct {
	StrategyId           string `json:"strategyId"`
	Category             string `json:"category"`
	Symbol               string `json:"symbol"`
	Side                 string `json:"side"`
	Size                 string `json:"size"`
	ExecutedSize         string `json:"executedSize"`
	ExecutedAvgPrice     string `json:"executedAvgPrice"`
	Status               int    `json:"status"`
	StrategyType         string `json:"strategyType"`
	Duration             int    `json:"duration"`
	ExecutedDuration     int    `json:"executedDuration"`
	ExecutedStartTimeE3  int    `json:"executedStartTimeE3"`
	ExecutedEndTimeE3    int    `json:"executedEndTimeE3"`
	CreatedTimeE3        int    `json:"createdTimeE3"`
	UpdatedTimeE3        int    `json:"updatedTimeE3"`
	IsRandom             bool   `json:"isRandom"`
	Interval             int    `json:"interval"`
	ReduceOnly           bool   `json:"reduceOnly"`
	PositionIdx          int    `json:"positionIdx"`
	LeverageType         int    `json:"leverageType"`
	ChasePrice           string `json:"chasePrice"`
	ChasePercentE4       int    `json:"chasePercentE4"`
	ChaseDistance        string `json:"chaseDistance"`
	MaxChasePrice        string `json:"maxChasePrice"`
	ChaseOrderPrice      string `json:"chaseOrderPrice"`
	TriggerPrice         string `json:"triggerPrice"`
	IsTriggered          bool   `json:"isTriggered"`
	PostOnly             int    `json:"postOnly"`
	TerminateType        int    `json:"terminateType"`
	TerminateRemark      string `json:"terminateRemark"`
	TriggerCount         int    `json:"triggerCount"`
	TradingCount         int    `json:"tradingCount"`
	RealizedPnl          string `json:"realizedPnl"`
	Mode                 string `json:"mode"`
	ParticipationRate    string `json:"participationRate"`
	ReferenceWindow      string `json:"referenceWindow"`
	DepthReference       int    `json:"depthReference"`
}

type TwapStrategyRequest struct {
	Category       string `json:"category"`
	Symbol         string `json:"symbol"`
	Side           string `json:"side"`
	Size           string `json:"size"`
	StrategyType   string `json:"strategyType"`
	Duration       int    `json:"duration"`
	Interval       int    `json:"interval"`
	IsRandom       bool   `json:"isRandom"`
	TriggerPrice   string `json:"triggerPrice"`
	MaxChasePrice  string `json:"maxChasePrice"`
	ChaseDistance  string `json:"chaseDistance"`
	ChasePercentE4 int    `json:"chasePercentE4"`
	ReduceOnly     bool   `json:"reduceOnly"`
	PositionIdx    int    `json:"positionIdx"`
	LeverageType   int    `json:"leverageType"`
}

type StrategyErrorResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}
