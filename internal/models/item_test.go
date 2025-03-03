package models_test

import (
	"cardapio-api/internal/models"
	"testing"
)

func TestItemModel(t *testing.T) {
	item := models.Item{
		Nome:      "Pizza",
		Preco:     25.50,
		Descricao: "Pizza de mussarela",
	}

	if item.Nome != "Pizza" || item.Preco <= 0 {
		t.Errorf("Falha ao criar modelo de item")
	}
}
