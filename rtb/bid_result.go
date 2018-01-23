package rtb

import (
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/tool/helper"
)

type BidResult struct {
	HighestBidResp         *protocol_buffer.BidResponse
	HighestBid             *protocol_buffer.BidResponse_SeatBid_Bid
	HighestApiResponseBody string
	//Because we have seatbid and bid as array, we need to store the highest one
	//And remove others
	HighestSeatBidPos int
	HighestBidPos     int
	DSP               *model.DSP
}

func (bidResult *BidResult) ToSSPResponse(profitMargin uint, newNurl string) protocol_buffer.BidResponse {
	bibResponse := bidResult.HighestBidResp
	highestSeatBid := bibResponse.Seatbid[bidResult.HighestSeatBidPos]
	highestBid := highestSeatBid.Bid[bidResult.HighestBidPos]

	price := *highestBid.Price
	priceCalculated := helper.CalculatePrice(price, profitMargin)
	highestBid.Price = &priceCalculated

	highestBid.Nurl = &newNurl
	bidArray := []*protocol_buffer.BidResponse_SeatBid_Bid{highestBid}
	highestSeatBid.Bid = bidArray
	seatBidArray := []*protocol_buffer.BidResponse_SeatBid{highestSeatBid}
	bibResponse.Seatbid = seatBidArray

	return *bibResponse
}
