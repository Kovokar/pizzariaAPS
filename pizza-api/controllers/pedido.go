package controllers

import (
	"net/http"
	"pizza-api/dbs"
	"pizza-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePedido(c *gin.Context) {
	var pedido models.Pedido

	// Verifica se os dados JSON foram enviados corretamente
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Cria o pedido no banco de dados
	pedido.DataPedido = time.Now()
	if err := dbs.GetDB().Create(&pedido).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}

	// Retorna o pedido criado
	c.JSON(http.StatusCreated, pedido)
}

func GetPedidos(c *gin.Context) {
	var pedidos []models.Pedido

	if err := dbs.GetDB().Find(&pedidos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar pedidos"})
		return
	}
	c.JSON(200, pedidos)
}
