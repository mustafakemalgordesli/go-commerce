package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Priorty   int       `json:"priorty"`
	ParentId  int       `json:"parentid"`
	IsActive  bool      `json:"isactive"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
