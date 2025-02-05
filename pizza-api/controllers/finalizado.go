package controllers

import (
	"net/http"
	"pizza-api/dbs"
	"pizza-api/models"

	"github.com/gin-gonic/gin"
)

func CreateItemFinalizado(c *gin.Context) {
	var filalizado models.Finalizado

	if err := c.ShouldBindJSON(&filalizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Calcular o preço do item
	pizza := models.Pizza{}
	tamanho := models.Tamanho{}

	if err := dbs.GetDB().First(&pizza, filalizado.PizzaID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro FRISTAR pizza"})
		return
	}

	if err := dbs.GetDB().First(&tamanho, filalizado.TamanhoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro FRISTAR tamanho"})
		return
	}

	filalizado.PrecoItem = pizza.PrecoBase * tamanho.Modificador * float64(filalizado.Quantidade)

	if err := dbs.GetDB().Create(&filalizado).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar filalizado"})
		return
	}

	c.JSON(http.StatusCreated, filalizado)
}

func GetItensFinalizado(c *gin.Context) {
	var itensFinalizados []models.Finalizado
	if err := dbs.GetDB().Find(&itensFinalizados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar itensFinalizados"})
		return
	}
	c.JSON(200, itensFinalizados)
}
