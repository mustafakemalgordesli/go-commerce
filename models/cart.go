package models

import "time"

type Cart struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userid" validate:"required"`
	IsActive  bool      `json:"isactive"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
