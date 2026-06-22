package models

type QueryDcpInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type BizDcpInfo struct {
	Product    string `json:"product"`
	DcpStatus  string `json:"dcpStatus"`
	TimeWindow string `json:"timeWindow"`
}
