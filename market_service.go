package bybit_connector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api/models"
)

// MarketKlinesService Market Kline (GET /v5/market/kline)
type MarketKlinesService struct {
	c        *Client
	category *models.Category
	symbol   string
	interval string
	start    *uint64
	end      *uint64
	limit    *int
}

// Category set category
func (s *MarketKlinesService) Category(category models.Category) *MarketKlinesService {
	s.category = &category
	return s
}

// Symbol set symbol
func (s *MarketKlinesService) Symbol(symbol string) *MarketKlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *MarketKlinesService) Interval(interval string) *MarketKlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *MarketKlinesService) Limit(limit int) *MarketKlinesService {
	s.limit = &limit
	return s
}

// Start set startTime
func (s *MarketKlinesService) Start(startTime uint64) *MarketKlinesService {
	s.start = &startTime
	return s
}

// End set endTime
func (s *MarketKlinesService) End(endTime uint64) *MarketKlinesService {
	s.end = &endTime
	return s
}

func (s *MarketKlinesService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketKlineResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/kline",
		secType:  secTypeNone,
	}
	if s.category != nil {
		r.setParam("category", *s.category)
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.start != nil {
		r.setParam("start", *s.start)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	j, err := newJSON(data)
	if err != nil {
		return nil, err
	}
	result := j.Get("result")
	res = new(models.MarketKlineResponse)
	res.Category = models.Category(result.Get("category").MustString())
	res.Symbol = result.Get("symbol").MustString()
	list := result.Get("list").MustArray()
	res.List = make([]*models.MarketKlineCandle, len(list))
	for i := range list {
		item := result.Get("list").GetIndex(i)
		if len(item.MustArray()) < 7 {
			return nil, fmt.Errorf("invalid kline response")
		}

		res.List[i] = &models.MarketKlineCandle{
			StartTime:  item.GetIndex(0).MustString(),
			OpenPrice:  item.GetIndex(1).MustString(),
			HighPrice:  item.GetIndex(2).MustString(),
			LowPrice:   item.GetIndex(3).MustString(),
			ClosePrice: item.GetIndex(4).MustString(),
			Volume:     item.GetIndex(5).MustString(),
			Turnover:   item.GetIndex(6).MustString(),
		}
	}

	return res, nil
}

// MarketMarkPriceKlineService Market mark price kline (GET /v5/market/mark-price-kline)
type MarketMarkPriceKlineService struct {
	c        *Client
	category *models.Category
	symbol   string
	interval string
	start    *uint64
	end      *uint64
	limit    *int
}

// Category set category
func (s *MarketMarkPriceKlineService) Category(category models.Category) *MarketMarkPriceKlineService {
	s.category = &category
	return s
}

// Symbol set symbol
func (s *MarketMarkPriceKlineService) Symbol(symbol string) *MarketMarkPriceKlineService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *MarketMarkPriceKlineService) Interval(interval string) *MarketMarkPriceKlineService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *MarketMarkPriceKlineService) Limit(limit int) *MarketMarkPriceKlineService {
	s.limit = &limit
	return s
}

// Start set startTime
func (s *MarketMarkPriceKlineService) Start(startTime uint64) *MarketMarkPriceKlineService {
	s.start = &startTime
	return s
}

// End set endTime
func (s *MarketMarkPriceKlineService) End(endTime uint64) *MarketMarkPriceKlineService {
	s.end = &endTime
	return s
}

func (s *MarketMarkPriceKlineService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketMarkPriceKlineResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/mark-price-kline",
		secType:  secTypeNone,
	}
	if s.category != nil {
		r.setParam("category", *s.category)
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.start != nil {
		r.setParam("start", *s.start)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	j, err := newJSON(data)
	if err != nil {
		return nil, err
	}
	result := j.Get("result")
	res = new(models.MarketMarkPriceKlineResponse)
	res.Category = models.Category(result.Get("category").MustString())
	res.Symbol = result.Get("symbol").MustString()
	list := result.Get("list").MustArray()
	res.List = make([]*models.MarketMarkPriceKlineCandle, len(list))
	for i := range list {
		item := result.Get("list").GetIndex(i)
		if len(item.MustArray()) < 5 {
			return nil, fmt.Errorf("invalid kline response")
		}

		res.List[i] = &models.MarketMarkPriceKlineCandle{
			StartTime:  item.GetIndex(0).MustString(),
			OpenPrice:  item.GetIndex(1).MustString(),
			HighPrice:  item.GetIndex(2).MustString(),
			LowPrice:   item.GetIndex(3).MustString(),
			ClosePrice: item.GetIndex(4).MustString(),
		}
	}

	return res, nil
}

// MarketIndexPriceKlineService Market index price kline (GET /v5/market/index-price-kline)
type MarketIndexPriceKlineService struct {
	c        *Client
	category *models.Category
	symbol   string
	interval string
	start    *uint64
	end      *uint64
	limit    *int
}

// Category set category
func (s *MarketIndexPriceKlineService) Category(category models.Category) *MarketIndexPriceKlineService {
	s.category = &category
	return s
}

// Symbol set symbol
func (s *MarketIndexPriceKlineService) Symbol(symbol string) *MarketIndexPriceKlineService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *MarketIndexPriceKlineService) Interval(interval string) *MarketIndexPriceKlineService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *MarketIndexPriceKlineService) Limit(limit int) *MarketIndexPriceKlineService {
	s.limit = &limit
	return s
}

// Start set startTime
func (s *MarketIndexPriceKlineService) Start(startTime uint64) *MarketIndexPriceKlineService {
	s.start = &startTime
	return s
}

// End set endTime
func (s *MarketIndexPriceKlineService) End(endTime uint64) *MarketIndexPriceKlineService {
	s.end = &endTime
	return s
}

func (s *MarketIndexPriceKlineService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketIndexPriceKlineResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/mark-price-kline",
		secType:  secTypeNone,
	}
	if s.category != nil {
		r.setParam("category", *s.category)
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.start != nil {
		r.setParam("start", *s.start)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	j, err := newJSON(data)
	if err != nil {
		return nil, err
	}
	result := j.Get("result")
	res = new(models.MarketIndexPriceKlineResponse)
	res.Category = models.Category(result.Get("category").MustString())
	res.Symbol = result.Get("symbol").MustString()
	list := result.Get("list").MustArray()
	res.List = make([]*models.MarketIndexPriceKlineCandle, len(list))
	for i := range list {
		item := result.Get("list").GetIndex(i)
		if len(item.MustArray()) < 5 {
			return nil, fmt.Errorf("invalid kline response")
		}

		res.List[i] = &models.MarketIndexPriceKlineCandle{
			StartTime:  item.GetIndex(0).MustString(),
			OpenPrice:  item.GetIndex(1).MustString(),
			HighPrice:  item.GetIndex(2).MustString(),
			LowPrice:   item.GetIndex(3).MustString(),
			ClosePrice: item.GetIndex(4).MustString(),
		}
	}

	return res, nil
}

// MarketPremiumIndexPriceKlineService Market premium index price kline (GET /v5/market/premium-index-price-kline)
type MarketPremiumIndexPriceKlineService struct {
	c        *Client
	category *models.Category
	symbol   string
	interval string
	start    *uint64
	end      *uint64
	limit    *int
}

// Category set category
func (s *MarketPremiumIndexPriceKlineService) Category(category models.Category) *MarketPremiumIndexPriceKlineService {
	s.category = &category
	return s
}

// Symbol set symbol
func (s *MarketPremiumIndexPriceKlineService) Symbol(symbol string) *MarketPremiumIndexPriceKlineService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *MarketPremiumIndexPriceKlineService) Interval(interval string) *MarketPremiumIndexPriceKlineService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *MarketPremiumIndexPriceKlineService) Limit(limit int) *MarketPremiumIndexPriceKlineService {
	s.limit = &limit
	return s
}

// Start set startTime
func (s *MarketPremiumIndexPriceKlineService) Start(startTime uint64) *MarketPremiumIndexPriceKlineService {
	s.start = &startTime
	return s
}

// End set endTime
func (s *MarketPremiumIndexPriceKlineService) End(endTime uint64) *MarketPremiumIndexPriceKlineService {
	s.end = &endTime
	return s
}

func (s *MarketPremiumIndexPriceKlineService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketPremiumIndexPriceKlineResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/mark-price-kline",
		secType:  secTypeNone,
	}
	if s.category != nil {
		r.setParam("category", *s.category)
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.start != nil {
		r.setParam("start", *s.start)
	}
	if s.end != nil {
		r.setParam("end", *s.end)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	j, err := newJSON(data)
	if err != nil {
		return nil, err
	}
	result := j.Get("result")
	res = new(models.MarketPremiumIndexPriceKlineResponse)
	res.Category = models.Category(result.Get("category").MustString())
	res.Symbol = result.Get("symbol").MustString()
	list := result.Get("list").MustArray()
	res.List = make([]*models.MarketPremiumIndexPriceKlineCandle, len(list))
	for i := range list {
		item := result.Get("list").GetIndex(i)
		if len(item.MustArray()) < 5 {
			return nil, fmt.Errorf("invalid kline response")
		}

		res.List[i] = &models.MarketPremiumIndexPriceKlineCandle{
			StartTime:  item.GetIndex(0).MustString(),
			OpenPrice:  item.GetIndex(1).MustString(),
			HighPrice:  item.GetIndex(2).MustString(),
			LowPrice:   item.GetIndex(3).MustString(),
			ClosePrice: item.GetIndex(4).MustString(),
		}
	}

	return res, nil
}

type InstrumentsInfoService struct {
	c        *Client
	category models.Category
	symbol   *string
	statsu   *models.SymbolStatus
	baseCoin *string
	limit    *int
	cursor   *string
}

// Category set category
func (s *InstrumentsInfoService) Category(category models.Category) *InstrumentsInfoService {
	s.category = category
	return s
}

// Symbol set symbol
func (s *InstrumentsInfoService) Symbol(symbol string) *InstrumentsInfoService {
	s.symbol = &symbol
	return s
}

// Status set status
func (s *InstrumentsInfoService) Status(status models.SymbolStatus) *InstrumentsInfoService {
	s.statsu = &status
	return s
}

// BaseCoin set baseCoin
func (s *InstrumentsInfoService) BaseCoin(baseCoin string) *InstrumentsInfoService {
	s.baseCoin = &baseCoin
	return s
}

// Limit set limit
func (s *InstrumentsInfoService) Limit(limit int) *InstrumentsInfoService {
	s.limit = &limit
	return s
}

// Cursor set cursor
func (s *InstrumentsInfoService) Cursor(cursor string) *InstrumentsInfoService {
	s.cursor = &cursor
	return s
}

func (s *InstrumentsInfoService) Do(ctx context.Context, opts ...RequestOption) (res *models.InstrumentInfoResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/instruments-info",
		secType:  secTypeNone,
	}

	r.setParam("category", s.category)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.statsu != nil {
		r.setParam("status", *s.statsu)
	}
	if s.baseCoin != nil {
		r.setParam("baseCoin", *s.baseCoin)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.cursor != nil {
		r.setParam("cursor", *s.cursor)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	type instrumentInfoResponse struct {
		InstrumentsInfo *models.InstrumentInfoResponse `json:"result"`
	}
	resp := new(instrumentInfoResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	return resp.InstrumentsInfo, nil
}

type MarketOrderBookService struct {
	c        *Client
	category models.Category
	symbol   string
	limit    *int
}

// Category set category
func (s *MarketOrderBookService) Category(category models.Category) *MarketOrderBookService {
	s.category = category
	return s
}

// Symbol set symbol
func (s *MarketOrderBookService) Symbol(symbol string) *MarketOrderBookService {
	s.symbol = symbol
	return s
}

func (s *MarketOrderBookService) Limit(limit int) *MarketOrderBookService {
	s.limit = &limit
	return s
}

func (s *MarketOrderBookService) Do(ctx context.Context, opts ...RequestOption) (res *models.OrderBookInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/orderbook",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(MarketOrderBookResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type MarketOrderBookResponse struct {
	RetCode    int                  `json:"retCode"`
	RetMsg     string               `json:"retMsg"`
	Result     models.OrderBookInfo `json:"result"`
	RetExtInfo struct{}             `json:"retExtInfo"`
	Time       int64                `json:"time"`
}

type MarketTickersService struct {
	c        *Client
	category models.Category
	symbol   *string
	baseCoin *string
	expDate  *string
}

func (s *MarketTickersService) Category(category models.Category) *MarketTickersService {
	s.category = category
	return s
}

func (s *MarketTickersService) Symbol(symbol string) *MarketTickersService {
	s.symbol = &symbol
	return s
}

func (s *MarketTickersService) BaseCoin(baseCoin string) *MarketTickersService {
	s.baseCoin = &baseCoin
	return s
}

func (s *MarketTickersService) ExpDate(expDate string) *MarketTickersService {
	s.expDate = &expDate
	return s
}

func (s *MarketTickersService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketTickers, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/tickers",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.baseCoin != nil {
		r.setParam("baseCoin", *s.baseCoin)
	}
	if s.expDate != nil {
		r.setParam("expDate", *s.expDate)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(MarketTickersResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type MarketTickersResponse struct {
	RetCode    int                  `json:"retCode"`
	RetMsg     string               `json:"retMsg"`
	Result     models.MarketTickers `json:"result"`
	RetExtInfo struct{}             `json:"retExtInfo"`
	Time       int64                `json:"time"`
}

type MarketFundingRatesService struct {
	c         *Client
	category  models.Category
	symbol    string
	startTime *uint64
	endTime   *uint64
	limit     *int
}

func (s *MarketFundingRatesService) Category(category models.Category) *MarketFundingRatesService {
	s.category = category
	return s
}

func (s *MarketFundingRatesService) Symbol(symbol string) *MarketFundingRatesService {
	s.symbol = symbol
	return s
}

func (s *MarketFundingRatesService) StartTime(startTime uint64) *MarketFundingRatesService {
	s.startTime = &startTime
	return s
}

func (s *MarketFundingRatesService) EndTime(endTime uint64) *MarketFundingRatesService {
	s.endTime = &endTime
	return s
}

func (s *MarketFundingRatesService) Limit(limit int) *MarketFundingRatesService {
	s.limit = &limit
	return s
}

func (s *MarketFundingRatesService) Do(ctx context.Context, opts ...RequestOption) (res *models.FundingRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/tickers",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(MarketFundingRatesResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type MarketFundingRatesResponse struct {
	RetCode    int                `json:"retCode"`
	RetMsg     string             `json:"retMsg"`
	Result     models.FundingRate `json:"result"`
	RetExtInfo struct{}           `json:"retExtInfo"`
	Time       int64              `json:"time"`
}

type GetPublicRecentTradesService struct {
	c          *Client
	category   models.Category
	symbol     *string
	baseCoin   *string
	optionType *string
	limit      *int
}

func (s *GetPublicRecentTradesService) Category(category models.Category) *GetPublicRecentTradesService {
	s.category = category
	return s
}

func (s *GetPublicRecentTradesService) Symbol(symbol string) *GetPublicRecentTradesService {
	s.symbol = &symbol
	return s
}

func (s *GetPublicRecentTradesService) BaseCoin(baseCoin string) *GetPublicRecentTradesService {
	s.baseCoin = &baseCoin
	return s
}

func (s *GetPublicRecentTradesService) OptionType(optionType string) *GetPublicRecentTradesService {
	s.optionType = &optionType
	return s
}

func (s *GetPublicRecentTradesService) Limit(limit int) *GetPublicRecentTradesService {
	s.limit = &limit
	return s
}

func (s *GetPublicRecentTradesService) Do(ctx context.Context, opts ...RequestOption) (res *models.PublicRecentTradeHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/recent-trade",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.baseCoin != nil {
		r.setParam("baseCoin", *s.baseCoin)
	}
	if s.optionType != nil {
		r.setParam("optionType", *s.optionType)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetPublicRecentTradesResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetPublicRecentTradesResponse struct {
	RetCode    int                             `json:"retCode"`
	RetMsg     string                          `json:"retMsg"`
	Result     models.PublicRecentTradeHistory `json:"result"`
	RetExtInfo struct{}                        `json:"retExtInfo"`
	Time       int64                           `json:"time"`
}

type GetOpenInterestsService struct {
	c            *Client
	category     models.Category
	symbol       string
	intervalTime string
	startTime    *uint64
	endTime      *uint64
	limit        *int
	cursor       *string
}

func (s *GetOpenInterestsService) Category(category models.Category) *GetOpenInterestsService {
	s.category = category
	return s
}

func (s *GetOpenInterestsService) Symbol(symbol string) *GetOpenInterestsService {
	s.symbol = symbol
	return s
}

func (s *GetOpenInterestsService) IntervalTime(intervalTime string) *GetOpenInterestsService {
	s.intervalTime = intervalTime
	return s
}

func (s *GetOpenInterestsService) StartTime(startTime uint64) *GetOpenInterestsService {
	s.startTime = &startTime
	return s
}

func (s *GetOpenInterestsService) EndTime(endTime uint64) *GetOpenInterestsService {
	s.endTime = &endTime
	return s
}

func (s *GetOpenInterestsService) Limit(limit int) *GetOpenInterestsService {
	s.limit = &limit
	return s
}

func (s *GetOpenInterestsService) Cursor(cursor string) *GetOpenInterestsService {
	s.cursor = &cursor
	return s
}

func (s *GetOpenInterestsService) Do(ctx context.Context, opts ...RequestOption) (res *models.OpenInterestInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/open-interest",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	r.setParam("symbol", s.symbol)
	r.setParam("intervalTime", s.intervalTime)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.cursor != nil {
		r.setParam("cursor", *s.cursor)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetOpenInterestsResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetOpenInterestsResponse struct {
	RetCode    int                     `json:"retCode"`
	RetMsg     string                  `json:"retMsg"`
	Result     models.OpenInterestInfo `json:"result"`
	RetExtInfo struct{}                `json:"retExtInfo"`
	Time       int64                   `json:"time"`
}

type GetHistoricalVolatilityService struct {
	c         *Client
	category  models.Category
	baseCoin  *string
	period    *string
	startTime *uint64
	endTime   *uint64
}

func (s *GetHistoricalVolatilityService) Category(category models.Category) *GetHistoricalVolatilityService {
	s.category = category
	return s
}

func (s *GetHistoricalVolatilityService) BaseCoin(baseCoin string) *GetHistoricalVolatilityService {
	s.baseCoin = &baseCoin
	return s
}

func (s *GetHistoricalVolatilityService) Period(period string) *GetHistoricalVolatilityService {
	s.period = &period
	return s
}

func (s *GetHistoricalVolatilityService) StartTime(startTime uint64) *GetHistoricalVolatilityService {
	s.startTime = &startTime
	return s
}

func (s *GetHistoricalVolatilityService) EndTime(endTime uint64) *GetHistoricalVolatilityService {
	s.endTime = &endTime
	return s
}

func (s *GetHistoricalVolatilityService) Do(ctx context.Context, opts ...RequestOption) (res *models.HistoricalVolatilityInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/historical-volatility",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	if s.baseCoin != nil {
		r.setParam("baseCoin", *s.baseCoin)
	}
	if s.period != nil {
		r.setParam("period", *s.period)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(models.HistoricalVolatilityInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetInsuranceInfoService struct {
	c    *Client
	coin *string
}

func (s *GetInsuranceInfoService) Coin(coin string) *GetInsuranceInfoService {
	s.coin = &coin
	return s
}

func (s *GetInsuranceInfoService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketInsuranceInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/insurance",
		secType:  secTypeNone,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetInsuranceInfoResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetInsuranceInfoResponse struct {
	RetCode    int                        `json:"retCode"`
	RetMsg     string                     `json:"retMsg"`
	Result     models.MarketInsuranceInfo `json:"result"`
	RetExtInfo struct{}                   `json:"retExtInfo"`
	Time       int64                      `json:"time"`
}

type GetRiskLimitService struct {
	c        *Client
	category models.Category
	symbol   *string
}

func (s *GetRiskLimitService) Category(category models.Category) *GetRiskLimitService {
	s.category = category
	return s
}

func (s *GetRiskLimitService) Symbol(symbol string) *GetRiskLimitService {
	s.symbol = &symbol
	return s
}

func (s *GetRiskLimitService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketRiskLimitInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/risk-limit",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetRiskLimitResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetRiskLimitResponse struct {
	RetCode    int                        `json:"retCode"`
	RetMsg     string                     `json:"retMsg"`
	Result     models.MarketRiskLimitInfo `json:"result"`
	RetExtInfo struct{}                   `json:"retExtInfo"`
	Time       int64                      `json:"time"`
}

type GetDeliveryPriceService struct {
	c        *Client
	category models.Category
	symbol   *string
	baseCoin *string
	limit    *int
	cursor   *string
}

func (s *GetDeliveryPriceService) Category(category models.Category) *GetDeliveryPriceService {
	s.category = category
	return s
}

func (s *GetDeliveryPriceService) Symbol(symbol string) *GetDeliveryPriceService {
	s.symbol = &symbol
	return s
}

func (s *GetDeliveryPriceService) BaseCoin(baseCoin string) *GetDeliveryPriceService {
	s.baseCoin = &baseCoin
	return s
}

func (s *GetDeliveryPriceService) Limit(limit int) *GetDeliveryPriceService {
	s.limit = &limit
	return s
}

func (s *GetDeliveryPriceService) Cursor(cursor string) *GetDeliveryPriceService {
	s.cursor = &cursor
	return s
}

func (s *GetDeliveryPriceService) Do(ctx context.Context, opts ...RequestOption) (res *models.DeliveryPriceInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/delivery-price",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.baseCoin != nil {
		r.setParam("baseCoin", *s.baseCoin)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.cursor != nil {
		r.setParam("cursor", *s.cursor)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetDeliveryPriceResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetDeliveryPriceResponse struct {
	RetCode    int                      `json:"retCode"`
	RetMsg     string                   `json:"retMsg"`
	Result     models.DeliveryPriceInfo `json:"result"`
	RetExtInfo struct{}                 `json:"retExtInfo"`
	Time       int64                    `json:"time"`
}

type GetMarketLSRatioService struct {
	c        *Client
	category models.Category
	baseCoin string
	period   string
	limit    *int
}

func (s *GetMarketLSRatioService) Category(category models.Category) *GetMarketLSRatioService {
	s.category = category
	return s
}

func (s *GetMarketLSRatioService) BaseCoin(baseCoin string) *GetMarketLSRatioService {
	s.baseCoin = baseCoin
	return s
}

func (s *GetMarketLSRatioService) Period(period string) *GetMarketLSRatioService {
	s.period = period
	return s
}

func (s *GetMarketLSRatioService) Limit(limit int) *GetMarketLSRatioService {
	s.limit = &limit
	return s
}

func (s *GetMarketLSRatioService) Do(ctx context.Context, opts ...RequestOption) (res *models.MarketLongShortRatioInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/account-ratio",
		secType:  secTypeNone,
	}
	r.setParam("category", s.category)
	r.setParam("baseCoin", s.baseCoin)
	r.setParam("period", s.period)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetMarketLSRatioResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetMarketLSRatioResponse struct {
	RetCode    int                             `json:"retCode"`
	RetMsg     string                          `json:"retMsg"`
	Result     models.MarketLongShortRatioInfo `json:"result"`
	RetExtInfo struct{}                        `json:"retExtInfo"`
	Time       int64                           `json:"time"`
}

type GetServerTimeService struct {
	c *Client
}

func (s *GetServerTimeService) Do(ctx context.Context, opts ...RequestOption) (res *models.ServerTimeResult, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/time",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(GetServerTimeResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

type GetServerTimeResponse struct {
	RetCode    int                     `json:"retCode"`
	RetMsg     string                  `json:"retMsg"`
	Result     models.ServerTimeResult `json:"result"`
	RetExtInfo struct{}                `json:"retExtInfo"`
	Time       int64                   `json:"time"`
}
