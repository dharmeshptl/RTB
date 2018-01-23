package request_checker

import (
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
)

type BannerTypeChecker struct {
}

func (checker *BannerTypeChecker) IsRequestValidForDSP(bidRequest *protocol_buffer.BidRequest, dsp *model.DSP) CheckerResult {
	if len(dsp.BannerTypes) == 0 {
		return NewCheckerResult(
			false,
			"DSP support no banner type",
			nil,
		)
	}

	for _, imp := range bidRequest.Imp {
		if imp.Banner != nil && dsp.HasBannerType(model.BannerTypeBanner) {
			return NewCheckerResult(
				true,
				"Bid request has banner type Banner and DSP support banner type Banner",
				nil,
			)
		} else if imp.Video != nil && dsp.HasBannerType(model.BannerTypeVideo) {
			return NewCheckerResult(
				true,
				"Bid request has banner type Video and DSP support banner type Video",
				nil,
			)
		} else if imp.Native != nil && dsp.HasBannerType(model.BannerTypeNative) {
			return NewCheckerResult(
				true,
				"Bid request has banner type Native and DSP support banner type Native",
				nil,
			)
		}
	}

	return NewCheckerResult(
		false,
		"Can not find any bidRequeset.Imp in bid request data. Invalid bid request",
		nil,
	)
}
