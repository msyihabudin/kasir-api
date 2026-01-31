package database

import (
	"context"
	// "database/sql"
	"log"

	// _ "github.com/lib/pq"
	"github.com/jackc/pgx/v5"
)

func InitDB(connStr string) (*pgx.Conn, error) {
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	return nil, err
	// }

	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }

	// db.SetMaxOpenConns(15)
	// db.SetMaxIdleConns(5)

	// log.Println("Connected to database")
	// return db, nil

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		return nil, err
	}

	log.Println("Connected to:", version)
	return conn, nil
}
