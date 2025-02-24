package models

// Usuario representa um usuário no sistema
type Usuario struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}
