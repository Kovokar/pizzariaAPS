package controllers

import (
	"net/http"
	"pizza-api/dbs"
	"pizza-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTamanho(c *gin.Context) {
	var tamanho models.Tamanho

	// Verifica se os dados JSON foram enviados corretamente
	if err := c.ShouldBindJSON(&tamanho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Cria o tamanho no banco de dados
	if err := dbs.GetDB().Create(&tamanho).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar tamanho"})
		return
	}

	// Retorna o tamanho criado
	c.JSON(http.StatusCreated, tamanho)
}

func GetTamanhos(c *gin.Context) {
	var tamanhos []models.Tamanho

	if err := dbs.GetDB().Find(&tamanhos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar cliente"})
		return
	}

	c.JSON(200, tamanhos)
}
