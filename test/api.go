package test

import (
	"encoding/json"
	"go_rtb/internal/router/middleware"
	"io"
	"net/http"
	"net/http/httptest"
)

func NewApiTest(h http.Handler) *ApiTest {
	return &ApiTest{h}
}

type ApiTest struct {
	h http.Handler
}

func (t *ApiTest) RequestTest(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	t.h.ServeHTTP(w, req)
	return w
}

func (t *ApiTest) buildRequest(method string, path string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Authorization", middleware.AdminKey)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func (t *ApiTest) Get(path string) *httptest.ResponseRecorder {
	return t.RequestTest(t.buildRequest(http.MethodGet, path, nil))
}

func (t *ApiTest) Post(path string, body io.Reader) *httptest.ResponseRecorder {
	return t.RequestTest(t.buildRequest(http.MethodPost, path, body))
}

func (t *ApiTest) Put(path string, body io.Reader) *httptest.ResponseRecorder {
	return t.RequestTest(t.buildRequest(http.MethodPut, path, body))
}

func (t *ApiTest) Patch(path string, body io.Reader) *httptest.ResponseRecorder {
	return t.RequestTest(t.buildRequest(http.MethodPatch, path, body))
}

func (t *ApiTest) Delete(path string) *httptest.ResponseRecorder {
	return t.RequestTest(t.buildRequest(http.MethodDelete, path, nil))
}

func ParseResponse(r io.Reader, v interface{}) error {
	//data := struct {
	//	Data interface{}
	//}{
	//	Data: v,
	//}

	decoder := json.NewDecoder(r)

	return decoder.Decode(&v)
}
