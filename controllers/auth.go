package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mustafakemalgordesli/go-commerce/controllers/request"
	"github.com/mustafakemalgordesli/go-commerce/controllers/response"
	"github.com/mustafakemalgordesli/go-commerce/models"
	"github.com/mustafakemalgordesli/go-commerce/pkg/events"
	"github.com/mustafakemalgordesli/go-commerce/pkg/helpers"
	"gorm.io/gorm"
)

var validate = validator.New()

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}

type Tokens struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

func (authController *AuthController) SignUp(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var registerRequest request.RegisterRequest

	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	if validationErr := validate.Struct(registerRequest); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": validationErr.Error()})
		return
	}

	var count int64

	dbRes := authController.db.Model(&models.User{}).Where("email = ?", registerRequest.Email).Count(&count)

	if dbRes.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": dbRes.Error})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Email is exist"})
		return
	}

	user, err := registerRequest.ToModel()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	result := authController.db.Create(&user)

	if err != result.Error {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	go func() {
		events.PublishMail(user.Email)
	}()

	accessToken, err := helpers.GenerateAccessToken(user.Id)
	refreshToken, err := helpers.GenerateRefreshToken(user.Id)

	var authResponse response.AuthResponse

	authResponse = authResponse.ToModel(user)

	responseData := gin.H{
		"data": authResponse,
		"tokens": map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
		"success": true,
	}

	c.JSON(200, responseData)
}
