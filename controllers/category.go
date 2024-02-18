package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
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

func (categoryController CategoryController) AddCategory(c *gin.Context) {
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

func (categoryController CategoryController) GetAllCategories(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	number, err := strconv.Atoi(c.Query("number"))
	if err != nil || number == 0 {
		number = 1
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil || number == 0 {
		size = 8
	}

	categories, err := services.GetCategoriesPagination(ctx, categoryController.db, number, size)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	responseData := gin.H{
		"data":    categories,
		"success": true,
	}

	c.JSON(200, responseData)
}
