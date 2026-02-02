package services

import (
	"errors"

	"github.com/RiosHectorM/iso-audit-backend/internal/core/domain"
	"github.com/RiosHectorM/iso-audit-backend/internal/core/ports"
)

var (
	ErrUnauthorizedAction = errors.New("user not authorized for this action")
	ErrAuditNotFound      = errors.New("audit not found")
)

type AuditService struct {
	repo ports.AuditRepository
}

// NewAuditService es nuestro constructor para la inyección de dependencias
func NewAuditService(repo ports.AuditRepository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

// GetUserDashboard trae las auditorías según el rol del usuario
func (s *AuditService) GetUserDashboard(userID string) ([]domain.Audit, error) {
	// Aquí podríamos agregar lógica de negocio extra antes de llamar al repo
	return s.repo.GetByUserID(userID)
}

// AssignStaff permite a un ADMIN asignar personal a un sector
func (s *AuditService) AssignStaff(adminID string, assignment domain.Assignment) error {
	// 1. Validar que quien asigna sea ADMIN (Simulación, esto vendría del contexto/token)
	// 2. Ejecutar la asignación
	return s.repo.AssignUser(assignment)
}

// CreateAudit permite crear una nueva auditoría (para testear)
func (s *AuditService) CreateAudit(audit *domain.Audit) error {
	return s.repo.Create(audit)
}
