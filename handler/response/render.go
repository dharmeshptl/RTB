package response

import (
	"encoding/json"
	chiRender "github.com/go-chi/render"
	"go_rtb/internal/handler/response/render"
	"net/http"
)

func RenderJsonString(w http.ResponseWriter, res string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func RenderEmpty(w http.ResponseWriter, r *http.Request) {
	chiRender.Status(r, http.StatusNoContent)
	chiRender.Render(w, r, render.NewEmptyResponse())
}

func RenderJson(w http.ResponseWriter, res ApiResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	for key, value := range res.Headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(res.Code)
	if res.Data == nil || res.Code == http.StatusNoContent {
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(res.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
