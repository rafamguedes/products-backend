package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/seuusuario/api-rest-go/database"
	"github.com/seuusuario/api-rest-go/handlers"
	"github.com/seuusuario/api-rest-go/repositories"
	"github.com/seuusuario/api-rest-go/routes"
)

func main() {
	database.Connect()
	defer database.DB.Close()

	// Configurar o ambiente
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar o router
	router := gin.Default()

	// Middleware para logging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Inicializar repositórios
	productRepo := repositories.NewProductRepository(database.DB)

	// Inicializar handlers
	productHandler := handlers.NewProductHandler(productRepo)

	// Configurar rotas
	routes.SetupRoutes(router, productHandler)

	// Obter porta do ambiente ou usar padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando na porta %s", port)
	log.Printf("Acesse http://localhost:%s/health para verificar o status", port)
	log.Printf("API disponível em http://localhost:%s/api/v1/products", port)

	// Iniciar o servidor
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
