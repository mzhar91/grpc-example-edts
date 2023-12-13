package models

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

type Order struct {
	ID         uuid.UUID   `json:"id" db:"id"`
	Price      float64     `json:"price" db:"price"`
	Username   string      `json:"username" db:"username"`
	Status     string      `json:"status" db:"status"`
	CreatedBy  string      `json:"createdBy" db:"created_by"`
	CreatedAt  int64       `json:"createdAt" db:"created_at"`
	ModifiedBy null.String `json:"modifiedBy" db:"modified_by"`
	ModifiedAt null.Int    `json:"modifiedAt" db:"modified_at"`
}

type OrderCreate struct {
	ID        uuid.UUID `db:"id"`
	Price     float64   `db:"price"`
	Username  string    `db:"username"`
	Status    string    `db:"status"`
	CreatedBy string    `db:"created_by"`
	CreatedAt int64     `db:"created_at"`
}

type OrderStatusUpdate struct {
	Status     string `db:"status"`
	ModifiedBy string `db:"modified_by"`
	ModifiedAt int64  `db:"modified_at"`
}

type OrderPost struct {
	Price float64 `json:"price" validate:"required"`
}

type OrderStatusPatch struct {
	Status string `json:"status"`
}
