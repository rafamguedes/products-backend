package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func Connect() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		host := getEnv("PGHOST", "localhost")
		port := getEnv("PGPORT", "5432")
		user := getEnv("PGUSER", "admin")
		password := getEnv("PGPASSWORD", "admin123")
		dbname := getEnv("PGDATABASE", "products_db")

		dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
			host, port, user, password, dbname)
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("erro ao conectar com o banco de dados: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("erro ao verificar conexão com o banco de dados: %w", err)
	}

	log.Println("Conectado ao banco de dados PostgreSQL com sucesso!")

	if err := RunMigrations(); err != nil {
		return fmt.Errorf("erro ao executar migrations: %w", err)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func RunMigrations() error {
	migrationsDir, err := filepath.Abs("migrations")
	if err != nil {
		return fmt.Errorf("erro ao localizar diretório de migrations: %w", err)
	}

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("diretório de migrations não encontrado: %s", migrationsDir)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("erro ao configurar dialect: %w", err)
	}

	if err := goose.Up(DB, migrationsDir); err != nil {
		return fmt.Errorf("erro ao executar migrations: %w", err)
	}

	log.Println("Migrations executadas com sucesso!")
	return nil
}
