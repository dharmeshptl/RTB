package payload

import (
	"go_rtb/internal/model"
	"net/http"
)

type DSPUpdatePayload struct {
	*model.DSP
	OmitID interface{} `json:"id,omitempty"`
}

func (payload *DSPUpdatePayload) Bind(r *http.Request) error {
	return nil
}
