package services

import (
	"errors"

	"github.com/mustafakemalgordesli/go-commerce/models"
	"gorm.io/gorm"
)

func CreateCart(db *gorm.DB, cart models.Cart) (models.Cart, error) {
	result := db.Create(&cart)

	if result.Error != nil {
		return models.Cart{}, result.Error
	}

	if !(result.RowsAffected > 0) {
		return models.Cart{}, errors.New("Cart not created")
	}

	return cart, nil
}
