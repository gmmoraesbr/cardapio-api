package repositories_test

import (
	"cardapio-api/config"
	"cardapio-api/internal/repositories"
	"log"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBuscarPorID(t *testing.T) {

	// Definir valores padrão para evitar erro ao conectar ao MongoDB
	if os.Getenv("MONGO_URI") == "" {
		os.Setenv("MONGO_URI", "")
	}
	if os.Getenv("DATABASE_NAME") == "" {
		os.Setenv("DATABASE_NAME", "")
	}

	// Carregar configurações
	cfg := config.LoadConfig()

	// Conectar ao MongoDB e garantir que a conexão foi feita
	db := config.ConectarMongo(cfg.MongoURI, cfg.Database)
	if db == nil {
		log.Fatal("❌ Erro: Banco de dados não inicializado!")
	}

	repo := repositories.NovoCardapioRepository(db)

	// Criando um ObjectID diretamente
	id := primitive.NewObjectID()

	_, err := repo.BuscarPorID(id) // Agora passamos um ObjectID válido
	if err == nil {
		t.Errorf("Esperado erro ao buscar item inexistente")
	}
}
