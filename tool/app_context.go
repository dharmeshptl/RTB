package tool

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"go_rtb/internal/env"
	"go_rtb/internal/handler/response"
)

var schemaDecoder = schema.NewDecoder()

type AppContext struct {
	e *env.Env
	w http.ResponseWriter
	r *http.Request
}

func NewAppContext(e *env.Env, w http.ResponseWriter, r *http.Request) AppContext {
	return AppContext{e, w, r}
}

// DecodeURLParam Parses url params to target
func (c *AppContext) DecodeURLParam(target interface{}) error {
	return schemaDecoder.Decode(target, c.r.URL.Query())
}

// DecodePayload Decode payload to target
func (c *AppContext) DecodePayload(target interface{}) error {
	decoder := json.NewDecoder(c.r.Body)
	return decoder.Decode(target)
}

func (c *AppContext) GetRequest() *http.Request {
	return c.r
}

func (c *AppContext) GetResponseWriter() http.ResponseWriter {
	return c.w
}

func (c *AppContext) GetRequestEnv() *env.Env {
	return c.e
}

// Validate Shortcuts to validate a struct
func (c *AppContext) Validate(s interface{}) error {
	return c.e.GetValidator().Struct(s)
}

// PanicOnError
func (c *AppContext) PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// Func Handler function
type HandlerFunc func(ctx AppContext) response.ApiResponse
