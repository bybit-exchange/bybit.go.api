package models

type DeliveryRecordInfo struct {
	Category       string                `json:"category"`
	List           []DeliveryRecordEntry `json:"list"`
	NextPageCursor string                `json:"nextPageCursor"`
}

type DeliveryRecordEntry struct {
	DeliveryTime  int64  `json:"deliveryTime"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Position      string `json:"position"`
	DeliveryPrice string `json:"deliveryPrice"`
	Strike        string `json:"strike"`
	Fee           string `json:"fee"`
	DeliveryPnl   string `json:"deliveryRpl"`
}

type USDCSettlementInfo struct {
	Category       string            `json:"category"`
	List           []SettlementEntry `json:"list"`
	NextPageCursor string            `json:"nextPageCursor"`
}

type SettlementEntry struct {
	Symbol          string `json:"symbol"`
	Side            string `json:"side"`
	Size            string `json:"size"`
	SessionAvgPrice string `json:"sessionAvgPrice"`
	MarkPrice       string `json:"markPrice"`
	RealisedPnl     string `json:"realisedPnl"`
	CreatedTime     string `json:"createdTime"`
}

type AssetInfo struct {
	Spot SpotInfo `json:"spot"`
}

type SpotInfo struct {
	Status string      `json:"status"`
	Assets []AssetItem `json:"assets"`
}

type AssetItem struct {
	Coin     string `json:"coin"`
	Frozen   string `json:"frozen"`
	Free     string `json:"free"`
	Withdraw string `json:"withdraw"`
}

type AllCoinsBalance struct {
	AccountType string          `json:"accountType"`
	MemberId    string          `json:"memberId"`
	Balance     []BalanceDetail `json:"balance"`
}

type BalanceDetail struct {
	Coin            string `json:"coin"`
	WalletBalance   string `json:"walletBalance"`
	TransferBalance string `json:"transferBalance"`
	Bonus           string `json:"bonus"`
}

type SingleCoinBalance struct {
	AccountType string      `json:"accountType"`
	BizType     int         `json:"bizType"`
	AccountId   string      `json:"accountId"`
	MemberId    string      `json:"memberId"`
	Balance     CoinBalance `json:"balance"`
}

type CoinBalance struct {
	Coin                  string `json:"coin"`
	WalletBalance         string `json:"walletBalance"`
	TransferBalance       string `json:"transferBalance"`
	Bonus                 string `json:"bonus"`
	TransferSafeAmount    string `json:"transferSafeAmount"`
	LtvTransferSafeAmount string `json:"ltvTransferSafeAmount"`
}

type TransferResult struct {
	TransferId string `json:"transferId"`
}

type InternalTransferInfo struct {
	List           []TransferDetail `json:"list"`
	NextPageCursor string           `json:"nextPageCursor"`
}

type TransferDetail struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type SubUidsInfo struct {
	SubMemberIds             []string `json:"subMemberIds"`
	TransferableSubMemberIds []string `json:"transferableSubMemberIds"`
}

type UniversalTransferInfo struct {
	List           []UniversalTransferDetail `json:"list"`
	NextPageCursor string                    `json:"nextPageCursor"`
}

type UniversalTransferDetail struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromMemberId    string `json:"fromMemberId"`
	ToMemberId      string `json:"toMemberId"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type WithdrawAssetResult struct {
	Id string `json:"id"`
}

type CancelWithdrawAssetResult struct {
	Status string `json:"status"`
}

type SetDepositAccountResult struct {
	Status string `json:"status"`
}

type AllowDepositCoinInfo struct {
	ConfigList     []DepositCoinConfig `json:"configList"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type DepositCoinConfig struct {
	Coin               string `json:"coin"`
	Chain              string `json:"chain"`
	CoinShowName       string `json:"coinShowName"`
	ChainType          string `json:"chainType"`
	BlockConfirmNumber int    `json:"blockConfirmNumber"`
	MinDepositAmount   string `json:"minDepositAmount"`
}

type DepositRecords struct {
	Rows           []DepositRecord `json:"rows"`
	NextPageCursor string          `json:"nextPageCursor"`
}

type DepositRecord struct {
	Coin              string `json:"coin"`
	Chain             string `json:"chain"`
	Amount            string `json:"amount"`
	TxID              string `json:"txID"`
	Status            int    `json:"status"`
	ToAddress         string `json:"toAddress"`
	Tag               string `json:"tag"`
	DepositFee        string `json:"depositFee"`
	SuccessAt         string `json:"successAt"`
	Confirmations     string `json:"confirmations"`
	TxIndex           string `json:"txIndex"`
	BlockHash         string `json:"blockHash"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
	DepositType       int    `json:"depositType"`
}

type SubDepositResult struct {
	Coin   string             `json:"coin"`
	Chains []DepositChainInfo `json:"chains"`
}

// MasterDepositResult represents the structure for master deposit results.
type MasterDepositResult struct {
	Coin   string             `json:"coin"`
	Chains []DepositChainInfo `json:"chains"`
}

// DepositChainInfo represents the shared structure for deposit chain information.
type DepositChainInfo struct {
	ChainType         string `json:"chainType"`
	AddressDeposit    string `json:"addressDeposit"`
	TagDeposit        string `json:"tagDeposit"`
	Chain             string `json:"chain"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
}

// CoinInfoResult represents the structure for coin info results.
type CoinInfoResult struct {
	Rows []CoinInfoRow `json:"rows"`
}

// CoinInfoRow represents the structure for each row of coin information.
type CoinInfoRow struct {
	Name         string          `json:"name"`
	Coin         string          `json:"coin"`
	RemainAmount string          `json:"remainAmount"`
	Chains       []CoinChainInfo `json:"chains"`
}

// CoinChainInfo represents the structure for each chain's information for a coin.
type CoinChainInfo struct {
	Chain                 string `json:"chain"`
	ChainType             string `json:"chainType"`
	Confirmation          string `json:"confirmation"`
	WithdrawFee           string `json:"withdrawFee"`
	DepositMin            string `json:"depositMin"`
	WithdrawMin           string `json:"withdrawMin"`
	MinAccuracy           string `json:"minAccuracy"`
	ChainDeposit          string `json:"chainDeposit"`
	ChainWithdraw         string `json:"chainWithdraw"`
	WithdrawPercentageFee string `json:"withdrawPercentageFee"`
}

// WithdrawRecords represents the structure for withdrawal records.
type WithdrawRecords struct {
	Rows           []WithdrawRecord `json:"rows"`
	NextPageCursor string           `json:"nextPageCursor"`
}

// WithdrawRecord represents the structure for each withdrawal record.
type WithdrawRecord struct {
	WithdrawID   string `json:"withdrawId"`
	TxID         string `json:"txID"`
	WithdrawType string `json:"withdrawType"`
	Coin         string `json:"coin"`
	Chain        string `json:"chain"`
	Amount       string `json:"amount"`
	WithdrawFee  string `json:"withdrawFee"`
	Status       string `json:"status"`
	ToAddress    string `json:"toAddress"`
	Tag          string `json:"tag"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

// WithdrawableAmount represents the structure for information about withdrawable amounts.
type WithdrawableAmount struct {
	LimitAmountUsd      string              `json:"limitAmountUsd"`
	WithdrawableAmounts []WithdrawableAsset `json:"withdrawableAmount"`
}

// WithdrawableAsset represents the structure for withdrawable amounts for each wallet type.
type WithdrawableAsset struct {
	SPOT *WalletInfo `json:"SPOT,omitempty"` // Omitted if empty
	FUND *WalletInfo `json:"FUND,omitempty"` // Omitted if empty
}

// WalletInfo represents the wallet information for withdrawable assets.
type WalletInfo struct {
	Coin               string `json:"coin"`
	WithdrawableAmount string `json:"withdrawableAmount"`
	AvailableBalance   string `json:"availableBalance"`
}

type CreatePayResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type PayResultResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type MockStatusResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type FxConvertResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type RefundResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type PayoutResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type PayOrder struct {
	MerchantId      string        `json:"merchantId"`
	ClientId        string        `json:"clientId"`
	PaymentType     string        `json:"paymentType"`
	MerchantTradeNo string        `json:"merchantTradeNo"`
	PayId           string        `json:"payId"`
	Status          string        `json:"status"`
	Amount          string        `json:"amount"`
	Currency        string        `json:"currency"`
	CurrencyType    string        `json:"currencyType"`
	CreateTime      int           `json:"createTime"`
	PaymentTime     int           `json:"paymentTime"`
	FinishTime      int           `json:"finishTime"`
	RefundOrders    []interface{} `json:"refundOrders"`
	Remark          string        `json:"remark"`
}

type RefundOrder struct {
	RefundId         string `json:"refundId"`
	RefundType       string `json:"refundType"`
	MerchantTradeNo  string `json:"merchantTradeNo"`
	MerchantRefundNo string `json:"merchantRefundNo"`
	PayId            string `json:"payId"`
	RefundStatus     string `json:"refundStatus"`
	RefundCurrency   string `json:"refundCurrency"`
	Amount           string `json:"amount"`
	CreateTime       int    `json:"createTime"`
}

type AgreementSignResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementUnsignResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementPayResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementPayWithSignResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementQueryResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementListResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementPayQueryResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementPayListResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementRefundResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AgreementBaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AgreementMonetaryAmount struct {
	Total        string `json:"total"`
	Currency     string `json:"currency"`
	CurrencyType string `json:"currency_type"`
	Chain        string `json:"chain"`
}

type Env struct {
	TerminalType   string `json:"terminalType"`
	Device         string `json:"device"`
	BrowserVersion string `json:"browserVersion"`
	Ip             string `json:"ip"`
}

type Customer struct {
	Uid            string `json:"uid"`
	ExternalUserId string `json:"externalUserId"`
	UserName       string `json:"userName"`
	RegisterTime   string `json:"registerTime"`
	KycCountry     string `json:"kycCountry"`
	Remarks        string `json:"remarks"`
}

type Payee struct {
	Uid string `json:"uid"`
}

type CoinListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type ReferencePriceResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type QuoteApplyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TradeExecuteResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TradeQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type TradeHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type BalanceQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type FiatCoin struct {
	Coin               string `json:"coin"`
	FullName           string `json:"fullName"`
	Icon               string `json:"icon"`
	IconNight          string `json:"iconNight"`
	Precision          int    `json:"precision"`
	Disable            bool   `json:"disable"`
	SingleFromMinLimit string `json:"singleFromMinLimit"`
	SingleFromMaxLimit string `json:"singleFromMaxLimit"`
}

type CryptoCoin struct {
	Coin               string `json:"coin"`
	FullName           string `json:"fullName"`
	Icon               string `json:"icon"`
	IconNight          string `json:"iconNight"`
	Precision          int    `json:"precision"`
	Disable            bool   `json:"disable"`
	SingleFromMinLimit string `json:"singleFromMinLimit"`
	SingleFromMaxLimit string `json:"singleFromMaxLimit"`
}

type BalanceInfo struct {
	TotalBalance  string      `json:"totalBalance"`
	Balance       string      `json:"balance"`
	FrozenBalance string      `json:"frozenBalance"`
	Currency      interface{} `json:"currency"`
}

type CurrencyInfo struct {
	CurrencyType string `json:"currencyType"`
	CurrencyCode string `json:"currencyCode"`
	CurrencyName string `json:"currencyName"`
}

type QuotaInfo struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type PriceQuote struct {
	UnitPrice     string      `json:"unitPrice"`
	PaymentMethod string      `json:"paymentMethod"`
	Quota         interface{} `json:"quota"`
}

type QuoteApplyRequest struct {
	FromCoin        string `json:"fromCoin"`
	FromCoinType    string `json:"fromCoinType"`
	ToCoin          string `json:"toCoin"`
	ToCoinType      string `json:"toCoinType"`
	RequestAmount   string `json:"requestAmount"`
	RequestCoinType string `json:"requestCoinType"`
}

type TradeExecuteRequest struct {
	QuoteTxId         string `json:"quoteTxId"`
	SubUserId         string `json:"subUserId"`
	WebhookUrl        string `json:"webhookUrl"`
	MerchantRequestId string `json:"merchantRequestId"`
}

type BasicResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
	ExtCode string      `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	TimeNow string      `json:"time_now"`
}

type TradingPreferenceSet struct {
	HasUnPostAd               string `json:"hasUnPostAd"`
	IsKyc                     string `json:"isKyc"`
	IsEmail                   string `json:"isEmail"`
	IsMobile                  string `json:"isMobile"`
	HasRegisterTime           string `json:"hasRegisterTime"`
	RegisterTimeThreshold     string `json:"registerTimeThreshold"`
	OrderFinishNumberDay30    string `json:"orderFinishNumberDay30"`
	CompleteRateDay30         string `json:"completeRateDay30"`
	NationalLimit             string `json:"nationalLimit"`
	HasOrderFinishNumberDay30 string `json:"hasOrderFinishNumberDay30"`
	HasCompleteRateDay30      string `json:"hasCompleteRateDay30"`
	HasNationalLimit          string `json:"hasNationalLimit"`
}

type GetCoinBalanceResponse struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Result  interface{} `json:"result"`
}

type GetAdsResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type PostAdResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type UpdateAdResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type AdItem struct {
	Id                   string        `json:"id"`
	AccountId            string        `json:"accountId"`
	UserId               string        `json:"userId"`
	NickName             string        `json:"nickName"`
	TokenId              string        `json:"tokenId"`
	CurrencyId           string        `json:"currencyId"`
	Side                 int           `json:"side"`
	PriceType            int           `json:"priceType"`
	Price                string        `json:"price"`
	Premium              string        `json:"premium"`
	LastQuantity         string        `json:"lastQuantity"`
	Quantity             string        `json:"quantity"`
	FrozenQuantity       string        `json:"frozenQuantity"`
	ExecutedQuantity     string        `json:"executedQuantity"`
	MinAmount            string        `json:"minAmount"`
	MaxAmount            string        `json:"maxAmount"`
	Remark               string        `json:"remark"`
	Status               int           `json:"status"`
	CreateDate           string        `json:"createDate"`
	Payments             []string      `json:"payments"`
	HiddenReason         string        `json:"hiddenReason"`
	TradingPreferenceSet interface{}   `json:"tradingPreferenceSet"`
	UpdateDate           string        `json:"updateDate"`
	FeeRate              string        `json:"feeRate"`
	PaymentPeriod        int           `json:"paymentPeriod"`
	ItemType             string        `json:"itemType"`
	PaymentTerms         []interface{} `json:"paymentTerms"`
}

type GetMyAdsResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type GetAdDetailResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type OrderItem struct {
	Id                  string      `json:"id"`
	Side                int         `json:"side"`
	TokenId             string      `json:"tokenId"`
	OrderType           string      `json:"orderType"`
	Amount              string      `json:"amount"`
	CurrencyId          string      `json:"currencyId"`
	Price               string      `json:"price"`
	Fee                 string      `json:"fee"`
	TargetNickName      string      `json:"targetNickName"`
	TargetUserId        string      `json:"targetUserId"`
	Status              int         `json:"status"`
	CreateDate          string      `json:"createDate"`
	TransferLastSeconds string      `json:"transferLastSeconds"`
	UserId              string      `json:"userId"`
	SellerRealName      string      `json:"sellerRealName"`
	BuyerRealName       string      `json:"buyerRealName"`
	Extension           interface{} `json:"extension"`
}

type GetAllOrdersResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type PaymentTermItem struct {
	Id          string `json:"id"`
	RealName    string `json:"realName"`
	PaymentType int    `json:"paymentType"`
	BankName    string `json:"bankName"`
	BranchName  string `json:"branchName"`
	AccountNo   string `json:"accountNo"`
	Qrcode      string `json:"qrcode"`
}

type GetOrderDetailResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type UploadFileResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type GetChatMessageResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type UserInfo struct {
	NickName             string `json:"nickName"`
	DefaultNickName      bool   `json:"defaultNickName"`
	IsOnline             bool   `json:"isOnline"`
	KycLevel             string `json:"kycLevel"`
	Email                string `json:"email"`
	Mobile               string `json:"mobile"`
	LastLogoutTime       string `json:"lastLogoutTime"`
	RecentRate           string `json:"recentRate"`
	TotalFinishCount     int    `json:"totalFinishCount"`
	TotalFinishSellCount int    `json:"totalFinishSellCount"`
	TotalFinishBuyCount  int    `json:"totalFinishBuyCount"`
	RecentFinishCount    int    `json:"recentFinishCount"`
	AverageReleaseTime   string `json:"averageReleaseTime"`
	AverageTransferTime  string `json:"averageTransferTime"`
	AccountCreateDays    int    `json:"accountCreateDays"`
	FirstTradeDays       int    `json:"firstTradeDays"`
	RealName             string `json:"realName"`
	RecentTradeAmount    string `json:"recentTradeAmount"`
	TotalTradeAmount     string `json:"totalTradeAmount"`
	RegisterTime         string `json:"registerTime"`
	AuthStatus           int    `json:"authStatus"`
	KycCountryCode       string `json:"kycCountryCode"`
	Blocked              string `json:"blocked"`
	GoodAppraiseRate     string `json:"goodAppraiseRate"`
	GoodAppraiseCount    int    `json:"goodAppraiseCount"`
	BadAppraiseCount     int    `json:"badAppraiseCount"`
	VipLevel             int    `json:"vipLevel"`
	UserId               string `json:"userId"`
	RealNameEn           string `json:"realNameEn"`
}

type GetCounterpartyInfoResponse struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	Result  interface{} `json:"result"`
}

type PaymentMethod struct {
	Id                        string      `json:"id"`
	RealName                  string      `json:"realName"`
	PaymentType               string      `json:"paymentType"`
	BankName                  string      `json:"bankName"`
	BranchName                string      `json:"branchName"`
	AccountNo                 string      `json:"accountNo"`
	Qrcode                    string      `json:"qrcode"`
	Online                    string      `json:"online"`
	Visible                   int         `json:"visible"`
	PayMessage                string      `json:"payMessage"`
	FirstName                 string      `json:"firstName"`
	LastName                  string      `json:"lastName"`
	SecondLastName            string      `json:"secondLastName"`
	Clabe                     string      `json:"clabe"`
	DebitCardNumber           string      `json:"debitCardNumber"`
	Concept                   string      `json:"concept"`
	CountNo                   string      `json:"countNo"`
	PaymentExt1               string      `json:"paymentExt1"`
	PaymentExt2               string      `json:"paymentExt2"`
	PaymentExt3               string      `json:"paymentExt3"`
	PaymentExt4               string      `json:"paymentExt4"`
	PaymentExt5               string      `json:"paymentExt5"`
	PaymentExt6               string      `json:"paymentExt6"`
	PaymentTemplateVersion    int         `json:"paymentTemplateVersion"`
	HasPaymentTemplateChanged bool        `json:"hasPaymentTemplateChanged"`
	PaymentConfigVo           interface{} `json:"paymentConfigVo"`
	RealNameVerified          bool        `json:"realNameVerified"`
	Channel                   string      `json:"channel"`
	CurrencyBalance           []string    `json:"currencyBalance"`
}

type GetUserPaymentResponse struct {
	RetCode int           `json:"ret_code"`
	RetMsg  string        `json:"ret_msg"`
	Result  []interface{} `json:"result"`
}

type FundingDetailApiResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type FundingDetailApiResult struct {
	NextPageCursor string               `json:"nextPageCursor"`
	List           []FundingDetailApiBO `json:"list"`
}

type FundingDetailApiBO struct {
	MemberId       string `json:"memberId"`
	Currency       string `json:"currency"`
	IoDirection    string `json:"ioDirection"`
	TxnAmt         string `json:"txnAmt"`
	AfterAmt       string `json:"afterAmt"`
	CreateTime     string `json:"createTime"`
	ShowBusiType   string `json:"showBusiType"`
	ShowBusiTypeEn string `json:"showBusiTypeEn"`
	Description    string `json:"description"`
	DescriptionEn  string `json:"descriptionEn"`
}

type AssetBaseResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type CoinExchangeTradeInfo struct {
	TradeNo      string `json:"tradeNo"`
	Status       string `json:"status"`
	QuoteTxId    string `json:"quoteTxId"`
	ExchangeRate string `json:"exchangeRate"`
	FromCoin     string `json:"fromCoin"`
	FromCoinType string `json:"fromCoinType"`
	ToCoin       string `json:"toCoin"`
	ToCoinType   string `json:"toCoinType"`
	FromAmount   string `json:"fromAmount"`
	ToAmount     string `json:"toAmount"`
	CreatedAt    string `json:"createdAt"`
	SubUserId    string `json:"subUserId"`
}

type BatchSettingCollateralResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type BatchSettingResult struct {
	List []BatchSettingResultItem `json:"list"`
}

type BatchSettingResultItem struct {
	Coin             string `json:"coin"`
	CollateralSwitch string `json:"collateralSwitch"`
}

type QueryCoinChainInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CoinChainDetail struct {
	Name         string        `json:"name"`
	Coin         string        `json:"coin"`
	RemainAmount string        `json:"remainAmount"`
	Chains       []ChainDetail `json:"chains"`
}

type ChainDetail struct {
	ChainType             string `json:"chainType"`
	Confirmation          string `json:"confirmation"`
	WithdrawFee           string `json:"withdrawFee"`
	DepositMin            string `json:"depositMin"`
	WithdrawMin           string `json:"withdrawMin"`
	Chain                 string `json:"chain"`
	ChainDeposit          string `json:"chainDeposit"`
	ChainWithdraw         string `json:"chainWithdraw"`
	MinAccuracy           string `json:"minAccuracy"`
	WithdrawPercentageFee string `json:"withdrawPercentageFee"`
	ContractAddress       string `json:"contractAddress"`
	SafeConfirmNumber     string `json:"safeConfirmNumber"`
}

type QueryDepositRecordsResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type DepositRecordItem struct {
	Coin                string `json:"coin"`
	Chain               string `json:"chain"`
	Amount              string `json:"amount"`
	TxID                string `json:"txID"`
	Status              int    `json:"status"`
	ToAddress           string `json:"toAddress"`
	Tag                 string `json:"tag"`
	DepositFee          string `json:"depositFee"`
	SuccessAt           string `json:"successAt"`
	Confirmations       string `json:"confirmations"`
	TxIndex             string `json:"txIndex"`
	BlockHash           string `json:"blockHash"`
	BatchReleaseLimit   string `json:"batchReleaseLimit"`
	DepositType         string `json:"depositType"`
	FromAddress         string `json:"fromAddress"`
	TaxDepositRecordsId string `json:"taxDepositRecordsId"`
	TaxStatus           int    `json:"taxStatus"`
	Id                  string `json:"id"`
}

type QueryDepositAddressResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type DepositAddress struct {
	ChainType         string `json:"chainType"`
	AddressDeposit    string `json:"addressDeposit"`
	TagDeposit        string `json:"tagDeposit"`
	Chain             string `json:"chain"`
	BatchReleaseLimit string `json:"batchReleaseLimit"`
	ContractAddress   string `json:"contractAddress"`
}

type QuerySubMemberDepositAddressResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QuerySubMemberDepositRecordsResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryInternalDepositRecordsResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type InternalDepositRecordItem struct {
	Id                  string `json:"id"`
	Amount              string `json:"amount"`
	Type                int    `json:"type"`
	Coin                string `json:"coin"`
	Address             string `json:"address"`
	Status              int    `json:"status"`
	CreatedTime         string `json:"createdTime"`
	FromMemberId        string `json:"fromMemberId"`
	TxID                string `json:"txID"`
	TaxDepositRecordsId string `json:"taxDepositRecordsId"`
	TaxStatus           int    `json:"taxStatus"`
}

type SetDefaultDepositToAccountResp struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type InterTransferResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type UniversalTransferResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type InterTransferListQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type InterTransferItem struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type UniversalTransferListQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type UniversalTransferItem struct {
	TransferId      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromMemberId    string `json:"fromMemberId"`
	ToMemberId      string `json:"toMemberId"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type TransferCoinListQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SubMemberListQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AccountCoinBalanceQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type AssetInfoQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SpotAsset struct {
	Coin     string `json:"coin"`
	Free     string `json:"free"`
	Frozen   string `json:"frozen"`
	Withdraw string `json:"withdraw"`
}

type UserAssetInfoQueryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type UserCoinBalance struct {
	Coin            string `json:"coin"`
	WalletBalance   string `json:"walletBalance"`
	TransferBalance string `json:"transferBalance"`
	Bonus           string `json:"bonus"`
}

type SendWithdrawResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryWithdrawRecordsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type WithdrawRecordItem struct {
	Coin         string `json:"coin"`
	Chain        string `json:"chain"`
	Amount       string `json:"amount"`
	TxID         string `json:"txID"`
	Status       string `json:"status"`
	ToAddress    string `json:"toAddress"`
	Tag          string `json:"tag"`
	WithdrawFee  string `json:"withdrawFee"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	WithdrawId   string `json:"withdrawId"`
	WithdrawType int    `json:"withdrawType"`
	Tax          string `json:"tax"`
	TaxRate      string `json:"taxRate"`
	TaxType      string `json:"taxType"`
}

type CancelWithdrawResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetWithdrawableAmountResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type WithdrawableAmountItem struct {
	Coin               string `json:"coin"`
	WithdrawableAmount string `json:"withdrawableAmount"`
	AvailableBalance   string `json:"availableBalance"`
}

type GetVASPListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type VASPItem struct {
	VaspEntityId string `json:"vaspEntityId"`
	VaspName     string `json:"vaspName"`
}

type QueryWithdrawAddressesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type WithdrawAddressItem struct {
	Coin        string `json:"coin"`
	Chain       string `json:"chain"`
	Address     string `json:"address"`
	Tag         string `json:"tag"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
	AddressType int    `json:"addressType"`
	Verified    int    `json:"verified"`
	CreateAt    string `json:"createAt"`
}

type GetOptionAssetInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type OptionAssetInfoItem struct {
	Coin       string `json:"coin"`
	TotalDelta string `json:"totalDelta"`
	TotalRPL   string `json:"totalRPL"`
	TotalUPL   string `json:"totalUPL"`
	AssetIM    string `json:"assetIM"`
	AssetMM    string `json:"assetMM"`
	SendTime   int    `json:"sendTime"`
}

type GetPayInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PayInfoCollateralItem struct {
	Coin            string `json:"coin"`
	AvailableSize   string `json:"availableSize"`
	AvailableValue  string `json:"availableValue"`
	CoinScale       int    `json:"coinScale"`
	BorrowSize      string `json:"borrowSize"`
	SpotHedgeAmount string `json:"spotHedgeAmount"`
	AssetFrozen     string `json:"assetFrozen"`
}

type GetUserSettingConfigResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetWalletBalanceResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type WalletBalanceItem struct {
	AccountType            string                  `json:"accountType"`
	AccountIMRate          string                  `json:"accountIMRate"`
	AccountMMRate          string                  `json:"accountMMRate"`
	TotalEquity            string                  `json:"totalEquity"`
	TotalWalletBalance     string                  `json:"totalWalletBalance"`
	TotalMarginBalance     string                  `json:"totalMarginBalance"`
	TotalAvailableBalance  string                  `json:"totalAvailableBalance"`
	TotalPerpUPL           string                  `json:"totalPerpUPL"`
	TotalInitialMargin     string                  `json:"totalInitialMargin"`
	TotalMaintenanceMargin string                  `json:"totalMaintenanceMargin"`
	Coin                   []WalletBalanceCoinItem `json:"coin"`
}

type WalletBalanceCoinItem struct {
	Coin                string `json:"coin"`
	Equity              string `json:"equity"`
	UsdValue            string `json:"usdValue"`
	WalletBalance       string `json:"walletBalance"`
	Locked              string `json:"locked"`
	SpotHedgingQty      string `json:"spotHedgingQty"`
	BorrowAmount        string `json:"borrowAmount"`
	AccruedInterest     string `json:"accruedInterest"`
	TotalOrderIM        string `json:"totalOrderIM"`
	TotalPositionIM     string `json:"totalPositionIM"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Bonus               string `json:"bonus"`
	MarginCollateral    bool   `json:"marginCollateral"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	SpotBorrow          string `json:"spotBorrow"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	Free                string `json:"free"`
	ColRes              string `json:"colRes"`
}

type GetAccountWithdrawalResponse struct {
	RetCode                int         `json:"retCode"`
	RetMsg                 string      `json:"retMsg"`
	AvailableWithdrawal    string      `json:"availableWithdrawal"`
	AvailableWithdrawalMap interface{} `json:"availableWithdrawalMap"`
}

type GetAssetOverviewResponse struct {
	RetCode     int                        `json:"retCode"`
	RetMsg      string                     `json:"retMsg"`
	TotalEquity string                     `json:"totalEquity"`
	List        []AssetOverviewAccountItem `json:"list"`
	RetExtInfo  interface{}                `json:"retExtInfo"`
	Time        int                        `json:"time"`
}

type AssetOverviewAccountItem struct {
	AccountType       string                    `json:"accountType"`
	SnapshotTime      string                    `json:"snapshotTime"`
	ValuationCurrency string                    `json:"valuationCurrency"`
	TotalEquity       string                    `json:"totalEquity"`
	Categories        []AssetOverviewCategory   `json:"categories"`
	CoinDetail        []AssetOverviewCoinDetail `json:"coinDetail"`
}

type AssetOverviewCategory struct {
	Category   string                    `json:"category"`
	Equity     string                    `json:"equity"`
	CoinDetail []AssetOverviewCoinDetail `json:"coinDetail"`
}

type AssetOverviewCoinDetail struct {
	Coin   string      `json:"coin"`
	Equity string      `json:"equity"`
	ExtMap interface{} `json:"extMap"`
}

type GetDeliveryRecordResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type DeliveryRecordItem struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	DeliveryTime  int    `json:"deliveryTime"`
	Strike        string `json:"strike"`
	Fee           string `json:"fee"`
	Position      string `json:"position"`
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryRpl   string `json:"deliveryRpl"`
}

type GetPortfolioMarginResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type PortfolioMarginWallet struct {
	Equity            string `json:"equity"`
	CashBalance       string `json:"cashBalance"`
	MarginBalance     string `json:"marginBalance"`
	AvailableBalance  string `json:"availableBalance"`
	TotalRPL          string `json:"totalRPL"`
	TotalSessionRPL   string `json:"totalSessionRPL"`
	TotalSessionUPL   string `json:"totalSessionUPL"`
	AccountIM         string `json:"accountIM"`
	AccountMM         string `json:"accountMM"`
	ExperienceBalance string `json:"experienceBalance"`
	PerpUPL           string `json:"perpUPL"`
	AccountMMRate     string `json:"accountMMRate"`
	AccountIMRate     string `json:"accountIMRate"`
}

type GetSettlementRecordResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SettlementRecordItem struct {
	Symbol          string `json:"symbol"`
	Side            string `json:"side"`
	Size            string `json:"size"`
	SessionAvgPrice string `json:"sessionAvgPrice"`
	MarkPrice       string `json:"markPrice"`
	RealisedPnl     string `json:"realisedPnl"`
	CreatedTime     string `json:"createdTime"`
}

type GetTotalMembersAssetsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type TotalMembersAssetItem struct {
	Uid    int                         `json:"uid"`
	IsM    bool                        `json:"isM"`
	Stat   int                         `json:"stat"`
	Origb  string                      `json:"origb"`
	Quoteb string                      `json:"quoteb"`
	Items  []TotalMembersAssetTypeItem `json:"items"`
}

type TotalMembersAssetTypeItem struct {
	Type   string `json:"type"`
	Origb  string `json:"origb"`
	Quoteb string `json:"quoteb"`
	Stat   int    `json:"stat"`
}

type GetExecutionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type ExecutionItem struct {
	Symbol          string `json:"symbol"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	Side            string `json:"side"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	CreateType      string `json:"createType"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	ExecType        string `json:"execType"`
	ExecFee         string `json:"execFee"`
	FeeRate         string `json:"feeRate"`
	FeeCurrency     string `json:"feeCurrency"`
	IsMaker         bool   `json:"isMaker"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	BlockTradeId    string `json:"blockTradeId"`
	ClosedSize      string `json:"closedSize"`
	Seq             int    `json:"seq"`
	ExtraFees       string `json:"extraFees"`
}

type GetSpreadExecutionListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SpreadExecutionItem struct {
	Symbol      string                 `json:"symbol"`
	OrderId     string                 `json:"orderId"`
	OrderLinkId string                 `json:"orderLinkId"`
	Side        string                 `json:"side"`
	ExecId      string                 `json:"execId"`
	ExecPrice   string                 `json:"execPrice"`
	ExecQty     string                 `json:"execQty"`
	ExecType    string                 `json:"execType"`
	ExecTime    string                 `json:"execTime"`
	Legs        []SpreadLegTradeDetail `json:"legs"`
}

type SpreadLegTradeDetail struct {
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	Category    string `json:"category"`
	ExecId      string `json:"execId"`
	ExecPrice   string `json:"execPrice"`
	ExecQty     string `json:"execQty"`
	ExecValue   string `json:"execValue"`
	ExecFee     string `json:"execFee"`
	ExecFeeV2   string `json:"execFeeV2"`
	FeeCurrency string `json:"feeCurrency"`
	ExecType    string `json:"execType"`
	ExecTime    string `json:"execTime"`
}

type GetSpreadMaxQtyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CoinListQueryResponse struct {
	Coins []CoinInfo `json:"coins"`
}

type ConvertExecuteResponse struct {
	ExchangeStatus string `json:"exchangeStatus"`
	QuoteTxId      string `json:"quoteTxId"`
}

type ConvertHistoryQueryResponse struct {
	List []ConvertResult `json:"list"`
}

type ConvertResult struct {
	AccountType    string      `json:"accountType"`
	ExchangeTxId   string      `json:"exchangeTxId"`
	UserId         string      `json:"userId"`
	FromCoin       string      `json:"fromCoin"`
	FromCoinType   string      `json:"fromCoinType"`
	FromAmount     string      `json:"fromAmount"`
	ToCoin         string      `json:"toCoin"`
	ToCoinType     string      `json:"toCoinType"`
	ToAmount       string      `json:"toAmount"`
	ExchangeStatus string      `json:"exchangeStatus"`
	ExtInfo        interface{} `json:"extInfo"`
	ConvertRate    string      `json:"convertRate"`
	CreatedAt      string      `json:"createdAt"`
}

type CoinConvertLimitQueryResponse struct {
	Min       string `json:"min"`
	Max       string `json:"max"`
	ErrorCode string `json:"errorCode"`
}

type QueryResultResponse struct {
	Result interface{} `json:"result"`
}

type LimitOrderCallbackResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type QueryOrderByPageResponse struct {
	OrderBody      []ExchangeBody `json:"orderBody"`
	NextPageCursor string         `json:"nextPageCursor"`
}

type ExchangeBody struct {
	FromCoin     string `json:"fromCoin"`
	FromAmount   string `json:"fromAmount"`
	ToCoin       string `json:"toCoin"`
	ToAmount     string `json:"toAmount"`
	ExchangeRate string `json:"exchangeRate"`
	CreatedTime  int    `json:"createdTime"`
	ExchangeTxId string `json:"exchangeTxId"`
}

type QueryOrderFromOpenApiResponse struct {
	NextPageCursor string      `json:"nextPageCursor"`
	TotalCount     int         `json:"totalCount"`
	Order          []OrderBody `json:"order"`
}

type OrderBody struct {
	FromCoin       string `json:"fromCoin"`
	FromAmount     string `json:"fromAmount"`
	ToCoin         string `json:"toCoin"`
	ToAmount       string `json:"toAmount"`
	ExchangeRate   string `json:"exchangeRate"`
	CalcType       int    `json:"calcType"`
	CreatedTime    int    `json:"createdTime"`
	ExchangeStatus string `json:"exchangeStatus"`
	ExchangeTxId   string `json:"exchangeTxId"`
}

type TaxAndFee struct {
	TaxFeeName   string `json:"taxFeeName"`
	TaxFeeRate   string `json:"taxFeeRate"`
	TaxFeeAmount string `json:"taxFeeAmount"`
	TaxFeeCoin   string `json:"taxFeeCoin"`
}

type SmallAssetConvertResponse struct {
	QuoteId      string `json:"quoteId"`
	ExchangeTxId string `json:"exchangeTxId"`
	SubmitTime   string `json:"submitTime"`
	Status       string `json:"status"`
	Msg          string `json:"msg"`
}

type QuerySmallAssetConvertOrderResponse struct {
	Cursor     string              `json:"cursor"`
	Size       string              `json:"size"`
	LastPage   string              `json:"lastPage"`
	TotalCount string              `json:"totalCount"`
	Records    []MainConvertRecord `json:"records"`
}

type MainConvertRecord struct {
	AccountType     string             `json:"accountType"`
	ExchangeTxId    string             `json:"exchangeTxId"`
	ToCoin          string             `json:"toCoin"`
	ToAmount        string             `json:"toAmount"`
	SubRecords      []SubConvertRecord `json:"subRecords"`
	Status          string             `json:"status"`
	CreatedAt       string             `json:"createdAt"`
	ExchangeSource  string             `json:"exchangeSource"`
	FeeCoin         string             `json:"feeCoin"`
	TotalFeeAmount  string             `json:"totalFeeAmount"`
	TotalTaxFeeInfo interface{}        `json:"totalTaxFeeInfo"`
}

type SubConvertRecord struct {
	FromCoin   string      `json:"fromCoin"`
	FromAmount string      `json:"fromAmount"`
	ToCoin     string      `json:"toCoin"`
	ToAmount   string      `json:"toAmount"`
	FeeCoin    string      `json:"feeCoin"`
	FeeAmount  string      `json:"feeAmount"`
	Status     string      `json:"status"`
	TaxFeeInfo interface{} `json:"taxFeeInfo"`
}

type TaxFeeInfo struct {
	TotalAmount string       `json:"totalAmount"`
	FeeCoin     string       `json:"feeCoin"`
	TaxFeeItems []TaxFeeItem `json:"taxFeeItems"`
}

type TaxFeeItem struct {
	FeeType   string `json:"feeType"`
	FeeRate   string `json:"feeRate"`
	FeeAmount string `json:"feeAmount"`
	Region    string `json:"region"`
	FeeCoin   string `json:"feeCoin"`
}

type QuerySmallAssetListResponse struct {
	SmallAssetCoins []SmallAssetCoin `json:"smallAssetCoins"`
	SupportToCoins  []string         `json:"supportToCoins"`
}

type SmallAssetCoin struct {
	FromCoin         string      `json:"fromCoin"`
	SupportConvert   int         `json:"supportConvert"`
	AvailableBalance string      `json:"availableBalance"`
	BaseValue        string      `json:"baseValue"`
	ToCoin           string      `json:"toCoin"`
	ToAmount         string      `json:"toAmount"`
	ExchangeRate     string      `json:"exchangeRate"`
	FeeInfo          interface{} `json:"feeInfo"`
	TaxFeeInfo       interface{} `json:"taxFeeInfo"`
}

type SmallAssetQuoteResponse struct {
	QuoteId string      `json:"quoteId"`
	Result  interface{} `json:"result"`
}

type SmallAssetQuoteResult struct {
	QuoteCreateTime string           `json:"quoteCreateTime"`
	QuoteExpireTime string           `json:"quoteExpireTime"`
	ExchangeCoins   []SmallAssetCoin `json:"exchangeCoins"`
	TotalFeeInfo    interface{}      `json:"totalFeeInfo"`
	TotalTaxFeeInfo interface{}      `json:"totalTaxFeeInfo"`
}

type FeeInfo struct {
	FeeCoin string `json:"feeCoin"`
	Amount  string `json:"amount"`
	FeeRate string `json:"feeRate"`
}

type AccountRepayResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type TradeAssetListResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type TradeUserAssetDto struct {
	ChainCode            string `json:"chainCode"`
	ChainIconUrl         string `json:"chainIconUrl"`
	TokenAddress         string `json:"tokenAddress"`
	TokenCode            string `json:"tokenCode"`
	TokenSymbol          string `json:"tokenSymbol"`
	TokenDecimals        int    `json:"tokenDecimals"`
	TokenIconUrlDay      string `json:"tokenIconUrlDay"`
	TokenIconUrlNight    string `json:"tokenIconUrlNight"`
	TokenAmount          string `json:"tokenAmount"`
	TokenAmountUsd       string `json:"tokenAmountUsd"`
	TradeFlag            int    `json:"tradeFlag"`
	Pnl                  string `json:"pnl"`
	PnlRatio             string `json:"pnlRatio"`
	CostPrice            string `json:"costPrice"`
	LastPrice            string `json:"lastPrice"`
	CostTotalValue       string `json:"costTotalValue"`
	AssetStatus          int    `json:"assetStatus"`
	AnnouncementUrl      string `json:"announcementUrl"`
	EstimatedOfflineTime int    `json:"estimatedOfflineTime"`
	DelistingTime        int    `json:"delistingTime"`
}

type Good struct {
	ShoppingName string `json:"shoppingName"`
	MccCode      string `json:"mccCode"`
	GoodsName    string `json:"goodsName"`
	GoodsDetail  string `json:"goodsDetail"`
}

type RiskInfo struct {
	TerminalType string `json:"terminalType"`
}

type RefundOrderItem struct {
	RefundType       string                  `json:"refundType"`
	MerchantTradeNo  string                  `json:"merchantTradeNo"`
	PayId            string                  `json:"payId"`
	MerchantRefundNo string                  `json:"merchantRefundNo"`
	RefundAmount     string                  `json:"refundAmount"`
	Env              RefundOrderItemEnv      `json:"env"`
	RiskInfo         RiskInfo                `json:"riskInfo"`
	Customer         RefundOrderItemCustomer `json:"customer"`
}

type RefundOrderItemEnv struct {
}

type RefundOrderItemCustomer struct {
}

type AgreementLimitConfig struct {
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	CurrencyType string `json:"currency_type"`
	Chain        string `json:"chain"`
}

type AgreementPeriodLimit struct {
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	CurrencyType string `json:"currency_type"`
	Chain        string `json:"chain"`
	PeriodType   string `json:"period_type"`
}

type AgreementCryptoPaymentInfo struct {
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Chain        string `json:"chain"`
	ExchangeRate string `json:"exchange_rate"`
	RateTime     string `json:"rate_time"`
}

type AgreementOrderInfo struct {
	OrderTitle    string `json:"order_title"`
	OrderDesc     string `json:"order_desc"`
	GoodsName     string `json:"goods_name"`
	GoodsId       string `json:"goods_id"`
	GoodsCategory string `json:"goods_category"`
}

type AgreementSceneInfo struct {
	DeviceId string                     `json:"device_id"`
	DeviceIp string                     `json:"device_ip"`
	Location AgreementSceneInfoLocation `json:"location"`
}

type AgreementSceneInfoLocation struct {
}

type AgreementRiskInfo struct {
	UserIp            string `json:"user_ip"`
	DeviceFingerprint string `json:"device_fingerprint"`
	UserAgent         string `json:"user_agent"`
}

type AgreementSignParams struct {
	MerchantUserId      string                         `json:"merchant_user_id"`
	SceneCode           string                         `json:"scene_code"`
	ProductCode         string                         `json:"product_code"`
	ExternalAgreementNo string                         `json:"external_agreement_no"`
	SignValidTime       string                         `json:"sign_valid_time"`
	SingleLimit         AgreementLimitConfig           `json:"single_limit"`
	PeriodLimits        []AgreementPeriodLimit         `json:"period_limits"`
	SignNotifyUrl       string                         `json:"sign_notify_url"`
	ReturnUrl           string                         `json:"return_url"`
	SignExpireMinutes   int                            `json:"sign_expire_minutes"`
	ExtraParams         AgreementSignParamsExtraParams `json:"extra_params"`
}

type AgreementSignParamsExtraParams struct {
}

type AgreementPayParams struct {
	AgreementNo  string                     `json:"agreement_no"`
	OutTradeNo   string                     `json:"out_trade_no"`
	SceneCode    string                     `json:"scene_code"`
	Amount       AgreementCryptoPaymentInfo `json:"amount"`
	OrderInfo    AgreementOrderInfo         `json:"order_info"`
	SceneInfo    AgreementSceneInfo         `json:"scene_info"`
	PayNotifyUrl string                     `json:"pay_notify_url"`
	RiskInfo     AgreementRiskInfo          `json:"risk_info"`
}

type GetAdsRequest struct {
	TokenId    string `json:"tokenId"`
	CurrencyId string `json:"currencyId"`
	Side       string `json:"side"`
	Page       string `json:"page"`
	Size       string `json:"size"`
}

type PostAdRequest struct {
	TokenId              string                            `json:"tokenId"`
	CurrencyId           string                            `json:"currencyId"`
	Side                 string                            `json:"side"`
	PriceType            string                            `json:"priceType"`
	Premium              string                            `json:"premium"`
	Price                string                            `json:"price"`
	MinAmount            string                            `json:"minAmount"`
	MaxAmount            string                            `json:"maxAmount"`
	Remark               string                            `json:"remark"`
	TradingPreferenceSet PostAdRequestTradingPreferenceSet `json:"tradingPreferenceSet"`
	PaymentIds           []string                          `json:"paymentIds"`
	Quantity             string                            `json:"quantity"`
	PaymentPeriod        string                            `json:"paymentPeriod"`
	ItemType             string                            `json:"itemType"`
}

type PostAdRequestTradingPreferenceSet struct {
}

type RemoveAdRequest struct {
	ItemId string `json:"itemId"`
}

type UpdateAdRequest struct {
	Id                   string                              `json:"id"`
	PriceType            string                              `json:"priceType"`
	Premium              string                              `json:"premium"`
	Price                string                              `json:"price"`
	MinAmount            string                              `json:"minAmount"`
	MaxAmount            string                              `json:"maxAmount"`
	Remark               string                              `json:"remark"`
	TradingPreferenceSet UpdateAdRequestTradingPreferenceSet `json:"tradingPreferenceSet"`
	PaymentIds           []string                            `json:"paymentIds"`
	ActionType           string                              `json:"actionType"`
	Quantity             string                              `json:"quantity"`
	PaymentPeriod        string                              `json:"paymentPeriod"`
}

type UpdateAdRequestTradingPreferenceSet struct {
}

type GetMyAdsRequest struct {
	ItemId     string `json:"itemId"`
	Status     string `json:"status"`
	Side       string `json:"side"`
	TokenId    string `json:"tokenId"`
	Page       string `json:"page"`
	Size       string `json:"size"`
	CurrencyId string `json:"currencyId"`
}

type GetAdDetailRequest struct {
	ItemId string `json:"itemId"`
}

type GetAllOrdersRequest struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Status    int    `json:"status"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
	TokenId   string `json:"tokenId"`
	Side      int    `json:"side"`
}

type GetPendingOrdersRequest struct {
	Status    int    `json:"status"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
	TokenId   string `json:"tokenId"`
	Side      int    `json:"side"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

type GetOrderDetailRequest struct {
	OrderId string `json:"orderId"`
}

type MarkOrderPaidRequest struct {
	OrderId     string `json:"orderId"`
	PaymentType string `json:"paymentType"`
	PaymentId   string `json:"paymentId"`
}

type ReleaseAssetsRequest struct {
	OrderId string `json:"orderId"`
}

type SendChatMessageRequest struct {
	Message     string `json:"message"`
	ContentType string `json:"contentType"`
	OrderId     string `json:"orderId"`
	MsgUuid     string `json:"msgUuid"`
	FileName    string `json:"fileName"`
}

type GetChatMessageRequest struct {
	OrderId     string `json:"orderId"`
	CurrentPage string `json:"currentPage"`
	Size        string `json:"size"`
}

type GetCounterpartyInfoRequest struct {
	OriginalUid string `json:"originalUid"`
	OrderId     string `json:"orderId"`
}
