package services

import (
	"context"
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

func GetProductsPagination(ctx context.Context, db *gorm.DB, pageNumber int, pageSize int) ([]models.Product, error) {
	var products []models.Product
	dbRes := db.WithContext(ctx).Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&products)
	if dbRes.Error != nil {
		return nil, db.Error
	}
	return products, nil
}
