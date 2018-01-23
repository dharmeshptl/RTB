package model

type NUrl struct {
	//Generated token, also use as primary key for table
	Token string `as:"token"`
	//nurl we build and send to ssp
	DSPNurl  string  `as:"dsp_nurl"`
	SSPId    uint    `as:"ssp_id"`
	DSPId    uint    `as:"dsp_id"`
	BidPrice float64 `as:"bid_price"`
	//Store current ssp's profit margin cause, it can be changed
	ProfitMargin uint `as:"profit_margin"`
	//Make sure that ssp can only call win confirm once
	//Because we calculate money and other stuff in this request
	IsUsed bool `as:"is_used"`
}

func (model *NUrl) GetId() string {
	return model.Token
}
