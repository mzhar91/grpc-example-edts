package mysql

import (
	"context"
	"fmt"
	"reflect"
	
	"github.com/huandu/go-sqlbuilder"
	uuid "github.com/satori/go.uuid"
	
	_payment "github.com/grpc-example-edts/payment/domains/payment"
	_repository "github.com/grpc-example-edts/payment/helpers/repository"
	_models "github.com/grpc-example-edts/payment/models"
)

type mysqlRepository struct{}

func NewMysqlRepository() _payment.MysqlRepository {
	return &mysqlRepository{}
}

func (m *mysqlRepository) Create(ctx context.Context, tx *_repository.Use, param *_models.PaymentCreate) error {
	cols, values, err := _repository.Values(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}
	
	ib := sqlbuilder.MySQL.NewInsertBuilder()
	ib.InsertInto("payment")
	ib.Cols(cols...)
	ib.Values(values...)
	
	query, args := ib.Build()
	
	stmt, err := tx.Trans.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		err = fmt.Errorf("Weird behaviour. Total affected: %d", affect)
		return err
	}
	
	return nil
}

func (m *mysqlRepository) UpdateStatus(ctx context.Context, tx *_repository.Use, id uuid.UUID, param *_models.PaymentStatusUpdate) error {
	set, err := _repository.Set(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}
	
	ub := sqlbuilder.MySQL.NewUpdateBuilder()
	ub.Update("payment")
	ub.Set(set...)
	ub.Where(
		ub.Equal("id", id),
	)
	
	query, args := ub.Build()
	
	stmt, err := tx.Trans.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		err = fmt.Errorf("Weird behaviour. Total affected: %d", affect)
		return err
	}
	
	return nil
}

func (m *mysqlRepository) ReadByID(ctx context.Context, conn *_repository.Use, id uuid.UUID) (*_models.Payment, error) {
	sb := sqlbuilder.MySQL.NewSelectBuilder()
	sb.Select(
		"p.id",
		"p.order_id",
		"p.price",
		"p.status",
		"p.created_by",
		"p.created_at",
		"p.modified_by",
		"p.modified_at",
	)
	sb.From("payment AS p")
	sb.Where(
		sb.Equal("p.id", id),
	)
	
	query, args := sb.Build()
	
	result, err := m.fetchSingle(ctx, conn, query, args...)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}
