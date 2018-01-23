package filter

import (
	"fmt"
	"strconv"
)

type StatFilter struct {
	DSPId    uint `schema:"dsp_id"`
	SSPId    uint `schema:"ssp_id"`
	StatHour uint `schema:"stat_hour"`
}

func NewStatFilter() *StatFilter {
	return &StatFilter{}
}

func (filter *StatFilter) ToLogID() string {
	return fmt.Sprintf(
		"%s_%s_%s",
		strconv.Itoa(int(filter.StatHour)),
		strconv.Itoa(int(filter.SSPId)),
		strconv.Itoa(int(filter.DSPId)),
	)
}
