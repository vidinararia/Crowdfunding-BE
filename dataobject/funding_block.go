package dataobject

type FundingBlock struct {
	UserID           int    `json:"user_id"`
	FundingID        int    `json:"funding_id"`
	Amount           int64  `json:"amount"`
	Timestamp        int64  `json:"timestamp"`
	Hash             string `json:"hash"`
	PreviousNodeHash string `json:"previous_node_hash"`
}

type FundingBlockchain []FundingBlock
