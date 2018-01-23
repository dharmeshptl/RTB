package helper

import "math"

func CalculatePrice(price float64, profitMargin uint) float64 {
	return price * toFixed((float64(1)-(float64(profitMargin)/100)), 2)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
