package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

	// Cria a tabela se não existir
	if err := createTables(); err != nil {
		return fmt.Errorf("erro ao criar tabelas: %w", err)
	}

	// Insere dados iniciais
	if err := seedInitialData(); err != nil {
		return fmt.Errorf("erro ao inserir dados iniciais: %w", err)
	}

	return nil
}

func createTables() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price DECIMAL(10,2) NOT NULL,
			category VARCHAR(100),
			stock_quantity INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func seedInitialData() error {
	// Verifica se a tabela já tem dados
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		// A tabela provavelmente não existe ou há outro erro
		return nil
	}

	if count > 0 {
		// Já existem dados, não precisa inserir novamente
		return nil
	}

	// Insere os dados iniciais
	_, err = DB.Exec(`
		INSERT INTO products (name, description, price, category, stock_quantity) VALUES
		('Smartphone Samsung Galaxy S23', 'Smartphone Android com 256GB de armazenamento', 2999.99, 'Eletrônicos', 50),
		('Notebook Dell XPS 13', 'Laptop ultrafino com processador Intel i7', 4999.99, 'Computadores', 25),
		('Fone de Ouvido Sony WH-1000XM4', 'Fone com cancelamento de ruído ativo', 799.99, 'Áudio', 100),
		('Tablet Apple iPad Air', 'Tablet com tela de 10.9 polegadas', 2499.99, 'Eletrônicos', 30),
		('Teclado Mecânico Logitech MX Keys', 'Teclado sem fio para produtividade', 349.99, 'Acessórios', 75)
	`)
	return err
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
