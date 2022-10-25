package message

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Code     int        `json:"code"`
	Message  string     `json:"message"`
	Type     string     `json:"type"`
	StarTime *time.Time `json:"start_time,omitempty"`
	Data     any        `json:"data,omitempty"`
}

func SuccessResponseSwitcher(ctx *gin.Context, httpCode int, message string, data any) {
	var response SuccessResponse
	switch httpCode {
	case http.StatusOK:
		response = SuccessResponse{
			Code:    00,
			Message: message,
			Type:    "SUCCESS",
			Data:    data,
		}
	case http.StatusAccepted:
		response = SuccessResponse{
			Code:    02,
			Message: message,
			Type:    "ACCEPTED",
			Data:    data,
		}
	}
	ctx.JSONP(httpCode, response)
}

type ErrorResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	InvalidArg struct {
		ErrorType    string `json:"error_type"`
		ErrorMessage string `json:"error_message"`
	} `json:"invalid_arg"`
}

func ErrorResponseSwitcher(ctx *gin.Context, httpCode int, message string, errorType string) {
	var response ErrorResponse
	switch httpCode {
	case http.StatusBadRequest:
		response = ErrorResponse{
			Code:    96,
			Message: message,
			Type:    "BAD_REQUEST",
			InvalidArg: struct {
				ErrorType    string "json:\"error_type\""
				ErrorMessage string "json:\"error_message\""
			}{
				ErrorType:    errorType,
				ErrorMessage: message,
			},
		}
	case http.StatusUnauthorized:
		response = ErrorResponse{
			Code:    97,
			Message: message,
			Type:    "UNAUTHENTICATED",
			InvalidArg: struct {
				ErrorType    string "json:\"error_type\""
				ErrorMessage string "json:\"error_message\""
			}{
				ErrorType:    errorType,
				ErrorMessage: message,
			},
		}
	case http.StatusForbidden:
		response = ErrorResponse{
			Code:    98,
			Message: message,
			Type:    "FORBIDDEN",
			InvalidArg: struct {
				ErrorType    string "json:\"error_type\""
				ErrorMessage string "json:\"error_message\""
			}{
				ErrorType:    errorType,
				ErrorMessage: message,
			},
		}
	default:
		response = ErrorResponse{
			Code:    99,
			Message: message,
			Type:    "INTERNAL_SERVER_ERROR",
			InvalidArg: struct {
				ErrorType    string "json:\"error_type\""
				ErrorMessage string "json:\"error_message\""
			}{
				ErrorType:    errorType,
				ErrorMessage: message,
			},
		}
	}

	ctx.AbortWithStatusJSON(httpCode, response)
}
