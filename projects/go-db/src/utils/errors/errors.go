package errors

// TODO: tests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BaseResponse base response
// todo: move into another package?
type BaseResponse struct {
	Success bool `json:"success"`
}

// ErrorResponse error response
type ErrorResponse struct {
	BaseResponse
	Error serverError `json:"error"`
}

type serverError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetErrorResponse get error response
// todo: make packages with responses? | todo: rename?
func getErrorResponse(code int, msg string) ErrorResponse {
	serverErr := serverError{code, msg}
	return ErrorResponse{BaseResponse{false}, serverErr}
}

// GetErrorResponseJSONStr get error response in JSON string
func GetErrorResponseJSONStr(code int, msg string) string {
	r := getErrorResponse(code, msg)
	value, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return "" // TODO!
	}

	return string(value)
}

// GetBadRequestErrorResponseJSONStr get bad request error in JSON string
func GetBadRequestErrorResponseJSONStr() string {
	return GetErrorResponseJSONStr(http.StatusBadRequest, "bad request")
}
