package mysql

import (
	"context"
	"fmt"
	"reflect"
	
	"github.com/huandu/go-sqlbuilder"
	uuid "github.com/satori/go.uuid"
	
	_order "github.com/grpc-example-edts/order/domains/order"
	_repository "github.com/grpc-example-edts/order/helpers/repository"
	_models "github.com/grpc-example-edts/order/models"
)

type mysqlRepository struct{}

func NewMysqlRepository() _order.MysqlRepository {
	return &mysqlRepository{}
}

func (m *mysqlRepository) Create(ctx context.Context, tx *_repository.Use, param *_models.OrderCreate) error {
	cols, values, err := _repository.Values(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}
	
	ib := sqlbuilder.MySQL.NewInsertBuilder()
	ib.InsertInto("`order`")
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

func (m *mysqlRepository) UpdateStatus(ctx context.Context, tx *_repository.Use, id uuid.UUID, param *_models.OrderStatusUpdate) error {
	set, err := _repository.Set(reflect.ValueOf(param).Elem())
	if err != nil {
		return err
	}
	
	ub := sqlbuilder.MySQL.NewUpdateBuilder()
	ub.Update("`order`")
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

func (m *mysqlRepository) ReadByID(ctx context.Context, conn *_repository.Use, id uuid.UUID) (*_models.Order, error) {
	sb := sqlbuilder.MySQL.NewSelectBuilder()
	sb.Select(
		"o.id",
		"o.username",
		"o.price",
		"o.status",
		"o.created_by",
		"o.created_at",
		"o.modified_by",
		"o.modified_at",
	)
	sb.From("order AS o")
	sb.Where(
		sb.Equal("o.id", id),
	)
	
	query, args := sb.Build()
	
	result, err := m.fetchSingle(ctx, conn, query, args...)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}

func (m *mysqlRepository) ReadByUsername(ctx context.Context, conn *_repository.Use, username string) ([]*_models.Order, error) {
	sb := sqlbuilder.MySQL.NewSelectBuilder()
	sb.Select(
		"o.id",
		"o.username",
		"o.price",
		"o.status",
		"o.created_by",
		"o.created_at",
		"o.modified_by",
		"o.modified_at",
	)
	sb.From("order AS o")
	sb.Where(
		sb.Equal("o.username", username),
	)
	
	query, args := sb.Build()
	
	result, err := m.fetchList(ctx, conn, query, args...)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}
