package models

type Pizza struct {
    ID        uint    `json:"id"`
    Sabor     string  `json:"sabor"`
    PrecoBase float64 `json:"preco_base"`
}
