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

func (r *ProductRepository) GetAll() ([]*models.Product, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, name, price, stock FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
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
		"INSERT INTO product (name, price, stock) VALUES ($1, $2, $3) RETURNING id",
		p.Name, p.Price, p.Stock,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	p.ID = id
	return p, nil
}

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(
		context.Background(),
		"SELECT id, name, price, stock FROM product WHERE id = $1",
		id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(p *models.Product) (*models.Product, error) {
	_, err := r.db.Exec(
		context.Background(),
		"UPDATE product SET name = $1, price = $2, stock = $3 WHERE id = $4",
		p.Name, p.Price, p.Stock, p.ID)
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
