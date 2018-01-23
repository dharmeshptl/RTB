package helper_test

import (
	"github.com/stretchr/testify/assert"
	"go_rtb/internal/tool/helper"
	"testing"
)

func TestDoesStringArrayContain(t *testing.T) {
	cases := []struct {
		needle         string
		haystack       []string
		expectedResult bool
	}{
		{"USA", []string{"USA", "VNM"}, true},
		{"USA", []string{"VAT", "VNM"}, false},
	}

	for _, c := range cases {
		result := helper.DoesStringArrayContain(c.needle, c.haystack)
		assert.Equal(t, c.expectedResult, result)
	}
}
