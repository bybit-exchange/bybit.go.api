package models

// SpotMarginDataResult holds data for spot margin including VIP level specific coin list.
type SpotMarginDataResult struct {
	VipCoinList []struct {
		List []struct {
			Borrowable         bool   `json:"borrowable"`
			CollateralRatio    string `json:"collateralRatio"`
			Currency           string `json:"currency"`
			HourlyBorrowRate   string `json:"hourlyBorrowRate"`
			LiquidationOrder   string `json:"liquidationOrder"`
			MarginCollateral   bool   `json:"marginCollateral"`
			MaxBorrowingAmount string `json:"maxBorrowingAmount"`
		} `json:"list"`
		VipLevel string `json:"vipLevel"`
	} `json:"vipCoinList"`
}

// ClassicalSpotMarginCoinResult holds information about coins in classical spot margin.
type ClassicalSpotMarginCoinResult struct {
	List []struct {
		Coin             string `json:"coin"`
		ConversionRate   string `json:"conversionRate"`
		LiquidationOrder int    `json:"liquidationOrder"`
	} `json:"list"`
}

// ClassicalSpotMarginBorrowCoinResult holds the borrowing precision information for coins in classical spot margin.
type ClassicalSpotMarginBorrowCoinResult struct {
	List []struct {
		Coin               string `json:"coin"`
		BorrowingPrecision int    `json:"borrowingPrecision"`
		RepaymentPrecision int    `json:"repaymentPrecision"`
	} `json:"list"`
}

// ClassicalSpotMarginInterestResult contains information about the interest rate on spot margin.
type ClassicalSpotMarginInterestResult struct {
	Coin           string `json:"coin"`
	InterestRate   string `json:"interestRate"`
	LoanAbleAmount string `json:"loanAbleAmount"`
	MaxLoanAmount  string `json:"maxLoanAmount"`
}

// ClassicalSpotMarginLoanResult holds the loan account information in the classical spot margin.
type ClassicalSpotMarginLoanResult struct {
	AcctBalanceSum  string `json:"acctBalanceSum"`
	DebtBalanceSum  string `json:"debtBalanceSum"`
	LoanAccountList []struct {
		Free         string `json:"free"`
		Interest     string `json:"interest"`
		Loan         string `json:"loan"`
		RemainAmount string `json:"remainAmount"`
		Locked       string `json:"locked"`
		TokenId      string `json:"tokenId"`
		Total        string `json:"total"`
	} `json:"loanAccountList"`
	RiskRate     string `json:"riskRate"`
	Status       int    `json:"status"`       // Use int for integer type
	SwitchStatus int    `json:"switchStatus"` // Use int for integer type
}

type SpotMarginBorrowOrders struct {
	List []struct {
		AccountId       string `json:"accountId"`       // Account ID
		Coin            string `json:"coin"`            // Coin name
		CreatedTime     int64  `json:"createdTime"`     // Borrow order created timestamp (ms)
		Id              string `json:"id"`              // Borrow order ID
		InterestAmount  string `json:"interestAmount"`  // Total interest
		InterestBalance string `json:"interestBalance"` // Outstanding interest
		LoanAmount      string `json:"loanAmount"`      // Principal amount
		LoanBalance     string `json:"loanBalance"`     // Outstanding principal
		RemainAmount    string `json:"remainAmount"`    // Remaining debt = interestBalance + loanBalance
		Status          int    `json:"status"`          // Status 1: uncleared, 2: cleared
		Type            int    `json:"type"`            // Order Type 1: manual loan, 2: auto loan
	} `json:"list"`
}

type SpotBorrowOrderResult struct {
	TransactId string `json:"transactId"`
}

type SpotRepayOrderResult struct {
	RepayId string `json:"repayId"`
}

type SpotToggleMarginResult struct {
	SwitchStatus string `json:"switchStatus"`
}

type SpotMarginLeverageResult struct {
	Leverage string `json:"leverage"`
}

type SpotMarginStateResult struct {
	SpotLeverage   string `json:"spotLeverage"`
	SpotMarginMode string `json:"spotMarginMode"`
}

type HistoricalInterestResult struct {
	List []InterestRateRecord `json:"list"`
}

type InterestRateRecord struct {
	Timestamp        int    `json:"timestamp"`
	Currency         string `json:"currency"`
	HourlyBorrowRate string `json:"hourlyBorrowRate"`
	VipLevel         string `json:"vipLevel"`
}

type PositionTiersResult struct {
	List []PositionTierItem `json:"list"`
}

type PositionTierItem struct {
	Currency               string              `json:"currency"`
	PositionTiersRatioList []PositionTierRatio `json:"positionTiersRatioList"`
}

type PositionTierRatio struct {
	Tier        string `json:"tier"`
	BorrowLimit string `json:"borrowLimit"`
	PositionMMR string `json:"positionMMR"`
	PositionIMR string `json:"positionIMR"`
	MaxLeverage string `json:"maxLeverage"`
}

type TieredCollateralRatioResult struct {
	List []CollateralTierItem `json:"list"`
}

type CollateralTierItem struct {
	Currency            string                    `json:"currency"`
	CollateralRatioList []SpotCollateralRatioTier `json:"collateralRatioList"`
}

type SpotCollateralRatioTier struct {
	MinQty          string `json:"minQty"`
	MaxQty          string `json:"maxQty"`
	CollateralRatio string `json:"collateralRatio"`
}

type VipMarginResult struct {
	VipCoinList []VipCoinGroup `json:"vipCoinList"`
}

type VipCoinGroup struct {
	VipLevel string          `json:"vipLevel"`
	List     []VipCoinDetail `json:"list"`
}

type VipCoinDetail struct {
	Borrowable         bool   `json:"borrowable"`
	CollateralRatio    string `json:"collateralRatio"`
	Currency           string `json:"currency"`
	HourlyBorrowRate   string `json:"hourlyBorrowRate"`
	LiquidationOrder   string `json:"liquidationOrder"`
	MarginCollateral   bool   `json:"marginCollateral"`
	MaxBorrowingAmount string `json:"maxBorrowingAmount"`
}

type SetLeverageRequest struct {
	Leverage string `json:"leverage"`
	Currency string `json:"currency"`
}

type SwitchModeResponse struct {
	SpotMarginMode string `json:"spotMarginMode"`
}

type GetSpotMarginCoinStateResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CoinStateItem struct {
	Currency     string `json:"currency"`
	SpotLeverage string `json:"spotLeverage"`
}

type GetAutoRepayModeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AutoRepayModeItem struct {
	Currency      string `json:"currency"`
	AutoRepayMode string `json:"autoRepayMode"`
}

type GetMaxBorrowableResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetRepaymentAvailableAmountResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetSpotMarginTradeStateResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryFixedBorrowContractsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type FixedBorrowContractItem struct {
	LoanId                  string `json:"loanId"`
	OrderId                 string `json:"orderId"`
	BorrowCurrency          string `json:"borrowCurrency"`
	AnnualRate              string `json:"annualRate"`
	Term                    string `json:"term"`
	ResidualPrincipal       string `json:"residualPrincipal"`
	InterestPaid            string `json:"interestPaid"`
	ResidualPenaltyInterest string `json:"residualPenaltyInterest"`
	BorrowTime              string `json:"borrowTime"`
	RepaymentTime           string `json:"repaymentTime"`
	Status                  int    `json:"status"`
	RepayType               string `json:"repayType"`
	StrategyType            string `json:"strategyType"`
}

type QueryFixedBorrowOrdersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type FixedBorrowOrderItem struct {
	OrderId       string `json:"orderId"`
	OrderCurrency string `json:"orderCurrency"`
	OrderQty      string `json:"orderQty"`
	FilledQty     string `json:"filledQty"`
	AnnualRate    string `json:"annualRate"`
	Term          int    `json:"term"`
	State         int    `json:"state"`
	RepayType     string `json:"repayType"`
	StrategyType  string `json:"strategyType"`
	OrderTime     string `json:"orderTime"`
}

type QueryFixedBorrowMarketResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type FixedBorrowMarketItem struct {
	OrderCurrency string `json:"orderCurrency"`
	Term          int    `json:"term"`
	AnnualRate    string `json:"annualRate"`
	Qty           string `json:"qty"`
}

type RenewFixedBorrowResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryBorrowLiabilityResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SetAutoRepayModeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}
