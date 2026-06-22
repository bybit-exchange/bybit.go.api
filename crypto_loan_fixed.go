package bybit_connector

import (
	"context"
	"net/http"
)

func (s *BybitClientRequest) GetCryptoLoanFixedSupplyOrderQuote(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/crypto-loan-fixed/supply-order-quote",
		secType:  secTypeNone,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
