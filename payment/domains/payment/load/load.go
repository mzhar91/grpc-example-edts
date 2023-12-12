package load

import (
	"time"
	
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	
	_config "github.com/grpc-example-edts/payment/config"
	_grpc "github.com/grpc-example-edts/payment/domains/payment/delivery/grpc"
	_http "github.com/grpc-example-edts/payment/domains/payment/delivery/http"
	_paymentMysql "github.com/grpc-example-edts/payment/domains/payment/repository/mysql"
	_usecase "github.com/grpc-example-edts/payment/domains/payment/usecase"
	_mysql "github.com/grpc-example-edts/payment/helpers/repository/mysql"
	_paymentPB "github.com/grpc-example-edts/payment/pb/server/payment"
)

func Load(e *echo.Echo, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_mysql.Repository{
		Payment: _paymentMysql.NewMysqlRepository(),
	}
	
	ucase := _usecase.NewUcase(repo, connection, timeoutContext)
	
	_http.NewHandler(e, ucase)
}

func GrpcLoad(s *grpc.Server, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_mysql.Repository{
		Payment: _paymentMysql.NewMysqlRepository(),
	}
	
	ucase := _usecase.NewUcase(repo, connection, timeoutContext)
	guide := _grpc.NewGuide(ucase)
	
	_paymentPB.RegisterPaymentServiceServer(s, guide)
}
