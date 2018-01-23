package model

const StatImpressionField string = "impression"
const StatSpendByHourField string = "spend_by_hour"
const StatEarnByHourField string = "earn_by_hour"

type Stat struct {
	//Compose of YearMonthDayHour_SSPId_DSPId
	ID       string `json:"id" as:"id"`
	SSPId    uint   `json:"ssp_id" as:"ssp_id"`
	DSPId    uint   `json:"dsp_id" as:"dsp_id"`
	StatHour uint   `json:"stat_hour" as:"stat_hour"`
	//Revenue = earning - spend
	//Revenue    uint64 `json:"revenue" as:"revenue"`
	Impression  uint64  `json:"impression" as:"impression"`
	SpendByHour float64 `json:"spend_by_hour" as:"spend_by_hour"`
	EarnByHour  float64 `json:"earn_by_hour" as:"earn_by_hour"`
}

func (model *Stat) GetID() string {
	return model.ID
}
