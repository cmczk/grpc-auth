package main

import (
	"log/slog"

	"github.com/cmczk/grpc-auth/sso/internal/config"
	"github.com/cmczk/grpc-auth/sso/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Setup(cfg.Env)

	log.Info(
		"starting app",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPC.Port),
	)
	log.Debug("debug messages are enabled")

	// init app
	// run grpc-server
}
