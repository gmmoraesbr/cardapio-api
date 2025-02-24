package main

import (
	"cardapio-api/config"
	"cardapio-api/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env:", err)
	}

	// Conectar ao MongoDB
	config.ConectarMongo()

	r := gin.Default()

	// 🔥 Adicionar Middleware de CORS 🔥
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permitir qualquer origem (pode restringir a um domínio específico)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar rotas públicas e privadas
	routes.PublicRoutes(r)
	privateGroup := r.Group("/")
	routes.PrivateRoutes(privateGroup)

	// Rodar o servidor
	r.Run(":8080")
}
