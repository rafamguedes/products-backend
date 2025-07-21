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

// @title Products Backend API Golang
// @version 1.0.0
// @description API para gerenciamento de produtos usando Gin e Swagger

// @host products-backend-production-a43e.up.railway.app
// @schemes https
func main() {
	database.Connect()
	defer database.DB.Close()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	productRepo := repositories.NewProductRepository(database.DB)

	productHandler := handlers.NewProductHandler(productRepo)

	routes.SetupRoutes(router, productHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando na porta %s", port)
	log.Printf("Acesse http://localhost:%s/health para verificar o status", port)
	log.Printf("API disponível em http://localhost:%s/api/v1/products", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
