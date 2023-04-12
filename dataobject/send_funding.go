package dataobject

type SendFundingRequest struct {
	UserID    int   `json:"user_id"`
	FundingID int   `json:"funding_id"`
	Amount    int64 `json:"amount"`
}
