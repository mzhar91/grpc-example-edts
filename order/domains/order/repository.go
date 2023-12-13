package order

import (
	"context"
	
	uuid "github.com/satori/go.uuid"
	
	_repository "github.com/grpc-example-edts/order/helpers/repository"
	_models "github.com/grpc-example-edts/order/models"
)

type MysqlRepository interface {
	ReadByID(ctx context.Context, conn *_repository.Use, id uuid.UUID) (res *_models.Order, err error)
	ReadByUsername(ctx context.Context, conn *_repository.Use, username string) (res []*_models.Order, err error)
	Create(ctx context.Context, tx *_repository.Use, param *_models.OrderCreate) (err error)
	UpdateStatus(ctx context.Context, tx *_repository.Use, id uuid.UUID, param *_models.OrderStatusUpdate) (err error)
}
