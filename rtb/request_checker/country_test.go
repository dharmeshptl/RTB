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

var countryChecker request_checker.CountryChecker

func TestCountryChecker_IsRequestValidForDSP(t *testing.T) {
	countryChecker = request_checker.CountryChecker{}

	bidRequestWithoutGeo := data_generator.GenerateSimpleBannerBidRequestWithoutGeo()

	dspWithoutCountryCodes := data_generator.GenerateDSP()
	dspWithoutCountryCodes.CountryCodes = []string{}

	dspWithUSCountryCode := data_generator.GenerateDSP()
	dspWithUSCountryCode.CountryCodes = []string{"USA", "VAT"}

	bidRequestWithGeoUS := data_generator.GenerateSimpleBannerBidRequest()

	bidRequestWithGeoFromVN := data_generator.GenerateSimpleBannerBidRequest()
	*bidRequestWithGeoFromVN.Device.Geo.Country = "VNM"

	dspWithCountryCodesVN := data_generator.GenerateDSP()
	dspWithCountryCodesVN.CountryCodes = []string{"VNM", "VAT"}

	cases := []struct {
		dsp            *model.DSP
		bidRequest     *protocol_buffer.BidRequest
		expectedResult bool
	}{
		{dspWithoutCountryCodes, bidRequestWithoutGeo, true},
		{dspWithUSCountryCode, bidRequestWithoutGeo, true},
		{dspWithUSCountryCode, bidRequestWithGeoUS, true},
		{dspWithUSCountryCode, bidRequestWithGeoFromVN, false},
		{dspWithCountryCodesVN, bidRequestWithoutGeo, false},
	}

	for _, c := range cases {
		result := countryChecker.IsRequestValidForDSP(c.bidRequest, c.dsp)
		helper.PanicOnError(result.GetError())
		assert.Equal(t, c.expectedResult, result.Success())
	}
}
