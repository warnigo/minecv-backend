package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents the standard API response structure
type Response struct {
	StatusCode int               `json:"status_code"`
	IsError    bool              `json:"is_error"`
	IsSuccess  bool              `json:"is_success"`
	Errors     map[string]string `json:"errors,omitempty"`
	Data       interface{}       `json:"data,omitempty"`
	Message    string            `json:"message"`
}

// Respond formats the response and sends it to the client
func Respond(c *gin.Context, statusCode int, isError bool, isSuccess bool, data interface{}, errors map[string]string, message string) {
	c.JSON(statusCode, Response{
		StatusCode: statusCode,
		IsError:    isError,
		IsSuccess:  isSuccess,
		Errors:     errors,
		Data:       data,
		Message:    message,
	})
}

// RespondValidationError handles validation errors specifically
func RespondValidationError(c *gin.Context, validationErrors map[string]string, message string) {
	Respond(c, http.StatusBadRequest, true, false, nil, validationErrors, message)
}

// RespondError handles generic errors
func RespondError(c *gin.Context, statusCode int, message string) {
	Respond(c, statusCode, true, false, nil, nil, message)
}

// RespondSuccess handles successful responses
func RespondSuccess(c *gin.Context, statusCode int, data interface{}, message string) {
	Respond(c, statusCode, false, true, data, nil, message)
}
