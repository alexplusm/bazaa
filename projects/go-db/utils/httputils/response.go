package httputils

import (
	"net/http"
)

type ResponseTemplate struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	ResponseTemplate
	Error serverError `json:"error"`
}

type SuccessResponse struct {
	ResponseTemplate
	Data interface{} `json:"data"`
}

type serverError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BuildSuccessResponse(data interface{}) interface{} {
	return SuccessResponse{ResponseTemplate{true}, data}
}

func BuildSuccessWithoutBodyResponse() interface{} {
	return ResponseTemplate{true}
}

func BuildErrorResponse(code int, message string) interface{} {
	serverErr := serverError{code, message}
	return ErrorResponse{ResponseTemplate{false}, serverErr}
}

func BuildBadRequestErrorResponse() interface{} {
	return BuildBadRequestErrorResponseWithMgs("")
}

func BuildInternalServerErrorResponse() interface{} {
	return BuildErrorResponse(http.StatusInternalServerError, "internal server error")
}

func BuildBadRequestErrorResponseWithMgs(message string) interface{} {
	msg := "bad request"
	if message != "" {
		msg += ": " + message
	}
	return BuildErrorResponse(http.StatusBadRequest, msg)
}
