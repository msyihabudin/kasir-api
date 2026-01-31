package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(connStr string) (*sql.DB, error) {
	// var err error
	// db, err = sql.Open("postgres", conn)
	// if err != nil {
	// 	return nil, err
	// }

	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }

	// db.SetMaxOpenConns(25)
	// db.SetMaxIdleConns(5)

	// log.Println("Database connected successfully")
	// return db, nil

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		return nil, err
	}

	log.Println("Connected to:", version)
	return db, nil
}
