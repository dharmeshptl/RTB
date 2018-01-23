package data_generator

import (
	"go_rtb/internal/model"
	"go_rtb/internal/tool/helper"
	"math/rand"
)

func GenerateStat(ssp *model.SSP, dsp *model.DSP) *model.Stat {
	primaryKey := helper.GetLogId(ssp.ID, dsp.ID)
	stat := &model.Stat{
		ID:          primaryKey,
		SSPId:       ssp.ID,
		DSPId:       dsp.ID,
		StatHour:    helper.GetCurrentStatHour(),
		Impression:  rand.Uint64(),
		SpendByHour: rand.Float64(),
		EarnByHour:  rand.Float64(),
	}

	return stat
}
