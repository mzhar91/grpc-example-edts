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
	_orderPB.UnimplementedAccountServiceServer
	mu sync.Mutex
}

// NewHandler represent new handler
func NewGuide(pu _order.Usecase) _orderPB.AccountServiceServer {
	return &routeGuideAccount{
		OrderUseCase: pu,
	}
}

func (r *routeGuideAccount) EditOrder(ctx context.Context, req *_orderPB.EditOrderRequest) (*_orderPB.EditOrderResponse, error) {
	if req == nil {
		return nil, status.Errorf(
			codes.Canceled,
			fmt.Sprintf("Forbidden"),
		)
	}
	
	log.Printf("EditOrder function was invoked with %v\n", req)
	
	_, err := r.OrderUseCase.EditOrder(
		ctx, req.Id, _models.OrderPatch{
			Status: req.Status,
		},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Canceled,
			"failed to edit order caused %v",
			err.(*_apiError.APIError).Message,
		)
	}
	
	res := &_orderPB.EditOrderResponse{
		Id: req.Id,
	}
	
	return res, nil
}
