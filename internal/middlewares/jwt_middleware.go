package middlewares

import (
	"cardapio-api/config"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica o token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
			c.Abort()
			return
		}

		// Remover "Bearer " do token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("Método de assinatura inválido", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Verifica tempo de expiração
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			c.Abort()
			return
		}

		// Adiciona os claims ao contexto
		c.Set("usuario", claims["usuario"])
		c.Next()
	}
}

// GenerateTokens cria um access token e um refresh token
func GenerateTokens(usuario string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario": usuario,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario": usuario,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(config.SecretRefresh))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// RefreshTokenHandler gera um novo access token usando um refresh token
func RefreshTokenHandler(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	token, err := jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Método de assinatura inválido", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(config.SecretRefresh), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expirado"})
		return
	}

	usuario := claims["usuario"].(string)
	newAccessToken, newRefreshToken, err := GenerateTokens(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar novo token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
