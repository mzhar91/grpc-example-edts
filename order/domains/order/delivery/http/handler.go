package http

import (
	"context"
	"net/http"
	
	_apiError "github.com/grpc-example-edts/order/helpers/apierror"
	_message "github.com/grpc-example-edts/order/helpers/message"
	"github.com/labstack/echo"
	
	_order "github.com/grpc-example-edts/order/domains/order"
	_request "github.com/grpc-example-edts/order/helpers/request"
	_response "github.com/grpc-example-edts/order/helpers/response"
	_models "github.com/grpc-example-edts/order/models"
)

type Handler struct {
	OrderUseCase _order.Usecase
}

// NewHandler represent new handler
func NewHandler(e *echo.Echo, us _order.Usecase) {
	handler := &Handler{
		OrderUseCase: us,
	}
	
	gu := e.Group("/")
	gu.GET("healthcheck", handler.healthcheck)
	gu.GET("id/:id", handler.getOrderByID)
	gu.GET("username/:username", handler.getOrderByUsername)
	gu.POST("", handler.addOrder)
}

// Healthcheck to check if the service is running
func (a *Handler) healthcheck(c echo.Context) error {
	msg := map[string]interface{}{"message": "OK"}
	
	return _response.SuccessWithMessage(c, http.StatusOK, nil, msg)
}

// getOrderByID show order by ID
func (a *Handler) getOrderByID(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	id := c.Param("id")
	data, err := a.OrderUseCase.GetOrderByID(
		ctx, id,
	)
	if err != nil {
		if errApi, ok := err.(*_apiError.APIError); ok {
			return _response.Failed(c, errApi.Status, err)
		}
		
		return _response.Failed(c, http.StatusInternalServerError, err)
	}
	
	return _response.Success(c, http.StatusOK, data)
}

// addOrder function
func (a *Handler) addOrder(c echo.Context) error {
	var req _models.OrderPost
	
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
	
	data, err := a.OrderUseCase.AddOrder(ctx, req)
	if err != nil {
		if errApi, ok := err.(*_apiError.APIError); ok {
			return _response.Failed(c, errApi.Status, err)
		}
		
		return _response.Failed(c, http.StatusInternalServerError, err)
	}
	
	return _response.Success(c, http.StatusOK, data)
}

// getOrderByUsername show order by Username
func (a *Handler) getOrderByUsername(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	
	id := c.Param("id")
	data, err := a.OrderUseCase.GetOrderByID(
		ctx, id,
	)
	if err != nil {
		if errApi, ok := err.(*_apiError.APIError); ok {
			return _response.Failed(c, errApi.Status, err)
		}
		
		return _response.Failed(c, http.StatusInternalServerError, err)
	}
	
	return _response.Success(c, http.StatusOK, data)
}
