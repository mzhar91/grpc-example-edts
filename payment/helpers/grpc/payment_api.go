package grpc

import (
	"context"
	"fmt"
	"log"
	
	"google.golang.org/grpc"
	
	_config "github.com/grpc-example-edts/payment/config"
	_orderApiPB "github.com/grpc-example-edts/payment/pb/client/order-api"
)

type dialPaymentApi struct {
	Order _orderApiPB.OrderServiceClient
}

var initPaymentApi *dialPaymentApi = nil

func DialOrderAPI(context context.Context) (*dialPaymentApi, context.Context, *grpc.ClientConn) {
	conn, err := grpc.Dial(_config.Env.OrderApiHostGRPC, grpc.WithInsecure())
	if err != nil {
		log.Println(fmt.Printf("did not connect: %s", err))
	}
	
	initPaymentApi = &dialPaymentApi{
		Order: _orderApiPB.NewOrderServiceClient(conn),
	}
	
	return initPaymentApi, context, conn
}
