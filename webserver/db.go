package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func initDB() error {
	_, err := db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS targets (
			id SERIAL PRIMARY KEY,
			address TEXT NOT NULL,
			alive BOOLEAN NOT NULL	
		)
	`)
	return err
}