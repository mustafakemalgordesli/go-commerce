package services

import (
	"github.com/mustafakemalgordesli/go-commerce/models"
	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User

	dbRes := db.Model(&models.User{}).Where("email = ?", email).First(&user)

	return user, dbRes.Error
}
