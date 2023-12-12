package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
	
	_config "github.com/grpc-example-edts/order/config"
	_order "github.com/grpc-example-edts/order/domains/order"
	_apiError "github.com/grpc-example-edts/order/helpers/apierror"
	_repository "github.com/grpc-example-edts/order/helpers/repository"
	_mysql "github.com/grpc-example-edts/order/helpers/repository/mysql"
	_models "github.com/grpc-example-edts/order/models"
)

type ucase struct {
	orderRepo      _order.MysqlRepository
	contextTimeout time.Duration
	dbConn         *sql.DB
	debug          bool
}

func NewUcase(mysql *_mysql.Repository, connection *_config.Connection, timeout time.Duration) _order.Usecase {
	return &ucase{
		orderRepo:      mysql.Order,
		dbConn:         connection.Database,
		contextTimeout: timeout,
		debug:          _config.Env.Debug,
	}
}

func (a *ucase) AddOrder(ctx context.Context, param _models.OrderPost) (string, error) {
	var result string
	
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	
	// transaction
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			conn := &_repository.Use{
				Trans: tx,
			}
			now := time.Now()
			id := uuid.NewV4()
			
			err := a.orderRepo.Create(
				ctx, conn, &_models.Order{
					ID:        id,
					Username:  strconv.Itoa(rand.Intn(100)) + "@mail.com",
					Price:     param.Price,
					Status:    "inactive",
					CreatedBy: "system",
					CreatedAt: int(now.Unix()),
				},
			)
			if err != nil {
				log.Printf(err.Error())
				
				return err, http.StatusInternalServerError
			}
			
			result = id.String()
			
			return nil, http.StatusOK
		},
	)
	if err != nil {
		if a.debug || code == http.StatusUnprocessableEntity {
			return result, _apiError.WithMessage(0, fmt.Sprintf("Create Order failed caused: %v", err.Error()), code)
		}
		
		return result, _apiError.WithMessage(0, "Create Order failed", code)
	}
	
	return result, nil
}

func (a *ucase) EditOrder(ctx context.Context, id string, param _models.OrderPatch) (string, error) {
	var result string
	
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			conn := &_repository.Use{
				Trans: tx,
			}
			now := time.Now()
			
			// check if account exist
			_, err := a.orderRepo.ReadByID(ctx, conn, uuid.FromStringOrNil(id))
			if err != nil {
				log.Printf(err.Error())
				
				return err, http.StatusInternalServerError
			}
			
			err = a.orderRepo.Update(
				ctx, conn, uuid.FromStringOrNil(id), &_models.Order{
					Status:     param.Status,
					ModifiedBy: null.StringFrom("system"),
					ModifiedAt: null.IntFrom(now.Unix()),
				},
			)
			if err != nil {
				log.Printf(err.Error())
				
				return err, http.StatusInternalServerError
			}
			
			result = id
			
			return nil, http.StatusOK
		},
	)
	
	if err != nil {
		if a.debug || code == http.StatusUnprocessableEntity {
			return result, _apiError.WithMessage(0, fmt.Sprintf("Update Order failed caused: %v", err.Error()), code)
		}
		
		return result, _apiError.WithMessage(0, "Update Order failed", code)
	}
	
	return result, nil
}

func (a *ucase) GetOrderByID(ctx context.Context, id string) (*_models.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	
	conn := &_repository.Use{
		Db: a.dbConn,
	}
	
	result, err := a.orderRepo.ReadByID(ctx, conn, uuid.FromStringOrNil(id))
	if err != nil {
		log.Printf(err.Error())
		
		if a.debug {
			return nil, _apiError.WithMessage(0, err.Error(), http.StatusOK)
		}
		
		return nil, _apiError.WithMessage(0, "Order not found", http.StatusOK)
	}
	
	return result, nil
}

func (a *ucase) GetOrderByUsername(ctx context.Context, username string) ([]*_models.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	
	conn := &_repository.Use{
		Db: a.dbConn,
	}
	
	result, err := a.orderRepo.ReadByUsername(ctx, conn, username)
	if err != nil {
		log.Printf(err.Error())
		
		if a.debug {
			return nil, _apiError.WithMessage(0, err.Error(), http.StatusOK)
		}
		
		return nil, _apiError.WithMessage(0, "Order not found", http.StatusOK)
	}
	
	return result, nil
}
