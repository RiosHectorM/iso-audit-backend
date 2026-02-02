package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RiosHectorM/iso-audit-backend/internal/core/domain"
	"github.com/RiosHectorM/iso-audit-backend/internal/core/services"
)

type AuditHandler struct {
	service *services.AuditService
}

func NewAuditHandler(service *services.AuditService) *AuditHandler {
	return &AuditHandler{service: service}
}

func (h *AuditHandler) CreateAudit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Title  string `json:"title"`
		Norm   string `json:"norm"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	audit := &domain.Audit{
		Title:     req.Title,
		Norm:      req.Norm,
		Status:    req.Status,
		CreatedAt: time.Now(),
	}

	if err := h.service.CreateAudit(audit); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(audit)
}

func (h *AuditHandler) GetAuditsByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	audits, err := h.service.GetUserDashboard(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(audits)
}

func (h *AuditHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
