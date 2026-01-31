package repositories

import (
	"context"
	// "database/sql"
	"kasir-api/models"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]*models.ProductWithCategory, error) {
	rows, err := r.db.Query(
		context.Background(),
		"SELECT p.id, p.name, p.price, p.stock, c.name as category_name FROM product p LEFT JOIN category c ON p.category_id = c.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.ProductWithCategory
	for rows.Next() {
		var p models.ProductWithCategory
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryName); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *ProductRepository) Create(p *models.Product) (*models.Product, error) {
	var id string
	err := r.db.QueryRow(
		context.Background(),
		"INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Price, p.Stock, p.CategoryID,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	p.ID = id
	return p, nil
}

func (r *ProductRepository) GetByID(id string) (*models.ProductWithCategory, error) {
	var p models.ProductWithCategory
	err := r.db.QueryRow(
		context.Background(),
		"SELECT p.id, p.name, p.price, p.stock, c.name FROM product p LEFT JOIN category c ON p.category_id = c.id WHERE p.id = $1",
		id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryName)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(p *models.Product) (*models.Product, error) {
	_, err := r.db.Exec(
		context.Background(),
		"UPDATE product SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5",
		p.Name, p.Price, p.Stock, p.CategoryID, p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) Delete(id string) error {
	_, err := r.db.Exec(
		context.Background(),
		"DELETE FROM product WHERE id = $1", id)
	return err
}
