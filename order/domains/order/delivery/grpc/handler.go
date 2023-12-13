package grpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	
	_order "github.com/grpc-example-edts/order/domains/order"
	_apiError "github.com/grpc-example-edts/order/helpers/apierror"
	_models "github.com/grpc-example-edts/order/models"
	_orderPB "github.com/grpc-example-edts/order/pb/server/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type routeGuideAccount struct {
	OrderUseCase _order.Usecase
	_orderPB.UnimplementedOrderServiceServer
	mu sync.Mutex
}

// NewHandler represent new handler
func NewGuide(pu _order.Usecase) _orderPB.OrderServiceServer {
	return &routeGuideAccount{
		OrderUseCase: pu,
	}
}

func (r *routeGuideAccount) EditOrderStatus(ctx context.Context, req *_orderPB.EditOrderStatusRequest) (*_orderPB.EditOrderStatusResponse, error) {
	var orderStatus string
	
	if req == nil {
		return nil, status.Errorf(
			codes.Canceled,
			fmt.Sprintf("Forbidden"),
		)
	}
	
	log.Printf("EditStatusOrder function was invoked with %v\n", req)
	
	if req.Status {
		orderStatus = "active"
	} else {
		orderStatus = "cancel"
	}
	
	_, err := r.OrderUseCase.EditStatusOrder(
		ctx, req.Id, _models.OrderStatusPatch{
			Status: orderStatus,
		},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Canceled,
			"failed to edit status order caused %v",
			err.(*_apiError.APIError).Message,
		)
	}
	
	res := &_orderPB.EditOrderStatusResponse{
		Id: req.Id,
	}
	
	return res, nil
}
