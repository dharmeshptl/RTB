package model

import "go_rtb/internal/tool/helper"

const BannerTypeVideo string = "video"
const BannerTypeBanner string = "banner"
const BannerTypeNative string = "native"

const DeviceTypeMobile string = "mobile"
const DeviceTypeDesktop string = "desktop"

const AdsTypeWeb string = "web"
const AdsTypeInApp string = "in_app"

type DSP struct {
	ID           uint     `json:"id" as:"dsp_id" validate:"required"`
	Name         string   `json:"name" as:"name" validate:"required"`
	QPSLimit     float64  `json:"qps_limit" as:"qps_limit" validate:"required"`
	MinFloor     float32  `json:"min_floor" as:"min_floor" validate:"required"`
	Spend        float64  `json:"spend" as:"spend"`
	EndpointURL  string   `json:"endpoint_url" as:"endpoint_url" validate:"required"`
	Company      string   `json:"company" as:"company" validate:"required"`
	Region       string   `json:"region" as:"region" validate:"required"`
	BannerTypes  []string `json:"banner_types" as:"banner_types" validate:"string_in_array=banner;video;native"`
	CountryCodes []string `json:"country_codes" as:"country_codes" validate:"required"`
	DeviceTypes  []string `json:"device_types" as:"device_types" validate:"string_in_array=mobile;desktop"`
	AdsTypes     []string `json:"ads_types" as:"ads_types" validate:"string_in_array=web;in_app"`
	AdsSizes     []string `json:"ads_sizes" as:"ads_sizes"`
}

func (model *DSP) GetID() uint {
	return model.ID
}

func (dsp *DSP) HasBannerType(bannerToCheck string) bool {
	for _, bannerType := range dsp.BannerTypes {
		if bannerType == bannerToCheck {
			return true
		}
	}
	return false
}

func (dsp *DSP) HasDeviceTypeMobile() bool {
	return helper.DoesStringArrayContain(DeviceTypeMobile, dsp.DeviceTypes)
}

func (dsp *DSP) HasDeviceTypeDesktop() bool {
	return helper.DoesStringArrayContain(DeviceTypeDesktop, dsp.DeviceTypes)
}

func (dsp *DSP) IsValidForCountry(countryCode string) bool {
	return helper.DoesStringArrayContain(countryCode, dsp.CountryCodes)
}
