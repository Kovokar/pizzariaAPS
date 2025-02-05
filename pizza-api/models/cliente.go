package models

type Cliente struct {
    ID       uint   `json:"id"`
    Nome     string `json:"nome"`
    Endereco string `json:"endereco"`
    Telefone string `json:"telefone"`
    Bairro   string `json:"bairro"`
}

