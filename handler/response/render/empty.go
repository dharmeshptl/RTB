package render

import "net/http"

type EmptyResponseRender struct {
}

func NewEmptyResponse() *EmptyResponseRender {
	return &EmptyResponseRender{}
}

func (rd *EmptyResponseRender) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
