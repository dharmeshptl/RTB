package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePrice(t *testing.T) {
	assertor := assert.New(t)
	cases := []struct {
		price        float64
		profitMargin uint
		expected     float64
	}{
		{0.00011, 7, 0.0001023},
	}

	for _, c := range cases {
		result := CalculatePrice(c.price, c.profitMargin)
		assertor.Equal(c.expected, result)
	}
}
