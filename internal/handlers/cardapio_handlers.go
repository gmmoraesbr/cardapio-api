package handlers

import (
	"cardapio-api/internal/models"
	"cardapio-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura do Handler (Controller)
type CardapioHandler struct {
	Service *services.CardapioService
}

// Adicionando o método ListarCardapio corretamente
func (h *CardapioHandler) ListarCardapio(c *gin.Context) {
	itens, err := h.Service.BuscarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar itens"})
		return
	}
	c.JSON(http.StatusOK, itens)
}

// Adicionando o método ObterItemPorID corretamente
func (h *CardapioHandler) ObterItemPorID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.BuscarPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// Criar um novo Handler
func NovoCardapioHandler(service *services.CardapioService) *CardapioHandler {
	return &CardapioHandler{Service: service}
}

// ✅ Adicionar um novo item ao cardápio
func (h *CardapioHandler) AdicionarItem(c *gin.Context) {
	var novoItem models.Item
	if err := c.ShouldBindJSON(&novoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	item, err := h.Service.AdicionarItem(novoItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao inserir item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// ✅ Atualizar um item do cardápio
func (h *CardapioHandler) AtualizarItem(c *gin.Context) {
	id := c.Param("id")
	var itemAtualizado models.Item

	if err := c.ShouldBindJSON(&itemAtualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	if err := h.Service.AtualizarItem(id, itemAtualizado); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Item atualizado com sucesso"})
}

// ✅ Remover um item do cardápio
func (h *CardapioHandler) RemoverItem(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.RemoverItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Item removido com sucesso"})
}
