package bybit_connector

import (
	"context"
	"net/http"
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
