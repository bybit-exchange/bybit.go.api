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
