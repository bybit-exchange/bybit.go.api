package models

type ProjectListResponse struct {
	RetCode    int               `json:"retCode"`
	RetMsg     string            `json:"retMsg"`
	Result     ProjectListResult `json:"result"`
	RetExtInfo interface{}       `json:"retExtInfo"`
	Time       int               `json:"time"`
}

type ProjectListResult struct {
	List           []ProjectItemDto `json:"list"`
	NextPageCursor string           `json:"nextPageCursor"`
}

type ProjectItemDto struct {
	ProjectId             string        `json:"projectId"`
	ActivityCoin          string        `json:"activityCoin"`
	RewardCoin            string        `json:"rewardCoin"`
	Status                int           `json:"status"`
	TotalReward           string        `json:"totalReward"`
	StakeStartTime        string        `json:"stakeStartTime"`
	StakeEndTime          string        `json:"stakeEndTime"`
	TotalParticipantCount string        `json:"totalParticipantCount"`
	Pools                 []PoolInfoDto `json:"pools"`
}

type PoolInfoDto struct {
	PoolId            string `json:"poolId"`
	StakeCoin         string `json:"stakeCoin"`
	Apr               string `json:"apr"`
	TotalStakedAmount string `json:"totalStakedAmount"`
	ParticipantCount  string `json:"participantCount"`
}

type ActivityLogReqDto struct {
	StakeCoin string `json:"stakeCoin"`
	Type      int    `json:"type"`
	Status    int    `json:"status"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	PageSize  int    `json:"pageSize"`
	Current   int    `json:"current"`
}

type ActivityLogResponse struct {
	RetCode    int                `json:"retCode"`
	RetMsg     string             `json:"retMsg"`
	Result     ActivityLogRespDto `json:"result"`
	RetExtInfo interface{}        `json:"retExtInfo"`
	Time       int                `json:"time"`
}

type ActivityLogRespDto struct {
	List     []ActivityLogItemDto `json:"list"`
	Total    int                  `json:"total"`
	PageSize int                  `json:"pageSize"`
	Current  int                  `json:"current"`
}

type ActivityLogItemDto struct {
	StakeCoin   string `json:"stakeCoin"`
	RewardCoin  string `json:"rewardCoin"`
	Amount      string `json:"amount"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	CreatedTime string `json:"createdTime"`
	OrderId     string `json:"orderId"`
}

type CurrentStakingResponse struct {
	RetCode    int                   `json:"retCode"`
	RetMsg     string                `json:"retMsg"`
	Result     CurrentStakingRespDto `json:"result"`
	RetExtInfo interface{}           `json:"retExtInfo"`
	Time       int                   `json:"time"`
}

type CurrentStakingRespDto struct {
	TotalInvestmentUsd string                  `json:"totalInvestmentUsd"`
	TotalEarningsUsd   string                  `json:"totalEarningsUsd"`
	TodayEarningsUsd   string                  `json:"todayEarningsUsd"`
	List               []CurrentStakingItemDto `json:"list"`
}

type CurrentStakingItemDto struct {
	StakeCoin      string `json:"stakeCoin"`
	RewardCoin     string `json:"rewardCoin"`
	StakeAmount    string `json:"stakeAmount"`
	TotalReward    string `json:"totalReward"`
	AutoRedeemDate string `json:"autoRedeemDate"`
}

type UserHistoryReqDto struct {
	StakeCoin  string `json:"stakeCoin"`
	RewardCoin string `json:"rewardCoin"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	PageSize   int    `json:"pageSize"`
	Current    int    `json:"current"`
}

type UserHistoryResponse struct {
	RetCode    int                `json:"retCode"`
	RetMsg     string             `json:"retMsg"`
	Result     UserHistoryRespDto `json:"result"`
	RetExtInfo interface{}        `json:"retExtInfo"`
	Time       int                `json:"time"`
}

type UserHistoryRespDto struct {
	List     []HistoryItemDto `json:"list"`
	Total    int              `json:"total"`
	PageSize int              `json:"pageSize"`
	Current  int              `json:"current"`
}

type HistoryItemDto struct {
	StakeCoin      string `json:"stakeCoin"`
	RewardCoin     string `json:"rewardCoin"`
	TotalReward    string `json:"totalReward"`
	StakeStartTime string `json:"stakeStartTime"`
	StakeEndTime   string `json:"stakeEndTime"`
}

type PuzzleProjectListResponse struct {
	RetCode    int                     `json:"retCode"`
	RetMsg     string                  `json:"retMsg"`
	Result     PuzzleProjectListResult `json:"result"`
	RetExtInfo interface{}             `json:"retExtInfo"`
	Time       int                     `json:"time"`
}

type PuzzleProjectListResult struct {
	List           []PuzzleProjectItem `json:"list"`
	NextPageCursor string              `json:"nextPageCursor"`
}

type PuzzleProjectItem struct {
	ProjectId        string `json:"projectId"`
	ActivityCoin     string `json:"activityCoin"`
	RewardCoin       string `json:"rewardCoin"`
	Status           int    `json:"status"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	TotalReward      string `json:"totalReward"`
	ParticipantCount string `json:"participantCount"`
}

type GetProjectListResponse struct {
	RetCode    int                  `json:"retCode"`
	RetMsg     string               `json:"retMsg"`
	Result     GetProjectListResult `json:"result"`
	RetExtInfo interface{}          `json:"retExtInfo"`
	Time       int                  `json:"time"`
}

type GetProjectListResult struct {
	List           []ProjectItemDTO `json:"list"`
	NextPageCursor string           `json:"nextPageCursor"`
}

type ProjectItemDTO struct {
	Status                int    `json:"status"`
	ProjectId             string `json:"projectId"`
	ActivityCoin          string `json:"activityCoin"`
	RewardCoin            string `json:"rewardCoin"`
	TotalReward           string `json:"totalReward"`
	ParticipantCount      string `json:"participantCount"`
	RegistrationStartTime string `json:"registrationStartTime"`
	ActivityEndTime       string `json:"activityEndTime"`
}

type GetUserActivityParamsResponse struct {
	RetCode    int                         `json:"retCode"`
	RetMsg     string                      `json:"retMsg"`
	Result     GetUserActivityParamsResult `json:"result"`
	RetExtInfo interface{}                 `json:"retExtInfo"`
	Time       int                         `json:"time"`
}

type GetUserActivityParamsResult struct {
	List []UserActivityItemDTO `json:"list"`
}

type UserActivityItemDTO struct {
	ProjectId    string             `json:"projectId"`
	ActivityCoin string             `json:"activityCoin"`
	TradeTask    TradeTaskDetailDTO `json:"tradeTask"`
}

type TradeTaskDetailDTO struct {
	TradeRequiredAmount   string `json:"tradeRequiredAmount"`
	TradeUnit             string `json:"tradeUnit"`
	TradeToken            string `json:"tradeToken"`
	TradedAmount          string `json:"tradedAmount"`
	MaxRewardAmount       string `json:"maxRewardAmount"`
	EstimatedRewardAmount string `json:"estimatedRewardAmount"`
	RewardCoin            string `json:"rewardCoin"`
}
