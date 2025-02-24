package controllers

import (
	"cardapio-api/config"
	"cardapio-api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Listar todos os itens do cardápio
func ListarCardapio(c *gin.Context) {
	collection := config.GetCollection("itens")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar itens"})
		return
	}
	defer cursor.Close(context.TODO())

	var itens []models.Item
	if err = cursor.All(context.TODO(), &itens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao processar dados"})
		return
	}

	c.JSON(http.StatusOK, itens)
}

// Buscar um item do cardápio pelo ID
func ObterItemPorID(c *gin.Context) {
	collection := config.GetCollection("itens")

	// Capturar o ID passado na URL
	idParam := c.Param("id")

	// Converter o ID para ObjectID do MongoDB
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	// Buscar o item no banco
	var item models.Item
	filter := bson.M{"_id": id}
	err = collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Item não encontrado"})
		return
	}

	// Retornar o item encontrado
	c.JSON(http.StatusOK, item)
}

// Adicionar um novo item ao cardápio
func AdicionarItem(c *gin.Context) {
	collection := config.GetCollection("itens")
	var novoItem models.Item

	if err := c.ShouldBindJSON(&novoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	novoItem.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), novoItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao inserir item"})
		return
	}

	c.JSON(http.StatusCreated, novoItem)
}

// Atualizar um item existente do cardápio
func AtualizarItem(c *gin.Context) {
	collection := config.GetCollection("itens")
	var itemAtualizado models.Item

	if err := c.ShouldBindJSON(&itemAtualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": itemAtualizado}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Item atualizado com sucesso"})
}

// Remover um item do cardápio
func RemoverItem(c *gin.Context) {
	collection := config.GetCollection("itens")

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil || result.DeletedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao remover item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Item removido com sucesso"})
}
