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

var deviceTypeChecker request_checker.DeviceTypeChecker

func TestDeviceTypeChecker_IsRequestValidForDSP(t *testing.T) {
	deviceTypeChecker = request_checker.DeviceTypeChecker{}
	mobileBidRequestWithDeviceType := data_generator.GenerateMobileBidRequest()
	mobileBidRequestWithoutDeviceType := data_generator.GenerateMobileBidRequestWithoutDeviceType()

	desktopBidRequestWithoutDeviceType := data_generator.GenerateSimpleBannerBidRequest()

	dspDeviceTypeMobile := data_generator.GenerateDSP()
	dspDeviceTypeMobile.DeviceTypes = []string{model.DeviceTypeMobile}

	dspDeviceTypeDesktop := data_generator.GenerateDSP()
	dspDeviceTypeDesktop.DeviceTypes = []string{model.DeviceTypeDesktop}

	dspHasAllDeviceType := data_generator.GenerateDSP()
	dspHasAllDeviceType.DeviceTypes = []string{
		model.DeviceTypeDesktop,
		model.DeviceTypeMobile,
	}

	dspHasNoDeviceType := data_generator.GenerateDSP()
	dspHasNoDeviceType.DeviceTypes = []string{}

	desktopBidRequestWithDeviceType := data_generator.GenerateSimpleBannerBidRequestWithoutGeo()

	cases := []struct {
		bidRequest     *protocol_buffer.BidRequest
		dsp            *model.DSP
		expectedResult bool
	}{
		{mobileBidRequestWithDeviceType, dspDeviceTypeMobile, true},
		{mobileBidRequestWithDeviceType, dspHasAllDeviceType, true},
		{mobileBidRequestWithDeviceType, dspDeviceTypeDesktop, false},

		{mobileBidRequestWithoutDeviceType, dspDeviceTypeMobile, true},
		{mobileBidRequestWithoutDeviceType, dspHasAllDeviceType, true},
		{mobileBidRequestWithoutDeviceType, dspDeviceTypeDesktop, false},

		{desktopBidRequestWithoutDeviceType, dspDeviceTypeDesktop, true},
		{desktopBidRequestWithoutDeviceType, dspHasAllDeviceType, true},
		{desktopBidRequestWithoutDeviceType, dspDeviceTypeMobile, false},

		{desktopBidRequestWithDeviceType, dspDeviceTypeDesktop, true},
		{desktopBidRequestWithDeviceType, dspHasAllDeviceType, true},
		{desktopBidRequestWithDeviceType, dspDeviceTypeMobile, false},

		{desktopBidRequestWithDeviceType, dspHasNoDeviceType, false},
		{desktopBidRequestWithoutDeviceType, dspHasNoDeviceType, false},
	}

	for _, c := range cases {
		result := deviceTypeChecker.IsRequestValidForDSP(c.bidRequest, c.dsp)
		helper.PanicOnError(result.GetError())
		assert.Equal(t, c.expectedResult, result.Success())
	}
}
