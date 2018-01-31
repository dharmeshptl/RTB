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

var adsTypeChecker request_checker.AdsTypeChecker

func TestAdsTypeChecker_IsRequestValidForDSP(t *testing.T) {
	adsTypeChecker = request_checker.AdsTypeChecker{}

	webBidRequest := data_generator.GenerateSimpleBannerBidRequest()
	nativeBidRequest := data_generator.GenerateSimpleBannerBidRequest()
	*nativeBidRequest.Device.Ua = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre"

	dspWithNoAdsType := data_generator.GenerateDSP()
	dspWithNoAdsType.AdsTypes = []string{}

	dspWebAdsTypeDsp := data_generator.GenerateDSP()
	dspWebAdsTypeDsp.AdsTypes = []string{model.AdsTypeWeb}

	dspInAppAdsType := data_generator.GenerateDSP()
	dspInAppAdsType.AdsTypes = []string{model.AdsTypeInApp}

	cases := []struct {
		bidRequest     *protocol_buffer.BidRequest
		dsp            *model.DSP
		expectedResult bool
	}{
		{webBidRequest, dspWebAdsTypeDsp, true},
		{webBidRequest, dspInAppAdsType, false},

		{nativeBidRequest, dspWebAdsTypeDsp, false},
		{nativeBidRequest, dspInAppAdsType, true},

		{webBidRequest, dspWithNoAdsType, false},
		{nativeBidRequest, dspWithNoAdsType, false},
	}

	for _, c := range cases {
		result := adsTypeChecker.IsRequestValidForDSP(c.bidRequest, c.dsp)
		helper.PanicOnError(result.GetError())
		assert.Equal(t, c.expectedResult, result.Success())
	}
}
