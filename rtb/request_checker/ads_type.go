package request_checker

import (
	"github.com/avct/uasurfer"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
)

type AdsTypeChecker struct {
}

func (checker *AdsTypeChecker) IsRequestValidForDSP(bidRequest *protocol_buffer.BidRequest, dsp *model.DSP) CheckerResult {
	if len(dsp.AdsTypes) == 0 {
		return NewCheckerResult(
			false,
			"DSP support no ads type",
			nil,
		)
	}

	userAgent := uasurfer.Parse(*bidRequest.Device.Ua)
	for _, adsType := range dsp.AdsTypes {
		if userAgent.Browser.Name != uasurfer.BrowserUnknown && adsType == model.AdsTypeWeb {
			return NewCheckerResult(
				true,
				"User agent is web browser and DSP support Web ads type",
				nil,
			)
		} else if userAgent.Browser.Name == uasurfer.BrowserUnknown && adsType == model.AdsTypeInApp {
			return NewCheckerResult(
				true,
				"User agent is not web brower (request from app) and DSP support InApp ads type",
				nil,
			)
		}
	}

	return NewCheckerResult(
		false,
		"Can not detect user agent browser from bid request",
		nil,
	)
}
