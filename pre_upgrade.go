package bybit_connector

import (
	"context"
	"github.com/0xArchiviste/bybit.go.api/handlers"
	"net/http"
)

func (s *BybitClientRequest) GetPreUpgradeOrderHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/pre-upgrade/order/history",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetPreUpgradeTradeHistory(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/pre-upgrade/execution/list",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetPreUpgradeClosedPnl(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/pre-upgrade/position/closed-pnl",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetPreUpgradeTransactionLog(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/pre-upgrade/account/transaction-log",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetPreUpgradeOptionDeliveryRecord(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/pre-upgrade/asset/delivery-record",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetPreUpgradeUsdcSettlement(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/pre-upgrade/asset/settlement-record",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
