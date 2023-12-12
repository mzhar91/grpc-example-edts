package load

import (
	"time"
	
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	
	_config "github.com/grpc-example-edts/order/config"
	_grpc "github.com/grpc-example-edts/order/domains/order/delivery/grpc"
	_http "github.com/grpc-example-edts/order/domains/order/delivery/http"
	_orderMysql "github.com/grpc-example-edts/order/domains/order/repository/mysql"
	_usecase "github.com/grpc-example-edts/order/domains/order/usecase"
	_mysql "github.com/grpc-example-edts/order/helpers/repository/mysql"
	_orderPB "github.com/grpc-example-edts/order/pb/server/order"
)

func Load(e *echo.Echo, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_mysql.Repository{
		Order: _orderMysql.NewMysqlRepository(),
	}
	
	ucase := _usecase.NewUcase(repo, connection, timeoutContext)
	
	_http.NewHandler(e, ucase)
}

func GrpcLoad(s *grpc.Server, connection *_config.Connection, timeoutContext time.Duration) {
	repo := &_mysql.Repository{
		Order: _orderMysql.NewMysqlRepository(),
	}
	
	ucase := _usecase.NewUcase(repo, connection, timeoutContext)
	guide := _grpc.NewGuide(ucase)
	
	_orderPB.RegisterAccountServiceServer(s, guide)
}
