# Branch Review Fixes Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Correct the reviewed DTO wire contracts, numeric fixed-loan inventory decoding, timestamp widths, and branch hygiene defects.

**Architecture:** Keep each API family's public models in its existing file. Add focused model tests that assert encoded JSON field sets, deserialize representative server payloads, and use reflection plus a 32-bit build to enforce `int64` timestamps; do not change request routing or `InsLoanDelayLiqStatus`.

**Tech Stack:** Go, standard-library `encoding/json`, `reflect`, and `testing`; Git.

---

### Task 1: Correct Prediction request DTO contracts

**Files:**
- Create: `models/alphaResponse_test.go`
- Modify: `models/alphaResponse.go:120-213`

- [ ] **Step 1: Write the failing JSON contract tests**

Create table-driven tests in `models/alphaResponse_test.go` that marshal populated requests and compare decoded JSON objects:

```go
package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPredictionRequestJSONContracts(t *testing.T) {
	tests := []struct {
		name string
		request interface{}
		want map[string]interface{}
	}{
		{
			name: "position history",
			request: PredictionPositionHistoryRequest{
				TokenId: "token", EventId: "event", Result: 1,
				Days: 7, Limit: 20, PageIndex: 2,
			},
			want: map[string]interface{}{
				"tokenId": "token", "eventId": "event", "result": float64(1),
				"days": float64(7), "limit": float64(20), "pageIndex": float64(2),
			},
		},
		{
			name: "position list",
			request: PredictionPositionListRequest{
				TokenId: "token", EventId: "event", Limit: 20, PageIndex: 2,
			},
			want: map[string]interface{}{
				"tokenId": "token", "eventId": "event",
				"limit": float64(20), "pageIndex": float64(2),
			},
		},
		{
			name: "price history",
			request: PredictionPriceHistoryRequest{TokenId: "token", Interval: "1h"},
			want: map[string]interface{}{"tokenId": "token", "interval": "1h"},
		},
		{
			name: "side market list",
			request: PredictionSideMarketListRequest{
				EventId: "event", SortBy: "volume", MarketType: 1, Limit: 20, PageIndex: 2,
			},
			want: map[string]interface{}{
				"eventId": "event", "sortBy": "volume", "marketType": float64(1),
				"limit": float64(20), "pageIndex": float64(2),
			},
		},
		{
			name: "group stage detail",
			request: PredictionGroupStageDetailRequest{EventType: 1, GroupName: "A"},
			want: map[string]interface{}{"eventType": float64(1), "groupName": "A"},
		},
		{
			name: "match list",
			request: PredictionMatchListRequest{
				EventType: 1, Status: 2, StartTime: 1760340000000, EndTime: 1760426400000,
				StageCode: "GROUP", Limit: 20, PageIndex: 2,
			},
			want: map[string]interface{}{
				"eventType": float64(1), "status": float64(2),
				"startTime": float64(1760340000000), "endTime": float64(1760426400000),
				"stageCode": "GROUP", "limit": float64(20), "pageIndex": float64(2),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payload, err := json.Marshal(test.request)
			if err != nil {
				t.Fatalf("marshal request: %v", err)
			}
			var got map[string]interface{}
			if err := json.Unmarshal(payload, &got); err != nil {
				t.Fatalf("decode marshalled request: %v", err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("JSON fields = %#v, want %#v", got, test.want)
			}
		})
	}
}
```

- [ ] **Step 2: Run the contract test and verify it fails to compile**

Run: `go test ./models -run TestPredictionRequestJSONContracts -count=1`

Expected: FAIL because fields such as `TokenId`, `Result`, `SortBy`, `GroupName`, and `Status` are missing from the current structs.

- [ ] **Step 3: Replace the six request DTO definitions with the documented contracts**

Use these definitions in `models/alphaResponse.go`:

```go
type PredictionPositionHistoryRequest struct {
	TokenId   string                   `json:"tokenId"`
	EventId   string                   `json:"eventId"`
	Result    PredictionPositionResult `json:"result"`
	Days      int                      `json:"days"`
	Limit     int                      `json:"limit"`
	PageIndex int                      `json:"pageIndex"`
}

type PredictionPositionResult int

type PredictionPositionListRequest struct {
	TokenId   string `json:"tokenId"`
	EventId   string `json:"eventId"`
	Limit     int    `json:"limit"`
	PageIndex int    `json:"pageIndex"`
}

type PredictionPriceHistoryRequest struct {
	TokenId  string                  `json:"tokenId"`
	Interval PredictionPriceInterval `json:"interval"`
}

type PredictionSideMarketListRequest struct {
	EventId    string               `json:"eventId"`
	SortBy     PredictionSortBy     `json:"sortBy"`
	MarketType PredictionMarketType `json:"marketType"`
	Limit      int                  `json:"limit"`
	PageIndex  int                  `json:"pageIndex"`
}

type PredictionSortBy string
type PredictionMarketType int

type PredictionGroupStageDetailRequest struct {
	EventType PredictionEventType `json:"eventType"`
	GroupName string              `json:"groupName"`
}

type PredictionMatchListRequest struct {
	EventType PredictionEventType  `json:"eventType"`
	Status    PredictionMatchStatus `json:"status"`
	StartTime int64                `json:"startTime"`
	EndTime   int64                `json:"endTime"`
	StageCode PredictionStageCode  `json:"stageCode"`
	Limit     int                  `json:"limit"`
	PageIndex int                  `json:"pageIndex"`
}

type PredictionMatchStatus int
```

- [ ] **Step 4: Format and run the focused test**

Run: `gofmt -w models/alphaResponse.go models/alphaResponse_test.go`

Run: `go test ./models -run TestPredictionRequestJSONContracts -count=1`

Expected: PASS.

- [ ] **Step 5: Commit the Prediction contract fix**

```bash
git add models/alphaResponse.go models/alphaResponse_test.go
git commit -m "fix: align prediction request models"
```

### Task 2: Separate Crypto Loan fixed inventory from Spot Margin

**Files:**
- Create: `models/cryptoLoanResponse.go`
- Create: `models/cryptoLoanResponse_test.go`

- [ ] **Step 1: Write the failing numeric-term decoding test**

Create `models/cryptoLoanResponse_test.go`:

```go
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
```

- [ ] **Step 2: Run the test and verify it fails to compile**

Run: `go test ./models -run 'Test(CryptoLoan|SpotMargin)Fixed' -count=1`

Expected: FAIL because `CryptoLoanFixedAvailableInventoryResponse` does not exist.

- [ ] **Step 3: Add a Crypto Loan-specific response and retain the Spot Margin type**

Add the following to the new API-specific `models/cryptoLoanResponse.go`:

```go
type CryptoLoanFixedAvailableInventoryResponse struct {
	RetCode    int                                     `json:"retCode"`
	RetMsg     string                                  `json:"retMsg"`
	Result     CryptoLoanFixedAvailableInventoryResult `json:"result"`
	RetExtInfo interface{}                             `json:"retExtInfo"`
	Time       int64                                   `json:"time"`
}

type CryptoLoanFixedAvailableInventoryResult struct {
	Currency           string `json:"currency"`
	Term               int    `json:"term"`
	AnnualRate         string `json:"annualRate"`
	AvailableInventory string `json:"availableInventory"`
	UpdateTime         string `json:"updateTime"`
}
```

Leave `QueryFixedAvailableInventoryResponse` and `FixedAvailableInventoryResult.Term string` unchanged for Spot Margin compatibility.

- [ ] **Step 4: Format and run the focused tests**

Run: `gofmt -w models/cryptoLoanResponse.go models/cryptoLoanResponse_test.go`

Run: `go test ./models -run 'Test(CryptoLoan|SpotMargin)Fixed' -count=1`

Expected: PASS.

- [ ] **Step 5: Commit the response-model separation**

```bash
git add models/cryptoLoanResponse.go models/cryptoLoanResponse_test.go
git commit -m "fix: decode crypto loan fixed inventory term"
```

### Task 3: Make new millisecond timestamp fields 64-bit

**Files:**
- Create: `models/timestampResponse_test.go`
- Modify: `models/alphaResponse.go`
- Modify: `models/referralResponse.go:55-77,127-135`
- Modify: `models/spotMarginResponse.go:329-346`
- Modify: `models/spotXResponse.go`
- Modify: `models/taxResponse.go:8-14`

- [ ] **Step 1: Write a reflection test for every affected timestamp field**

Create `models/timestampResponse_test.go` with a helper that checks the named field is `int64`, then enumerate all new response `Time` fields plus the timestamp-bearing records and requests:

```go
package models

import (
	"reflect"
	"testing"
)

func TestReviewedTimestampFieldsUseInt64(t *testing.T) {
	tests := []struct {
		value interface{}
		field string
	}{
		{PredictionBuyResponse{}, "Time"},
		{PredictionEngineStatusResponse{}, "Time"},
		{PredictionEventDetailResponse{}, "Time"},
		{PredictionOrderBookResponse{}, "Time"},
		{PredictionOrderEstimateResponse{}, "Time"},
		{PredictionOrderListResponse{}, "Time"},
		{PredictionPayTokenListResponse{}, "Time"},
		{PredictionPortfolioSummaryResponse{}, "Time"},
		{PredictionPositionHistoryResponse{}, "Time"},
		{PredictionPositionListResponse{}, "Time"},
		{PredictionPriceHistoryResponse{}, "Time"},
		{PredictionSellResponse{}, "Time"},
		{PredictionSideMarketListResponse{}, "Time"},
		{PredictionGroupStageDetailResponse{}, "Time"},
		{PredictionMatchListResponse{}, "Time"},
		{PredictionTimelineStagesResponse{}, "Time"},
		{PredictionTokenPriceResponse{}, "Time"},
		{ReferralCommissionRecord{}, "PaymentTime"},
		{GetReferralUserIdentityProfileResponse{}, "Time"},
		{ReferralUserInvitationRecord{}, "SignTime"},
		{QueryFixedBorrowPenaltyInterestResponse{}, "Time"},
		{CancelFixedBorrowOrderResponse{}, "Time"},
		{ProjectListResponse{}, "Time"},
		{ActivityLogResponse{}, "Time"},
		{CurrentStakingResponse{}, "Time"},
		{UserHistoryResponse{}, "Time"},
		{PuzzleProjectListResponse{}, "Time"},
		{GetProjectListResponse{}, "Time"},
		{GetUserActivityParamsResponse{}, "Time"},
		{TaxBatchCreateRequest{}, "StartTime"},
		{TaxBatchCreateRequest{}, "EndTime"},
		{PredictionMatchListRequest{}, "StartTime"},
		{PredictionMatchListRequest{}, "EndTime"},
	}

	for _, test := range tests {
		t.Run(reflect.TypeOf(test.value).Name()+"."+test.field, func(t *testing.T) {
			field, ok := reflect.TypeOf(test.value).FieldByName(test.field)
			if !ok {
				t.Fatalf("field %s is missing", test.field)
			}
			if field.Type.Kind() != reflect.Int64 {
				t.Fatalf("type = %v, want int64", field.Type)
			}
		})
	}
}
```

- [ ] **Step 2: Run the timestamp test and verify it fails**

Run: `go test ./models -run TestReviewedTimestampFieldsUseInt64 -count=1`

Expected: FAIL on the first reviewed field still declared as `int`.

- [ ] **Step 3: Change only timestamp-semantic fields to `int64`**

In the listed model files, change each field enumerated by the test from `int` to `int64`. Keep counters, enum values, paging fields, status fields, and durations as `int`.

- [ ] **Step 4: Format, test, and verify a 32-bit build**

Run: `gofmt -w models/alphaResponse.go models/referralResponse.go models/spotMarginResponse.go models/spotXResponse.go models/taxResponse.go models/timestampResponse_test.go`

Run: `go test ./models -run 'Test(ReviewedTimestampFieldsUseInt64|PredictionRequestJSONContracts|CryptoLoanFixed)' -count=1`

Expected: PASS.

Run: `GOOS=linux GOARCH=386 go test -c ./models -o /tmp/bybit-models-linux-386.test`

Expected: successful cross-compilation, including millisecond literals beyond the 32-bit `int` range.

- [ ] **Step 5: Commit the timestamp-width corrections**

```bash
git add models/alphaResponse.go models/referralResponse.go models/spotMarginResponse.go models/spotXResponse.go models/taxResponse.go models/timestampResponse_test.go
git commit -m "fix: use int64 for new millisecond timestamps"
```

### Task 4: Clean branch artifacts and run full verification

**Files:**
- Modify: `RELEASE_NOTE.md:268`
- Delete: `docs/superpowers/specs/2026-11-09-remove-rwa-stocks-apis-design.md`

- [ ] **Step 1: Remove the trailing space and obsolete design file**

Change the final release-note line to exactly:

```text
Deprecated leverage token endpoints and legacy crypto loan endpoints
```

Delete `docs/superpowers/specs/2026-11-09-remove-rwa-stocks-apis-design.md`, which documents already completed work and is unrelated to the SDK update being reviewed.

- [ ] **Step 2: Run formatting and the complete unit test suite**

Run: `gofmt -w models/alphaResponse.go models/alphaResponse_test.go models/cryptoLoanResponse.go models/cryptoLoanResponse_test.go models/referralResponse.go models/spotMarginResponse.go models/spotXResponse.go models/taxResponse.go models/timestampResponse_test.go`

Run: `go test -count=1 ./...`

Expected: PASS for every package.

- [ ] **Step 3: Run static and patch hygiene checks**

Run: `go vet ./...`

Expected: exit code 0 with no diagnostics.

Run: `git diff --check main...HEAD`

Expected: exit code 0 with no whitespace errors after the cleanup commit is included.

Run: `GOOS=linux GOARCH=386 go test -c ./models -o /tmp/bybit-models-linux-386.test`

Expected: successful 32-bit cross-compilation of the model tests.

- [ ] **Step 4: Verify removed RWA stocks interfaces remain absent and the known environment issue is untouched**

Run: `rg -n '/v5/rwa/stocks/' . --glob '!docs/superpowers/plans/*' --glob '!docs/superpowers/specs/*'`

Expected: no matches.

Run: `git diff main...HEAD -- ins_loan.go | rg 'delay-liq-status|MethodGet'`

Expected: the existing `GET /v5/ins-loan/delay-liq-status` implementation remains present and unchanged by this fix.

- [ ] **Step 5: Commit branch cleanup and the implementation plan**

```bash
git add RELEASE_NOTE.md docs/superpowers/plans/2026-11-09-branch-review-fixes.md
git add -u docs/superpowers/specs/2026-11-09-remove-rwa-stocks-apis-design.md
git commit -m "chore: clean reviewed branch artifacts"
```

No state-changing testnet endpoints are called during implementation. The previously completed authenticated read-only smoke run remains the live-service evidence, and no API credential is written into the repository.
