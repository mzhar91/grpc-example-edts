package mysql

import (
	_payment "github.com/grpc-example-edts/payment/domains/payment"
)

type Repository struct {
	Payment _payment.MysqlRepository
}
