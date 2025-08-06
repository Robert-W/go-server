package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/robert-w/go-server/internal/logger"
	"github.com/robert-w/go-server/internal/server"
)

func main() {
	logger.SetDefault()

	// Listen to these signals and create a context to use for a graceful shutdown
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	srv := server.New(ctx)
	// Run the server in a go routine so the main go routine can catch the
	// interrupt signals
	go func() {
		err := srv.Run()

		if err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "ServerError", err)
		}
	}()

	// Wait for a signal and start shutting things down
	<-ctx.Done()

	srv.Shutdown()
	slog.Info("Server shutting down")
}
