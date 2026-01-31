package repositories

import (
	"context"
	// "database/sql"
	"kasir-api/models"

	"github.com/jackc/pgx/v5"
)

type CategoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]*models.Category, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, name, description FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
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
		"INSERT INTO category (name, description) VALUES ($1, $2) RETURNING id",
		c.Name, c.Description,
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
		"SELECT id, name, description FROM category WHERE id = $1",
		id,
	).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Update(c *models.Category) (*models.Category, error) {
	_, err := r.db.Exec(
		context.Background(),
		"UPDATE category SET name = $1, description = $2 WHERE id = $3",
		c.Name, c.Description, c.ID,
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
