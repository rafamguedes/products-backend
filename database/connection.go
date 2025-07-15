package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Em produção (Railway), as variáveis já estarão disponíveis no ambiente
	// Localmente, carregamos do .env se existir
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Não foi possível carregar o arquivo .env:", err)
		}
	}

	host := getEnv("PGHOST", getEnv("DB_HOST", "localhost"))
	port := getEnv("PGPORT", getEnv("DB_PORT", "5432"))
	user := getEnv("PGUSER", getEnv("DB_USER", "postgres"))
	password := getEnv("PGPASSWORD", getEnv("DB_PASSWORD", ""))
	dbname := getEnv("PGDATABASE", getEnv("DB_NAME", "products_db"))

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
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
