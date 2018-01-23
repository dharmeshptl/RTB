package payload

import (
	"go_rtb/internal/model"
	"net/http"
)

type SSPUpdatePayload struct {
	*model.SSP
	OmitID interface{} `json:"id,omitempty"`
}

func (payload *SSPUpdatePayload) Bind(r *http.Request) error {
	return nil
}
