package model

const SSPLogImpressionField string = "impression"
const SSPLogSumRequestField string = "sum_request"
const SSPLogSumResponseField string = "sum_response"
const SSPLogSpendByHourField string = "spend_by_hour"

type SSPLog struct {
	//Compose of YearMonthDayHour_SSPId
	ID          string  `json:"id" as:"id"`
	SSPId       uint    `json:"ssp_id" as:"ssp_id"`
	StatHour    uint    `json:"stat_hour" as:"stat_hour"`
	BidQps      float64 `json:"bid_qps" as:"bid_qps"`
	Qps         float64 `json:"qps" as:"qps"`
	Impression  uint64  `json:"impression" as:"impression"`
	SumRequest  uint64  `json:"sum_request" as:"sum_request"`
	SumResponse uint64  `json:"sum_response" as:"sum_response"`
	SpendByHour float64 `json:"spend_by_hour" as:"spend_by_hour"`
}

func (model *SSPLog) GetID() string {
	return model.ID
}
