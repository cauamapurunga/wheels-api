package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"wheels-api/config"
	"wheels-api/model"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "Requer autenticação Bearer Token."})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "Formato do token inválido. Use 'Bearer [token]'."})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &usecase.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.GetJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			log.Printf("Error parsing token: %v", err)
			c.JSON(http.StatusUnauthorized, model.Response{Message: "Token inválido ou expirado."})
			c.Abort()
			return
		}

		// Adiciona as claims ao contexto para uso posterior nos handlers
		c.Set("user_claims", claims)

		c.Next()
	}
}
