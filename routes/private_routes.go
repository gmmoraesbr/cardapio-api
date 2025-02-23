package routes

import (
	"cardapio-api/controllers"
	"cardapio-api/middlewares"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(r *gin.RouterGroup) {
	r.Use(middlewares.AutenticarToken())

	r.POST("/cardapio", controllers.AdicionarItem)
	r.PUT("/cardapio/:id", controllers.AtualizarItem)  // Atualiza item por ID
	r.DELETE("/cardapio/:id", controllers.RemoverItem) // Remove item por ID
}
