package middleware

import (
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
			return config.GetJWTSecret(), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwt.ErrTokenExpired {
				c.JSON(http.StatusUnauthorized, model.Response{Message: "Token inválido ou expirado."})
			} else {
				c.JSON(http.StatusBadRequest, model.Response{Message: "Token malformado."})
			}
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "Token inválido."})
			c.Abort()
			return
		}

		// Adiciona as claims ao contexto para uso posterior nos handlers
		c.Set("user_claims", claims)

		c.Next()
	}
}
