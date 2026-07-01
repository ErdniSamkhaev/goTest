package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// глобальная переменная для пула
var db *pgxpool.Pool

func initDB() error {
	_, err := db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS links (
			id SERIAL PRIMARY KEY,
			short_code TEXT NOT NULL UNIQUE,
			original_url TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	return err
}
