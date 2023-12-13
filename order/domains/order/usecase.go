package order

import (
	"context"
	
	_models "github.com/grpc-example-edts/order/models"
)

type Usecase interface {
	GetOrderByID(ctx context.Context, id string) (res *_models.Order, err error)
	GetOrderByUsername(ctx context.Context, username string) (res []*_models.Order, err error)
	AddOrder(ctx context.Context, param _models.OrderPost) (string, error)
	EditStatusOrder(ctx context.Context, id string, param _models.OrderStatusPatch) (string, error)
}
