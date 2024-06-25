package bybit_connector

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/bybit-exchange/bybit.go.api/handlers"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ServerResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo struct{}    `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

// Client define API client
type Client struct {
	APIKey     string
	APISecret  string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

type doFunc func(req *http.Request) (*http.Response, error)

type ClientOption func(*Client)

// WithDebug print more details in debug mode
func WithDebug(debug bool) ClientOption {
	return func(c *Client) {
		c.Debug = debug
	}
}

// WithBaseURL is a client option to set the base URL of the Bybit HTTP client.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", " ")
	return string(s)
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func GetCurrentTime() int64 {
	now := time.Now()
	unixNano := now.UnixNano()
	timeStamp := unixNano / int64(time.Millisecond)
	return timeStamp
}

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// NewBybitHttpClient NewClient Create client function for initialising new Bybit client
func NewBybitHttpClient(apiKey string, APISecret string, options ...ClientOption) *Client {
	c := &Client{
		APIKey:     apiKey,
		APISecret:  APISecret,
		BaseURL:    MAINNET,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, Name, log.LstdFlags),
	}

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	queryString := r.query.Encode()
	header := http.Header{}
	body := &bytes.Buffer{}
	if r.params != nil {
		body = bytes.NewBuffer(r.params)
	}
	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))

	if r.secType == secTypeSigned {
		timeStamp := GetCurrentTime()
		header.Set(signTypeKey, "2")
		header.Set(apiRequestKey, c.APIKey)
		header.Set(timestampKey, strconv.FormatInt(timeStamp, 10))
		if r.recvWindow == "" {
			r.recvWindow = "5000"
		}
		header.Set(recvWindowKey, r.recvWindow)

		var signatureBase []byte
		if r.method == "POST" {
			header.Set("Content-Type", "application/json")
			signatureBase = []byte(strconv.FormatInt(timeStamp, 10) + c.APIKey + r.recvWindow + string(r.params[:]))
		} else {
			signatureBase = []byte(strconv.FormatInt(timeStamp, 10) + c.APIKey + r.recvWindow + queryString)
		}
		hmac256 := hmac.New(sha256.New, []byte(c.APISecret))
		hmac256.Write(signatureBase)
		signature := hex.EncodeToString(hmac256.Sum(nil))
		header.Set(signatureKey, signature)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, body)
	r.fullURL = fullURL
	r.body = body
	r.header = header
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		var (
			apiErr = new(handlers.APIError)
		)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) NewInstrumentsInfoService() *InstrumentsInfoService {
	return &InstrumentsInfoService{c: c}
}

// NewMarketKlineService Market Kline Endpoints
func (c *Client) NewMarketKlineService() *MarketKlinesService {
	return &MarketKlinesService{c: c}
}

// NewMarketMarkPriceKlineService Market Mark Price Kline Endpoints
func (c *Client) NewMarketMarkPriceKlineService() *MarketMarkPriceKlineService {
	return &MarketMarkPriceKlineService{c: c}
}

// NewMarketIndexPriceKlineService Market Index Price Kline Endpoints
func (c *Client) NewMarketIndexPriceKlineService() *MarketIndexPriceKlineService {
	return &MarketIndexPriceKlineService{c: c}
}

// NewMarketPremiumIndexPriceKlineService Market Premium Index Price Kline Endpoints
func (c *Client) NewMarketPremiumIndexPriceKlineService() *MarketPremiumIndexPriceKlineService {
	return &MarketPremiumIndexPriceKlineService{c: c}
}

func (c *Client) NewOrderBookService() *MarketOrderBookService {
	return &MarketOrderBookService{c: c}
}

func (c *Client) NewTickersService() *MarketTickersService {
	return &MarketTickersService{c: c}
}

func (c *Client) NewFundingTatesService() *MarketFundingRatesService {
	return &MarketFundingRatesService{c: c}
}

func (c *Client) NewGetPublicRecentTradesService() *GetPublicRecentTradesService {
	return &GetPublicRecentTradesService{c: c}
}

// GetOpenInterestsServicdde
func (c *Client) NewGetOpenInterestsService() *GetOpenInterestsService {
	return &GetOpenInterestsService{c: c}
}

// GetHistoricalVolatilityService
func (c *Client) NewGetHistoricalVolatilityService() *GetHistoricalVolatilityService {
	return &GetHistoricalVolatilityService{c: c}
}

// GetInsuranceInfoService
func (c *Client) NewGetInsuranceInfoService() *GetInsuranceInfoService {
	return &GetInsuranceInfoService{c: c}
}

// GetRiskLimitService
func (c *Client) NewGetRiskLimitService() *GetRiskLimitService {
	return &GetRiskLimitService{c: c}
}

// GetDeliveryPriceService
func (c *Client) NewGetDeliveryPriceService() *GetDeliveryPriceService {
	return &GetDeliveryPriceService{c: c}
}

// GetMarketLSRatioService
func (c *Client) NewGetMarketLSRatioService() *GetMarketLSRatioService {
	return &GetMarketLSRatioService{c: c}
}

// GetServerTimeService
func (c *Client) NewGetServerTimeService() *GetServerTimeService {
	return &GetServerTimeService{c: c}
}

// NewPlaceOrderService Trade Endpoints
func (c *Client) NewPlaceOrderService(category, symbol, side, orderType, qty string) *Order {
	return &Order{
		c:         c,
		category:  category,
		symbol:    symbol,
		side:      side,
		orderType: orderType,
		qty:       qty,
	}
}

func (c *Client) NewTradeService(params map[string]interface{}) *TradeClient {
	return &TradeClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewPositionService(params map[string]interface{}) *PositionClient {
	return &PositionClient{
		c:      c,
		params: params,
	}

}

func (c *Client) NewPreUpgradeService(params map[string]interface{}) *PreUpgradeClient {
	return &PreUpgradeClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewAccountService(params map[string]interface{}) *AccountClient {
	return &AccountClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewAccountServiceNoParams() *AccountClient {
	return &AccountClient{
		c: c,
	}
}

func (c *Client) NewAssetService(params map[string]interface{}) *AssetClient {
	return &AssetClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewUserService(params map[string]interface{}) *UserServiceClient {
	return &UserServiceClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewUserServiceNoParams() *UserServiceClient {
	return &UserServiceClient{
		c: c,
	}
}

func (c *Client) NewBrokerService(params map[string]interface{}) *BrokerServiceClient {
	return &BrokerServiceClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewLendingService(params map[string]interface{}) *LendingServiceClient {
	return &LendingServiceClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewLendingServiceNoParams() *LendingServiceClient {
	return &LendingServiceClient{
		c: c,
	}
}

func (c *Client) NewSpotLeverageService(params map[string]interface{}) *SpotLeverageClient {
	return &SpotLeverageClient{
		c:      c,
		params: params,
	}
}

func (c *Client) NewSpotMarginDataService(params map[string]interface{}, isUta bool) *SpotMarginClient {
	return &SpotMarginClient{
		c:      c,
		isUta:  isUta,
		params: params,
	}
}
