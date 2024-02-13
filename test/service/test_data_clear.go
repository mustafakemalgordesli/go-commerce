package service

import (
	"context"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TruncateTestData(ctx context.Context, db *gorm.DB) {
	truncateProductResult := db.WithContext(ctx).Exec("TRUNCATE products RESTART IDENTITY")
	if truncateProductResult.Error != nil {
		log.Error(truncateProductResult.Error)
	} else {
		log.Info("Products table truncated")
	}

	truncateCategoriesResult := db.WithContext(ctx).Exec("TRUNCATE categories RESTART IDENTITY")
	if truncateCategoriesResult.Error != nil {
		log.Error(truncateCategoriesResult.Error)
	} else {
		log.Info("Categories table truncated")
	}
}
