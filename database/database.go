package database

import (
	"context"
	"log"

	// _ "github.com/lib/pq"
	"github.com/jackc/pgx/v5"
)

func InitDB(connStr string) (*pgx.Conn, error) {
	// var err error
	// db, err = sql.Open("postgres", connStr)
	// if err != nil {
	// 	return nil, err
	// }

	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }

	// db.SetMaxOpenConns(15)
	// db.SetMaxIdleConns(5)

	// var version string
	// if err := db.QueryRow("SELECT version()").Scan(&version); err != nil {
	// 	return nil, err
	// }

	// log.Println("Connected to:", version)
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
