package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vidinararia/crowd-funding-be/dataobject"
)

// returns previous node hash
func AddBlockToBlockchain(nodeUid string, block dataobject.FundingBlock) string {
	blockchain := GetBlockchain(nodeUid)
	prevNodeHash := getPrevNodeHash(blockchain)

	block.PreviousNodeHash = prevNodeHash
	blockchain = append(blockchain, block)

	jsonifiedBlockchain, _ := json.Marshal(blockchain)

	err := os.WriteFile(fmt.Sprintf("./%s.json", nodeUid), jsonifiedBlockchain, 0677)
	if err != nil {
		log.Println(err)
		return ""
	}

	return prevNodeHash
}

func GetBlockchain(nodeUid string) dataobject.FundingBlockchain {
	blockchain := make([]dataobject.FundingBlock, 0)

	b, err := os.ReadFile(fmt.Sprintf("./%s.json", nodeUid))
	if err != nil {
		return blockchain
	}

	json.Unmarshal(b, &blockchain)
	return blockchain
}

func isFirstBlock(b dataobject.FundingBlockchain) bool {
	return len(b) == 0
}

func getPrevNodeHash(b dataobject.FundingBlockchain) string {
	if isFirstBlock(b) {
		return ""
	}

	return b[len(b)-1].Hash
}
