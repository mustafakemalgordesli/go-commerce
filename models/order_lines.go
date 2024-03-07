package models

import "time"

type OrderLines struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"userid" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	Price     uint64    `json:"price" validate:"required"`
	IsActive  bool      `json:"isactive"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
