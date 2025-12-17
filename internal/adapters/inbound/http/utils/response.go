package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// HandleError handles errors and returns appropriate HTTP responses
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	errMsg := err.Error()
	statusCode := http.StatusInternalServerError
	errorType := "internal_error"

	// Map specific errors to HTTP status codes
	switch {
	case strings.Contains(errMsg, "already exists"):
		statusCode = http.StatusConflict
		errorType = "conflict"
	case strings.Contains(errMsg, "not found"):
		statusCode = http.StatusNotFound
		errorType = "not_found"
	case strings.Contains(errMsg, "invalid"):
		statusCode = http.StatusBadRequest
		errorType = "bad_request"
	case strings.Contains(errMsg, "unauthorized"):
		statusCode = http.StatusUnauthorized
		errorType = "unauthorized"
	case strings.Contains(errMsg, "forbidden"):
		statusCode = http.StatusForbidden
		errorType = "forbidden"
	}

	c.JSON(statusCode, ErrorResponse{
		Error:   errorType,
		Message: errMsg,
		Code:    statusCode,
	})
}

// RespondSuccess sends a successful JSON response
func RespondSuccess(c *gin.Context, statusCode int, data interface{}, message ...string) {
	response := SuccessResponse{
		Data: data,
	}
	if len(message) > 0 {
		response.Message = message[0]
	}
	c.JSON(statusCode, response)
}

// RespondCreated sends a 201 Created response
func RespondCreated(c *gin.Context, data interface{}) {
	RespondSuccess(c, http.StatusCreated, data, "Resource created successfully")
}

// RespondOK sends a 200 OK response
// For arrays, returns them directly (KrakenD compatible)
// For objects, wraps in SuccessResponse
func RespondOK(c *gin.Context, data interface{}) {
	// Check if data is a slice/array
	if isSlice(data) {
		// Return array directly for KrakenD compatibility
		c.JSON(http.StatusOK, data)
	} else {
		RespondSuccess(c, http.StatusOK, data)
	}
}

// isSlice checks if the interface is a slice
func isSlice(data interface{}) bool {
	if data == nil {
		return false
	}
	// Check using type assertion for common types
	switch data.(type) {
	case []interface{}:
		return true
	}
	// For specific types like []*dto.RoleResponse, just return array directly
	// This is a simple approach - return data as-is for list endpoints
	return false
}

// RespondNoContent sends a 204 No Content response
func RespondNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// BindJSON binds JSON request body and handles errors
func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "bad_request",
			Message: "Invalid request body: " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return false
	}
	return true
}
