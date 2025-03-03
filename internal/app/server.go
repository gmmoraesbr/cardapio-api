package app

import (
	"cardapio-api/config"
	"cardapio-api/internal/handlers"
	"cardapio-api/internal/repositories"
	"cardapio-api/internal/services"
	"cardapio-api/pkg/logger"
	"cardapio-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer inicia o servidor do aplicativo, configurando:
// - configurações do aplicativo (config.LoadConfig())
// - conexão ao MongoDB (config.ConectarMongo)
// - repositório, serviço e controller
// - servidor (gin.Default())
// - rotas (routes.SetupRoutes)
//
// Ele também configura proxies de forma segura e define uma porta padrão
// caso a variável esteja vazia.
//
// Em seguida, ele inicia o servidor na porta correta.
//
// Se houver um erro ao iniciar o servidor, ele é registrado no log.
func StartServer() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Conectar ao MongoDB e garantir que a conexão foi feita
	db := config.ConectarMongo(cfg.MongoURI, cfg.Database)
	if db == nil {
		logger.Error("❌ Erro: Banco de dados não inicializado!")
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

	logger.Info("✅ Conectado ao MongoDB com sucesso!")
	logger.Info("🚀 Servidor rodando na porta" + porta)

	// Iniciar o servidor na porta correta
	err := r.Run(":" + porta)
	if err != nil {
		logger.Error("❌ Erro ao iniciar servidor:" + err.Error())
	}
}
