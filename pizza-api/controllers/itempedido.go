package controllers

import (
	"net/http"
	"pizza-api/dbs"
	"pizza-api/models"

	"github.com/gin-gonic/gin"
)

func CreateItemPedido(c *gin.Context) {
	var itemPedido models.ItensPedido

	if err := c.ShouldBindJSON(&itemPedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Calcular o preço do item
	pizza := models.Pizza{}
	tamanho := models.Tamanho{}

	if err := dbs.GetDB().First(&pizza, itemPedido.PizzaID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro FRISTAR pizza"})
		return
	}

	if err := dbs.GetDB().First(&tamanho, itemPedido.TamanhoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro FRISTAR taanho"})
		return
	}

	itemPedido.PrecoItem = pizza.PrecoBase * tamanho.Modificador * float64(itemPedido.Quantidade)

	if err := dbs.GetDB().Create(&itemPedido).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar itemPedido"})
		return
	}

	c.JSON(http.StatusCreated, itemPedido)
}

func GetItensPedidos(c *gin.Context) {
	var itensPedidos []models.ItensPedido
	if err := dbs.GetDB().Find(&itensPedidos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar itensPedidos"})
		return
	}
	c.JSON(200, itensPedidos)
}
