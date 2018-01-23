package data_generator

import "go_rtb/internal/rtb"

func GeneraSimpleBidResult() *rtb.BidResult {
	bidResponse := GenerateSimpleBidResponse()
	bidResult := &rtb.BidResult{
		HighestBidResp:         bidResponse,
		HighestBid:             bidResponse.Seatbid[0].Bid[0],
		HighestApiResponseBody: "",
		HighestSeatBidPos:      0,
		HighestBidPos:          0,
		DSP:                    GenerateDSP(),
	}

	return bidResult
}
