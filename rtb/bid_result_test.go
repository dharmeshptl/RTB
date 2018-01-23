package rtb_test

import (
	"github.com/stretchr/testify/assert"
	"go_rtb/internal/rtb"
	"go_rtb/internal/test/data_generator"
	"testing"
)

func TestBidResult_ToSSPResponse(t *testing.T) {
	assertor := assert.New(t)
	bidResponse := data_generator.GenerateSimpleBidResponse()
	*bidResponse.Seatbid[0].Bid[0].Price = 0.0003

	bidResult := &rtb.BidResult{
		HighestBidResp:         bidResponse,
		HighestBid:             bidResponse.Seatbid[0].Bid[0],
		HighestApiResponseBody: "",
		HighestSeatBidPos:      0,
		HighestBidPos:          0,
		DSP:                    data_generator.GenerateDSP(),
	}

	returnBidResponse := bidResult.ToSSPResponse(7, "new_url_test")

	assertor.Equal(bidResponse.Id, returnBidResponse.Id)
	assertor.Equal(0.000279, *returnBidResponse.Seatbid[0].Bid[0].Price)
	assertor.Equal("new_url_test", *returnBidResponse.Seatbid[0].Bid[0].Nurl)
}
