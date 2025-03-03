package repositories

import (
	"cardapio-api/internal/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface do repositório do cardápio
type CardapioRepository interface {
	BuscarTodos() ([]models.Item, error)
	BuscarPorID(id primitive.ObjectID) (*models.Item, error)
	Inserir(item models.Item) error
	Atualizar(id primitive.ObjectID, item models.Item) error
	Remover(id primitive.ObjectID) error
}

// Implementação do repositório usando MongoDB
type MongoCardapioRepository struct {
	collection *mongo.Collection
}

// Criar um novo repositório
func NovoCardapioRepository(db *mongo.Database) CardapioRepository {
	return &MongoCardapioRepository{
		collection: db.Collection("itens"),
	}
}

// Buscar todos os itens
func (r *MongoCardapioRepository) BuscarTodos() ([]models.Item, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var itens []models.Item
	if err = cursor.All(context.TODO(), &itens); err != nil {
		return nil, err
	}

	return itens, nil
}

// Buscar um item pelo ID
func (r *MongoCardapioRepository) BuscarPorID(id primitive.ObjectID) (*models.Item, error) {
	var item models.Item
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return nil, errors.New("item não encontrado")
	}
	return &item, nil
}

// Inserir um novo item
func (r *MongoCardapioRepository) Inserir(item models.Item) error {
	_, err := r.collection.InsertOne(context.TODO(), item)
	return err
}

// Atualizar um item
func (r *MongoCardapioRepository) Atualizar(id primitive.ObjectID, item models.Item) error {
	result, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": item})
	if err != nil || result.MatchedCount == 0 {
		return errors.New("erro ao atualizar item")
	}
	return nil
}

// Remover um item
func (r *MongoCardapioRepository) Remover(id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		return errors.New("erro ao remover item")
	}
	return nil
}
