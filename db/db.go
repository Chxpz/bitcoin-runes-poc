package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var Dbpool *pgxpool.Pool

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	Dbpool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS runes_poc (
		id SERIAL PRIMARY KEY,
		tx_hash VARCHAR(64) NOT NULL,
		etch_hash VARCHAR(64) NOT NULL,
		mint_hash VARCHAR(64) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := Dbpool.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Unable to create table: %v", err)
	}
}

func SaveTransaction(txHash, etchHash, mintHash string) error {
	query := `INSERT INTO runes_poc (tx_hash, etch_hash, mint_hash) VALUES ($1, $2, $3)`
	_, err := Dbpool.Exec(context.Background(), query, txHash, etchHash, mintHash)
	return err
}
