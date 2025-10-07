package middleware

import (
	"net/http"
	"os"
	"strings"
	"wheels-api/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// Busca o token da variável de ambiente uma vez na inicialização
	authToken := os.Getenv("AUTH_TOKEN")

	// Checagem se a variável de ambiente existe
	if authToken == "" {
		// Retorna um handler que sempre falha se a configuração estiver faltando
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, model.Response{Message: "Configuração de autenticação faltando no servidor (AUTH_TOKEN)."})
			c.Abort()
		}
	}

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

		token := parts[1]

		if token != authToken {
			c.JSON(http.StatusUnauthorized, model.Response{Message: "Token inválido."})
			c.Abort()
			return
		}

		c.Next()
	}
}
