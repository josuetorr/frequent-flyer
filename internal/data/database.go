package data

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool = pgxpool.Pool

func Init(ctx context.Context) *DBPool {
	user := os.Getenv("POSTGRES_USER")
	psswrd := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_HOST_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", user, psswrd, port, dbName)

	dbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Could not create connection pool to postgres... %s", err.Error())
	}

	var exists bool
	err = dbPool.QueryRow(ctx, fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)).Scan(&exists)
	if err != nil {
		log.Fatal("Query failed:", err)
	}

	if !exists {
		fmt.Println("Database does not exist. Creating...")
		_, err := dbPool.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			log.Fatal("Failed to create database:", err)
		}
		fmt.Println("Database created successfully!")
	}

	return dbPool
}
