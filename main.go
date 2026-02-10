package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mini-bank/config"
	"mini-bank/db"
)

func main() {
    // 1. Load Config
    cfg, err := config.LoadConfig(".env")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    // 2. Connect Database
    dbConfig := db.Config{
		DBDriver:      cfg.DBDriver,
		DBSource:      cfg.DBSource,
		ServerAddress: cfg.ServerAddress,
	}

	database, err := db.NewDB(dbConfig)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }
    defer database.Close()

    // 3. Setup Simple HTTP Server (Health Check)
    server := &http.Server{
        Addr:    cfg.ServerAddress,
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("🏦 Mini Bank API is Running!"))
        }),
    }

    // 4. Graceful Shutdown (Menutup server dengan rapi jika di-kill)
    go func() {
        log.Printf("🚀 Starting server on %s", cfg.ServerAddress)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server error: %v", err)
        }
    }()

    // Menunggu sinyal interrupt (Ctrl+C)
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exited")
}