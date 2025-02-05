package models

type Tamanho struct {
    ID         uint    `json:"id"`
    Nome       string  `json:"nome"`
    Modificador float64 `json:"modificador"` // Ex: 1.5 para pizza média
}
