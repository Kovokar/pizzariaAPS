package controllers

import (
	"net/http"
	"pizza-api/dbs"
	"pizza-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateCliente(c *gin.Context) {
	var cliente models.Cliente

	// Verifica se os dados JSON foram enviados corretamente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Cria o cliente no banco de dados
	if err := dbs.GetDB().Create(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	// Retorna o cliente criado
	c.JSON(http.StatusCreated, cliente)
}

func GetClientes(c *gin.Context) {
	var clientes []models.Cliente

	if err := dbs.GetDB().Find(&clientes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(200, clientes)
}
