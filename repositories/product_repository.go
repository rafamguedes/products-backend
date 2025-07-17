package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/seuusuario/api-rest-go/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := `
		SELECT id, name, description, price, category, stock_quantity, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.StockQuantity,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT id, name, description, price, category, stock_quantity, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var product models.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
		&product.StockQuantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("produto com ID %d não encontrado", id)
		}
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Create(req models.CreateProductRequest) (*models.Product, error) {
	query := `
		INSERT INTO products (name, description, price, category, stock_quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, description, price, category, stock_quantity, created_at, updated_at
	`

	now := time.Now()
	row := r.db.QueryRow(query, req.Name, req.Description, req.Price, req.Category, req.StockQuantity, now, now)

	var product models.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
		&product.StockQuantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(id int, req models.UpdateProductRequest) (*models.Product, error) {
	existing, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.Price > 0 {
		existing.Price = req.Price
	}
	if req.Category != "" {
		existing.Category = req.Category
	}
	if req.StockQuantity >= 0 {
		existing.StockQuantity = req.StockQuantity
	}
	existing.UpdatedAt = time.Now()

	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, category = $4, stock_quantity = $5, updated_at = $6
		WHERE id = $7
		RETURNING id, name, description, price, category, stock_quantity, created_at, updated_at
	`

	row := r.db.QueryRow(query, existing.Name, existing.Description, existing.Price,
		existing.Category, existing.StockQuantity, existing.UpdatedAt, id)

	var product models.Product
	err = row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
		&product.StockQuantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto com ID %d não encontrado", id)
	}

	log.Printf("Produto com ID %d removido com sucesso", id)
	return nil
}

func (r *ProductRepository) GetByCategory(category string) ([]models.Product, error) {
	query := `
		SELECT id, name, description, price, category, stock_quantity, created_at, updated_at
		FROM products
		WHERE category = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.StockQuantity,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) FindByFilter(filter models.ProductFilter, nextToken models.NextTokenRequest) ([]models.Product, int, error) {
	baseQuery := `
		SELECT id, name, description, price, category, stock_quantity, created_at, updated_at
		FROM products
		WHERE 1=1
	`

	countQuery := `
		SELECT COUNT(*)
		FROM products
		WHERE 1=1
	`

	var args []interface{}
	var conditions string
	argIndex := 1

	// Build filter conditions
	if filter.Name != "" {
		conditions += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+filter.Name+"%")
		argIndex++
	}

	if filter.Category != "" {
		conditions += fmt.Sprintf(" AND category = $%d", argIndex)
		args = append(args, filter.Category)
		argIndex++
	}

	if filter.MinPrice != nil {
		conditions += fmt.Sprintf(" AND price >= $%d", argIndex)
		args = append(args, *filter.MinPrice)
		argIndex++
	}

	if filter.MaxPrice != nil {
		conditions += fmt.Sprintf(" AND price <= $%d", argIndex)
		args = append(args, *filter.MaxPrice)
		argIndex++
	}

	if filter.MinStock != nil {
		conditions += fmt.Sprintf(" AND stock_quantity >= $%d", argIndex)
		args = append(args, *filter.MinStock)
		argIndex++
	}

	if filter.MaxStock != nil {
		conditions += fmt.Sprintf(" AND stock_quantity <= $%d", argIndex)
		args = append(args, *filter.MaxStock)
		argIndex++
	}

	if nextToken.Row > 0 {
		if nextToken.Order == "asc" {
			conditions += fmt.Sprintf(" AND id > $%d", argIndex)
		} else {
			conditions += fmt.Sprintf(" AND id < $%d", argIndex)
		}
		args = append(args, nextToken.Row)
		argIndex++
	}

	// Get total count (create a copy of args without the limit parameter)
	var total int
	var countArgs []interface{}
	if len(args) > 0 {
		countArgs = make([]interface{}, len(args))
		copy(countArgs, args)
	}
	err := r.db.QueryRow(countQuery+conditions, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	orderClause := " ORDER BY id DESC"
	if nextToken.Order == "asc" {
		orderClause = " ORDER BY id ASC"
	}

	limit := nextToken.Limit
	if limit == 0 {
		limit = 10
	}

	limitClause := fmt.Sprintf(" LIMIT $%d", argIndex)
	args = append(args, limit+1)

	finalQuery := baseQuery + conditions + orderClause + limitClause
	rows, err := r.db.Query(finalQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.StockQuantity,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	return products, total, nil
}
