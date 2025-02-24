package main

import (
	"context"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/josuetorr/frequent-flyer/server/data"
)

func main() {
	// TODO: call correct env file depending on env
	godotenv.Load(".dev.env")

	ctx := context.Background()
	dbPool := data.Init(ctx)
	defer dbPool.Close()
}
