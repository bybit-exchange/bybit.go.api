package models

import (
	"encoding/json"
	"testing"
)

func TestCryptoLoanFixedAvailableInventoryDecodesNumericTerm(t *testing.T) {
	const payload = `{"retCode":0,"result":{"currency":"USDT","term":30,"annualRate":"0.08","availableInventory":"1000","updateTime":"1760340000000"},"time":1760340000000}`

	var response CryptoLoanFixedAvailableInventoryResponse
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal fixed inventory response: %v", err)
	}
	if got, want := response.Result.Term, 30; got != want {
		t.Fatalf("Term = %d, want %d", got, want)
	}
}

func TestSpotMarginFixedInventoryKeepsStringTerm(t *testing.T) {
	var response QueryFixedAvailableInventoryResponse
	const payload = `{"result":{"term":"30"}}`
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal spot margin fixed inventory response: %v", err)
	}
	if got, want := response.Result.Term, "30"; got != want {
		t.Fatalf("Term = %q, want %q", got, want)
	}
}
