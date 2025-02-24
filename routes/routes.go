package routes

import (
	"cardapio-api/internal/handlers"
	"cardapio-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// Configurar todas as rotas da API
func SetupRoutes(r *gin.Engine, cardapio_handler *handlers.CardapioHandler) {

	// Criar grupo de rotas públicas
	public := r.Group("/")
	{
		public.POST("/login", handlers.Login)
		public.GET("/cardapio", cardapio_handler.ListarCardapio)
		public.GET("/cardapio/:id", cardapio_handler.ObterItemPorID)
	}

	// Criar grupo de rotas privadas (com autenticação JWT)
	private := r.Group("/")
	private.Use(middlewares.AutenticarToken()) // Proteção JWT
	{
		private.POST("/cardapio", cardapio_handler.AdicionarItem)
		private.PUT("/cardapio/:id", cardapio_handler.AtualizarItem)
		private.DELETE("/cardapio/:id", cardapio_handler.RemoverItem)
	}
}
