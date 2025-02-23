package main

import (
	"cardapio-api/config"
	"cardapio-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// Conectar ao MongoDB
	config.ConectarMongo()

	r := gin.Default()

	// Configurar rotas públicas e privadas
	routes.PublicRoutes(r)
	privateGroup := r.Group("/")
	routes.PrivateRoutes(privateGroup)

	r.Run(":8080")
}
