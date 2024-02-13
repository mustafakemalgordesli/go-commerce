package services

import (
	"errors"

	"github.com/mustafakemalgordesli/go-commerce/models"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product) (models.Product, error) {
	result := db.Create(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	if !(result.RowsAffected > 0) {
		return models.Product{}, errors.New("Product not created")
	}

	return product, nil
}
