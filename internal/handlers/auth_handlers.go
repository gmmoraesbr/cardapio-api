package handlers

import (
	"cardapio-api/config"
	"cardapio-api/internal/models"
	"cardapio-api/pkg/logger"
	"cardapio-api/pkg/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login gera um token JWT para o usuário autenticado
func Login(c *gin.Context) {
	var credenciais models.Usuario

	// Tenta decodificar as credenciais do usuário
	if err := c.ShouldBindJSON(&credenciais); err != nil {
		logger.Error("Erro ao decodificar credenciais JSON: " + err.Error())
		response.Error(c, "Credenciais inválidas", err)
		return
	}

	// Validação do usuário e senha (hardcoded para exemplo)
	if credenciais.Usuario != "admin" || credenciais.Senha != "1234" {
		logger.Error("Tentativa de login com credenciais inválidas: Usuário=" + credenciais.Usuario)
		response.Error(c, "Usuário ou senha incorretos", nil)
		return
	}

	// Criar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario": credenciais.Usuario,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	// Assinar token com SecretKey
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		logger.Error("Erro ao gerar token JWT: " + err.Error())
		response.Error(c, "Erro ao gerar token", err)
		return
	}

	logger.Info("Login bem-sucedido para usuário: " + credenciais.Usuario)
	response.Success(c, "Login realizado com sucesso", gin.H{"token": tokenString})
}
