package http

import (
	"context"
	"net/http"
	
	_apiError "github.com/grpc-example-edts/payment/helpers/apierror"
	_message "github.com/grpc-example-edts/payment/helpers/message"
	"github.com/labstack/echo"
	
	_payment "github.com/grpc-example-edts/payment/domains/payment"
	_request "github.com/grpc-example-edts/payment/helpers/request"
	_response "github.com/grpc-example-edts/payment/helpers/response"
	_models "github.com/grpc-example-edts/payment/models"
)

type Handler struct {
	PaymentUseCase _payment.Usecase
}

// NewHandler represent new handler
func NewHandler(e *echo.Echo, us _payment.Usecase) {
	handler := &Handler{
		PaymentUseCase: us,
	}
	
	gu := e.Group("/")
	gu.GET("healthcheck", handler.healthcheck)
	gu.PATCH("confirm/:id", handler.confirmPayment)
}

// Healthcheck to check if the service is running
func (a *Handler) healthcheck(c echo.Context) error {
	msg := map[string]interface{}{"message": "OK"}
	
	return _response.SuccessWithMessage(c, http.StatusOK, nil, msg)
}

// confirmPayment function
func (a *Handler) confirmPayment(c echo.Context) error {
	var req _models.PaymentConfirmPatch
	
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	if err := c.Bind(&req); err != nil {
		return _response.Failed(c, http.StatusBadRequest, _apiError.FromErrorCode(_message.IncorrectFormat))
	}
	
	if okBind, err := _request.IsRequestValid(&req); !okBind {
		return _response.Failed(c, http.StatusBadRequest, err)
	}
	
	id := c.Param("id")
	data, err := a.PaymentUseCase.ConfirmPayment(ctx, id, req)
	if err != nil {
		if errApi, ok := err.(*_apiError.APIError); ok {
			return _response.Failed(c, errApi.Status, err)
		}
		
		return _response.Failed(c, http.StatusInternalServerError, err)
	}
	
	return _response.Success(c, http.StatusOK, data)
}
