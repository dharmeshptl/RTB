package request_checker_test

import (
	"github.com/stretchr/testify/assert"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
	"go_rtb/internal/rtb/request_checker"
	"go_rtb/internal/test/data_generator"
	"go_rtb/internal/tool/helper"
	"testing"
)

var bannerTypeChecker request_checker.BannerTypeChecker

func TestBannerTypeChecker_IsRequestValidForDSP(t *testing.T) {
	bannerTypeChecker = request_checker.BannerTypeChecker{}

	dspSupportNoBannerType := data_generator.GenerateDSP()
	dspSupportNoBannerType.BannerTypes = []string{}

	dspSupportAllBannerType := data_generator.GenerateDSP()
	dspSupportAllBannerType.BannerTypes = []string{
		model.BannerTypeBanner,
		model.BannerTypeVideo,
		model.BannerTypeNative,
	}

	simpleBannerBidRequest := data_generator.GenerateSimpleBannerBidRequest()
	videoBidRequest := data_generator.GenerateVideoBidRequest()
	nativeBidRequest := data_generator.GenerateNativeAdsBidRequest()

	cases := []struct {
		bidRequest     *protocol_buffer.BidRequest
		dsp            *model.DSP
		expectedResult bool
	}{
		{simpleBannerBidRequest, dspSupportNoBannerType, false},
		{videoBidRequest, dspSupportNoBannerType, false},
		{nativeBidRequest, dspSupportNoBannerType, false},

		{simpleBannerBidRequest, dspSupportAllBannerType, true},
		{videoBidRequest, dspSupportAllBannerType, true},
		{nativeBidRequest, dspSupportAllBannerType, true},
	}

	for _, c := range cases {
		result := bannerTypeChecker.IsRequestValidForDSP(c.bidRequest, c.dsp)
		helper.PanicOnError(result.GetError())
		assert.Equal(t, c.expectedResult, result.Success())
	}
}
