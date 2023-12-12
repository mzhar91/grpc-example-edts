package response

import (
	"github.com/labstack/echo"

	"github.com/grpc-example-edts/payment/helpers/apierror"
	"github.com/grpc-example-edts/payment/helpers/pagination"
)

// Success will write a default template response when returning a success response
func Success(c echo.Context, status int, data interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    status,
			"message": nil,
		},
		"data": data,
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}

// SuccessWithMessage will write a default template response when returning a success response
func SuccessWithMessage(c echo.Context, status int, data interface{}, params map[string]interface{}) error {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    params["code"],
			"message": params["message"],
		},
		"data": data,
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}

// SuccessWithPagination will write a default template response when returning a success response with pagination
func SuccessWithPagination(c echo.Context, status int, data interface{}, params map[string]interface{}) error {
	limit := params["limit"].(int)
	page := params["page"].(int)

	totalRows := params["totalRows"].(int)

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    params["code"],
			"message": params["message"],
		},
		"data":       data,
		"pagination": pagination.MyPagination(limit, totalRows, page),
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}

// Failed will write a default template response when returning a failed response
func Failed(c echo.Context, status int, err error) error {
	// if status/1e2 == 4 {
	// 	logger.Warn("%v", err)
	// } else {
	// 	logger.Err("%v", err)
	// }

	var errResponse map[string]interface{}

	if err != nil {
		errCode := 0
		errMsg := err.Error()
		if f, ok := err.(*apierror.APIError); ok {
			errCode = f.Code
			if status == 0 {
				status = f.Status
			}
			errMsg = f.Message
		}

		if errCode == 0 {
			errCode = status
		}

		errResponse = map[string]interface{}{
			"status":  status,
			"message": errMsg,
			"code":    errCode,
		}
	}

	response := errResponse

	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(status, response)
}
