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
