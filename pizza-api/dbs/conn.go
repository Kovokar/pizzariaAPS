// db.go
package dbs

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variável global para armazenar a conexão com o banco de dados
var DB *gorm.DB

// Configurações do banco de dados
type DBConfig struct {
	User     string 
	Password string
	DbName   string
	Host     string
	Port     string
	SSLMode  string
}

// Função para inicializar a conexão com o banco de dados
func InitDB(config DBConfig) {
	// Cria a string de conexão (DSN)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", 
		config.User, config.Password, config.DbName, config.Host, config.Port, config.SSLMode)

	// Conecta ao banco de dados
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
}

// Função para obter a instância da conexão com o banco de dados
func GetDB() *gorm.DB {
	return DB
}
