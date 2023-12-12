package mysql

import (
	"context"
	"database/sql"
	"fmt"
	
	_repository "github.com/grpc-example-edts/order/helpers/repository"
	_models "github.com/grpc-example-edts/order/models"
)

func (m *mysqlRepository) fetchSingle(ctx context.Context, conn *_repository.Use, query string, args ...interface{}) (*_models.Order, error) {
	var rows *sql.Rows
	var err error
	
	if conn.Db != nil {
		rows, err = conn.Db.QueryContext(ctx, query, args...)
	} else if conn.Trans != nil {
		rows, err = conn.Trans.QueryContext(ctx, query, args...)
	}
	
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	
	for rows.Next() {
		t := new(_models.Order)
		
		err = rows.Scan(
			&t.ID,
			&t.Username,
			&t.Price,
			&t.Status,
			&t.CreatedBy,
			&t.CreatedAt,
			&t.ModifiedBy,
			&t.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}
		
		return t, nil
	}
	
	return nil, fmt.Errorf("Order not found")
}

func (m *mysqlRepository) fetchList(ctx context.Context, conn *_repository.Use, query string, args ...interface{}) ([]*_models.Order, error) {
	var rows *sql.Rows
	var err error
	
	if conn.Db != nil {
		rows, err = conn.Db.QueryContext(ctx, query, args...)
	} else if conn.Trans != nil {
		rows, err = conn.Trans.QueryContext(ctx, query, args...)
	}
	
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	
	result := make([]*_models.Order, 0)
	
	for rows.Next() {
		t := new(_models.Order)
		
		err = rows.Scan(
			&t.ID,
			&t.Username,
			&t.Price,
			&t.Status,
			&t.CreatedBy,
			&t.CreatedAt,
			&t.ModifiedBy,
			&t.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}
		
		result = append(result, t)
	}
	
	return result, nil
}
