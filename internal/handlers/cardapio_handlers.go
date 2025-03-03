package handlers

import (
	"cardapio-api/internal/models"
	"cardapio-api/internal/services"
	"cardapio-api/pkg/logger"
	"cardapio-api/pkg/response"

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
		logger.Error("Erro ao buscar itens do cardápio: " + err.Error())
		response.Error(c, "Erro ao buscar itens", err)
		return
	}
	response.Success(c, "Itens carregados com sucesso", itens)
}

// Adicionando o método ObterItemPorID corretamente
func (h *CardapioHandler) ObterItemPorID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.Service.BuscarPorID(id)
	if err != nil {
		logger.Error("Erro ao buscar item por ID " + id + ": " + err.Error())
		response.Error(c, "Item não encontrado", err)
		return
	}
	response.Success(c, "Item encontrado", item)
}

// Criar um novo Handler
func NovoCardapioHandler(service *services.CardapioService) *CardapioHandler {
	return &CardapioHandler{Service: service}
}

// ✅ Adicionar um novo item ao cardápio
func (h *CardapioHandler) AdicionarItem(c *gin.Context) {
	var novoItem models.Item
	if err := c.ShouldBindJSON(&novoItem); err != nil {
		logger.Error("Erro ao decodificar JSON: " + err.Error())
		response.Error(c, "Dados inválidos", err)
		return
	}

	item, err := h.Service.AdicionarItem(novoItem)
	if err != nil {
		logger.Error("Erro ao inserir item no banco: " + err.Error())
		response.Error(c, "Erro ao inserir item", err)
		return
	}

	response.Success(c, "Item adicionado com sucesso", item)
}

// ✅ Atualizar um item do cardápio
func (h *CardapioHandler) AtualizarItem(c *gin.Context) {
	id := c.Param("id")
	var itemAtualizado models.Item

	if err := c.ShouldBindJSON(&itemAtualizado); err != nil {
		logger.Error("Erro ao decodificar JSON para atualização: " + err.Error())
		response.Error(c, "Dados inválidos", err)
		return
	}

	if err := h.Service.AtualizarItem(id, itemAtualizado); err != nil {
		logger.Error("Erro ao atualizar item " + id + ": " + err.Error())
		response.Error(c, "Erro ao atualizar item", err)
		return
	}

	response.Success(c, "Item atualizado com sucesso", nil)
}

// ✅ Remover um item do cardápio
func (h *CardapioHandler) RemoverItem(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.RemoverItem(id); err != nil {
		logger.Error("Erro ao remover item " + id + ": " + err.Error())
		response.Error(c, "Erro ao remover item", err)
		return
	}

	response.Success(c, "Item removido com sucesso", nil)
}
