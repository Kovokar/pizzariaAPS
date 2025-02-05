package controllers

import (
	"net/http"
	"pizza-api/dbs"
	"pizza-api/models"

	"github.com/gin-gonic/gin"
)

func CreatePizza(c *gin.Context) {
	var pizza models.Pizza

	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := dbs.GetDB().Create(&pizza).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, pizza)
}

func GetPizzas(c *gin.Context) {
	var pizzas []models.Pizza
    if err := dbs.GetDB().Find(&pizzas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar pizzas"})
		return
	}
	c.JSON(200, pizzas)
}
