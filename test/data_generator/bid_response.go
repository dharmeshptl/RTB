package data_generator

import (
	"github.com/golang/protobuf/jsonpb"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/tool/helper"
)

func GenerateSimpleBidResponse() *protocol_buffer.BidResponse {
	bidResponse := protocol_buffer.BidResponse{}
	jsonStr := `{
  "id": "1234567890",
  "bidid": "abc1123",
  "cur": "USD",
  "seatbid": [
    {
      "seat": "512",
      "bid": [
        {
          "id": "1",
          "impid": "102",
          "price": 9.43,
          "nurl": "http://adserver.com/winnotice?impid=102",
          "iurl": "http://adserver.com/pathtosampleimage",
          "adomain": [
            "advertiserdomain.com"
          ],
          "cid": "campaign111",
          "crid": "creative112",
          "attr": [
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            12
          ]
        }
      ]
    }
  ]
}`

	if err := jsonpb.UnmarshalString(jsonStr, &bidResponse); err != nil {
		helper.PanicOnError(err)
	}

	return &bidResponse
}

func GenerateDirectDealBidResponse() *protocol_buffer.BidResponse {
	bidResponse := protocol_buffer.BidResponse{}
	jsonStr := `{
  "id": "1234567890",
  "bidid": "abc1123",
  "cur": "USD",
  "seatbid": [
    {
      "seat": "512",
      "bid": [
        {
          "id": "1",
          "impid": "102",
          "price": 5,
          "dealid": "ABC-1234-6789",
          "nurl": "http: //adserver.com/winnotice?impid=102",
          "adomain": [
            "advertiserdomain.com"
          ],
          "iurl": "http: //adserver.com/pathtosampleimage",
          "cid": "campaign111",
          "crid": "creative112",
          "adid": "314",
          "attr": [
            1,
            2,
            3,
            4
          ]
        }
      ]
    }
  ]
}`

	if err := jsonpb.UnmarshalString(jsonStr, &bidResponse); err != nil {
		helper.PanicOnError(err)
	}

	return &bidResponse
}

func GenerateNativeAdsBidResponse() *protocol_buffer.BidResponse {
	bidResponse := protocol_buffer.BidResponse{}
	jsonStr := `{
  "id": "123",
  "seatbid": [
    {
      "bid": [
        {
          "id": "12345",
          "impid": "2",
          "price": 3,
          "nurl": "http://example.com/winnoticeurl",
          "adm": "...Native Spec response as an encoded string..."
        }
      ]
    }
  ]
}`

	if err := jsonpb.UnmarshalString(jsonStr, &bidResponse); err != nil {
		helper.PanicOnError(err)
	}

	return &bidResponse
}
