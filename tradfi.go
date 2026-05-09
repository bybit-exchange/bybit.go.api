package bybit_connector

import (
	"context"
)

// Symbol types for TradFi-adjacent products (see Bybit v5 enum: symbolType).
const (
	SymbolTypeXStocks   = "xstocks"   // Tokenized equities on spot
	SymbolTypeCommodity = "commodity" // Commodity linear contracts
)

func mergeClientParams(s *BybitClientRequest, extra map[string]interface{}) *BybitClientRequest {
	out := make(map[string]interface{})
	for k, v := range s.params {
		out[k] = v
	}
	for k, v := range extra {
		out[k] = v
	}
	return &BybitClientRequest{c: s.c, params: out, isUta: s.isUta}
}

// GetTradFiSpotInstrumentsInfo lists xStocks (tokenized equity) spot instruments.
// Equivalent to instruments-info with category=spot and symbolType=xstocks.
func (s *BybitClientRequest) GetTradFiSpotInstrumentsInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category":   "spot",
		"symbolType": SymbolTypeXStocks,
	}).GetInstrumentInfo(ctx, opts...)
}

// GetTradFiSpotTickers queries tickers for xStocks spot pairs (params e.g. symbol).
func (s *BybitClientRequest) GetTradFiSpotTickers(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category":   "spot",
		"symbolType": SymbolTypeXStocks,
	}).GetMarketTickers(ctx, opts...)
}

// GetTradFiSpotOrderBook queries order book for an xStocks spot symbol.
func (s *BybitClientRequest) GetTradFiSpotOrderBook(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category":   "spot",
		"symbolType": SymbolTypeXStocks,
	}).GetOrderBookInfo(ctx, opts...)
}

// GetTradFiSpotRecentTrades queries recent trades for an xStocks spot symbol.
func (s *BybitClientRequest) GetTradFiSpotRecentTrades(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category":   "spot",
		"symbolType": SymbolTypeXStocks,
	}).GetPublicRecentTrades(ctx, opts...)
}

// GetTradFiSpotKline queries klines for xStocks spot (pass interval, symbol, etc. in merged params).
func (s *BybitClientRequest) GetTradFiSpotKline(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category":   "spot",
		"symbolType": SymbolTypeXStocks,
	}).GetMarketKline(ctx, opts...)
}

// GetCommodityLinearInstrumentsInfo lists linear instruments with symbolType=commodity.
func (s *BybitClientRequest) GetCommodityLinearInstrumentsInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category":   "linear",
		"symbolType": SymbolTypeCommodity,
	}).GetInstrumentInfo(ctx, opts...)
}

// PlaceTradFiSpotOrder places a spot order for an xStocks pair via unified v5 order API.
// Request params must include symbol, side, orderType, qty, etc.; category=spot is set automatically.
func (s *BybitClientRequest) PlaceTradFiSpotOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category": "spot",
	}).PlaceOrder(ctx, opts...)
}

// GetTradFiSpotOpenOrders lists open spot orders; use for xStocks symbols (category=spot preset).
func (s *BybitClientRequest) GetTradFiSpotOpenOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category": "spot",
	}).GetOpenOrders(ctx, opts...)
}

// GetTradFiSpotOrderHistory lists historical spot orders (xStocks: category=spot preset).
func (s *BybitClientRequest) GetTradFiSpotOrderHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	return mergeClientParams(s, map[string]interface{}{
		"category": "spot",
	}).GetOrderHistory(ctx, opts...)
}
