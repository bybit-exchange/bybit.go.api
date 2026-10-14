package bybit_connector

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRequestSetParamsGETCollections(t *testing.T) {
	r := &request{method: http.MethodGet}
	r.setParams(params{
		"coins": []string{"BTC", "ETH"},
		"terms": [2]int{7, 14},
		"limit": 50,
		"empty": []string{},
	})

	if got, want := r.query["coins"], []string{"BTC", "ETH"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("coins = %v, want %v", got, want)
	}
	if got, want := r.query["terms"], []string{"7", "14"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("terms = %v, want %v", got, want)
	}
	if got := r.query.Get("limit"); got != "50" {
		t.Fatalf("limit = %q, want %q", got, "50")
	}
	if _, ok := r.query["empty"]; ok {
		t.Fatal("empty slice must not emit a query parameter")
	}
}
