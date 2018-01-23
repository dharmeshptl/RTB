package model

const DSPLogSumRequestField string = "sum_request"
const DSPLogSumResponseField string = "sum_response"
const DSPLogSpendByHourField string = "spend_by_hour"

type DSPLog struct {
	//Compose of YearMonthDayHour_DSPId
	ID          string  `json:"id" as:"id"`
	DSPId       uint    `json:"dsp_id" as:"dsp_id"`
	StatHour    uint    `json:"stat_hour" as:"stat_hour"`
	BidQps      float64 `json:"bid_qps" as:"bid_qps"`
	Qps         float64 `json:"qps" as:"qps"`
	SumRequest  uint64  `json:"sum_request" as:"sum_request"`
	SumResponse uint64  `json:"sum_response" as:"sum_response"`
	SpendByHour float64 `json:"spend_by_hour" as:"spend_by_hour"`
}

func (model *DSPLog) GetID() string {
	return model.ID
}
