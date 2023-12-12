package grpc

import (
	"context"
	"fmt"
	"log"
	
	"google.golang.org/grpc"
	
	_config "github.com/grpc-example-edts/order/config"
	_paymentApiPB "github.com/grpc-example-edts/order/pb/client/payment-api"
)

type dialPaymentApi struct {
	Payment _paymentApiPB.PaymentServiceClient
}

var initPaymentApi *dialPaymentApi = nil

func DialCustomerAPI(context context.Context, authorization string) (*dialPaymentApi, *grpc.ClientConn) {
	conn, err := grpc.Dial(_config.Env.PaymentApiHostGRPC, grpc.WithInsecure())
	if err != nil {
		log.Println(fmt.Printf("did not connect: %s", err))
	}
	
	initPaymentApi = &dialPaymentApi{
		Payment: _paymentApiPB.NewPaymentServiceClient(conn),
	}
	
	return initPaymentApi, conn
}
