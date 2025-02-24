package server

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/josuetorr/frequent-flyer/server/data"
)

func main() {
	ctx := context.Background()
	dbPool := data.Init(ctx)
	defer dbPool.Close()
}
