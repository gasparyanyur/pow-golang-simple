// cmd/server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"wow/internal/pow"
	"wow/internal/repository"
	"wow/internal/resthttp/handler"
	"wow/internal/service"
	"wow/pkg/utils"

	"wow/internal/config"
)

func main() {

	cfgPath := os.Getenv("CONFIG")
	if cfgPath == "" {
		log.Fatal("$CONFIG must be set")
	}

	// Load configuration
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		util.Error("Failed to load config: " + err.Error())
		log.Fatal(err)
	}

	// Initialize repository, service, and handler
	repo := repository.NewQuoteRepository()
	powService := pow.NewProofOfWorkService(cfg.Target)
	quoteService := service.NewQuoteService(repo, powService)
	restHTTP := handler.NewHandler(quoteService)
	restHTTPRouteEngine := handler.NewRouter()

	// Setup the router and routes
	restHTTPRouteGroup := restHTTPRouteEngine.Group("/api/v1")
	restHTTPRouteGroup.POST("/quote", restHTTP.GetQuote)
	restHTTPRouteGroup.GET("/challenge", restHTTP.GetChallenge)

	server := &http.Server{
		Addr:    cfg.ServerAddress,   // The port to run the server on
		Handler: restHTTPRouteEngine, // The Gin router as the HTTP handler
	}
	// Create a channel to listen for OS signals (for graceful shutdown)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine so we can listen for termination signals
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting Gin server:", err)
		}
	}()
	// Wait for a termination signal (SIGINT or SIGTERM)
	sig := <-sigChan
	fmt.Println("Received signal:", sig)

	// Graceful shutdown - attempt to stop the server
	fmt.Println("Shutting down the server gracefully...")

	// Set a timeout for shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Call server.Shutdown() to shut down gracefully
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Error during shutdown:", err)
	} else {
		fmt.Println("Server shut down gracefully.")
	}
}
