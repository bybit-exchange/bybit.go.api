package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPredictionRequestJSONContracts(t *testing.T) {
	tests := []struct {
		name    string
		request interface{}
		want    map[string]interface{}
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
			name:    "price history",
			request: PredictionPriceHistoryRequest{TokenId: "token", Interval: "1h"},
			want:    map[string]interface{}{"tokenId": "token", "interval": "1h"},
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
			name:    "group stage detail",
			request: PredictionGroupStageDetailRequest{EventType: 1, GroupName: "A"},
			want:    map[string]interface{}{"eventType": float64(1), "groupName": "A"},
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
