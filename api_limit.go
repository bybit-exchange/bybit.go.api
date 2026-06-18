package bybit_connector

import (
	"context"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api/handlers"
)

// SetApiRateLimit sets the API rate limit
func (s *BybitClientRequest) SetApiRateLimit(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/apilimit/set",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// GetApiRateLimit queries the API rate limit
func (s *BybitClientRequest) GetApiRateLimit(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/apilimit/query",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// QueryCap
func (s *BybitClientRequest) QueryCap(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/apilimit/query-cap",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// QueryAll
func (s *BybitClientRequest) QueryAll(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/apilimit/query-all",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) QueryBrokerAllUidDetails(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/broker/apilimit/query-all",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) QueryBrokerCap(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/broker/apilimit/query-cap",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) SetBrokerApiLimit(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/broker/apilimit/set",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
