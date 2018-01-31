package value_object_test

import (
	"github.com/stretchr/testify/assert"
	"go_rtb/internal/tool/helper"
	"go_rtb/internal/value_object"
	"testing"
)

func TestNewBannerFromJsonString(t *testing.T) {
	data := []string{
		`{"w":100,"h":150,"pos":2}`,
		`{"w":100,"h":200,"pos":1}`,
		`{"w":100,"h":200}`,
	}
	bannerList, err := value_object.NewBannerFromJsonString(data)
	helper.PanicOnError(err)

	assert.Equal(t, bannerList[0].GetHeight(), int32(150))
	assert.Equal(t, bannerList[1].GetWidth(), int32(100))
	assert.Equal(t, bannerList[0].GetPosition(), int32(2))
	assert.Equal(t, bannerList[2].GetPosition(), int32(0))
}
