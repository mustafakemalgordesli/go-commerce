package service

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/mustafakemalgordesli/go-commerce/models"
	"gorm.io/gorm"
)

var INSERT_PRODUCTS = []*models.Product{}

func TestDataInitialize(ctx context.Context, db *gorm.DB) {
	insertProductsResult := db.WithContext(ctx).Create(INSERT_PRODUCTS)
	if insertProductsResult.Error != nil {
		log.Error(insertProductsResult.Error)
	} else {
		log.Info(fmt.Sprintf("Products data created with %d rows", insertProductsResult.RowsAffected))
	}

	var INSERT_CATEGORIES = []*models.Category{
		{
			Name:    "Technology",
			Priorty: 1,
		},
	}

	insertCategoriesResult := db.WithContext(ctx).Create(INSERT_CATEGORIES)
	if insertCategoriesResult.Error != nil {
		log.Error(insertCategoriesResult.Error)
	} else {
		log.Info(fmt.Sprintf("Categories data created with %d rows", insertCategoriesResult.RowsAffected))
	}
}
