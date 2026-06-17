package bybit_connector

import (
	"context"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api/handlers"
)

func (s *BybitClientRequest) QueryDcpInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/account/query-dcp-info",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
