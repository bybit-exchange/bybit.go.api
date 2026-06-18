package models

type BorrowOrderQuoteResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type BorrowQuote struct {
	OrderCurrency string `json:"orderCurrency"`
	Term          int    `json:"term"`
	AnnualRate    string `json:"annualRate"`
	Qty           string `json:"qty"`
}

type FixedBorrowResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}
