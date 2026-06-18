package models

type TransactionLogInfo struct {
	List           []TransactionLogEntry `json:"list"`
	NextPageCursor string                `json:"nextPageCursor"`
}

type TransactionLogEntry struct {
	Symbol          string `json:"symbol"`
	Category        string `json:"category"`
	Side            string `json:"side"`
	TransactionTime string `json:"transactionTime"`
	Type            string `json:"type"`
	Qty             string `json:"qty"`
	Size            string `json:"size"`
	Currency        string `json:"currency"`
	TradePrice      string `json:"tradePrice"`
	Funding         string `json:"funding"`
	Fee             string `json:"fee"`
	CashFlow        string `json:"cashFlow"`
	Change          string `json:"change"`
	CashBalance     string `json:"cashBalance"`
	FeeRate         string `json:"feeRate"`
	BonusChange     string `json:"bonusChange"`
	TradeID         string `json:"tradeId"`
	OrderID         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
}

type UpgradeUtaInfo struct {
	UnifiedUpdateStatus string            `json:"unifiedUpdateStatus"`
	UnifiedUpdateMsg    *UnifiedUpdateMsg `json:"unifiedUpdateMsg,omitempty"`
}

type UnifiedUpdateMsg struct {
	Msg []string `json:"msg"`
}

type WalletBalanceInfo struct {
	List []AccountInfo `json:"list"`
}

type WalletAccountInfo struct {
	AccountType            string     `json:"accountType"`
	AccountLTV             string     `json:"accountLTV"`
	AccountIMRate          string     `json:"accountIMRate"`
	AccountMMRate          string     `json:"accountMMRate"`
	TotalEquity            string     `json:"totalEquity"`
	TotalWalletBalance     string     `json:"totalWalletBalance"`
	TotalMarginBalance     string     `json:"totalMarginBalance"`
	TotalAvailableBalance  string     `json:"totalAvailableBalance"`
	TotalPerpUPL           string     `json:"totalPerpUPL"`
	TotalInitialMargin     string     `json:"totalInitialMargin"`
	TotalMaintenanceMargin string     `json:"totalMaintenanceMargin"`
	Coins                  []CoinInfo `json:"coin"`
}

type CoinInfo struct {
	Coin                string `json:"coin"`
	Equity              string `json:"equity"`
	UsdValue            string `json:"usdValue"`
	WalletBalance       string `json:"walletBalance"`
	Free                string `json:"free"`
	Locked              string `json:"locked"`
	BorrowAmount        string `json:"borrowAmount"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	AccruedInterest     string `json:"accruedInterest"`
	TotalOrderIM        string `json:"totalOrderIM"`
	TotalPositionIM     string `json:"totalPositionIM"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Bonus               string `json:"bonus"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	MarginCollateral    bool   `json:"marginCollateral"`
}

type CoinGreeks struct {
	CoinGreeks []CoinGreekInfo `json:"coin"`
}

type CoinGreekInfo struct {
	BaseCoin   string `json:"baseCoin"`
	TotalDelta string `json:"totalDelta"`
	TotalGamma string `json:"totalGamma"`
	TotalVega  string `json:"totalVega"`
	TotalTheta string `json:"totalTheta"`
}

type CollateralInfo struct {
	List []CollateralItem `json:"list"`
}

type CollateralItem struct {
	Currency           string `json:"currency"`
	HourlyBorrowRate   string `json:"hourlyBorrowRate"`
	MaxBorrowingAmount string `json:"maxBorrowingAmount"`
	FreeBorrowingLimit string `json:"freeBorrowingLimit"`
	FreeBorrowAmount   string `json:"freeBorrowAmount"`
	BorrowAmount       string `json:"borrowAmount"`
	AvailableToBorrow  string `json:"availableToBorrow"`
	Borrowable         bool   `json:"borrowable"`
	BorrowUsageRate    string `json:"borrowUsageRate"`
	MarginCollateral   bool   `json:"marginCollateral"`
	CollateralSwitch   bool   `json:"collateralSwitch"`
	CollateralRatio    string `json:"collateralRatio"`
}

type BorrowHistory struct {
	List           []BorrowHistoryItem `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type BorrowHistoryItem struct {
	Currency                  string `json:"currency"`
	CreatedTime               int64  `json:"createdTime"` // Using int64 for milliseconds timestamp
	BorrowCost                string `json:"borrowCost"`
	HourlyBorrowRate          string `json:"hourlyBorrowRate"`
	InterestBearingBorrowSize string `json:"interestBearingBorrowSize"`
	CostExemption             string `json:"costExemption"`
	BorrowAmount              string `json:"borrowAmount"`
	UnrealisedLoss            string `json:"unrealisedLoss"`
	FreeBorrowedAmount        string `json:"freeBorrowedAmount"`
}

type FeeRatesInfo struct {
	Category string        `json:"category"`
	List     []FeeRateItem `json:"list"`
}

type FeeRateItem struct {
	Symbol       string `json:"symbol"`
	BaseCoin     string `json:"baseCoin"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
}

type AccountInfo struct {
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	MarginMode          string `json:"marginMode"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	SmpGroup            int    `json:"smpGroup"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
	UpdatedTime         string `json:"updatedTime"`
}

type MMPStateInfo struct {
	Result []MMPStateItem `json:"result"`
}

type MMPStateItem struct {
	BaseCoin     string `json:"baseCoin"`
	MmpEnabled   bool   `json:"mmpEnabled"`
	Window       string `json:"window"`
	FrozenPeriod string `json:"frozenPeriod"`
	QtyLimit     string `json:"qtyLimit"`
	DeltaLimit   string `json:"deltaLimit"`
}

type MarginMode struct {
	Reasons        []ReasonItem `json:"reasons"`
	MmpFrozenUntil string       `json:"mmpFrozenUntil"`
	MmpFrozen      bool         `json:"mmpFrozen"`
}

type ReasonItem struct {
	ReasonCode string `json:"reasonCode"`
	ReasonMsg  string `json:"reasonMsg"`
}

type QueryBrokerAllUidResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type ApiLimitInfo struct {
	Uids    string `json:"uids"`
	BizType string `json:"bizType"`
	Rate    int    `json:"rate"`
}

type QueryBrokerCapResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type BrokerCapInfo struct {
	BizType   string `json:"bizType"`
	TotalRate string `json:"totalRate"`
	EbCap     string `json:"ebCap"`
	UidCap    string `json:"uidCap"`
}

type SetApiLimitResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type ApiLimitInfoResult struct {
	Uids    string `json:"uids"`
	BizType string `json:"bizType"`
	Rate    int    `json:"rate"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type BatchSetCollateralResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetAccountInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetAccountInfoResult struct {
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	MarginMode          string `json:"marginMode"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
	SpotHedgingStatus   string `json:"spotHedgingStatus"`
	UpdatedTime         string `json:"updatedTime"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	SmpGroup            int    `json:"smpGroup"`
}

type BorrowHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type BorrowHistoryResult struct {
	List           []BorrowHistoryItem `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type GetCoinGreeksResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type CoinGreeksItem struct {
	BaseCoin   string `json:"baseCoin"`
	TotalDelta string `json:"totalDelta"`
	TotalGamma string `json:"totalGamma"`
	TotalVega  string `json:"totalVega"`
	TotalTheta string `json:"totalTheta"`
}

type GetCollateralInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type CollateralInfoItem struct {
	Currency            string `json:"currency"`
	HourlyBorrowRate    string `json:"hourlyBorrowRate"`
	MaxBorrowingAmount  string `json:"maxBorrowingAmount"`
	FreeBorrowingLimit  string `json:"freeBorrowingLimit"`
	FreeBorrowAmount    string `json:"freeBorrowAmount"`
	BorrowAmount        string `json:"borrowAmount"`
	OtherBorrowAmount   string `json:"otherBorrowAmount"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	Borrowable          bool   `json:"borrowable"`
	BorrowUsageRate     string `json:"borrowUsageRate"`
	MarginCollateral    bool   `json:"marginCollateral"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	FreeBorrowingAmount string `json:"freeBorrowingAmount"`
	CollateralRatio     string `json:"collateralRatio"`
}

type GetFeeRateResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetFeeRateResult struct {
	Category string        `json:"category"`
	List     []FeeRateItem `json:"list"`
}

type GetMmpStateResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type MmpStateItem struct {
	BaseCoin       string `json:"baseCoin"`
	MmpEnabled     bool   `json:"mmpEnabled"`
	Window         string `json:"window"`
	FrozenPeriod   string `json:"frozenPeriod"`
	QtyLimit       string `json:"qtyLimit"`
	DeltaLimit     string `json:"deltaLimit"`
	MmpFrozenUntil string `json:"mmpFrozenUntil"`
	MmpFrozen      bool   `json:"mmpFrozen"`
}

type GetSmpGroupResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetTransactionLogResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TransactionLogItem struct {
	Id              string `json:"id"`
	Symbol          string `json:"symbol"`
	Category        string `json:"category"`
	Side            string `json:"side"`
	TransactionTime string `json:"transactionTime"`
	Type            string `json:"type"`
	TransSubType    string `json:"transSubType"`
	Qty             string `json:"qty"`
	Size            string `json:"size"`
	Currency        string `json:"currency"`
	TradePrice      string `json:"tradePrice"`
	Funding         string `json:"funding"`
	Fee             string `json:"fee"`
	CashFlow        string `json:"cashFlow"`
	Change          string `json:"change"`
	CashBalance     string `json:"cashBalance"`
	FeeRate         string `json:"feeRate"`
	BonusChange     string `json:"bonusChange"`
	TradeId         string `json:"tradeId"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	ExtraFees       string `json:"extraFees"`
}

type GetUserSettingsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetUserSettingsResult struct {
	LpaSpot bool `json:"lpaSpot"`
	LpaPerp bool `json:"lpaPerp"`
}

type ManualBorrowResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type ManualBorrowResult struct {
	Coin   string `json:"coin"`
	Amount string `json:"amount"`
}

type ManualRepayResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type OneClickRepayResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type OneClickRepayResultItem struct {
	Coin         string `json:"coin"`
	RepaymentQty string `json:"repaymentQty"`
}

type SetMarginModeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type SetMmpResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type UpgradeToUtaProResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type UpgradeToUtaProResult struct {
	UnifiedUpdateStatus string      `json:"unifiedUpdateStatus"`
	UnifiedUpdateMsg    interface{} `json:"unifiedUpdateMsg"`
}

type UpgradeToUtaProUpdateMsg struct {
	Msg []string `json:"msg"`
}

type InfoResponse struct {
	RetCode int        `json:"retCode"`
	RetMsg  string     `json:"retMsg"`
	Result  InfoResult `json:"result"`
}

type InfoResult struct {
	MarginMode          string `json:"marginMode"`
	UpdatedTime         string `json:"updatedTime"`
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	SmpGroup            int    `json:"smpGroup"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
	SpotHedgingStatus   string `json:"spotHedgingStatus"`
}

type UtaDemoApplyMoneyResponse struct {
	ResultCode              string                    `json:"resultCode"`
	UtaDemoApplyMoneyConfig []UtaDemoApplyMoneyConfig `json:"utaDemoApplyMoneyConfig"`
	OrderStatus             string                    `json:"orderStatus"`
	RetMsg                  string                    `json:"retMsg"`
}

type UtaDemoApplyMoneyConfig struct {
	Coin      string `json:"coin"`
	AmountStr string `json:"amountStr"`
}

type OpenRepaymentResp struct {
	List []OpenRepaymentResult `json:"list"`
}

type OpenRepaymentResult struct {
	Coin         string `json:"coin"`
	RepaymentQty string `json:"repaymentQty"`
}

type SetHedgingModeResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type SetMarginModeLimitReason struct {
	Reasons []interface{} `json:"reasons"`
}

type SetMarginModeLimit struct {
	ReasonCode string `json:"reasonCode"`
	ReasonMsg  string `json:"reasonMsg"`
}

type UpgradeUtaOpenApiResponse struct {
	UnifiedUpdateStatus string      `json:"unifiedUpdateStatus"`
	UnifiedUpdateMsg    interface{} `json:"unifiedUpdateMsg"`
}

type GetTradeInfoForAnalysisResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type TradeAnalysisDailySummary struct {
	Day              string `json:"day"`
	SumBuyExecValue  string `json:"sumBuyExecValue"`
	SumSellExecValue string `json:"sumSellExecValue"`
	SumExecValue     string `json:"sumExecValue"`
}

type AccountBorrowResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type BatchSetCollateralResultItem struct {
	Coin             string `json:"coin"`
	CollateralSwitch string `json:"collateralSwitch"`
}

type GetBorrowHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type GetBorrowHistoryResult struct {
	List           []GetBorrowHistoryItem `json:"list"`
	NextPageCursor string                 `json:"nextPageCursor"`
}

type GetBorrowHistoryItem struct {
	Currency                  string `json:"currency"`
	CreatedTime               int64  `json:"createdTime"`
	BorrowCost                string `json:"borrowCost"`
	HourlyBorrowRate          string `json:"hourlyBorrowRate"`
	InterestBearingBorrowSize string `json:"interestBearingBorrowSize"`
	CostExemption             string `json:"costExemption"`
	BorrowAmount              string `json:"borrowAmount"`
	UnrealisedLoss            string `json:"unrealisedLoss"`
	FreeBorrowedAmount        string `json:"freeBorrowedAmount"`
}

type GetDcpInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type DcpInfo struct {
	Product    string `json:"product"`
	DcpStatus  string `json:"dcpStatus"`
	TimeWindow string `json:"timeWindow"`
}

type TransactionLogResult struct {
	List           []TransactionLogItem `json:"list"`
	NextPageCursor string               `json:"nextPageCursor"`
}

type GetTransferableAmountResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type ManualRepayResult struct {
	ResultStatus string `json:"resultStatus"`
}

type NoConvertRepayResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type NoConvertRepayResult struct {
	ResultStatus string `json:"resultStatus"`
}

type SetMarginModeResult struct {
	Reasons []interface{} `json:"reasons"`
}

type SetMarginModeReason struct {
	ReasonCode string `json:"reasonCode"`
	ReasonMsg  string `json:"reasonMsg"`
}
