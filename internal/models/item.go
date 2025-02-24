package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Item representa um item no cardápio
type Item struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome      string             `json:"nome"`
	Preco     float64            `json:"preco"`
	Descricao string             `json:"descricao"`
}
