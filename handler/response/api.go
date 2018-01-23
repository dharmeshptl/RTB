package response

import (
	"errors"
	"go_rtb/internal/tool/helper"
	"go_rtb/internal/tool/validation"
	"net/http"
)

type ApiResponse struct {
	Code    int               `json:"-"`
	Data    interface{}       `json:"data"`
	Headers map[string]string `json:"-"`
}

// swagger:model ErrorResponse
type Error struct {
	Message    string              `json:"message,omitempty"`
	Code       int                 `json:"code,omitempty"`
	StatusCode int                 `json:"status_code,omitempty"`
	Errors     map[string][]string `json:"errors,omitempty"`
}

// Ok 200-OK with data in body
func Ok(data interface{}) ApiResponse {
	return ApiResponse{http.StatusOK, data, nil}
}

// Accepted 202-Accepted with data in body
func Accepted(data interface{}) ApiResponse {
	return ApiResponse{http.StatusAccepted, data, nil}
}

// PartialContent 206-Partial Content with data in body
func PartialContent(data interface{}) ApiResponse {
	return ApiResponse{http.StatusPartialContent, data, nil}
}

func NoContent() ApiResponse {
	return ApiResponse{http.StatusNoContent, nil, nil}
}

// BadRequest 400-Bad Request
func BadRequest(err error) ApiResponse {
	return ErrorResponse(err, http.StatusBadRequest)
}

// ValidationError 422-Unprocessable Enttity
func ValidationError(err error) ApiResponse {
	errMessages := validation.ParseValidationErr(err)
	var response ApiResponse
	response.Code = http.StatusUnprocessableEntity
	response.Data = Error{
		Message:    http.StatusText(http.StatusUnprocessableEntity),
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     errMessages,
	}
	return response
}

// ErrorResponse Custom error response
func ErrorResponse(err error, status int) ApiResponse {
	var response ApiResponse
	response.Code = status
	response.Data = Error{
		Message:    err.Error(),
		StatusCode: status,
	}
	return response
}

func RecordNotFoundError(err error, message string) ApiResponse {
	if helper.IsRecordNotFoundError(err) {
		return ErrorResponse(errors.New(message), http.StatusNotFound)
	}

	return ErrorResponse(err, http.StatusInternalServerError)
}
