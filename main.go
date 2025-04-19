package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	version = "dev" // will be injected at build time
)

func main() {
	port := "9000" // default port
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	server := echo.New()
	server.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, version)
	})
	server.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running\n")
	})

	// Start server in a goroutine
	go func() {
		log.Printf("Server v%s starting at http://localhost:%s", version, port)
		if err := server.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server shut down cleanly.")
}
