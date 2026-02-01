package main

import (
	"log/slog"
	"os"

	"github.com/RiosHectorM/iso-audit-backend/internal/platform/config"
	"github.com/RiosHectorM/iso-audit-backend/internal/platform/storage/postgres"
)

func main() {
	// 1. Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. Configuración
	cfg := config.Load()

	// 3. Database
	db, err := postgres.NewConnection(cfg.DBDSN)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("🚀 API Audit ISO started", "port", cfg.Port, "env", cfg.Env)

	// Aquí iría el inicio del servidor HTTP (Gin/Echo)
}
