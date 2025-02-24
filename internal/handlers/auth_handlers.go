package handlers

import (
	"cardapio-api/config"
	"cardapio-api/internal/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login gera um token JWT para o usuário autenticado
func Login(c *gin.Context) {
	var credenciais models.Usuario

	if err := c.ShouldBindJSON(&credenciais); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Credenciais inválidas"})
		return
	}

	if credenciais.Usuario != "admin" || credenciais.Senha != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário ou senha incorretos"})
		return
	}

	// Criar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario": credenciais.Usuario,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	// Assinar token com SecretKey
	tokenString, err := token.SignedString(config.SecretKey) // ✅ Agora funciona
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
