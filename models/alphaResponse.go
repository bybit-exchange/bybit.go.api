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
	Time       int64       `json:"time"`
}

type PredictionEngineStatusResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
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
	Time       int64       `json:"time"`
}

type PredictionOrderBookRequest struct {
	TokenIds []string `json:"tokenIds"`
}

type PredictionOrderBookResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     []interface{} `json:"result"`
	RetExtInfo interface{}   `json:"retExtInfo"`
	Time       int64         `json:"time"`
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
	Time       int64       `json:"time"`
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
	Time       int64       `json:"time"`
}

type PredictionPayTokenListResponse struct {
	RetCode    int                     `json:"retCode"`
	RetMsg     string                  `json:"retMsg"`
	Result     []PredictionPayTokenDto `json:"result"`
	RetExtInfo interface{}             `json:"retExtInfo"`
	Time       int64                   `json:"time"`
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
	Time       int64       `json:"time"`
}

type PredictionPositionHistoryRequest struct {
	TokenId   string                   `json:"tokenId"`
	EventId   string                   `json:"eventId"`
	Result    PredictionPositionResult `json:"result"`
	Days      int                      `json:"days"`
	Limit     int                      `json:"limit"`
	PageIndex int                      `json:"pageIndex"`
}

type PredictionPositionResult int

type PredictionPositionHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PredictionPositionListRequest struct {
	TokenId   string `json:"tokenId"`
	EventId   string `json:"eventId"`
	Limit     int    `json:"limit"`
	PageIndex int    `json:"pageIndex"`
}

type PredictionPositionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PredictionPriceHistoryRequest struct {
	TokenId  string                  `json:"tokenId"`
	Interval PredictionPriceInterval `json:"interval"`
}

type PredictionPriceInterval string

type PredictionPriceHistoryResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     []interface{} `json:"result"`
	RetExtInfo interface{}   `json:"retExtInfo"`
	Time       int64         `json:"time"`
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
	Time       int64       `json:"time"`
}

type PredictionSideMarketListRequest struct {
	EventId    string               `json:"eventId"`
	SortBy     PredictionSortBy     `json:"sortBy"`
	MarketType PredictionMarketType `json:"marketType"`
	Limit      int                  `json:"limit"`
	PageIndex  int                  `json:"pageIndex"`
}

type PredictionSortBy string

type PredictionMarketType int

type PredictionEventType int

type PredictionSideMarketListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PredictionGroupStageDetailRequest struct {
	EventType PredictionEventType `json:"eventType"`
	GroupName string              `json:"groupName"`
}

type PredictionStageCode string

type PredictionGroupStageDetailResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PredictionMatchListRequest struct {
	EventType PredictionEventType   `json:"eventType"`
	Status    PredictionMatchStatus `json:"status"`
	StartTime int64                 `json:"startTime"`
	EndTime   int64                 `json:"endTime"`
	StageCode PredictionStageCode   `json:"stageCode"`
	Limit     int                   `json:"limit"`
	PageIndex int                   `json:"pageIndex"`
}

type PredictionMatchStatus int

type PredictionMatchListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PredictionTimelineStagesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type PredictionTokenPriceRequest struct {
	TokenIds []string `json:"tokenIds"`
}

type PredictionTokenPriceResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     []interface{} `json:"result"`
	RetExtInfo interface{}   `json:"retExtInfo"`
	Time       int64         `json:"time"`
}
