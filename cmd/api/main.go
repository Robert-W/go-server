package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	dotenv "github.com/joho/godotenv"
	"github.com/robert-w/go-server/internal/logger"
	"github.com/robert-w/go-server/internal/server"
)

func main() {
	// Ignore any error here, we just use this locally for convenience
	_ = dotenv.Load()

	logger.SetDefault()

	// Signals I want to catch for graceful shutdown
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	srv, err := server.New(ctx)
	if err != nil {
		slog.Error("Server failed creation", "ServerError", err)
		os.Exit(1)
	}

	// Run the server in a go routine so the main go routine can catch the
	// interrupt signals and so we dont block the main routine
	go func() {
		err := srv.Run()

		if err != nil && err != http.ErrServerClosed {
			// Force our server to start the shutdown process
			slog.Error("Server failed to start", "ServerError", err)
			stop()
		}
	}()

	// Wait for a signal and start shutting things down
	<-ctx.Done()

	srv.Shutdown()
	slog.Info("Server shutting down")
}
