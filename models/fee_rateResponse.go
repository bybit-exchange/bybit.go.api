package models

type FeeRateEntity struct {
	BaseCoin     string `json:"baseCoin"`
	Symbol       string `json:"symbol"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
}

type GroupFeeRateEntity struct {
	GroupName       string          `json:"groupName"`
	WeightingFactor int             `json:"weightingFactor"`
	SymbolsNumbers  int             `json:"symbolsNumbers"`
	Symbols         []string        `json:"symbols"`
	FeeRates        interface{}     `json:"feeRates"`
	UpdateTime      int             `json:"updateTime"`
}

type FeeRateDetailMap struct {
	Pro          []FeeRateDetail `json:"pro"`
	MarketMaker  []FeeRateDetail `json:"marketMaker"`
}

type FeeRateDetail struct {
	Level        string `json:"level"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
	MakerRebate  string `json:"makerRebate"`
}
