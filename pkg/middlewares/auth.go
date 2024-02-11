package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mustafakemalgordesli/go-commerce/config"
)

func VerifyJWT() gin.HandlerFunc {
	configs := config.GetConfig()

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(tokenString, " ")
		tokenString = splitToken[1]

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(configs.Server.AccessTokenSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false})
			return
		}

		expiresAt := int(claims["ExpiresAt"].(float64))

		fmt.Println(int(time.Now().Unix()))
		if expiresAt < int(time.Now().Unix()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false})
			return
		}

		id := int(claims["id"].(float64))
		c.Set("userId", id)
		c.Next()
	}
}
