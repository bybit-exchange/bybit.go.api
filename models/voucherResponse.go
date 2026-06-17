package models

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
