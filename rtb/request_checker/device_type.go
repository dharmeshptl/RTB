package request_checker

import (
	"fmt"
	"github.com/avct/uasurfer"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
)

type DeviceTypeChecker struct {
}

func (checker *DeviceTypeChecker) IsRequestValidForDSP(bidRequest *protocol_buffer.BidRequest, dsp *model.DSP) CheckerResult {
	if len(dsp.DeviceTypes) == 0 {
		return NewCheckerResult(
			false,
			"DSP supports not device type",
			nil,
		)
	}

	//If device_type is not sent in bid_request -> parser from user browser
	if bidRequest.Device.Devicetype == nil {
		userAgent := uasurfer.Parse(*bidRequest.Device.Ua)
		if dsp.HasDeviceTypeMobile() && checker.isUserAgentMobileDevice(userAgent) {
			return NewCheckerResult(
				true,
				"Bid request has not Device.DeviceType, user agent is mobile browser, dsp support it",
				nil,
			)
		} else if dsp.HasDeviceTypeDesktop() && checker.isUserAgentMobileDevice(userAgent) == false {
			return NewCheckerResult(
				true,
				"Bid request has not Device.DeviceType, user agent is desktop browser, dsp support it",
				nil,
			)
		}
	} else {
		bidRequestDeviceType := *bidRequest.Device.Devicetype
		if dsp.HasDeviceTypeMobile() && checker.isDeviceTypeMobile(bidRequestDeviceType) {
			return NewCheckerResult(
				true,
				fmt.Sprintf(
					"Device type is mobile type: %d, dsp support it",
					int32(*bidRequest.Device.Devicetype),
				),
				nil,
			)
		} else if dsp.HasDeviceTypeDesktop() && checker.isDeviceTypeMobile(bidRequestDeviceType) == false {
			return NewCheckerResult(
				true,
				fmt.Sprintf(
					"Device type is desktop type: %d, dsp support it",
					int32(*bidRequest.Device.Devicetype),
				),
				nil,
			)
		}
	}

	return NewCheckerResult(
		false,
		"Unknown error when checking device type for bid request",
		nil,
	)
}

func (checker *DeviceTypeChecker) isUserAgentMobileDevice(userAgent *uasurfer.UserAgent) bool {
	if userAgent.DeviceType == uasurfer.DevicePhone || userAgent.DeviceType == uasurfer.DeviceTablet {
		return true
	}
	return false
}

func (checker *DeviceTypeChecker) isDeviceTypeMobile(deviceType protocol_buffer.DeviceType) bool {
	mobileDeviceTypes := []protocol_buffer.DeviceType{
		protocol_buffer.DeviceType_MOBILE,
		protocol_buffer.DeviceType_HIGHEND_PHONE,
		protocol_buffer.DeviceType_TABLET,
	}

	for _, ele := range mobileDeviceTypes {
		if ele == deviceType {
			return true
		}
	}

	return false
}
