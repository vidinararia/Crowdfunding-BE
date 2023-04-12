package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vidinararia/crowd-funding-be/dataobject"
	"github.com/vidinararia/crowd-funding-be/usecase"
	"github.com/vidinararia/crowd-funding-be/utils"
)

var blockchainNodes = []string{
	"node-1",
	"node-2",
	"node-3",
	"node-4",
	"node-5",
	"node-6",
}

func InitRouter(ec *echo.Echo) {
	ec.Use(middleware.CORS())

	ec.POST("/send-fund", func(c echo.Context) error {
		var sendFund dataobject.SendFundingRequest
		if err := c.Bind(&sendFund); err != nil {
			return c.JSON(400, dataobject.ResponseBody{
				Code: 1,
				Data: "invalid request body being sent",
			})
		}

		block := usecase.NewBlock(sendFund.UserID, sendFund.FundingID, sendFund.Amount)
		hash := usecase.HashBlock(block)
		block.Hash = hash

		prevHash := ""
		for _, v := range blockchainNodes {
			prevHash = usecase.AddBlockToBlockchain(v, block)
		}
		block.PreviousNodeHash = prevHash

		return c.JSON(200, dataobject.ResponseBody{
			Code: 0,
			Data: block,
		})
	})

	ec.GET("/all-chain", func(c echo.Context) error {
		blockchain := usecase.GetBlockchainWithConsensus(blockchainNodes...)
		return c.JSON(200, dataobject.ResponseBody{
			Code: 0,
			Data: blockchain,
		})
	})

	ec.GET("/crowd-funding/:id/total-fund", func(c echo.Context) error {
		crowdFundingID := utils.HttpParamsInt(c, "id")
		blockchain := usecase.GetBlockchainWithConsensus(blockchainNodes...)
		totalFund := int64(0)
		funderToAmountMap := make(map[int]int64)
		funders := make([]dataobject.Funder, 0)

		for _, v := range blockchain {
			if v.FundingID == crowdFundingID {
				totalFund += v.Amount
				funderToAmountMap[v.UserID] += v.Amount
			}
		}

		for i, v := range funderToAmountMap {
			funders = append(funders, dataobject.Funder{
				UserID:         i,
				RaisedByFunder: v,
			})
		}

		return c.JSON(200, dataobject.ResponseBody{
			Code: 0,
			Data: dataobject.CrowdFundingTotalResponse{
				TotalRaised: totalFund,
				Funders:     funders,
			},
		})
	})

	ec.GET("/crowd-funding/:id/transactions", func(c echo.Context) error {
		crowdFundingID := utils.HttpParamsInt(c, "id")
		blockchain := usecase.GetBlockchainWithConsensus(blockchainNodes...)
		transactions := make(dataobject.FundingBlockchain, 0)

		for _, v := range blockchain {
			if v.FundingID == crowdFundingID {
				transactions = append(transactions, v)
			}
		}

		return c.JSON(200, dataobject.ResponseBody{
			Code: 0,
			Data: transactions,
		})
	})
}
