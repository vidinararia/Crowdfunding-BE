package usecase

import (
	"math"

	"github.com/vidinararia/crowd-funding-be/dataobject"
)

func GetBlockchainWithConsensus(nodeIds ...string) dataobject.FundingBlockchain {
	var blockchains []dataobject.FundingBlockchain
	for _, v := range nodeIds {
		blockchains = append(blockchains, GetBlockchain(v))
	}

	max := 0
	for _, v := range blockchains {
		max = int(math.Max(float64(max), float64(len(v))))
	}

	consensusBlockchain := make(dataobject.FundingBlockchain, 0)

	for i := 0; i < max; i++ {
		majority := make(map[string]int)
		highest := 0
		var highestBlock dataobject.FundingBlock

		for _, bchain := range blockchains {
			if i >= len(bchain) {
				continue
			}

			currNodeHash := HashBlock(bchain[i])
			majority[currNodeHash]++

			if majority[currNodeHash] > highest {
				highest = majority[currNodeHash]
				highestBlock = bchain[i]
			}
		}

		consensusBlockchain = append(consensusBlockchain, highestBlock)
	}

	return consensusBlockchain
}
