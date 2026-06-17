package bybit_connector

import (
	"context"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api/handlers"
)

func (s *BybitClientRequest) CloseComboBot(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/fcombobot/close",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
