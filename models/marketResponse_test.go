package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOptionBaseCoinsResponseUsesInt64Timestamps(t *testing.T) {
	var response OptionBaseCoinsResponse

	timeField, ok := reflect.TypeOf(response).FieldByName("Time")
	if !ok {
		t.Fatal("OptionBaseCoinsResponse.Time field is missing")
	}
	if timeField.Type.Kind() != reflect.Int64 {
		t.Fatalf("OptionBaseCoinsResponse.Time type = %v, want int64", timeField.Type)
	}

	optionOnlineTimeField, ok := reflect.TypeOf(OptionBaseCoin{}).FieldByName("OptionOnlineTime")
	if !ok {
		t.Fatal("OptionBaseCoin.OptionOnlineTime field is missing")
	}
	if optionOnlineTimeField.Type.Kind() != reflect.Int64 {
		t.Fatalf("OptionBaseCoin.OptionOnlineTime type = %v, want int64", optionOnlineTimeField.Type)
	}

	const payload = `{"time":1760340000000,"result":{"list":[{"optionOnlineTime":1760340000000}]}}`
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal option base coins response: %v", err)
	}
	if got, want := response.Time, int64(1760340000000); got != want {
		t.Fatalf("Time = %d, want %d", got, want)
	}
	if got, want := response.Result.List[0].OptionOnlineTime, int64(1760340000000); got != want {
		t.Fatalf("OptionOnlineTime = %d, want %d", got, want)
	}
}

func TestNewResponseModelsUseInt64Timestamps(t *testing.T) {
	tests := []struct {
		name     string
		response interface{}
	}{
		{name: "ListAutoSavingsResponse", response: &ListAutoSavingsResponse{}},
		{name: "DelayLiqStatusResponse", response: &DelayLiqStatusResponse{}},
		{name: "GetFlexibleAvailableInventoryResponse", response: &GetFlexibleAvailableInventoryResponse{}},
		{name: "QueryFixedAvailableInventoryResponse", response: &QueryFixedAvailableInventoryResponse{}},
		{name: "QueryReferralCodeApiResponse", response: &QueryReferralCodeApiResponse{}},
	}

	const (
		payload  = `{"time":1760340000000}`
		wantTime = int64(1760340000000)
	)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responseType := reflect.TypeOf(test.response).Elem()
			timeField, ok := responseType.FieldByName("Time")
			if !ok {
				t.Fatalf("%s.Time field is missing", test.name)
			}
			if timeField.Type.Kind() != reflect.Int64 {
				t.Fatalf("%s.Time type = %v, want int64", test.name, timeField.Type)
			}

			if err := json.Unmarshal([]byte(payload), test.response); err != nil {
				t.Fatalf("unmarshal %s: %v", test.name, err)
			}
			if got := reflect.ValueOf(test.response).Elem().FieldByName("Time").Int(); got != wantTime {
				t.Fatalf("%s.Time = %d, want %d", test.name, got, wantTime)
			}
		})
	}
}
