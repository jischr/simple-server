package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	version = "dev" // will be injected at build time
	port = "9000" // default port
)

func main() {
	server := echo.New()
	server.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, version)
	})
	server.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Worker is running\n")
	})

	// Start server in a goroutine
	go func() {
		log.Printf("Worker v%s starting at http://localhost:%s", version, port)
		if err := server.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting worker: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Shutting down worker...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Worker shutdown failed: %v", err)
	}

	log.Println("Worker shut down cleanly.")
}