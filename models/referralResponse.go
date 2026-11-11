package models

type EmptyCampaignRequest struct {
}

type CampaignResponse struct {
	RetCode int            `json:"retCode"`
	RetMsg  string         `json:"retMsg"`
	Result  CampaignResult `json:"result"`
}

type CampaignResult struct {
	RebateRule ReferralCampaignRebateRatioRules `json:"rebateRule"`
	Tasks      []string                         `json:"tasks"`
}

type ReferralCampaignRebateRatioRules struct {
	DepositAmount        int                               `json:"depositAmount"`
	TradeAmount          int                               `json:"tradeAmount"`
	RebateRatioRules     []ReferralCampaignRebateRatioRule `json:"rebateRatioRules"`
	RebateDurationDays   int                               `json:"rebateDurationDays"`
	TradefiDepositAmount int                               `json:"tradefiDepositAmount"`
	TradefiTradeAmount   int                               `json:"tradefiTradeAmount"`
}

type ReferralCampaignRebateRatioRule struct {
	RebateRatio          int  `json:"rebateRatio"`
	ReferralCount        int  `json:"referralCount"`
	HasTradeCondition    bool `json:"hasTradeCondition"`
	TradeAmountCondition int  `json:"tradeAmountCondition"`
}

type GetReferralCommissionRecordsRequest struct {
	UserIdList  []int        `json:"userIdList"`
	StartTime   string       `json:"startTime"`
	EndTime     string       `json:"endTime"`
	Page        int          `json:"page"`
	PageSize    int          `json:"pageSize"`
	ProductLine BusinessUnit `json:"productLine"`
}

type BusinessUnit int

type GetReferralCommissionRecordsResponse struct {
	RetCode int                                `json:"retCode"`
	RetMsg  string                             `json:"retMsg"`
	Result  GetReferralCommissionRecordsResult `json:"result"`
}

type GetReferralCommissionRecordsResult struct {
	Records []ReferralCommissionRecord `json:"records"`
	Total   int                        `json:"total"`
}

type ReferralCommissionRecord struct {
	UserId                 string           `json:"userId"`
	PaymentTime            int64            `json:"paymentTime"`
	Rate                   string           `json:"rate"`
	BusinessUnit           BusinessUnit     `json:"businessUnit"`
	CommissionOriginCoin   string           `json:"commissionOriginCoin"`
	CommissionOriginAmount string           `json:"commissionOriginAmount"`
	CommissionUsdtAmount   string           `json:"commissionUsdtAmount"`
	CommissionStatus       CommissionStatus `json:"commissionStatus"`
	IsBoostM2              bool             `json:"isBoostM2"`
}

type CommissionStatus int

type GetReferralUserIdentityProfileRequest struct {
}

type GetReferralUserIdentityProfileResponse struct {
	RetCode    int                                  `json:"retCode"`
	RetMsg     string                               `json:"retMsg"`
	Result     GetReferralUserIdentityProfileResult `json:"result"`
	RetExtInfo interface{}                          `json:"retExtInfo"`
	Time       int64                                `json:"time"`
}

type GetReferralUserIdentityProfileResult struct {
	IsAffiliate            bool                                 `json:"isAffiliate"`
	IsM1                   bool                                 `json:"isM1"`
	InviterAffiliateId     int                                  `json:"inviterAffiliateId"`
	InviterAffiliateUserId int                                  `json:"inviterAffiliateUserId"`
	InviterUserId          int                                  `json:"inviterUserId"`
	CommissionInfo         ReferralIdentityCommissionInfoResult `json:"commissionInfo"`
}

type ReferralIdentityCommissionInfoResult struct {
	TotalCommission                         string `json:"totalCommission"`
	YesterdayCommission                     string `json:"yesterdayCommission"`
	TotalFriends                            int    `json:"totalFriends"`
	QualifiedFriends                        int    `json:"qualifiedFriends"`
	WalletBalance                           string `json:"walletBalance"`
	QualifiedFriendsQuarterTotalTradeVolume string `json:"qualifiedFriendsQuarterTotalTradeVolume"`
	CurQuarterActiveFriendsCount            int    `json:"curQuarterActiveFriendsCount"`
	HasOldReferral                          int    `json:"hasOldReferral"`
	DegradeType                             int    `json:"degradeType"`
	CommissionRate                          string `json:"commissionRate"`
	UpgradeUnifiedAccount                   int    `json:"upgradeUnifiedAccount"`
	WithdrawalSwitch                        int    `json:"withdrawalSwitch"`
	LastQuarterQualifiedFriendsCount        int    `json:"lastQuarterQualifiedFriendsCount"`
}

type GetReferralUserInvitationsRequest struct {
	UserIdList   []int        `json:"userIdList"`
	StartTime    string       `json:"startTime"`
	EndTime      string       `json:"endTime"`
	Page         int          `json:"page"`
	PageSize     int          `json:"pageSize"`
	FriendStatus FriendStatus `json:"friendStatus"`
}

type FriendStatus int

type GetReferralUserInvitationsResponse struct {
	RetCode int                              `json:"retCode"`
	RetMsg  string                           `json:"retMsg"`
	Result  GetReferralUserInvitationsResult `json:"result"`
}

type GetReferralUserInvitationsResult struct {
	Records []ReferralUserInvitationRecord `json:"records"`
	Total   int                            `json:"total"`
}

type ReferralUserInvitationRecord struct {
	UserIdMask        string       `json:"userIdMask"`
	SignTime          int64        `json:"signTime"`
	UserStatus        FriendStatus `json:"userStatus"`
	TotalCommission   string       `json:"totalCommission"`
	UserId            string       `json:"userId"`
	TaskProgress      string       `json:"taskProgress"`
	NotShowCommission bool         `json:"notShowCommission"`
	TotalBaseRewards  string       `json:"totalBaseRewards"`
}
