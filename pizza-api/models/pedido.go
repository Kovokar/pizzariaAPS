package models

import "time"

type Pedido struct {
    ID           uint      `json:"id"`
    ClienteID    uint      `json:"cliente_id"`
    Status       string    `json:"status"` // "aberto", "finalizado", "cancelado"
    DataPedido   time.Time `json:"data_pedido"`
    ValorTotal   float64   `json:"valor_total"`
}
