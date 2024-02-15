package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mustafakemalgordesli/go-commerce/controllers/request"
	"github.com/mustafakemalgordesli/go-commerce/services"
	"gorm.io/gorm"
)

type CategoryController struct {
	db *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{
		db: db,
	}
}

func (categoryController CategoryController) Add(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var categoryRequest request.CategoryCreateRequest

	if err := c.BindJSON(&categoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if validationErr := validate.Struct(categoryRequest); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": validationErr.Error()})
		return
	}

	if categoryRequest.ParentId != 0 {
		parentCategory, err := services.GetCategoryById(ctx, categoryController.db, categoryRequest.ParentId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err})
			return
		}

		if parentCategory == nil || parentCategory.IsActive == false {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}
	}

	category := categoryRequest.ToModel()

	fmt.Println(category.IsActive)

	fmt.Println(category)

	responseData := gin.H{

		"success": true,
	}

	c.JSON(200, responseData)
}
