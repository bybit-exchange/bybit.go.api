package bybit_connector

import (
	"context"
	"errors"
	"net/http"

	"github.com/fqds/bybit-go-api-mf/handlers"
)

func (s *BybitClientRequest) GetSpotMarginData(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	var endpoint string
	if !s.isUta {
		endpoint = "/v5/spot-margin-trade/data"
	} else {
		endpoint = "/v5/spot-cross-margin-trade/data"
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetTieredCollateralData(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/collateral",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpotMarginInterests(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	var endpoint string
	if s.isUta {
		endpoint = "/v5/spot-margin-trade/interest-rate-history"
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		endpoint = "/v5/spot-cross-margin-trade/loan-info"
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) SetSpotMarginLeverage(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if !s.isUta {
		return nil, errors.New("this function only works for UTA accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/spot-margin-trade/set-leverage",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BybitClientRequest) GetSpotMarginState(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if !s.isUta {
		return nil, errors.New("this function only works for UTA accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/state",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BybitClientRequest) ToggleSpotMarginTrade(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	var endpoint string
	if !s.isUta {
		endpoint = "/v5/spot-margin-trade/switch-mode"
	} else {
		endpoint = "/v5/spot-cross-margin-trade/data"
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Deprecated: GetSpotMarginCoin is deprecated.
func (s *BybitClientRequest) GetSpotMarginCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/pledge-token",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetSpotMarginBorrowCoin is deprecated.
func (s *BybitClientRequest) GetSpotMarginBorrowCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/borrow-token",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetSpotMarginLoanAccountInfo is deprecated.
func (s *BybitClientRequest) GetSpotMarginLoanAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/account",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetSpotMarginBorrowOrders is deprecated.
func (s *BybitClientRequest) GetSpotMarginBorrowOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/orders",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: GetSpotMarginRepaymentOrders is deprecated.
func (s *BybitClientRequest) GetSpotMarginRepaymentOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/repay-history",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: BorrowSpotMarginLoan is deprecated.
func (s *BybitClientRequest) BorrowSpotMarginLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/spot-cross-margin-trade/loan",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

// Deprecated: RepaySpotMarginLoan is deprecated.
func (s *BybitClientRequest) RepaySpotMarginLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/spot-cross-margin-trade/repay",
		secType:  secTypeSigned,
	}
	data, err := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}

func (s *BybitClientRequest) GetSpotMarginAutoRepayMode(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if !s.isUta {
		return nil, errors.New("this function only works for UTA accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-margin-trade/get-auto-repay-mode",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BybitClientRequest) SetSpotMarginAutoRepayMode(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if !s.isUta {
		return nil, errors.New("this function only works for UTA accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/spot-margin-trade/set-auto-repay-mode",
		secType:  secTypeSigned,
	}
	r.setParams(s.params)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
