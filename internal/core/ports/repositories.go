package ports

import "github.com/RiosHectorM/iso-audit-backend/internal/core/domain"

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	Save(user *domain.User) error
}

type AuditRepository interface {
	// Para el panel principal: trae auditorías filtradas por permiso
	GetByUserID(userID string) ([]domain.Audit, error)

	// Para la Auditora Líder: crear y asignar
	Create(audit *domain.Audit) error
	AssignUser(assignment domain.Assignment) error

	// Para el Auditor de Campo: ver el detalle de su sector
	GetByID(id string) (*domain.Audit, error)
}