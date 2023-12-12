package models

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

type Payment struct {
	ID         uuid.UUID   `json:"id" db:"id"`
	OrderID    uuid.UUID   `json:"orderId" db:"order_id"`
	Price      float64     `json:"price" db:"price"`
	Status     string      `json:"status" db:"status"`
	CreatedBy  string      `json:"createdBy" db:"created_by"`
	CreatedAt  int         `json:"createdAt" db:"created_at"`
	ModifiedBy null.String `json:"modifiedBy" db:"modified_by"`
	ModifiedAt null.Int    `json:"modifiedAt" db:"modified_at"`
}

type PaymentPost struct {
	ID      string  `json:"id" validate:"required"`
	Price   float64 `json:"price" validate:"required"`
	OrderID string  `json:"orderId" validate:"required"`
}

type PaymentConfirmPatch struct {
	Status bool `json:"status" validate:"required"`
}
