package middleware

import (
	"errors"
	"go_rtb/internal/handler/response"
	"net/http"
)

const AdminKey = "8c20cfdf3d8e42e36f60aa26de116ac524f109de4954e3d45d31308c17f2351989f885b601da5bb7"

func AdminProtector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adminApiKey := r.Header.Get("Authorization")
		if adminApiKey == "" || adminApiKey != AdminKey {
			res := response.ErrorResponse(errors.New("Invalid access"), http.StatusForbidden)
			response.RenderJson(w, res)
			return
		}

		next.ServeHTTP(w, r)
	})
}
