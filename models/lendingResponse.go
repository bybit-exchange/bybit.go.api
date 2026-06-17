package models

// LendingCoin represents information about a specific coin available for lending.
type LendingCoin struct {
	Coin            string `json:"coin"`
	MaxRedeemQty    string `json:"maxRedeemQty"`
	MinPurchaseQty  string `json:"minPurchaseQty"`
	Precision       string `json:"precision"`
	Rate            string `json:"rate"`
	LoanToPoolRatio string `json:"loanToPoolRatio"`
	ActualApy       string `json:"actualApy"`
}

// LendingCoinInfoResult represents the list of coins available for lending.
type LendingCoinInfoResult struct {
	List []LendingCoin `json:"list"`
}

// DepositFund represents the status of a deposit into a lending product.
type DepositFund struct {
	Coin        string `json:"coin"`
	CreatedTime string `json:"createdTime"`
	OrderId     string `json:"orderId"`
	Quantity    string `json:"quantity"`
	SerialNo    string `json:"serialNo"`
	Status      string `json:"status"`
	UpdatedTime string `json:"updatedTime"`
}

// RedeemFund represents the status of a redemption request from a lending product.
type RedeemFund struct {
	Coin         string `json:"coin"`
	CreatedTime  string `json:"createdTime"`
	OrderId      string `json:"orderId"`
	PrincipalQty string `json:"principalQty"`
	SerialNo     string `json:"serialNo"`
	Status       string `json:"status"`
	UpdatedTime  string `json:"updatedTime"`
}

// CancelRedeemFund represents the status of a cancelled redemption request from a lending product.
type CancelRedeemFund struct {
	OrderId     string `json:"orderId"`
	SerialNo    string `json:"serialNo"`
	UpdatedTime string `json:"updatedTime"`
}

// LendingOrderRecord represents an individual lending order record.
type LendingOrderRecord struct {
	Coin        string `json:"coin"`
	CreatedTime string `json:"createdTime"`
	OrderId     string `json:"orderId"`
	Quantity    string `json:"quantity"`
	SerialNo    string `json:"serialNo"`
	Status      string `json:"status"`
	UpdatedTime string `json:"updatedTime"`
}

// LendingOrdersRecordsResult represents the result of a query for lending order records.
type LendingOrdersRecordsResult struct {
	List []LendingOrderRecord `json:"list"`
}

type MarginProductInfo struct {
	ProductId              string         `json:"productId"`
	Leverage               string         `json:"leverage"`
	SupportSpot            int            `json:"supportSpot"`
	SupportContract        int            `json:"supportContract"`
	SupportMarginTrading   int            `json:"supportMarginTrading"`
	WithdrawLine           string         `json:"withdrawLine"`
	TransferLine           string         `json:"transferLine"`
	SpotBuyLine            string         `json:"spotBuyLine"`
	SpotSellLine           string         `json:"spotSellLine"`
	ContractOpenLine       string         `json:"contractOpenLine"`
	LiquidationLine        string         `json:"liquidationLine"`
	StopLiquidationLine    string         `json:"stopLiquidationLine"`
	ContractLeverage       string         `json:"contractLeverage"`
	TransferRatio          string         `json:"transferRatio"`
	SpotSymbols            []string       `json:"spotSymbols"`
	ContractSymbols        []string       `json:"contractSymbols"`
	SupportUSDCContract    int            `json:"supportUSDCContract"`
	SupportUSDCOptions     int            `json:"supportUSDCOptions"`
	USDTPerpetualOpenLine  string         `json:"USDTPerpetualOpenLine"`
	USDCContractOpenLine   string         `json:"USDCContractOpenLine"`
	USDCOptionsOpenLine    string         `json:"USDCOptionsOpenLine"`
	USDTPerpetualCloseLine string         `json:"USDTPerpetualCloseLine"`
	USDCContractCloseLine  string         `json:"USDCContractCloseLine"`
	USDCOptionsCloseLine   string         `json:"USDCOptionsCloseLine"`
	USDCContractSymbols    []string       `json:"USDCContractSymbols"`
	USDCOptionsSymbols     []string       `json:"USDCOptionsSymbols"`
	MarginLeverage         string         `json:"marginLeverage"`
	USDTPerpetualLeverage  []LeverageInfo `json:"USDTPerpetualLeverage"`
	USDCContractLeverage   []LeverageInfo `json:"USDCContractLeverage"`
}

type LeverageInfo struct {
	Symbol   string `json:"symbol"`
	Leverage string `json:"leverage"`
}

type TokenInfo struct {
	Token            string         `json:"token"`
	ConvertRatioList []ConvertRatio `json:"convertRatioList"`
}

type ConvertRatio struct {
	Ladder       string `json:"ladder"`
	ConvertRatio string `json:"convertRatio"`
}

type MarginToken struct {
	ProductId string      `json:"productId"`
	TokenInfo []TokenInfo `json:"tokenInfo"`
}

// LoanInfo To do repay & loand to value
type LoanInfo struct {
	OrderId               string   `json:"orderId"`
	OrderProductId        string   `json:"orderProductId"`
	ParentUid             string   `json:"parentUid"`
	LoanTime              string   `json:"loanTime"`
	LoanCoin              string   `json:"loanCoin"`
	LoanAmount            string   `json:"loanAmount"`
	UnpaidAmount          string   `json:"unpaidAmount"`
	UnpaidInterest        string   `json:"unpaidInterest"`
	RepaidAmount          string   `json:"repaidAmount"`
	RepaidInterest        string   `json:"repaidInterest"`
	InterestRate          string   `json:"interestRate"`
	Status                string   `json:"status"`
	Leverage              string   `json:"leverage"`
	SupportSpot           string   `json:"supportSpot"`
	SupportContract       string   `json:"supportContract"`
	WithdrawLine          string   `json:"withdrawLine"`
	TransferLine          string   `json:"transferLine"`
	SpotBuyLine           string   `json:"spotBuyLine"`
	SpotSellLine          string   `json:"spotSellLine"`
	ContractOpenLine      string   `json:"contractOpenLine"`
	LiquidationLine       string   `json:"liquidationLine"`
	StopLiquidationLine   string   `json:"stopLiquidationLine"`
	ContractLeverage      string   `json:"contractLeverage"`
	TransferRatio         string   `json:"transferRatio"`
	SpotSymbols           []string `json:"spotSymbols"`
	ContractSymbols       []string `json:"contractSymbols"`
	SupportUSDCContract   string   `json:"supportUSDCContract"`
	SupportUSDCOptions    string   `json:"supportUSDCOptions"`
	SupportMarginTrading  string   `json:"supportMarginTrading"`
	USDTPerpetualOpenLine string   `json:"USDTPerpetualOpenLine"`
	USDCContractOpenLine  string   `json:"USDCContractOpenLine"`
}

type DcpSetTimewindowResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AdjustLtvResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AdjustmentHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AdjustmentHistoryItem struct {
	CollateralCurrency string `json:"collateralCurrency"`
	AdjustId           int    `json:"adjustId"`
	AdjustTime         int    `json:"adjustTime"`
	PreLTV             string `json:"preLTV"`
	AfterLTV           string `json:"afterLTV"`
	Direction          int    `json:"direction"`
	Amount             string `json:"amount"`
	Status             int    `json:"status"`
}

type CollateralDataResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CurrencyLiquidationConfig struct {
	Currency         string `json:"currency"`
	LiquidationOrder int    `json:"liquidationOrder"`
}

type CollateralRatioConfig struct {
	Currencies          string                `json:"currencies"`
	CollateralRatioList []CollateralRatioTier `json:"collateralRatioList"`
}

type CollateralRatioTier struct {
	MinValue        string `json:"minValue"`
	MaxValue        string `json:"maxValue"`
	CollateralRatio string `json:"collateralRatio"`
}

type LoanableDataResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LoanableDataItem struct {
	Currency                       string `json:"currency"`
	VipLevel                       string `json:"vipLevel"`
	FlexibleBorrowable             bool   `json:"flexibleBorrowable"`
	FlexibleBorrowingAccuracy      int    `json:"flexibleBorrowingAccuracy"`
	MinFlexibleBorrowingAmount     string `json:"minFlexibleBorrowingAmount"`
	FlexibleAnnualizedInterestRate string `json:"flexibleAnnualizedInterestRate"`
	FixedBorrowable                bool   `json:"fixedBorrowable"`
	FixedBorrowingAccuracy         int    `json:"fixedBorrowingAccuracy"`
	MinFixedBorrowingAmount        string `json:"minFixedBorrowingAmount"`
	MaxBorrowingAmount             string `json:"maxBorrowingAmount"`
	AnnualizedInterestRate7D       string `json:"annualizedInterestRate7D"`
	AnnualizedInterestRate14D      string `json:"annualizedInterestRate14D"`
	AnnualizedInterestRate30D      string `json:"annualizedInterestRate30D"`
	AnnualizedInterestRate60D      string `json:"annualizedInterestRate60D"`
	AnnualizedInterestRate90D      string `json:"annualizedInterestRate90D"`
	AnnualizedInterestRate180D     string `json:"annualizedInterestRate180D"`
}

type MaxCollateralAmountResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type MaxLoanResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type BorrowItem struct {
	LoanCurrency               string `json:"loanCurrency"`
	FlexibleTotalDebt          string `json:"flexibleTotalDebt"`
	FlexibleTotalDebtUSD       string `json:"flexibleTotalDebtUSD"`
	FlexibleHourlyInterestRate string `json:"flexibleHourlyInterestRate"`
	FixedTotalDebt             string `json:"fixedTotalDebt"`
	FixedTotalDebtUSD          string `json:"fixedTotalDebtUSD"`
}

type SupplyItem struct {
	Currency  string `json:"currency"`
	Amount    string `json:"amount"`
	AmountUSD string `json:"amountUSD"`
}

type PostCryptoLoanFixedBorrowOrderCancelResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type PostCryptoLoanFixedSupplyResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type FlexibleBorrowResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PostCryptoLoanFlexibleRepayResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type GetCryptoLoanFlexibleRepaymentHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type FlexibleRepaymentHistoryResult struct {
	List           []FlexibleRepaymentRecord `json:"list"`
	NextPageCursor string                    `json:"nextPageCursor"`
}

type FlexibleRepaymentRecord struct {
	RepayId         string `json:"repayId"`
	LoanCurrency    string `json:"loanCurrency"`
	RepayAmount     string `json:"repayAmount"`
	PrincipalAmount string `json:"principalAmount"`
	InterestAmount  string `json:"interestAmount"`
	RepayTime       int    `json:"repayTime"`
	RepayType       int    `json:"repayType"`
	Status          int    `json:"status"`
}

type MaxLoanRequest struct {
	Currency       string        `json:"currency"`
	CollateralList []interface{} `json:"collateralList"`
}

type CollateralInput struct {
	Ccy    string `json:"ccy"`
	Amount string `json:"amount"`
}

type FullyRepayResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type RenewLoanResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type FlexibleBorrowRequest struct {
	LoanCurrency   string        `json:"loanCurrency"`
	LoanAmount     string        `json:"loanAmount"`
	CollateralList []interface{} `json:"collateralList"`
}

type PostCryptoLoanFlexibleRepayCollateralResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
	Time    int         `json:"time"`
}

type PostCryptoLoanFlexibleRepayCollateralResult struct {
	RepayId string `json:"repayId"`
}

type AssociationUidResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CoinDeltaAmountResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CoinDeltaDetail struct {
	Coin                     string `json:"coin"`
	CoinDeltaSize            string `json:"coinDeltaSize"`
	CoinDeltaAvailableAmount string `json:"coinDeltaAvailableAmount"`
	CoinDeltaAmount          string `json:"coinDeltaAmount"`
}

type EnsureTokensConvertResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type TokenConvertDetail struct {
	ProductId string            `json:"productId"`
	TokenInfo []TokenLadderInfo `json:"tokenInfo"`
}

type TokenLadderInfo struct {
	Token            string               `json:"token"`
	ConvertRatioList []ConvertRatioLadder `json:"convertRatioList"`
}

type ConvertRatioLadder struct {
	Ladder       string `json:"ladder"`
	ConvertRatio string `json:"convertRatio"`
}

type EnsureTokensResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type MarginTokenInfo struct {
	ProductId     string         `json:"productId"`
	SpotToken     []TokenConvert `json:"spotToken"`
	ContractToken []TokenConvert `json:"contractToken"`
}

type TokenConvert struct {
	Token        string `json:"token"`
	ConvertRatio string `json:"convertRatio"`
}

type LoanOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LoanOrderInfo struct {
	OrderId                 string        `json:"orderId"`
	OrderProductId          string        `json:"orderProductId"`
	ParentUid               string        `json:"parentUid"`
	LoanTime                int           `json:"loanTime"`
	LoanCoin                string        `json:"loanCoin"`
	LoanAmount              string        `json:"loanAmount"`
	UnpaidAmount            string        `json:"unpaidAmount"`
	UnpaidInterest          string        `json:"unpaidInterest"`
	RepaidAmount            string        `json:"repaidAmount"`
	RepaidInterest          string        `json:"repaidInterest"`
	InterestRate            string        `json:"interestRate"`
	Status                  string        `json:"status"`
	Leverage                string        `json:"leverage"`
	SupportSpot             string        `json:"supportSpot"`
	SupportContract         string        `json:"supportContract"`
	WithdrawLine            string        `json:"withdrawLine"`
	TransferLine            string        `json:"transferLine"`
	SpotBuyLine             string        `json:"spotBuyLine"`
	SpotSellLine            string        `json:"spotSellLine"`
	ContractOpenLine        string        `json:"contractOpenLine"`
	LiquidationLine         string        `json:"liquidationLine"`
	StopLiquidationLine     string        `json:"stopLiquidationLine"`
	ContractLeverage        string        `json:"contractLeverage"`
	TransferRatio           string        `json:"transferRatio"`
	SpotSymbols             []string      `json:"spotSymbols"`
	ContractSymbols         []string      `json:"contractSymbols"`
	SupportUSDCContract     string        `json:"supportUSDCContract"`
	SupportUSDCOptions      string        `json:"supportUSDCOptions"`
	USDTPerpetualOpenLine   string        `json:"USDTPerpetualOpenLine"`
	USDCContractOpenLine    string        `json:"USDCContractOpenLine"`
	USDCOptionsOpenLine     string        `json:"USDCOptionsOpenLine"`
	USDTPerpetualCloseLine  string        `json:"USDTPerpetualCloseLine"`
	USDCContractCloseLine   string        `json:"USDCContractCloseLine"`
	USDCOptionsCloseLine    string        `json:"USDCOptionsCloseLine"`
	USDCContractSymbols     []string      `json:"USDCContractSymbols"`
	USDCOptionsSymbols      []string      `json:"USDCOptionsSymbols"`
	MarginLeverage          string        `json:"marginLeverage"`
	USDTPerpetualLeverage   []interface{} `json:"USDTPerpetualLeverage"`
	USDCContractLeverage    []interface{} `json:"USDCContractLeverage"`
	SupportMarginTrading    string        `json:"supportMarginTrading"`
	DeferredLiquidationLine string        `json:"deferredLiquidationLine"`
	DeferredLiquidationTime string        `json:"deferredLiquidationTime"`
	ReserveToken            string        `json:"reserveToken"`
	ReserveQuantity         string        `json:"reserveQuantity"`
}

type SymbolLeverage struct {
	Symbol   string `json:"symbol"`
	Leverage string `json:"leverage"`
}

type LtvConvertResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LtvConvertInfo struct {
	Ltv             string               `json:"ltv"`
	ParentUid       string               `json:"parentUid"`
	SubAccountUids  []int                `json:"subAccountUids"`
	UnpaidAmount    string               `json:"unpaidAmount"`
	UnpaidInfo      []UnpaidInfo         `json:"unpaidInfo"`
	Balance         string               `json:"balance"`
	BalanceInfo     []BalanceConvertInfo `json:"balanceInfo"`
	Rst             string               `json:"rst"`
	LiqStatus       int                  `json:"liqStatus"`
}

type UnpaidInfo struct {
	Token          string `json:"token"`
	UnpaidQty      string `json:"unpaidQty"`
	UnpaidInterest string `json:"unpaidInterest"`
}

type BalanceConvertInfo struct {
	Token           string `json:"token"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	ConvertedAmount string `json:"convertedAmount"`
}

type ProductInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type ProductInfo struct {
	ProductId               string           `json:"productId"`
	Leverage                string           `json:"leverage"`
	SupportSpot             int              `json:"supportSpot"`
	SupportContract         int              `json:"supportContract"`
	WithdrawLine            string           `json:"withdrawLine"`
	TransferLine            string           `json:"transferLine"`
	SpotBuyLine             string           `json:"spotBuyLine"`
	SpotSellLine            string           `json:"spotSellLine"`
	ContractOpenLine        string           `json:"contractOpenLine"`
	LiquidationLine         string           `json:"liquidationLine"`
	StopLiquidationLine     string           `json:"stopLiquidationLine"`
	ContractLeverage        string           `json:"contractLeverage"`
	TransferRatio           string           `json:"transferRatio"`
	SpotSymbols             []string         `json:"spotSymbols"`
	ContractSymbols         []string         `json:"contractSymbols"`
	SupportUSDCContract     int              `json:"supportUSDCContract"`
	SupportUSDCOptions      int              `json:"supportUSDCOptions"`
	USDTPerpetualOpenLine   string           `json:"USDTPerpetualOpenLine"`
	USDCContractOpenLine    string           `json:"USDCContractOpenLine"`
	USDCOptionsOpenLine     string           `json:"USDCOptionsOpenLine"`
	USDTPerpetualCloseLine  string           `json:"USDTPerpetualCloseLine"`
	USDCContractCloseLine   string           `json:"USDCContractCloseLine"`
	USDCOptionsCloseLine    string           `json:"USDCOptionsCloseLine"`
	USDCContractSymbols     []string         `json:"USDCContractSymbols"`
	USDCOptionsSymbols      []string         `json:"USDCOptionsSymbols"`
	MarginLeverage          string           `json:"marginLeverage"`
	USDTPerpetualLeverage   []SymbolLeverage `json:"USDTPerpetualLeverage"`
	USDCContractLeverage    []SymbolLeverage `json:"USDCContractLeverage"`
	SupportMarginTrading    int              `json:"supportMarginTrading"`
	DeferredLiquidationLine string           `json:"deferredLiquidationLine"`
	DeferredLiquidationTime string           `json:"deferredLiquidationTime"`
	ProductType             string           `json:"productType"`
}

type RepaidHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type RepayInfo struct {
	RepayOrderId string `json:"repayOrderId"`
	RepaidTime   int    `json:"repaidTime"`
	Token        string `json:"token"`
	Quantity     string `json:"quantity"`
	Interest     string `json:"interest"`
	BusinessType string `json:"businessType"`
	Status       string `json:"status"`
}

type RepayLoanResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetAdvanceProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DualAssetsProduct struct {
	Category                    string  `json:"category"`
	ProductId                   string  `json:"productId"`
	BaseCoin                    string  `json:"baseCoin"`
	QuoteCoin                   string  `json:"quoteCoin"`
	ExpectReceiveAt             string  `json:"expectReceiveAt"`
	Duration                    string  `json:"duration"`
	Status                      string  `json:"status"`
	IsVipProduct                bool    `json:"isVipProduct"`
	SubscribeStartAt            string  `json:"subscribeStartAt"`
	SubscribeEndAt              string  `json:"subscribeEndAt"`
	ApplyStartAt                string  `json:"applyStartAt"`
	SettlementTime              string  `json:"settlementTime"`
	MinPurchaseQuoteAmount      string  `json:"minPurchaseQuoteAmount"`
	MinPurchaseBaseAmount       string  `json:"minPurchaseBaseAmount"`
	RemainingAmountQuote        string  `json:"remainingAmountQuote"`
	RemainingAmountBase         string  `json:"remainingAmountBase"`
	OrderPrecisionDigitalQuote  int     `json:"orderPrecisionDigitalQuote"`
	OrderPrecisionDigitalBase   int     `json:"orderPrecisionDigitalBase"`
}

type GetProductExtraInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type ProductOffer struct {
	ProductId     string        `json:"productId"`
	CurrentPrice  string        `json:"currentPrice"`
	BuyLowPrice   []interface{} `json:"buyLowPrice"`
	SellHighPrice []interface{} `json:"sellHighPrice"`
}

type PlaceAdvanceOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type GetAdvancePositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DualAssetsPosition struct {
	PositionId       string `json:"positionId"`
	ProductId        string `json:"productId"`
	Category         string `json:"category"`
	BaseCoin         string `json:"baseCoin"`
	QuoteCoin        string `json:"quoteCoin"`
	InvestCoin       string `json:"investCoin"`
	Amount           string `json:"amount"`
	ApyE8            string `json:"apyE8"`
	Direction        string `json:"direction"`
	TargetPrice      string `json:"targetPrice"`
	SettlementTime   string `json:"settlementTime"`
	Status           string `json:"status"`
	OrderId          string `json:"orderId"`
	Duration         string `json:"duration"`
	ExpectReturnCoin string `json:"expectReturnCoin"`
	ExpectReturnAmount string `json:"expectReturnAmount"`
	AccountType      string `json:"accountType"`
	ToAccountType    string `json:"toAccountType"`
	YieldStartAt     string `json:"yieldStartAt"`
	YieldEndAt       string `json:"yieldEndAt"`
}

type GetAdvanceOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DualAssetsOrder struct {
	OrderId          string `json:"orderId"`
	OrderLinkId      string `json:"orderLinkId"`
	ProductId        string `json:"productId"`
	Category         string `json:"category"`
	OrderType        string `json:"orderType"`
	Amount           string `json:"amount"`
	Coin             string `json:"coin"`
	BaseCoin         string `json:"baseCoin"`
	QuoteCoin        string `json:"quoteCoin"`
	Status           string `json:"status"`
	CreatedTime      string `json:"createdTime"`
	UpdatedTime      string `json:"updatedTime"`
	Direction        string `json:"direction"`
	TargetPrice      string `json:"targetPrice"`
	SettlementTime   string `json:"settlementTime"`
	EstimateApyE8    string `json:"estimateApyE8"`
	Duration         string `json:"duration"`
	AccountType      string `json:"accountType"`
	ToAccountType    string `json:"toAccountType"`
	SelectApyE8      string `json:"selectApyE8"`
	IsVip            bool   `json:"isVip"`
	SettlementCoin   string `json:"settlementCoin"`
	SettlementAmount string `json:"settlementAmount"`
	OrderMode        string `json:"orderMode"`
	SettlementPrice  string `json:"settlementPrice"`
	RefundStatus     string `json:"refundStatus"`
	TrialBonusAmount string `json:"trialBonusAmount"`
	TrialBonusPnl    string `json:"trialBonusPnl"`
}

type GetRedeemEstAmountListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RedeemEstItem struct {
	Success          bool   `json:"success"`
	PositionId       string `json:"positionId"`
	EstRedeemAmount  string `json:"estRedeemAmount"`
	EstRedeemTime    string `json:"estRedeemTime"`
	SlippageRate     string `json:"slippageRate"`
}

type GetDoubleWinLeverageResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type GetSmartLeverageProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SmartLeverageProduct struct {
	Category              string `json:"category"`
	ProductId             string `json:"productId"`
	InvestCoin            string `json:"investCoin"`
	UnderlyingAsset       string `json:"underlyingAsset"`
	Direction             string `json:"direction"`
	Leverage              string `json:"leverage"`
	Duration              string `json:"duration"`
	ExpectReceiveAt       string `json:"expectReceiveAt"`
	SubscribeStartAt      string `json:"subscribeStartAt"`
	SubscribeEndAt        string `json:"subscribeEndAt"`
	SettlementTime        string `json:"settlementTime"`
	MinPurchaseAmount     string `json:"minPurchaseAmount"`
	RemainingAmount       string `json:"remainingAmount"`
	OrderPrecisionDigital int    `json:"orderPrecisionDigital"`
}

type GetSmartLeverageProductExtraInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SmartLeverageProductExtraInfo struct {
	Category            string `json:"category"`
	ProductId           string `json:"productId"`
	BreakevenPrice      string `json:"breakevenPrice"`
	CurrentPrice        string `json:"currentPrice"`
	ExpireAt            string `json:"expireAt"`
	MaxInvestmentAmount string `json:"maxInvestmentAmount"`
}

type GetSmartLeveragePositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SmartLeveragePosition struct {
	PositionId      string `json:"positionId"`
	ProductId       string `json:"productId"`
	Category        string `json:"category"`
	InvestCoin      string `json:"investCoin"`
	UnderlyingAsset string `json:"underlyingAsset"`
	Direction       string `json:"direction"`
	Leverage        string `json:"leverage"`
	Amount          string `json:"amount"`
	BreakevenPrice  string `json:"breakevenPrice"`
	InitialPrice    string `json:"initialPrice"`
	Duration        string `json:"duration"`
	SettlementTime  string `json:"settlementTime"`
	CreatedTime     string `json:"createdTime"`
	Status          string `json:"status"`
	Redeemable      bool   `json:"redeemable"`
	AccountType     string `json:"accountType"`
	OrderLinkId     string `json:"orderLinkId"`
	OrderId         string `json:"orderId"`
}

type GetSmartLeverageOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SmartLeverageOrder struct {
	OrderId          string `json:"orderId"`
	OrderLinkId      string `json:"orderLinkId"`
	ProductId        string `json:"productId"`
	Category         string `json:"category"`
	OrderType        string `json:"orderType"`
	InvestCoin       string `json:"investCoin"`
	Amount           string `json:"amount"`
	UnderlyingAsset  string `json:"underlyingAsset"`
	Direction        string `json:"direction"`
	Leverage         string `json:"leverage"`
	BreakevenPrice   string `json:"breakevenPrice"`
	InitialPrice     string `json:"initialPrice"`
	SettlementTime   string `json:"settlementTime"`
	Duration         string `json:"duration"`
	CreatedTime      string `json:"createdTime"`
	Status           string `json:"status"`
	SettlementPrice  string `json:"settlementPrice"`
	Pnl              string `json:"pnl"`
	RefundStatus     string `json:"refundStatus"`
	AccountType      string `json:"accountType"`
	ToAccountType    string `json:"toAccountType"`
}

type GetDoubleWinProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DoubleWinProduct struct {
	Category              string `json:"category"`
	ProductId             string `json:"productId"`
	InvestCoin            string `json:"investCoin"`
	UnderlyingAsset       string `json:"underlyingAsset"`
	Duration              string `json:"duration"`
	ExpectReceiveAt       string `json:"expectReceiveAt"`
	SubscribeStartAt      string `json:"subscribeStartAt"`
	SubscribeEndAt        string `json:"subscribeEndAt"`
	SettlementTime        string `json:"settlementTime"`
	MinPurchaseAmount     string `json:"minPurchaseAmount"`
	OrderPrecisionDigital int    `json:"orderPrecisionDigital"`
	IsRfqProduct          bool   `json:"isRfqProduct"`
	LowerPriceBuffer      string `json:"lowerPriceBuffer"`
	UpperPriceBuffer      string `json:"upperPriceBuffer"`
	MinDeviationRatio     string `json:"minDeviationRatio"`
	MaxDeviationRatio     string `json:"maxDeviationRatio"`
	PriceTickSize         string `json:"priceTickSize"`
}

type GetDoubleWinProductExtraInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DoubleWinProductExtraInfo struct {
	Category            string `json:"category"`
	ProductId           string `json:"productId"`
	Leverage            string `json:"leverage"`
	CurrentPrice        string `json:"currentPrice"`
	ExpireTime          string `json:"expireTime"`
	MaxInvestmentAmount string `json:"maxInvestmentAmount"`
}

type GetDoubleWinPositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DoubleWinPosition struct {
	PositionId      string `json:"positionId"`
	ProductId       string `json:"productId"`
	Category        string `json:"category"`
	InvestCoin      string `json:"investCoin"`
	UnderlyingAsset string `json:"underlyingAsset"`
	Amount          string `json:"amount"`
	Leverage        string `json:"leverage"`
	InitialPrice    string `json:"initialPrice"`
	LowerPrice      string `json:"lowerPrice"`
	UpperPrice      string `json:"upperPrice"`
	Duration        string `json:"duration"`
	SettlementTime  string `json:"settlementTime"`
	CreatedTime     string `json:"createdTime"`
	Status          string `json:"status"`
	Redeemable      bool   `json:"redeemable"`
	AccountType     string `json:"accountType"`
	OrderId         string `json:"orderId"`
}

type GetDoubleWinOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DoubleWinOrder struct {
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	ProductId       string `json:"productId"`
	Category        string `json:"category"`
	OrderType       string `json:"orderType"`
	InvestCoin      string `json:"investCoin"`
	Amount          string `json:"amount"`
	UnderlyingAsset string `json:"underlyingAsset"`
	InitialPrice    string `json:"initialPrice"`
	LowerPrice      string `json:"lowerPrice"`
	UpperPrice      string `json:"upperPrice"`
	Leverage        string `json:"leverage"`
	SettlementTime  string `json:"settlementTime"`
	Duration        string `json:"duration"`
	CreatedTime     string `json:"createdTime"`
	UpdatedTime     string `json:"updatedTime"`
	Status          string `json:"status"`
	SettlementPrice string `json:"settlementPrice"`
	Pnl             string `json:"pnl"`
	RefundStatus    string `json:"refundStatus"`
	AccountType     string `json:"accountType"`
	ToAccountType   string `json:"toAccountType"`
}

type GetDiscountBuyProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type AdvDiscountBuyProduct struct {
	Category              string `json:"category"`
	ProductId             int    `json:"productId"`
	Coin                  string `json:"coin"`
	UnderlyingAsset       string `json:"underlyingAsset"`
	SettlementTime        int    `json:"settlementTime"`
	Duration              string `json:"duration"`
	IsVipProduct          bool   `json:"isVipProduct"`
	SubscribeStartAt      int    `json:"subscribeStartAt"`
	SubscribeEndAt        int    `json:"subscribeEndAt"`
	MinPurchaseAmount     string `json:"minPurchaseAmount"`
	RemainingAmount       string `json:"remainingAmount"`
	OrderPrecisionDigital int    `json:"orderPrecisionDigital"`
	ExpectReceiveAt       int    `json:"expectReceiveAt"`
}

type GetDiscountBuyProductExtraInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DiscountBuyPriceOfferItem struct {
	Category            string `json:"category"`
	ProductId           int    `json:"productId"`
	CurrentPrice        string `json:"currentPrice"`
	PurchasePrice       string `json:"purchasePrice"`
	KnockoutPrice       string `json:"knockoutPrice"`
	KnockoutCouponE8    int    `json:"knockoutCouponE8"`
	MaxInvestmentAmount string `json:"maxInvestmentAmount"`
	InstUid             int    `json:"instUid"`
	ExpiredAt           int    `json:"expiredAt"`
}

type GetDiscountBuyPositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type AdvDiscountBuyPosition struct {
	PositionId       string `json:"positionId"`
	ProductId        int    `json:"productId"`
	Category         string `json:"category"`
	Coin             string `json:"coin"`
	UnderlyingAsset  string `json:"underlyingAsset"`
	Amount           string `json:"amount"`
	PurchasePrice    string `json:"purchasePrice"`
	KnockoutPrice    string `json:"knockoutPrice"`
	KnockoutCouponE8 int    `json:"knockoutCouponE8"`
	Status           string `json:"status"`
	OrderId          string `json:"orderId"`
	Duration         string `json:"duration"`
	SettlementTime   int    `json:"settlementTime"`
	AccountType      string `json:"accountType"`
	ToAccountType    string `json:"toAccountType"`
	SettleType       string `json:"settleType"`
	ExpectReceiveAt  int    `json:"expectReceiveAt"`
}

type GetDiscountBuyOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type AdvDiscountBuyOrder struct {
	OrderId          string `json:"orderId"`
	OrderLinkId      string `json:"orderLinkId"`
	ProductId        int    `json:"productId"`
	Category         string `json:"category"`
	OrderType        string `json:"orderType"`
	Amount           string `json:"amount"`
	Coin             string `json:"coin"`
	UnderlyingAsset  string `json:"underlyingAsset"`
	Status           string `json:"status"`
	CreatedTime      string `json:"createdTime"`
	PurchasePrice    string `json:"purchasePrice"`
	KnockoutPrice    string `json:"knockoutPrice"`
	KnockoutCouponE8 int    `json:"knockoutCouponE8"`
	Duration         string `json:"duration"`
	SettlementTime   int    `json:"settlementTime"`
	AccountType      string `json:"accountType"`
	ToAccountType    string `json:"toAccountType"`
	SettlementPrice  string `json:"settlementPrice"`
	SettlementCoin   string `json:"settlementCoin"`
	SettlementAmount string `json:"settlementAmount"`
	SettleType       string `json:"settleType"`
	IsVip            bool   `json:"isVip"`
	RefundStatus     string `json:"refundStatus"`
}

type PlaceTokenOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TokenOrderListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TokenOrder struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
	OrderType   string `json:"orderType"`
	FromCoin    string `json:"fromCoin"`
	ToCoin      string `json:"toCoin"`
	FromAmount  string `json:"fromAmount"`
	ToAmount    string `json:"toAmount"`
	ServiceFee  string `json:"serviceFee"`
	Status      string `json:"status"`
	CreatedTime int    `json:"createdTime"`
}

type TokenProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TokenProduct struct {
	ProductId        int    `json:"productId"`
	Coin             string `json:"coin"`
	MintFeeRateE8    int    `json:"mintFeeRateE8"`
	RedeemFeeRateE8  int    `json:"redeemFeeRateE8"`
	MinInvestment    string `json:"minInvestment"`
	UserHolding      string `json:"userHolding"`
	LeftQuota        string `json:"leftQuota"`
	CanMint          bool   `json:"canMint"`
	SavingsBalance   string `json:"savingsBalance"`
	AprE8            int    `json:"aprE8"`
	BonusAprE8       int    `json:"bonusAprE8"`
	BonusMaxAmount   string `json:"bonusMaxAmount"`
	BaseCoinPrecision int   `json:"baseCoinPrecision"`
	TokenPrecision   int    `json:"tokenPrecision"`
}

type TokenPositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TokenPosition struct {
	TotalAmount    string `json:"totalAmount"`
	TotalYield     string `json:"totalYield"`
	YesterdayYield string `json:"yesterdayYield"`
	AprE8          int    `json:"aprE8"`
	BonusAprE8     int    `json:"bonusAprE8"`
	BonusMaxAmount string `json:"bonusMaxAmount"`
	HasQuota       bool   `json:"hasQuota"`
}

type TokenDailyYieldResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DailyYieldRecord struct {
	Yield       string `json:"yield"`
	BonusYield  string `json:"bonusYield"`
	Status      string `json:"status"`
	CreatedTime int    `json:"createdTime"`
}

type TokenHourlyYieldResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type HourlyYieldRecord struct {
	EffectiveAmount string `json:"effectiveAmount"`
	Yield           string `json:"yield"`
	RewardType      int    `json:"rewardType"`
	AprE8           int    `json:"aprE8"`
	HourlyDate      int    `json:"hourlyDate"`
	CreatedTime     int    `json:"createdTime"`
}

type TokenAprHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type AprRecord struct {
	Timestamp int `json:"timestamp"`
	AprE8     int `json:"apyE8"`
}

type GetFixedTermProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PlaceFixedTermOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RedeemFixedTermResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type GetFixedTermPositionResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type GetFixedTermOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type FixedTermProduct struct {
	ProductId                string        `json:"productId"`
	Category                 string        `json:"category"`
	Coin                     string        `json:"coin"`
	Duration                 string        `json:"duration"`
	Status                   string        `json:"status"`
	TieredApyList            []interface{} `json:"tieredApyList"`
	MinStakeAmount           string        `json:"minStakeAmount"`
	MaxStakeAmount           string        `json:"maxStakeAmount"`
	Precision                int           `json:"precision"`
	SubscribeStartAt         int           `json:"subscribeStartAt"`
	SubscribeEndAt           int           `json:"subscribeEndAt"`
	AllowEarlyRedemption     bool          `json:"allowEarlyRedemption"`
	EarlyRedemptionApy       string        `json:"earlyRedemptionApy"`
	RedemptionLimitDuration  string        `json:"redemptionLimitDuration"`
	AllowAutoReinvest        bool          `json:"allowAutoReinvest"`
	InterestCoinApyList      []interface{} `json:"interestCoinApyList"`
	IsVip                    bool          `json:"isVip"`
	CreditTime               int           `json:"creditTime"`
	SpecialUserGroupRequired bool          `json:"specialUserGroupRequired"`
	SpecialUserGroupInfo     string        `json:"specialUserGroupInfo"`
}

type FixedTermPosition struct {
	PositionId          string        `json:"positionId"`
	ProductId           string        `json:"productId"`
	Category            string        `json:"category"`
	Coin                string        `json:"coin"`
	Amount              string        `json:"amount"`
	EffectiveAmount     string        `json:"effectiveAmount"`
	Duration            string        `json:"duration"`
	Status              string        `json:"status"`
	SettlementTime      int           `json:"settlementTime"`
	CreatedAt           int           `json:"createdAt"`
	OrderId             string        `json:"orderId"`
	EarlyRedeemInfo     interface{}   `json:"earlyRedeemInfo"`
	AllowAutoReinvest   bool          `json:"allowAutoReinvest"`
	AutoReinvest        string        `json:"autoReinvest"`
	InterestCoinApyList []interface{} `json:"interestCoinApyList"`
}

type FixedTermOrder struct {
	OrderId        string        `json:"orderId"`
	OrderLinkId    string        `json:"orderLinkId"`
	OrderType      string        `json:"orderType"`
	Status         string        `json:"status"`
	ProductId      string        `json:"productId"`
	Category       string        `json:"category"`
	Coin           string        `json:"coin"`
	Amount         string        `json:"amount"`
	Duration       string        `json:"duration"`
	AccountType    string        `json:"accountType"`
	SettlementTime int           `json:"settlementTime"`
	CreatedAt      int           `json:"createdAt"`
	YieldInfoList  []interface{} `json:"yieldInfoList"`
}

type TieredApy struct {
	Min string `json:"min"`
	Max string `json:"max"`
	Apy string `json:"apy"`
}

type InterestCoinApy struct {
	Coin               string `json:"coin"`
	Apy                string `json:"apy"`
	ExpectUnitEarning  string `json:"expectUnitEarning"`
	CurrentPrice       string `json:"currentPrice"`
}

type PositionInterestCoinApy struct {
	Coin                string `json:"coin"`
	Apy                 string `json:"apy"`
	ExpectReturnEarning string `json:"expectReturnEarning"`
	Price               string `json:"price"`
}

type YieldInfo struct {
	Coin      string `json:"coin"`
	Amount    string `json:"amount"`
	Status    string `json:"status"`
	CreatedAt int    `json:"createdAt"`
	Apy       string `json:"apy"`
}

type EarlyRedeemInfo struct {
	AllowEarlyRedeem       bool   `json:"allowEarlyRedeem"`
	EarlyRedeemEarning     string `json:"earlyRedeemEarning"`
	ReturnCoin             string `json:"returnCoin"`
	RedemptionLimitDuration string `json:"redemptionLimitDuration"`
}

type PlaceFixedTermOrderRequest struct {
	ProductId   string `json:"productId"`
	Category    string `json:"category"`
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	AccountType string `json:"accountType"`
	OrderLinkId string `json:"orderLinkId"`
	AutoInvest  bool   `json:"autoInvest"`
}

type RedeemFixedTermRequest struct {
	ProductId  string `json:"productId"`
	Category   string `json:"category"`
	PositionId string `json:"positionId"`
}

type SetAutoInvestRequest struct {
	ProductId  string `json:"productId"`
	Category   string `json:"category"`
	PositionId string `json:"positionId"`
	Status     string `json:"status"`
}

type PaginatedList struct {
	NextPageCursor string `json:"nextPageCursor"`
}

type RwaProductListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RwaProduct struct {
	ProductId        int    `json:"productId"`
	Coin             string `json:"coin"`
	AssetSymbol      string `json:"assetSymbol"`
	Manager          string `json:"manager"`
	BaseApr          string `json:"baseApr"`
	BonusApr         string `json:"bonusApr"`
	SavingType       string `json:"savingType"`
	Duration         int    `json:"duration"`
	Nav              string `json:"nav"`
	MinStakeAmount   string `json:"minStakeAmount"`
	MaxStakeAmount   string `json:"maxStakeAmount"`
	UserMaxAmount    string `json:"userMaxAmount"`
	UserQuota        string `json:"userQuota"`
	MinRedeemShare   string `json:"minRedeemShare"`
	RedeemFeeRate    string `json:"redeemFeeRate"`
	SubscriptionFee  string `json:"subscriptionFee"`
	ExtLink          string `json:"extLink"`
	AmountPrecision  int    `json:"amountPrecision"`
	SharePrecision   int    `json:"sharePrecision"`
}

type PlaceRwaOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RwaPositionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RwaPosition struct {
	ProductId              int    `json:"productId"`
	Coin                   string `json:"coin"`
	AssetSymbol            string `json:"assetSymbol"`
	EffectiveShare         string `json:"effectiveShare"`
	ProcessingStakeAmount  string `json:"processingStakeAmount"`
	ProcessingRedeemShare  string `json:"processingRedeemShare"`
	BonusEarned            string `json:"bonusEarned"`
	Nav                    string `json:"nav"`
	HoldAmount             string `json:"holdAmount"`
	Duration               int    `json:"duration"`
}

type RwaOrderListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RwaOrder struct {
	OrderId        string `json:"orderId"`
	OrderLinkId    string `json:"orderLinkId"`
	OrderType      string `json:"orderType"`
	ProductId      int    `json:"productId"`
	Coin           string `json:"coin"`
	StakeAmount    string `json:"stakeAmount"`
	RedeemShares   string `json:"redeemShares"`
	Status         string `json:"status"`
	AccountType    string `json:"accountType"`
	CreatedTime    int    `json:"createdTime"`
	UpdatedTime    int    `json:"updatedTime"`
	SettledShares  string `json:"settledShares"`
	SettledAmount  string `json:"settledAmount"`
}

type RwaNavChartResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RwaNavPoint struct {
	Date string `json:"date"`
	Nav  string `json:"nav"`
}

type EarnProduct struct {
	Category                   string        `json:"category"`
	EstimateApr                string        `json:"estimateApr"`
	Coin                       string        `json:"coin"`
	MinStakeAmount             string        `json:"minStakeAmount"`
	MaxStakeAmount             string        `json:"maxStakeAmount"`
	Precision                  string        `json:"precision"`
	ProductId                  string        `json:"productId"`
	Status                     string        `json:"status"`
	Duration                   string        `json:"duration"`
	Term                       int           `json:"term"`
	SwapCoin                   string        `json:"swapCoin"`
	SwapCoinPrecision          string        `json:"swapCoinPrecision"`
	StakeExchangeRate          string        `json:"stakeExchangeRate"`
	RedeemExchangeRate         string        `json:"redeemExchangeRate"`
	MinRedeemAmount            string        `json:"minRedeemAmount"`
	MaxRedeemAmount            string        `json:"maxRedeemAmount"`
	RedeemProcessingMinute     string        `json:"redeemProcessingMinute"`
	StakeTime                  string        `json:"stakeTime"`
	InterestCalculationTime    string        `json:"interestCalculationTime"`
	RewardDistributionType     string        `json:"rewardDistributionType"`
	RewardIntervalMinute       int           `json:"rewardIntervalMinute"`
	BonusEvents                []interface{} `json:"bonusEvents"`
}

type GetProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PlaceOrderRequest struct {
	Category         string      `json:"category"`
	OrderType        string      `json:"orderType"`
	AccountType      string      `json:"accountType"`
	Amount           string      `json:"amount"`
	Coin             string      `json:"coin"`
	ProductId        string      `json:"productId"`
	OrderLinkId      string      `json:"orderLinkId"`
	RedeemPositionId string      `json:"redeemPositionId"`
	ToAccountType    string      `json:"toAccountType"`
	InterestCard     interface{} `json:"interestCard"`
}

type PlaceOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type EarnOrder struct {
	Coin                 string `json:"coin"`
	OrderValue           string `json:"orderValue"`
	OrderType            string `json:"orderType"`
	OrderId              string `json:"orderId"`
	OrderLinkId          string `json:"orderLinkId"`
	Status               string `json:"status"`
	CreatedAt            string `json:"createdAt"`
	UpdatedAt            string `json:"updatedAt"`
	ProductId            string `json:"productId"`
	SwapOrderValue       string `json:"swapOrderValue"`
	EstimateRedeemTime   string `json:"estimateRedeemTime"`
	EstimateStakeTime    string `json:"estimateStakeTime"`
}

type OrderHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type EarnPosition struct {
	Coin                              string `json:"coin"`
	ProductId                         string `json:"productId"`
	Amount                            string `json:"amount"`
	TotalPnl                          string `json:"totalPnl"`
	ClaimableYield                    string `json:"claimableYield"`
	Id                                string `json:"id"`
	Status                            string `json:"status"`
	OrderId                           string `json:"orderId"`
	EstimateRedeemTime                string `json:"estimateRedeemTime"`
	EstimateStakeTime                 string `json:"estimateStakeTime"`
	EstimateInterestCalculationTime   string `json:"estimateInterestCalculationTime"`
	SettlementTime                    string `json:"settlementTime"`
	AutoReinvest                      string `json:"autoReinvest"`
}

type YieldRecord struct {
	ProductId              string `json:"productId"`
	Coin                   string `json:"coin"`
	Id                     string `json:"id"`
	Amount                 string `json:"amount"`
	YieldType              string `json:"yieldType"`
	DistributionMode       string `json:"distributionMode"`
	EffectiveStakingAmount string `json:"effectiveStakingAmount"`
	OrderId                string `json:"orderId"`
	Status                 string `json:"status"`
	CreatedAt              string `json:"createdAt"`
}

type YieldHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type HourlyYieldResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type ModifyEarnPositionRequest struct {
	Category     string `json:"category"`
	ProductId    int    `json:"productId"`
	PositionId   int    `json:"positionId"`
	AutoReinvest int    `json:"autoReinvest"`
}

type AprHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type InterestCard struct {
	AwardId                 int    `json:"awardId"`
	SpecCode                string `json:"specCode"`
	Coin                    string `json:"coin"`
	Apy                     string `json:"apy"`
	Duration                int    `json:"duration"`
	ClaimedAt               int    `json:"claimedAt"`
	ExpireAt                int    `json:"expireAt"`
	UsedAt                  int    `json:"usedAt"`
	Status                  string `json:"status"`
	CurrentPnl              string `json:"currentPnl"`
	LimitPnl                string `json:"limitPnl"`
	PositionEffectiveAmount string `json:"positionEffectiveAmount"`
	ProductId               int    `json:"productId"`
	Category                string `json:"category"`
}

type AwardCard struct {
	AwardId             int    `json:"awardId"`
	SpecCode            string `json:"specCode"`
	ClaimedAt           int    `json:"claimedAt"`
	UsedAt              int    `json:"usedAt"`
	ExpireAt            int    `json:"expireAt"`
	Status              string `json:"status"`
	Amount              string `json:"amount"`
	LimitPnlPercentage  string `json:"limitPnlPercentage"`
	BaseCoin            string `json:"baseCoin"`
	QuoteCoin           string `json:"quoteCoin"`
	Direction           int    `json:"direction"`
	Category            string `json:"category"`
}

type ListCouponsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type InterestCardRef struct {
	AwardId  int    `json:"awardId"`
	SpecCode string `json:"specCode"`
}

type GetHoldToEarnProductResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type GetHoldToEarnYieldHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type HoldToEarnProduct struct {
	CoinName        string        `json:"coinName"`
	Yields          []interface{} `json:"yields"`
	Status          string        `json:"status"`
	Apy             string        `json:"apy"`
	AnnouncementUrl string        `json:"announcementUrl"`
}

type HoldToEarnYieldCoin struct {
	CoinName string `json:"coinName"`
	Apy      string `json:"apy"`
}

type AirdropDailyPnl struct {
	CoinName        string `json:"coinName"`
	YieldCoinName   string `json:"yieldCoinName"`
	EffectiveAmount string `json:"effectiveAmount"`
	Pnl             string `json:"pnl"`
	Apy             string `json:"apy"`
	CreatedAt       int    `json:"createdAt"`
}

type PaginatedResult struct {
	NextPageCursor string `json:"nextPageCursor"`
}

type CoinApy struct {
	Coin   int  `json:"coin"`
	ApyE8  int  `json:"apy_e8"`
	Reward bool `json:"reward"`
}

type LMProduct struct {
	ProductId              string        `json:"productId"`
	BaseCoin               string        `json:"baseCoin"`
	QuoteCoin              string        `json:"quoteCoin"`
	Status                 string        `json:"status"`
	MaxLeverage            int           `json:"maxLeverage"`
	MinInvestmentQuote     string        `json:"minInvestmentQuote"`
	MinInvestmentBase      string        `json:"minInvestmentBase"`
	MaxInvestmentQuote     string        `json:"maxInvestmentQuote"`
	MaxInvestmentBase      string        `json:"maxInvestmentBase"`
	MinWithdrawalAmount    string        `json:"minWithdrawalAmount"`
	BaseCoinPrecision      int           `json:"baseCoinPrecision"`
	QuoteCoinPrecision     int           `json:"quoteCoinPrecision"`
	MinReinvestAmount      string        `json:"minReinvestAmount"`
	YieldCoins             []string      `json:"yieldCoins"`
	ApyE8                  int           `json:"apyE8"`
	Apy7dE8                int           `json:"apy7dE8"`
	PoolLiquidityValue     string        `json:"poolLiquidityValue"`
	DailyYield             string        `json:"dailyYield"`
	SlippageLevels         []string      `json:"slippageLevels"`
	SlippageRateE8List     []int         `json:"slippageRateE8List"`
	ApyBreakdown           []interface{} `json:"apyBreakdown"`
	Apy7dBreakdown         []interface{} `json:"apy7dBreakdown"`
}

type GetLMProductsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type LMOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type LMPosition struct {
	PositionId              string `json:"positionId"`
	ProductId               string `json:"productId"`
	BaseCoin                string `json:"baseCoin"`
	QuoteCoin               string `json:"quoteCoin"`
	QuoteAmount             string `json:"quoteAmount"`
	BaseAmount              string `json:"baseAmount"`
	PrincipalQuoteAmount    string `json:"principalQuoteAmount"`
	PrincipalBaseAmount     string `json:"principalBaseAmount"`
	PrincipalLiquidityValue string `json:"principalLiquidityValue"`
	LeveragedValue          string `json:"leveragedValue"`
	Loan                    string `json:"loan"`
	ClaimableYield          string `json:"claimableYield"`
	CurrentApr              string `json:"currentApr"`
	Leverage                string `json:"leverage"`
	Margin                  string `json:"margin"`
	LiquidationPrice        string `json:"liquidationPrice"`
	CurrentPriceY           string `json:"currentPriceY"`
	Status                  string `json:"status"`
	CreatedTime             string `json:"createdTime"`
}

type GetLMPositionsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type LMOrder struct {
	OrderId      string `json:"orderId"`
	OrderLinkId  string `json:"orderLinkId"`
	ProductId    string `json:"productId"`
	OrderType    string `json:"orderType"`
	BaseCoin     string `json:"baseCoin"`
	QuoteCoin    string `json:"quoteCoin"`
	QuoteAmount  string `json:"quoteAmount"`
	BaseAmount   string `json:"baseAmount"`
	Status       string `json:"status"`
	RemoveType   string `json:"removeType"`
	RemoveRate   int    `json:"removeRate"`
	Leverage     string `json:"leverage"`
	SlippageValue string `json:"slippageValue"`
	CreatedTime  string `json:"createdTime"`
}

type GetLMOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type LMYieldRecord struct {
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	BaseCoin    string `json:"baseCoin"`
	QuoteCoin   string `json:"quoteCoin"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	CreatedTime string `json:"createdTime"`
}

type GetLMYieldRecordsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type LMLiquidationRecord struct {
	BaseCoin         string `json:"baseCoin"`
	QuoteCoin        string `json:"quoteCoin"`
	BaseAmount       string `json:"baseAmount"`
	QuoteAmount      string `json:"quoteAmount"`
	LiquidationPrice string `json:"liquidationPrice"`
	LiquidationTime  string `json:"liquidationTime"`
}

type GetLMLiquidationRecordsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstListFundsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstFundInfo struct {
	FundId            string   `json:"fundId"`
	FundName          string   `json:"fundName"`
	Coin              string   `json:"coin"`
	Status            string   `json:"status"`
	TotalEquity       string   `json:"totalEquity"`
	TotalShares       string   `json:"totalShares"`
	CurrentNav        string   `json:"currentNav"`
	CurrentAPR        string   `json:"currentAPR"`
	AccountUid        string   `json:"accountUid"`
	SubAccountList    []string `json:"subAccountList"`
	ProfitShareRate   string   `json:"profitShareRate"`
	ManagementFeeRate string   `json:"managementFeeRate"`
	UncollectedProfit string   `json:"uncollectedProfit"`
	CollectedProfit   string   `json:"collectedProfit"`
	TotalLoan         string   `json:"totalLoan"`
	CreatedTime       string   `json:"createdTime"`
}

type PwmInstSettleProfitResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstCreateFundResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstCreateInvestmentPlanResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstGetInvestmentPlansResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstInvestmentPlanInfo struct {
	PlanId                  string        `json:"planId"`
	PlanName                string        `json:"planName"`
	PlanType                string        `json:"planType"`
	SubscriptionUid         string        `json:"subscriptionUid"`
	Status                  string        `json:"status"`
	Source                  string        `json:"source"`
	CurrentAssetUsd         string        `json:"currentAssetUsd"`
	AccumulateYieldUsd      string        `json:"accumulateYieldUsd"`
	InvestmentDistribution  []interface{} `json:"investmentDistribution"`
	CreatedTime             string        `json:"createdTime"`
}

type PwmInstManageInvestmentPlanResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstListOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstOrderInfo struct {
	OrderId     string `json:"orderId"`
	FundId      string `json:"fundId"`
	FundName    string `json:"fundName"`
	AccountUid  string `json:"accountUid"`
	OrderType   string `json:"orderType"`
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	Shares      string `json:"shares"`
	Status      string `json:"status"`
	CreatedTime string `json:"createdTime"`
}

type PwmInstManageOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInstCreateSubAccountResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmFundTransferResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmQueryFundTransferResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmFundTransferRecord struct {
	TransferId string `json:"transferId"`
	Status     string `json:"status"`
	FromUserId int    `json:"fromUserId"`
	ToUserId   int    `json:"toUserId"`
	Amount     string `json:"amount"`
	Coin       string `json:"coin"`
}

type PwmListInvestmentPlansResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmGetPlanDetailResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmGetNewPlanDetailResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmClaimResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmAssetTrendResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmFundNavResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmSubscribeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInvestMoreResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmRedeemResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmListOrderResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmListProductCardsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmCreateCustomPlanResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PwmInvestmentPlanSummary struct {
	PlanId                 string        `json:"planId"`
	PlanName               string        `json:"planName"`
	PlanType               string        `json:"planType"`
	Status                 string        `json:"status"`
	Source                 string        `json:"source"`
	CurrentAssetUsd        string        `json:"currentAssetUsd"`
	AccumulateYieldUsd     string        `json:"accumulateYieldUsd"`
	InvestmentDistribution []interface{} `json:"investmentDistribution"`
	CreatedTime            string        `json:"createdTime"`
}

type PwmInvestmentDistribution struct {
	Category      string `json:"category"`
	ProductId     string `json:"productId"`
	Coin          string `json:"coin"`
	CurrentAmount string `json:"currentAmount"`
}

type PwmPositionItem struct {
	Category         string `json:"category"`
	ProductId        string `json:"productId"`
	Coin             string `json:"coin"`
	CurrentAmount    string `json:"currentAmount"`
	AccumulateYield  string `json:"accumulateYield"`
	Apr              string `json:"apr"`
	Duration         int    `json:"duration"`
	MaturityTime     string `json:"maturityTime"`
	AutoReinvest     bool   `json:"autoReinvest"`
	StakeAmount      string `json:"stakeAmount"`
	PositionId       int    `json:"positionId"`
	Status           string `json:"status"`
}

type PwmEquityFundPositionItem struct {
	Category        string   `json:"category"`
	ProductId       string   `json:"productId"`
	FundName        string   `json:"fundName"`
	Coin            string   `json:"coin"`
	Tags            []string `json:"tags"`
	Nav             string   `json:"nav"`
	UserShares      string   `json:"userShares"`
	ShareValue      string   `json:"shareValue"`
	HoldingValue    string   `json:"holdingValue"`
	AccumulateYield string   `json:"accumulateYield"`
	Apr30d          string   `json:"apr30d"`
	AprTotal        string   `json:"aprTotal"`
	SharpRatio      string   `json:"sharpRatio"`
	MaxDrawdown     string   `json:"maxDrawdown"`
	CreatedTime     string   `json:"createdTime"`
	RunningDays     int      `json:"runningDays"`
	PositionId      int      `json:"positionId"`
	Status          string   `json:"status"`
}

type PwmCategoryPositionGroup struct {
	TotalInvestmentUsd  string        `json:"totalInvestmentUsd"`
	AccumulateYieldUsd  string        `json:"accumulateYieldUsd"`
	WeightedAvgApr      string        `json:"weightedAvgApr"`
	Items               []interface{} `json:"items"`
}

type PwmEquityFundPositionGroup struct {
	TotalInvestmentUsd string        `json:"totalInvestmentUsd"`
	AccumulateYieldUsd string        `json:"accumulateYieldUsd"`
	WeightedAvgApr     string        `json:"weightedAvgApr"`
	Items              []interface{} `json:"items"`
}

type PwmPlanPositions struct {
	MultiCoinsEarning interface{}   `json:"multiCoinsEarning"`
	FixedYield        interface{}   `json:"fixedYield"`
	EquityFunds       interface{}   `json:"equityFunds"`
	OnchainEarn       interface{}   `json:"onchainEarn"`
	FundingAccount    []interface{} `json:"fundingAccount"`
}

type PwmCoinAmount struct {
	Coin   string `json:"coin"`
	Amount string `json:"amount"`
}

type PwmFundIntroduction struct {
	Description              string `json:"description"`
	HistoricalYieldRateMax   string `json:"historicalYieldRateMax"`
	HistoricalYieldRateMin   string `json:"historicalYieldRateMin"`
	SharpRatio               string `json:"sharpRatio"`
	MaxDrawback              string `json:"maxDrawback"`
	LockupPeriod             string `json:"lockupPeriod"`
}

type PwmConfiguredProductItem struct {
	Category        string        `json:"category"`
	ProductId       string        `json:"productId"`
	FundName        string        `json:"fundName"`
	Coin            string        `json:"coin"`
	ConfiguredAmount string       `json:"configuredAmount"`
	Apr             string        `json:"apr"`
	Duration        int           `json:"duration"`
	Tags            []string      `json:"tags"`
	Introduction    interface{}   `json:"introduction"`
}

type PwmConfiguredProductGroup struct {
	ConfiguredAmountUsd string        `json:"configuredAmountUsd"`
	Items               []interface{} `json:"items"`
}

type PwmNewPlanProducts struct {
	MultiCoinsEarning interface{} `json:"multiCoinsEarning"`
	FixedYield        interface{} `json:"fixedYield"`
	EquityFunds       interface{} `json:"equityFunds"`
	OnchainEarn       interface{} `json:"onchainEarn"`
}

type PwmClaimRequest struct {
	PlanId        string `json:"planId"`
	ToAccountType string `json:"toAccountType"`
	OrderLinkId   string `json:"orderLinkId"`
}

type PwmSubscribeRequest struct {
	PlanId      string `json:"planId"`
	AccountType string `json:"accountType"`
	OrderLinkId string `json:"orderLinkId"`
}

type PwmInvestMoreRequest struct {
	PlanId      string `json:"planId"`
	AccountType string `json:"accountType"`
	Category    string `json:"category"`
	ProductId   string `json:"productId"`
	Amount      string `json:"amount"`
	OrderLinkId string `json:"orderLinkId"`
}

type PwmRedeemRequest struct {
	PlanId      string `json:"planId"`
	Category    string `json:"category"`
	ProductId   string `json:"productId"`
	Shares      string `json:"shares"`
	Amount      string `json:"amount"`
	OrderLinkId string `json:"orderLinkId"`
	PositionId  int    `json:"positionId"`
}

type PwmAssetDataPoint struct {
	Date           string `json:"date"`
	AssetValueUsd  string `json:"assetValueUsd"`
}

type PwmNavDataPoint struct {
	Date string `json:"date"`
	Nav  string `json:"nav"`
}

type PwmOrderDetail struct {
	OrderId     string `json:"orderId"`
	PlanId      string `json:"planId"`
	Type        string `json:"type"`
	AccountType string `json:"accountType"`
	Coin        string `json:"coin"`
	Amount      string `json:"amount"`
	Category    string `json:"category"`
	ProductId   string `json:"productId"`
	Status      string `json:"status"`
	OrderTime   string `json:"orderTime"`
}

type PwmProductCard struct {
	Category            string   `json:"category"`
	ProductId           string   `json:"productId"`
	FundName            string   `json:"fundName"`
	Coin                string   `json:"coin"`
	Apr                 string   `json:"apr"`
	AprRangeLow         string   `json:"aprRangeLow"`
	AprRangeHigh        string   `json:"aprRangeHigh"`
	Tags                []string `json:"tags"`
	Introduction        string   `json:"introduction"`
	Aum                 string   `json:"aum"`
	MinInvestmentAmount string   `json:"minInvestmentAmount"`
	MaxInvestmentAmount string   `json:"maxInvestmentAmount"`
	Duration            int      `json:"duration"`
	MaxDrawdown         string   `json:"maxDrawdown"`
	SharpRatio          string   `json:"sharpRatio"`
	EstAPR              string   `json:"estAPR"`
}

type PwmProductTypeGroup struct {
	Type  string        `json:"type"`
	Cards []interface{} `json:"cards"`
}

type PwmCreateCustomPlanProductItem struct {
	Category  string `json:"category"`
	ProductId string `json:"productId"`
	FundName  string `json:"fundName"`
	Amount    string `json:"amount"`
}

type PwmCreateCustomPlanRequest struct {
	AccountType string        `json:"accountType"`
	Products    []interface{} `json:"products"`
	OrderLinkId string        `json:"orderLinkId"`
}

type LPOrderListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LPOrderDto struct {
	OrderType      int    `json:"orderType"`
	OrderNo        string `json:"orderNo"`
	OrderStatus    int    `json:"orderStatus"`
	PoolAddress    string `json:"poolAddress"`
	PoolName       string `json:"poolName"`
	PositionId     int    `json:"positionId"`
	TokenCode      string `json:"tokenCode"`
	TokenSymbol    string `json:"tokenSymbol"`
	TokenIconUrlDay   string `json:"tokenIconUrlDay"`
	TokenIconUrlNight string `json:"tokenIconUrlNight"`
	Amount         string `json:"amount"`
	ChainCode      string `json:"chainCode"`
	ChainIconUrl   string `json:"chainIconUrl"`
	GasTokenSymbol string `json:"gasTokenSymbol"`
	GasOnchain     string `json:"gasOnchain"`
	GasUsd         string `json:"gasUsd"`
	PlatformFee    string `json:"platformFee"`
	PlatformFeeUsd string `json:"platformFeeUsd"`
	CreateTime     int    `json:"createTime"`
	ExecutionTime  int    `json:"executionTime"`
	FailureReason  string `json:"failureReason"`
	DercRatio      string `json:"dercRatio"`
}

type LPPayTokenListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LPPayTokenDto struct {
	TokenCode         string `json:"tokenCode"`
	TokenSymbol       string `json:"tokenSymbol"`
	ChainCode         string `json:"chainCode"`
	ChainIconUrl      string `json:"chainIconUrl"`
	Decimals          int    `json:"decimals"`
	AvailableBalance  string `json:"availableBalance"`
	TokenIconUrlDay   string `json:"tokenIconUrlDay"`
	TokenIconUrlNight string `json:"tokenIconUrlNight"`
	MinStakeAmount    string `json:"minStakeAmount"`
	MaxStakeAmount    string `json:"maxStakeAmount"`
}

type LPPayTokenPriceResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type TokenPriceDto struct {
	TokenCode   string `json:"tokenCode"`
	TokenSymbol string `json:"tokenSymbol"`
	PriceUsd    string `json:"priceUsd"`
	ChainCode   string `json:"chainCode"`
	UpdateTime  int    `json:"updateTime"`
}

type LPPoolInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LPPoolListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LPPoolDto struct {
	PoolAddress        string `json:"poolAddress"`
	PoolName           string `json:"poolName"`
	PoolTag            string `json:"poolTag"`
	Apy                string `json:"apy"`
	Tvl                string `json:"tvl"`
	Token0Symbol       string `json:"token0Symbol"`
	Token0IconUrlDay   string `json:"token0IconUrlDay"`
	Token0IconUrlNight string `json:"token0IconUrlNight"`
	Token1Symbol       string `json:"token1Symbol"`
	Token1IconUrlDay   string `json:"token1IconUrlDay"`
	Token1IconUrlNight string `json:"token1IconUrlNight"`
	ChainCode          string `json:"chainCode"`
	ChainIconUrl       string `json:"chainIconUrl"`
}

type LPPositionListRequest struct{}

type LPPositionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LPPositionDto struct {
	PositionId       int    `json:"positionId"`
	PoolAddress      string `json:"poolAddress"`
	PoolName         string `json:"poolName"`
	StakedAmount     string `json:"stakedAmount"`
	StakedTokenCode  string `json:"stakedTokenCode"`
	StakedTokenSymbol string `json:"stakedTokenSymbol"`
	CurrentValueUsd  string `json:"currentValueUsd"`
	EarnedRewardsUsd string `json:"earnedRewardsUsd"`
	UnrealizedPnl    string `json:"unrealizedPnl"`
	RealizedPnl      string `json:"realizedPnl"`
	Apy              string `json:"apy"`
	Token0Amount     string `json:"token0Amount"`
	Token0Symbol     string `json:"token0Symbol"`
	Token1Amount     string `json:"token1Amount"`
	Token1Symbol     string `json:"token1Symbol"`
	RangeUpper       string `json:"rangeUpper"`
	RangeLower       string `json:"rangeLower"`
	CreateTime       int    `json:"createTime"`
	UpdateTime       int    `json:"updateTime"`
	Status           int    `json:"status"`
}

type LPRedeemResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LPStakeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}
