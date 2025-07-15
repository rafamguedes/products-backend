package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "admin")
	password := getEnv("DB_PASSWORD", "admin123")
	dbname := getEnv("DB_NAME", "products_db")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao verificar conexão com o banco de dados:", err)
	}

	log.Println("Conectado ao banco de dados PostgreSQL com sucesso!")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
