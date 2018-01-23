package filter

import (
	"fmt"
	"strconv"
)

type DSPLogFilter struct {
	DSPId    uint `schema:"dsp_id"`
	StatHour uint `schema:"stat_hour"`
}

func NewDSPLogFilter() *DSPLogFilter {
	return &DSPLogFilter{}
}

func (filter *DSPLogFilter) ToLogID() string {
	return fmt.Sprintf(
		"%s_%s",
		strconv.Itoa(int(filter.StatHour)),
		strconv.Itoa(int(filter.DSPId)),
	)
}
