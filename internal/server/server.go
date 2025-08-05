package server

import (
	"log/slog"

	"github.com/robert-w/go-server/internal/logger"
)

func Run() error {
	logger.SetDefault()

	slog.Info("Lets start our server")

	return nil
}
