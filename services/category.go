package services

import (
	"errors"

	"github.com/mustafakemalgordesli/go-commerce/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, category models.Category) (models.Category, error) {
	result := db.Create(&category)

	if result.Error != nil {
		return models.Category{}, result.Error
	}

	if !(result.RowsAffected > 0) {
		return models.Category{}, errors.New("Category not created")
	}

	return category, nil
}

func GetAllCategories(db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category
	dbRes := db.Find(&categories)
	if dbRes.Error != nil {
		return []models.Category{}, dbRes.Error
	}
	return categories, nil
}
