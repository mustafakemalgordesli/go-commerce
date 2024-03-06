package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mustafakemalgordesli/go-commerce/services"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (productController ProductController) GetAllProductsByPagination(c *gin.Context) {
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

	products, err := services.GetProductsPagination(ctx, productController.db, number, size)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	responseData := gin.H{
		"data":    products,
		"success": true,
	}

	c.JSON(200, responseData)
}
