package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		host := getEnv("PGHOST", "localhost")
		port := getEnv("PGPORT", "5432")
		user := getEnv("PGUSER", "admin")
		password := getEnv("PGPASSWORD", "admin123")
		dbname := getEnv("PGDATABASE", "products_db")

		dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			host, port, user, password, dbname)
	}

	if os.Getenv("RAILWAY_ENVIRONMENT") == "production" || os.Getenv("DATABASE_URL") != "" {
		if !strings.Contains(dbURL, "sslmode=") {
			dbURL += " sslmode=require"
		}
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
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
