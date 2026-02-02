package main

import (
	"log/slog"
	"net/http"
	"os"

	api "github.com/RiosHectorM/iso-audit-backend/internal/adapters/http"
	"github.com/RiosHectorM/iso-audit-backend/internal/adapters/storage/postgres"
	"github.com/RiosHectorM/iso-audit-backend/internal/core/services"
	"github.com/RiosHectorM/iso-audit-backend/internal/platform/config"
)

func main() {
	// 1. Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. Config
	cfg := config.Load()

	// 3. Database
	dbConn, err := postgres.NewConnection(cfg.DBDSN)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// 4. Adapters & Services (Dependency Injection)
	auditRepo := postgres.NewAuditRepository(dbConn)
	auditService := services.NewAuditService(auditRepo)
	auditHandler := api.NewAuditHandler(auditService)

	// 5. Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/health", auditHandler.Health)
	mux.HandleFunc("/audits", auditHandler.CreateAudit)
	mux.HandleFunc("/audits/user", auditHandler.GetAuditsByUser)

	slog.Info("🚀 API Audit ISO started", "port", cfg.Port, "env", cfg.Env)

	// 6. Start Server
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
