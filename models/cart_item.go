package models

import "time"

type CartItem struct {
	Id        int       `json:"id"`
	CartId    int       `json:"userid" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	IsActive  bool      `json:"isactive"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
