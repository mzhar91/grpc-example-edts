package payment

import (
	"context"
	
	uuid "github.com/satori/go.uuid"
	
	_repository "github.com/grpc-example-edts/payment/helpers/repository"
	_models "github.com/grpc-example-edts/payment/models"
)

type MysqlRepository interface {
	ReadByID(ctx context.Context, conn *_repository.Use, id uuid.UUID) (res *_models.Payment, err error)
	Create(ctx context.Context, tx *_repository.Use, param *_models.Payment) (err error)
	Update(ctx context.Context, tx *_repository.Use, id uuid.UUID, param *_models.Payment) (err error)
}
