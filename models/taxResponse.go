package models

type ApiError struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type TaxBatchCreateRequest struct {
	StartTime         int64           `json:"startTime"`
	EndTime           int64           `json:"endTime"`
	Items             []TaxExportItem `json:"items"`
	SourceInstitution string          `json:"sourceInstitution"`
	ExportFileType    string          `json:"exportFileType"`
}

type TaxExportItem struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type TaxBatchCreateResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type TaxBatchQueryResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}
