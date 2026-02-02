package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RiosHectorM/iso-audit-backend/internal/adapters/storage/postgres/db"
	"github.com/RiosHectorM/iso-audit-backend/internal/core/domain"
	"github.com/RiosHectorM/iso-audit-backend/internal/core/ports"
	"github.com/google/uuid"
)

type AuditRepository struct {
	q *db.Queries
}

func NewAuditRepository(conn *sql.DB) *AuditRepository {
	return &AuditRepository{
		q: db.New(conn),
	}
}

// Ensure implementation
var _ ports.AuditRepository = (*AuditRepository)(nil)

func (r *AuditRepository) GetByUserID(userID string) ([]domain.Audit, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	// Use context.Background() for now, ideally passed from service
	audits, err := r.q.GetAuditsByUserID(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	var result []domain.Audit
	for _, a := range audits {
		result = append(result, domain.Audit{
			ID:        a.ID.String(),
			Title:     a.Title,
			Norm:      a.Norm,
			Status:    a.Status,
			CreatedAt: a.CreatedAt,
		})
	}
	return result, nil
}

func (r *AuditRepository) Create(audit *domain.Audit) error {
	var id uuid.UUID
	var err error

	if audit.ID != "" {
		id, err = uuid.Parse(audit.ID)
		if err != nil {
			return fmt.Errorf("invalid audit id: %w", err)
		}
	} else {
		id = uuid.New()
	}

	newAudit, err := r.q.CreateAudit(context.Background(), db.CreateAuditParams{
		ID:        id,
		Title:     audit.Title,
		Norm:      audit.Norm,
		Status:    audit.Status,
		CreatedAt: audit.CreatedAt,
	})
	if err != nil {
		return err
	}

	audit.ID = newAudit.ID.String()
	audit.CreatedAt = newAudit.CreatedAt
	return nil
}

func (r *AuditRepository) AssignUser(assignment domain.Assignment) error {
	uid, err := uuid.Parse(assignment.UserID)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}
	aid, err := uuid.Parse(assignment.AuditID)
	if err != nil {
		return fmt.Errorf("invalid audit id: %w", err)
	}

	return r.q.AssignUserToAudit(context.Background(), db.AssignUserToAuditParams{
		UserID:   uid,
		AuditID:  aid,
		SectorID: assignment.SectorID,
	})
}

func (r *AuditRepository) GetByID(id string) (*domain.Audit, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid audit id: %w", err)
	}

	a, err := r.q.GetAuditByID(context.Background(), uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &domain.Audit{
		ID:        a.ID.String(),
		Title:     a.Title,
		Norm:      a.Norm,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
	}, nil
}
