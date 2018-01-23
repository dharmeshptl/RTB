package data_generator

import (
	"github.com/icrowley/fake"
	"go_rtb/internal/model"
	"math/rand"
)

func GenerateSSP() *model.SSP {
	ssp := model.SSP{
		ID:           randomUInt(),
		Name:         fake.FullName(),
		ApiKey:       fake.Word(),
		Company:      fake.Company(),
		Region:       fake.Word(),
		MinFloor:     rand.Float64(),
		ProfitMargin: randomUInt(),
		Active:       true,
		DSPIds:       []uint{},
	}

	return &ssp
}
