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
	CreatedAt  int         `json:"createdAt" db:"created_at"`
	ModifiedBy null.String `json:"modifiedBy" db:"modified_by"`
	ModifiedAt null.Int    `json:"modifiedAt" db:"modified_at"`
}

type OrderPost struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price" validate:"required"`
	Username string  `json:"username" validate:"required"`
}

type OrderPatch struct {
	Status string `json:"status"`
}
