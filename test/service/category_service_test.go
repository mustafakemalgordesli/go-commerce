package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mustafakemalgordesli/go-commerce/config"
	"github.com/mustafakemalgordesli/go-commerce/models"
	"github.com/mustafakemalgordesli/go-commerce/pkg/database"
	"github.com/mustafakemalgordesli/go-commerce/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	if err := config.SetupPath("./../.."); err != nil {
		log.Fatalf("config.Setup() error: %s", err)
		os.Exit(1)
	}
	if err := database.Setup(); err != nil {
		log.Fatalf("database.Setup() error: %s", err)
		fmt.Println("database" + err.Error())
		os.Exit(1)
	}
	db = database.GetDB()
	fmt.Println("Before all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func setup(ctx context.Context, db *gorm.DB) {
	TestDataInitialize(ctx, db)
}
func clear(ctx context.Context, db *gorm.DB) {
	TruncateTestData(ctx, db)
}

func TestAddCategory(t *testing.T) {
	expectedCategory := models.Category{
		Id:      1,
		Name:    "Technology",
		Priorty: 1,
	}
	newCategory := models.Category{
		Id:      1,
		Name:    "Technology",
		Priorty: 1,
	}
	t.Run("AddCategory", func(t *testing.T) {
		services.CreateCategory(db, newCategory)
		actualCategoriess, err := services.GetAllCategories(db)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(actualCategoriess))
		assert.Equal(t, expectedCategory.Id, actualCategoriess[0].Id)
		assert.Equal(t, expectedCategory.Name, actualCategoriess[0].Name)
		assert.Equal(t, expectedCategory.Priorty, actualCategoriess[0].Priorty)
	})

	clear(ctx, db)
}
