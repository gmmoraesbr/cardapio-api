package services

import (
	"cardapio-api/internal/models"
	"cardapio-api/internal/repositories"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Estrutura do serviço do cardápio
type CardapioService struct {
	Repo repositories.CardapioRepository
}

// Criar um novo serviço
func NovoCardapioService(repo repositories.CardapioRepository) *CardapioService {
	return &CardapioService{Repo: repo}
}

// Buscar todos os itens do cardápio
func (s *CardapioService) BuscarTodos() ([]models.Item, error) {
	return s.Repo.BuscarTodos()
}

// Buscar um item por ID
func (s *CardapioService) BuscarPorID(id string) (*models.Item, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	return s.Repo.BuscarPorID(objID)
}

// Adicionar um novo item
func (s *CardapioService) AdicionarItem(item models.Item) (*models.Item, error) {
	item.ID = primitive.NewObjectID()
	err := s.Repo.Inserir(item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Atualizar um item existente
func (s *CardapioService) AtualizarItem(id string, itemAtualizado models.Item) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	return s.Repo.Atualizar(objID, itemAtualizado)
}

// Remover um item do cardápio
func (s *CardapioService) RemoverItem(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	return s.Repo.Remover(objID)
}
