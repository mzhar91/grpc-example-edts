package mysql

import (
	_order "github.com/grpc-example-edts/order/domains/order"
)

type Repository struct {
	Order _order.MysqlRepository
}
