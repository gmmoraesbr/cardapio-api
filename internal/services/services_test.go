package services_test

import (
	"cardapio-api/config"
	"cardapio-api/internal/models"
	"cardapio-api/internal/repositories"
	"cardapio-api/internal/services"
	"log"
	"os"
	"testing"
)

func TestNovoCardapioService(t *testing.T) {
	service := services.NovoCardapioService(nil)
	if service == nil {
		t.Errorf("Erro ao inicializar o serviço")
	}
}

func TestAdicionarItem(t *testing.T) {

	// Definir valores padrão para evitar erro ao conectar ao MongoDB
	if os.Getenv("MONGO_URI") == "" {
		os.Setenv("MONGO_URI", "")
	}
	if os.Getenv("DATABASE_NAME") == "" {
		os.Setenv("DATABASE_NAME", "")
	}

	// Carregar configurações
	cfg := config.LoadConfig()
	log.Println(cfg)
	// Conectar ao MongoDB e garantir que a conexão foi feita
	db := config.ConectarMongo(cfg.MongoURI, cfg.Database)
	if db == nil {
		log.Fatal("❌ Erro: Banco de dados não inicializado!")
	}

	repo := repositories.NovoCardapioRepository(db)
	if repo == nil {
		t.Fatalf("Erro ao criar repositório para teste")
	}

	// Criar serviço com repositório válido
	service := services.NovoCardapioService(repo)

	item := models.Item{Nome: "Hamburguer", Preco: 20.00, Descricao: "Hamburguer artesanal"}

	_, err := service.AdicionarItem(item)
	if err != nil {
		t.Errorf("Erro inesperado ao adicionar item: %v", err)
	}
}
