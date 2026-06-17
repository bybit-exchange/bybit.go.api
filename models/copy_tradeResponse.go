package models

type CreateCopyTradeBindResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CreateCopyMt5BindResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetCopyTradingClassicLeaderboardResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CopyTradingClassicLeaderboardItem struct {
	LeaderMark           string `json:"leaderMark"`
	Nickname             string `json:"nickname"`
	ThirtyDayRoi         string `json:"thirtyDayRoi"`
	ThirtyDayMaxDrawdown string `json:"thirtyDayMaxDrawdown"`
	ThirtyDaySharpeRatio string `json:"thirtyDaySharpeRatio"`
}

type GetCopyTradingTradFiLeaderboardResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CopyTradingTradFiLeaderboardItem struct {
	ProviderMark         string `json:"providerMark"`
	Nickname             string `json:"nickname"`
	ThirtyDayRoe         string `json:"thirtyDayRoe"`
	ThirtyDayMaxDrawdown string `json:"thirtyDayMaxDrawdown"`
	ThirtyDaySharpeRatio string `json:"thirtyDaySharpeRatio"`
}

type CreateCopyTradeBindRequest struct {
	LeaderMark  string `json:"leaderMark"`
	InvestmentE8 string `json:"investmentE8"`
}
