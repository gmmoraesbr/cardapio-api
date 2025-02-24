package repositories

import (
	"cardapio-api/config"
	"cardapio-api/internal/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface para o repositório
type CardapioRepository interface {
	BuscarTodos() ([]models.Item, error)
	BuscarPorID(id string) (models.Item, error)
}

// Implementação do repositório no MongoDB
type MongoCardapioRepository struct {
	collection *mongo.Collection
}

func NovoCardapioRepository(db *mongo.Database) CardapioRepository {
	return &MongoCardapioRepository{
		collection: db.Collection("itens"),
	}
}

// Busca todos os itens do cardápio
func (repo *MongoCardapioRepository) BuscarTodos() ([]models.Item, error) {
	cursor, err := repo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Erro ao buscar itens:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var itens []models.Item
	if err = cursor.All(context.TODO(), &itens); err != nil {
		log.Println("Erro ao processar dados:", err)
		return nil, err
	}

	return itens, nil
}

// Busca um item por ID
func (repo *MongoCardapioRepository) BuscarPorID(id string) (models.Item, error) {
	var item models.Item
	err := repo.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}
