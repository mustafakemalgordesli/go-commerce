package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mustafakemalgordesli/go-commerce/config"
)

func GenerateAccessToken(userId int) (string, error) {
	configs := config.GetConfig()

	secretKey := []byte(configs.Server.AccessTokenSecret)

	expDuration := configs.Server.AccessTokenExpireDuration

	return generateToken(userId, secretKey, expDuration)
}

func GenerateRefreshToken(userId int) (string, error) {
	configs := config.GetConfig()

	secretKey := []byte(configs.Server.RefreshTokenSecret)

	expDuration := configs.Server.RefreshTokenExpireDuration

	return generateToken(userId, secretKey, expDuration)
}

func generateToken(userId int, key []byte, expDuration int) (string, error) {

	expirationTime := time.Now().Add(time.Duration(expDuration) * time.Minute)
	//expirationTime := time.Now().Add(time.Duration(expDuration*60) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        userId,
		"ExpiresAt": jwt.NewNumericDate(expirationTime),
	})

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
