package bybit_connector

import (
	"context"
	"net/http"

	"github.com/0xArchiviste/bybit.go.api/handlers"
)

func (s *BybitClientRequest) GetEarnProductInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/earn/product",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) RedeemEarnOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/earn/place-order",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetEarnRedeemOrder(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/earn/order",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetEarnRedeemPosition(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/earn/position",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetYieldHistory
func (s *BybitClientRequest) GetYieldHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/earn/yield",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetHourlyYieldHistory
func (s *BybitClientRequest) GetHourlyYieldHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/earn/hourly-yield",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
