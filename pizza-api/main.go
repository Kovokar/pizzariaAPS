package main

import (
	"fmt"
	"log"
	"pizza-api/controllers"
	"pizza-api/dbs"
	"pizza-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados
	config := dbs.DBConfig{
		User:     "postgres",
		Password: "Girassol0@",
		DbName:   "gepeto",
		Host:     "localhost",
		Port:     "5432",
		SSLMode:  "disable",
	}

	dbs.InitDB(config)

	// AutoMigrate para criar as tabelas
	if err := dbs.GetDB().AutoMigrate(&models.Cliente{}, &models.Pizza{}, &models.Tamanho{}, &models.Pedido{}, &models.ItensPedido{}, &models.Finalizado{}); err != nil {
		log.Fatal("Erro ao migrar os modelos:", err)
	}

	fmt.Println("Tabelas criadas com sucesso!")

	// Inicializa o router Gin
	router := gin.Default()

	// 🔥 Habilita CORS
	router.Use(cors.Default())

	// Definição de rotas
	router.POST("/clientes", controllers.CreateCliente)
	router.GET("/clientes", controllers.GetClientes)
	router.POST("/pizzas", controllers.CreatePizza)
	router.GET("/pizzas", controllers.GetPizzas)
	router.POST("/tamanhos", controllers.CreateTamanho)
	router.GET("/tamanhos", controllers.GetTamanhos)
	router.POST("/pedidos", controllers.CreatePedido)
	router.GET("/pedidos", controllers.GetPedidos)

	router.POST("/finalizado", controllers.CreateItemFinalizado)
	router.GET("/finalizados", controllers.GetItensFinalizado)

	router.POST("/itenspedidos", controllers.CreateItemPedido)
	router.GET("/itenspedidos", controllers.GetItensPedidos)
	//------------------------------------------------
	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
