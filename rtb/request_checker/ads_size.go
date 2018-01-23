package request_checker

import (
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/value_object"
)

type AdsSizeChecker struct {
}

func (checker *AdsSizeChecker) IsRequestValidForDSP(bidRequest *protocol_buffer.BidRequest, dsp *model.DSP) CheckerResult {
	//Native ads does not have size
	if len(dsp.AdsSizes) == 0 && !dsp.HasBannerType(model.BannerTypeNative) {
		return NewCheckerResult(true, "DSP accept all ads size and does not have banner type native", nil)
	}

	bannerList, err := value_object.NewBannerFromJsonString(dsp.AdsSizes)
	if err != nil {
		return NewCheckerResult(false, "Can not decode ads size info from dsp", err)
	}
	for _, imp := range bidRequest.Imp {
		if imp.Banner != nil && dsp.HasBannerType(model.BannerTypeBanner) {
			for _, banner := range bannerList {
				if banner.EqualWithBidBanner(imp.Banner) {
					return NewCheckerResult(
						true,
						"DSP has banner type Banner, and size is equal with bid request Banner",
						nil,
					)
				}
			}
		} else if imp.Video != nil && dsp.HasBannerType(model.BannerTypeVideo) {
			for _, banner := range bannerList {
				if banner.EqualWithBidBannerVideo(imp.Video) {
					return NewCheckerResult(
						true,
						"DSP has banner type Video, and size is equal with bid request Video",
						nil,
					)
				}
			}
		} else if imp.Native != nil && dsp.HasBannerType(model.BannerTypeNative) {
			return NewCheckerResult(
				true,
				"DSP has banner type native",
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
