package payment

import (
	"context"
	
	_models "github.com/grpc-example-edts/payment/models"
)

type Usecase interface {
	AddPayment(ctx context.Context, param _models.PaymentPost) (string, error)
	ConfirmPayment(ctx context.Context, id string, param _models.PaymentConfirmPatch) (string, error)
}
