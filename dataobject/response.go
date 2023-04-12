package dataobject

type ResponseBody struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type CrowdFundingTotalResponse struct {
	TotalRaised int64    `json:"total_raised"`
	Funders     []Funder `json:"funders"`
}

type Funder struct {
	UserID         int   `json:"user_id"`
	RaisedByFunder int64 `json:"raised_by_funder"`
}
