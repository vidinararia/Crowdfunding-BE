package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/vidinararia/crowd-funding-be/dataobject"
)

func NewBlock(userID int, fundingID int, amount int64) dataobject.FundingBlock {
	return dataobject.FundingBlock{
		UserID:    userID,
		FundingID: fundingID,
		Amount:    amount,
	}
}

func HashBlock(block dataobject.FundingBlock) string {
	block.Hash = ""
	block.PreviousNodeHash = ""

	jsonifiedBlock, _ := json.Marshal(block)
	md5Hash := md5.Sum(jsonifiedBlock)

	return hex.EncodeToString(md5Hash[:])
}
