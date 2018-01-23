package data_generator

import (
	"github.com/icrowley/fake"
	"go_rtb/internal/model"
)

func GenerateNUrl(ssp *model.SSP, dsp *model.DSP) *model.NUrl {
	nurl := model.NUrl{
		Token:        randomUUID(),
		DSPNurl:      fake.Word(),
		IsUsed:       false,
		SSPId:        ssp.ID,
		DSPId:        dsp.ID,
		BidPrice:     0.0000123,
		ProfitMargin: 7,
	}

	return &nurl
}
