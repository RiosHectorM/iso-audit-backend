package domain

import "errors"

type Role string

const (
	RoleAdmin    Role = "ADMIN"    // Auditor Líder: Control total
	RoleField    Role = "FIELD"    // Auditor de Campo: Redacción técnica de hallazgos
	RoleObserver Role = "OBSERVER" // Colaborador: Solo fotos e info adicional
	RoleViewer   Role = "VIEWER"   // Cliente: Solo lectura de informes finales
)

var ErrInvalidRole = errors.New("invalid role")

type User struct {
	ID       string
	Email    string
	Password string
	Role     Role
}

type Assignment struct {
	UserID   string
	AuditID  string
	SectorID string
}
