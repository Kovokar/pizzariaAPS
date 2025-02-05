package models

type Finalizado struct {
	ID         uint    `json:id`
	PizzaID    uint    `json:"pizza_id"`
	TamanhoID  uint    `json:"tamanho_id"`
	Quantidade int     `json:"quantidade"`
	PrecoItem  float64 `json:"preco_item"`
	ClienteId  uint    `json:"cliente_id"`
}
