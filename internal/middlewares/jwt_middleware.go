package middlewares

import (
	"cardapio-api/config"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AutenticarToken é um middleware que valida o JWT
func AutenticarToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pegar token do header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token não fornecido"})
			c.Abort()
			return
		}

		// Remover "Bearer " do token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Validar token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
