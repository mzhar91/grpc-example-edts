package load

import (
	"time"
	
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	
	_config "github.com/grpc-example-edts/payment/config"
	_loadOrder "github.com/grpc-example-edts/payment/domains/payment/load"
)

func Load(e *echo.Echo, connection *_config.Connection, timeoutContext time.Duration) {
	_loadOrder.Load(e, connection, timeoutContext)
}

func GrpcLoad(s *grpc.Server, connection *_config.Connection, timeoutContext time.Duration) {
	_loadOrder.GrpcLoad(s, connection, timeoutContext)
}
