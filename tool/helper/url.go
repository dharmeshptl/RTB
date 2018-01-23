package helper

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

var schemaDecoder = schema.NewDecoder()

// DecodeURLParam Parses url params to target
func DecodeURLParam(target interface{}, r *http.Request) error {
	return schemaDecoder.Decode(target, r.URL.Query())
}

// DecodePayload Decode payload to target
func DecodePayload(target interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(target)
}
