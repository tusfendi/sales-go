package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtAuthMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(key, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "gagal", "error": "Anda tidak ada Akses"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func TokenValid(key string, c *gin.Context) error {
	tokenString := ExtractToken(c)
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return err
	}

	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		c.Set("id", claims["id"])
		c.Set("email", claims["email"])
		c.Set("name", claims["name"])
	}

	result.Claims.Valid()

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
