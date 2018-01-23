package rtb

import "go_rtb/internal/protocol_buffer"

type RtbRequest struct {
	BidRequest *protocol_buffer.BidRequest
}
