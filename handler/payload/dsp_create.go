package payload

import (
	"go_rtb/internal/model"
	"net/http"
)

type DSPCreatePayload struct {
	*model.DSP
}

func (payload *DSPCreatePayload) Bind(r *http.Request) error {
	return nil
}
