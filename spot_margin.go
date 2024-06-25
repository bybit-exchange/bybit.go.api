package bybit_connector

import (
	"context"
	"errors"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api/handlers"
)

type SpotMarginClient struct {
	c      *Client
	isUta  bool
	params map[string]interface{}
}

func (s *SpotMarginClient) GetSpotMarginData(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) GetSpotMarginCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) GetSpotMarginBorrowCoin(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) GetSpotMarginInterests(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if s.isUta {
		return nil, errors.New("this function only works for classical accounts")
	}
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/spot-cross-margin-trade/loan-info",
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

func (s *SpotMarginClient) GetSpotMarginLoanAccountInfo(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) GetSpotMarginBorrowOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) GetSpotMarginRepaymentOrders(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) BorrowSpotMarginLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) RepaySpotMarginLoan(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) SetSpotMarginLeverage(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) GetSpotMarginState(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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

func (s *SpotMarginClient) ToggleSpotMarginTrade(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
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
