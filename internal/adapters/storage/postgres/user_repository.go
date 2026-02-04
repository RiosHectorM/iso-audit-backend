package postgres

import (
	"context"
	"database/sql"
	"github.com/RiosHectorM/iso-audit-backend/internal/adapters/storage/postgres/db"
	"github.com/RiosHectorM/iso-audit-backend/internal/core/domain"
	"github.com/google/uuid"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{queries: db.New(conn)}
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	u, err := r.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:       u.ID.String(),
		Email:    u.Email,
		Password: u.Password,
		Role:     domain.Role(u.Role),
	}, nil
}

// Nota: Necesitarás implementar Create para registrar tu primer usuario ADMIN