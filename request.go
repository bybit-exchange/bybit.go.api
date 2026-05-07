package bybit_connector

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type secType int

const (
	secTypeNone   secType = iota
	secTypeSigned         // private request
)

type params map[string]interface{}

// request define an API request
type request struct {
	method     string
	endpoint   string
	query      url.Values
	recvWindow string
	secType    secType
	header     http.Header
	params     []byte
	fullURL    string
	body       io.Reader
}

// addParam add param with key/value to query string
func (r *request) addParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// setParam set param with key/value to query string
func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set params with key/values to query string or body
func (r *request) setParams(m params) *request {
	if r.method == http.MethodGet {
		for k, v := range m {
			r.setParam(k, v)
		}
	} else if r.method == http.MethodPost {
		jsonData, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		r.params = jsonData
	}

	return r
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	return nil
}

// WithRecvWindow Append `WithRecvWindow(insert_recvWindow)` to request to modify the default recvWindow value
func WithRecvWindow(recvWindow string) RequestOption {
	return func(r *request) {
		r.recvWindow = recvWindow
	}
}

// RequestOption define option type for request
type RequestOption func(*request)
