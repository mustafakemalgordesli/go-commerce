package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email" validate:"email,required"`
	FirstName string    `json:"firstname" validate:"required,min=3"`
	LastName  string    `json:"lastname"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
