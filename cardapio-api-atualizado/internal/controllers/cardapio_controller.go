package controllers

import (
	"cardapio-api/internal/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura do Controller
type CardapioController struct {
	Repo repositories.CardapioRepository
}

// Buscar todos os itens do cardápio
func (ctrl *CardapioController) ListarCardapio(c *gin.Context) {
	itens, err := ctrl.Repo.BuscarTodos()
	if err != nil {
		log.Println("Erro ao listar itens:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar itens"})
		return
	}
	c.JSON(http.StatusOK, itens)
}

// Buscar um item por ID
func (ctrl *CardapioController) ObterItemPorID(c *gin.Context) {
	id := c.Param("id")
	item, err := ctrl.Repo.BuscarPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Item não encontrado"})
		return
	}
	c.JSON(http.StatusOK, item)
}
