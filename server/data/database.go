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
	port := os.Getenv("POSTGRES_HOST_POST")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", user, psswrd, port, dbName)
	dbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Could not create connection pool to postgres... %s", err.Error())
	}

	return dbPool
}
