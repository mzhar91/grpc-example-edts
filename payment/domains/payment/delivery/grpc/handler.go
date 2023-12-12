package grpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	
	_payment "github.com/grpc-example-edts/payment/domains/payment"
	_apiError "github.com/grpc-example-edts/payment/helpers/apierror"
	_models "github.com/grpc-example-edts/payment/models"
	_paymentPB "github.com/grpc-example-edts/payment/pb/server/payment"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type routeGuideAccount struct {
	PaymentUseCase _payment.Usecase
	_paymentPB.UnimplementedPaymentServiceServer
	mu sync.Mutex
}

// NewHandler represent new handler
func NewGuide(pu _payment.Usecase) _paymentPB.PaymentServiceServer {
	return &routeGuideAccount{
		PaymentUseCase: pu,
	}
}

func (r *routeGuideAccount) AddPayment(ctx context.Context, req *_paymentPB.AddPaymentRequest) (*_paymentPB.AddPaymentResponse, error) {
	if req == nil {
		return nil, status.Errorf(
			codes.Canceled,
			fmt.Sprintf("Forbidden"),
		)
	}
	
	log.Printf("AddPayment function was invoked with %v\n", req)
	
	paymentId := uuid.NewV4().String()
	
	_, err := r.PaymentUseCase.AddPayment(
		ctx, _models.PaymentPost{
			ID:      paymentId,
			Price:   float64(req.Price),
			OrderID: req.OrderId,
		},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Canceled,
			"failed to edit order caused %v",
			err.(*_apiError.APIError).Message,
		)
	}
	
	res := &_paymentPB.AddPaymentResponse{
		Id: paymentId,
	}
	
	return res, nil
}
