package data_generator

import (
	"github.com/icrowley/fake"
	"go_rtb/internal/model"
	"math/rand"
)

var (
	BannerTypes = []string{
		"banner",
		"video",
		"native",
	}
	CountryCodes = []string{
		"VNM",
		"GBR",
		"USA",
	}
	DeviceTypes = []string{
		"mobile",
		"desktop",
	}
	AdsTypes = []string{
		"web",
		"in_app",
	}
	AdsSizes = []string{
		`{"w":110,"h":123,"pos":1}`,
		`{"w":210,"h":123,"pos":13}`,
		`{"w":300,"h":250,"pos":0}`,
	}
)

func GenerateDSP() *model.DSP {
	dsp := model.DSP{
		ID:           randomUInt(),
		Name:         fake.FemaleFullName(),
		QPSLimit:     rand.Float64(),
		MinFloor:     rand.Float32(),
		Spend:        rand.Float64(),
		EndpointURL:  fake.Word(),
		Company:      fake.Company(),
		Region:       fake.Word(),
		BannerTypes:  randomStrArrFromList(BannerTypes),
		CountryCodes: randomStrArrFromList(CountryCodes),
		DeviceTypes:  randomStrArrFromList(DeviceTypes),
		AdsTypes:     randomStrArrFromList(AdsTypes),
		AdsSizes:     randomStrArrFromList(AdsSizes),
	}

	return &dsp
}
