package models

// BrokerEarningInfo represents an individual record of broker earnings.
type BrokerEarningInfo struct {
	UserId   string `json:"userId"`
	BizType  string `json:"bizType"`
	Symbol   string `json:"symbol"`
	Coin     string `json:"coin"`
	Earning  string `json:"earning"`
	OrderId  string `json:"orderId"`
	ExecTime string `json:"execTime"`
}

// BrokerEarningResult represents the paginated result of broker earnings.
type BrokerEarningResult struct {
	List           []BrokerEarningInfo `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type BrokerAccountInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type BaseFeeRebateRate struct {
	Spot        string `json:"spot"`
	Derivatives string `json:"derivatives"`
}

type MarkupFeeRebateRate struct {
	Spot        string `json:"spot"`
	Derivatives string `json:"derivatives"`
	Convert     string `json:"convert"`
}

type BrokerEarningResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type TotalEarningCat struct {
	Spot        []TotalEarning `json:"spot"`
	Derivatives []TotalEarning `json:"derivatives"`
	Options     []TotalEarning `json:"options"`
	Convert     []TotalEarning `json:"convert"`
	Total       []TotalEarning `json:"total"`
}

type TotalEarning struct {
	Coin    string `json:"coin"`
	Earning string `json:"earning"`
}

type RebateDetail struct {
	UserId         string `json:"userId"`
	BizType        string `json:"bizType"`
	Symbol         string `json:"symbol"`
	Coin           string `json:"coin"`
	Earning        string `json:"earning"`
	MarkupEarning  string `json:"markupEarning"`
	BaseFeeEarning string `json:"baseFeeEarning"`
	OrderId        string `json:"orderId"`
	ExecTime       string `json:"execTime"`
	ExecId         string `json:"execId"`
}

type SetApiLimitRequest struct {
	List []ApiLimitInfo `json:"list"`
}

type AwardInfo struct {
	Id             string `json:"id"`
	Coin           string `json:"coin"`
	AmountUnit     string `json:"amountUnit"`
	ProductLine    string `json:"productLine"`
	SubProductLine string `json:"subProductLine"`
	TotalAmount    string `json:"totalAmount"`
	UsedAmount     string `json:"usedAmount"`
}

type DistributionRecord struct {
	AccountId     string `json:"accountId"`
	AwardId       string `json:"awardId"`
	SpecCode      string `json:"specCode"`
	Amount        string `json:"amount"`
	IsClaimed     bool   `json:"isClaimed"`
	StartAt       string `json:"startAt"`
	EndAt         string `json:"endAt"`
	EffectiveAt   string `json:"effectiveAt"`
	IneffectiveAt string `json:"ineffectiveAt"`
	UsedAmount    string `json:"usedAmount"`
}
