package filter

import (
	"fmt"
	"strconv"
)

type SSPLogFilter struct {
	SSPId    uint `schema:"ssp_id"`
	StatHour uint `schema:"stat_hour"`
}

func NewSSPLogFilter() *SSPLogFilter {
	return &SSPLogFilter{}
}

func (filter *SSPLogFilter) ToLogID() string {
	return fmt.Sprintf(
		"%s_%s",
		strconv.Itoa(int(filter.StatHour)),
		strconv.Itoa(int(filter.SSPId)),
	)
}
