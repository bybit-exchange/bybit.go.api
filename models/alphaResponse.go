package models

type PredictionBuyRequest struct {
	TokenId      string              `json:"tokenId"`
	Amount       string              `json:"amount"`
	PayTokenCode string              `json:"payTokenCode"`
	OrderType    PredictionOrderType `json:"orderType"`
	Slippage     string              `json:"slippage"`
	EventId      string              `json:"eventId"`
}

type PredictionOrderType int

type PredictionBuyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionEngineStatusResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionEventDetailRequest struct {
	EventId        string `json:"eventId"`
	Slug           string `json:"slug"`
	HasMoreMarkets bool   `json:"hasMoreMarkets"`
}

type PredictionEventDetailResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionOrderBookRequest struct {
	TokenIds []string `json:"tokenIds"`
}

type PredictionOrderBookResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     []interface{} `json:"result"`
	RetExtInfo interface{}   `json:"retExtInfo"`
	Time       int           `json:"time"`
}

type PredictionOrderEstimateRequest struct {
	TokenId      string              `json:"tokenId"`
	Side         PredictionSide      `json:"side"`
	EventId      string              `json:"eventId"`
	Amount       string              `json:"amount"`
	OrderType    PredictionOrderType `json:"orderType"`
	PayTokenCode string              `json:"payTokenCode"`
}

type PredictionSide int

type PredictionOrderEstimateResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionOrderListRequest struct {
	Status    PredictionOrderStatus `json:"status"`
	TokenId   string                `json:"tokenId"`
	Limit     int                   `json:"limit"`
	PageIndex int                   `json:"pageIndex"`
	EventId   string                `json:"eventId"`
	Side      PredictionSide        `json:"side"`
	Days      int                   `json:"days"`
}

type PredictionOrderStatus int

type PredictionOrderListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionPayTokenListResponse struct {
	RetCode    int                     `json:"retCode"`
	RetMsg     string                  `json:"retMsg"`
	Result     []PredictionPayTokenDto `json:"result"`
	RetExtInfo interface{}             `json:"retExtInfo"`
	Time       int                     `json:"time"`
}

type PredictionPayTokenDto struct {
	TokenCode         string   `json:"tokenCode"`
	Symbol            string   `json:"symbol"`
	TokenDecimals     int      `json:"tokenDecimals"`
	TokenIconUrlDay   string   `json:"tokenIconUrlDay"`
	TokenIconUrlNight string   `json:"tokenIconUrlNight"`
	SupportChains     []string `json:"supportChains"`
}

type PredictionPortfolioSummaryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionPositionHistoryRequest struct {
	Limit     int    `json:"limit"`
	PageIndex int    `json:"pageIndex"`
	Direction string `json:"direction"`
}

type PredictionPositionHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionPositionListRequest struct {
	Limit     int    `json:"limit"`
	PageIndex int    `json:"pageIndex"`
	Direction string `json:"direction"`
}

type PredictionPositionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionPriceHistoryRequest struct {
	TokenIds []string                `json:"tokenIds"`
	EventId  string                  `json:"eventId"`
	Interval PredictionPriceInterval `json:"interval"`
	Fidelity int                     `json:"fidelity"`
}

type PredictionPriceInterval string

type PredictionPriceHistoryResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     []interface{} `json:"result"`
	RetExtInfo interface{}   `json:"retExtInfo"`
	Time       int           `json:"time"`
}

type PredictionSellRequest struct {
	TokenId     string              `json:"tokenId"`
	Size        string              `json:"size"`
	OrderType   PredictionOrderType `json:"orderType"`
	Slippage    string              `json:"slippage"`
	EventId     string              `json:"eventId"`
	ToTokenCode string              `json:"toTokenCode"`
}

type PredictionSellResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionSideMarketListRequest struct {
	EventType PredictionEventType `json:"eventType"`
}

type PredictionEventType int

type PredictionSideMarketListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionGroupStageDetailRequest struct {
	EventType PredictionEventType `json:"eventType"`
	StageCode PredictionStageCode `json:"stageCode"`
}

type PredictionStageCode string

type PredictionGroupStageDetailResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionMatchListRequest struct {
	EventType PredictionEventType `json:"eventType"`
}

type PredictionMatchListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionTimelineStagesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PredictionTokenPriceRequest struct {
	TokenIds []string `json:"tokenIds"`
}

type PredictionTokenPriceResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     []interface{} `json:"result"`
	RetExtInfo interface{}   `json:"retExtInfo"`
	Time       int           `json:"time"`
}
