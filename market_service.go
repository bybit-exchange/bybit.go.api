package bybit_connector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api/models"
)

func (s *BybitClientRequest) GetServerTime(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/time",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetMarketKline(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/kline",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func GetMarketKlineResponse(err error, data []byte, res *models.MarketKlineResponse) (*models.MarketKlineResponse, *models.MarketKlineResponse, error) {
	j, err := newJSON(data)
	if err != nil {
		return nil, nil, err
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
			return nil, nil, fmt.Errorf("invalid kline response")
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
	return res, nil, nil
}

func (s *BybitClientRequest) GetMarkPriceKline(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/mark-price-kline",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func GetMarkPriceKline(err error, data []byte, res *models.MarketMarkPriceKlineResponse) (*models.MarketMarkPriceKlineResponse, error) {
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

func (s *BybitClientRequest) GetIndexPriceKline(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/index-price-kline",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func GetIndexPriceKline(err error, data []byte, res *models.MarketIndexPriceKlineResponse) (*models.MarketIndexPriceKlineResponse, error) {
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

func (s *BybitClientRequest) GetPremiumIndexPriceKline(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/premium-index-price-kline",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func GetPremiumIndexKline(err error, data []byte, res *models.MarketPremiumIndexPriceKlineResponse) (*models.MarketPremiumIndexPriceKlineResponse, error) {
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

func (s *BybitClientRequest) GetInstrumentInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/instruments-info",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetOrderBookInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/orderbook",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetRPIOrderBook
func (s *BybitClientRequest) GetRPIOrderBook(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/rpi_orderbook",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetMarketTickers(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/tickers",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetFundingRateHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/funding/history",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetPublicRecentTrades(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/recent-trade",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetOpenInterests(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/open-interest",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetHistoryVolatility(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/historical-volatility",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetMarketInsurance(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/insurance",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetMarketRiskLimits(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/risk-limit",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetDeliveryPrice(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/delivery-price",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetNewDeliveryPrice
func (s *BybitClientRequest) GetNewDeliveryPrice(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/new-delivery-price",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetLongShortRatio(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/account-ratio",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetOrderPriceLimit(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/price-limit",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetADLAlerts
func (s *BybitClientRequest) GetADLAlerts(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/adlAlert",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetIndexPriceComponents
func (s *BybitClientRequest) GetIndexPriceComponents(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/index-price-components",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetFeeGroupStructure
func (s *BybitClientRequest) GetFeeGroupStructure(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/market/fee-group-info",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpreadInstrumentsInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spread/instrument",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpreadOrderbook(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spread/orderbook",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpreadRecentTrades(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spread/recent-trade",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpreadTickers(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spread/tickers",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
