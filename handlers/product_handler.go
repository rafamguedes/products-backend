package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seuusuario/api-rest-go/models"
	"github.com/seuusuario/api-rest-go/repositories"
)

type ProductHandler struct {
	productRepo *repositories.ProductRepository
}

func NewProductHandler(productRepo *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepo: productRepo}
}

// GetProducts godoc
// @Summary Lista todos os produtos
// @Description Retorna todos os produtos cadastrados
// @Tags produtos
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar produtos",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": len(products),
	})
}

// GetProduct godoc
// @Summary Busca um produto por ID
// @Description Retorna um produto específico pelo ID
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	product, err := h.productRepo.GetByID(id)
	if err != nil {
		if err.Error() == "produto com ID "+idStr+" não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Produto não encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar produto",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// CreateProduct godoc
// @Summary Cria um novo produto
// @Description Adiciona um novo produto ao sistema
// @Tags produtos
// @Accept json
// @Produce json
// @Param product body models.CreateProductRequest true "Dados do produto"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	product, err := h.productRepo.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar produto",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Produto criado com sucesso",
		"data":    product,
	})
}

// UpdateProduct godoc
// @Summary Atualiza um produto existente
// @Description Atualiza os dados de um produto específico
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body models.UpdateProductRequest true "Dados do produto"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	product, err := h.productRepo.Update(id, req)
	if err != nil {
		if err.Error() == "produto com ID "+idStr+" não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Produto não encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao atualizar produto",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto atualizado com sucesso",
		"data":    product,
	})
}

// DeleteProduct godoc
// @Summary Remove um produto
// @Description Remove um produto específico do sistema
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = h.productRepo.Delete(id)
	if err != nil {
		if err.Error() == "produto com ID "+idStr+" não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Produto não encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao remover produto",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto removido com sucesso",
	})
}

// GetProductsByCategory godoc
// @Summary Lista produtos por categoria
// @Description Retorna todos os produtos de uma categoria específica
// @Tags produtos
// @Accept json
// @Produce json
// @Param category path string true "Categoria"
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products/category/{category} [get]
func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")

	products, err := h.productRepo.GetByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar produtos por categoria",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     products,
		"total":    len(products),
		"category": category,
	})
}
