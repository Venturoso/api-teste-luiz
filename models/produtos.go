package models

type Produtos struct {
	ID        int     `json:"ID"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float32 `json:"valor"`
	Ativo     string  `json:"ativo"`
}
