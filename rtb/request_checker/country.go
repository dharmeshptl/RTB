package request_checker

import (
	"fmt"
	"go_rtb/internal/model"
	"go_rtb/internal/protocol_buffer"
)

const defaultCountryCode string = "USA"

type CountryChecker struct {
}

func (checker *CountryChecker) IsRequestValidForDSP(bidRequest *protocol_buffer.BidRequest, dsp *model.DSP) CheckerResult {
	if len(dsp.CountryCodes) == 0 {
		return NewCheckerResult(
			true,
			"DSP supports bid request from any country",
			nil,
		)
	}

	//No country code in request data, consider it comes from US
	if bidRequest.Device.Geo == nil || *bidRequest.Device.Geo.Country == "" {
		if dsp.IsValidForCountry(defaultCountryCode) {
			return NewCheckerResult(
				true,
				fmt.Sprintf(
					"BidRequest.Device.Geo is empty, default country is set to: %s, and DSP support: %s",
					defaultCountryCode,
					defaultCountryCode,
				),
				nil,
			)
		}
		return NewCheckerResult(
			false,
			fmt.Sprintf(
				"BidRequest.Device.Geo is empty, default country is set to: %s, and DSP does NOT support: %s",
				defaultCountryCode,
				defaultCountryCode,
			),
			nil,
		)
	}

	for _, country := range dsp.CountryCodes {
		if *bidRequest.Device.Geo.Country == country {
			return NewCheckerResult(
				true,
				fmt.Sprintf(
					"BidRequest country is :%s and DSP support it",
					*bidRequest.Device.Geo.Country,
				),
				nil,
			)
		}
	}

	return NewCheckerResult(
		false,
		"Unknown error when checking country for bid request",
		nil,
	)
}
