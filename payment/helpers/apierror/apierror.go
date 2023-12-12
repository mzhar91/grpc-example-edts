package apierror

import (
	"fmt"

	"github.com/grpc-example-edts/payment/helpers/message"
)

// APIError define error to be returned in response
type APIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"-"`
}

var errorDictionary map[int]string

func (e *APIError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// WithMessage return APIError with message
func WithMessage(code int, message string, status ...int) error {
	httpStatus := checkStatus(status)
	return &APIError{
		Code:    code,
		Status:  httpStatus,
		Message: message,
	}
}

// WithData return APIError with data and message
func WithData(code int, message string, data interface{}, status ...int) error {
	httpStatus := checkStatus(status)
	return &APIError{
		Code:    code,
		Status:  httpStatus,
		Message: message,
		Data:    data,
	}
}

// FromError return APIError from error
func FromError(code int, err error, status ...int) error {
	httpStatus := checkStatus(status)
	return &APIError{
		Code:    code,
		Status:  httpStatus,
		Message: err.Error(),
	}
}

// FromErrorCode return APIError from error
func FromErrorCode(code int, status ...int) error {
	httpStatus := checkStatus(status)
	return &APIError{
		Code:    code,
		Status:  httpStatus,
		Message: errorDictionary[code],
	}
}

// FromErrorCodeWithParam return APIError from error
func FromErrorCodeWithParam(code int, param []interface{}, status ...int) error {
	httpStatus := checkStatus(status)
	return &APIError{
		Code:    code,
		Status:  httpStatus,
		Message: fmt.Sprintf(errorDictionary[code], param...),
	}
}

// Setup error message dictionary
func Setup() {
	errorDictionary = message.InitErrorMessage()
}

func checkStatus(status []int) int {
	if len(status) == 0 {
		return 0
	}
	return status[0]
}
