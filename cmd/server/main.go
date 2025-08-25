package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saurabhraut1212/kitchen_kirana/internal/config"
	"github.com/saurabhraut1212/kitchen_kirana/internal/db"
	"github.com/saurabhraut1212/kitchen_kirana/internal/routes"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()
	cfg := config.Load()

	// Connect to MongoDB
	client, database := db.Connect(cfg)
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}
	}()

	// Setup Gin
	router := gin.Default()
	routes.Register(router, database)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Run server in a goroutine
	go func() {
		log.Printf("Listening on :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Context with timeout to finish ongoing requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting gracefully")
}
