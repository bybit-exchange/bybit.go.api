package models

// UserVolumeInfo represents the structure for a user's trading and deposit volume information.
type UserVolumeInfo struct {
	UID                 string `json:"uid"`
	VipLevel            string `json:"vipLevel"`
	TakerVol30Day       string `json:"takerVol30Day"`
	MakerVol30Day       string `json:"makerVol30Day"`
	TradeVol30Day       string `json:"tradeVol30Day"`
	DepositAmount30Day  string `json:"depositAmount30Day"`
	TakerVol365Day      string `json:"takerVol365Day"`
	MakerVol365Day      string `json:"makerVol365Day"`
	TradeVol365Day      string `json:"tradeVol365Day"`
	DepositAmount365Day string `json:"depositAmount365Day"`
	TotalWalletBalance  string `json:"totalWalletBalance"` // This should be an integer value representing a range, not a string.
	DepositUpdateTime   string `json:"depositUpdateTime"`
	VolUpdateTime       string `json:"volUpdateTime"`
}

type SubMember struct {
	AccountMode int    `json:"accountMode"`
	MemberType  int    `json:"memberType"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
	Uid         string `json:"uid"`
	Username    string `json:"username"`
}

type GetAffiliateCustomOpenInfoV5SuccessResponse struct {
	RetCode int                           `json:"retCode"`
	RetMsg  string                        `json:"retMsg"`
	Result  AffiliateCustomOpenInfoResult `json:"result"`
}

type AffiliateCustomOpenInfoResult struct {
	Uid                  string      `json:"uid"`
	VipLevel             string      `json:"vipLevel"`
	TakerVol30Day        string      `json:"takerVol30Day"`
	MakerVol30Day        string      `json:"makerVol30Day"`
	TradeVol30Day        string      `json:"tradeVol30Day"`
	DepositAmount30Day   string      `json:"depositAmount30Day"`
	TakerVol365Day       string      `json:"takerVol365Day"`
	MakerVol365Day       string      `json:"makerVol365Day"`
	TradeVol365Day       string      `json:"tradeVol365Day"`
	DepositAmount365Day  string      `json:"depositAmount365Day"`
	TotalWalletBalance   string      `json:"totalWalletBalance"`
	DepositUpdateTime    string      `json:"depositUpdateTime"`
	VolUpdateTime        string      `json:"volUpdateTime"`
	KycLevel             int         `json:"KycLevel"`
	TradfiTradeVol30Day  string      `json:"tradfiTradeVol30Day"`
	TradfiTradeVol365Day string      `json:"tradfiTradeVol365Day"`
	Commissions30Day     interface{} `json:"commissions30Day"`
	Commissions365Day    interface{} `json:"commissions365Day"`
	PaySendAmount30Day   string      `json:"paySendAmount30Day"`
	PayFtt               string      `json:"payFtt"`
	CardFtt              string      `json:"cardFtt"`
}

type CreateSubAPIKeyRequest struct {
	Subuid      int         `json:"subuid"`
	ReadOnly    int         `json:"readOnly"`
	Ips         string      `json:"ips"`
	Permissions interface{} `json:"permissions"`
	Note        string      `json:"note"`
}

type CreateSubAPIKeyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CreateSubMemberRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	MemberType int    `json:"memberType"`
	Switch     int    `json:"switch"`
	IsUta      bool   `json:"isUta"`
	Note       string `json:"note"`
}

type CreateSubMemberResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type DeleteAPIKeyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type DeleteSubAPIKeyRequest struct {
	Subuid int    `json:"subuid"`
	Apikey string `json:"apikey"`
}

type DeleteSubAPIKeyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type DeleteSubMemberV5Request struct {
	Subuid int `json:"subuid"`
}

type DeleteSubMemberV5Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type FrozenSubMemberRequest struct {
	Subuid int `json:"subuid"`
	Frozen int `json:"frozen"`
}

type FrozenSubMemberResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetAffiliateCustomOpenInfoV5Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type ListSubAPIKeysV5Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryAPIKeyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryEscrowSubMembersV5Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QueryReferralsResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type ReferralRecord struct {
	Id         int `json:"id"`
	InviteeUid int `json:"inviteeUid"`
	Status     int `json:"status"`
	CreatedAt  int `json:"createdAt"`
	UpdatedAt  int `json:"updatedAt"`
}

type QuerySubMembersV5Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type QuerySubMembersResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type SignAgreementRequest struct {
	Category int  `json:"category"`
	Agree    bool `json:"agree"`
}

type SignAgreementResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type UpdateAPIKeyRequest struct {
	ReadOnly    int         `json:"readOnly"`
	Ips         string      `json:"ips"`
	Permissions interface{} `json:"permissions"`
}

type UpdateAPIKeyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type UpdateSubAPIKeyRequest struct {
	Subuid      int         `json:"subuid"`
	Apikey      string      `json:"apikey"`
	ReadOnly    int         `json:"readOnly"`
	Ips         string      `json:"ips"`
	Permissions interface{} `json:"permissions"`
	Note        string      `json:"note"`
}

type UpdateSubAPIKeyResponse struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type GetAffiliateSubListResult struct {
	List           []AffiliateSubItem `json:"list"`
	NextPageCursor string             `json:"nextPageCursor"`
}
