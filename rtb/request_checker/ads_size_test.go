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

var adsSizeChecker request_checker.AdsSizeChecker

func TestAdsSizeChecker_IsRequestValidForDSP(t *testing.T) {
	adsSizeChecker = request_checker.AdsSizeChecker{}

	dspWithNoAdsSize := data_generator.GenerateDSP()
	dspWithNoAdsSize.AdsSizes = []string{}
	dspWithNoAdsSize.BannerTypes = []string{
		model.BannerTypeBanner,
		model.BannerTypeVideo,
	}

	dspSupportBannerTypeBanner := data_generator.GenerateDSP()
	dspSupportBannerTypeBanner.BannerTypes = []string{
		model.BannerTypeBanner,
		model.BannerTypeVideo,
	}
	dspSupportBannerTypeBanner.AdsSizes = []string{
		`{"w":100,"h":150,"pos":2}`,
		`{"w":100,"h":200,"pos":1}`,
	}

	dspSupportBannerTypeVideo := data_generator.GenerateDSP()
	dspSupportBannerTypeVideo.BannerTypes = []string{
		model.BannerTypeBanner,
		model.BannerTypeVideo,
	}
	dspSupportBannerTypeVideo.AdsSizes = []string{
		`{"w":100,"h":225,"pos":12}`,
		`{"w":100,"h":200,"pos":1}`,
	}

	dspSupportBannerTypeNative := data_generator.GenerateDSP()
	dspSupportBannerTypeNative.BannerTypes = []string{
		model.BannerTypeBanner,
		model.BannerTypeVideo,
		model.BannerTypeNative,
	}

	simpleBannerBidRequest := data_generator.GenerateSimpleBannerBidRequest()
	*simpleBannerBidRequest.Imp[0].Banner.W = 100
	*simpleBannerBidRequest.Imp[0].Banner.H = 200
	*simpleBannerBidRequest.Imp[0].Banner.Pos = 1

	videoBidRequest := data_generator.GenerateVideoBidRequest()
	*videoBidRequest.Imp[1].Video.W = 100
	*videoBidRequest.Imp[1].Video.H = 225
	*videoBidRequest.Imp[1].Video.Pos = 12

	nativeBidRequest := data_generator.GenerateNativeAdsBidRequest()

	cases := []struct {
		bidRequest     *protocol_buffer.BidRequest
		dsp            *model.DSP
		expectedResult bool
	}{
		{simpleBannerBidRequest, dspWithNoAdsSize, true},
		{videoBidRequest, dspWithNoAdsSize, true},
		{nativeBidRequest, dspWithNoAdsSize, true},

		{simpleBannerBidRequest, dspSupportBannerTypeBanner, true},
		{videoBidRequest, dspSupportBannerTypeVideo, true},
		{nativeBidRequest, dspSupportBannerTypeNative, true},
	}

	for _, c := range cases {
		result := adsSizeChecker.IsRequestValidForDSP(c.bidRequest, c.dsp)
		helper.PanicOnError(result.GetError())
		assert.Equal(t, c.expectedResult, result.Success())
	}
}
