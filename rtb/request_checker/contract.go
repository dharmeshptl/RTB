package request_checker

import (
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
)

type DSPRequestChecker interface {
	IsRequestValidForDSP(bidRequest *protocol_buffer.BidRequest, dsp *model.DSP) CheckerResult
}
