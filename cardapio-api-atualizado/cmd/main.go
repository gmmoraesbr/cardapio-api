package main

import (
	"cardapio-api/config"
	"cardapio-api/internal/controllers"
	"cardapio-api/internal/repositories"
	"cardapio-api/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Configurar MongoDB
	db := config.ConectarMongo(cfg.MongoURI, cfg.Database)

	// Criar repositório e controller
	repo := repositories.NovoCardapioRepository(db)
	controller := controllers.CardapioController{Repo: repo}

	// Configurar servidor
	r := gin.Default()
	r.Use(cors.Default())

	// Definir rotas
	routes.PublicRoutes(r)

	// Iniciar servidor
	log.Println("🚀 Servidor rodando na porta", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
