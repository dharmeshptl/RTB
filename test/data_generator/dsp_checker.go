package data_generator

import "go_rtb/internal/rtb/request_checker"

func GenerateCheckerList() []request_checker.DSPRequestChecker {
	checkerList := []request_checker.DSPRequestChecker{
		&request_checker.AdsSizeChecker{},
		&request_checker.AdsTypeChecker{},
		&request_checker.BannerTypeChecker{},
		&request_checker.CountryChecker{},
		&request_checker.DeviceTypeChecker{},
	}

	return checkerList
}
