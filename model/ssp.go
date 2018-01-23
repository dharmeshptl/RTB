package model

type SSP struct {
	ID           uint    `json:"id" as:"ssp_id" validate:"required"`
	Name         string  `json:"name" as:"name" validate:"required"`
	ApiKey       string  `json:"api_key" as:"api_key" validate:"required"`
	Company      string  `json:"company" as:"company" validate:"required"`
	Region       string  `json:"region" as:"region" validate:"required"`
	MinFloor     float64 `json:"min_floor" as:"min_floor" validate:"required"`
	ProfitMargin uint    `json:"profit_margin" as:"profit_margin" validate:"required"`
	Active       bool    `json:"active" as:"active" validate:"required"`
	DSPIds       []uint  `json:"dsp_ids" as:"dsp_ids" validate:"required"`
}

func (model *SSP) GetId() uint {
	return model.ID
}
