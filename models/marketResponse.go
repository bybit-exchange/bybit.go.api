package models

type GetServerTimeResponse struct {
	RetCode    int              `json:"retCode"`
	RetMsg     string           `json:"retMsg"`
	Result     ServerTimeResult `json:"result"`
	RetExtInfo struct{}         `json:"retExtInfo"`
	Time       int64            `json:"time"`
}

type ServerTimeResult struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

type MarketKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
	Volume     string `json:"volume"`
	Turnover   string `json:"turnover"`
}

type MarketKlineResponse struct {
	Category Category             `json:"category"`
	Symbol   string               `json:"symbol"`
	List     []*MarketKlineCandle `json:"list"`
}

type MarketMarkPriceKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
}

type MarketMarkPriceKlineResponse struct {
	Category Category                      `json:"category"`
	Symbol   string                        `json:"symbol"`
	List     []*MarketMarkPriceKlineCandle `json:"list"`
}

type MarketIndexPriceKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
}

type MarketIndexPriceKlineResponse struct {
	Category Category                       `json:"category"`
	Symbol   string                         `json:"symbol"`
	List     []*MarketIndexPriceKlineCandle `json:"list"`
}

type MarketPremiumIndexPriceKlineCandle struct {
	StartTime  string `json:"startTime"`
	OpenPrice  string `json:"openPrice"`
	HighPrice  string `json:"highPrice"`
	LowPrice   string `json:"lowPrice"`
	ClosePrice string `json:"closePrice"`
}

type MarketPremiumIndexPriceKlineResponse struct {
	Category Category                              `json:"category"`
	Symbol   string                                `json:"symbol"`
	List     []*MarketPremiumIndexPriceKlineCandle `json:"list"`
}

type InstrumentInfoResponse struct {
	Category       Category      `json:"category"`
	NextPageCursor string        `json:"nextPageCursor"`
	List           []interface{} `json:"list"`
}

type SpotInstrument struct {
	Symbol             string         `json:"symbol"`
	ContractType       string         `json:"contractType"`
	OptionType         string         `json:"optionType"`
	Innovation         string         `json:"innovation"`
	Status             SymbolStatus   `json:"status"`
	BaseCoin           string         `json:"baseCoin"`
	QuoteCoin          string         `json:"quoteCoin"`
	LaunchTime         string         `json:"launchTime"`
	DeliveryTime       string         `json:"deliveryTime"`
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`
	PriceScale         string         `json:"priceScale"`
	MarginTrading      string         `json:"marginTrading"`
	LeverageFilter     LeverageFilter `json:"leverageFilter"`
	PriceFilter        PriceFilter    `json:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `json:"lotSizeFilter"`
	UnifiedMarginTrade bool           `json:"unifiedMarginTrade"`
	FundingInterval    int            `json:"fundingInterval"`
	SettleCoin         string         `json:"settleCoin"`
	CopyTrading        string         `json:"copyTrading"`
}

type LeverageFilter struct {
	MinLeverage  string `json:"minLeverage"`
	MaxLeverage  string `json:"maxLeverage"`
	LeverageStep string `json:"leverageStep"`
}

type PriceFilter struct {
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`
}

type LotSizeFilter struct {
	MaxOrderQty         string `json:"maxOrderQty"`
	MinOrderQty         string `json:"minOrderQty"`
	QtyStep             string `json:"qtyStep"`
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
	BasePrecision       string `json:"basePrecision"`
	QuotePrecision      string `json:"quotePrecision"`
	MaxOrderAmt         string `json:"maxOrderAmt"`
	MinOrderAmt         string `json:"minOrderAmt"`
}

type MarketOrderBookResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     OrderBookInfo `json:"result"`
	RetExtInfo struct{}      `json:"retExtInfo"`
	Time       int64         `json:"time"`
}

// type OrderBookEntry struct {
// Price string `json:"0"`
// Size  string `json:"1"`
// }

type OrderBookEntry []string

type OrderBookInfo struct {
	Symbol    string           `json:"s"`
	Bids      []OrderBookEntry `json:"b"`
	Asks      []OrderBookEntry `json:"a"`
	Timestamp int64            `json:"ts"`
	UpdateID  int64            `json:"u"`
}

type TickerInfo struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice"`
	MarkPrice              string `json:"markPrice"`
	PrevPrice24h           string `json:"prevPrice24h"`
	Price24hPcnt           string `json:"price24hPcnt"`
	HighPrice24h           string `json:"highPrice24h"`
	LowPrice24h            string `json:"lowPrice24h"`
	PrevPrice1h            string `json:"prevPrice1h"`
	OpenInterest           string `json:"openInterest"`
	OpenInterestValue      string `json:"openInterestValue"`
	Turnover24h            string `json:"turnover24h"`
	Volume24h              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate"`
	NextFundingTime        string `json:"nextFundingTime"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
	BasisRate              string `json:"basisRate"`
	Basis                  string `json:"basis"`
	DeliveryFeeRate        string `json:"deliveryFeeRate"`
	DeliveryTime           string `json:"deliveryTime"`
	Ask1Size               string `json:"ask1Size"`
	Bid1Price              string `json:"bid1Price"`
	Ask1Price              string `json:"ask1Price"`
	Bid1Size               string `json:"bid1Size"`
	Ask1Iv                 string `json:"ask1Iv"`
	Bid1Iv                 string `json:"bid1Iv"`
	MarkIv                 string `json:"markIv"`
	UnderlyingPrice        string `json:"underlyingPrice"`
	TotalVolume            string `json:"totalVolume"`
	TotalTurnover          string `json:"totalTurnover"`
	Change24h              string `json:"change24h"`
	UsdIndexPrice          string `json:"usdIndexPrice"`
}

type MarketTickersResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     MarketTickers `json:"result"`
	RetExtInfo struct{}      `json:"retExtInfo"`
	Time       int64         `json:"time"`
}

type MarketTickers struct {
	Category string        `json:"category"`
	List     []*TickerInfo `json:"list"`
}

type MarketFundingRatesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     FundingRate `json:"result"`
	RetExtInfo struct{}    `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type FundingRateInfo struct {
	Symbol               string `json:"symbol"`
	FundingRate          string `json:"fundingRate"`
	FundingRateTimestamp string `json:"fundingRateTimestamp"`
}

type FundingRate struct {
	Category string            `json:"category"`
	List     []FundingRateInfo `json:"list"`
}

type TradeInfo struct {
	ExecId       string `json:"execId"`
	Symbol       string `json:"symbol"`
	Price        string `json:"price"`
	Size         string `json:"size"`
	Side         string `json:"side"`
	Time         string `json:"time"`
	IsBlockTrade bool   `json:"isBlockTrade"`
}

type PublicRecentTradeHistory struct {
	Category string      `json:"category"`
	List     []TradeInfo `json:"list"`
}

type GetPublicRecentTradesResponse struct {
	RetCode    int                      `json:"retCode"`
	RetMsg     string                   `json:"retMsg"`
	Result     PublicRecentTradeHistory `json:"result"`
	RetExtInfo struct{}                 `json:"retExtInfo"`
	Time       int64                    `json:"time"`
}

type GetOpenInterestsResponse struct {
	RetCode    int              `json:"retCode"`
	RetMsg     string           `json:"retMsg"`
	Result     OpenInterestInfo `json:"result"`
	RetExtInfo struct{}         `json:"retExtInfo"`
	Time       int64            `json:"time"`
}

type OpenInterestInfo struct {
	Category       string         `json:"category"`
	Symbol         string         `json:"symbol"`
	List           []OpenInterest `json:"list"`
	NextPageCursor string         `json:"nextPageCursor"`
}

type OpenInterest struct {
	OpenInterest string `json:"openInterest"`
	Timestamp    string `json:"timeStamp"`
}

type VolatilityData struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

type HistoricalVolatilityInfo struct {
	Category string           `json:"category"`
	List     []VolatilityData `json:"result"`
}

type GetInsuranceInfoResponse struct {
	RetCode    int                 `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	Result     MarketInsuranceInfo `json:"result"`
	RetExtInfo struct{}            `json:"retExtInfo"`
	Time       int64               `json:"time"`
}

type InsuranceData struct {
	Coin    string `json:"coin"`
	Balance string `json:"balance"`
	Value   string `json:"value"`
}

type MarketInsuranceInfo struct {
	UpdatedTime string          `json:"updatedTime"`
	List        []InsuranceData `json:"list"`
}

type GetRiskLimitResponse struct {
	RetCode    int                 `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	Result     MarketRiskLimitInfo `json:"result"`
	RetExtInfo struct{}            `json:"retExtInfo"`
	Time       int64               `json:"time"`
}

type RiskLimitData struct {
	Id                int    `json:"id"`
	Symbol            string `json:"symbol"`
	RiskLimitValue    string `json:"riskLimitValue"`
	MaintenanceMargin string `json:"maintenanceMargin"`
	InitialMargin     string `json:"initialMargin"`
	IsLowestRisk      int    `json:"isLowestRisk"`
	MaxLeverage       string `json:"maxLeverage"`
}

type MarketRiskLimitInfo struct {
	Category string          `json:"category"`
	List     []RiskLimitData `json:"list"`
}

type GetDeliveryPriceResponse struct {
	RetCode    int               `json:"retCode"`
	RetMsg     string            `json:"retMsg"`
	Result     DeliveryPriceInfo `json:"result"`
	RetExtInfo struct{}          `json:"retExtInfo"`
	Time       int64             `json:"time"`
}

type DeliveryPriceData struct {
	Symbol        string `json:"symbol"`
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryTime  string `json:"deliveryTime"`
}

type DeliveryPriceInfo struct {
	Category       string              `json:"category"`
	List           []DeliveryPriceData `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type GetMarketLSRatioResponse struct {
	RetCode    int                      `json:"retCode"`
	RetMsg     string                   `json:"retMsg"`
	Result     MarketLongShortRatioInfo `json:"result"`
	RetExtInfo struct{}                 `json:"retExtInfo"`
	Time       int64                    `json:"time"`
}

type LongShortRatioData struct {
	Symbol    string `json:"symbol"`
	BuyRatio  string `json:"buyRatio"`
	SellRatio string `json:"sellRatio"`
	Timestamp string `json:"timestamp"`
}

type MarketLongShortRatioInfo struct {
	List []LongShortRatioData `json:"list"`
}

type AdlAlertResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type AdlAlertResult struct {
	UpdatedTime string           `json:"updatedTime"`
	List        []AdlAlertRecord `json:"list"`
}

type AdlAlertRecord struct {
	Coin                string `json:"coin"`
	Symbol              string `json:"symbol"`
	Balance             string `json:"balance"`
	MaxBalance          string `json:"maxBalance"`
	InsurancePnlRatio   string `json:"insurancePnlRatio"`
	PnlRatio            string `json:"pnlRatio"`
	AdlTriggerThreshold string `json:"adlTriggerThreshold"`
	AdlStopRatio        string `json:"adlStopRatio"`
}

type DeliveryPriceResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type DeliveryPriceResult struct {
	Category       string                `json:"category"`
	NextPageCursor string                `json:"nextPageCursor"`
	List           []DeliveryPriceRecord `json:"list"`
}

type DeliveryPriceRecord struct {
	Symbol        string `json:"symbol"`
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryTime  string `json:"deliveryTime"`
}

type FeeGroupInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type FeeGroupInfoResult struct {
	List []FeeGroup `json:"list"`
}

type FeeGroup struct {
	GroupName       string      `json:"groupName"`
	WeightingFactor int         `json:"weightingFactor"`
	SymbolsNumbers  int         `json:"symbolsNumbers"`
	Symbols         []string    `json:"symbols"`
	FeeRates        interface{} `json:"feeRates"`
	UpdateTime      string      `json:"updateTime"`
}

type FeeRates struct {
	Pro         []FeeRateLevel `json:"pro"`
	MarketMaker []FeeRateLevel `json:"marketMaker"`
}

type FeeRateLevel struct {
	Level        string `json:"level"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
	MakerRebate  string `json:"makerRebate"`
}

type FundingRateHistoryResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type FundingRateHistoryResult struct {
	Category string              `json:"category"`
	List     []FundingRateRecord `json:"list"`
}

type FundingRateRecord struct {
	Symbol               string `json:"symbol"`
	FundingRate          string `json:"fundingRate"`
	FundingRateTimestamp string `json:"fundingRateTimestamp"`
}

type HistoricalVolatilityResponse struct {
	RetCode    int                `json:"retCode"`
	RetMsg     string             `json:"retMsg"`
	RetExtInfo interface{}        `json:"retExtInfo"`
	Time       int                `json:"time"`
	Category   string             `json:"category"`
	Result     []VolatilityRecord `json:"result"`
}

type VolatilityRecord struct {
	Period int    `json:"period"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

type IndexComponentsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type IndexComponentsResult struct {
	IndexName  string           `json:"indexName"`
	LastPrice  string           `json:"lastPrice"`
	UpdateTime string           `json:"updateTime"`
	Components []IndexComponent `json:"components"`
}

type IndexComponent struct {
	Exchange        string `json:"exchange"`
	SpotPair        string `json:"spotPair"`
	EquivalentPrice string `json:"equivalentPrice"`
	Multiplier      string `json:"multiplier"`
	Price           string `json:"price"`
	Weight          string `json:"weight"`
}

type IndexPriceKlineResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type IndexPriceKlineResult struct {
	Category string          `json:"category"`
	Symbol   string          `json:"symbol"`
	List     [][]interface{} `json:"list"`
}

type IndexKlineEntry struct {
	Items []string `json:"items"`
}

type BaseResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type LotSizeFilterLinear struct {
	MaxOrderQty         string `json:"maxOrderQty"`
	MinOrderQty         string `json:"minOrderQty"`
	QtyStep             string `json:"qtyStep"`
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
	MaxMktOrderQty      string `json:"maxMktOrderQty"`
	MinNotionalValue    string `json:"minNotionalValue"`
}

type LotSizeFilterSpot struct {
	BasePrecision             string `json:"basePrecision"`
	QuotePrecision            string `json:"quotePrecision"`
	MinOrderAmt               string `json:"minOrderAmt"`
	MaxOrderAmt               string `json:"maxOrderAmt"`
	MaxLimitOrderQty          string `json:"maxLimitOrderQty"`
	MaxMarketOrderQty         string `json:"maxMarketOrderQty"`
	PostOnlyMaxLimitOrderSize string `json:"postOnlyMaxLimitOrderSize"`
}

type RiskParameters struct {
	PriceLimitRatioX string `json:"priceLimitRatioX"`
	PriceLimitRatioY string `json:"priceLimitRatioY"`
}

type InstrumentLinearInverse struct {
	Symbol             string      `json:"symbol"`
	ContractType       string      `json:"contractType"`
	Status             string      `json:"status"`
	BaseCoin           string      `json:"baseCoin"`
	QuoteCoin          string      `json:"quoteCoin"`
	LaunchTime         string      `json:"launchTime"`
	DeliveryTime       string      `json:"deliveryTime"`
	DeliveryFeeRate    string      `json:"deliveryFeeRate"`
	PriceScale         string      `json:"priceScale"`
	LeverageFilter     interface{} `json:"leverageFilter"`
	PriceFilter        interface{} `json:"priceFilter"`
	LotSizeFilter      interface{} `json:"lotSizeFilter"`
	UnifiedMarginTrade bool        `json:"unifiedMarginTrade"`
	FundingInterval    int         `json:"fundingInterval"`
	SettleCoin         string      `json:"settleCoin"`
	CopyTrading        string      `json:"copyTrading"`
	UpperFundingRate   string      `json:"upperFundingRate"`
	LowerFundingRate   string      `json:"lowerFundingRate"`
	IsPreListing       bool        `json:"isPreListing"`
	PreListingInfo     interface{} `json:"preListingInfo"`
	RiskParameters     interface{} `json:"riskParameters"`
}

type InstrumentSpot struct {
	Symbol         string      `json:"symbol"`
	BaseCoin       string      `json:"baseCoin"`
	QuoteCoin      string      `json:"quoteCoin"`
	Status         string      `json:"status"`
	MarginTrading  string      `json:"marginTrading"`
	StTag          string      `json:"stTag"`
	LotSizeFilter  interface{} `json:"lotSizeFilter"`
	PriceFilter    interface{} `json:"priceFilter"`
	RiskParameters interface{} `json:"riskParameters"`
}

type InstrumentOption struct {
	Symbol        string      `json:"symbol"`
	OptionsType   string      `json:"optionsType"`
	Status        string      `json:"status"`
	BaseCoin      string      `json:"baseCoin"`
	QuoteCoin     string      `json:"quoteCoin"`
	SettleCoin    string      `json:"settleCoin"`
	LaunchTime    string      `json:"launchTime"`
	DeliveryTime  string      `json:"deliveryTime"`
	PriceFilter   interface{} `json:"priceFilter"`
	LotSizeFilter interface{} `json:"lotSizeFilter"`
}

type InstrumentsInfoResult struct {
	Category       string        `json:"category"`
	NextPageCursor string        `json:"nextPageCursor"`
	List           []interface{} `json:"list"`
}

type InstrumentsInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type InsuranceResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type InsuranceResult struct {
	UpdatedTime string            `json:"updatedTime"`
	List        []InsuranceRecord `json:"list"`
}

type InsuranceRecord struct {
	Coin    string `json:"coin"`
	Symbols string `json:"symbols"`
	Balance string `json:"balance"`
	Value   string `json:"value"`
}

type KlineResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type KlineResult struct {
	Category string       `json:"category"`
	Symbol   string       `json:"symbol"`
	List     []KlineEntry `json:"list"`
}

type KlineEntry struct {
	Items []string `json:"items"`
}

type LongShortRatioResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type LongShortRatioResult struct {
	List           []LongShortRatioRecord `json:"list"`
	NextPageCursor string                 `json:"nextPageCursor"`
}

type LongShortRatioRecord struct {
	Symbol    string `json:"symbol"`
	BuyRatio  string `json:"buyRatio"`
	SellRatio string `json:"sellRatio"`
	Timestamp string `json:"timestamp"`
}

type MarkPriceKlineResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type MarkPriceKlineResult struct {
	Category string          `json:"category"`
	Symbol   string          `json:"symbol"`
	List     [][]interface{} `json:"list"`
}

type MarkKlineEntry struct {
	Items string `json:"items"`
}

type NewDeliveryPriceResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type NewDeliveryPriceResult struct {
	Category string                   `json:"category"`
	List     []NewDeliveryPriceRecord `json:"list"`
}

type NewDeliveryPriceRecord struct {
	DeliveryPrice string `json:"deliveryPrice"`
	DeliveryTime  string `json:"deliveryTime"`
}

type OpenInterestResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type OpenInterestResult struct {
	Category       string               `json:"category"`
	Symbol         string               `json:"symbol"`
	List           []OpenInterestRecord `json:"list"`
	NextPageCursor string               `json:"nextPageCursor"`
}

type OpenInterestRecord struct {
	OpenInterest string `json:"openInterest"`
	Timestamp    string `json:"timestamp"`
}

type OrderPriceLimitResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type OrderPriceLimitResult struct {
	Symbol  string `json:"symbol"`
	BuyLmt  string `json:"buyLmt"`
	SellLmt string `json:"sellLmt"`
	Ts      string `json:"ts"`
}

type OrderbookResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type OrderbookResult struct {
	S   string          `json:"s"`
	B   [][]interface{} `json:"b"`
	A   [][]interface{} `json:"a"`
	Ts  int             `json:"ts"`
	U   int             `json:"u"`
	Seq int             `json:"seq"`
	Cts int             `json:"cts"`
}

type OrderbookLevel struct {
	Items []string `json:"items"`
}

type PremiumIndexKlineResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type PremiumIndexKlineResult struct {
	Category string          `json:"category"`
	Symbol   string          `json:"symbol"`
	List     [][]interface{} `json:"list"`
}

type PremiumIndexKlineEntry struct {
	Items []string `json:"items"`
}

type RecentTradeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RecentTradeResult struct {
	Category string        `json:"category"`
	List     []TradeRecord `json:"list"`
}

type TradeRecord struct {
	ExecId       string `json:"execId"`
	Symbol       string `json:"symbol"`
	Price        string `json:"price"`
	Size         string `json:"size"`
	Side         string `json:"side"`
	Time         string `json:"time"`
	IsBlockTrade bool   `json:"isBlockTrade"`
	IsRPITrade   bool   `json:"isRPITrade"`
	Seq          string `json:"seq"`
	MP           string `json:"mP"`
	IP           string `json:"iP"`
	MIv          string `json:"mIv"`
	Iv           string `json:"iv"`
}

type RiskLimitResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RiskLimitResult struct {
	Category       string          `json:"category"`
	List           []RiskLimitTier `json:"list"`
	NextPageCursor string          `json:"nextPageCursor"`
}

type RiskLimitTier struct {
	Id                int    `json:"id"`
	Symbol            string `json:"symbol"`
	RiskLimitValue    string `json:"riskLimitValue"`
	MaintenanceMargin string `json:"maintenanceMargin"`
	InitialMargin     string `json:"initialMargin"`
	IsLowestRisk      int    `json:"isLowestRisk"`
	MaxLeverage       string `json:"maxLeverage"`
	MmDeduction       string `json:"mmDeduction"`
}

type RpiOrderbookResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type RpiOrderbookResult struct {
	S   string          `json:"s"`
	B   [][]interface{} `json:"b"`
	A   [][]interface{} `json:"a"`
	Ts  int             `json:"ts"`
	U   int             `json:"u"`
	Seq int             `json:"seq"`
	Cts int             `json:"cts"`
}

type RpiOrderbookLevel struct {
	Items []string `json:"items"`
}

type TickerLinearInverse struct {
	Symbol              string `json:"symbol"`
	LastPrice           string `json:"lastPrice"`
	IndexPrice          string `json:"indexPrice"`
	MarkPrice           string `json:"markPrice"`
	PrevPrice24h        string `json:"prevPrice24h"`
	Price24hPcnt        string `json:"price24hPcnt"`
	HighPrice24h        string `json:"highPrice24h"`
	LowPrice24h         string `json:"lowPrice24h"`
	PrevPrice1h         string `json:"prevPrice1h"`
	OpenInterest        string `json:"openInterest"`
	OpenInterestValue   string `json:"openInterestValue"`
	Turnover24h         string `json:"turnover24h"`
	Volume24h           string `json:"volume24h"`
	FundingRate         string `json:"fundingRate"`
	NextFundingTime     string `json:"nextFundingTime"`
	Bid1Price           string `json:"bid1Price"`
	Bid1Size            string `json:"bid1Size"`
	Ask1Price           string `json:"ask1Price"`
	Ask1Size            string `json:"ask1Size"`
	FundingIntervalHour string `json:"fundingIntervalHour"`
	FundingCap          string `json:"fundingCap"`
}

type TickerSpot struct {
	Symbol        string `json:"symbol"`
	Bid1Price     string `json:"bid1Price"`
	Bid1Size      string `json:"bid1Size"`
	Ask1Price     string `json:"ask1Price"`
	Ask1Size      string `json:"ask1Size"`
	LastPrice     string `json:"lastPrice"`
	PrevPrice24h  string `json:"prevPrice24h"`
	Price24hPcnt  string `json:"price24hPcnt"`
	HighPrice24h  string `json:"highPrice24h"`
	LowPrice24h   string `json:"lowPrice24h"`
	Turnover24h   string `json:"turnover24h"`
	Volume24h     string `json:"volume24h"`
	UsdIndexPrice string `json:"usdIndexPrice"`
}

type TickerOption struct {
	Symbol       string `json:"symbol"`
	Bid1Price    string `json:"bid1Price"`
	Bid1Size     string `json:"bid1Size"`
	Bid1Iv       string `json:"bid1Iv"`
	Ask1Price    string `json:"ask1Price"`
	Ask1Size     string `json:"ask1Size"`
	Ask1Iv       string `json:"ask1Iv"`
	LastPrice    string `json:"lastPrice"`
	HighPrice24h string `json:"highPrice24h"`
	LowPrice24h  string `json:"lowPrice24h"`
	MarkPrice    string `json:"markPrice"`
	IndexPrice   string `json:"indexPrice"`
	MarkIv       string `json:"markIv"`
	OpenInterest string `json:"openInterest"`
	Turnover24h  string `json:"turnover24h"`
	Volume24h    string `json:"volume24h"`
	Delta        string `json:"delta"`
	Gamma        string `json:"gamma"`
	Vega         string `json:"vega"`
	Theta        string `json:"theta"`
}

type TickersResult struct {
	Category string        `json:"category"`
	List     []interface{} `json:"list"`
}

type TickersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type ServerTimeResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SpreadInstrumentsInfoResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SpreadInstrumentsInfoResult struct {
	List           []SpreadInstrument `json:"list"`
	NextPageCursor string             `json:"nextPageCursor"`
}

type SpreadInstrument struct {
	Symbol       string      `json:"symbol"`
	ContractType string      `json:"contractType"`
	Status       string      `json:"status"`
	BaseCoin     string      `json:"baseCoin"`
	QuoteCoin    string      `json:"quoteCoin"`
	SettleCoin   string      `json:"settleCoin"`
	TickSize     string      `json:"tickSize"`
	MinPrice     string      `json:"minPrice"`
	MaxPrice     string      `json:"maxPrice"`
	LotSize      string      `json:"lotSize"`
	MinSize      string      `json:"minSize"`
	MaxSize      string      `json:"maxSize"`
	LaunchTime   string      `json:"launchTime"`
	DeliveryTime string      `json:"deliveryTime"`
	Legs         []SpreadLeg `json:"legs"`
}

type SpreadLeg struct {
	Symbol       string `json:"symbol"`
	ContractType string `json:"contractType"`
}

type SpreadOrderbookResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SpreadOrderbookResult struct {
	S   string          `json:"s"`
	B   [][]interface{} `json:"b"`
	A   [][]interface{} `json:"a"`
	Ts  int             `json:"ts"`
	U   int             `json:"u"`
	Seq int             `json:"seq"`
	Cts int             `json:"cts"`
}

type SpreadRecentTradesResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SpreadRecentTradesResult struct {
	List []SpreadTradeRecord `json:"list"`
}

type SpreadTradeRecord struct {
	ExecId string `json:"execId"`
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Size   string `json:"size"`
	Side   string `json:"side"`
	Time   string `json:"time"`
	Seq    string `json:"seq"`
}

type SpreadTickersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
	Result     interface{} `json:"result"`
}

type SpreadTickersResult struct {
	List []SpreadTicker `json:"list"`
}

type SpreadTicker struct {
	Symbol       string `json:"symbol"`
	BidPrice     string `json:"bidPrice"`
	BidSize      string `json:"bidSize"`
	AskPrice     string `json:"askPrice"`
	AskSize      string `json:"askSize"`
	LastPrice    string `json:"lastPrice"`
	HighPrice24h string `json:"highPrice24h"`
	LowPrice24h  string `json:"lowPrice24h"`
	PrevPrice24h string `json:"prevPrice24h"`
	Volume24h    string `json:"volume24h"`
}
