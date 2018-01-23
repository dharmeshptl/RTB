package payload

import (
	"go_rtb/internal/model"
	"net/http"
)

type SSPCreatePayload struct {
	*model.SSP
}

func (payload *SSPCreatePayload) Bind(r *http.Request) error {
	return nil
}
