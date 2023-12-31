package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	
	_orderApiPB "github.com/grpc-example-edts/payment/pb/client/order-api"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	
	_config "github.com/grpc-example-edts/payment/config"
	_payment "github.com/grpc-example-edts/payment/domains/payment"
	_apiError "github.com/grpc-example-edts/payment/helpers/apierror"
	_grpcHelper "github.com/grpc-example-edts/payment/helpers/grpc"
	_repository "github.com/grpc-example-edts/payment/helpers/repository"
	_mysql "github.com/grpc-example-edts/payment/helpers/repository/mysql"
	_models "github.com/grpc-example-edts/payment/models"
)

type ucase struct {
	paymentRepo    _payment.MysqlRepository
	contextTimeout time.Duration
	dbConn         *sql.DB
	debug          bool
}

func NewUcase(mysql *_mysql.Repository, connection *_config.Connection, timeout time.Duration) _payment.Usecase {
	return &ucase{
		paymentRepo:    mysql.Payment,
		dbConn:         connection.Database,
		contextTimeout: timeout,
		debug:          _config.Env.Debug,
	}
}

func (a *ucase) AddPayment(ctx context.Context, param _models.PaymentPost) (string, error) {
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
			
			err := a.paymentRepo.Create(
				ctx, conn, &_models.PaymentCreate{
					ID:        id,
					OrderID:   uuid.FromStringOrNil(param.OrderID),
					Price:     param.Price,
					Status:    "pending",
					CreatedBy: "system",
					CreatedAt: now.Unix(),
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
			return result, _apiError.WithMessage(0, fmt.Sprintf("Create Payment failed caused: %v", err.Error()), code)
		}
		
		return result, _apiError.WithMessage(0, "Create Payment failed", code)
	}
	
	return result, nil
}

func (a *ucase) ConfirmPayment(ctx context.Context, id string, param _models.PaymentConfirmPatch) (string, error) {
	var result string
	
	err, code := _repository.WithTransaction(
		a.dbConn, func(tx _repository.Transaction) (error, int) {
			var status string
			var orderStatus bool
			
			conn := &_repository.Use{
				Trans: tx,
			}
			now := time.Now()
			
			// check if account exist
			getPayment, err := a.paymentRepo.ReadByID(ctx, conn, uuid.FromStringOrNil(id))
			if err != nil {
				log.Printf(err.Error())
				
				return err, http.StatusInternalServerError
			}
			
			if param.Status {
				status = "paid"
				orderStatus = true
			} else {
				status = "cancel"
				orderStatus = false
			}
			
			err = a.paymentRepo.UpdateStatus(
				ctx, conn, uuid.FromStringOrNil(id), &_models.PaymentStatusUpdate{
					Status:     status,
					ModifiedBy: "system",
					ModifiedAt: now.Unix(),
				},
			)
			if err != nil {
				log.Printf(err.Error())
				
				return err, http.StatusInternalServerError
			}
			
			// calling grpc function for update order status
			grpcOrder, ctxPayment, grpcConn := _grpcHelper.DialOrderAPI(ctx)
			defer func(grpcConn *grpc.ClientConn) {
				err := grpcConn.Close()
				if err != nil {
					log.Printf(err.Error())
				}
			}(grpcConn)
			
			updateOrder, err := grpcOrder.Order.EditOrderStatus(
				ctxPayment, &_orderApiPB.EditOrderStatusRequest{
					Id:     getPayment.OrderID.String(),
					Status: orderStatus,
				},
			)
			if err != nil {
				log.Printf(err.Error())
				
				return err, http.StatusInternalServerError
			}
			
			result = id + " | " + updateOrder.Id
			
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
