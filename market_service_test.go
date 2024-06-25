package bybit_connector

import (
	"net/http"
	"testing"

	"github.com/bybit-exchange/bybit.go.api/models"
	"github.com/stretchr/testify/suite"
)

type marketTestSuite struct {
	baseTestSuite
}

func TestMarketKline(t *testing.T) {
	suite.Run(t, new(marketTestSuite))
}

func (s *marketTestSuite) TestMarketKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSD",
        "category": "inverse",
        "list": [
            [
                "1670608800000",
                "17071",
                "17073",
                "17027",
                "17055.5",
                "268611",
                "15.74462667"
            ],
            [
                "1670605200000",
                "17071.5",
                "17071.5",
                "17061",
                "17071",
                "4177",
                "0.24469757"
            ],
            [
                "1670601600000",
                "17086.5",
                "17088",
                "16978",
                "17071.5",
                "6356",
                "0.37288112"
            ]
        ]
    },
    "retExtInfo": {},
    "time": 1672025956592
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSD"
	interval := "1"
	start := uint64(1499040000000)
	end := uint64(1499040000001)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"interval": interval,
			"start":    start,
			"end":      end,
			"limit":    limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarketKlineService().
		Category(category).Symbol(symbol).Interval(interval).
		Start(start).End(end).Limit(limit).
		Do(newContext())
	e1 := &models.MarketKlineResponse{
		Symbol:   "BTCUSD",
		Category: "inverse",
		List: []*models.MarketKlineCandle{
			{
				StartTime:  "1670608800000",
				OpenPrice:  "17071",
				HighPrice:  "17073",
				LowPrice:   "17027",
				ClosePrice: "17055.5",
				Volume:     "268611",
				Turnover:   "15.74462667",
			},
			{
				StartTime:  "1670605200000",
				OpenPrice:  "17071.5",
				HighPrice:  "17071.5",
				LowPrice:   "17061",
				ClosePrice: "17071",
				Volume:     "4177",
				Turnover:   "0.24469757",
			},
			{
				StartTime:  "1670601600000",
				OpenPrice:  "17086.5",
				HighPrice:  "17088",
				LowPrice:   "16978",
				ClosePrice: "17071.5",
				Volume:     "6356",
				Turnover:   "0.37288112",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketKlineEqual(e1, res)
}

func (s *marketTestSuite) assertMarketKlineEqual(expected, actual *models.MarketKlineResponse) {
	r := s.r()
	r.Equal(expected.Symbol, actual.Symbol, "Symbol")
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(len(expected.List), len(actual.List), "List")
	r.Equal(expected.List, actual.List)
}
func (s *marketTestSuite) TestMarketMarkPriceKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSDT",
        "category": "linear",
        "list": [
            [
            "1670608800000",
            "17164.16",
            "17164.16",
            "17121.5",
            "17131.64"
            ]
        ]
    },
    "retExtInfo": {},
    "time": 1672026361839
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSDT"
	interval := "1"
	start := uint64(1499040000000)
	end := uint64(1499040000001)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"interval": interval,
			"start":    start,
			"end":      end,
			"limit":    limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarketMarkPriceKlineService().
		Category(category).Symbol(symbol).Interval(interval).
		Start(start).End(end).Limit(limit).
		Do(newContext())
	e1 := &models.MarketMarkPriceKlineResponse{
		Symbol:   "BTCUSDT",
		Category: "linear",
		List: []*models.MarketMarkPriceKlineCandle{
			{
				StartTime:  "1670608800000",
				OpenPrice:  "17164.16",
				HighPrice:  "17164.16",
				LowPrice:   "17121.5",
				ClosePrice: "17131.64",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketMarkPriceKlineEqual(e1, res)
}

func (s *marketTestSuite) assertMarketMarkPriceKlineEqual(expected, actual *models.MarketMarkPriceKlineResponse) {
	r := s.r()
	r.Equal(expected.Symbol, actual.Symbol, "Symbol")
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(len(expected.List), len(actual.List), "List")
	r.Equal(expected.List, actual.List)
}

func (s *marketTestSuite) TestMarketIndexPriceKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSDZ22",
        "category": "inverse",
        "list": [
            [
                "1670608800000",
                "17167.00",
                "17167.00",
                "17161.90",
                "17163.07"
            ],
            [
                "1670608740000",
                "17166.54",
                "17167.69",
                "17165.42",
                "17167.00"
            ]
        ]
    },
    "retExtInfo": {},
    "time": 1672026471128
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSDZ22"
	interval := "1"
	start := uint64(1499040000000)
	end := uint64(1499040000001)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"interval": interval,
			"start":    start,
			"end":      end,
			"limit":    limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarketIndexPriceKlineService().
		Category(category).Symbol(symbol).Interval(interval).
		Start(start).End(end).Limit(limit).
		Do(newContext())
	e1 := &models.MarketIndexPriceKlineResponse{
		Symbol:   "BTCUSDZ22",
		Category: models.CategoryInverse,
		List: []*models.MarketIndexPriceKlineCandle{
			{
				StartTime:  "1670608800000",
				OpenPrice:  "17167.00",
				HighPrice:  "17167.00",
				LowPrice:   "17161.90",
				ClosePrice: "17163.07",
			},
			{
				StartTime:  "1670608740000",
				OpenPrice:  "17166.54",
				HighPrice:  "17167.69",
				LowPrice:   "17165.42",
				ClosePrice: "17167.00",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketIndexPriceKlineEqual(e1, res)
}

func (s *marketTestSuite) assertMarketIndexPriceKlineEqual(expected, actual *models.MarketIndexPriceKlineResponse) {
	r := s.r()
	r.Equal(expected.Symbol, actual.Symbol, "Symbol")
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(len(expected.List), len(actual.List), "List")
	r.Equal(expected.List, actual.List)
}

func (s *marketTestSuite) TestMarketPremiumIndexPriceKline() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCPERP",
        "category": "linear",
        "list": [
            [
                "1672026540000",
                "0.000000",
                "0.000000",
                "0.000000",
                "0.000000"
            ],
            [
                "1672026480000",
                "0.000000",
                "0.000000",
                "0.000000",
                "0.000000"
            ]
            ]
    },
    "retExtInfo": {},
    "time": 1672026605042
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCPERP"
	interval := "1"
	start := uint64(1499040000000)
	end := uint64(1499040000001)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"interval": interval,
			"start":    start,
			"end":      end,
			"limit":    limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarketPremiumIndexPriceKlineService().
		Category(category).Symbol(symbol).Interval(interval).
		Start(start).End(end).Limit(limit).
		Do(newContext())
	e1 := &models.MarketPremiumIndexPriceKlineResponse{
		Symbol:   "BTCPERP",
		Category: "linear",
		List: []*models.MarketPremiumIndexPriceKlineCandle{
			{
				StartTime:  "1672026540000",
				OpenPrice:  "0.000000",
				HighPrice:  "0.000000",
				LowPrice:   "0.000000",
				ClosePrice: "0.000000",
			},
			{
				StartTime:  "1672026480000",
				OpenPrice:  "0.000000",
				HighPrice:  "0.000000",
				LowPrice:   "0.000000",
				ClosePrice: "0.000000",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketPremiumIndexPriceKlineEqual(e1, res)
}

func (s *marketTestSuite) assertMarketPremiumIndexPriceKlineEqual(expected, actual *models.MarketPremiumIndexPriceKlineResponse) {
	r := s.r()
	r.Equal(expected.Symbol, actual.Symbol, "Symbol")
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(len(expected.List), len(actual.List), "List")
	r.Equal(expected.List, actual.List)
}

func (s *marketTestSuite) TestInstrumentsInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "spot",
        "list": [
            {
                    "symbol": "BTCUSDT",
                    "baseCoin": "BTC",
                    "quoteCoin": "USDT",
                    "innovation": "0",
                    "status": "Trading",
                    "marginTrading": "both",
                    "lotSizeFilter": {
                    "basePrecision": "0.000001",
                    "quotePrecision": "0.00000001",
                    "minOrderQty": "0.000048",
                    "maxOrderQty": "71.73956243",
                    "minOrderAmt": "1",
                    "maxOrderAmt": "2000000"
                    },
                    "priceFilter": {
                    "tickSize": "0.01"
                }
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672712468011
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSD"
	status := models.SymbolStatusTrading
	baseCoin := "BTC"
	limit := 10
	cursor := "cursor"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"status":   status,
			"baseCoin": baseCoin,
			"limit":    limit,
			"cursor":   cursor,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewInstrumentsInfoService().
		Category(category).Symbol(symbol).Status(status).
		BaseCoin(baseCoin).Limit(limit).Cursor(cursor).
		Do(newContext())

	e1 := &models.InstrumentInfoResponse{
		Category:       "spot",
		NextPageCursor: "",
		List: []models.Instrument{
			{
				Symbol:        "BTCUSDT",
				BaseCoin:      "BTC",
				QuoteCoin:     "USDT",
				Innovation:    "0",
				Status:        "Trading",
				MarginTrading: "both",
				LotSizeFilter: models.LotSizeFilter{
					BasePrecision:  "0.000001",
					QuotePrecision: "0.00000001",
					MinOrderQty:    "0.000048",
					MaxOrderQty:    "71.73956243", MinOrderAmt: "1",
					MaxOrderAmt: "2000000",
				},
				PriceFilter: models.PriceFilter{
					TickSize: "0.01",
				},
			},
		},
	}

	s.r().NoError(err)
	s.assertInstrumentsInfoEqual(e1, res)
}

func (s *marketTestSuite) assertInstrumentsInfoEqual(expected, actual *models.InstrumentInfoResponse) {
	r := s.r()
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.NextPageCursor, actual.NextPageCursor, "NextPageCursor")
	r.Equal(len(expected.List), len(actual.List), "List")
	r.Equal(expected.List, actual.List)
}

func (s *marketTestSuite) TestOrderBookInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "SUCCESS",
    "result": {
        "s": "BTC-30DEC22-18000-C",
        "b": [
            [
                "5",
                "3.12"
            ]
        ],
        "a": [
            [
                "175",
                "4.88"
            ]
        ],
        "u": 1203433656,
        "ts": 1672043188375
    },
    "retExtInfo": {},
    "time": 1672043199230
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSD"
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"limit":    limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewOrderBookService().
		Category(category).Symbol(symbol).Limit(limit).
		Do(newContext())

	e1 := &models.OrderBookInfo{
		Symbol: "BTC-30DEC22-18000-C",
		Bids: []models.OrderBookEntry{
			{
				"5",
				"3.12",
			},
		},
		Asks: []models.OrderBookEntry{
			{
				"175",
				"4.88",
			},
		},
		Timestamp: 1672043188375,
		UpdateID:  1203433656,
	}

	s.r().NoError(err)
	s.assertOrderbookInfoEqual(e1, res)
}

func (s *marketTestSuite) assertOrderbookInfoEqual(expected, actual *models.OrderBookInfo) {
	r := s.r()
	r.Equal(expected.Symbol, actual.Symbol, "Symbol")
	r.Equal(expected.Timestamp, actual.Timestamp, "Timestamp")
	r.Equal(expected.UpdateID, actual.UpdateID, "UpdateID")
	r.Equal(expected.Bids, actual.Bids)
	r.Equal(expected.Asks, actual.Asks)
}

func (s *marketTestSuite) TestInverseTickersInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "inverse",
        "list": [
            {
                "symbol": "BTCUSD",
                "lastPrice": "16597.00",
                "indexPrice": "16598.54",
                "markPrice": "16596.00",
                "prevPrice24h": "16464.50",
                "price24hPcnt": "0.008047",
                "highPrice24h": "30912.50",
                "lowPrice24h": "15700.00",
                "prevPrice1h": "16595.50",
                "openInterest": "373504107",
                "openInterestValue": "22505.67",
                "turnover24h": "2352.94950046",
                "volume24h": "49337318",
                "fundingRate": "-0.001034",
                "nextFundingTime": "1672387200000",
                "predictedDeliveryPrice": "",
                "basisRate": "",
                "deliveryFeeRate": "",
                "deliveryTime": "0",
                "ask1Size": "1",
                "bid1Price": "16596.00",
                "ask1Price": "16597.50",
                "bid1Size": "1",
                "basis": ""
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672376496682
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.testTickersInfo([]*models.TickerInfo{
		{
			Symbol:                 "BTCUSD",
			LastPrice:              "16597.00",
			IndexPrice:             "16598.54",
			MarkPrice:              "16596.00",
			PrevPrice24h:           "16464.50",
			Price24hPcnt:           "0.008047",
			HighPrice24h:           "30912.50",
			LowPrice24h:            "15700.00",
			PrevPrice1h:            "16595.50",
			OpenInterest:           "373504107",
			OpenInterestValue:      "22505.67",
			Turnover24h:            "2352.94950046",
			Volume24h:              "49337318",
			FundingRate:            "-0.001034",
			NextFundingTime:        "1672387200000",
			PredictedDeliveryPrice: "",
			BasisRate:              "",
			DeliveryFeeRate:        "",
			DeliveryTime:           "0",
			Ask1Size:               "1",
			Bid1Price:              "16596.00",
			Ask1Price:              "16597.50",
			Bid1Size:               "1",
			Basis:                  "",
		},
	})
}

func (s *marketTestSuite) TestOptionTickersInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "option",
        "list": [
            {
                "symbol": "BTC-30DEC22-18000-C",
                "bid1Price": "0",
                "bid1Size": "0",
                "bid1Iv": "0",
                "ask1Price": "435",
                "ask1Size": "0.66",
                "ask1Iv": "5",
                "lastPrice": "435",
                "highPrice24h": "435",
                "lowPrice24h": "165",
                "markPrice": "0.00000009",
                "indexPrice": "16600.55",
                "markIv": "0.7567",
                "underlyingPrice": "16590.42",
                "openInterest": "6.3",
                "turnover24h": "2482.73",
                "volume24h": "0.15",
                "totalVolume": "99",
                "totalTurnover": "1967653",
                "delta": "0.00000001",
                "gamma": "0.00000001",
                "vega": "0.00000004",
                "theta": "-0.00000152",
                "predictedDeliveryPrice": "0",
                "change24h": "86"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672376592395
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.testTickersInfo([]*models.TickerInfo{
		{
			Symbol:          "BTC-30DEC22-18000-C",
			Bid1Price:       "0",
			Bid1Size:        "0",
			Bid1Iv:          "0",
			Ask1Price:       "435",
			Ask1Size:        "0.66",
			Ask1Iv:          "5",
			LastPrice:       "435",
			HighPrice24h:    "435",
			LowPrice24h:     "165",
			MarkPrice:       "0.00000009",
			IndexPrice:      "16600.55",
			MarkIv:          "0.7567",
			UnderlyingPrice: "16590.42",
			OpenInterest:    "6.3",
			Turnover24h:     "2482.73",
			Volume24h:       "0.15",
			TotalVolume:     "99",
			TotalTurnover:   "1967653",
			// Delta: "0.00000001",
			// Gamma: "0.00000001",
			// Vega: "0.00000004",
			// Theta: "-0.00000152",
			PredictedDeliveryPrice: "0",
			Change24h:              "86",
		},
	})
}

func (s *marketTestSuite) TestSpotTickersInfo() {
	data := []byte(`{
	    "retCode": 0,
	    "retMsg": "OK",
	    "result": {
	        "category": "spot",
	        "list": [
	            {
	                "symbol": "BTCUSDT",
	                "bid1Price": "20517.96",
	                "bid1Size": "2",
	                "ask1Price": "20527.77",
	                "ask1Size": "1.862172",
	                "lastPrice": "20533.13",
	                "prevPrice24h": "20393.48",
	                "price24hPcnt": "0.0068",
	                "highPrice24h": "21128.12",
	                "lowPrice24h": "20318.89",
	                "turnover24h": "243765620.65899866",
	                "volume24h": "11801.27771",
	                "usdIndexPrice": "20784.12009279"
	            }
	        ]
	    },
	    "retExtInfo": {},
	    "time": 1673859087947
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.testTickersInfo([]*models.TickerInfo{
		{
			Symbol:        "BTCUSDT",
			Bid1Price:     "20517.96",
			Bid1Size:      "2",
			Ask1Price:     "20527.77",
			Ask1Size:      "1.862172",
			LastPrice:     "20533.13",
			PrevPrice24h:  "20393.48",
			Price24hPcnt:  "0.0068",
			HighPrice24h:  "21128.12",
			LowPrice24h:   "20318.89",
			Turnover24h:   "243765620.65899866",
			Volume24h:     "11801.27771",
			UsdIndexPrice: "20784.12009279",
		},
	})
}

func (s *marketTestSuite) testTickersInfo(e1 []*models.TickerInfo) {
	category := models.CategoryInverse
	symbol := "BTCUSD"
	baseCoin := "BTC"
	expDate := "2022-12-30"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"baseCoin": baseCoin,
			"expDate":  expDate,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewTickersService().
		Category(category).Symbol(symbol).BaseCoin(baseCoin).ExpDate(expDate).
		Do(newContext())

	s.r().NoError(err)
	s.assertTickersEqual(e1, res.List)
}

func (s *marketTestSuite) assertTickersEqual(expected, actual []*models.TickerInfo) {
	r := s.r()
	r.Equal(expected, actual)
}

func (s *marketTestSuite) TestFundingRates() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "linear",
        "list": [
            {
                "symbol": "ETHPERP",
                "fundingRate": "0.0001",
                "fundingRateTimestamp": "1672041600000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672051897447
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSD"
	startTime := uint64(1499040000000)
	endTime := uint64(1499040000001)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category":  category,
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFundingTatesService().
		Category(category).Symbol(symbol).
		StartTime(startTime).EndTime(endTime).Limit(limit).
		Do(newContext())

	e1 := &models.FundingRate{
		Category: "linear",
		List: []models.FundingRateInfo{
			{
				Symbol:               "ETHPERP",
				FundingRate:          "0.0001",
				FundingRateTimestamp: "1672041600000",
			},
		},
	}

	s.r().NoError(err)
	s.assertFundingRatesEqual(e1, res)
}

func (s *marketTestSuite) assertFundingRatesEqual(expected, actual *models.FundingRate) {
	r := s.r()
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.List, actual.List, "List")
}

func (s *marketTestSuite) TestGetPublicRecentTrades() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "spot",
        "list": [
            {
                "execId": "2100000000007764263",
                "symbol": "BTCUSDT",
                "price": "16618.49",
                "size": "0.00012",
                "side": "Buy",
                "time": "1672052955758",
                "isBlockTrade": false
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672053054358
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSD"
	baseCoin := "BTC"
	optionType := "optionType"
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		e.setParams(params{
			"category":   category,
			"symbol":     symbol,
			"baseCoin":   baseCoin,
			"optionType": optionType,
			"limit":      limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetPublicRecentTradesService().
		Category(category).Symbol(symbol).BaseCoin(baseCoin).
		OptionType(optionType).Limit(limit).
		Do(newContext())

	e1 := &models.PublicRecentTradeHistory{
		Category: "spot",
		List: []models.TradeInfo{
			{
				ExecId:       "2100000000007764263",
				Symbol:       "BTCUSDT",
				Price:        "16618.49",
				Size:         "0.00012",
				Side:         "Buy",
				Time:         "1672052955758",
				IsBlockTrade: false,
			},
		},
	}

	s.r().NoError(err)
	s.assertPublicRecentTradeHistoryEqual(e1, res)
}

func (s *marketTestSuite) assertPublicRecentTradeHistoryEqual(expected, actual *models.PublicRecentTradeHistory) {
	r := s.r()
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.List, actual.List, "List")
}

func (s *marketTestSuite) TestGetOpenInterests() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "symbol": "BTCUSD",
        "category": "inverse",
        "list": [
            {
                "openInterest": "461134384.00000000",
                "timestamp": "1669571400000"
            },
            {
                "openInterest": "461134292.00000000",
                "timestamp": "1669571100000"
            }
        ],
        "nextPageCursor": ""
    },
    "retExtInfo": {},
    "time": 1672053548579
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSD"
	intervalTime := "intervalTime"
	startTime := uint64(1499040000000)
	endTime := uint64(1499040000001)
	limit := 10
	cursor := "cursor"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet

		e.setParams(params{
			"category":     category,
			"symbol":       symbol,
			"intervalTime": intervalTime,
			"startTime":    startTime,
			"endTime":      endTime,
			"limit":        limit,
			"cursor":       cursor,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetOpenInterestsService().
		Category(category).Symbol(symbol).IntervalTime(intervalTime).
		StartTime(startTime).EndTime(endTime).Limit(limit).Cursor(cursor).
		Do(newContext())

	e1 := &models.OpenInterestInfo{
		Symbol:   "BTCUSD",
		Category: "inverse",
		List: []models.OpenInterest{
			{
				OpenInterest: "461134384.00000000",
				Timestamp:    "1669571400000",
			},
			{
				OpenInterest: "461134292.00000000",
				Timestamp:    "1669571100000",
			},
		},
		NextPageCursor: "",
	}

	s.r().NoError(err)
	s.assertOpenInterestInfoEqual(e1, res)
}

func (s *marketTestSuite) assertOpenInterestInfoEqual(expected, actual *models.OpenInterestInfo) {
	r := s.r()
	r.Equal(expected.Symbol, actual.Symbol, "Symbol")
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.List, actual.List, "List")
	r.Equal(expected.NextPageCursor, actual.NextPageCursor, "NextPageCursor")
}

func (s *marketTestSuite) TestGetHistoricalVolatility() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "SUCCESS",
    "category": "option",
    "result": [
        {
            "period": 7,
            "value": "0.27545620",
            "time": "1672232400000"
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	baseCoin := "BTC"
	period := "period"
	startTime := uint64(1499040000000)
	endTime := uint64(1499040000001)
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet

		e.setParams(params{
			"category":  category,
			"baseCoin":  baseCoin,
			"period":    period,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetHistoricalVolatilityService().
		Category(category).BaseCoin(baseCoin).Period(period).
		StartTime(startTime).EndTime(endTime).
		Do(newContext())

	e1 := &models.HistoricalVolatilityInfo{
		Category: "option",
		List: []models.VolatilityData{
			{
				Period: 7,
				Value:  "0.27545620",
				Time:   "1672232400000",
			},
		},
	}

	s.r().NoError(err)
	s.assertHistoricalVolatilityInfoEqual(e1, res)
}

func (s *marketTestSuite) assertHistoricalVolatilityInfoEqual(expected, actual *models.HistoricalVolatilityInfo) {
	r := s.r()
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.List, actual.List, "List")
}

func (s *marketTestSuite) TestGetInsuranceInfo() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "updatedTime": "1672012800000",
        "list": [
            {
                "coin": "ETH",
                "balance": "0.00187332",
                "value": "0"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672053931991
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "BTC"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet

		e.setParams(params{
			"coin": coin,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetInsuranceInfoService().
		Coin(coin).Do(newContext())

	e1 := &models.MarketInsuranceInfo{
		UpdatedTime: "1672012800000",
		List: []models.InsuranceData{
			{
				Coin:    "ETH",
				Balance: "0.00187332",
				Value:   "0",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketInsuranceInfoEqual(e1, res)
}

func (s *marketTestSuite) assertMarketInsuranceInfoEqual(expected, actual *models.MarketInsuranceInfo) {
	r := s.r()
	r.Equal(expected.UpdatedTime, actual.UpdatedTime, "UpdatedTime")
	r.Equal(expected.List, actual.List, "List")
}

func (s *marketTestSuite) TestGetRiskLimit() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "category": "inverse",
        "list": [
            {
                "id": 1,
                "symbol": "BTCUSD",
                "riskLimitValue": "150",
                "maintenanceMargin": "0.5",
                "initialMargin": "1",
                "isLowestRisk": 1,
                "maxLeverage": "100.00"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672054488010
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet

		e.setParams(params{
			"category": category,
			"symbol":   symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetRiskLimitService().
		Category(category).Symbol(symbol).
		Do(newContext())

	e1 := &models.MarketRiskLimitInfo{
		Category: "inverse",
		List: []models.RiskLimitData{
			{
				Id:                1,
				Symbol:            "BTCUSD",
				RiskLimitValue:    "150",
				MaintenanceMargin: "0.5",
				InitialMargin:     "1",
				IsLowestRisk:      1,
				MaxLeverage:       "100.00",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketRiskLimitInfoEqual(e1, res)
}

func (s *marketTestSuite) assertMarketRiskLimitInfoEqual(expected, actual *models.MarketRiskLimitInfo) {
	r := s.r()
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.List, actual.List, "List")
}

func (s *marketTestSuite) TestGetDeliveryPrice() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "success",
    "result": {
        "category": "option",
        "nextPageCursor": "",
        "list": [
            {
                "symbol": "ETH-26DEC22-1400-C",
                "deliveryPrice": "1220.728594450",
                "deliveryTime": "1672041600000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1672055336993
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	symbol := "BTCUSDT"
	baseCoin := "BTC"
	limit := 10
	cursor := "cursor"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet

		e.setParams(params{
			"category": category,
			"symbol":   symbol,
			"baseCoin": baseCoin,
			"limit":    limit,
			"cursor":   cursor,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetDeliveryPriceService().
		Category(category).Symbol(symbol).BaseCoin(baseCoin).
		Limit(limit).Cursor(cursor).
		Do(newContext())

	e1 := &models.DeliveryPriceInfo{
		Category:       "option",
		NextPageCursor: "",
		List: []models.DeliveryPriceData{
			{
				Symbol:        "ETH-26DEC22-1400-C",
				DeliveryPrice: "1220.728594450",
				DeliveryTime:  "1672041600000",
			},
		},
	}

	s.r().NoError(err)
	s.assertDeliveryPriceInfoEqual(e1, res)
}

func (s *marketTestSuite) assertDeliveryPriceInfoEqual(expected, actual *models.DeliveryPriceInfo) {
	r := s.r()
	r.Equal(expected.Category, actual.Category, "Category")
	r.Equal(expected.NextPageCursor, actual.NextPageCursor, "NextPageCursor")
	r.Equal(expected.List, actual.List, "List")
}

func (s *marketTestSuite) TestGetMarketLSRatio() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "list": [
            {
                "symbol": "BTCUSDT",
                "buyRatio": "0.5777",
                "sellRatio": "0.4223",
                "timestamp": "1695772800000"
            }
        ]
    },
    "retExtInfo": {},
    "time": 1695785131028
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	category := models.CategoryInverse
	baseCoin := "BTC"
	period := "period"
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet

		e.setParams(params{
			"category": category,
			"baseCoin": baseCoin,
			"period":   period,
			"limit":    limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetMarketLSRatioService().
		Category(category).BaseCoin(baseCoin).
		Period(period).Limit(limit).
		Do(newContext())

	e1 := &models.MarketLongShortRatioInfo{
		List: []models.LongShortRatioData{
			{
				Symbol:    "BTCUSDT",
				BuyRatio:  "0.5777",
				SellRatio: "0.4223",
				Timestamp: "1695772800000",
			},
		},
	}

	s.r().NoError(err)
	s.assertMarketLongShortRatioInfoEqual(e1, res)
}

func (s *marketTestSuite) assertMarketLongShortRatioInfoEqual(expected, actual *models.MarketLongShortRatioInfo) {
	r := s.r()
	r.Equal(expected.List, actual.List, "List")
}
func (s *marketTestSuite) TestGetServerTime() {
	data := []byte(`{
    "retCode": 0,
    "retMsg": "OK",
    "result": {
        "timeSecond": "1688639403",
        "timeNano": "1688639403423213947"
    },
    "retExtInfo": {},
    "time": 1688639403423
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		e.method = http.MethodGet
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetServerTimeService().Do(newContext())

	e1 := &models.ServerTimeResult{
		TimeSecond: "1688639403",
		TimeNano:   "1688639403423213947",
	}

	s.r().NoError(err)
	s.assertServerTimeEqual(e1, res)
}

func (s *marketTestSuite) assertServerTimeEqual(expected, actual *models.ServerTimeResult) {
	r := s.r()
	r.Equal(expected.TimeSecond, actual.TimeSecond, "TimeSecond")
	r.Equal(expected.TimeNano, actual.TimeNano, "TimeNano")
}
