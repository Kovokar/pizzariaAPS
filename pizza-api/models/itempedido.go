package models

type ItensPedido struct {
    ID        uint    `json:"id"`
    PedidoID  uint    `json:"pedido_id"`
    PizzaID   uint    `json:"pizza_id"`
    TamanhoID uint    `json:"tamanho_id"`
    Quantidade int    `json:"quantidade"`
    PrecoItem float64 `json:"preco_item"`
}
