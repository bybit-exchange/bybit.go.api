package models

type CryptoLoanFixedAvailableInventoryResponse struct {
	RetCode    int                                     `json:"retCode"`
	RetMsg     string                                  `json:"retMsg"`
	Result     CryptoLoanFixedAvailableInventoryResult `json:"result"`
	RetExtInfo interface{}                             `json:"retExtInfo"`
	Time       int64                                   `json:"time"`
}

type CryptoLoanFixedAvailableInventoryResult struct {
	Currency           string `json:"currency"`
	Term               int    `json:"term"`
	AnnualRate         string `json:"annualRate"`
	AvailableInventory string `json:"availableInventory"`
	UpdateTime         string `json:"updateTime"`
}
