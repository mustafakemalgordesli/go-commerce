package models

import (
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Price       uint64    `json:"price" validate:"required"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageurl"`
	CategoryId  int       `json:"categoryid"`
	IsActive    bool      `json:"isactive"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}
