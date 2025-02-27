package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/server/routes"
)

func main() {
	// TODO: call correct env file depending on env
	godotenv.Load(".dev.env")
	log := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	slog.SetDefault(log)

	ctx := context.Background()
	dbPool := data.Init(ctx)
	defer dbPool.Close()

	r := routes.RegisterRoutes(dbPool)

	fs := http.FileServer(http.Dir("web/public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	port := "3000"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			_ = fmt.Errorf("Server error: %s", err)
		}
	}()

	fmt.Println("Server started on port:", port)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	fmt.Println("Server exited.")
}
