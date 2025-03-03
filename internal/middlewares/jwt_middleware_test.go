package middlewares_test

import (
	"cardapio-api/internal/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAutenticarToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(middlewares.AutenticarToken())
	r.GET("/protegido", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })

	t.Run("Sem token retorna 401", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/protegido", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Esperado 401, obteve %d", w.Code)
		}
	})
}
