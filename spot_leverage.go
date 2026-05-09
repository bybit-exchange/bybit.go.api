package bybit_connector

import (
	"context"
	"github.com/0xArchiviste/bybit.go.api/handlers"
	"net/http"
)

// Deprecated: GetLeverageTokenInfo is deprecated.
func (s *BybitClientRequest) GetLeverageTokenInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-lever-token/info",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetLeverageTokenOrders is deprecated.
func (s *BybitClientRequest) GetLeverageTokenOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-lever-token/order-record",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetLeverageTokenMarket is deprecated.
func (s *BybitClientRequest) GetLeverageTokenMarket(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-lever-token/reference",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: PurchaseLeverageToken is deprecated.
func (s *BybitClientRequest) PurchaseLeverageToken(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/spot-lever-token/purchase",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: RedeemLeverageToken is deprecated.
func (s *BybitClientRequest) RedeemLeverageToken(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/spot-lever-token/redeem",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
