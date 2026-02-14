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
	List []WalletAccountInfo `json:"list"`
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
	InterestBearingBorrowSize string `json:"InterestBearingBorrowSize"`
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
