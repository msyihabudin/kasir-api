package repositories

import (
	"context"
	// "database/sql"
	"github.com/jackc/pgx/v5"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]*models.Category, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *CategoryRepository) Create(c *models.Category) (*models.Category, error) {
	var id string
	err := r.db.QueryRow(
		context.Background(),
		"INSERT INTO category (name) VALUES ($1) RETURNING id",
		c.Name,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	c.ID = id
	return c, nil
}

func (r *CategoryRepository) GetByID(id string) (*models.Category, error) {
	var c models.Category
	err := r.db.QueryRow(
		context.Background(),
		"SELECT id, name FROM category WHERE id = $1",
		id,
	).Scan(&c.ID, &c.Name)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Update(c *models.Category) (*models.Category, error) {
	_, err := r.db.Exec(
		context.Background(),
		"UPDATE category SET name = $1 WHERE id = $2",
		c.Name, c.ID,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CategoryRepository) Delete(id string) error {
	_, err := r.db.Exec(
		context.Background(),
		"DELETE FROM category WHERE id = $1",
		id,
	)
	return err
}
