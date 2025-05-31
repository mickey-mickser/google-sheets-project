package responses

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ResponseError describes a standardized error payload.
// Code is a machine-readable string; Detail is a human-readable message.
type ResponseError struct {
	Detail string `json:"detail"`
	Code   string `json:"code"`
}

// Predefined error responses with codes
var (
	// 000 – Internal server error
	InternalErrorResponse = ResponseError{Detail: "Something bad happened", Code: "000"}
	// 002 – External ID is already in use
	ExternalIDInUse = ResponseError{Detail: "External id is used", Code: "002"}
	// 003 – User wallet not found
	UserWalletNotExist = ResponseError{Detail: "Such user wallet does not exist", Code: "003"}
	// 004 – Withdrawal not found
	WithdrawalNotExist = ResponseError{Detail: "Such withdrawal does not exist", Code: "004"}
)

// Constructors for dynamic errors

// NewBadRequestError returns a 400 error with code 001\ n
func NewBadRequestError(err error) ResponseError {
	return ResponseError{Detail: err.Error(), Code: "001"}
}

// NewMissingParamError returns a 400 error indicating a missing parameter
func NewMissingParamError(param string) ResponseError {
	return ResponseError{Detail: fmt.Sprintf("%s is required", param), Code: "005"}
}

// NewInvalidBodyError returns a 400 error for invalid JSON body
func NewInvalidBodyError(err error) ResponseError {
	return ResponseError{Detail: err.Error(), Code: "006"}
}

// NewUnauthorizedError returns a 401 error with custom detail
func NewUnauthorizedError(detail string) ResponseError {
	return ResponseError{Detail: detail, Code: "401"}
}

// NewNotFoundError returns a 404 error with custom detail
func NewNotFoundError(detail string) ResponseError {
	return ResponseError{Detail: detail, Code: "404"}
}

// NewConflictError returns a 409 error with custom detail
func NewConflictError(detail string) ResponseError {
	return ResponseError{Detail: detail, Code: "409"}
}
func BadRequestBind(c *gin.Context, handlerName string, err error) {
	logrus.WithField("handler", handlerName).
		WithError(err).
		Error("failed to bind request body")
	BadRequest(c, NewInvalidBodyError(err))
}
