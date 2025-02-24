package main

import (
	"cardapio-api/config"
	"cardapio-api/internal/handlers"
	"cardapio-api/internal/repositories"
	"cardapio-api/internal/services"
	"cardapio-api/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Conectar ao MongoDB e garantir que a conexão foi feita
	db := config.ConectarMongo(cfg.MongoURI, cfg.Database)
	if db == nil {
		log.Fatal("❌ Erro: Banco de dados não inicializado!")
	}

	// Criar repositório, serviço e controller
	repo := repositories.NovoCardapioRepository(db)
	service := services.NovoCardapioService(repo)     // Criar o serviço
	cardapio := handlers.NovoCardapioHandler(service) // Criar o handler

	// Configurar servidor
	r := gin.Default()
	r.Use(cors.Default())

	// Adicione esta linha para configurar proxies de forma segura:
	r.SetTrustedProxies(nil) // Isso desativa a confiança em todos os proxies

	// Configurar rotas
	routes.SetupRoutes(r, cardapio)

	porta := cfg.ServerPort
	if porta == "" {
		porta = "8080" // Define um padrão caso a variável esteja vazia
	}

	log.Println("✅ Conectado ao MongoDB com sucesso!")
	log.Println("🚀 Servidor rodando na porta", porta)

	// Iniciar o servidor na porta correta
	err := r.Run(":" + porta)
	if err != nil {
		log.Fatal("❌ Erro ao iniciar servidor:", err)
	}
}
