package models

type GetAffiliateSubListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AffiliateSubItem struct {
	SubAffId            string      `json:"subAffId"`
	UserId              string      `json:"userId"`
	Name                string      `json:"name"`
	Email               string      `json:"email"`
	CommissionsVol      interface{} `json:"commissionsVol"`
	CommissionsForUsdt  string      `json:"commissionsForUsdt"`
	BecameAffTime       string      `json:"becameAffTime"`
	StartDate           string      `json:"startDate"`
	EndDate             string      `json:"endDate"`
}

type GetAffiliateUserListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetAffiliateUserInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AffiliateUserItem struct {
	UserId               string      `json:"userId"`
	RegisterTime         string      `json:"registerTime"`
	Source               string      `json:"source"`
	Remarks              string      `json:"remarks"`
	IsKyc                bool        `json:"isKyc"`
	TakerVol30Day        string      `json:"takerVol30Day"`
	MakerVol30Day        string      `json:"makerVol30Day"`
	TradeVol30Day        string      `json:"tradeVol30Day"`
	TradfiTradeVol30Day  string      `json:"tradfiTradeVol30Day"`
	Commissions30Day     interface{} `json:"commissions30Day"`
	DepositAmount30Day   string      `json:"depositAmount30Day"`
	DepositAmount365Day  string      `json:"depositAmount365Day"`
	TakerVol365Day       string      `json:"takerVol365Day"`
	MakerVol365Day       string      `json:"makerVol365Day"`
	TradeVol365Day       string      `json:"tradeVol365Day"`
	TradfiTradeVol365Day string      `json:"tradfiTradeVol365Day"`
	Commissions365Day    interface{} `json:"commissions365Day"`
	TakerVol             string      `json:"takerVol"`
	MakerVol             string      `json:"makerVol"`
	TradeVol             string      `json:"tradeVol"`
	TradfiTradeVol       string      `json:"tradfiTradeVol"`
	CommissionsVol       interface{} `json:"commissionsVol"`
	StartDate            string      `json:"startDate"`
	EndDate              string      `json:"endDate"`
}

type AffiliateUserInfoResult struct {
	Uid                 string      `json:"uid"`
	VipLevel            string      `json:"vipLevel"`
	TakerVol30Day       string      `json:"takerVol30Day"`
	MakerVol30Day       string      `json:"makerVol30Day"`
	TradeVol30Day       string      `json:"tradeVol30Day"`
	DepositAmount30Day  string      `json:"depositAmount30Day"`
	TakerVol365Day      string      `json:"takerVol365Day"`
	MakerVol365Day      string      `json:"makerVol365Day"`
	TradeVol365Day      string      `json:"tradeVol365Day"`
	DepositAmount365Day string      `json:"depositAmount365Day"`
	TotalWalletBalance  string      `json:"totalWalletBalance"`
	KycLevel            int         `json:"KycLevel"`
	Commissions30Day    interface{} `json:"commissions30Day"`
	Commissions365Day   interface{} `json:"commissions365Day"`
	DepositUpdateTime   string      `json:"depositUpdateTime"`
	VolUpdateTime       string      `json:"volUpdateTime"`
}
